package sysinfo

import (
	"model"
	"testing"
	"time"
)

func NoTestDiskUsage(t *testing.T) {
	dvCh := make(chan *model.DataValue, 1)
	diskUsage := NewDiskUsage(1)
	diskUsage.StartGetData(dvCh)
	go func() {
		for v := range dvCh {
			t.Log("diskUsage:", v.Value, "%")
		}
	}()
	time.Sleep(2 * time.Second)
	diskUsage.Stop()
	time.Sleep(2 * time.Second)
}
