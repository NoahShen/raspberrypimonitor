package sysinfo

import (
	"code.google.com/p/goconf/conf"
	"com.cosm"
	"testing"
	"time"
)

func TestDiskUsage(t *testing.T) {
	c, err := conf.ReadConfigFile("../config/config.test")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	interval, err2 := c.GetInt("diskUsage", "interval")
	if err2 != nil {
		t.Log(err2)
		t.FailNow()
	}

	dsCh := make(chan *cosm.Datastream, 1)
	diskUsage := NewDiskUsage(interval, dsCh)
	diskUsage.StartGetData()
	go func() {
		for ds := range dsCh {
			t.Log("diskUsage:", ds.CurrentValue, "%")
		}
	}()
	time.Sleep(10 * time.Second)
	diskUsage.Stop()
	time.Sleep(2 * time.Second)
}
