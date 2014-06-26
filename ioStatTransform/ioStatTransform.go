package ioStatTransform
import(
   "fmt"
    "../diskStat"
   )

func TransformStat(channel <-chan diskStat.DiskStat) (err error) {
	lastRawStat: map[string]diskStat.DiskStat{}
for {
		stat := <- channel
		lastStat, in := lastRawStat[stat.device]
		if in {
			//do calculations
			fmt.Println(lastStat)
		} else {
			lastRawStat[stat.device] = stat
		}
	}
}