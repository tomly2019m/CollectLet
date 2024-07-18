package main

import (
	"CollectLet/exporterApi"
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {

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
}
