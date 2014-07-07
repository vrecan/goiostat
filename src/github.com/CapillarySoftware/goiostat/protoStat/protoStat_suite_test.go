package protoStat_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestProtoStat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProtoStat Suite")
}
