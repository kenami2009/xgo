package middleware

import (
	"fmt"
	"xgo/framework"
)

func Recovery() framework.ControllerHandler {
	fmt.Println("middleware", "recovery")
	return func(c *framework.XContext) error {
		defer func() {
			if err := recover(); err != nil {
				c.Json(500, err)
			}
		}()

		if err := c.Next(); err != nil {
			return err
		}

		return nil
	}
}
