package main

import "fmt"

func main() {
	fmt.Println("for")

	// 1
	for i := 0; i < 5; i++ {
		fmt.Printf("i= %v \n", i)
	}

	// 2
	j := 0
	for j < 5 {
		fmt.Printf("j= %v \n", j)
		j++
	}

	// 3
	h := 0
	for {
		if h >= 5 {
			break
		}
		fmt.Printf("h= %v \n", h)
		h++
	}

	// 字符串遍历
	var str string = "hello, guoguo"
	for i := 0; i < len(str); i++ {
		fmt.Printf("i = %v char = %c \n", i, str[i])
	}

	// for-range
	str = "for range"
	for k, v := range str {
		fmt.Printf("k=%v v=%c \n", k, v)
	}

	// 字符串有汉字，因为go是utf-8的编码，一个汉字3个字节
	// 传统的循环式根据字节来进行便利的，会乱码
	str = "hello, 中国"
	for i := 0; i < len(str); i++ {
		fmt.Printf("i = %v char = %c \n", i, str[i])
	}

	//  []rune(s), 它可以将字符串转化成 unicode 码点
	//  []byte(s) 和 []rune(s) 的区别： byte 表示一个字节，rune 表示四个字节
	str2 := []rune(str);
	for i := 0; i < len(str2); i++ {
		fmt.Printf("i = %v char = %c \n", i, str2[i])
	}


	// for-range是根据字符来遍历的
	for k, v := range str {
		fmt.Printf("k=%v v=%c \n", k, v)
	}


	

}
