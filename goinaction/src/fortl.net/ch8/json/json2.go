package main

import (
	"encoding/json"
	"log"
	"fmt"
)

// JSON 包含要反序列化的样例字符串
var JSON = `{
  "name": "Gopher",
	"title": "programmer",
	"contact": {
	"home": "415.333.3333",
	"cell": "415.555.5555"
	}
}`
//无法为 JSON 的格式声明一个结构类型，而是需要更加灵活的方式来处理 JSON 文档。
//在这种情况下，可以将 JSON 文档解码到一个 map 变量中
func main() {
	var c map[string]interface{}

	err := json.Unmarshal([]byte(JSON), &c)

	if err != nil {
		log.Println("ERROR", err)
		return
	}

	fmt.Println(c)

	fmt.Println("H:", c["contact"].(map[string]interface{})["home"])
}