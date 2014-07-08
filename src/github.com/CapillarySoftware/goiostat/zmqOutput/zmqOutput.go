package zmqOutput

import (
	"code.google.com/p/goprotobuf/proto"
	"errors"
	"fmt"
	. "github.com/CapillarySoftware/goiostat/diskStat"
	. "github.com/CapillarySoftware/goiostat/protoStat"
	zmq "github.com/pebbe/zmq4"
)

type ZmqOutput struct {
	sendSocket *zmq.Socket
	err        error
}

func (z *ZmqOutput) Connect(url string) {
	z.sendSocket, z.err = zmq.NewSocket(zmq.PUSH)
	z.sendSocket.Connect(url)
	// fmt.Println(*z)
}

func (z *ZmqOutput) send(data *[]byte) (r int, err error) {
	// fmt.Println(*z)
	if nil == z.sendSocket {
		err = errors.New("Nil Socket, can't send")
		return
	}
	r, err = z.sendSocket.SendBytes(*data, 0)
	if nil != err {
		fmt.Println(err)
	}
	return
}

func (z *ZmqOutput) Close() {
	if nil != z.sendSocket {
		z.sendSocket.Close()
	}
}

func (z *ZmqOutput) SendStats(eStat *ExtendedIoStats) (err error) {
	if nil == z.sendSocket {
		err = errors.New("Nil socket, call zmqOutput.Connect() before trying to send stats")
		return
	}
	var (
		stats []ProtoStat
	)

	stats, err = GetProtoStats(eStat)
	if nil != err {
		return //return the error
	}
	for _, stat := range stats {
		data, err := proto.Marshal(&stat)
		if nil != err {
			errors.New("Failed to marshal stat message : ")
		}
		_, err = z.send(&data)
	}
	return
}
