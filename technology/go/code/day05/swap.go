package main

import "fmt"

func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

func swap2(n1 int, n2 int) (n3 int, n4 int) {
	n3 = n2
	n4 = n1
	return
}

func main() {
	fmt.Println("swap")
	n1 := 10
	n2 := 20
	swap(&n1, &n2)
	fmt.Println("swap", n1, n2)

	n3, n4 := swap2(n1, n2)
	fmt.Println("swap", n1, n2)
	fmt.Println("n3=", n3, "n4=", n4)
}
