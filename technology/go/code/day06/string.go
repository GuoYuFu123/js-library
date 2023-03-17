package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 统计字符串的长度，按字节 len(str)
	// golang的编码统一为utf-8（ASCII的字符(字母和数字占一个字节，汉字占3个字节)）
	str := "cat eat fish哈"
	fmt.Println("len=", len(str))

	// 字符串遍历
	str2 := "hello北京"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符串= %c \n", r[i])
	}

	// 字符串转整数
	// Atoi是ParseInt(s, 10, 0)的简写。
	n, err := strconv.Atoi("123a")
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("n=", n)
	}

	// 整数转字符串
	// Itoa是FormatInt(i, 10) 的简写。
	str = strconv.Itoa(1234)
	fmt.Println("str", str)

	// 字符串转切片字节
	bytes := []byte("hello我")
	fmt.Printf("byte=%v\n", bytes)

	str = string([]byte{97, 98, 99})
	fmt.Println("str=", str)

	// 10 进制 转 2 8 16进制
	str = strconv.FormatInt(10, 2)
	fmt.Println("str 2 = ", str)
	str = strconv.FormatInt(10, 8)
	fmt.Println("str 8 = ", str)
	str = strconv.FormatInt(10, 16)
	fmt.Println("str 16 = ", str)

	// 查找字符串中是否存在子串
	isExist := strings.Contains("hello word", "hello")
	fmt.Println("isExist=", isExist)

	// 统计一个字符串中有几个字符
	num := strings.Count("hello hello", "e")
	fmt.Println("num=", num)

	// 比较字符串大小写
	b := strings.EqualFold("abc", "ABC")
	fmt.Println("b=", b)
	fmt.Println("abc == ABC :", "acb" == "ABC")

	// 返回指定字符串第一次出现的位置
	// index or -1
	index := strings.Index("hello_abc", "a1bc")
	fmt.Println("index=", index)

	// 查找字符串末次的位置
	// index or -1
	index = strings.LastIndex("hello, hello", "o")
	fmt.Println("index=", index)

	// 替换指定字符串
	str2 = strings.Replace("go,go,hello", "go", "to", 1)
	fmt.Println("replace=", str2)
	// 全部替换
	str2 = strings.Replace("go,go,hello", "go", "to", -1)
	fmt.Println("replace=", str2)
	str2 = strings.ReplaceAll("go,go,hello", "go", "to")
	fmt.Println("replace=", str2)

	// 字符串分隔
	strArr := strings.Split("hello,hello,wolrd", ",")
	fmt.Println("strArr=", strArr)

	// 转换大小写
	str = strings.ToLower("hello,WORLD")
	fmt.Println("lower=", str)
	str = strings.ToUpper("hello,WORLD")
	fmt.Println("upper=", str)

	// 左右字符串空格去掉
	str = strings.TrimSpace(" hello world")
	fmt.Println("str=", str)

	// 将左右指定字符串去掉
	str = strings.Trim("hello helloh", "h")
	fmt.Println(str)

	// 将左侧指定字符串去掉
	str = strings.TrimLeft("hello helloh", "h")
	fmt.Println(str)

	// 将右侧指定字符串去掉
	str = strings.TrimRight("hello helloh", "h")
	fmt.Println(str)

	// 指定字段开头
	b = strings.HasPrefix("https://xxx", "https")
	fmt.Println("b=", b)
	// 指定字段结尾
	b = strings.HasSuffix("xxx.jpg", "jpg")
	fmt.Println("b=", b)
}
