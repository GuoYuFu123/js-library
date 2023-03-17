package main

import "fmt"

// ifelse
func main() {
	// 双分支
	var age int = 20

	if age >= 18 {
		fmt.Println("成年人")
	} else {
		fmt.Println("未成年人")
	}

	// go 支持在if中直接定义一个变量，比如下面
	if age2 := 0; age2 > 18 {
		fmt.Println("成年人")
	} else {
		fmt.Println("未成年人")
	}

	if age >= 18 {
		fmt.Println("成年人")
		if sex := "男"; sex == "男" {
			fmt.Println("男人")
		} else {
			fmt.Println("女人")
		}
	} else {
		fmt.Println("未成年人")
	}
}
