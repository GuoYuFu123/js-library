package main

import "fmt"

type Student struct {
	Name string
}

// 编写一个函数，可以判断输入的参数是什么类型
func TypeJudge(items ...interface{}) {
	for _, x := range items {
		switch x.(type) {
		case bool:
			fmt.Println("bool", x)
		case float32:
			fmt.Println("float32", x)
		case float64:
			fmt.Println("float64", x)
		case int, int32, int64:
			fmt.Println("int, int32,int64", x)
		case string:
			fmt.Println("string", x)
		case Student:
			fmt.Println("Student", x)
		case *Student:
			fmt.Println("*Student", x)
		default:
			fmt.Println("type 不确定~")
		}
	}
}

func main() {
	fmt.Println("interface")
	var n1 float32 = 1.1
	var n2 float64 = 1.2
	var n3 int16 = 1
	var n4 int32 = 2
	var s string = "sss"
	var stu = Student{"guo"}

	TypeJudge(n1, n2, n3, n4, s, stu, &stu)

}
