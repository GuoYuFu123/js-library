package main

import (
	"fmt"
	"sort"
)

type Stu struct {
	name string
}

func main() {
	fmt.Println("map")
	// 声明map变量
	// var a map[string]int
	// 在使用map时，需要先进行make，make的作用就是给map分配数据空间
	a := make(map[string]int, 10)
	a["a"] = 1
	a["a"] = 2
	fmt.Println("a", a)
	fmt.Println("a len", len(a)) // 1

	// 第一种方式
	var m1 map[string]int
	m1 = make(map[string]int)
	m1["m11"] = 1
	m1["m12"] = 2
	fmt.Println("m1", m1)

	// 第二种方式
	m2 := make(map[string]int)
	m2["m21"] = 1
	m2["m22"] = 2
	fmt.Println("m2", m2)
	// 第三种方式
	m3 := map[string]int{
		"m31": 1,
		"m32": 2,
	}
	fmt.Println("m3", m3)

	// 测试
	m4 := make(map[string]map[string]string)
	m4["小明"] = make(map[string]string, 1)
	m4["小明"]["name"] = "xiaoming"
	m4["小红"] = make(map[string]string, 1)
	m4["小红"]["name"] = "小红"

	fmt.Println("m4=", m4)

	m5 := map[string]int{
		"m51": 1,
		"m52": 2,
	}
	fmt.Println("m5", m5)
	m5["m53"] = 3
	fmt.Println("m5", m5)
	m5["m51"] = 100
	fmt.Println("m5", m5)
	delete(m5, "m51")
	// 当delete指定的key不存在时，删除不会操作，也不会报错
	fmt.Println("m5=", m5)

	// 如果我们希望一次性删除所有的key
	// 1、遍历所有的key，如何逐一删除【遍历】
	// 2、直接make一个新的空间
	m5 = make(map[string]int)

	// key的查找
	val, ok := m5["m52"]
	if ok {
		fmt.Println("val", val)
	} else {
		fmt.Println("没有key【m51】")
	}

	// map的遍历
	for k, v := range m5 {
		fmt.Println("k=", k, "v=", v)
	}

	monsters := make([]map[string]string, 2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "100"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "牛魔王2"
		monsters[1]["age"] = "1200"
	}
	// 数组越界
	// if monsters[2] ==nil {
	// 	monsters[2] = make(map[string]string,2)
	// 	monsters[2]["name"] = "牛魔王3"
	// 	monsters[2]["age"] = "2200"
	// }

	// 需要动态append

	mewMonster := map[string]string{
		"name": "孙悟空",
		"age":  "10000",
	}

	monsters = append(monsters, mewMonster)

	fmt.Println("monsters", monsters)

	// golang中没有一个专门的方针针对map的key进行排序
	// golang中的map默认是无序的，注意也不是按照添加的顺序存放的，所以遍历出来的结果可能不一致
	// golang中的map排序，是先将key进行排序，然后根据key的值遍历输出即可
	m6 := make(map[int]int, 10)
	m6[10] = 100
	m6[1] = 1
	m6[4] = 4
	m6[6] = 6
	fmt.Println("m6", m6)
	

	// map的key数组
	var keys []int;
	for k ,_ := range m6 {
		keys = append(keys, k)
	}
	// 排序
	sort.Ints(keys)
	fmt.Println("keys", keys)

	for _, k := range keys {
		fmt.Println("key= ", k, "v=", m6[k])
	}



	stus := make(map[string]Stu , 10)

	stu1:=Stu{"大明"}
	stu2:=Stu{"小明"}
	stus["1"] = stu1
	stus["2"] = stu2

	for k, v := range stus {
		fmt.Println("key= ", k, "v=", v)
	}
}
