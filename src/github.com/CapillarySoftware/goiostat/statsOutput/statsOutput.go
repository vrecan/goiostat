package statsOutput

//statsOutput is a simple goroutine that uses the output interface and
//sends the stat to the interface given.

import (
	"fmt"
	"github.com/CapillarySoftware/goiostat/diskStat"
	"github.com/CapillarySoftware/goiostat/outputInterface"
)

//Output takes an input channel and sends the data to the output interface.
func Output(channel <-chan *diskStat.ExtendedIoStats, output outputInterface.Output) {
	for {
		stat := <-channel
		err := output.SendStats(stat)
		if nil != err {
			fmt.Println("Failed to send stat to selected output: ", err)
		}
	}
}
