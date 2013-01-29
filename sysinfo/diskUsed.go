package sysinfo

import (
	"com.cosm"
	"log"
	"strconv"
	"syscall"
	"time"
)

type DiskUsed struct {
	Id        string
	Stop      chan int
	interval  int
	dsChannel chan<- *cosm.Datastream
}

func NewDiskUsed(interval int, dsCh chan<- *cosm.Datastream) *DiskUsed {
	diskUsed := &DiskUsed{Id: "diskUsed"}
	diskUsed.Stop = make(chan int, 1)
	diskUsed.interval = interval
	diskUsed.dsChannel = dsCh
	return diskUsed
}

func (self *DiskUsed) GetData() {
	for {
		select {
		case <-time.After((time.Duration)(self.interval) * time.Second):
			usedPercent, err := getDiskUsage()

			if err != nil {
				log.Println("get disk used error", err)
				continue
			}
			formated := strconv.FormatFloat(usedPercent*100, 'f', 2, 64)
			datestream := &cosm.Datastream{Id: self.Id, CurrentValue: formated}
			self.dsChannel <- datestream
		case <-self.Stop:
			break
		}
	}
}

func getDiskUsage() (float64, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/", &fs)
	if err != nil {
		return 0.0, err
	}
	all := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)
	used := all - free

	usedPercent := float64(used) / float64(all)
	return usedPercent, nil
}
