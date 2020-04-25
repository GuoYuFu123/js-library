## manifest

author: 果果    data：2020.04.25 17:06

##### html属性 manifest定义

```
manifest 属性规定文档的缓存 manifest 的位置。

HTML5 引入了应用程序缓存，这意味着 Web 应用程序可以被缓存，然后在无互联网连接的时候进行访问。

应用程序缓存使得应用程序有三个优点：

离线浏览 - 用户可以在离线时使用应用程序
快速 - 缓存的资源可以更快地加载
减少服务器加载 - 浏览器只从服务器上下载已更新/已更改的资源
manifest 属性应该被 Web 应用程序中您想要缓存的每个页面包含。

manifest 文件是一个简单的文本文件，列举出了浏览器用于离线访问而缓存的资源。
```

##### 如何使用呢

![manifest](./\manifest.jpg)

##### 测试一下

1、浏览器中在哪看manifest缓存

![浏览器看缓存](.\浏览器manifest.jpg)

2、是否走缓存了呢？

![manifest缓存](D:\Program Files\phpStudy\WWW\pwa\markdown\manifest-test.jpg)

####  结语

 ```javascrip
技术交流，共同进步，欢迎fork和star！
 ```
