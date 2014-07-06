package main

import (
	"os"
	"bufio"
	"flag"
	"github.com/CapillarySoftware/goiostat/diskStat"
	"github.com/CapillarySoftware/goiostat/ioStatTransform"
	"github.com/CapillarySoftware/goiostat/statsOutput"
	"github.com/CapillarySoftware/goiostat/logOutput"
	"log"
	"strings"
	"time"
	// "fmt"
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
	output := logOutput.LogOutput{}
	go ioStatTransform.TransformStat(statsTransformChannel, statsOutputChannel)

	go statsOutput.Output(statsOutputChannel, output)

	for {
		readAndSendStats(statsTransformChannel)
		time.Sleep(time.Second * time.Duration(*interval))

	}
	close(statsTransformChannel)
	close(statsOutputChannel)
}

func readAndSendStats(statsTransformChannel chan diskStat.DiskStat) {

	file, err := os.Open(linuxDiskStats)
	if nil != err {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		stat, err := diskStat.LineToStat(line)
		if nil != err {
			log.Fatal(err)
		}
		statsTransformChannel <- stat
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
