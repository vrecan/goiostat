package main

import (
   "os"
   // "os/signal"
   // "syscall"
   "bufio"
   "log"
   "strings"
   "time"
   "./diskStat"
   "./ioStatTransform"
   "flag"
   // "fmt"

)
/**
Go version of iostat, pull stats from proc and optionally log or send to a zeroMQ
*/

var interval = flag.Int("interval", 5, "Interval that stats should be reported.")

const linuxDiskStats = "/proc/diskstats"

func main() {
  flag.Parse()
  for {
    my_channel := make(chan diskStat.DiskStat, 1)
    go ioStatTransform.TransformStat(my_channel)

    // // Handle SIGINT and SIGTERM.
    // ch := make(chan os.Signal)
    // signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

    file,err := os.Open(linuxDiskStats)
    if nil != err {
  		log.Fatal(err)
  	}
    
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	line := strings.Fields(scanner.Text())
    	stat,err := diskStat.LineToStat(line)
    	if(nil != err) {
    		log.Fatal(err)
    	}
    	my_channel <- stat
    }
    file.Close()
    if err := scanner.Err(); err != nil {
      log.Fatal(err)
  	}
    time.Sleep(time.Second * time.Duration(*interval))
  }

}
