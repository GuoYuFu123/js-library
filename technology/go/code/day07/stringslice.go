package main

import "fmt"

func main() {
	// string 底层是一个byte数组，因此string也可以是进行切片处理
	str := "hello@guoguo";
	// 使用slice获取到guoguo
	slice := str[6:]
	fmt.Println("slice=",slice)

	// string 是不可变额，也就是说不能通过 str[0]='z' 方式来修改字符串
	// 如果需要修改字符串，可以先将 string-> []byte / []rune -> 修改 -> 重写为string
	arr1 := []byte(str);
	arr1[0] ='A';
	str = string(arr1)
	fmt.Println("str=",str)

	// 细节：我们转为 []byte 后，可以处理英文和数字，但是不能处理中文
	// 原因，[]byte 是通过字节来处理的，而一个汉字是3个字节，因此会出现乱码
	// 可以转为 []rune, 是通过字符来处理

	arr2 := []rune(str);
	arr2[0] ='爱';
	str = string(arr2)
	fmt.Println("str=",str)
} 