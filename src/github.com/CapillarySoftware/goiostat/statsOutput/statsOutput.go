package statsOutput

import (
	"github.com/CapillarySoftware/goiostat/diskStat"
	"github.com/CapillarySoftware/goiostat/outputInterface"
	"fmt"
)

func Output(channel <-chan *diskStat.ExtendedIoStats, output outputInterface.Output) {
	for {
		stat := <-channel
		err := output.SendStats(stat)
		if(nil != err) {
			fmt.Println("Failed to send stat to selected output")
		}
	}
}
