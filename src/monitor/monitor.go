package main

import (
	"code.google.com/p/goconf/conf"
	"log"
	"model"
	"os"
	"sysinfo"
)

type Config struct {
	RestUrl             string
	ApiKey              string
	FeedId              string
	DiskUsageInterval   int
	MemoryUsageInterval int
	LoadAverageInterval int
}

type Monitor struct {
	collectors []model.DataCollector
	stopCh     chan int
	dvCh       chan *model.DataValue
	config     *Config
}

func loadConfig(configPath string) (*Config, error) {
	var c *conf.ConfigFile
	var err error
	c, err = conf.ReadConfigFile(configPath)
	if err != nil {
		return nil, err
	}

	config := new(Config)
	if config.RestUrl, err = c.GetString("default", "url"); err != nil {
		return nil, err
	}

	if config.ApiKey, err = c.GetString("key", "all"); err != nil {
		return nil, err
	}

	if config.FeedId, err = c.GetString("default", "feedid"); err != nil {
		return nil, err
	}

	if config.DiskUsageInterval, err = c.GetInt("diskUsage", "interval"); err != nil {
		return nil, err
	}

	if config.MemoryUsageInterval, err = c.GetInt("memoryUsage", "interval"); err != nil {
		return nil, err
	}

	if config.LoadAverageInterval, err = c.GetInt("loadAverage", "interval"); err != nil {
		return nil, err
	}
	return config, nil
}

func InitMonitor(configPath string) (*Monitor, error) {

	config, e0 := loadConfig(configPath)
	if e0 != nil {
		return nil, e0
	}
	monitor := new(Monitor)
	monitor.config = config
	monitor.dvCh = make(chan *model.DataValue, 1)
	monitor.stopCh = make(chan int, 1)
	collector1 := createDiskCollector(monitor.config)
	monitor.collectors = append(monitor.collectors, collector1)

	collector2 := createLoadAverage(monitor.config)
	monitor.collectors = append(monitor.collectors, collector2)

	collector3 := createMemoryUsage(monitor.config)
	monitor.collectors = append(monitor.collectors, collector3)
	return monitor, nil
}

func (self *Monitor) Start() {
	for _, collector := range self.collectors {
		collector.StartGetData(self.dvCh)
	}

	go func() {
		for v := range self.dvCh {
			UpdateDatastreams(self.config, v)
		}
	}()
	<-self.stopCh
}

func (self *Monitor) Stop() {
	self.stopCh <- 1
}

func createDiskCollector(c *Config) model.DataCollector {
	return sysinfo.NewDiskUsage(c.DiskUsageInterval)
}

func createLoadAverage(c *Config) model.DataCollector {
	return sysinfo.NewLoadAverage(c.LoadAverageInterval)
}

func createMemoryUsage(c *Config) model.DataCollector {
	return sysinfo.NewMemoryUsage(c.MemoryUsageInterval)
}

func main() {
	log.SetOutput(os.Stdout)

	monitor, err := InitMonitor("../config/config.test")
	if err != nil {
		log.Println("init error:", err)
	}
	monitor.Start()
}
