package logOutput

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/CapillarySoftware/goiostat/diskStat"
	. "github.com/CapillarySoftware/goiostat/protocols"
	"github.com/dustin/go-humanize"

	"os"
)

type LogOutput struct {
	Proto Protocol
}

func (l *LogOutput) SendStats(stat *ExtendedIoStats) (err error) {

	switch l.Proto {
	case PStdOut:
		{
			stdOut(stat)
		}
	case PJson:
		{
			jsonOut(stat)
		}
	default:
		{
			err = errors.New("LogOutput doesn't support the type given... ")
			return
		}
	}
	return
}

func stdOut(stat *ExtendedIoStats) (err error) {
	fmt.Printf("%s:  rrqm/s %.2f wrqm/s %.2f r/s %.2f w/s %.2f rsize/s %s wsize/s %s avgrq-sz %.2f avgqu-sz %.2f, await %.2f r_await %.2f w_await %.2f svctm %.2f util %.2f%% \n\n",
		stat.Device, stat.ReadsMerged, stat.WritesMerged, stat.Reads, stat.Writes, humanize.Bytes(uint64(stat.SectorsRead)),
		humanize.Bytes(uint64(stat.SectorsWrite)), stat.Arqsz, stat.AvgQueueSize, stat.Await, stat.RAwait, stat.WAwait, stat.Svctm, stat.Util)
	return
}

func jsonOut(stat *ExtendedIoStats) (err error) {
	var d []byte
	d, err = json.Marshal(*stat)
	os.Stdout.Write(d)
	return
}
