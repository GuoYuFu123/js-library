package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

/**
1) 在主线程(可以理解成进程)中，开启一个 goroutine, 该协程每隔 1 秒输出 "hello,world"
2) 在主线程中也每隔一秒输出"hello,golang", 输出 10 次后，退出程序
3) 要求主线程和 goroutine 同时执行.
*/

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("goroutine")

	// 获取当前的cpu数量
	num := runtime.NumCPU()
	// 设置CPU的数量
	runtime.GOMAXPROCS(num - 1)
	fmt.Println("num", num)

	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
