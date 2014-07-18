package diskStat_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDiskStat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DiskStat Suite")
}
