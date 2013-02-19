package sysinfo

import (
	"model"
	"testing"
	"time"
)

func TestLoadAverage(t *testing.T) {
	dvCh := make(chan *model.DataValue, 1)
	loadAverage := NewLoadAverage(1)
	loadAverage.StartGetData(dvCh)
	go func() {
		for v := range dvCh {
			t.Log("load average:", v.Value)
		}
	}()
	time.Sleep(2 * time.Second)
	loadAverage.Stop()
	time.Sleep(2 * time.Second)
}
