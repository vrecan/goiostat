package zmqOutput_test

import (
	. "github.com/CapillarySoftware/goiostat/diskStat"
	. "github.com/CapillarySoftware/goiostat/zmqOutput"
	. "github.com/onsi/ginkgo"
	// . "github.com/CapillarySoftware/goiostat/protoStat"
	// "fmt"
	. "github.com/onsi/gomega"
	// zmq "github.com/pebbe/zmq3"
)

func sendStats(output ZmqOutput, eStat *ExtendedIoStats, sendCount int) {
	for i := 0; i <= sendCount; i++ {
		output.SendStats(eStat)
	}
}

var _ = Describe("ZmqOutput", func() {
	eStat := ExtendedIoStats{
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

	url := "ipc:///tmp/testOutput1.ipc"

	It("Testing basic send stats", func() {
		output := ZmqOutput{}
		output.Connect(url)
		defer output.Close()

		err := output.SendStats(&eStat)
		Expect(err).Should(BeNil())
	})

	It("Call sendStats without initializing socket", func() {
		output := ZmqOutput{}
		defer output.Close()
		err := output.SendStats(&eStat)
		Expect(err).ShouldNot(BeNil())
	})

	//this test validates zmq works but also sucks for an integration test
	// It("Send to recv socket and validate we get what we expect", func() {
	// 	output := ZmqOutput{}
	// 	defer output.Close()
	// 	output.Connect(url)

	// 	recv, err := zmq.NewSocket(zmq.PULL)
	// 	Expect(err).Should(BeNil())
	// 	defer recv.Close()

	// 	recv.Bind(url)
	// 	go sendStats(output, &eStat, 1)

	// 	for i := 0; i <= 12; i++ {
	// 		s, err := recv.RecvBytes(0)
	// 		fmt.Println("bytes: ", s)
	// 		Expect(err).Should(BeNil())
	// 	}
	// })
})
