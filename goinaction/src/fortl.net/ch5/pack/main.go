package main

import (
	"fortl.net/ch5/pack/counters"
	"fmt"
	"fortl.net/ch5/pack/entities"
)

func main() {

	//第一，公开或者未公开的标识符，不是一个值。
	//第二，短变量声明操作符，有能力捕获引用的类型，并创建一个未公开的类型的变量。
	//永远不能显式创建一个未公开的类型的变量，不过短变量声明操作符可以这么做
	counter := counters.New(10)
	fmt.Println(counter)

	//u := entities.User{Name:"bill", em}

	a := entities.Admin{Rights:10}

	//设置未公开的内部类型的公开字段的值
	a.Name = "bill"
	a.Email = "bill@qq.com"
}


