package middleware_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoMiddleware(t *testing.T) {
	RegisterFailHandler(Fail)
	config.DefaultReporterConfig.NoColor = true
	RunSpecs(t, "GoMiddleware Suite")
}
