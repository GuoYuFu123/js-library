相信很多人都遇到过登陆了SSH后想要去查找相关命令的资料，过了一会回到命令行下时发现SSH已经失去了响应，必须重新连接登陆才行。这个就很烦诶。上网查了一下，原来这是SSH默认的安全设置，经过一定的时间客户端未发送数据就断开连接。个人感觉这个没啥用，就想把它关掉，方法嘛，可以在服务器端设置间隔一定的时间发送信号给客户端，如果客户端响应了这个信号，就保持连接。如果客户端连续几次都没有响应服务端发来的信号，服务端就断开连接。通过设置间隔时间的长短和未响应的次数就可以实现一直保持连接。具体的步骤如下：

###### 配置

- 首先连上服务器，以root身份登陆，输入下面这条命令（如果没有vim 可以用自带的vi 或者下载安装vim）

```undefined
vim /etc/ssh/sshd_config   //打开SSH的配置文件
```

- 看一看sshd_config 中有没有：

  ```undefined
  TCPKeepAlive yes  #表示TCP保持连接不断开
  ```

  如果没有，就在新的一行中加入上面这条命令(# 号后面的是注释，可以不用添加进去)

- 继续加入以下代码：

  ```undefined
  ClientAliveInterval 60
  ClientAliveCountMax 3
  ```

  其中，ClientAliveInterval 60 的意思是：服务端向客户端发送信号的间隔，单位是秒。例如我这里指定的是60，就是服务端每隔60秒就向客户端发送一个信号，如果客户端响应了这个信号，就保持连接。

  ClientAliveCountMax 3 的意思是客户端未响应服务端发来的信号的最大次数。例如我这里指定的是3次，那么客户端如果超过3次没有响应服务端发来的信号，就会断开与服务端的连接。

###### 结束

好啦，就这三行代码，加入到sshd_config 中即可，最终编辑完的效果就是这样：

```undefined
TCPKeepAlive yes
ClientAliveInterval 30
ClientAliveCountMax 3
```

保存并退出编辑器，最后重启SSH服务：

```undefined
service ssh restart  #或者：
/etc/init.d/ssh restart
```

上面命令不好用，用下面的

```bash
执行以下命令，重启sshd服务，使配置生效。

- CentOS6操作系统

  **# service sshd restart**

- CentOS7/EulerOS操作系统

  **# systemctl restart sshd**
```

这样子再也不用怕SSH会话会超时断开啦 !



参考： [华为云](https://support.huaweicloud.com/trouble-ecs/ecs_trouble_0306.html)  [博客](https://blog.sunriseydy.top/technology/linux/ssh-keep-alive/)

