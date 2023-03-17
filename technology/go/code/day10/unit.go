package main

import "fmt"

func addUpper(n int) int {
	res := 0
	for i := 0; i < n; i++ {
		res += i
	}
	return res
}

func main() {
	fmt.Println("单元测试")

	count := addUpper(10)
	fmt.Println("count=", count)
	if count == 45 {
		fmt.Println("succ")
	} else {
		fmt.Println("err")
	}
}
