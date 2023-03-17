package main

import "fmt"

func fb(n int) int {
	if n <= 1 {
		return 1
	} else {
		return fb(n-1) + fb(n-2)
	}
}

func eat(n int) int {
	if n == 10 {
		return 1
	} else {
		return (eat(n+1) + 1) * 2
	}

}

func main() {

	// 斐波那契
	fmt.Println("fb=", fb(4))

	// 猴子吃桃
	fmt.Println("eat=",eat(1))
}
