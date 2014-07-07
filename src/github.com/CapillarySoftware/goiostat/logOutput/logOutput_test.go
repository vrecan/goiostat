package logOutput_test

import (
	. "github.com/CapillarySoftware/goiostat/logOutput"
	. "github.com/CapillarySoftware/goiostat/diskStat"
	. "github.com/CapillarySoftware/goiostat/outputInterface"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testInterface(output Output, stats *ExtendedIoStats) {
	err := output.SendStats(stats)
	Expect(err).ShouldNot(Equal(BeNil()))
}

var _ = Describe("Test LogOutput Interface", func() {

	It("basic interface test", func(){


	output := LogOutput{}
	stats := ExtendedIoStats{
		"Device",
		float64(0),
		float64(0),
		float64(0),
		float64(0),
		float64(0),
		float64(0),
		float64(0),
		float64(0),
		float64(0),
		float64(0),
		float64(0),
		float64(0),
		float64(0),
	}
	testInterface(output, &stats)
	})
})
