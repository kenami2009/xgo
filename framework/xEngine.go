package framework

/*
web服务引擎

*/
import (
	"errors"
	"net/http"
)

var noDefineHttpMethod = errors.New("未定义的Http方法")

// 定义XEngine类型，实现Handler接口

type XEngine struct {
	Routers     map[string]*Tree
	middlewares ControllerHandlerChain
}

func NewEngine() *XEngine {
	return &XEngine{
		Routers: map[string]*Tree{
			"GET":  NewTree(),
			"POST": NewTree(),
		},
	}
}

//请求入口函数
func (x *XEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//用户请求上下文
	ctx := &XContext{
		index:          -1,
		handlers:       x.middlewares,
		request:        r,
		responseWriter: w,
		logger:         NewLogger(),
	}
	ctx.logger.Info(ctx, "msg", "ee", "22", "33")
	_ = ctx

	//查找路由获取对应Handler处理请求
	handlers, err := x.FindRouterHandlers(ctx)

	if err != nil {
		//404
		ctx.Json(404, "Page Not Found.")
	}
	ctx.SetHandlers(handlers)
	if err := ctx.Next(); err != nil {
		ctx.Json(500, "Server Error.")
	}
}

//添加路由
func (x *XEngine) AddRouter(method, url string, chain ControllerHandlerChain) error {
	tree, ok := x.Routers[method]
	if !ok {
		return noDefineHttpMethod
	}
	err := tree.AddRouter(url, chain)
	if err != nil {
		return err
	}
	return nil
}

//查找路由
func (x *XEngine) FindRouterHandlers(ctx *XContext) (handlers ControllerHandlerChain, err error) {
	httpMethod := ctx.request.Method
	uri := ctx.request.RequestURI
	tree, ok := x.Routers[httpMethod]
	if !ok {
		err = noDefineHttpMethod
	}
	handlers = tree.FindHandlers(uri)
	return
}

// 注册中间件
func (x *XEngine) Use(middleware ControllerHandler) {
	x.middlewares = append(x.middlewares, middleware)
}

//注册路由（GET,POST）
func (x *XEngine) GET(url string, handlerFunc ControllerHandler) {
	handlers := append(x.middlewares, handlerFunc)
	x.AddRouter("GET", url, handlers)
}

func (x *XEngine) POST(url string, handlerFunc ControllerHandler) {
	handlers := append(x.middlewares, handlerFunc)
	x.AddRouter("POST", url, handlers)
}

//启动服务
func (x *XEngine) Run(addr string) error {
	return http.ListenAndServe(addr, x)
}

//关闭服务
func (x *XEngine) Shutdown() {
}
