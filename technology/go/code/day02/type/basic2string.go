package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 基本数据类型转string

	// 第一种方式： fmt.Sprintf
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var c byte = 'h'
	var str string // 空的str

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str = %q \n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str = %q \n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str = %q \n", str, str)

	str = fmt.Sprintf("%c", c)
	fmt.Printf("str type %T str = %q \n", str, str)

	// 第二种方式， strconv 函数
	var num3 int = 99
	var num4 float64 = 23.2343
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str = %q \n", str, str)

	// 说明：
	// ”f“: 格式，'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）
	// 10: 精度，小数点后面的个数
	// 64/32: 来源类型（32：float32、64：float64），会据此进行舍入。
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str = %q \n", str, str)

	str = strconv.FormatBool(b2);
	fmt.Printf("str type %T str = %q \n", str, str)


	// int to string
	str = strconv.Itoa(num3);
	fmt.Printf("str type %T str = %q \n", str, str)
}
