package main

import "fmt"

func main() {
	fmt.Println("assert")

	var x interface{}
	var b2 float32 = 1.1
	x = b2 // 空接口，可以接收任意类型

	//  x => float32
	y := x.(float32)
	fmt.Println("y", y)

	// x => float64 error
	z, ok := x.(float64)
	if ok == true {
		fmt.Println("z", z)
	} else {
		fmt.Println("convert error")
	}
	fmt.Println("continue")

}
