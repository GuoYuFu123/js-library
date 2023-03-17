package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)

	lock sync.Mutex
)

func test(n int) {
	
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	//这里我们将 res 放入到 myMap
	myMap[n] = res //concurrent map writes?
	lock.Unlock()
}
func main() {
	fmt.Println("factorial")

	for i := 1; i <= 200; i++ {	
		go test(i)
	}

	//休眠 10 秒钟【第二个问题 】
	time.Sleep(time.Second * 10)

	//这里我们输出结果,变量这个结果
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()

}
