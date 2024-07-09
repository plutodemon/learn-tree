package designPattern

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func demo1() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- 1
		wg.Done()

	}()
	go func() {
		fmt.Println("a:", <-ch)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("end")
}

func demo2() {
	// 创建一个通道来接收信号
	sigChan := make(chan os.Signal, 1)

	// 注册要接收的信号
	// kill pid 是发送SIGTERM的信号 ; kill -9 pid 是发送SIGKILL的信号(无法捕获)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	// 启动一个协程来处理信号
	go func() {

		//select {
		//case sig := <-sigChan:
		//	fmt.Printf("Received signal: %s\n", sig)
		//	// 执行清理逻辑，如关闭资源、保存状态等
		//}

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

func demo3(n int) {
	c1 := make(chan bool, 1)
	c2 := make(chan bool, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(n int) {
		defer wg.Done()
		for i := 2; i <= n; i += 2 {
			<-c1
			fmt.Println(i)
			c2 <- true
		}
	}(n)

	go func(n int) {
		defer wg.Done()
		for i := 1; i <= n; i += 2 {
			<-c2
			fmt.Println(i)
			c1 <- true
		}
	}(n)

	// 启动第一个goroutine，因为从1开始，所以先启动打印奇数的协程
	c2 <- true

	wg.Wait()
}
