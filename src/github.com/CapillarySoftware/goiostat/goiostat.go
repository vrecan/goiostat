package main

//main  goiostat application that allows you to send send extended iostat info
//over zmq using protobuffers or json.

import (
	"bufio"
	"flag"
	// "fmt"
	"github.com/CapillarySoftware/goiostat/diskStat"
	"github.com/CapillarySoftware/goiostat/ioStatTransform"
	"github.com/CapillarySoftware/goiostat/logOutput"
	"github.com/CapillarySoftware/goiostat/outputInterface"
	. "github.com/CapillarySoftware/goiostat/protocols"
	"github.com/CapillarySoftware/goiostat/statsOutput"
	"github.com/CapillarySoftware/goiostat/zmqOutput"
	"log"
	"os"
	"strings"
	"time"
)

/**
Go version of iostat, pull stats from proc and optionally log or send to a zeroMQ
*/

var interval = flag.Int("interval", 5, "Interval that stats should be reported.")
var outputType = flag.String("output", "stdout", "output should be one of the following types (stdout,zmq)")
var zmqUrl = flag.String("zmqUrl", "tcp://localhost:5400", "ZmqUrl valid formats (tcp://localhost:[port], ipc:///location/file.ipc)")
var protocolType = flag.String("protocol", "", "Valid protocol types are (protobuffers, json")

const linuxDiskStats = "/proc/diskstats"

func main() {
	flag.Parse()
	statsTransformChannel := make(chan *diskStat.DiskStat, 10)
	statsOutputChannel := make(chan *diskStat.ExtendedIoStats, 10)

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
	var output outputInterface.Output
	proto := PStdOut

	switch *protocolType {
	case "protobuffers":
		{
			proto = PProtoBuffers
		}
	case "json":
		{
			proto = PJson
		}
	default:
		{
			if *outputType == "zmq" {
				proto = PProtoBuffers
			} else if *outputType == "stdout" {
				proto = PStdOut
			}
		}
	}

	switch *outputType {
	case "zmq":
		zmq := &zmqOutput.ZmqOutput{Proto: proto}
		zmq.Connect(*zmqUrl)
		defer zmq.Close()
		output = zmq
	default:
		output = &logOutput.LogOutput{proto}
	}

	go ioStatTransform.TransformStat(statsTransformChannel, statsOutputChannel)

	go statsOutput.Output(statsOutputChannel, output)

	for {
		readAndSendStats(statsTransformChannel)
		time.Sleep(time.Second * time.Duration(*interval))

	}
	close(statsTransformChannel)
	close(statsOutputChannel)
}

func readAndSendStats(statsTransformChannel chan *diskStat.DiskStat) {

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
		statsTransformChannel <- &stat
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
