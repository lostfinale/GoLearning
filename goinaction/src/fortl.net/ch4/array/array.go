package main

import "fmt"

//定义一个数组
var array [5]int



func main() {
	//array := [5]int{1,2,3,4,5}
	array := [...]int{1,2,3}
	fmt.Println(array)

	array1 := [5]int{1:10, 4:10}
	array1[0] = 100
	fmt.Println(array1)


	array3 := [5]*int{0:new(int), 1:new(int)}
	*array3[0] = 10
	*array3[1] = 20
	fmt.Println(array3)

	//复制数组
	array4 := array1
	fmt.Println(array4)

	//数组变量的类型包括数组长度和每个元素的类型。只有这两部分都相同的数组，才是类型相
	//同的数组，才能互相赋值


	//var array5 [4]string

	//array6 := [5]string{"1","2","3"}
	//array5 = array6 错误编译不通过


	//-------------------------------
	//复制数组指针，只会复制指针的值，而不会复制指针所指向的值
	var a1 [3]*string
	a2 := [3]*string{new(string), new(string), new(string)}


	*a2[0] = "R"
	*a2[1] = "B"
	*a2[2] = "G"

	a1 = a2

	*a1[0] = "P"
	for _,v := range a1 {
		fmt.Print("     ", *v)
	}
	fmt.Println()
	for _,v := range a2 {
		fmt.Print("     ", *v)
	}
	fmt.Println("\n--------------------------------------------")

	//--------------------------------------
	//多维数组

	//var aa [4][2]int
	//使用数组字面量来声明并初始化一个二维整型数组
	aa := [4][2]int{{1,2},{1,3},{1,4},{1,5}}
	fmt.Println(aa)
	//声明并初始化外层数组中索引为 1 个和 3 的元素
	aa1 := [4][2]int {1:{1,2},3:{2,4}}
	fmt.Println(aa1)
	// 声明并初始化外层数组和内层数组的单个元素
	aa2 := [4][2]int{1:{0:20}, 3:{1:30}}
	fmt.Println(aa2)


	var aa3 [2][2]int
	aa3[0][0] = 10

	var aa4 [2][2]int
	aa4 = aa3
	fmt.Println(aa4)

	//因为每个数组都是一个值，所以可以独立复制某个维度
	var aa5 [2]int = aa4[1]
	fmt.Println(aa5)
	var value int = aa4[1][0]
	fmt.Println(value)

	//---------------------------------
	//在函数间传递数组

	var b [1e6]int
	foo(&b)


}

func foo(array *[1e6]int) {

}
