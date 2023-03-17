package main

import "fmt"

func main() {
	fmt.Println("查找")

	name := "i"
	names := [3]string{"i", "love", "you"}
	// 第一种方式
	for i := 0; i < len(names); i++ {
		if name == names[i] {
			fmt.Println("index = ", i)
			break;
		} else if i ==len(names) -1{
			fmt.Println("no find")
		}
	}

	// 第二种方式
	index := -1;
	for i := 0; i < len(names); i++ {
		if name == names[i] {
			index = i
			break;
		} 
	}
	if(index != -1){
		fmt.Println("index = ", index)
	}
}

