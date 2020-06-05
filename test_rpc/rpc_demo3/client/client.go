package client

import (
	"context"
	"errors"
	"io"
	"log"
	"rpc_test.com/codec"
	"rpc_test.com/protocol"
	"rpc_test.com/transport"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type RpcClient interface {
	// 定义调用方法
	// serviceMethod 服务名.方法名
	// arg 参数名称
	// replay 返回值
	Call(ctx context.Context, serviceMethod string, arg interface{}, replay interface{}) error
	Go(ctx context.Context, serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call
	Close() error
}

type Call struct {
	ServiceMethod string      // 服务名.方法名
	Args          interface{} // 参数
	Reply         interface{} // 返回值（指针类型）
	Error         error       // 错误信息
	Done          chan *Call  // 在调用结束时激活
}

// 定义客服端实例
type simpleClient struct {
	codec        codec.Codec        // 使用的编码方式
	tr           io.ReadWriteCloser //使用的传输服务
	pendingCalls sync.Map           // 保存缓存
	mutex        sync.Mutex         // 使用锁
	shutdown     bool               //是否关闭
	option       Option
	seq          uint64 //唯一编号
}

func NewRpcClient(network string, addr string, option Option) (RpcClient, error) {
	s := new(simpleClient)
	// 获取codec
	s.codec = codec.GetCodec(option.SerializeType)
	// 获取tr 进行链接
	tr := transport.NewTransport(option.TransportType)
	err := tr.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	s.tr = tr
	// 启动一个接受数据的协程
	go s.input()
	return s, nil
}

func (s *simpleClient) Call(ctx context.Context, serviceMethod string, args interface{}, reply interface{}) error {
	// 添加context
	seq := atomic.AddUint64(&s.seq, 1)
	ctx = context.WithValue(ctx, protocol.RequestSeqKey, seq)

	canFn := func() {}
	// 判断是否有过期时间
	if s.option.RequestTimeout != time.Duration(0) {
		ctx, canFn = context.WithTimeout(ctx, s.option.RequestTimeout)
		metaDataInterface := ctx.Value(protocol.MetaDataKey)
		var metaData map[string]string
		if metaDataInterface == nil {
			metaData = make(map[string]string)
		} else {
			metaData = metaDataInterface.(map[string]string)
		}
		metaData[protocol.RequestTimeoutKey] = s.option.RequestTimeout.String()
		ctx = context.WithValue(ctx, protocol.MetaDataKey, metaData)
	}

	done := make(chan *Call, 1)
	call := s.Go(ctx, serviceMethod, args, reply, done)
	select {
	// 如果有取消的操作 直接取消
	case <-ctx.Done():
		canFn()
		s.pendingCalls.Delete(seq)
		call.Error = errors.New("client request time out")
	// 等待信息返回
	case <-call.Done:
	}
	return call.Error
}

func (c *simpleClient) Go(ctx context.Context, serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call {
	call := new(Call)
	call.ServiceMethod = serviceMethod
	call.Args = args
	call.Reply = reply

	if done == nil {
		done = make(chan *Call, 10) // buffered.
	} else {
		if cap(done) == 0 {
			log.Panic("rpc: done channel is unbuffered")
		}
	}
	call.Done = done

	c.send(ctx, call)

	return call
}

func (c *simpleClient) send(ctx context.Context, call *Call) {
	seq := ctx.Value(protocol.RequestSeqKey).(uint64)
	// 存储到缓存中
	c.pendingCalls.Store(seq, call)

	request := protocol.NewMessage(c.option.ProtocolType)
	// 构造请求数据
	request.Seq = seq
	request.MessageType = protocol.MessageTypeRequest
	serviceMethod := strings.SplitN(call.ServiceMethod, ".", 2)
	request.ServiceName = serviceMethod[0]
	request.MethodName = serviceMethod[1]
	request.SerializeType = codec.MessagePack
	request.CompressType = protocol.CompressTypeNone
	if ctx.Value(protocol.MetaDataKey) != nil {
		request.MetaData = ctx.Value(protocol.MetaDataKey).(map[string]string)
	}
	// 数据进行序列化
	requestData, err := c.codec.Encode(call.Args)
	if err != nil {
		log.Println(err)
		c.pendingCalls.Delete(seq)
		call.Error = err
		call.done()
		return
	}
	request.Data = requestData
	// 数据进行协议包装
	data := protocol.EncodeMessage(c.option.ProtocolType, request)
	// 通过transport发送数据
	_, err = c.tr.Write(data)
	// 如果发生错误的话
	if err != nil {
		log.Println(err)
		c.pendingCalls.Delete(seq)
		call.Error = err
		call.done()
		return
	}
}

func (c *Call) done() {
	c.Done <- c
}

func (s *simpleClient) Close() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.shutdown = true

	s.pendingCalls.Range(func(key, value interface{}) bool {
		call, ok := value.(*Call)
		if ok {
			call.Error = errors.New("client closer")
			call.done()
		}
		s.pendingCalls.Delete(key)
		return true
	})
	return nil
}

func (s *simpleClient) input() {
	var err error
	var response *protocol.Message
	for err == nil {
		response, err = protocol.DecodeMessage(s.option.ProtocolType, s.tr)
		if err != nil {
			break
		}

		seq := response.Seq
		callInterface, _ := s.pendingCalls.Load(seq)
		call := callInterface.(*Call)
		s.pendingCalls.Delete(seq)

		switch {
		case call == nil:
			//请求已经被清理掉了，可能是已经超时了
		case response.Error != "":
			call.Error = errors.New(response.Error)
			call.done()
		default:
			err = s.codec.Decode(response.Data, call.Reply)
			if err != nil {
				call.Error = errors.New("reading body " + err.Error())
			}
			call.done()
		}
	}
}
