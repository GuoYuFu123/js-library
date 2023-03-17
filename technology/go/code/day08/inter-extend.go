package main

import "fmt"

// 结构体
type Monkey struct {
	Name string
}

// interface
type BridAble interface {
	Flying()
}
type FishAble interface {
	Swimming()
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树~")
}

func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "通过学习，会飞~")
}

func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "通过学习，会游泳~")
}

// 继承猴子的小猴子
type LittleMonkey struct {
	Monkey
}

func main() {
	fmt.Println("inter-vs-extend")

	littleMonkey := LittleMonkey{
		Monkey: Monkey{
			Name: "悟空",
		},
	}
	littleMonkey.climbing()
	littleMonkey.Flying()
	littleMonkey.Swimming()
}
