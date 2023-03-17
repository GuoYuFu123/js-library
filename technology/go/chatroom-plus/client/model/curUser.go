package model

import (
	"chatroom/common/message"
	"net"
)

type CurUser struct { //当前发送消息的用户
	Conn         net.Conn
	message.User //匿名字段
}

var MyCurUser CurUser

// 由于CurUSer在许多处都会使用到，有且只用改一个变量
// 故将其定义成一个全局变量
// CurUser的作用实际上是用来管理和维护客户信息的
