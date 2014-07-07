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

func (z ZmqOutput) Connect(url string) {
	z.sendSocket, z.err = zmq.NewSocket(zmq.PULL)
	z.sendSocket.Connect(url)
}

func (z ZmqOutput) Close() {
	if nil != z.sendSocket {
		z.sendSocket.Close()
	}
}

func (z ZmqOutput) SendStats(eStat *ExtendedIoStats) (err error) {
	if nil == z.sendSocket {
		err = errors.New("Nil socket, call zmqOutput.Connect() before trying to send stats")
		return
	}
	var (
		stats []ProtoStat
	)

	stats, err = GetProtoStats(eStat)
	if nil != err {
		fmt.Println(err)
	}
	for _, stat := range stats {
		data, err := proto.Marshal(&stat)
		if nil != err {
			fmt.Println("Failed to marshal stat message : ", stat)
		}
		//just print the encoded data for now... soon this will actually send a queue
		fmt.Println(data)
	}
	return
}

// //  Socket to receive messages on
// receiver, _ := zmq.NewSocket(zmq.PULL)
// defer receiver.Close()
// receiver.Connect("tcp://localhost:5557")

// //  Socket to send messages to
// sender, _ := zmq.NewSocket(zmq.PUSH)
// defer sender.Close()
// sender.Connect("tcp://localhost:5558")

// //  Process tasks forever
// for {
// 	s, _ := receiver.Recv(0)

// 	//  Simple progress indicator for the viewer
// 	fmt.Print(s + ".")

// 	//  Do the work
// 	msec, _ := strconv.Atoi(s)
// 	time.Sleep(time.Duration(msec) * time.Millisecond)

// 	//  Send results to sink
// 	sender.Send("", 0)
// }
