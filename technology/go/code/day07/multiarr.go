package main

import "fmt"

func main() {
	fmt.Println("multiarr")

	// 二维数组
	// 4行 6列
	var arr [4][6]int

	arr[1][2] = 1
	arr[2][2] = 2
	arr[3][2] = 3
	fmt.Println("arr=", arr)

	// 遍历
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			// fmt.Println("i=", i, "j=", j, "v=", arr[i][j])
			fmt.Print(arr[i][j])
			fmt.Print(" ")
		}
		fmt.Println("")

	}

	// 初始化
	var arr2 [2][2]int = [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("arr2=", arr2)

	var arr3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr3=", arr3)
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			// fmt.Println("i=", i, "j=", j, "v=", arr3[i][j])
			fmt.Print(arr3[i][j])
			fmt.Print(" ")
		}
		fmt.Println("")
	}

	fmt.Println("for-range")
	for _, v := range arr3 {
		for _, v2 := range v {
			fmt.Print(v2)
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
