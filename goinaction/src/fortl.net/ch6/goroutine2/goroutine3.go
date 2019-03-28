package main

import (
	"sync"
	"runtime"
	"fmt"
)

var (
	counter int
	wait sync.WaitGroup
)

func main() {
	wait.Add(2)
	go incCounter(1)
	go incCounter(2)
	wait.Wait()
	fmt.Println("Final Counter:",  counter)
}

func incCounter(id int) {
	defer wait.Done()
	for count := 0; count < 2; count++ {
		value := counter
		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()
		value++
		counter = value
	}
}
