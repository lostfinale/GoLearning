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

func sendNotification(n notifier) {
	n.notify()
}

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


	sendNotification(bill)





}