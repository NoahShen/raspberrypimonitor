package sysinfo

import (
	"com.cosm"
	"log"
	"os/exec"
	"strings"
	"time"
)

const (
	LOAD_AVERAGE_STRING = "load average:"
)

type LoadAverage struct {
	Id        string
	stopCh    chan int
	interval  int
	dsChannel chan<- *cosm.Datastream
}

func NewLoadAverage(interval int, dsCh chan<- *cosm.Datastream) *LoadAverage {
	loadAverage := &LoadAverage{Id: "loadaverage", stopCh: make(chan int, 1)}
	loadAverage.interval = interval
	loadAverage.dsChannel = dsCh
	return loadAverage
}

func (self *LoadAverage) StartGetData() {
	go func() {
		for {
			select {
			case <-time.After((time.Duration)(self.interval) * time.Second):
				loadaverage, err := getLoadAverage()

				if err != nil {
					log.Println("get load average error", err)
					continue
				}
				datestream := &cosm.Datastream{Id: self.Id, CurrentValue: loadaverage}
				self.dsChannel <- datestream
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
