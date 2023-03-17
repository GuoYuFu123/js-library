package main

import "fmt"

func test() int {
	return 10
}

func main() {
	fmt.Println("赋值运算符")

	// 数值交换
	a := 9
	b := 2
	fmt.Printf("a=%v,b=%v \n", a, b)
	t := a
	a = b
	b = t
	fmt.Printf("a=%v,b=%v \n", a, b)

	a += 10
	fmt.Printf("a=%v \n", a)

	a = test() + 10
	fmt.Printf("a=%v \n", a)

	// 有2个变量，进行数值交换，但是不允许使用中间变量
	fmt.Println("面试题")
	var j int = 10
	var k int = 20
	j = j + k // j = j + k
	k = j - k // k = j + k - k  =>  k = j;
	j = j - k // j = j + k - j  =>  j = k

	fmt.Printf("j = %v,k = %v \n", j, k)

	var aaa int = 100

	ptr := &aaa

	fmt.Printf("&ptr = %v %p  %v", aaa, ptr, *ptr)
}
