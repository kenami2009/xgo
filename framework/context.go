package framework

import (
	"encoding/json"
	"net/http"
)

var jsonContentType = []string{"application/json; charset=utf-8"}

type XContext struct {
	request        *http.Request
	responseWriter http.ResponseWriter
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
