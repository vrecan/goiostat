package outputInterface
import (
	"github.com/CapillarySoftware/goiostat/diskStat"
	// "errors"
)

type Output interface {
	SendStats(*diskStat.ExtendedIoStats)(err error)
}