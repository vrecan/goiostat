package zmqOutput_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestZmqOutput(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ZmqOutput Suite")
}
