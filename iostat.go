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
   "./statsOutputLog"
   "flag"
   // "fmt"
   // "runtime"

)
/**
Go version of iostat, pull stats from proc and optionally log or send to a zeroMQ
*/

var interval = flag.Int("interval", 5, "Interval that stats should be reported.")

const linuxDiskStats = "/proc/diskstats"

func main() {
  flag.Parse()
  // // Handle SIGINT and SIGTERM.
  // ch := make(chan os.Signal)
  // signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
  statsTransformChannel := make(chan diskStat.DiskStat, 10)
  statsOutputChannel := make(chan diskStat.ExtendedIoStats, 10)
  go ioStatTransform.TransformStat(statsTransformChannel, statsOutputChannel)
  go statsOutputLog.Output(statsOutputChannel)

  for {
    readAndSendStats(statsTransformChannel)
    time.Sleep(time.Second * time.Duration(*interval))

  }
  close(statsTransformChannel)
}

func readAndSendStats(statsTransformChannel chan diskStat.DiskStat) {

    file,err := os.Open(linuxDiskStats)
    if nil != err {
      log.Fatal(err)
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      line := strings.Fields(scanner.Text())
      stat,err := diskStat.LineToStat(line)
      if(nil != err) {
        log.Fatal(err)
      }
      statsTransformChannel <- stat
    }

    if err := scanner.Err(); err != nil {
      log.Fatal(err)
    }
}