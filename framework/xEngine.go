package framework
/*
web服务引擎

 */
import (
	"net/http"
)

// 定义XEngine类型，实现Handler接口
type XEngine struct {
	Routers *XRouter
}

func NewEngine() *XEngine {
	return &XEngine{
		Routers: NewRouter(),
	}
}

//请求入口函数
func (x *XEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpMethod := r.Method
	uri := r.RequestURI
	//用户请求上下文
	ctx := &XContext{
		request: r,
		responseWriter: w,
	}
	_ = ctx

	//查找路由获取对应Handler处理请求
	handlerFunc, err := x.Routers.FindHandlerFunc(httpMethod,uri)
	if err != nil {
		//404
	}

	handlerFunc(ctx)

}
//注册路由（GET,POST）
func (x *XEngine) GET(url string, handlerFunc HandlerFunc) {
	x.Routers.AddRouter("GET", url, handlerFunc)
}
func (x *XEngine) POST(url string, handlerFunc HandlerFunc) {
	x.Routers.AddRouter("POST", url, handlerFunc)
}



//启动服务
func (x *XEngine) Run(addr string) error {
	return http.ListenAndServe(addr, x)
}

//关闭服务
func (x *XEngine) Shutdown() {
}
