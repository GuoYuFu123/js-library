package main

import (
	"errors"
	"fmt"
)

func testError() {
	// 使用defer + recover 来捕获和处理异常

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err=", err)
			// 处理异常
			fmt.Println("send email")
		}
	}()
	num1 := 100
	num2 := 0
	res := num1 / num2
	fmt.Print("res", res)
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		// 返回一个自定义错误
		return errors.New("读取文件失败")
	}
}

func test02() {
	err := readConf("config.ini")
	if err != nil {
		// 抛出异常，终止程序
		panic(err)
	}
	fmt.Print("test02 done")
}

func main() {
	fmt.Println("error handle")
	// testError()

	test02()
}  
