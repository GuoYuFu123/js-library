package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建一个新文件，写入5句 hello world
	fmt.Println("方式1")

	file, err := os.OpenFile("./w1.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	// 及时关闭句柄
	defer file.Close()
	str := "hello world \n"
	// 使用带缓冲的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// 因为 writer 带缓冲，因为调用WriteString方法时，其实是写入缓冲中的，所以我们需要调用Flush方法，将缓冲中的数据写入文件
	writer.Flush(); // 刷入文件

}
