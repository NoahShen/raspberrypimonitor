package main

import (
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"testing"
)

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// disk usage of path/disk
func diskUsage(path string) (*DiskStatus, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return nil, err
	}
	disk := &DiskStatus{}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return disk, nil
}

func NoTestGetDiskStatus(t *testing.T) {
	diskStatus, err := diskUsage("/")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log("All:", diskStatus.All)
	t.Log("Free:", diskStatus.Free)
	t.Log("Used:", diskStatus.Used)

	diskUsedPercent := float64(diskStatus.Used) / float64(diskStatus.All)
	t.Log("Used%:", diskUsedPercent)
	formated := strconv.FormatFloat(diskUsedPercent*100, 'f', 2, 64)
	t.Log("Used:", formated, "%")
}

type MemStatus struct {
	All  uint32 `json:"all"`
	Used uint32 `json:"used"`
	Free uint32 `json:"free"`
	Self uint64 `json:"self"`
}

func MemStat(t *testing.T) (*MemStatus, error) {
	//go self usage
	memStat := new(runtime.MemStats)
	runtime.ReadMemStats(memStat)
	mem := &MemStatus{}
	mem.Self = memStat.Alloc

	//system memory usage
	sysInfo := new(syscall.Sysinfo_t)
	err := syscall.Sysinfo(sysInfo)
	if err != nil {
		return nil, err
	}

	t.Log(sysInfo.Loads)

	mem.All = sysInfo.Totalram
	mem.Free = sysInfo.Freeram
	mem.Used = mem.All - mem.Free
	return mem, nil
}

func NoTestGetMemStatus(t *testing.T) {
	memStatus, err := MemStat(t)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log("All:", memStatus.All)
	t.Log("Free:", memStatus.Free)
	t.Log("Used:", memStatus.Used)

	memoryUsedPercent := float64(memStatus.Used) / float64(memStatus.All)
	formated := strconv.FormatFloat(memoryUsedPercent*100, 'f', 2, 64)
	t.Log("Used:", formated, "%")
}

func TestGetLoadAverage(t *testing.T) {
	out, err := exec.Command("uptime").Output()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	uptimeResult := string(out)
	t.Logf("The output is : %v\n", uptimeResult)
	loadAverageStr := "load average:"
	i := strings.Index(uptimeResult, loadAverageStr)
	loadValue := uptimeResult[i+len(loadAverageStr):]
	t.Logf("loadValue is : %v\n", loadValue)
	splitsAverage := strings.Split(loadValue, ",")
	oneMinuteLoad := strings.TrimSpace(splitsAverage[0])
	t.Logf("oneMinuteLoad is : %v\n", oneMinuteLoad)
}
