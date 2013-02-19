package sysinfo

import (
	"model"
	"testing"
	"time"
)

func NoTestMemoryUsage(t *testing.T) {
	dvCh := make(chan *model.DataValue, 1)
	memoryUsage := NewMemoryUsage(1)
	memoryUsage.StartGetData(dvCh)
	go func() {
		for v := range dvCh {
			t.Log("memoryUsage:", v.Value, "%")
		}
	}()
	time.Sleep(2 * time.Second)
	memoryUsage.Stop()
	time.Sleep(2 * time.Second)
}
