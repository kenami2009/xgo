package framework

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

var jsonContentType = []string{"application/json; charset=utf-8"}

type XContext struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	logger         *Logger
}

var _ context.Context = &XContext{}

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
	r := x.responseWriter
	header := x.responseWriter.Header()

	header["Content-Type"] = jsonContentType
	r.WriteHeader(httpStatus)
	result, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	r.Write(result)

}
