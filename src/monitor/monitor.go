package main

import (
	"code.google.com/p/goconf/conf"
	"com.cosm"
	"log"
	"os"
	"sysinfo"
)

type Config struct {
	RestUrl string
	ApiKey  string
	FeedId  string
}
type Monitor struct {
	collectors []DataCollector
	stopCh     chan int
	dsCh       chan *cosm.Datastream
	c          *Config
}

func InitMonitor(configPath string) (*Monitor, error) {
	var c *conf.ConfigFile
	var err error
	c, err = conf.ReadConfigFile(configPath)
	if err != nil {
		return nil, err
	}

	config := new(Config)
	config.RestUrl, err = c.GetString("default", "url")
	if err != nil {
		return nil, err
	}
	config.ApiKey, err = c.GetString("key", "all")
	if err != nil {
		return nil, err
	}

	config.FeedId, err = c.GetString("default", "feedid")
	if err != nil {
		return nil, err
	}

	monitor := new(Monitor)
	monitor.c = config
	monitor.dsCh = make(chan *cosm.Datastream, 1)
	monitor.stopCh = make(chan int, 1)
	collector1, e1 := createDiskCollector(c, monitor.dsCh)
	if e1 != nil {
		return nil, e1
	}
	monitor.collectors = append(monitor.collectors, collector1)

	collector2, e2 := createLoadAverage(c, monitor.dsCh)
	if e2 != nil {
		return nil, e2
	}
	monitor.collectors = append(monitor.collectors, collector2)

	collector3, e3 := createMemoryUsage(c, monitor.dsCh)
	if e3 != nil {
		return nil, e3
	}
	monitor.collectors = append(monitor.collectors, collector3)
	return monitor, nil
}

func (self *Monitor) Start() {
	for _, collector := range self.collectors {
		collector.StartGetData()
	}

	go func() {
		for ds := range self.dsCh {
			UpdateDatastreams(self.c, ds)
		}
	}()
	<-self.stopCh
}

func (self *Monitor) Stop() {
	self.stopCh <- 1
}

func createDiskCollector(c *conf.ConfigFile, dsCh chan *cosm.Datastream) (DataCollector, error) {
	interval, err := c.GetInt("diskUsage", "interval")
	if err != nil {
		return nil, err
	}
	diskUsage := sysinfo.NewDiskUsage(interval, dsCh)
	return diskUsage, nil
}

func createLoadAverage(c *conf.ConfigFile, dsCh chan *cosm.Datastream) (DataCollector, error) {
	interval, err := c.GetInt("loadAverage", "interval")
	if err != nil {
		return nil, err
	}
	loadAverage := sysinfo.NewLoadAverage(interval, dsCh)
	return loadAverage, nil
}

func createMemoryUsage(c *conf.ConfigFile, dsCh chan *cosm.Datastream) (DataCollector, error) {
	interval, err := c.GetInt("memoryUsage", "interval")
	if err != nil {
		return nil, err
	}
	memoryUsage := sysinfo.NewMemoryUsage(interval, dsCh)
	return memoryUsage, nil
}

func main() {
	log.SetOutput(os.Stdout)

	monitor, err := InitMonitor("../config/config.test")
	if err != nil {
		log.Println("init error:", err)
	}

	monitor.Start()
}
