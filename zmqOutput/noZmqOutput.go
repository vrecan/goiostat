// +build !zmq

package zmqOutput

//zmqOutput Package that allows you to send stats over zeromq.

import (
	"errors"
	. "github.com/CapillarySoftware/goiostat/diskStat"
	. "github.com/CapillarySoftware/goiostat/protocols"
)

type ZmqOutput struct {
}

func NewZmqOutput(url *string, proto Protocol) (zmq *ZmqOutput, err error) {
	err = errors.New("ZeroMQ is not compiled! Please compile with go build -tags 'zmq' if you want to use ZeroMQ.")
	return
}

func (z *ZmqOutput) Connect(url string) {
	return
}

func (z *ZmqOutput) send(data *[]byte) (r int, err error) {
	err = errors.New("ZeroMQ is not compiled! Please compile with go build -tags 'zmq' if you want to use ZeroMQ.")
	return
}

func (z *ZmqOutput) Close() {
	return
}

func (z *ZmqOutput) SendStats(eStat *ExtendedIoStats) (err error) {
	err = errors.New("ZeroMQ is not compiled! Please compile with go build -tags 'zmq' if you want to use ZeroMQ.")
	return
}

func (z *ZmqOutput) SendProtoBuffers(eStat *ExtendedIoStats) (err error) {
	err = errors.New("ZeroMQ is not compiled! Please compile with go build -tags 'zmq' if you want to use ZeroMQ.")
	return
}

func (z *ZmqOutput) SendJson(eStat *ExtendedIoStats) (err error) {
	err = errors.New("ZeroMQ is not compiled! Please compile with go build -tags 'zmq' if you want to use ZeroMQ.")
	return
}
