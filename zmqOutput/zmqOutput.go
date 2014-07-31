package zmqOutput

//zmqOutput Package that allows you to send stats over zeromq.

import (
	"code.google.com/p/goprotobuf/proto"
	"encoding/json"
	"errors"
	. "github.com/CapillarySoftware/goiostat/diskStat"
	. "github.com/CapillarySoftware/goiostat/protoStat"
	. "github.com/CapillarySoftware/goiostat/protocols"
	log "github.com/cihub/seelog"
	zmq "github.com/pebbe/zmq3"
)

type ZmqOutput struct {
	Proto      Protocol
	sendSocket *zmq.Socket
	err        error
}

func (z *ZmqOutput) Connect(url string) {
	z.sendSocket, z.err = zmq.NewSocket(zmq.PUSH)
	z.sendSocket.Connect(url)
}

func (z *ZmqOutput) send(data *[]byte) (r int, err error) {
	// fmt.Println(*z)
	if nil == z.sendSocket {
		err = errors.New("Nil Socket, can't send")
		return
	}
	r, err = z.sendSocket.SendBytes(*data, 0)
	if nil != err {
		log.Error(err)
	}
	return
}

func (z *ZmqOutput) Close() {
	if nil != z.sendSocket {
		z.sendSocket.Close()
	}
}

func (z *ZmqOutput) SendStats(eStat *ExtendedIoStats) (err error) {
	switch z.Proto {
	case PProtoBuffers:
		{
			err = z.SendProtoBuffers(eStat)
		}
	case PJson:
		{
			err = z.SendJson(eStat)
		}

	default:
		{
			err = errors.New("zmqOutput doesn't support the type given... ")
			return
		}
	}
	return
}

func (z *ZmqOutput) SendProtoBuffers(eStat *ExtendedIoStats) (err error) {
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
		data, mErr := proto.Marshal(&stat)
		if nil != mErr {
			err = mErr
			return
		}
		_, err = z.send(&data)
	}
	return
}

func (z *ZmqOutput) SendJson(eStat *ExtendedIoStats) (err error) {
	if nil == z.sendSocket {
		err = errors.New("Nil socket, call zmqOutput.Connect() before trying to send stats")
		return
	}

	data, err := json.Marshal(&eStat)
	if nil != err {
		return
	}
	_, err = z.send(&data)
	return
}
