package main

import "fmt"

type Person struct {
	Name string
}

// 方法是绑定在自定义的类型上的
func (p Person) eat() {
	fmt.Println("p.name", p.Name)
}
func (p Person) getSun(n1 int, n2 int) int {
	fmt.Println("p.name getSun", p.Name)
	return n1 + n2
}

func main() {
	fmt.Println("method")

	p := Person{"张三"}
	p.eat()
	p.getSun(1, 2)
}
