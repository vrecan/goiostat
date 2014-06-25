package ioStatTransform
import(
   "fmt"
    "../diskStat"
   )

func TransformStat(channel <-chan diskStat.DiskStat) (err error) {
for {
		stat := <- channel
		fmt.Println(stat)
	}
}



