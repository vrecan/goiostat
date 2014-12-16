// +build !nano

package nanoMsgOutput

//Nanomsg output Package that allows you to send stats over nanomsg.

import (
	"errors"
	. "github.com/CapillarySoftware/goiostat/diskStat"
	. "github.com/CapillarySoftware/goiostat/protocols"
)

type NanoMsgOutput struct {
}

func NewNanoMsgOutput(url *string, proto Protocol) (nano *NanoMsgOutput, err error) {
	err = errors.New("Nanomsg is not compiled! Please compile with go build -tags 'nano' if you want to use nanomsg.")
	return
}

//Method to connect to queue
func (this *NanoMsgOutput) Connect(url string) (err error) {
	err = errors.New("Nanomsg is not compiled! Please compile with go build -tags 'nano' if you want to use nanomsg.")
	return
}

//Send byte data over queue
func (this *NanoMsgOutput) send(data *[]byte) (r int, err error) {
	err = errors.New("Nanomsg is not compiled! Please compile with go build -tags 'nano' if you want to use nanomsg.")
	return
}

//Close the socket
func (this *NanoMsgOutput) Close() {
	return
}

//Send stats by given format
func (this *NanoMsgOutput) SendStats(eStat *ExtendedIoStats) (err error) {
	err = errors.New("Nanomsg is not compiled! Please compile with go build -tags 'nano' if you want to use nanomsg.")
	return
}

//Send stats in protobuffer format
func (this *NanoMsgOutput) SendProtoBuffers(eStat *ExtendedIoStats) (err error) {
	err = errors.New("Nanomsg is not compiled! Please compile with go build -tags 'nano' if you want to use nanomsg.")
	return
}

//Send stats in json format
func (this *NanoMsgOutput) SendJson(eStat *ExtendedIoStats) (err error) {
	err = errors.New("Nanomsg is not compiled! Please compile with go build -tags 'nano' if you want to use nanomsg.")
	return
}
