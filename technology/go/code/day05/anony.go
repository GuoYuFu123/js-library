package main

import "fmt"

// 首字母大写，全局调用
var FnAll = func(n1 int, n2 int) int {
	return n1 + n2
}

func main() {

	// 匿名函数直接调用
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res=", res)

	// 匿名赋值给变量，通过变量来调用
	fn := func(n1 int, n2 int) int {
		return n1 + n2
	}
	res1 := fn(10, 30)
	fmt.Println("res1=", res1)

	// 全局匿名函数
	res2 := FnAll(10, 40)
	fmt.Println("res2=", res2)
}
