package collector

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func GetTotalMemory() uint64 {
	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
	}
	return virtualMemory.Total
}

func GetUsedMemory() uint64 {
	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
	}
	return virtualMemory.Used
}

func GetFreeMemory() uint64 {
	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
	}
	return virtualMemory.Available
}

func GetFreeDiskSpace() uint64 {
	usage, err := disk.Usage("/")
	if err != nil {
		fmt.Println(err)
	}
	return usage.Free
}
