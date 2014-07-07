package zmqOutput

import(
	"fmt"
	. "github.com/CapillarySoftware/goiostat/diskStat"
	"code.google.com/p/goprotobuf/proto"
	"reflect"
	"bytes"

)

type ZmqOutput struct {
}

func (l ZmqOutput) SendStats (stat ExtendedIoStats) {


	deviceName := stat.Device
	s := reflect.ValueOf(&stat).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		if(i == 0) {continue}
		var buf bytes.Buffer
		buf.WriteString(deviceName)
		buf.WriteString("_")
	    f := s.Field(i)
	    // fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	    buf.WriteString(typeOfT.Field(i).Name)
	    name := buf.String()
	    value := uint64(f.Float())
	    // fmt.Println(&value)
	    protoStat := ProtoStat{Key: &name, Value: &value}
	    fmt.Println(*protoStat.Key)
	    data, err := proto.Marshal(&protoStat)
	    if(nil != err) {
	    	fmt.Println(err)
	    }
	    fmt.Println(data)
	}
}