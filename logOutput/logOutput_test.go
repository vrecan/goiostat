package logOutput_test

import (
	. "github.com/CapillarySoftware/goiostat/diskStat"
	. "github.com/CapillarySoftware/goiostat/logOutput"
	. "github.com/CapillarySoftware/goiostat/outputInterface"
	. "github.com/CapillarySoftware/goiostat/protocols"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"fmt"
)

func testInterface(output Output, stats *ExtendedIoStats) {
	err := output.SendStats(stats)
	Expect(err).ShouldNot(Equal(BeNil()))
}

var _ = Describe("Test LogOutput Interface", func() {

	It("basic interface test", func() {

		output := &LogOutput{PStdOut}
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

	It("Fail Protobuffers output test", func() {
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
		output := &LogOutput{PProtoBuffers}
		err := output.SendStats(&stats)
		
		Expect(err).ShouldNot(BeNil())
	})

	It("Basic json output test", func() {
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
		output := &LogOutput{PJson}
		err := output.SendStats(&stats)
		fmt.Println(err)
		
		Expect(err).Should(BeNil())
	})	
})
