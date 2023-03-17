package main

import "fmt"

func main() {
	fmt.Println("switch")

	var key byte
	fmt.Println("输入a,b,c,d,e")
	// 指定格式来接收
	fmt.Scanf("%c", &key)

	// fallthrough 直接穿透下一层
	switch key {
	case 'a':
		fmt.Println("a")
		fallthrough
	case 'b':
		fmt.Println("b")
	case 'c', 'd':
		fmt.Println("c or d")
	default:
		fmt.Println("error")
	}

	// switch 后面不写表达式
	var age = 1
	switch {
	case age < 10:
		fmt.Println("age<10")
	default:
		fmt.Println("age>10")
	}

	// switch 变量后面定义一个变量，分号结束，不推荐
	switch grade := 90; {
	case grade > 90:
		fmt.Println(">90")
	default:
		fmt.Println("default")
	}

	// switch 类型判断
	var x interface{}
	var y = 10
	x = y

	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型~ ： %T", i)
	case int:
		fmt.Printf("int")
	case float32:
		fmt.Printf("float32")
	case float64:
		fmt.Printf("float64")
	case bool:
		fmt.Printf("bool")
	case func(int):
		fmt.Println("func(int)")
	default:
		fmt.Println("未知~")

	}
}
