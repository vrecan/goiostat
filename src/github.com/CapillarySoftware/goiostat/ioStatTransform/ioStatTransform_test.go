package ioStatTransform_test

import (
	. "github.com/CapillarySoftware/goiostat/ioStatTransform"

	"fmt"
	. "github.com/CapillarySoftware/goiostat/diskStat"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type error interface {
	Error() string
}

var _ = Describe("IoStatTransform", func() {

	Describe("Unit tests", func() {
		It("validate partition regex doesn't get hit when using raid controllers", func() {
			device := "md-1"
			Expect(IsPartition(&device)).Should(BeFalse())
		})
		It("validate partition regeg against device", func() {
			device := "sda"
			Expect(IsPartition(&device)).Should(BeFalse())
		})  

		It("validate partition regeg against partition", func() {
			device := "sda1"
			Expect(IsPartition(&device)).Should(BeTrue())
		}) 
	})

	Describe("IntegrationTest", func() {

		var (
			goodLine              []string
			stats                 DiskStat
			err                   error
			statsTransformChannel chan *DiskStat
			statsOutputChannel    chan *ExtendedIoStats
			outStat               *ExtendedIoStats
		)

		BeforeEach(func() {
			goodLine = []string{"1", "2", "Device", "4", "5", "6", "7", "8", "9",
				"10", "11", "12", "13", "14"}
			stats, err = LineToStat(goodLine)
			statsTransformChannel = make(chan *DiskStat, 10)
			statsOutputChannel = make(chan *ExtendedIoStats, 10)

		})

		It("Go routine test", func() {
			expStat := ExtendedIoStats{
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
			nstat, _ := LineToStat(goodLine)
			defer close(statsTransformChannel)
			defer close(statsOutputChannel)

			go TransformStat(statsTransformChannel, statsOutputChannel)

			Expect(&stats).ShouldNot(Equal(BeNil()))
			Expect(&nstat).ShouldNot(Equal(BeNil()))
			statsTransformChannel <- &stats
			statsTransformChannel <- &nstat
			outStat = <-statsOutputChannel
			fmt.Println(outStat)
			Expect(outStat).Should(Equal(&expStat))

		})
	})

})
