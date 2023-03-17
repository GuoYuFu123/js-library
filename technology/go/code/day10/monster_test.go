package main

import (
	"../utils"
	"testing"
)

// 测试用例,测试 Store 方法
func TestStore(t *testing.T) {
	// 先创建一个 Monster 实例
	monsterr := &utils.Monsterr{
		Name:  "红孩儿",
		Age:   10,
		Skill: "吐火.",
	}
	res := monsterr.Store()
	if !res {
		t.Fatalf("monsterr.Store() 错误，希望为=%v 实际为=%v", true, res)
	}

	t.Logf("monsterr.Store() 测试成功!")
}

func TestReStore(t *testing.T) { //测试数据是很多，测试很多次，才确定函数，模块..
	//先创建一个 Monster 实例 ， 不需要指定字段的值
	var monsterr = &utils.Monsterr{}
	res := monsterr.ReStore()
	if !res {
		t.Fatalf("monsterr.ReStore() 错误，希望为=%v 实际为=%v", true, res)
	}
	// 进一步判断
	if monsterr.Name != "红孩儿" {
		t.Fatalf("monsterr.ReStore() 错误，希望为=%v 实际为=%v", "红孩儿", monsterr.Name)
	}
	t.Logf("monsterr.ReStore() 测试成功!")
}
