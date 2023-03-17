package main

import (
	"bufio"
	"fmt"
	"io"
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
	// 当函数退出的时，要即使的关闭file
	defer file.Close()

	// 创建一个 *Reader, 是带缓冲的
	/**
	const (
		defaultBufSize = 4096
	)
	*/
	reader := bufio.NewReader(file)
	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n') // 读到界限符就换行
		fmt.Print(str)
		if err == io.EOF {                  // io.EOF 标识文件的末尾
			break
		}
		
	}
	fmt.Println("\n文件读取结束")

}
