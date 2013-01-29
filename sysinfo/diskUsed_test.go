package sysinfo

import (
	"code.google.com/p/goconf/conf"
	"com.cosm"
	"testing"
	"time"
)

func TestDiskUsed(t *testing.T) {
	c, err := conf.ReadConfigFile("../config/config.test")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	interval, err2 := c.GetInt("diskUsed", "interval")
	if err2 != nil {
		t.Log(err2)
		t.FailNow()
	}

	dsCh := make(chan *cosm.Datastream, 1)
	diskUsed := NewDiskUsed(interval, dsCh)
	go diskUsed.GetData()
	go func() {
		for ds := range dsCh {
			t.Log("diskUsed:", ds.CurrentValue)
		}
	}()
	time.Sleep(10 * time.Second)
	diskUsed.Stop <- 1
	time.Sleep(2 * time.Second)
}
