package assginer

import (
	"CollectLet/logger"
	"CollectLet/util"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

const (
	cpuLimitPercent = 1000
	memLimitMB      = 1048576
)

var basePath = "/sys/fs/cgroup/"

var log = logger.GetLogger()
var logTagCGroup = "[cGroup]"

type LinuxCPU struct {
	// LimitType 限定cpu的方式
	LimitType string

	// Value 指定限定方式下的限定的值
	Value string

	// NumOfCores 限定CGroup的核心使用数量 -1 表示不限定
	NumOfCores int
}

type LinuxMemory struct {
	MemInBytes uint64
}

type LinuxDisk struct {
	ReadInBytes  uint64
	WriteInBytes uint64
}

type LinuxNetwork struct {
	RateInBytesPerSec uint64
}

type CGroup struct {
	Name    string
	PidSet  util.Set
	CPU     LinuxCPU
	Memory  LinuxMemory
	Disk    LinuxDisk
	Network LinuxNetwork
}

func (cg *CGroup) makeDir() {
	cGroupDir := filepath.Join(basePath, cg.Name)
	cpuPath := filepath.Join(cGroupDir, "cpu")
	memPath := filepath.Join(cGroupDir, "memory")
	diskPath := filepath.Join(cGroupDir, "disk")
	networkPath := filepath.Join(cGroupDir, "network")

	if err := os.MkdirAll(cpuPath, 0755); err != nil {
		log.Error("%s %s", logTagCGroup, err.Error())
	}

	if err := os.MkdirAll(memPath, 0755); err != nil {
		log.Error("%s %s", logTagCGroup, err.Error())
	}

	if err := os.MkdirAll(diskPath, 0755); err != nil {
		log.Error("%s %s", logTagCGroup, err.Error())
	}

	if err := os.MkdirAll(networkPath, 0755); err != nil {
		log.Error("%s %s", logTagCGroup, err.Error())
	}
}

func IsProcessAlive(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return true
	}
	return false
}

func (cg *CGroup) init() {

}

func (cg *CGroup) AddProcess(pid int) error {
	if !IsProcessAlive(pid) {
		log.Error("%s process %d not alive", logTagCGroup, pid)
	}
	cg.PidSet.Add(pid)
	return nil
}

func (cg *CGroup) RemoveProcess(pid int) error {
	if !IsProcessAlive(pid) {
		log.Error("%s process %d not alive", logTagCGroup, pid)
	}
	cg.PidSet.Remove(pid)
	return nil
}

// 设置cgroup资源限制
func (cg *CGroup) setCgroupLimit(path, filename, value string) error {
	return os.WriteFile(filepath.Join(path, filename), []byte(value), 0644)
}

// 将进程添加到cgroup
func addProcessToCgroup(path string, pid int) error {
	return os.WriteFile(filepath.Join(path, "tasks"), []byte(strconv.Itoa(pid)), 0644)
}

func (cg *CGroup) setCPULimit() {
	cpuPath := filepath.Join(basePath, "cpu")
	err := cg.setCgroupLimit(cpuPath, cg.CPU.LimitType, cg.CPU.Value)
	if err != nil {
		log.Error("%s %s", logTagCGroup, err.Error())
	}
}

func (cg *CGroup) setCPUCoresLimit() {
	// TODO
}

func (cg *CGroup) setMemoryLimit() {
	memPath := filepath.Join(basePath, "memory")
	err := cg.setCgroupLimit(memPath, "memory.limit_in_bytes", strconv.FormatUint(cg.Memory.MemInBytes, 10))
	if err != nil {
		log.Error("%s %s", logTagCGroup, err.Error())
	}
}

func (cg *CGroup) Apply() {
	cg.setCPULimit()
	cg.setMemoryLimit()
	cg.setCPUCoresLimit()
}
