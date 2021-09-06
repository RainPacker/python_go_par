package main

import (
	"fmt"
	"sync"
	"time"
)

var mutext sync.Mutex

// 读写锁
var rwmutext sync.RWMutex
var count int = 1

func add() {
	//互斥锁
	mutext.Lock()
	count += 1
	mutext.Unlock()
}
func read() {
	rwmutext.RLock()
	fmt.Println(count)
	rwmutext.RUnlock()
}
func main() {
	for i := 0; i < 10; i++ {
		go add()
	}
	for i := 0; i < 10; i++ {
		go read()
	}
	fmt.Println(count)
	time.Sleep(time.Second * 4)
}
