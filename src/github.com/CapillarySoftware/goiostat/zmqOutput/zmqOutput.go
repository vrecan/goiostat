package zmqOutput

import(
	"fmt"
	. "github.com/CapillarySoftware/goiostat/diskStat" 
	. "github.com/CapillarySoftware/goiostat/protoStat"

)
type ZmqOutput struct {
}

func (l ZmqOutput) SendStats (eStat *ExtendedIoStats) {
	var(
		stats []ProtoStat
		err error
		)

		stats, err = GetProtoStats(eStat)
		if(nil != err) {
			fmt.Println(err)
		}
	    fmt.Println(stats)
}