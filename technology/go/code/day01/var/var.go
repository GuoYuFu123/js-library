package main

import "fmt"

// var global1 = 100
// var global2 = "jack"

var (
	global1 = 100
	global2 = "jack"
)


func main()  {
	// 1 先类型定义，再类型赋值
	var num int
	num = 10
	fmt.Println(num);

	// 2  自行类型推导
	var age = 10
	fmt.Println(age)

	// 3  省略var ，相当于 var name string;  name="tom"
	name := "tom";
	fmt.Println(name)

	// 4  一次性声明多个变量[变量类型一样]
	var n1, n2,n3 int
	fmt.Println(n1,n2,n3)

	// 5  一次性声明多个变量,类型推导[变量类型不一样]
	var m1, m2, m3 = 100, "tom" , 888
	fmt.Println(m1,m2,m3)

	// 6  一次性声明多个变量,类型推导，省略var[变量类型不一样]
	b1, b2, b3 := 100, "to m" , 888
	fmt.Println(b1,b2,b3)

	// 7  全局变量声明
	fmt.Println(global1, global2)
}