package main

import (
	"sync"
	"runtime"
	"fmt"
)

var(
	cc1 int
	w4 sync.WaitGroup
	// mutex 用来定义一段代码临界区
	mutex sync.Mutex



)

// mutex 锁
func main() {

	w4.Add(2)

	go inccc(1)
	go inccc(2)

	w4.Wait()

	fmt.Println("Ret: ", cc1)
}

func inccc(id int) {
	defer w4.Done()

	for count:= 0; count <2; count++ {
		// 同一时刻只允许一个 goroutine 进入
		// 这个临界区
		mutex.Lock()
		{
			value := cc1
			runtime.Gosched()
			value++
			cc1 = value
		}
		// 释放锁，允许其他正在等待的 goroutine
		mutex.Unlock()
	}
}