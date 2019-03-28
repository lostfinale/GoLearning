package main

import (
	"encoding/json"
	"os"
	"log"
	"fmt"
)

type (

	server struct {
		ipAddress string `json:"ip"`
		sid int `json:"id"`
		sName string `json:"name"`
	}

	platServer struct {
		serverList []server `json:"servers"`
		oper string `json:"plat"`
	}
)

var file = "src/fortl.net/ch8/json/data.json"

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	file, err := os.Open(file)
	//file, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		log.Fatalln("Failed to open error json file:", err)
	}
	var ps platServer
	err = json.NewDecoder(file).Decode(&ps)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(ps)
}