package middleware_test

import (
	"github.com/labstack/echo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rightscale/go_middleware"
)

var _ = Describe("RequestID", func() {
	Describe("given a handler", func() {
		var handler echo.HandlerFunc
		var called bool

		BeforeEach(func() {
			handler = func(c *echo.Context) error {
				called = true
				return nil
			}
		})

		It("provides a middleware", func() {
			h := middleware.RequestID(handler)
			Ω(h).ShouldNot(BeNil())
		})

		It("calls the handler", func() {
			h := middleware.RequestID(handler)
			Ω(h).ShouldNot(BeNil())
			h(dummyContext())
			Ω(called).Should(BeTrue())
		})

		It("sets the request ID header if not set", func() {
			h := middleware.RequestID(handler)
			Ω(h).ShouldNot(BeNil())
			ctx := dummyContext()
			h(ctx)
			Ω(called).Should(BeTrue())
			header := ctx.Response().Header().Get("X-Request-ID")
			Ω(header).ShouldNot(BeEmpty())
		})

		It("reuses the request ID header if set", func() {
			h := middleware.RequestID(handler)
			Ω(h).ShouldNot(BeNil())
			ctx := dummyContext()
			ctx.Request().Header.Set("X-Request-Id", "foo")
			h(ctx)
			Ω(called).Should(BeTrue())
			header := ctx.Response().Header().Get("X-Request-ID")
			Ω(header).Should(Equal("foo"))
		})
	})

})
