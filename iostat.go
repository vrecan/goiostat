package main

import (
   "os"
   "bufio"
   "log"
   "strings"
   "time"
   "./diskStat"
   "./ioStatTransform"
   // "fmt"

)
/**
Go version of iostat, pull stats from proc and optionally log or send to a zeroMQ
*/


const linuxDiskStats = "/proc/diskstats"

func main() {
  for {
    my_channel := make(chan diskStat.DiskStat, 1000)
    go ioStatTransform.TransformStat(my_channel)

    file,err := os.Open(linuxDiskStats)
    if nil != err {
  		log.Fatal(err)
  	}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	// fmt.Println(scanner.Text())
    	line := strings.Fields(scanner.Text())
    	stat,err := diskStat.LineToStat(line)
    	if(nil != err) {
    		log.Fatal(err)
    	}
    	my_channel <- stat
    }
    if err := scanner.Err(); err != nil {
      log.Fatal(err)
  	}
    time.Sleep(5 * time.Second)
  }
}
