package main

import "fmt"

type Person struct {
	// 首字母大写可以被外界访问到
	Name string
	Age int
	Scores [5]float32
	ptr *int
	slice []int
	map1 map[string]string
}

func main() {
	fmt.Println("struct");

	// 定义结构体
	var p1 Person
	fmt.Println("p1",p1)

	if(p1.ptr == nil) {
		fmt.Println("p1.ptr is nil")
	}

	if(p1.slice == nil) {
		fmt.Println("p1.slice is nil")
	}

	if(p1.map1 == nil) {
		fmt.Println("p1.map1 is nil")
	}

	// 使用slice
	p1.slice = make([]int, 3)
	p1.slice[0] = 100;

	// 使用map
	p1.map1 = make(map[string]string, 0)
	p1.map1["key1"] = "value"
	fmt.Println("p1=", p1)

	// 复制，相互独立，但是设计到引用类型的，会修改地址，同步更新
	p2 := p1
	p2.Age = 200
	p2.map1["key1"] = "value2"
	fmt.Println("p1=", p1)
	fmt.Println("p2=" , p2)

}
