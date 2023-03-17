package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("cmd args flag")

	// 定义结构变量
	var user string
	var pwd string
	var host string
	var port int

	// &user, 就能接收用户命令行总 -u 后面的参数
	// u , 就是命令行中的 -u
	// "", 默认值
	// "用户默认值"  说明
	flag.StringVar(&user, "u", "", "用户名")
	flag.StringVar(&pwd, "pwd", "", "密码")
	flag.StringVar(&host, "host","localhost", "主机")
	flag.IntVar(&port, "port", 3306, "port")

	flag.Parse()

	fmt.Println(user,pwd,host,port)

}
