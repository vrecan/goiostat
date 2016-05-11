package main

//main  goiostat application that allows you to send send extended iostat info
//over zmq using protobuffers or json.

import (
	"bufio"
	"flag"
	"github.com/vrecan/goiostat/diskStat"
	"github.com/vrecan/goiostat/ioStatTransform"
	"github.com/vrecan/goiostat/logOutput"
	"github.com/vrecan/goiostat/nanoMsgOutput"
	"github.com/vrecan/goiostat/outputInterface"
	. "github.com/vrecan/goiostat/protocols"
	"github.com/vrecan/goiostat/statsOutput"
	"github.com/vrecan/goiostat/zmqOutput"
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

	if nil != err {
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
		output, err = zmqOutput.NewZmqOutput(queueUrl, proto)
	case "nano":
		output, err = nanoMsgOutput.NewNanoMsgOutput(queueUrl, proto)
	default:
		output = &logOutput.LogOutput{proto}
	}
	if nil != err {
		log.Error("Failed to setup output ", err)
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

//Read stats from proc and report stats
func readAndSendStats(statsTransformChannel chan *diskStat.DiskStat) {

	file, err := os.Open(linuxDiskStats)
	if nil != err {
		log.Error(err)
		return
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
