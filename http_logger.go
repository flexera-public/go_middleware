package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// Make it possible to mock logger for tests
type Logger interface {
	Print(val ...interface{})
	Printf(format string, val ...interface{})
}

// Simple HTTP logger echo middleware
// Note: this middleware leverages echo's context to retrieve the response status and size
func HttpLogger(logger Logger) echo.Middleware {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) *echo.HTTPError {
			msg := fmt.Sprintf(`Processing GET "%s"`, c.Request.URL.String())
			originIp := c.Request.Header.Get("X-Forwarded-For")
			if originIp == "" {
				originIp = c.Request.Header.Get("X-Originating-IP")
			}
			if originIp != "" {
				msg += fmt.Sprintf(" (for %s)", originIp)
			}
			if reqId := c.Get("RequestID"); reqId != nil {
				msg += fmt.Sprintf(" - Request ID: %v", reqId)
			}
			logger.Print(msg)
			start := time.Now()
			err := h(c)
			if err != nil {
				return err
			}
			elapsed := time.Since(start)
			var status string
			var size int
			if resp := c.Response; resp != nil {
				status = http.StatusText(resp.Status())
				size = resp.Size()
			}
			logger.Printf(`Completed in %s | %s | %d bytes`, elapsed, status, size)
			return nil
		}
	}
}
