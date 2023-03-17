package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("file")

	// 打开文件
	// 概念说明： file 的语法
	// 1、file 叫 file 对象
	// 2、file 叫 file 指针
	// 3、file 叫 file 文件句柄
	file, err := os.Open("./1.txt")
	if err != nil {
		fmt.Println("open file err", err)
	}
	// 输出文件，file就是一个指针
	fmt.Printf("file=%v", file)
	// 关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err", err)
	}
}
