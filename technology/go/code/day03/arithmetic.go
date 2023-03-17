package main

import "fmt"

func main() {
	fmt.Println("arithmetic")

	// 说明： 如果运算符都是整数，那么除后，去掉小数部分，保留整数部分
	fmt.Println(10 / 4)

	var n1 float32 = 10 / 4
	fmt.Println(n1)

	// 如果我们希望保留小数部分，就需要让浮点数参数运算
	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	var n3 = 10.4 * 4;
	fmt.Println(n3)

	// % 取模的使用
	// 计算公式： a % b = a - a / b * b
	fmt.Println("10%3", 10%3)
	fmt.Println("-10%3", -10%3)
	fmt.Println("10%-3", 10%-3)
	fmt.Println("-10%3", -10%3)

	// ++ 和 -- 的使用
	var i int = 10
	i++ // i = i + 1
	fmt.Println("i", i)

	i-- // i = i -1
	fmt.Println("i", i)
	
}
