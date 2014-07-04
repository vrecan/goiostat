package statsOutput_test

import (
	"github.com/CapillarySoftware/goiostat/diskStat"
	. "github.com/CapillarySoftware/goiostat/statsOutput"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "time"
)

var _ = Describe("StatsOutput", func() {
	var (
		stat = diskStat.ExtendedIoStats{}
	)
	var statsOutputChannel = make(chan diskStat.ExtendedIoStats, 10)

	BeforeEach(func() {
		stat = diskStat.ExtendedIoStats{
			Device:       "test",
			ReadsMerged:  0,
			WritesMerged: 0,
			Writes:       0,
			Reads:        0,
			SectorsRead:  0,
			SectorsWrite: 0,
			Arqsz:        0,
			AvgQueueSize: 0,
			Await:        0,
			RAwait:       0,
			WAwait:       0,
			Util:         0,
			Svctm:        0,
		}
	})

	Describe("Calling stats output function", func() {
		Context("With a single stat message on the queue", func() {
			statsOutputChannel <- stat
			It("should be a len of 0", func() {
				go Output(statsOutputChannel)
				close(statsOutputChannel)
				Expect(len(statsOutputChannel)).To(Equal(0))

			})
		})

		// Context("With fewer than 300 pages", func() {
		//     It("should be a short story", func() {
		//         Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
		//     })
		// })
	})
})
