package main

func main() {

	/**日志模块
	Log := logger.GetLogger()
	Log.Debug("Hello World")

	// Log必须在最后调用WaitForDone方法，否则最后一条log会出现无法记录的情况。
	Log.WaitForDone()
	*/

	/**
	ch1 := make(chan float64)
	ch2 := make(chan uint64)
	go collector.GetCPUUsage(ch1)
	go collector.GetTotalMemory(ch2)
	go collector.GetFreeMemory(ch2)

	fmt.Println("other things")
	cpuUsage := <-ch1
	totalMemory := <-ch2
	FreeMemory := <-ch2
	fmt.Println(cpuUsage)
	fmt.Println(totalMemory)
	fmt.Println(FreeMemory)
	*/

	/**
	// 对象转换******************************************************************************
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
	// ************************************************************************************************




	// start server ***********************************************************************************
	hs := exporterApi.NewHttpServer()
	hs.Start()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := hs.Server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting")
	// *************************************************************************************************
	*/
}
