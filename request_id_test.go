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
			handler = func(c *echo.Context) *echo.HTTPError {
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
			// Uncomment once it's possible to properly initialize
			// go contexts for testing, see https://github.com/labstack/echo/issues/51
			//h(dummyContext())
			//Ω(called).Should(BeTrue())
		})
	})

})
