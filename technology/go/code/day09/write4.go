package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开一个存在的文件，将原来的内容读出显示在终端，并且追加 5 句"hello,北京!"
	fmt.Println("方式1")

	file, err := os.OpenFile("./w1.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	// 及时关闭句柄
	defer file.Close()

	// 读取内容，显示终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF {
			break
		}
	}

	str := "hello shandong \n"
	// 使用带缓冲的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// 因为 writer 带缓冲，因为调用WriteString方法时，其实是写入缓冲中的，所以我们需要调用Flush方法，将缓冲中的数据写入文件
	writer.Flush() // 刷入文件

}
