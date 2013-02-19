package sysinfo

import (
	"log"
	"model"
	"os/exec"
	"strings"
	"time"
)

const (
	LOAD_AVERAGE_STRING = "load average:"
)

type LoadAverage struct {
	Id       string
	stopCh   chan int
	interval int
}

func NewLoadAverage(interval int) *LoadAverage {
	loadAverage := &LoadAverage{Id: "loadaverage", stopCh: make(chan int, 1), interval: interval}
	return loadAverage
}

func (self *LoadAverage) StartGetData(dvChannel chan<- *model.DataValue) {
	go func() {
		for {
			select {
			case <-time.After((time.Duration)(self.interval) * time.Second):
				loadaverage, err := getLoadAverage()
				if err != nil {
					log.Println("get load average error", err)
					continue
				}
				value := &model.DataValue{Id: self.Id, Value: loadaverage}
				dvChannel <- value
			case <-self.stopCh:
				break
			}
		}
	}()

}

func (self *LoadAverage) Stop() {
	self.stopCh <- 1
}

func getLoadAverage() (string, error) {
	out, err := exec.Command("uptime").Output()
	if err != nil {
		return "", err
	}
	uptimeResult := string(out)

	i := strings.Index(uptimeResult, LOAD_AVERAGE_STRING)
	loadValue := uptimeResult[i+len(LOAD_AVERAGE_STRING):]
	splitsAverage := strings.Split(loadValue, ",")
	oneMinuteLoad := strings.TrimSpace(splitsAverage[0])
	return oneMinuteLoad, nil
}
