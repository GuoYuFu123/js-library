package main

import "fmt"

func main() {
	fmt.Println("slice")

	var intArr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("intarr=%T,", intArr)
	slice := intArr[1:2]
	fmt.Printf("intarr=%T, intarr=%v \n", intArr, intArr)
	fmt.Printf("slice=%T, sliec= %v \n", slice, slice)
	fmt.Println("slice元素个数=", len(slice)) //1
	// slice 应用数组的指定索引起始位置到数组结尾就是 slice 的容量
	fmt.Println("slice的容量", cap(slice))   // 4

	// make
	slice1 := make([]int, 5,10)
	fmt.Println("slice1=", slice1)

	// 直接定义切片
	slice2 := []int {1,2,3,4,5}
	fmt.Println("slice2=", slice2)
	fmt.Println("slice元素个数=", len(slice2))
	fmt.Println("slice2的容量", cap(slice2))   
}
