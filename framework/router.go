package framework

import (
	"errors"
	"net/http"
)

type ControllerHandlerChain []ControllerHandler

var routerError = errors.New("未知访问，路由不存在。")

var httpMethods = []string{
	http.MethodGet, http.MethodPost,
}
