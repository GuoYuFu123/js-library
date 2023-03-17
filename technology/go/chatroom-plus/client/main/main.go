package main

import (
	"chatroom/client/process"
	"fmt"
	"os"
)

//定义两个变量，一个表示用户id, 一个表示用户密码
var userId int
var userPwd string
var userName string

func main() {

	//接收用户的选择
	var key int
	//判断是否继续显示菜单
	//var loop = true

	for true { //即输入有误时，继续显示该菜单
		fmt.Println("----------------欢迎登陆多人聊天系统------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)

			//完成登录
			//1.创建一个UserProcess的实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)

		//loop = false //不继续显示该菜单  //已经把最初的两级菜单合并到这了，所以此行没意义，可注释掉
		//然后 登录成功的二级菜单到了process\server.go里
		case 2:
			fmt.Println("注册用户")
			//loop = false //不继续显示该菜单
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名字(nickname):")
			fmt.Scanf("%s\n", &userName)
			//2.调用UserProcess，完成注册的请求,注册由服务器完成
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)

		case 3:
			fmt.Println("退出系统")
			//loop = false //不继续显示该菜单
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

	////根据用户的输入，显示新的提示信息
	//if key == 1 {
	//	//说明用户要登录
	//	//fmt.Println("请输入用户的id")
	//	//fmt.Scanf("%d\n", &userId)
	//	//fmt.Println("请输入用户的密码")
	//	//fmt.Scanf("%s\n", &userPwd)
	//
	//	//先把登录的函数（协议），写到另外一个文件，比如login.go
	//	//因为都在一个包（package main)下，所以可以直接调用
	//	//这里会需要重新调用
	//	login(userId, userPwd)
	//	//if err != nil {
	//	//	fmt.Println("登陆失败")
	//	//} else {
	//	//	fmt.Println("登陆成功")
	//	//}
	//} else if key == 2 {
	//	fmt.Println("进行用户注册的逻辑...")
	//}
}
