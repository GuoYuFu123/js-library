package main

import "fmt"

func main() {
	fmt.Println("chan")

	var intchan chan int
	intchan = make(chan int, 3)

	intchan <- 10
	intchan <- 20
	intchan <- 30

	num1 := <-intchan
	num2 := <-intchan
	// num3 := <-intchan

	// 1、不能使用普通的for循环
	// 2、在遍历时，如果 channel 没有关闭，则回出现 deadlock 的错误
	// 3、在遍历时，如果 channel 已经关闭，则会正常遍历数据，遍历完后，就会退出遍历。
	close(intchan)
	for v := range intchan {
		fmt.Println("v", v)
	}
	fmt.Println(num1, num2, intchan)
}
