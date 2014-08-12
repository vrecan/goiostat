package protoStat_test

import (
	. "github.com/CapillarySoftware/goiostat/diskStat"
	. "github.com/CapillarySoftware/goiostat/protoStat"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "fmt"
)

var _ = Describe("Conversions", func() {

	var _ = Describe("Validate Fields", func() {
		var (
			eStats ExtendedIoStats
			stats  *ProtoStats
			err    error
		)

		BeforeEach(func() {
			eStats = ExtendedIoStats{
				"Device",
				float64(0),
				float64(1),
				float64(2),
				float64(3),
				float64(4),
				float64(5),
				float64(6),
				float64(7),
				float64(8),
				float64(9),
				float64(10),
				float64(11),
				float64(12.1),
			}
			stats, err = GetProtoStat(&eStats)
			Expect(len(stats.Stats)).Should(Equal(13))
			Expect(err).Should(BeNil())
		})

		It("validate All Field Values for protoStat", func() {
			for index, protoStat := range stats.Stats {
				//validate we didn't lose percision
				if index == 12 {
					Expect(*protoStat.Value).Should(Equal(float64(12.1)))
				} else {
					Expect(*protoStat.Value).Should(Equal(float64(index)))
				}

			}
		})

		It("Validate FieldNames in right format", func() {
			Expect(*stats.Stats[0].Key).Should(Equal("Device_ReadsMerged"))
		})
	})
})
