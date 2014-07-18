package statsOutput_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStatsOutput(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StatsOutput Suite")
}
