package main

import (
	"fmt"
)



type user struct {
	name string
	email string
}



//notify 使用值接收者实现了一个方法
func (u user) notify() {
	fmt.Printf("发送邮件给：%s<%s>\n", u.name, u.email)
}
//changeEmail 使用指针接收者实现了一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {

	bill := user{
		name :"bill",
		email :"bill@qq.com",
	}


	bill.notify()

	//使用 bill 的值作为接收者进行调用，方法 notify 会接收到 bill 的值的一个副本
	lisa := &user{name: "lisa", email:"lisa@qq.com",}

	//也可以使用指针来调用使用值接收者声明的方法
	lisa.notify()

	bill.changeEmail("bill@163.com")
	bill.notify()
	lisa.changeEmail("lisa@163.com")
	lisa.notify()



	//Go 语言里有两种类型的接收者：值接收者和指针接收者

	//GO 语言里的内置类型有 数值类型、字符串类型和布尔类型
	//Go 语言里的引用类型有 切片、映射、通道、接口和函数类型
	//GO 语言中的结构类型是 struct自定义的类型

}
