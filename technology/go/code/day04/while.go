package main

import "fmt"

func main() {
	fmt.Println("for 实现 while ")

	// while
	i := 1
	for {
		if i > 5 {
			break
		}
		fmt.Printf("i = %v \n", i)
		i++
	}

	fmt.Println("do---while")
	// do...while 至少执行一次
	j := 1
	for {
		fmt.Printf("j = %v \n", j)
		j++
		if j > 1 {
			break
		}
	}
}
