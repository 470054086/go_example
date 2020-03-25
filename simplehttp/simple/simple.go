package simple

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Handler func(ctx *Context)

type HandleList []Handler

var (
	JSONBIND = &JsonBind{}
)

// 定义一个简单的http
type SimpleHttp struct {
	RouteGroup
	pool     sync.Pool
	Handlers []Handler
	router   map[string]HandleList
}

// context处理
type Context struct {
	Request *http.Request
	Response http.ResponseWriter
	http *SimpleHttp
}

// 获取bindjson的数据
func (c *Context) BindJson(d interface{}) error {
	return JSONBIND.bind(c.Request,d)
}

// 获取参数的方法
type Bind interface {
	Name() string//使用的什么方式
	Bind(*http.Request,interface{}) error //绑定的功能
}

// Json绑定方式
type JsonBind struct {
}
func (j *JsonBind) Name() string {
	return "json"
}
// bind 数据中心
func (j *JsonBind) bind(r *http.Request,d interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(d); err!= nil {
		return  err
	}
	return  nil
}

func Default() *SimpleHttp {
	s:= &SimpleHttp{
		Handlers:   []Handler{},
		router:     make(map[string]HandleList),
	}
	//放入pool中
	s.pool.New= func() interface{} {
		return &Context{
			http:    s,
		}
	}
	s.RouteGroup.http = s
	return s
}

func (s *SimpleHttp) addRouter(url string, h HandleList) {
	// todo 这里应该不存在锁的问题 因为都是统一启动
	s.router[url]  = h
}

// 启动服务
func (s *SimpleHttp) Run(port string) error  {
	return http.ListenAndServe(port,s)
}
// 实现ServeHttp 接口
func (s *SimpleHttp) ServeHTTP(w http.ResponseWriter , r *http.Request)   {
	//获取context
	context := s.pool.Get().(*Context)
	// 赋值为http默认的
	context.Response = w
	context.Request = r
	// 处理请求结果
	s.handlerRequest(context)
	//再次放入pool
	s.pool.Put(context)
}

func (s *SimpleHttp)  handlerRequest(ctx *Context) {
	url := ctx.Request.URL.Path
	if(url == "/favicon.ico") {
		return
	}
	if list,ok := s.router[url];!ok  {
		fmt.Println("获取不到路由了")
		return
	} else{
		for _,val := range list {
			// 解决第一个为空的问题
			val(ctx)
		}
	}
}
// 路由需要实现的方法
type IRouter interface {
	Use(...Handler) IRouter
	Get(string, ...Handler) IRouter
}

// 路由方法
type RouteGroup struct {
	Middleware []Handler   //防止中间件
	http     *SimpleHttp //http
	basePath string      //请求的路径
}
// 添加中间件
func (r *RouteGroup) Use(handler ...Handler) IRouter {
	r.Middleware = append(r.Middleware, handler...)
	return r
}
// get 请求
func (r *RouteGroup) Get(url string, handler ...Handler) IRouter {
	r.handler( url, handler)
	return r
}

/**
	处理多个请求
*/
func (r *RouteGroup) handler(url string, handler HandleList) IRouter {
	// 构建一个当前长度的slice
	handlerSlice := HandleList{}
	// 添加中间件的handler
	if r.Middleware  != nil {
		handlerSlice = append(handlerSlice, r.Middleware...)
	}
	// 添加当前传进来的handler
	handlerSlice = append(handlerSlice, handler...)
	// 将当前的handler和路由添加到路由处理器之中
	r.http.addRouter(url, handlerSlice)
	return r
}
