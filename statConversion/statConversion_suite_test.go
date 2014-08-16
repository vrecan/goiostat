package statConversion_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStatConversion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StatConversion Suite")
}
