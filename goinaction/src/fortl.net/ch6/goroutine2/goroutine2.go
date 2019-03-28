package main

import (
	"sync"
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(2)

	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("Terminating Program")
}



var wg sync.WaitGroup



// printPrime 显示 5000 以内的素数值
func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)

	}
	fmt.Println("Completed", prefix)
}