package main

//main  goiostat application that allows you to send send extended iostat info
//over zmq using protobuffers or json.

import (
	"bufio"
	"flag"
	"github.com/CapillarySoftware/goiostat/diskStat"
	"github.com/CapillarySoftware/goiostat/ioStatTransform"
	"github.com/CapillarySoftware/goiostat/logOutput"
	"github.com/CapillarySoftware/goiostat/nanoMsgOutput"
	"github.com/CapillarySoftware/goiostat/outputInterface"
	. "github.com/CapillarySoftware/goiostat/protocols"
	"github.com/CapillarySoftware/goiostat/statsOutput"
	"github.com/CapillarySoftware/goiostat/zmqOutput"
	log "github.com/cihub/seelog"
	"os"
	"strings"
	"time"
)

/**
Go version of iostat, pull stats from proc and optionally log or send to a zeroMQ
*/

var interval = flag.Int("interval", 5, "Interval that stats should be reported.")
var outputType = flag.String("output", "stdout", "output should be one of the following types (stdout,zmq,nano)")
var queueUrl = flag.String("queueUrl", "tcp://localhost:5400", "queueUrl valid formats (tcp://localhost:[port], ipc:///location/file.ipc)")
var protocolType = flag.String("protocol", "", "Valid protocol types are (protobuffers, json")

const linuxDiskStats = "/proc/diskstats"

func main() {
	defer log.Flush()
	logger, err := log.LoggerFromConfigAsFile("seelog.xml")

	if err != nil {
		log.Warn("Failed to load config", err)
	}
	log.ReplaceLogger(logger)
	flag.Parse()
	statsTransformChannel := make(chan *diskStat.DiskStat, 10)
	statsOutputChannel := make(chan *diskStat.ExtendedIoStats, 10)

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
		zmq.Connect(*queueUrl)
		defer zmq.Close()
		output = zmq
	case "nano":
		nano := &nanoMsgOutput.NanoMsgOutput{Proto: proto}
		nano.Connect(*queueUrl)
		output = nano
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
		log.Error(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		stat, err := diskStat.LineToStat(line)
		if nil != err {
			log.Error(err)
		}
		statsTransformChannel <- &stat
	}

	if err := scanner.Err(); err != nil {
		log.Error(err)
	}
}
