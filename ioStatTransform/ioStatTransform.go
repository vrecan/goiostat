package ioStatTransform
import(
   "fmt"
    "../diskStat"
    "errors"
   )
var LastRawStat = make(map[string]diskStat.DiskStat)	

func TransformStat(channel <-chan diskStat.DiskStat) (err error) {
for {
		stat := <- channel
		prevStat,in := LastRawStat[stat.Device]

		if in {
			//ignore partitions with no history of activity
			if(stat.ReadsCompleted == 0 && stat.WritesCompleted == 0) {
				continue
			}
			timeDiffMilli,err := getTimeDiffMilli(prevStat.RecordTime, stat.RecordTime)
			if(nil != err) { fmt.Println(err);continue}
			readsMerged,err := getOneSecondAvg(prevStat.ReadsMerged, stat.ReadsMerged, timeDiffMilli)
			if(nil != err) { fmt.Println(err);continue}			
			reads,err := getOneSecondAvg(prevStat.ReadsCompleted, stat.ReadsCompleted, timeDiffMilli)
			if(nil != err) { fmt.Println(err);continue}
			fmt.Printf( "%s:  rrqm/s %.2f r/s %.2f  timeDiff %f \n\n", stat.Device, readsMerged, reads, timeDiffMilli)

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
	fmt.Println(cur);
	fmt.Println(old);
	r = float64(float64(cur - old) / time) * 1000
	return
}

