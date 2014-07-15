package outputInterface

//outputInterface simple interface to allow you to send stats using
//multiple interfaces.

import (
	"github.com/CapillarySoftware/goiostat/diskStat"
	// "errors"
)

type Output interface {
	SendStats(*diskStat.ExtendedIoStats) (err error)
}
