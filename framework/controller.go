package framework

import (
	"net/http"
)

type HandlerFunc func(c *XContext) error

func IndexController(c *XContext) error {
	c.Json(http.StatusOK, map[string]interface{}{
		"name": "hello",
		"address":"nihao",
	})

	return nil
}
