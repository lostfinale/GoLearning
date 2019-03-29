package main

import (
	"encoding/json"
	"log"
	"fmt"
)

//列化 JSON 字符串
func main() {

	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home":"422.233.444",
		"cell":"223,224.992",
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}


	fmt.Println(string(data))

}
