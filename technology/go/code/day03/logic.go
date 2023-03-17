package main

import "fmt"

func main() {
	// 关系运算符
	var n1 int = 9
	var n2 int = 10

	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 > n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 <= n2)

	flag := n1 > n2
	fmt.Println("flag=", flag)

	// 逻辑运算法
	var age int = 40
	if age > 30 && age == 40 && age < 50 {
		fmt.Println("ok")
	}
	if age > 30 || age == 40 {
		fmt.Println("ok")
	}
	if !(age < 30) {
		fmt.Println("ok")
	}

}
