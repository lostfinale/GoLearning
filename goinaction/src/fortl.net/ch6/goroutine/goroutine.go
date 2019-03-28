package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main() {

	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(2)

	// wg 用来等待程序完成
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()
		for count:= 0; count < 3; count++ {
			for char := 'a'; char < 'a' + 26; char++{
				fmt.Printf("%c ", char)
			}
		}

	}()

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		defer wg.Done()
		for count:= 0; count < 3; count++ {
			for char := 'A'; char < 'A' + 26; char++{
				fmt.Printf("%c ", char)
			}
		}

	}()

	fmt.Println("Waiting to Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")

}
