package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 创建一个通道来接收信号
	sigChan := make(chan os.Signal, 1)

	// 注册要接收的信号
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 启动一个协程来处理信号
	go func() {
		sig := <-sigChan
		fmt.Println("Received signal:", sig)
		// 执行清理操作
		fmt.Println("Cleaning up...")
		time.Sleep(2 * time.Second)
		fmt.Println("Exiting...")
		os.Exit(0)
	}()

	fmt.Println("Running... Press Ctrl+C to exit")
	for {
		// 模拟长时间运行的任务
		time.Sleep(1 * time.Second)
	}
}
