package statsOutput

import (
	"github.com/CapillarySoftware/goiostat/diskStat"
	"github.com/CapillarySoftware/goiostat/outputInterface"
)

func Output(channel <-chan diskStat.ExtendedIoStats, output outputInterface.Output) {
	for {
		stat := <-channel
		output.SendStats(stat)
	}
}
