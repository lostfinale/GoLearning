package main

import (
	"encoding/json"
	"os"
	"log"
	"fmt"
)

type (

	Server struct {
		IpAddress string `json:"ip"`
		Sid int `json:"id"`
		SName string `json:"name"`
	}

	PlatServer struct {
		ServerList []Server `json:"list"`
		Oper string `json:"plat"`
	}

	Ret struct {
		R PlatServer `json:"server"`
	}
)

var file = "src/fortl.net/ch8/json/data.json"

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	//file, err := os.Open(file)
	file, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		log.Fatalln("Failed to open error json file:", err)
	}
	var ps Ret
	err = json.NewDecoder(file).Decode(&ps)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(ps)
}