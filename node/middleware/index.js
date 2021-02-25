/**
 * 实现一个middleware
 * 关联问题： 中间件的洋葱模型
 * 顺序： requset =》 洋葱内部（顺序）=》洋葱外部（倒叙） =》 resquest
 *
 * 中间件的理解：
 * 1、一个请求发送到服务器之后，它的生命周期是先收到request（请求），然后服务端处理，处理完了以后发送response（响应）回去
 * 2、app.use 加载用于处理http请求的middleware（中间件），当一个请求来的时候，会依次被这些 middlewares处理。
 * 3、中间件其是一个函数，在响应发送之前对请求进行一些操作
 * 4、框架内部维护一个函数数组，这个函数数组表示在发出响应之前要执行的所有函数，也就是中间件数组，使用app.use(fn)后，传进来的fn就会被扔到这个数组里，执行完毕后调用next()方法执行函数数组里的下一个函数，如果没有调用next()的话，就不会调用下一个函数了，也就是说调用就会被终止
 */

function _Koa() {
  let funcs = [];

  var _app = function (req, res) {
    var i = 0;
    function next() {
      var task = funcs[i++];
      if (!task) {
        return;
      }
      task(req, res, next);
    }
    next();
  };

  _app.use = function (task) {
    funcs.push(task);
  };

  return _app;
}

function middleware1(req, res, next) {
  console.log("middleware1 before next");
  next();
  console.log("middleware1 after next");
}
function middleware2(req, res, next) {
  console.log("middleware2 before next");
  next();
  console.log("middleware2 before next");
}

let app = _Koa();

app.use(middleware1);
app.use(middleware2);

app("MockQeq", "MockRes");

/**
 * 输出：
middleware1 before next
middleware2 before next
middleware2 before next
middleware1 after next
 */
