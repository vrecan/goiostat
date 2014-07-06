package outputInterface
import (
	"github.com/CapillarySoftware/goiostat/diskStat"
)

type Output interface {
	SendStats(diskStat.ExtendedIoStats)
}