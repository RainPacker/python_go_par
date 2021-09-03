package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test i=", i)
		time.Sleep(time.Second * 2)
	}
	wg.Done()
}
func main() {
	cupNum := runtime.NumCPU()
	fmt.Println("cpu num is", cupNum)
	runtime.GOMAXPROCS(cupNum)
	// 开启协程
	wg.Add(1)
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main i=", i)
		time.Sleep(time.Second)
	}
	wg.Wait()
	fmt.Println("主线程完成")
}
