package protoStat

//protoStat Package built to manage protobuffer stats messages.

import (
	"bytes"
	"errors"
	. "github.com/CapillarySoftware/goiostat/diskStat"
	log "github.com/cihub/seelog"
	"reflect"
)

//GetProtoStats get a slice of protostats from extendedIOStats.
func GetProtoStat(eStat *ExtendedIoStats) (stats *ProtoStats, err error) {
	var protoStat []*ProtoStat
	deviceName := eStat.Device
	stats = new(ProtoStats)

	s := reflect.ValueOf(eStat).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {

		f := s.Field(i)

		switch f.Kind() {
		case reflect.Float64:
			var buf bytes.Buffer
			buf.WriteString(deviceName)
			buf.WriteString("_")
			// fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
			buf.WriteString(typeOfT.Field(i).Name)
			name := buf.String()
			value := f.Float()
			msg := &ProtoStat{Key: &name, Value: &value}
			protoStat = append(protoStat, msg)
		case reflect.String:
			continue //this is expected to happen on the first index which has the name
		default:
			var buf bytes.Buffer
			kind := f.Kind().String()
			buf.WriteString("Invalid type: ")
			buf.WriteString(kind)
			buf.WriteString(" given")
			errors.New(buf.String())
		}
	}
	log.Info("Info", protoStat)
	log.Flush()
	stats.Stats = protoStat
	return
}
