package main

import "fmt"

func main() {
	// 创建一个映射，键的类型是 string，值的类型是 int
	m1 := make(map[string]int)
	fmt.Println(m1)
	m2 := map[string]int{"R":111, "G":222, "B":333}
	fmt.Println(m2)

	//切片、函数以及包含切片的结构类型这些类型由于具有引用语义，
	//不能作为映射的键，使用这些类型会造成编译错误
	//m3 := map[[]string]int{}

	m4 := map[int][]string{}
	fmt.Println(m4)

	m5 := map[string]string{}
	m5["A"] = "Z"
	m5["B"] = "Y"
	fmt.Println(m5)

	//nil 映射
	//不能用于存储键值对，否则，会产生一个语言运行时错误
	//var colors map[string]string
	//colors["R"] = "111"

	// 从映射获取值并判断键是否存在
	value, exists := m2["R"]
	if exists {
		p(value)
	}

	//从映射获取值，并通过该值判断键是否存在
	v1 := m2["D"]
	if v1 != 0 {
		p(v1)
	}


	//使用 range 迭代映射
	m6 := map[string]string {
		"AliceBlue": "#f0f8ff",
		"Coral": "#ff7F50",
		"DarkGray": "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	for key, value := range m6 {
		fmt.Printf("key: %s, value: %s\n", key, value)
		delete(m6, "key")
	}
	p("")

	delete(m6, "DarkGray")
	for key, value := range m6 {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}
	p("")
	remove(m6, "ForestGreen")
	for key, value := range m6 {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}


}

func p(v interface{}) {
	fmt.Println(v)
}

func remove(m map[string]string, key string) {
	delete(m, key)
}
