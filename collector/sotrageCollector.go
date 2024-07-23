package collector

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func GetTotalMemory(ch chan<- uint64) {
	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
	}
	ch <- virtualMemory.Total
}

func GetUsedMemory(ch chan<- uint64) {
	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
	}
	ch <- virtualMemory.Used
}

func GetFreeMemory(ch chan<- uint64) {
	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
	}
	ch <- virtualMemory.Available
}

func GetFreeDiskSpace(ch chan<- uint64) {
	usage, err := disk.Usage("/")
	if err != nil {
		fmt.Println(err)
	}
	ch <- usage.Free
}
