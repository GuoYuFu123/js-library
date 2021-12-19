### commit镜像

```shell
docker commit 提交容器成为一个新的副本

# 命令和git的原理类似
docker commit -m="提交的描述信息" -a="作者" 容器ID 要创建的目标镜像名:[标签名]

-a：提交的镜像作者；
-c：使用Dockerfile指令来创建镜像；
-m：提交时的说明文字；
-p：在commit时，将容器暂停。
```

实战测试

```shell
# 1、启动一个默认的tomcat

# 2、发现这个默认的tomcat是没有webapps应用的，镜像的原因，官网的镜像默认 webapps 目录下是没有文件的

# 3、cp进去基本的文件

# 4、将我们操作过的容器通过commit提交为一个新的镜像，然后我们使用我们自己的镜像就可以了
```

```shell
[root@guoguo ~]# docker pull tomcat

[root@guoguo ~]# docker run -d -p 8088:8080 --name tomcat01 tomcat

[root@guoguo ~]# docker exec -it tomcat01 /bin/bash

root@f7bea1ba5ae9:/usr/local/tomcat# cp -r  webapps.dist/* webapps

[root@guoguo ~]# ctrl + P and ctrl + q

[root@guoguo ~]# docker commit -m="add webapps App" -a="guoguo" f7bea1ba5ae9 guoguo/tomcat:1.0

[root@guoguo ~]# docker images
REPOSITORY            TAG       IMAGE ID       CREATED          SIZE
guoguo/tomcat         1.0       c37ffb5b6c1c   36 minutes ago   684MB

# 启动自己做的容器
[root@guoguo ~] docker run -it -p 8888:8080 guoguo/tomcat:1.0

```

> 提示：
>
> ```
> docker commit -m="add webapps App" -a="guoguo" f7bea1ba5ae9 guoguo/tomcat:1.0
> ```
>
> 命令的最后tomcat:1.0为自定义的镜像名和版本，也可以在前面加入Namespace的，也就如guoguo/tomcat:1.0，这个abc就叫Namespace，就相当与Java中，类前面的包名。我们拉取镜像的时候也经常会看到一些这样命名的包，如用到过的`portainer/portainer`工具。
>
> 如果是有命名，通过name启动的时候需加上tag，不然找不到指定的image
>
> 或者直接通过imageId启动 

至此，docker就已经入门啦











