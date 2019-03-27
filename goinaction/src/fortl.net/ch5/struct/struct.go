package main

import "fmt"

type user struct {
	name string
	email string
	ext int
	privileged bool
}

type admin struct {
	person user
	level string
}

//在代码清单 5-2 的第 10 行，关键字 var 创建了类型为 user 且名为 bill 的变量。当声明
//变量时，这个变量对应的值总是会被初始化。这个值要么用指定的值初始化，要么用零值（即变
//量类型的默认值）做初始化。对数值类型来说，零值是 0；对字符串来说，零值是空字符串；对
//布尔类型，零值是 false。对这个例子里的结构，结构里每个字段都会用零值初始化

var bill user



//基于 int64 声明一个新类型
//把 int64 类型叫作 Duration 的基础类型。不过，虽然 int64 是基础类型，
//Go 并不认为 Duration 和 int64 是同一种类型。这两个类型是完全不同的有区别的类型。
type Duration int64

func main() {

	//使用结构字面量创建结构类型的值
	lisa := user {
		name:"Lisa",
		email: "lisa@qq.com",
		ext: 123,
		privileged: true,
	}
	p(lisa)

	//不使用字段名，创建结构类型的值 必须保证顺序
	l1 := user {"Lisa","lisa@qq.com", 123,true}

	p(l1)


	l2 := admin{
		person : user {
			name:"Lisa",
			email: "lisa@qq.com",
			ext: 123,
			privileged: true,
		},
		level : "super",
	}

	p(l2)


	var dur Duration
	//编译器不会对不同类型的值做隐式转换 需要手动强制转换
	dur = Duration(int64(1000))

	p(dur)

}

func p(v interface{}) {
	fmt.Println(v)
}