## 什么是Service Work

author: 果果    data：2020.04.25 15:50

##### 背景

```
随着web的快速发展，用户对于站点的体验期望也是越来越来高，作为一个前端工程师，就必须为性能的提高而耗费心
思；CDN、CSS Sprite、文件的合并压缩、异步加载、资源缓存等方式的目的就是减少用户的感知，提高用户的体验；
```

##### service work

```
众所周知，浏览器中的js都是运行在单一主线程上的，在同一时间内只能做一件事情。这个时候web work出现了，它是
独立于主线程之外的，但是它是临时的，最后service work来了，我们来看一下她有哪些功能和特性；
一个独立的 worker 线程，独立于当前网页进程，有自己独立的 worker context。
1、一旦被 install，就永远存在，除非被手动 unregister
2、用到的时候可以直接唤醒，不用的时候自动睡眠
3、可编程拦截代理请求和返回，缓存文件，缓存的文件可以被网页进程取到（包括网络离线状态）
4、离线内容开发者可控
5、能向客户端推送消息
6、不能直接操作 DOM
7、必须在 HTTPS 环境下才能工作
8、异步实现，内部大都是通过 Promise 实现
```



#### 一块来大展身手

```
1、必须是https环境，本地调试localhost或者127.0.0.1环境也是可以的，
2、依赖于cache api进行实现的
3、依赖于h5的fetch Api
4、依赖于promise进行实现
```

##### 注册service

```
if ('serviceWorker' in navigator) {
    window.addEventListener('load', function () {
        const sw = navigator.serviceWorker;
        sw.register('/pwa/sw.js', {scope: '/pwa/'})
    	.then(function (registration) {
            // 注册成功
            console.log('ServiceWorker registration successful with scope: ', registration);
                })
        .catch(function (err) {
            // 注册失败
            console.log('ServiceWorker registration failed: ', err);
        }); 
    });
 }
```

1、先判断浏览器是否支持sw，否则就免谈了，

2、在页面load时候，注册./sw.js的service work，页面每次加载都会注册，浏览器会根据注册情况作出相应的处理

3、scope参数就是一个作用域，默认作用于当前站点

##### 注册成功

![注册成功](./\注册成功.jpg)

##### 单纯的注册成功是没有什么用的，如何来缓存呢？这就需要我们依赖于cache来进行缓存了

```
console.log('******************sw.js**************************')
const version = 1;
const cacheStorageKey = "testCache-" + version;

// 这是需要预缓存的资源，也可以是appshell,可以通过webpack的插件来生成
const cacheList = ["/", "index.html", 'data.json'];

// 注册成功的时候，以版本名为key主动缓存静态资源
console.log(self)
self.addEventListener("install", function(e) {
  console.log("Cache event!");
  e.waitUntil(
    caches.open(cacheStorageKey).then(function(cache) {
      console.log("Adding to Cache:", cacheList);
      return cache.addAll(cacheList);
    })
    .then(function() {
      // 注册成功跳过等待，酌情处理
      // console.log('Skip waiting!')
      // return self.skipWaiting()
    })
  );
});

// 当新的serviceWorker被激活时，删除旧版本的缓存
self.addEventListener("activate", event => {
  console.log("Activate event");
  event.waitUntil(
    caches
      .keys()
      .then(cacheNames => {
        return cacheNames.filter(cacheName => cacheStorageKey !== cacheName);
      })
      .then(cachesToDelete => {
        return Promise.all(
          cachesToDelete.map(cacheToDelete => {
            return caches.delete(cacheToDelete);
          })
        );
      })
      .then(() => {
        console.log("Clients claims.");
        // 立即接管所有页面，酌情处理
        // 会导致新的sw接管旧的页面，同时旧版本的缓存已被清空
        self.clients.claim();
      })
  );
});

// 发起请求时去根据uri去匹配缓存，无法命中缓存则发起请求
self.addEventListener("fetch", function(e) {
  e.respondWith(
    caches.match(e.request).then(function(response) {
      return response || fetch(e.request);
    })
  );
});

self.addEventListener("message", event => {
  if (event.data === "skipWaiting") {
    console.log("Skip waiting!");
    self.skipWaiting();
  }
});
```

##### 我们缓存的资源都在哪里了呢？都在cache的缓存里

![资源缓存](./\资源缓存.jpg)

#### 测试一下

1、在不开启service work的情况下访问资源

![不适用service work](D:\Program Files\phpStudy\WWW\pwa\markdown\不适用work.jpg)

2、使用service work访问资源

![](D:\Program Files\phpStudy\WWW\pwa\markdown\使用work.jpg)

####  结语

 ```javascrip
技术交流，共同进步，欢迎fork和star！
 ```
[参考链接](https://www.bookstack.cn/read/pwa-doc/README.md )

