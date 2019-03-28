package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

var wg1 sync.WaitGroup


func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)

	wg1.Add(2)

	// 启动两个选手
	go player("Lily", court)
	go player("Dane", court)

	// 发球
	court <- 1

	// 等待游戏结束
	wg1.Wait()
}

func player(name string, court chan int) {

	defer wg1.Done()

	for {
		ball, ok := <- court
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)

		if n %13 == 0 {
			// 关闭通道，表示我们输了
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}

		// 显示击球数，并将击球数加 1
		fmt.Printf("Player %s Hit %d\n", name, ball)

		ball++
		// 将球打向对手
		court <- ball

	}


}

