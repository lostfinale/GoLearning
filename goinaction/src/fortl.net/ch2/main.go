package main

import (
	"log"
	"os"
	_ "fortl.net/ch2/matchers"
	"fortl.net/ch2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}

