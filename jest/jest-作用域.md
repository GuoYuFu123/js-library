### jest-learn

*2021.05.30  19: 30 付国宇*

#### 一、hook钩子函数

1.  beforeAll：    在所有测试用例之前进行执行
2.  afterAll ：      在所有测试用例之后执行一次
3. beforeEach： 在 每次测试之前执行一次
4. afterEach ：   在每次测试之后执行一次
5. describe ：   组的概念， 进行测试分组类似作用域，，默认全局就是个全局作用域

注： 在describe嵌套的test中，hook函数的执行顺序

```bash
父[1，3]--->  子【1，3，4，2】---> 父 [4,2]
```

#### 二、hookscoped规律总结

```bash
/**
 * 作用域规则：
 * 钩子函数在父级分组可作用域子集，类似继承
 * 钩子函数同级分组作用域互不干扰，各起作用
 * 先执行外部的钩子函数，再执行内部的钩子函数
 */
```

#### 三、eg.示例

```javascript
describe("全局作用域", () => {
  console.log("我是全局作用域");
  beforeAll(() => {
    console.log("兄弟走起，大保健～～～");
  });

  beforeEach(() => {
    console.log("给钱，叫公主");
  });

  afterEach(() => {
    console.log("技术不错");
  });

  afterAll(() => {
    console.log("兄弟，撤了～～～");
  });

  describe("小明相关", () => {
    afterEach(() => {
      console.log("给了50元小费");
    });
    test("测试按脚", () => {
      console.log('小明按脚');
      expect('小明按脚').toBe("小明按脚");
    });
  });

  describe("小红相关", () => {
    afterEach(() => {
      console.log("给了80元小费");
    });
    test("测试按摩", () => {
      expect('小红按摩').toBe("小红按摩");
    });
  });
});

```





 

 



