package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("file")

	// 使用 ioutil.ReadFile 一次性将文件读取到位
	content, err := ioutil.ReadFile("./1.txt")
	if err != nil {
		fmt.Println("readfile err", err)
	}
	// 把读取到内容显示在终端
	//content =  []uint8%
	fmt.Printf("%v %T", string(content), content)

	// 我们没有显示的open文件，所以我们也不需要显示的close文件
	// 因为，文件的 open 和 close 被封装在 readfile 函数内部啦
}
