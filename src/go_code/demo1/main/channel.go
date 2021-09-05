package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 向渠道放数据
func putNum(size int, ch chan int) {
	for i := 0; i < size; i++ {
		ch <- i
	}
	// 关闭
	close(ch)
	wg.Done()
}

// 取数据并判断 是否是 一个素数
func getPrimalAndCheck(nums chan int, primChan chan int, existChan chan bool) {
	for num := range nums {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
			}
		}
		if flag {
			//  是素数
			primChan <- num
		}
	}
	//每个线程结束 将执行结束标识放到 结束渠道中
	existChan <- true
	wg.Done()
}

// 读取素数 管道 并打印
func PrintPrimal(primChan chan int) {
	for v := range primChan {
		fmt.Println("素数", v)
	}
	wg.Done()
}
func main() {
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
