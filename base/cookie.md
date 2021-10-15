1、cookie是客户端存储，因为http协议是无状态的，cookie的存在，让http协议记录稳定的状态信息有了可能

2、cookie主要用于一下三个方便

- 会话状态管理（如用户登录状态、购物车、游戏分数或其它需要记录的信息）
- 个性化设置（如用户自定义设置、主题等）
- 浏览器行为跟踪（如跟踪分析用户行为等）

3、会话级cookie和持久级cookie，可以通过设置过期时间来区分

4、cookie的同源策略

```
同端口，同协议，同域名
```

5、cookie跨域问题【cookie默认会携带同域的cookie，跨域的不会】

（1）如果不设置cookie，http请求不会自动携带cookie，需要前端设置

```java
withCredentials: true
```

（2）服务端也需要设置cors允许凭证，不然也接收不到，[nestjs]

```javascript
app.enableCors({
    credentials: true,
    methods: ['POST', 'GET', 'PUT', 'DELETE', 'OPTIONS'],
    allowedHeaders: [
      'Content-Type',
      'Access-Token',
      'X-XSRF-TOKEN',
      'User-Agent',
      'Referer',
      'Accept',
    ],
  });
```

以上2者必须同步设置

6、cookie遵循同源策略，所以在非同域名的复杂请求中，第5步设置完，在请求中携带的请求域名下的cookie

```javascript
web: www.guoguo.com
server: www.guo.com
🌟web->server: 携带的是server端的cookie
```

7、SameSite

`SameSite` Cookie 允许服务器要求某个 cookie 在跨站请求时不会被发送，

SameSite 可以有下面三种值： 默认为none

```
None。浏览器会在同站请求、跨站请求下继续发送 cookies，不区分大小写。
Strict。浏览器将只在访问相同站点时发送 cookie。（在原有 Cookies 的限制条件上的加强，如上文 “Cookie 的作用域” 所述）
Lax。与 Strict 类似，但用户从外部站点导航至URL时（例如通过链接）除外。 在新版本浏览器中，为默认选项，Same-site cookies 将会为一些跨站子请求保留，如图片加载或者 frames 的调用，但只有当用户从外部站点导航到URL时才会发送。如 link 链接
```

8、cookie大小：4或者5K