package service

import (
	"context"
	"errors"
	"io"
	"log"
	"reflect"
	codec2 "rpc_test.com/codec"
	"rpc_test.com/protocol"
	"rpc_test.com/transport"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

type RpcService interface {
	// 注册的方法 传递一个struct 生成一个对象
	Register(rcvr interface{}, metaData map[string]string) error
	// 启动一个服务
	Server(net string, addr string) error
	// 提供关闭方法
	Close() error
}

// 方法参数和返回值的约定
type methodType struct {
	method    reflect.Method //函数
	ArgType   reflect.Type   //函数的参数
	ReplyType reflect.Type   //函数的返回值
}

// service 记录的保存 方便通过反射下次直接获取
type service struct {
	name    string                 //service的名称
	typ     reflect.Type           // 反射之后的类型
	recv    reflect.Value          // 反射之后的值
	methods map[string]*methodType //名称和方法的对应map
}

// 定义一个服务类型
type simpleServer struct {
	codec      codec2.Codec              // 使用的序列化方式
	trs        transport.ServerTransport // 启动的传输方式服务
	mutex      sync.Mutex                //锁
	serviceMap sync.Map                  //保存service里面的方法
	shutdown   bool                      // 是否关闭
	option     Option                    // 传递的其他参数
}

func NewSimpleServer(option Option) *simpleServer {
	service := new(simpleServer)
	service.option = option
	service.codec = codec2.GetCodec(option.SerializeType)
	return service
}

func (s *simpleServer) Register(rcvr interface{}, metaData map[string]string) error {
	// 进行反射 获取到service的名称和下面公开的函数 然后绑定到map里面
	tye := reflect.TypeOf(rcvr)
	//  定义一个service
	srv := new(service)
	srv.name = tye.Name()
	srv.recv = reflect.ValueOf(rcvr)
	srv.typ = tye
	// 获取所有的方法 并且保存
	srv.methods = allMethods(tye, true)
	if len(srv.methods) == 0 {
		// 如果对应的类型没有任何符合规则的方法，扫描对应的指针类型
		// 也是从net.rpc包里抄来的
		method := allMethods(reflect.PtrTo(srv.typ), false)
		errorStr := ""
		if len(method) != 0 {
			errorStr = "rpcx.Register: type " + tye.Name() + " has no exported methods of suitable type (hint: pass a pointer to value of that type)"
		} else {
			errorStr = "rpcx.Register: type " + tye.Name() + " has no exported methods of suitable type"
		}
		return errors.New(errorStr)
	}
	// 根据服务名字 添加数字
	if _, duplicate := s.serviceMap.LoadOrStore(tye.Name(), srv); duplicate {
		return errors.New("rpc: service already defined: " + tye.Name())
	}
	return nil
}

func (s *simpleServer) Server(net string, addr string) error {
	// 获取到transport 启动服务
	s.trs = transport.NewServerTransport(s.option.TransportType)
	// 启动服务
	err := s.trs.Listen(net, addr)
	if err != nil {
		log.Println(err)
		return err
	}
	// 监听服务
	for {
		accept, err := s.trs.Accept()
		if err != nil {
			log.Println(err)
			return err
		}
		// 启动协程 处理任务
		go s.serveTransport(accept)
	}
}
func (s *simpleServer) serveTransport(tr transport.Transport) {
	for {
		// 根据协议 获取数据
		request, err := protocol.DecodeMessage(s.option.ProtocolType, tr)

		if err != nil {
			if err == io.EOF {
				log.Printf("client has closed this connection: %s", tr.RemoteAddr().String())
			} else if strings.Contains(err.Error(), "use of closed network connection") {
				log.Printf("rpcx: connection %s is closed", tr.RemoteAddr().String())
			} else {
				log.Printf("rpcx: failed to read request: %v", err)
			}
			return
		}
		// 复制一个数据
		response := request.Clone()
		response.MessageType = protocol.MessageTypeResponse
		// 获取传递的服务名和方法名
		sname := request.ServiceName
		mname := request.MethodName
		// 通过方法名字 获取保存到load中的数据
		srvInterface, ok := s.serviceMap.Load(sname)
		if !ok {
			s.writeErrorResponse(response, tr, "can not find service")
			return
		}
		srv, ok := srvInterface.(*service)
		if !ok {
			s.writeErrorResponse(response, tr, "not *service type")
			return
		}

		mtype, ok := srv.methods[mname]
		if !ok {
			s.writeErrorResponse(response, tr, "can not find method")
			return
		}
		argv := newValue(mtype.ArgType)
		replyv := newValue(mtype.ReplyType)

		ctx := context.Background()
		//  反序列化传递过来的数据
		err = s.codec.Decode(request.Data, argv)

		var returns []reflect.Value
		if mtype.ArgType.Kind() != reflect.Ptr {
			returns = mtype.method.Func.Call([]reflect.Value{srv.recv,
				reflect.ValueOf(ctx),
				reflect.ValueOf(argv).Elem(),
				reflect.ValueOf(replyv)})
		} else {
			returns = mtype.method.Func.Call([]reflect.Value{srv.recv,
				reflect.ValueOf(ctx),
				reflect.ValueOf(argv),
				reflect.ValueOf(replyv)})
		}
		if len(returns) > 0 && returns[0].Interface() != nil {
			err = returns[0].Interface().(error)
			s.writeErrorResponse(response, tr, err.Error())
			return
		}
		// 将数据进行编码
		responseData, err := codec2.GetCodec(request.SerializeType).Encode(replyv)
		if err != nil {
			s.writeErrorResponse(response, tr, err.Error())
			return
		}
		// 拼接返回值
		response.StatusCode = protocol.StatusOK
		response.Data = responseData
		// 将数据再次写回去
		_, err = tr.Write(protocol.EncodeMessage(s.option.ProtocolType, response))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (s *simpleServer) writeErrorResponse(response *protocol.Message, w io.Writer, err string) {
	response.Error = err
	log.Println(response.Error)
	response.StatusCode = protocol.StatusError
	response.Data = response.Data[:0]
	_, _ = w.Write(protocol.EncodeMessage(s.option.ProtocolType, response))
}

func newValue(t reflect.Type) interface{} {
	if t.Kind() == reflect.Ptr {
		return reflect.New(t.Elem()).Interface()
	} else {
		return reflect.New(t).Interface()
	}
}
// 关闭服务器
func (s *simpleServer) Close() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.shutdown = true
	err := s.trs.Close()
	// 删除下面所有绑定的方法
	s.serviceMap.Range(func(key, value interface{}) bool {
		s.serviceMap.Delete(key)
		return true
	})
	return err
}

var typeOfError = reflect.TypeOf((*error)(nil)).Elem()
var typeOfContext = reflect.TypeOf((*context.Context)(nil)).Elem()

// 构造方法名字和service的绑定
func allMethods(t reflect.Type, reportErr bool) map[string]*methodType {
	methods := make(map[string]*methodType)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		mtype := method.Type //方法类型
		mname := method.Name //方法名称
		// 方法必须是可以导出的
		if method.PkgPath != "" {
			continue
		}
		// 需要有四个参数: receiver, Context, args, *reply.
		if mtype.NumIn() != 4 {
			if reportErr {
				log.Println("method", mname, "has wrong number of ins:", mtype.NumIn())
			}
			continue
		}
		// 第一个参数必须是context.Context
		ctxType := mtype.In(1)
		if !ctxType.Implements(typeOfContext) {
			if reportErr {
				log.Println("method", mname, " must use context.Context as the first parameter")
			}
			continue
		}

		// 第二个参数是arg
		argType := mtype.In(2)
		if !isExportedOrBuiltinType(argType) {
			if reportErr {
				log.Println(mname, "parameter type not exported:", argType)
			}
			continue
		}
		// 第三个参数是返回值，必须是指针类型的
		replyType := mtype.In(3)
		if replyType.Kind() != reflect.Ptr {
			if reportErr {
				log.Println("method", mname, "reply type not a pointer:", replyType)
			}
			continue
		}
		// 返回值的类型必须是可导出的
		if !isExportedOrBuiltinType(replyType) {
			if reportErr {
				log.Println("method", mname, "reply type not exported:", replyType)
			}
			continue
		}

		// 必须有一个返回值
		if mtype.NumOut() != 1 {
			if reportErr {
				log.Println("method", mname, "has wrong number of outs:", mtype.NumOut())
			}
			continue
		}
		// 返回值类型必须是error
		if returnType := mtype.Out(0); returnType != typeOfError {
			if reportErr {
				log.Println("method", mname, "returns", returnType.String(), "not error")
			}
			continue
		}
		methods[mname] = &methodType{method: method, ArgType: argType, ReplyType: replyType}
	}
	return methods
}

// Is this type exported or a builtin?
func isExportedOrBuiltinType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return isExported(t.Name()) || t.PkgPath() == ""
}

// Is this an exported - upper case - name?
func isExported(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}
