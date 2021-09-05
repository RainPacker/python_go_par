package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 向渠道放数据
func putNum(size int, ch chan<- int) {
	for i := 0; i < size; i++ {
		ch <- i
	}
	// 关闭
	close(ch)
	wg.Done()
}

// 取数据并判断 是否是 一个素数
func getPrimalAndCheck(nums <-chan int, primChan chan<- int, existChan chan<- bool) {
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
func PrintPrimal(primChan chan int, name string) {
	for v := range primChan {
		fmt.Println("素数", v, " name=", name)
	}
	wg.Done()
}
func main() {
	var numCh = make(chan int, 3000)
	var primChan = make(chan int, 3000)
	wg.Add(1)
	go putNum(120000, numCh)
	var threadCount = 16
	var exitChan = make(chan bool, threadCount)
	for i := 0; i < threadCount; i++ {
		wg.Add(1)
		go getPrimalAndCheck(numCh, primChan, exitChan)
	}
	wg.Add(1)
	// 开启打印线程
	go PrintPrimal(primChan, "t1")
	wg.Add(1)
	// 开启打印线程
	go PrintPrimal(primChan, "t2")
	// wg.Add(1)
	// // 开启打印线程
	// go PrintPrimal(primChan, "t3")
	// wg.Add(1)
	// // 开启打印线程
	// go PrintPrimal(primChan, "t4")
	// wg.Add(1)
	// // 开启打印线程
	// go PrintPrimal(primChan, "t5")
	// wg.Add(1)
	// // 开启打印线程
	// go PrintPrimal(primChan, "t6")
	wg.Add(1)
	go func() {
		for i := 0; i < threadCount; i++ {
			<-exitChan
		}
		close(primChan)
		wg.Done()
	}()
	wg.Wait()

	// 只写 管道
	var chWritOnly = make(chan<- int, 1)
	// 只读 管道
	var chReadOnly = make(<-chan int, 1)
	fmt.Println(chWritOnly, chReadOnly)

	var ch1 = make(chan int, 10)
	var ch2 = make(chan string, 10)
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	for i := 0; i < 10; i++ {
		ch2 <- "hellogo " + fmt.Sprintf("-%d", i)
	}
	//使用多路复用 时 管道不用关闭
	for {
		select {
		case v := <-ch1:
			fmt.Printf("从 ch1 取到%v\n", v)
		case v := <-ch2:
			fmt.Printf("从 ch1 取到%v\n", v)
		default:
			// 结束循环
			return
		}
	}
	fmt.Println("主线程完成")
}
