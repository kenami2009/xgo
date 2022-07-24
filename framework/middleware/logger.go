package middleware

import (
	"fmt"
	"net"
	"strings"
	"xgo/framework"
)

func Logger() framework.ControllerHandler {
	fmt.Println("middleware", "log")
	return func(c *framework.XContext) error {
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		if err := c.Next(); err != nil {
			return err
		}

		ip, _, err := net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr))
		if err != nil {
			ip = ""
		}

		c.Info(ip, c.Request.Method, path, raw)

		return nil
	}
}
