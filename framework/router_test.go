package framework

import (
	"fmt"
	"testing"
)

func handlerTest(c *XContext) error {
	fmt.Println("111")
	return nil
}
func TestAddRouter(t *testing.T) {
	r := NewRouter()
	r.AddRouter("GET", "/test", handlerTest)
	_, err := r.FindHandlerFunc("GET", "/test")
	if err != nil {
		t.Fatal("未查询到路由")
	}
}
