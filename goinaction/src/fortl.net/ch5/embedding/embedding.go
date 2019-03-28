package main

import "fmt"

//嵌入类型（是不是和java的继承类似？）

type person struct {
	name string
	email string
}

func (p *person) send (){
	fmt.Printf("Sending user email to %s<%s>\n",
		 p.name,
		 p.email)
}

//一旦我们将 person 类型嵌入 master，我们就可以说 person 是外部类型 master 的内部类型。
type master struct {
	person //嵌入类型
	level string
}

func(m *master) send(){
	fmt.Printf("Sending master email to %s<%s>\n",
		m.name,
		m.email)
}

type sender interface {
	send()
}

func main() {
	p := master {
		person: person{
			name:"lily",
			email:"lily@qq.com",
		},
		level :"super",
	}

	// 我们可以直接访问内部类型的方法
	p.person.send()

	// 内部类型的方法也被提升到外部类型
	p.send()

	send(&p)
}


func send(s sender) {
	s.send()
}