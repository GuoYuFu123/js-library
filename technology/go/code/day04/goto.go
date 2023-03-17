package main

import "fmt"

func main() {
	var n int = 10

	if n == 10 {
		goto label1
	}
	fmt.Println("ok1")
	fmt.Println("ok2")
	fmt.Println("ok3")

label1:
	fmt.Println("ok4")
}
