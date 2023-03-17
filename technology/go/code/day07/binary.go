package main

import "fmt"

func binaryFind(val int, arr *[6]int, leftIndex int, rightIndex int) int {
	if leftIndex > rightIndex {
		return -1
	}

	middle := (leftIndex + rightIndex) / 2
	fmt.Println("middle", middle)

	if arr[middle] == val {
		return middle
	}

	if (*arr)[middle] > val {
		return binaryFind(val, arr, leftIndex, middle-1)
	}
	if (*arr)[middle] < val {
		return binaryFind(val, arr, middle+1, rightIndex)
	}
	return -1
}

func main() {
	fmt.Println("binary")

	arr := [6]int{1, 2, 3, 4, 5, 6}

	binaryFind(3, &arr, 0, len(arr)-1)
}
