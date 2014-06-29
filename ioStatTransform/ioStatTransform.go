package ioStatTransform
import(
   "fmt"
    "../diskStat"
    "../systemCall"
    "errors"
    "regexp"
    "github.com/dustin/go-humanize"
   )
var LastRawStat = make(map[string]diskStat.DiskStat)	
var partition = regexp.MustCompile(`\w.*\d`)
const oneSecondInMilli = 1000

type DiskStatDiff struct {
	Id int64
	PartId int64
	Device string
	ReadsCompleted float64
	ReadsMerged float64
	SectorsRead float64
	MillisReading float64
	WritesCompleted float64
	WritesMerged float64
	SectorsWrite float64
	MillisWriting float64
	// IoInProgress float64 //not used and calculated diff then all others
	MillisDoingIo float64
	WeightedMillisDoingIo float64
	RecordTime float64
    IoTotal float64
	SectorsTotal float64

}

func TransformStat(channel <-chan diskStat.DiskStat) (err error) {
for {
		stat := <- channel
		prevStat,in := LastRawStat[stat.Device]

		if in {
			//ignore partitions with no history of activity
			if((stat.ReadsCompleted == 0 && stat.WritesCompleted == 0) || partition.MatchString(stat.Device)) {
				continue
			}
			diffStat,err := getDiffDiskStat(prevStat, stat);
			if(nil != err) { fmt.Println(err, diffStat);continue}
			timeDiffMilli := getTimeDiffMilli(diffStat.RecordTime)
			readsMerged := getOneSecondAvg(diffStat.ReadsMerged, timeDiffMilli)
			writesMerged := getOneSecondAvg(diffStat.WritesMerged, timeDiffMilli)
			writes := getOneSecondAvg(diffStat.WritesCompleted, timeDiffMilli)
			reads := getOneSecondAvg(diffStat.ReadsCompleted, timeDiffMilli)			
			sectorsRead := getOneSecondAvg(diffStat.SectorsRead, timeDiffMilli)
			sectorsWrite := getOneSecondAvg(diffStat.SectorsWrite, timeDiffMilli)

			arqsz := getAvgRequestSize(diffStat.SectorsTotal, diffStat.IoTotal)
			avgQueueSize := getAvgQueueSize(diffStat.WeightedMillisDoingIo, timeDiffMilli)
			await := getAwait(prevStat.SectorsTotal, stat.SectorsTotal, prevStat.IoTotal, stat.IoTotal)

			util,err := getUtilization(diffStat.MillisDoingIo, timeDiffMilli)
			if(nil != err) { fmt.Println(err);continue}	

			svctm := getAvgServiceTime(prevStat.IoTotal, stat.IoTotal, timeDiffMilli, util)
			
			fmt.Printf( "%s:  rrqm/s %.2f wrqm/s %.2f r/s %.2f w/s %.2f rsize/s %s wsize/s %s avgrq-sz %.2f avgqu-sz %.2f await %.2f  svctm %.2f util %.2f%% \n\n", 
				stat.Device, readsMerged, writesMerged, reads, writes, humanize.Bytes(uint64(sectorsRead)), 
					humanize.Bytes(uint64(sectorsWrite)), arqsz, avgQueueSize, await, svctm, util)
		}
		LastRawStat[stat.Device] = stat
	}
}

func getTimeDiffMilli(diff float64) (r float64) {
	r = diff / 1000000.0
	return
}

func getOneSecondAvg(diff float64, time float64) (r float64) {
	r = float64(diff / time) * oneSecondInMilli
	return
}

// func getOneSecondAvgUint(old uint64, cur uint64, time float64) (r float64, err error) {
// 	if(old > cur) {
// 		err= errors.New("A stat has rolled over!")
// 		return
// 	}
// 	r = float64(float64(cur - old) / time) * oneSecondInMilli
// 	return
// }



func getDiffDiskStat(old diskStat.DiskStat, cur diskStat.DiskStat)(r DiskStatDiff, err error) {
	r.Id = cur.Id
	r.PartId = cur.PartId
	r.Device = cur.Device
	r.ReadsCompleted, err = getDiff(old.ReadsCompleted, cur.ReadsCompleted);
	if(nil != err){return}
	r.ReadsMerged, err = getDiff(old.ReadsCompleted, cur.ReadsCompleted);
	if(nil != err){return}
	// SectorsRead uint64
	r.SectorsRead, err = getDiffUint64(old.SectorsRead, cur.SectorsRead);
	if(nil != err){return}
	// MillisReading int64
		r.MillisReading, err = getDiff(old.MillisReading, cur.MillisReading);
	if(nil != err){return}
	// WritesCompleted int64
		r.WritesCompleted, err = getDiff(old.WritesCompleted, cur.WritesCompleted);
	if(nil != err){return}
	// WritesMerged int64
		r.WritesMerged, err = getDiff(old.WritesMerged, cur.WritesMerged);
	if(nil != err){return}
	// SectorsWrite uint64
		r.SectorsWrite, err = getDiffUint64(old.SectorsWrite, cur.SectorsWrite);
	if(nil != err){return}
	// MillisWriting int64
		r.MillisWriting, err = getDiff(old.MillisWriting, cur.MillisWriting);
	if(nil != err){return}
	// IoInProgress int64 //this stat seems to have old large then cur most of the time??? 
	// r.IoInProgress, err = getDiff(old.IoInProgress, cur.IoInProgress);
	if(nil != err){return}
	// MillisDoingIo int64
		// r.MillisDoingIo, err = getDiff(old.MillisDoingIo, cur.MillisDoingIo);
	if(nil != err){return}
	// WeightedMillisDoingIo 64
	r.WeightedMillisDoingIo, err = getDiff(old.WeightedMillisDoingIo, cur.WeightedMillisDoingIo);
	if(nil != err){r.WeightedMillisDoingIo = 0.00}
	// RecordTime int64
	r.RecordTime, err = getDiff(old.RecordTime, cur.RecordTime);
	if(nil != err){err = nil; fmt.Println(old.RecordTime, cur.RecordTime)}
 //    IoTotal int64
	r.IoTotal, err = getDiff(old.IoTotal, cur.IoTotal);
	if(nil != err){return}
	// SectorsTotal uint64
	r.SectorsTotal, err = getDiffUint64(old.SectorsTotal, cur.SectorsTotal);
	if(nil != err){return}
	return
}

func getDiff(old int64, cur int64)(r float64, err error) {
	if(old > cur) {err=errors.New("Old is newer then current... impressive!");return}
	r= float64(cur - old)
	return
}

func getDiffUint64(old uint64, cur uint64)(r float64, err error) {
	if(old > cur) {err=errors.New("Old is newer then current... impressive!");return}
	r= float64(cur - old)
	return
}

func getAvgRequestSize(diffSectorsTotal float64, diffIoTotal float64) (r float64) {
	if(0 == diffIoTotal) {
		r = 0.00
		return
	}

	r = float64(diffSectorsTotal) / float64(diffIoTotal)
	return
}

func getAvgQueueSize(diffWeightedMillisDoingIo float64, time float64) (r float64){
	if(0 == diffWeightedMillisDoingIo) {r=0.00; return}
	r = float64(diffWeightedMillisDoingIo) / time;
	return
}
	// xds->await = (sdc->nr_ios - sdp->nr_ios) ?
	// 	((sdc->rd_ticks - sdp->rd_ticks) + (sdc->wr_ticks - sdp->wr_ticks)) /
	// 	((double) (sdc->nr_ios - sdp->nr_ios)) : 0.0;
func getAwait(oldSectorsTotal uint64, curSectorsTotal uint64, oldIoTotal int64, curIoTotal int64) (r float64) {
	sectors := float64(curSectorsTotal - oldSectorsTotal)
	io := float64(curIoTotal - oldIoTotal)
	if 0 <= sectors || 0 <= io {r=0.00; return}
	r = sectors / io
	return

}

func getAvgServiceTime(oldIoTotal int64, curIoTotal int64, time float64, util float64) (r float64){
	hz := systemCall.GetClockTicksPerSecond()
	tput := float64(curIoTotal - oldIoTotal) * float64(hz) / time

	if(tput <= 0) {r=0.0; return}
	r = util / tput
	return


}

func getUtilization(diffMillisDoingIo float64, time float64) (r float64, err error) {
	r = (diffMillisDoingIo / (time * 100) * 10.0) * oneSecondInMilli
	if(r > 100.00) {
		r = 100.00;	
	} 
	return
}


