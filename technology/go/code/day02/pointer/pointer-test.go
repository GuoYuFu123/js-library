package main

import "fmt"


func main(){
	fmt.Println("---pointer---")
	// 值类型
	var num int = 9;
	// 获取当前值的内存的地址
	fmt.Println(&num) // // 0xc000194008 

	var ptr *int = &num;

	*ptr = 10

	fmt.Println("*ptr=", *ptr, "num=", num)

}