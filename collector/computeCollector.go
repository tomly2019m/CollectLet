package collector

import (
	"CollectLet/logger"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const statPath = "/proc/stat"

var logTagCompute = "[computeCollector]"

type cpuUsage struct {
	User      int64
	Nice      int64
	System    int64
	Idle      int64
	IOWait    int64
	IRQ       int64
	SoftIRQ   int64
	Steal     int64
	Guest     int64
	GuestNice int64
}

func init() {
	logTag = "[computeCollector]"
}

func readCPUUsage() (cpuUsage, error) {
	data, err := os.ReadFile(statPath)
	if err != nil {
		logger.GetLogger().Error("%s %s", logTagCompute, err.Error())
		return cpuUsage{}, err
	}

	lines := strings.Split(string(data), "\n")
	cpuLine := strings.Fields(lines[0])

	if cpuLine[0] != "cpu" {
		logger.GetLogger().Error("%s unexpected format: %s", logTagCompute, cpuLine[0])
		return cpuUsage{}, fmt.Errorf("unexpected format: %s", cpuLine[0])
	}

	user, err := strconv.ParseInt(cpuLine[1], 10, 64)
	if err != nil {
		return cpuUsage{}, err
	}
	nice, err := strconv.ParseInt(cpuLine[2], 10, 64)
	if err != nil {
		return cpuUsage{}, err
	}
	system, err := strconv.ParseInt(cpuLine[3], 10, 64)
	if err != nil {
		return cpuUsage{}, err
	}
	idle, err := strconv.ParseInt(cpuLine[4], 10, 64)
	if err != nil {
		return cpuUsage{}, err
	}
	ioWait, err := strconv.ParseInt(cpuLine[5], 10, 64)
	if err != nil {
		return cpuUsage{}, err
	}
	irq, err := strconv.ParseInt(cpuLine[6], 10, 64)
	if err != nil {
		return cpuUsage{}, err
	}
	softIRQ, err := strconv.ParseInt(cpuLine[7], 10, 64)
	if err != nil {
		return cpuUsage{}, err
	}
	steal, err := strconv.ParseInt(cpuLine[8], 10, 64)
	if err != nil {
		return cpuUsage{}, err
	}
	guest, err := strconv.ParseInt(cpuLine[9], 10, 64)
	if err != nil {
		return cpuUsage{}, err
	}
	guestNice, err := strconv.ParseInt(cpuLine[10], 10, 64)
	if err != nil {
		return cpuUsage{}, err
	}

	return cpuUsage{
		User:      user,
		Nice:      nice,
		System:    system,
		Idle:      idle,
		IOWait:    ioWait,
		IRQ:       irq,
		SoftIRQ:   softIRQ,
		Steal:     steal,
		Guest:     guest,
		GuestNice: guestNice,
	}, nil
}

func calculateUsage(prev, current cpuUsage) float64 {
	prevIdle := prev.Idle + prev.IOWait
	idle := current.Idle + current.IOWait

	prevNonIdle := prev.User + prev.Nice + prev.System + prev.IRQ + prev.SoftIRQ + prev.Steal
	nonIdle := current.User + current.Nice + current.System + current.IRQ + current.SoftIRQ + current.Steal

	prevTotal := prevIdle + prevNonIdle
	total := idle + nonIdle

	totald := total - prevTotal
	idled := idle - prevIdle

	return float64(totald-idled) / float64(totald) * 100
}

func GetCPUUsage() float64 {
	prevUsage, err := readCPUUsage()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(config.Compute.Freq) * time.Millisecond)

	currentUsage, err := readCPUUsage()
	if err != nil {
		logger.GetLogger().Error("%s %s", logTagCompute, err.Error())
	}
	cpuUsage := calculateUsage(prevUsage, currentUsage)
	return cpuUsage
}
