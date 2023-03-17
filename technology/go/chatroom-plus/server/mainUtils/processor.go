package mainUtils

import (
	"chatroom/common/message"
	process2 "chatroom/server/process" //如果文件夹名和包名不一样 导入时路径前面会自动加一个包名
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

//此文件总控服务器,完成调度

//先创建一个Processor的结构体
type Processor struct {
	Conn net.Conn
}

//处理消息
//编写一个ServerProcessMes函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {

	//看看是否能接收到客户端发送的群发消息
	fmt.Println("mes=", mes)

	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		//创建一个UserProcess实例
		up := &process2.UserProcess{ //这边因为调用process2包和定义process2函数都在这个源文件里，所以不冲突？
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		//创建一个SmsProcess实例完成转发群聊消息
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) Process2() (err error) { //如果被同一个包里的其他地方引用，首写小写就行，如果被其他包引用，首写必须大写
	//循环地读客户端发送的信息
	for {

		//这里我们将读取数据包，直接封装成一个函数readPkg(),返回Message，Err
		//创建一个Transfer实例，完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()

		if err != nil {
			if err == io.EOF { //即readPkg中conn.Read的err
				fmt.Println("客户端退出，服务器也正常退出...")
				return err
			} else {
				fmt.Println("readPkg err=", err) //其他错误信息
				return err
			}
		}
		//fmt.Println("mes=", mes)

		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
