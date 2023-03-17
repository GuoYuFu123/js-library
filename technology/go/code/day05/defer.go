package main

import "fmt"

func sum(n1 int, n2 int) int {
	// 当执行到defer是，可以理解为暂时不执行，然后将defer压入栈中，等函数执行完毕再执行defer
	defer fmt.Println("n1=", n1)
	defer fmt.Println("n2=", n2)
	res := n1 + n2
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}