package proto_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestProto(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Proto Suite")
}
