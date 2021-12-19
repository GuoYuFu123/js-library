### jest-learn

*2021.05.30  15: 30 付国宇*

#### 一、匹配器

```
1、toBe 严格相等,值引用 Object.is()
2、toEqual 值相等
3、toBeNull 
4、toBeUndefined
5、toBeDefined
6、toBeTruthy
7、toBeFalsy

// 数字
8、toBeGreaterThan 大于
9、toBeGreaterThanOrEqual 大于等于
10、toBeLessThan 小于
11、toBeLessThanOrEqual 小于等于

//  字符串
12、toMatch 匹配 （可以是字符串，可以是正则）

//  数组
13、toContain 是否包含

// 异常
14、toThrow 异常匹配器
15、not 不是
```

更多匹配器请[官网](https://jestjs.io/docs/expect)查找



#### 二、异步代码测试

###### 1、回调形式

```javascript
test("test done", (done) => {
  fetchData((data) => {
    expect(data).toBe("fetch");
    done();
  });
});
```

2、返回promise

```javascript
test("测试Promise", () => {
  return fetchPromiseData().then((res) => {
    expect(res).toBe(2);
  });
});
```

断言

```javascript
/**
 * 请确保添加 expect.assertions 来验证一定数量的断言被调用。
 */
test("测试Error", () => {
  expect.assertions(1); // 断言，必须走一次expect
  return fetchErrorData().catch((e) => {
    expect(e.toString().includes("404")).toBe(true);
  });
});
```

3、async-await

```javascript
test("测试async-await2", async () => {
  const data = await fetchPromiseData();
  return expect(data).toEqual(2);
});
```

