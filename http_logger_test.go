package middleware_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/labstack/echo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rightscale/go_middleware"
)

var _ = Describe("HttpLogger", func() {

	Describe("with a valid logger", func() {
		var logger middleware.Logger

		BeforeEach(func() {
			logger = log.New(os.Stdout, "", 0)
		})

		It("provides a middleware", func() {
			m := middleware.HttpLogger(logger)
			Ω(m).ShouldNot(BeNil())
		})

		Describe("given a handler", func() {
			var handler echo.HandlerFunc
			var called bool

			BeforeEach(func() {
				handler = func(c *echo.Context) *echo.HTTPError {
					called = true
					return nil
				}
			})

			It("calls the handler", func() {
				m := middleware.HttpLogger(logger)
				Ω(m).ShouldNot(BeNil())
				h := m.(func(echo.HandlerFunc) echo.HandlerFunc)(handler)
				Ω(h).ShouldNot(BeNil())
				h(dummyContext())
				Ω(called).Should(BeTrue())
			})

			Describe("logging", func() {
				var out []string

				BeforeEach(func() {
					logger = testLogger(&out)
				})

				It("logs", func() {
					m := middleware.HttpLogger(logger)
					Ω(m).ShouldNot(BeNil())
					h := m.(func(echo.HandlerFunc) echo.HandlerFunc)(handler)
					Ω(h).ShouldNot(BeNil())
					h(dummyContext())
					Ω(called).Should(BeTrue())
					Ω(out).Should(HaveLen(2))
					Ω(out[0]).Should(MatchRegexp("^Processing"))
					Ω(out[1]).Should(MatchRegexp("^Completed"))
				})

			})

		})
	})

})

// Dummy logger that keeps logged messages
func testLogger(out *[]string) middleware.Logger {
	return &tLogger{out: out}
}

// Test logger
type tLogger struct {
	out *[]string
}

func (t *tLogger) Print(v ...interface{}) {
	*t.out = append(*t.out, fmt.Sprint(v...))
}

func (t *tLogger) Printf(f string, v ...interface{}) {
	*t.out = append(*t.out, fmt.Sprintf(f, v...))
}

// Create dummy echo.Context with request for tests
// Note: echo makes it impossible to initialize the context response :(
func dummyContext() *echo.Context {
	req, _ := http.NewRequest("POST", "http://example.com", strings.NewReader("foo"))
	resp := &echo.Response{Writer: httptest.NewRecorder()}
	return echo.NewContext(req, resp, echo.New())
}
