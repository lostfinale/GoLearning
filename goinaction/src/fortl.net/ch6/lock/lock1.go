package main

import (
	"sync"
	"fmt"
	"time"
	"sync/atomic"
)

var (
	shutdown int64
	w2 sync.WaitGroup
)
//AddInt64 LoadInt64 StoreInt64  //原子函数
func main() {
	w2.Add(2)

	go doWork("A")
	go doWork("B")

	// 给定 goroutine 执行的时间
	time.Sleep(1 * time.Second)

	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	w2.Wait()
}


func doWork(name string) {
	defer w2.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Microsecond)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutding %s Down\n", name)
			break
		}
	}
}