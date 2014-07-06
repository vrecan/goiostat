package goiostat_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoiostat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goiostat Suite")
}
