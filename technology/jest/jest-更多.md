### jest-learn

*2021.05.30  22: 30 付国宇*

#### 一、[jest.fn](https://jestjs.io/docs/jest-object)

jest中的函数，具有跟踪性，可以测试是否被调用？调用几次？入参是否是期望值？

```javascript
it("should jest fn", () => {
    const jestFn = jest.fn((val) => {
      return val;
    });
    const pramas = 1;
    const res = jestFn(pramas);
    expect(res).toBe(1);
    expect(jestFn).toHaveBeenCalled();
    expect(jestFn).toHaveBeenCalledTimes(1);
    expect(jestFn).toHaveBeenCalledWith(pramas);
  });
```

#### 二、[jest.mock](https://jestjs.io/docs/jest-object)

jest中的mock函数，当不会真的去发起函数调用，会使用定义的函数进行mock返回

```javascript
it("should jest.mock", () => {
    jest.mock("../src/06mock", () => {
      return 5;
    });
    const mockFn = require("../src/06mock");
    expect(mockFn).toBe(5);
  });
```

#### 三、[jest.spyOn](https://jestjs.io/docs/jest-object)

`jest.spyOn()`方法同样创建一个mock函数，但是该mock函数不仅能够捕获函数的调用情况，还可以正常的执行被spy的函数。实际上，`jest.spyOn()`是`jest.fn()`的语法糖，它创建了一个和被spy的函数具有相同内部代码的mock函数。

```javascript
  it("should jest.spyOn", () => {
    const obj = require("../src/07spyOn");
    const spy = jest.spyOn(obj.default, "get");
    expect(spy()).toBe("牛");
    expect(spy).toHaveBeenCalled();
  });

  it("should jest.spyOn", () => {
    const obj = require("../src/07spyOn");
    const spy = jest.spyOn(obj.default, "get").mockReturnValue("好牛");
    expect(spy()).toBe("好牛");
    expect(spy).toHaveBeenCalled();
  });
```

#### 四、[更多](https://www.jestjs.cn/docs/more-resources)

1. 模拟延时

2. 模拟class

3. 模拟webpack，puppeteer等

   ...

```bash
未完待续...
```

