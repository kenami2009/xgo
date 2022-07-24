package framework

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

var jsonContentType = []string{"application/json; charset=utf-8"}

type XContext struct {
	index          int //中间件计数器
	handlers       ControllerHandlerChain
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	ctx            context.Context
	logger         *Logger
}

var _ context.Context = &XContext{}

func (x *XContext) Info(msg ...string) {
	if x.logger != nil {
		x.logger.Info(x, msg...)
	}
}

func (x *XContext) Next() error {
	x.index++
	if len(x.handlers) > x.index {
		if err := x.handlers[x.index](x); err != nil {
			return err
		}
	}
	return nil
}

func (x *XContext) SetHandlers(handlers ControllerHandlerChain) {
	x.handlers = handlers
}

func (x *XContext) BaseContext() context.Context {
	return x.ctx
}
func (x *XContext) Deadline() (time.Time, bool) {
	return x.BaseContext().Deadline()
}
func (x *XContext) Done() <-chan struct{} {
	return x.BaseContext().Done()
}
func (x *XContext) Value(key interface{}) interface{} {
	return x.ctx.Value(key)
}
func (x *XContext) Err() error {
	return x.ctx.Err()
}

//json输出

func (x *XContext) Json(httpStatus int, data interface{}) {
	//TODO
	r := x.ResponseWriter
	header := x.ResponseWriter.Header()

	header["Content-Type"] = jsonContentType
	r.WriteHeader(httpStatus)
	result, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	r.Write(result)

}
