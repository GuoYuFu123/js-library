package main

import "fmt"

func main() {
	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)

	fmt.Printf("i=%v n1=%v n2=%v n3=%v \n", i, n1, n2, n3)

	// 高精度 -> 低精度 ,编译不会报错，但是转换的结构会按溢出处理，与期望值不一致
	var num1 int64 = 99999;
	var num2 int8 = int8(num1);
	fmt.Println(num1, num2)

	// 练习
	var v1 int32 = 12;
	var v2 int64
	var v3 int8
	// v2 = v1 + 20;  // int32 ---> int63 compile error
	// v3 = v1 + 20   // int32 ---> int8  compile error
	v2 = int64(v1) + 20;  
	v3 = int8(v1) + 20   
	fmt.Println(v2, v3)

	// v3 = int8(v1) + 127 //【编译通过，但是结果不是127+12，按溢出处理】
	// v3 = int8(v1) + 128 //【编译不通过，因为128大于int8的范围】
}
