package utils //储存工具性质的函数
import (
	"fmt"
	"net"
)

//这里直接复制了server\utils\utils.go

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
)

//因为使用了分层加面向对象的思路，所以需要先将这些方法关联到结构体中
//用到的时候，直接创建实例，调用方法
//(其他文件中同理）
type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn   //链接
	Buf  [8096]byte //传输时，使用的缓冲
}

//读消息
func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	//buf := make([]byte, 8096)   //Transfer自带，这边就不用另外写了
	fmt.Println("读取客户端发送的数据...")
	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了conn，则不会阻塞
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil { //这边返回的err为io.EOF
		//err = errors.New("read pkg header error") //自定义错误
		return
	}

	//因为Read()括号内写入类型的原因，这边需要根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	//判断消息长度是否错误（是否丢包），以及是否有其他错误
	if n != int(pkgLen) || err != nil { //这边返回的err为io.EOF
		//err = errors.New("read pkg body error") //自定义错误
		return
	}

	//把pkgLen反序列化成->message.Message
	json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {

		fmt.Println("json Unmarshal err=", err)
		return
	}

	return
}

//写消息
func (this *Transfer) WritePkg(data []byte) (err error) {

	//发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail=", err)
		return
	}

	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail=", err)
		return
	}
	return
}
