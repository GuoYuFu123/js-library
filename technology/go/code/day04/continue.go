package main

import "fmt"

func main() {
	fmt.Println("continue")

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 2 {
				continue
			}
			fmt.Println("i=", i, "\tj=", j)
		}
	}

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("奇数=", i)
	}

	fmt.Println("---- continue label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue re
		}
	}

}
