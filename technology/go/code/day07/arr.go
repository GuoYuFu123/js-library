package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("arr")
	// 4种数组初始化

	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr1)

	var arr2 = [3]int{1, 2, 3}
	fmt.Println(arr2)

	// [...]是规定的写法
	var arr3 = [...]int{1, 2, 3}
	fmt.Println(arr3)

	var arr4 = [...]int{1: 800, 0: 100, 2: 900}
	fmt.Println(arr4)

	// 类型推导
	arr5 := [...]string{1: "tom", 0: "jack"}
	fmt.Println("strarr5", arr5)

	arr6 := [5]int{0, 1, 2, 3, 4}

	for i := 0; i < len(arr6); i++ {
		fmt.Println("index=", i, "vale=", arr6[i])
	}

	for index, value := range arr6 {
		fmt.Println("index=", index, "vale=", value)
	}

	var test = func(arr *[3]int) {
		(*arr)[0] = 10
	}
	arr7 := [3]int{1, 2, 3}
	fmt.Print("arr7", arr7)
	test(&arr7)
	fmt.Println("arr7", arr7)

	// 随机生成5个数，并反转打印，
	var arr8 [5]int
	lenth := len(arr8)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < lenth; i++ {
		arr8[i] = rand.Intn(100)
	}
	fmt.Println("交换前=", arr8)
	// 开始交换
	temp :=0;
	for i := 0; i < lenth/2; i++ {
		temp = arr8[lenth-1-i]
		arr8[lenth-1-i] = arr8[i]
		arr8[i] = temp
	}
	fmt.Println("交换后=", arr8)

}
