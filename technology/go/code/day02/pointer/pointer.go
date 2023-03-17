package main

import "fmt"


func main(){
	fmt.Println("---pointer---")
	// 值类型
	var num int = 3;
	// 获取当前值的内存的地址
	fmt.Println(&num) // // 0xc000194008 

	// 初始一个指针，指向num的内存的地址
	var ptr *int = &num;

	// 指针可以认为是一个存《内存地址》的变量，它本身也有自己的内存地址，取值用*ptr进行取值
	fmt.Println(ptr, &ptr, *ptr) // 0xc0000ac008 0xc0000a6020 3
	fmt.Printf("ptr type = %T",ptr)
}