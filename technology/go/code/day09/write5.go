package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 编程一个程序，将一个文件的内容，写入到另外一个文件。注:这两个文件已经存在了. 说明:使用 ioutil.ReadFile / ioutil.WriteFile 完成写文件的任务.
	fmt.Println("方式2")

	// 将w1写入w2
	file1Path := "./w1.txt"
	file2Path := "./w2.txt"

	content, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Println("file open err", err)
		return
	}
	err = ioutil.WriteFile(file2Path, content, 0666)
	if err != nil {
		fmt.Println("file write err", err)
	}

}
