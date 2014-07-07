package zmqOutput

import(
	"fmt"
	. "github.com/CapillarySoftware/goiostat/diskStat" 
	. "github.com/CapillarySoftware/goiostat/protoStat"
	"code.google.com/p/goprotobuf/proto"

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
		for _, stat := range stats {
			data, err := proto.Marshal(&stat)
			if(nil != err) {
				fmt.Println("Failed to marshal stat message : ", stat)
			}
			//just print the encoded data for now... soon this will actually send a queue
			fmt.Println(data)
		}
}