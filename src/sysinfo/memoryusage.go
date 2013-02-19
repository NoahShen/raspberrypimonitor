package sysinfo

import (
	"log"
	"model"
	"syscall"
	"time"
	"utils"
)

type MemoryUsage struct {
	Id       string
	stopCh   chan int
	interval int
}

func NewMemoryUsage(interval int) *MemoryUsage {
	memoryUsage := &MemoryUsage{Id: "memoryUsage", stopCh: make(chan int, 1), interval: interval}
	return memoryUsage
}

func (self *MemoryUsage) StartGetData(dvChannel chan<- *model.DataValue) {
	go func() {
		for {
			select {
			case <-time.After((time.Duration)(self.interval) * time.Second):
				usedPercent, err := getMemoryUsage()
				if err != nil {
					log.Println("get memory used error", err)
					continue
				}
				formated := utils.FormatFloatToPercent(usedPercent)
				value := &model.DataValue{Id: self.Id, Value: formated}
				dvChannel <- value
			case <-self.stopCh:
				break
			}
		}
	}()

}

func (self *MemoryUsage) Stop() {
	self.stopCh <- 1
}

func getMemoryUsage() (float64, error) {
	//system memory usage
	sysInfo := new(syscall.Sysinfo_t)
	err := syscall.Sysinfo(sysInfo)
	if err != nil {
		return 0.0, err
	}
	all := sysInfo.Totalram
	free := sysInfo.Freeram
	used := all - free

	usedPercent := float64(used) / float64(all)
	return usedPercent, nil
}
