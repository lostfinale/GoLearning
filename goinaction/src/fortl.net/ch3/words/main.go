package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"fortl.net/ch3/words/count"
)

func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)

	c := count.CountWorlds(text)

	fmt.Printf("总共有%d单词. \n", c)



}