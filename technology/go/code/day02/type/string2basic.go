package main

import (
	"fmt"
	"strconv"
)

// string 类型转基本类型
func main() {

	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b = %v \n", b, b)

	var str2 string = "12343"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T n1 = %v \n", n1, n1)
	fmt.Printf("n2 type %T n2 = %v \n", n2, n2)

	var str3 string = "123.1"
	var f1 float64
	var err error
	f1, err = strconv.ParseFloat(str3, 64)
	fmt.Println(err)
	fmt.Printf("f1 type %T f1 = %v \n", f1, f1)

	var t1 float32 = 24.234234
	var t2 float64 = float64(t1)
	fmt.Println("t1=", t1, "t2=", t2) // t1= 24.234234 t2= 24.234233856201172

	var t3 float64 = 141.234234123
	var t4 float32 = float32(t3) 
	fmt.Println("t3=", t3, "t4=", t4) // t3 1234.234234123 t4 1234.2343
}
