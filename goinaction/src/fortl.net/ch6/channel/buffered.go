package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

/**
	有缓冲的通道（buffered channel）是一种在被接收前能存储一个或者多个值的通道。这种类
	型的通道并不强制要求 goroutine 之间必须同时完成发送和接收。通道会阻塞发送和接收动作的
	条件也会不同。只有在通道中没有要接收的值时，接收动作才会阻塞。只有在通道没有可用缓冲
	区容纳被发送的值时，发送动作才会阻塞。这导致有缓冲的通道和无缓冲的通道之间的一个很大
	的不同：无缓冲的通道保证进行发送和接收的 goroutine 会在同一时间进行数据交换；有缓冲的
	通道没有这种保证。
 */

const (
	numberGoroutines = 4 // 要使用的 goroutine 的数量
	taskLoad = 10 // 要处理的工作的数量
)

var w3 sync.WaitGroup

// init 初始化包，Go 语言运行时会在其他代码执行之前
func init() {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
}

func main() {

	tasks := make(chan string, taskLoad)

	w3.Add(numberGoroutines)

	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}
	// 增加一组要完成的工作
	for post:=1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// 当所有工作都处理完时关闭通道
	// 以便所有 goroutine 退出
	//当通道关闭后，goroutine 依旧可以从通道接收数据，但是不能再向通道里发送数据。
	close(tasks)
	// 等待所有工作完成
	w3.Wait()

}
// worker 作为 goroutine 启动来处理
// 从有缓冲的通道传入的工作
func worker(tasks chan string, worker int) {

	defer w3.Done()

	for {
		task, ok := <-tasks
		if !ok {
			// 这意味着通道已经空了，并且已被关闭
			fmt.Printf("Worker: %d : Shuting Down\n", worker)
			return
		}

		// 显示我们开始工作了
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// 随机等一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Workder: %d : Completed %s\n", worker, task)

	}
}