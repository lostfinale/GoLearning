package main

import (
	"sync"
	"fmt"
	"time"
)

var w2 sync.WaitGroup

func main() {
	baton := make(chan int)

	w2.Add(1)

	// 为最后一位跑步者将计数加 1
	go Runner(baton)
	// 开始比赛
	baton <- 1
	w2.Wait()
}

// Runner 模拟接力比赛中的一位跑步者
func Runner(baton chan int) {
	var newRunner int

	// 等待接力棒
	runner := <-baton

	// 开始绕着跑道跑步
	fmt.Printf("Runner %d Running With Baton\n", runner)

	// 创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}

	// 围绕跑道跑
	time.Sleep(100 * time.Microsecond)

	// 比赛结束了吗？
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		w2.Done()
		return
	}

	// 将接力棒交给下一位跑步者
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner

}