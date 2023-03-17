package main

import "fmt"

func main() {
	fmt.Println("chan")

	var intchan chan int
	intchan = make(chan int, 3)

	intchan <- 10
	intchan <- 20
	intchan <- 30

	num1,ok := <-intchan
	num2 := <-intchan
	num3 := <-intchan
	fmt.Println(num1,ok, num2, num3, intchan)
}
