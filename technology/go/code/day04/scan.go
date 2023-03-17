package main

import "fmt"

func main() {
	fmt.Println("scan")

	var name string
	var age byte
	var sal float32
	var isPass bool

	// Scanln: 通过换行来进行获取输入

	// fmt.Println("请输入name ")
	// fmt.Scanln(&name)

	// fmt.Println("请输入age ")
	// fmt.Scanln(&age)

	// fmt.Println("请输入sal ")
	// fmt.Scanln(&sal)

	// fmt.Println("请输入isPass ")
	// fmt.Scanln(&isPass)

	// fmt.Printf("name= %v \nage= %v \nsal= %v \nisPass= %v", name, age, sal, isPass)

	// Scan: 通过格式化输入，通过空格来区分
	fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试，通过空格进行分隔")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("name= %v \nage= %v \nsal= %v \nisPass= %v", name, age, sal, isPass)

}
