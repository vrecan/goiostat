package zmqOutput_test

import (
	. "github.com/CapillarySoftware/goiostat/zmqOutput"
	. "github.com/CapillarySoftware/goiostat/diskStat"	
	. "github.com/onsi/ginkgo"
	// . "github.com/CapillarySoftware/goiostat/protoStat"
	// . "github.com/onsi/gomega"
)

var _ = Describe("ZmqOutput", func() {

	It("Testing basic send stats", func(){		
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
	output.SendStats(&stats)
	})
})
