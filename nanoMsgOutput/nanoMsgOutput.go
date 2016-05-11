// +build nano

package nanoMsgOutput

//Nanomsg output Package that allows you to send stats over nanomsg.

import (
	"code.google.com/p/goprotobuf/proto"
	"encoding/json"
	"errors"
	. "github.com/vrecan/goiostat/diskStat"
	. "github.com/vrecan/goiostat/protocols"
	"github.com/vrecan/goiostat/statConversion"
	. "github.com/vrecan/gostat/protoStat"
	log "github.com/cihub/seelog"
	nano "github.com/op/go-nanomsg"
	"time"
)

type NanoMsgOutput struct {
	Proto  Protocol
	socket *nano.PushSocket
	err    error
}

func NewNanoMsgOutput(url *string, proto Protocol) (nano *NanoMsgOutput, err error) {
	nano = &NanoMsgOutput{Proto: proto}
	nano.Connect(*url)
	return
}

//Method to connect to queue
func (this *NanoMsgOutput) Connect(url string) (err error) {

	this.socket, err = nano.NewPushSocket()
	if nil != err {
		return
	}
	_, err = this.socket.Connect(url)
	this.socket.SetSendTimeout(500 * time.Millisecond)
	return
}

//Send byte data over queue
func (this *NanoMsgOutput) send(data *[]byte) (r int, err error) {
	// fmt.Println(*z)
	if nil == this.socket {
		err = errors.New("Nil Socket, can't send")
		return
	}
	r, err = this.socket.Send(*data, 0)
	if nil != err {
		log.Error(err)
	}
	return
}

//Close the socket
func (this *NanoMsgOutput) Close() {

	if nil != this.socket {
		this.socket.Close()
	}
}

//Send stats by given format
func (this *NanoMsgOutput) SendStats(eStat *ExtendedIoStats) (err error) {
	switch this.Proto {
	case PProtoBuffers:
		{
			err = this.SendProtoBuffers(eStat)
		}
	case PJson:
		{
			err = this.SendJson(eStat)
		}

	default:
		{
			err = errors.New("zmqOutput doesn't support the type given... ")
			return
		}
	}
	return
}

//Send stats in protobuffer format
func (this *NanoMsgOutput) SendProtoBuffers(eStat *ExtendedIoStats) (err error) {
	if nil == this.socket {
		err = errors.New("Nil socket, call zmqOutput.Connect() before trying to send stats")
		return err
	}
	var (
		stats *ProtoStats
	)

	stats, err = statConversion.GetProtoStats(eStat)
	if nil != err {
		return //return the error
	}
	for _, stat := range stats.Stats {
		data, mErr := proto.Marshal(stat)
		if nil != mErr {
			err = mErr
			return
		}
		_, err = this.send(&data)
	}
	return
}

//Send stats in json format
func (this *NanoMsgOutput) SendJson(eStat *ExtendedIoStats) (err error) {
	if nil == this.socket {
		err = errors.New("Nil socket, call zmqOutput.Connect() before trying to send stats")
		return
	}

	data, err := json.Marshal(&eStat)
	if nil != err {
		return
	}
	_, err = this.send(&data)
	return
}
