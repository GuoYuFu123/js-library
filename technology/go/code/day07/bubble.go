package main

import "fmt"

// 冒泡排序

func bubbleSort(arr *[]int) {

	for i := 1; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i; j++ {
			if (*arr)[j] < (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}

}

func main() {
	fmt.Println("bubble-sort")

	arr := []int{10, 40, 20, 50, 30}
	bubbleSort(&arr)
	fmt.Println("arr", arr)
}
