package logOutput_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLogOutput(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LogOutput Suite")
}
