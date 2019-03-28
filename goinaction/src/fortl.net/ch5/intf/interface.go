package main

import (
	"os"
	"fmt"
)

func init() {
	if len(os.Args) != 2{
		fmt.Println("Usage: ./example2 <url>")
		os.Exit(-1)
	}
}


type notifier interface {
	notify()
}



type user struct {
	name string
	email string
}

//notify 使用值接收者实现了一个方法
func (u *user) notify() {
	fmt.Printf("发送邮件给：%s<%s>\n", u.name, u.email)
}
//changeEmail 使用指针接收者实现了一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}


type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", d)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {

	/*// 从 Web 服务器得到响应
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// 从 Body 复制到 Stdout
	io.Copy(os.Stdout, r.Body)

	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}*/


	bill := user{
		name :"bill",
		email :"bill@qq.com",
	}


	sendNotification(&bill)

	//方法集定义了一组关联到给定类型的值或者指针的方法。
	//定义方法时使用的接收者的类型决定了这个方法是关联到值，还是关联到指针，还是两个都关联。


	//从值的角度看
	//T 类型的值的方法集只包含值接收者声明的方法。
	//而指向 T 类型的指针的方法集既包含值接收者声明的方法，也包含指针接收者声明的方法

	//从接受者角度看
	//如果使用指针接收者来实现一个接口，那么只有指向那个类型的指针才能够实现对应的接口。
	//如果使用值接收者来实现一个接口，那么那个类型的值和指针都能够实现对应的接口。



	/*
	方法集的规则

	Values  Methods Receivers
	-----------------------------------------------
	T  		(t T)
	*T 		(t T) and (t *T)


	Methods Receivers Values
	-----------------------------------------------
	(t T) 	T and *T
	(t *T)  *T

	*/

	//编译器并不是总能自动获得一个值的地址
	//试图获取 duration 类型的值的地址，但是获取不到
	fmt.Println(duration(42).pretty())

}