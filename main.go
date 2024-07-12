package main

import (
	"CollectLet/cache"
	"fmt"
	"time"
)

func main() {

	lakeFactory := cache.NewLakeFactory()
	// 获取 Compute Lake
	computeLakeInterface, err := lakeFactory.GetObject("compute")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	computeLake, ok := computeLakeInterface.(*cache.Lake[cache.ComputeCache])
	if !ok {
		fmt.Println(err)
	}
	newItem := cache.ComputeCache{
		DataItem: cache.DataItem{
			Name:      "example",
			TimeStamp: time.Now().Unix(),
			Value:     "someValue",
		},
	}
	computeLake.Add(newItem)

	item, err := computeLake.Get()
	if err != nil {
		fmt.Println("获取项失败:", err)
	} else {
		fmt.Printf("获取到的项: %+v\n", item)
	}

	fmt.Printf("Compute Lake: %+v\n", computeLake)

	// 再次获取 Compute Lake，应该是同一个实例
	computeLake2, _ := lakeFactory.GetObject("compute")
	fmt.Println("Compute Lake is same instance:", computeLake == computeLake2)

	// 获取 Storage Lake
	storageLake, _ := lakeFactory.GetObject("storage")
	fmt.Printf("Storage Lake: %+v\n", storageLake)

	// 获取 Network Lake
	networkLake, _ := lakeFactory.GetObject("network")
	fmt.Printf("Network Lake: %+v\n", networkLake)

	// 尝试获取未知类型的 Lake
	_, err = lakeFactory.GetObject("unknown")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
