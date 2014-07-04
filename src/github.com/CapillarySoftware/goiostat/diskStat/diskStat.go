package diskStat

import (
	"errors"
	"strconv"
	"time"
)

type DiskStat struct {
	Id                    int64
	PartId                int64
	Device                string
	ReadsCompleted        int64
	ReadsMerged           int64
	SectorsRead           uint64
	MillisReading         int64
	WritesCompleted       int64
	WritesMerged          int64
	SectorsWrite          uint64
	MillisWriting         int64
	IoInProgress          int64
	MillisDoingIo         int64
	WeightedMillisDoingIo int64
	RecordTime            int64
	IoTotal               int64
	SectorsTotalRaw       uint64
}

type ExtendedIoStats struct {
	Device       string
	ReadsMerged  float64
	WritesMerged float64
	Writes       float64
	Reads        float64
	SectorsRead  float64
	SectorsWrite float64
	Arqsz        float64
	AvgQueueSize float64
	Await        float64
	RAwait       float64
	WAwait       float64
	Util         float64
	Svctm        float64
}

/*
  Simple function that takes a string and converts it to a stat
  TODO: Figure out better way to map string to struct
*/
func LineToStat(line []string) (stat DiskStat, err error) {
	var tmp int64
	var sectorsReadRaw uint64
	var sectorsWriteRaw uint64
	if len(line) < 14 {
		err = errors.New("Line is not a valid lenght for disk stats")
		return
	}

	stat.RecordTime = time.Now().UnixNano()

	tmp, err = strconv.ParseInt(line[0], 10, 64)
	if nil != err {
		return stat, err
	}
	stat.Id = tmp

	tmp, err = strconv.ParseInt(line[1], 10, 64)
	if nil != err {
		return stat, err
	}
	stat.PartId = tmp

	stat.Device = line[2]

	tmp, err = strconv.ParseInt(line[3], 10, 64)
	if nil != err {
		return stat, err
	}
	stat.ReadsCompleted = tmp

	tmp, err = strconv.ParseInt(line[4], 10, 64)
	if nil != err {
		return stat, err
	}
	stat.ReadsMerged = tmp

	sectorsReadRaw, err = strconv.ParseUint(line[5], 10, 64)
	if nil != err {
		return stat, err
	}
	stat.SectorsRead = (sectorsReadRaw * 1024) / 2

	tmp, err = strconv.ParseInt(line[6], 10, 64)
	if nil != err {
		return stat, err
	}
	stat.MillisReading = tmp

	tmp, err = strconv.ParseInt(line[7], 10, 64)
	if nil != err {
		return stat, err
	}
	stat.WritesCompleted = tmp

	tmp, err = strconv.ParseInt(line[8], 10, 64)
	if nil != err {
		return stat, err
	}
	stat.WritesMerged = tmp

	sectorsWriteRaw, err = strconv.ParseUint(line[9], 10, 64)
	if nil != err {
		return stat, err
	}
	stat.SectorsWrite = (sectorsWriteRaw * 1024) / 2

	tmp, err = strconv.ParseInt(line[10], 10, 64)
	if nil != err {
		return stat, err
	}
	stat.MillisWriting = tmp

	tmp, err = strconv.ParseInt(line[11], 10, 64)
	if nil != err {
		return stat, err
	}
	stat.IoInProgress = tmp

	tmp, err = strconv.ParseInt(line[12], 10, 64)
	if nil != err {
		return stat, err
	}
	stat.MillisDoingIo = tmp

	tmp, err = strconv.ParseInt(line[13], 10, 64)
	if nil != err {
		return stat, err
	}
	stat.WeightedMillisDoingIo = tmp

	stat.IoTotal = stat.ReadsCompleted + stat.WritesCompleted
	stat.SectorsTotalRaw = sectorsReadRaw + sectorsWriteRaw

	return stat, err

}
