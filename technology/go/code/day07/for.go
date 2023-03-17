package main

import "fmt"

func main() {
	fmt.Println("slice for")

	// 直接定义切片
	slice := []int {1,2,3,4,5}
	fmt.Println("slice=", slice)
	fmt.Println("slice元素个数=", len(slice))
	fmt.Println("slice的容量", cap(slice))   

	// for
	for i := 0; i < len(slice); i++ {
		fm t.Println("i=",i,"v=", slice[i])
	}

	//  for-range
	fmt.Println("for-range")
	for _, v := range slice {
		fmt.Println(v)
	}
	
}
