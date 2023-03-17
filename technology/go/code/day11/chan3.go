package main

import "fmt"

type Cat struct {
	Name string
}

func main() {
	fmt.Println("chan")

	var allchan chan interface{}
	allchan = make(chan interface{}, 3)

	cat1 := Cat{
		"tom",
	}
	cat2 := Cat{
		"tom2",
	}

	allchan <- cat1
	allchan <- cat2
	// 关闭之后，不能再写入数据到chan
	// close(allchan)
	// allchan <- 30

	// 但是可以继续读取数据
	cat11 := <-allchan
	cat22 := <-allchan
	num3 := <-allchan

	// 使用类型断言
	fmt.Println(cat11.(Cat).Name, cat22, num3, allchan)
}
