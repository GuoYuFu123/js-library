package main

import "fmt"

func getSequence() func() int {
	i := 0

	return func() int {
		i += 1
		return i
	}
}

func main() {
	fmt.Println("闭包累加")

	nextNumber := getSequence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

}
