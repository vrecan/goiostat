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
	statsTransformChannel := make(chan diskStat.DiskStat, 10)
	statsOutputChannel := make(chan diskStat.ExtendedIoStats, 10)

    // c := make(chan os.Signal, 1)
    // signal.Notify(c, os.Interrupt)
    // signal.Notify(c, syscall.SIGTERM)
    // go func() {
    //     <-c
    //     log.Info("Caught signal, shutting down")
    //     close(statsTransformChannel)
    //     close(statsOutputChannel)
    //     log.Info("Shutdown complete")
    //     os.Exit(0)
    // }()
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
