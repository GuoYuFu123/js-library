package main

import "fmt"

func main() {
	fmt.Println("copy")

	// 使用 append 内置函数，对slice进行动态追加
	// 内建函数append将元素追加到切片的末尾。若它有足够的容量，其目标就会重新切片以容纳新的元素。否则，就会分配一个新的基本数组。append返回更新后的切片，因此必须存储追加后的结果。

	var slice []int = []int{1, 2, 3, 4, 5}
	fmt.Println("slice=", slice)
	slice = append(slice, 6, 7, 8)
	fmt.Println("append slice=", slice)

	slice = append(slice, slice...)
	fmt.Println("append slice  =", slice)

	// 切片的拷贝
	// 内建函数copy将元素从来源切片复制到目标切片中，也能将字节从字符串复制到字节切片中。copy返回被复制的元素数量，它会是 len(src) 和 len(dst) 中较小的那个。来源和目标的底层内存可以重叠。
	var slice2 []int = []int{100, 200, 300}
	var slice3 = make([]int, 5, 10)
	copy(slice3, slice2)
	fmt.Println("copy slice2 =", slice2)
	fmt.Println("copy slice3 =", slice3)

}
