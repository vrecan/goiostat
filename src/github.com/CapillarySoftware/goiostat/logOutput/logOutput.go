package logOutput

import (
	"fmt"
	. "github.com/CapillarySoftware/goiostat/diskStat"
	"github.com/dustin/go-humanize"
)

type LogOutput struct {
}

func (l LogOutput) SendStats(stat *ExtendedIoStats) (err error) {
	fmt.Printf("%s:  rrqm/s %.2f wrqm/s %.2f r/s %.2f w/s %.2f rsize/s %s wsize/s %s avgrq-sz %.2f avgqu-sz %.2f, await %.2f r_await %.2f w_await %.2f svctm %.2f util %.2f%% \n\n",
		stat.Device, stat.ReadsMerged, stat.WritesMerged, stat.Reads, stat.Writes, humanize.Bytes(uint64(stat.SectorsRead)),
		humanize.Bytes(uint64(stat.SectorsWrite)), stat.Arqsz, stat.AvgQueueSize, stat.Await, stat.RAwait, stat.WAwait, stat.Svctm, stat.Util)
	return
}
