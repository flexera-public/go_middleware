package middleware

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/labstack/echo"
)

// Header used to get/set request id
const HeaderKey = "X-Request-Id"

// RequestID middleware
func RequestID(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		id := c.Request().Header.Get(HeaderKey)
		if id == "" {
			id = uuid.NewUUID().String()
		}
		c.Set("RequestID", id)
		err := h(c)

		// Set header before handling error so that the 500 response contains it
		// (be kind to prod support...)
		c.Response().Header().Set(HeaderKey, id)

		return err
	}
}
