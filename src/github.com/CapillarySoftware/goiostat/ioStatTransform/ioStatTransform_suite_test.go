package ioStatTransform_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIoStatTransform(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IoStatTransform Suite")
}
