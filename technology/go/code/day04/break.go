package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("break")

	n := 0

	for {
		rand.Seed(time.Now().UnixMicro())
		num := rand.Intn(100)
		n++
		fmt.Println("num=", num, "\tn=", n)
		if num == 99 {
			break
		}

	}


	// 多层循环嵌套，可以通过标签指定终止哪个标签
	label1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println("i=", i, "\tj=", j)
			if j == 2 {
				break label1
			}
		}
	}

}
