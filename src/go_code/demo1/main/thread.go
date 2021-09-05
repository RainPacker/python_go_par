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
func putVal(ch chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		// 将 i  放入管理
		ch <- i
		fmt.Println("放入数据", i)
		time.Sleep(time.Second * 1)
	}
	close(ch)
}
func getVal(ch chan int) {
	for v := range ch {
		fmt.Println("读取值", v)
	}
	wg.Done()
}

func checkIsExist(exitChan chan bool, count int) {

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
		fmt.Println(time.Now().Unix())
		//	time.Sleep(time.Second)
		fmt.Println(time.Now().Unix())
	}

	//管道
	var ch = make(chan int, 10)
	// for i := 0; i < 10 ; i++ {
	// 	ch <- i
	// }
	wg.Add(1)
	go putVal(ch)
	wg.Add(1)
	go getVal(ch)
	wg.Add(1)
	go getVal(ch)

	var numCh = make(chan int, 3000)
	var primChan = make(chan int, 3000)
	wg.Add(1)
	go putNum(100, numCh)
	var threadCount = 16
	var exitChan = make(chan bool, threadCount)
	for i := 0; i < threadCount; i++ {
		wg.Add(1)
		go getPrimalAndCheck(numCh, primChan, exitChan)
	}
	wg.Add(1)
	// 开启打印线程
	go PrintPrimal(primChan)
	wg.Add(1)
	go func() {
		for i := 0; i < threadCount; i++ {
			<-exitChan
		}
		close(primChan)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("主线程完成")

}
