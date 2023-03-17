package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("hello")

	/******************* 整数类型, 默认是int ****************/

	// 有符号整数
	// int8范围 -128 ~ 127 【第一位是0|1表示符号，+0和-0，就把-0给了-128】
	// int16,int32,int64 以此类推
	var i int8 = -128
	fmt.Println("i=", i)

	// 无符号整数
	// uint8 0 ~ 255
	// uint16, uint32,uint64,以此类推
	var k uint8 = 0
	fmt.Println("k=", k)

	// int， uint，rune，byte
	var a int = 9800
	fmt.Println("a=", a)

	// 查看数据类型
	var aa = 100
	fmt.Printf("aa的默认type= %T ", aa)

	// 查看某个变量占用的字节大小
	var bb int64 = 10
	fmt.Printf("bb 的 type %T bb 的字节数是 %d", bb, unsafe.Sizeof(bb))

	/*********************** 浮点类型， 默认float64 ***********************************/
	// 浮点数都是有符号的
	// 浮点数 = 符号位 + 指数位 +尾数位
	// 使用浮点数，有可能会精度丢失，尾数位部分丢失，64位的精度比32位的精度更准确
	var f float32 = 12.12
	fmt.Println("f= ", f)

	var ff float64 = -123123123.12312312
	fmt.Println("ff", ff)

	var fff = 3.1232
	fmt.Printf("fff type = %T", fff)

	f2 := 5.23
	f3 := .123123
	fmt.Println("f2", f2, "f3", f3)
	f4 := 5.12e2  //  5.12 * 10的2次方
	f5 := 5.12e2  //  5.12 * 10的2次方
	f6 := 5.12e-2 //  5.12 / 10的2次方

	fmt.Println(f4, f5, f6)

	/******************* 字符类型 *************************************/
	// 传统的字符串，是通过一串固定的字符连接起来的字符序列，但是go的字符串是由单个字节连接起来的
	var b1 byte = 'a'
	var b2 byte = '0'
	fmt.Println(b1, b2) // 97 48
	fmt.Printf("b1 = %c b2= %c", b1, b2)

	var b3 int = '好'
	fmt.Printf("b3=%c  b3的码值= %d \n", b3, b3)

	var b4 = 10 + 'a'
	fmt.Println("b4=", b4)

	/********************  布尔类型  **************************************/

	var isBool = false
	fmt.Println("isbool=", isBool)
	fmt.Println("bool的占用空间 = ", unsafe.Sizeof(isBool))

	/********************  字符串类型   ************************/
	// 1 字符串一旦被賦值了，字符串就不能再修改啦，但是可以整体赋值
	// 2 字符串可以直接通过双引号进行赋值，也可以通过字符串魔板进行赋值，不会对内容进行转义，会原样输出
	var str string = "北京"
	// str[0] = "123"
	fmt.Println("str = ", str)

	var str2 = `
		package main; "牛逼"
	`
	fmt.Println("str2", str2)

	// 3 字符串拼接
	str3 := str + str2
	str3 += "hah"
	fmt.Println("str3=", str3)

	// 4 当一个拼接操作很长的时候，我们可以分行写，但是要注意，需要将操作符，留在上一行，表示这个语句还没有结束

	str4 := "hello" + "word" + "word" + "word" +
		"word" + "word" + "word" + "word" +
		"word" + "word"
	fmt.Println("str4=", str4)

}
