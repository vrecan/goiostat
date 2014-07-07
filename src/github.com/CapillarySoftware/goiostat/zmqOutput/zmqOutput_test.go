package zmqOutput_test

import (
	. "github.com/CapillarySoftware/goiostat/zmqOutput"
	. "github.com/CapillarySoftware/goiostat/diskStat"
	// . "github.com/CapillarySoftware/goiostat/outputInterface"
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("ZmqOutput", func() {

	It("Testing reflection", func(){		
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
	output.SendStats(stats)
	})
})
