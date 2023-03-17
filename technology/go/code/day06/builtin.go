package main

import (
	"fmt"
)

func main() {
	fmt.Println("builtin")

	num1 := 100
	fmt.Printf("num1的类型%T, 值=%v, 地址=%v \n", num1, num1, num1)

	// new 来分配地址
	// 这个是new先分配了一个地址，然后返回这个指针地址，这个指针地址，也有自己的地址，赋值的话需要赋值给的这个指针
	num2 := new(int) // *int
	*num2 = 200
	fmt.Printf("num2的类型%T, 值=%v, 地址=%v , 指向=%v\n", num2, num2, &num2, num2)
}
