package main

import (
	"fmt"
)

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("write", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("read", v)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	fmt.Println("边读边写")

	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	
	// close(intChan)
	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		fmt.Println(ok)
		if !ok {
			break
		}
	}
}
