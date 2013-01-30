package sysinfo

import (
	"com.cosm"
	"log"
	"syscall"
	"time"
	"utils"
)

type DiskUsage struct {
	Id        string
	stopCh    chan int
	interval  int
	dsChannel chan<- *cosm.Datastream
}

func NewDiskUsage(interval int, dsCh chan<- *cosm.Datastream) *DiskUsage {
	diskUsed := &DiskUsage{Id: "diskUsage", stopCh: make(chan int, 1)}
	diskUsed.interval = interval
	diskUsed.dsChannel = dsCh
	return diskUsed
}

func (self *DiskUsage) StartGetData() {
	go func() {
		for {
			select {
			case <-time.After((time.Duration)(self.interval) * time.Second):
				usedPercent, err := getDiskUsage()

				if err != nil {
					log.Println("get disk used error", err)
					continue
				}
				formated := utils.FormatFloatToPercent(usedPercent)
				datestream := &cosm.Datastream{Id: self.Id, CurrentValue: formated}
				self.dsChannel <- datestream
			case <-self.stopCh:
				break
			}
		}
	}()

}

func (self *DiskUsage) Stop() {
	self.stopCh <- 1
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
