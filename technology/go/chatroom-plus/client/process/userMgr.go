package process

import (
	"chatroom/client/model"
	"chatroom/common/message"
	"fmt"
)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

//因为在客户端，很多地方会用到curUser,我们将其作为一个全局的
var CurUser model.CurUser //我们在用户登录成功（userProcess.go中）后，完成对CurUser的初始化

//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历一遍onlineUsers
	fmt.Println("当前在线用户列表:")
	for id := range onlineUsers {
		//如果不显示自己
		fmt.Println("用户id:\t", id)
	}
}

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	//适当优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

}
