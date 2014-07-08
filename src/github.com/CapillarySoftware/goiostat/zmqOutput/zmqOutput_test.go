package zmqOutput_test

import (
	. "github.com/CapillarySoftware/goiostat/diskStat"
	. "github.com/CapillarySoftware/goiostat/zmqOutput"
	. "github.com/onsi/ginkgo"
	// . "github.com/CapillarySoftware/goiostat/protoStat"
	// "fmt"
	. "github.com/onsi/gomega"
)

var _ = Describe("ZmqOutput", func() {

	It("Testing basic send stats", func() {
		output := ZmqOutput{}
		output.Connect("ipc:///tmp/zmqOutput.ipc")
		// defer output.Close()
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
		err := output.SendStats(&stats)
		Expect(err).Should(BeNil())
	})

	It("Call sendStats without initializing socket", func() {
		output := ZmqOutput{}
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
		err := output.SendStats(&stats)
		Expect(err).ShouldNot(BeNil())
	})
})
