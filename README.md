#Go echo Middleware

Set of commonly useful middlewares for the Go [echo](https://github.com/labstack/echo) web framework.

## RequestID

The RequestID middleware makes it possible to trace HTTP requests across multiple web services (a
"web transaction"). The idea consist of assigning a unique request id to a given web transaction.
This id is passed in the HTTP request header, each service that is part of the transaction must
forward the header to any other service it makes requests to. 
This middleware first checks whether the incoming request has a X-Request-ID and if not initializes
a new unique id. The request id is then set in the context and in the response headers.

Usage:
```
e := echo.New()
e.Use(middleware.RequestID) // Put that first so loggers can log request id (see below)
```

## HttpLogger

The HttpLogger middleware wraps an existing logger to log HTTP requests and their responses:
```
Processing GET "http://example.com/test"
Completed in 31ns | 200 OK | 100 bytes
```
HttpLogger can use any logger than implements the `Print` and `Printf` methods as defined by the go
`log` package:
```go
type Logger interface {
	func Print(v ...interface{})
	func Printf(format string, v ...interface{})
}
```
Example:
```go
e := echo.New()
l, _ := syslog.NewLogger(syslog.LOG_NOTICE|syslog.LOG_LOCAL0, 0)
e.Use(middleware.HttpLogger(logger)) // Log to syslog
```
If the RequestID middleware is loaded the HttpLogger also logs the value of the X-Request-ID header.
