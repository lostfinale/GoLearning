package main

import (
	"sync"
	"runtime"
	"sync/atomic"
	"fmt"
)

var (
	counter int64
	w1 sync.WaitGroup
)

func main() {
	w1.Add(2)
	go incCounter1(1)
	go incCounter1(2)
	w1.Wait()
	fmt.Println("Final Counter:",  counter)
}

func incCounter1(id int) {
	defer w1.Done()
	for count := 0; count < 2; count++ {
		// 安全地对 counter 加 1
		atomic.AddInt64(&counter, 1)
		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()
	}
}