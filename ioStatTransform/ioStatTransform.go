package ioStatTransform
import(
   "fmt"
    "../diskStat"
   )
var LastRawStat = make(map[string]diskStat.DiskStat)	

func TransformStat(channel <-chan diskStat.DiskStat) (err error) {
for {
		stat := <- channel
		prevStat,in := LastRawStat[stat.Device]

		if in {
			//do calculations
			// fmt.Println(stat)
			// fmt.Println(prevStat)
			if(stat.ReadsCompleted == 0 && stat.WritesCompleted == 0) {
				continue
			}
			timeDiff:= stat.RecordTime - prevStat.RecordTime
			rrq := stat.ReadsCompleted - prevStat.ReadsCompleted
			fmt.Println( stat.Device, " rrqm/s ", rrq, " timeDiff ", timeDiff / 1000000)

		} 
		LastRawStat[stat.Device] = stat
	}
} 