package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("init")

	// 第一种方式
	var p1 Person
	p1.Name = "p1"
	fmt.Println("p1", p1)

	// 第二种方式
	p2 := Person{Name: "p2", Age: 20}
	fmt.Println("p2", p2)

	// 第三种方式
	// p3是结构体指针
	// var p3 *Person = new(Person)
	var p3 *Person = new(Person)
	// 因为p3是一个指针，因此标准的给字段赋值方式
	// (*p3).Name = "aaa" 也可以 p3.Name = "bbb"
	// 原因，因为go的设计者，方便程序员使用方便，底层会对 p3.Name 进行处理， 加上取值运算 (*p3).Name
	(*p3).Name = "aaa"
	p3.Name = "bbb"
	fmt.Println("p3", *p3)

	// 第四种方式，同3
	var p4 *Person = &Person{}
	(*p4).Name = "aaa"
	p4.Name = "bbb"
	fmt.Println("p4", *p4)

	// json 序列化
	jsonStr, err := json.Marshal(p4)
	if err != nil {
		fmt.Println("json err")
	}
	fmt.Println("jsonStr", string(jsonStr))

}
