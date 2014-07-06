package systemCall_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSystemCall(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SystemCall Suite")
}
