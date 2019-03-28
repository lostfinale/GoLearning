package main

import "fmt"

//当一个资源需要在 goroutine 之间共享时，通道在 goroutine 之间架起了一个管道，并提供了
//确保同步交换数据的机制。声明通道时，需要指定将要被共享的数据的类型。可以通过通道共享
//内置类型、命名类型、结构类型和引用类型的值或者指针。
func main() {

	// 无缓冲的整型通道
	//unbuffered := make(chan int)

	// 有缓冲的字符串通道
	buffered := make(chan string, 10)

	buffered <- "A"

	value := <- buffered

	fmt.Println(value)

	//----------------------------
	//无缓冲的通道



}