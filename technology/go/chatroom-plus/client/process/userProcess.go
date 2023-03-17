package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
	//暂时不需要任何字段
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {

	//1.链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3.创建一个LoginMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.将registerMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//6.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes进行序列化,data即客户端需要发送的数据
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}

	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=", err)
	}

	//读返回信息
	mes, err = tf.ReadPkg() //mes就是RegisterResMes

	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将mes的Data部分反序列化成RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，你可重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

//原client\login.go 直接复制过来

//关联一个用户登录的方法
//写一个函数，完成登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	////开始定协议
	//fmt.Printf("userId=%d userPwd=%s\n", userId, userPwd)
	//
	//return nil

	//1.链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	//此处还没有进行序列化，所以首字母大写
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5.把data赋给mes.Data字段
	mes.Data = string(data) //因为上面Marshal处理得到的data为byte型切片，所以需要转化为string类型

	//6.将mes进行序列化,data即客户端需要发送的数据
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//7.1 先把data的长度发送给服务器
	//先获取到data的长度->转成一个表示长度的byte切片
	//因为Write()括号内写入类型的问题，这边要先做一个转换
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil { //四个字节
		fmt.Println("conn.Write(buf) fail", err)
		return
	}

	fmt.Printf("客户端，发送消息的长度%d,内容=%s", len(data), string(data))

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}

	//休眠20s
	//time.Sleep(10 * time.Second)
	//fmt.Println("休眠了20...")

	//这里还需要处理服务器端返回的消息

	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()

	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//初始化CurUser，即当前发送消息的用户
		CurUser.Conn = conn
		CurUser.UserId = userId //因为CurUser.User定义为了匿名字段，所以可以直接调用UserId
		CurUser.UserStatus = message.UserOnline

		//fmt.Println("登录成功")

		//可以显示当前在线用户列标配，遍历loginResMes.UserId
		fmt.Println("当前在线用户列表如下:")
		for _, v := range loginResMes.UsersId {

			//如果我们要求不显示自己在线，可以增加以下代码
			if v == userId {
				continue
			}

			fmt.Println("用户id\t", v)

			//完成客户端的onlineUsers的初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")

		//这里我们还需要在客户端启动一个协程
		//该协程保持和服务器端的通讯，如果服务器有数据推送给客户端
		//则接收并显示在客户端的终端
		go serverProcessMes(conn)

		//1.显示我们登录成功的菜单(循环显示)...
		for {
			ShowMenu() //因为userProcess.go和server.go在同一个包里，所以可以直接调
		}

	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}
