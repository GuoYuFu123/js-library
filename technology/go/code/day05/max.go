package main

import (
	"fmt"

	u "../utils"
)

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func getSumAndSub(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

func main() {
	fmt.Println("max=", max(3, 5), u.Max(8, 9))

	_, sub := getSumAndSub(5, 7)
	fmt.Println("sub=", sub)

}
