package diskStat_test

import (
	. "github.com/CapillarySoftware/goiostat/diskStat"

	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type error interface {
	Error() string
}

var _ = Describe("DiskStat", func() {
	var (
		stats DiskStat
		err   error
		goodLine []string
	)

	Describe("Parse valid Stat line", func() {
		BeforeEach(func() {
			goodLine := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9",
				"10", "11", "12", "13", "14"}
			stats, err = LineToStat(goodLine)
		})

		It("Nil check", func() {
			Expect(err).Should(BeNil())
			Expect(stats).ShouldNot(BeNil())
			fmt.Println(err)
			fmt.Println(stats)
		})

		It("Validate Fields", func() {
			// Id int64
			Expect(stats.Id).Should(Equal(int64(1)))
			// PartId int64
			Expect(stats.PartId).Should(Equal(int64(2)))
			// Device string
			Expect(stats.Device).Should(Equal("3"))
			// ReadsCompleted int64
			Expect(stats.ReadsCompleted).Should(Equal(int64(4)))
			// ReadsMerged int64
			Expect(stats.ReadsMerged).Should(Equal(int64(5)))
			// SectorsRead uint64 sectors are converted to bytes which should be * 1024 / 2
			Expect(stats.SectorsRead).Should(Equal(uint64(3072)))
			// MillisReading int64
			Expect(stats.MillisReading).Should(Equal(int64(7)))
			// WritesCompleted int64
			Expect(stats.WritesCompleted).Should(Equal(int64(8)))
			// WritesMerged int64
			Expect(stats.WritesMerged).Should(Equal(int64(9)))
			// SectorsWrite uint64 ectors are converted to bytes which should be * 1024 / 2
			Expect(stats.SectorsWrite).Should(Equal(uint64(5120)))
			// MillisWriting int64
			Expect(stats.MillisWriting).Should(Equal(int64(11)))
			// IoInProgress int64
			Expect(stats.IoInProgress).Should(Equal(int64(12)))
			// MillisDoingIo int64
			Expect(stats.MillisDoingIo).Should(Equal(int64(13)))
			// WeightedMillisDoingIo int64
			Expect(stats.WeightedMillisDoingIo).Should(Equal(int64(14)))
		})

		It("Validate Extended Fields", func() {
			// RecordTime int64
			// Expect(stats.RecordTime).Should(Equal(int64))
			// IoTotal int64
			Expect(stats.IoTotal).Should(Equal(int64(stats.ReadsCompleted + stats.WritesCompleted)))
			// SectorsTotal uint64 (input was 6 + 10 for raw stat)
			Expect(stats.SectorsTotalRaw).Should(Equal(uint64(16)))
		})

	})

	Describe("Invalid Lines Parsing", func(){
		BeforeEach(func() {
			goodLine = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9",
				"10", "11", "12", "13", "14"}
		})

			It("Empty test for all fields", func (){

				for index, value := range goodLine {
					goodLine[index] = ""
					stats, err= LineToStat(goodLine)
					Expect(err).ShouldNot(Equal(BeNil()))
					goodLine[index] = value
				}
		})	
	})
})
