package framework

import (
	"errors"
	"net/http"
	"strings"
)

type HandlerChain map[string]HandlerFunc

var routerError = errors.New("未知访问，路由不存在。")

var httpMethods = []string{
	http.MethodGet, http.MethodPost,
}

type XRouter struct {
	routers map[string]HandlerChain
}

//注册路由
func (r *XRouter) AddRouter(method, url string, h HandlerFunc) error {
	method = strings.ToUpper(method)
	methodRouter, ok := r.routers[method]
	if ok {
		methodRouter[url] = h
	} else {
		r.routers[method] = HandlerChain{
			url: h,
		}
	}
	return nil
}

//查找handler
func (r *XRouter) FindHandlerFunc(method, url string) (handlerFunc HandlerFunc, err error) {
	methodRouter, ok := r.routers[method]
	if !ok {
		err = routerError
		return
	} else {
		handlerFunc, ok = methodRouter[url]
		if !ok {
			err = routerError
			return
		}
	}
	return
}

func NewRouter() *XRouter {
	return &XRouter{
		routers: map[string]HandlerChain{
			"GET":  HandlerChain{},
			"POST": HandlerChain{},
		},
	}
}
