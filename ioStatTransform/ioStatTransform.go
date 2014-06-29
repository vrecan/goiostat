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

func TransformStat(channel <-chan diskStat.DiskStat) (err error) {
for {
		stat := <- channel
		prevStat,in := LastRawStat[stat.Device]

		if in {
			//ignore partitions with no history of activity
			if((stat.ReadsCompleted == 0 && stat.WritesCompleted == 0) || partition.MatchString(stat.Device)) {
				continue
			}
			timeDiffMilli,err := getTimeDiffMilli(prevStat.RecordTime, stat.RecordTime)
			if(nil != err) { fmt.Println(err);continue}

			readsMerged,err := getOneSecondAvg(prevStat.ReadsMerged, stat.ReadsMerged, timeDiffMilli)
			if(nil != err) { fmt.Println(err);continue}
			writesMerged,err := getOneSecondAvg(prevStat.WritesMerged, stat.WritesMerged, timeDiffMilli)
			if(nil != err) { fmt.Println(err);continue}

			writes,err := getOneSecondAvg(prevStat.WritesCompleted, stat.WritesCompleted, timeDiffMilli)
			if(nil != err) { fmt.Println(err);continue}
			reads,err := getOneSecondAvg(prevStat.ReadsCompleted, stat.ReadsCompleted, timeDiffMilli)
			if(nil != err) { fmt.Println(err);continue}
			
			sectorsRead,err := getOneSecondAvgUint(prevStat.SectorsRead, stat.SectorsRead, timeDiffMilli)
			if(nil != err) { fmt.Println(err);continue}						
			sectorsWrite,err := getOneSecondAvgUint(prevStat.SectorsWrite, stat.SectorsWrite, timeDiffMilli)
			if(nil != err) { fmt.Println(err);continue}	


			arqsz := getAvgRequestSize(prevStat.SectorsTotal, stat.SectorsTotal, prevStat.IoTotal, stat.IoTotal)
			avgQueueSize := getAvgQueueSize(prevStat.WeightedMillisDoingIo, stat.WeightedMillisDoingIo, timeDiffMilli)

			util,err := getUtilization(prevStat.MillisDoingIo, stat.MillisDoingIo, timeDiffMilli)
			if(nil != err) { fmt.Println(err);continue}	

			svctm := getAvgServiceTime(prevStat.IoTotal, stat.IoTotal, timeDiffMilli, util)
			
			fmt.Printf( "%s:  rrqm/s %.2f wrqm/s %.2f r/s %.2f w/s %.2f rsize/s %s wsize/s %s avgrq-sz %.2f avgqu-sz %.2f svctm %.2f util %.2f%% \n\n", 
				stat.Device, readsMerged, writesMerged, reads, writes, humanize.Bytes(uint64(sectorsRead)), 
					humanize.Bytes(uint64(sectorsWrite)), arqsz, avgQueueSize, svctm, util)
		}
		LastRawStat[stat.Device] = stat
	}
}

func getTimeDiffMilli( old int64,  cur int64) (r float64, err error){
	if(old >= cur) {
		err= errors.New("Time has moved backward or not moved at all... impressive!")
		return
	}
	r = float64(cur - old) / 1000000.0 
	return
}

func getOneSecondAvg(old int64, cur int64, time float64) (r float64, err error) {
	if(old > cur) {
		err= errors.New("A stat has rolled over!")
		return
	}
	r = float64(float64(cur - old) / time) * oneSecondInMilli
	return
}

func getOneSecondAvgUint(old uint64, cur uint64, time float64) (r float64, err error) {
	if(old > cur) {
		err= errors.New("A stat has rolled over!")
		return
	}
	r = float64(float64(cur - old) / time) * oneSecondInMilli
	return
}

func getAvgRequestSize(oldSectorsTotal uint64, curSectorsTotal uint64, oldIoTotal int64, curIoTotal int64) (r float64) {
	ioTotal := curIoTotal - oldIoTotal 
	sectorsTotal := curSectorsTotal - oldSectorsTotal 

	if(0 == ioTotal) {
		r = 0.00
		return
	}

	r = float64(sectorsTotal) / float64(ioTotal)
	return
}

func getAvgQueueSize(oldWeightedMillisDoingIo int64, curWeightedMillisDoingIo int64, time float64) (r float64){
	if(oldWeightedMillisDoingIo > curWeightedMillisDoingIo) {
		r = 0.00
		return
	}
	r = float64(curWeightedMillisDoingIo - oldWeightedMillisDoingIo) / time;
	return
}

func getAvgServiceTime(oldIoTotal int64, curIoTotal int64, time float64, util float64) (r float64){
	hz := systemCall.GetClockTicksPerSecond()
	tput := float64(curIoTotal - oldIoTotal) * float64(hz) / time

	if(tput <= 0) {r=0.0; return}
	r = util / tput
	return


}

func getUtilization(old int64, cur int64, time float64) (r float64, err error) {
	if(old > cur) {
		err = errors.New("No IO Happened?")
		return
	}
	r =  (float64(cur - old) / (time * 100) * 10.0) * oneSecondInMilli
	return
}


