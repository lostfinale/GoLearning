package main

import "fmt"

type notifier1 interface {
	notify()
}

type user1 struct {
	email string
	name  string
}

func (u *user1) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)

}

type admin1 struct {
	email string
	name  string
}

func (a *admin1) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

func main() {
	bill := user1 {"1@qq.com", "bill"}
	send(&bill)
	lisa := admin1{"2@qq.com", "lisa"}
	send(&lisa)
}

func send(notify notifier1) {
	notify.notify()
}

