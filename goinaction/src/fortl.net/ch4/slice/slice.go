package main

import "fmt"

func main() {


	//make 和切片字面量
	//使用长度声明一个字符串切片
	s1 := make([]string, 5)
	fmt.Println(s1)
	// 创建一个整型切片
	// 其长度为 3 个元素，容量为 5 个元素
	s2 := make([]string, 3, 5)
	fmt.Println(s2)

	//容量小于长度的切片会在编译时报错



	//通过切片字面量来声明切片

	s3 := []string{"R", "G", "B"}
	fmt.Println(s3)

	s4 := []int {1,2,3,4,5,6}
	fmt.Println(s4)

	//使用空字符串初始化第 100 个元素
	s5 := []int{99:10}
	fmt.Println(len(s5))

	// 创建 nil 整型切片
	var s6 []int
	fmt.Println(s6 == nil)

	// 使用 make 创建空的整型切片
	s7 := make([]int, 0)
	fmt.Println(s7 == nil)


	// 使用切片字面量创建空的整型切片
	s8 := []int{}
	fmt.Println(s8 == nil)


	//----------------------------------------------------
	//赋值和切片
	ss := []int{10, 20, 30, 40, 50}
	ss[1] = 25

	fmt.Println(ss)


	ss2 := []int{10, 20, 30, 40, 50}
	ss3 := ss2[1:4]
	fmt.Println(ss3)



	ss4 := []int{10,20,30,40,50}
	ss5 := ss4[2:5]
	fmt.Println(ss5)
	ss4[4] = 51

	fmt.Println(ss5)

	//使用 append 向切片增加元素
	a1 := []int{10,20,30,40,50}
	a2 := a1[1:3]
	a3 := append(a2, 60)
	a4 := append(a3, 60)
	//果切片的底层数组没有足够的可用容量，append 函数会创建一个新的底层数组，将被引
	//用的现有的值复制到新数组里，再追加新的值
	a5 := append(a4, 60)
	a6 := append(a5, 60)
	a6[1] = 999
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
	fmt.Println(a6)


	a7 := []string {"A", "B", "C", "D", "E"}
	a8 := a7[2:3:3]
	fmt.Println(a8)
	fmt.Println(cap(a8))

	a9 := a7[2:3:3]
	// 将两个切片追加在一起，并显示结果
	var a10 = append(a9, a7...)
	fmt.Println(a10)

	//因为迭代返回的变量是一个迭代过程中根据切片依次赋值的新变量，所以 value 的地址总
	//是相同的
	for index, value := range a10 {
		fmt.Printf("%d->%d\n", &index, &value)
	}

	for _, value := range a10 {
		fmt.Printf("%s ", value)
	}
	fmt.Println()


	for index := 2; index < len(a10); index++ {
		fmt.Printf("%s ", a10[index])
	}
	fmt.Println("\n\n---------------多维切片--------------------\n")
	//两个特殊的内置函数 len 和 cap，可以用于处理数组、切片和通道。对于切片，函数 len
	//返回切片的长度，函数 cap 返回切片的容量。


	//---------------------------------------------------------------------
	//多维切片

	c := [][]int{{10}, {100, 200}}
	fmt.Println(c)

	c[0] = append(c[0], 20)
	fmt.Println(c)

	c1 := make([]int , 1e6)
	c2 := foo(c1)
	fmt.Println(c2)
}

func foo(slice []int) []int {
	return slice
}