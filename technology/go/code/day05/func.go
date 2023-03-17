package main

import "fmt"

func getSum(a int, b int) int {
	return a + b
}

type typeFn = func(int, int) int

func funcParam(fn typeFn, n1 int, n2 int) int {
	return fn(n1, n2)
}

func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func countSum(n1 int, args ...int) (sum int) {
	sum = n1
	for _, v := range args {
		sum += v
	}
	return
}

func main() {
	fmt.Println("函数")
	sum := funcParam(getSum, 50, 50)
	fmt.Println("sum=", sum)

	type myInt int
	var n1 myInt
	var n2 int
	n1 = 100
	n2 = int(n1)

	fmt.Println("n1=", n1, "n2=", n2)

	// go支持可变参数
	count := countSum(10, 1, 1, 1, 1, 1)
	fmt.Println("count=", count)

}
