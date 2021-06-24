### jest-learn

*2021.02.28  15: 30 付国宇*

#### 一、认识前端自动化

随着前端的发展，前端设计的领域已经越来越多，也越来越复杂。这就对我们前端工程化能力，提出了更高的要求。 好的前端工程化一般包括三个大的方面：

- 前端自动化测试（前提条件）
- 高质量的代码设计
- 高质量的代码实现

#### 二、市面主流的前端测试框架

1. Jasmine （茉莉花）： JavaScript测试框架（BDD-集成测试开发框架），这个也算是比较早的测试框架；
2. MOCHA（摩卡）: 它是一个功能丰富的JavaScript测试框架，运行在`Node.js`和浏览器中，使异步测试变得简单有趣。也是非常优秀的框架；
3. Jest：目前最流行的前端测试框架，几乎国内所有的大型互联网公司都在使用；

#### 三、jest前端测试框架的优点

- 比较新：喜新厌旧是人的天性，新技术出来后我们总想动手尝试一下
- 基础很好：框架基础好就是性能好、功能多、简单易用，Jest在这三个方面你可以完全放心。绝对是一把好手，干就完事了。
- 速度快： 单独模块测试功能，比如说有两个模块A和B，以前都测试过了，这时候你只改动A模块，才次测试，模块B不会再跑一次，而是直接测试A模块。这就好比你去‘大宝剑’，你所有技师都试过了一次，下次你再来，直接就找最好的就行了。不用再测试一遍。（安心、放心）
- API简单 ：看一下官网，你就会发现API非常简单，数量也少。
- 隔离性好：Jest里会有很多的测试文件等待我们使用，Jest的执行环境都是隔离，这样就避免不同的测试文件执行的时候互相影响而造成出错。
- IDE整合：Jest直接可以和很多编辑器（VSCode）进行融合，让测试变的更加简单。
- 多项目并行：比如我们写了Node.js的后台项目，用React写了一个前台项目，Jest是支持他们并行运行，让我们的效率更加提高了。
- 快出覆盖率：（测试代码覆盖率） 对于一个项目的测试都要出覆盖率的，Jest就可以快速出这样的覆盖率统计结果，非常好用。

#### 四：测试区别

- **单元测试**：英文是(unit testing) 单,是指对软件中的最小可测试单元进行检查和验证。前端所说的单元测试就是对一个模块进行测试。也就是说前端测试的时候，你测试的东西一定是一个模块。
- **集成测试**：也叫组装测试或者联合测试。在单元测试的基础上，将所有模块按照涉及要求组装成为子系统或系统，进行集成测试。

#### 五、环境搭建+初始化配置

1. 初始化一个项目 npm init
2. 安装jest框架，因为这个只是开发使用 yarn add jest -D
3. npx jest --init 或者 全局安装jest初始化

#### 六、基础必会方法

- test方法：Jest封装的测试方法，一般填写两个参数，描述和测试方法，这个是测试的最小单元，不允许嵌套
- expect方法 ：预期方法，就是你调用了什么方法，传递了什么参数，得到的预期是什么(代码中详细讲解)。

```javascript
test('test',() => {
	expect("1").toBe("1");
})
```



#### 七、基本配置和覆盖率文件详解

1.jest.config.ts

```javascript
export default {
  clearMocks: true, // 是否每次运行清除mock
  collectCoverageFrom: ['src/*.{js,ts}', 'src/**/*.{js,ts}'], //覆盖率文件
	coverageDirectory: 'coverage', // 生成覆盖率的文件名
	coverageProvider: 'v8', 
	coverageThreshold: { // 覆盖率阈值
    global: {
      branches: 90,
      functions: 95,
      lines: 95,
      statements: 95,
    },
 },
 moduleFileExtensions: ['ts', 'tsx', 'js', 'json', 'jsx', 'node'], // 文件扩展
 preset: 'ts-jest', //ts
 setupFilesAfterEnv: [ // 每次test的启动文件，类似main.ts
    "<rootDir>/__tests__/boot.ts"
 ],
 testEnvironment: 'jest-environment-jsdom', // js
 testRegex: '(/__tests__/.+\\.(test|spec))\\.(ts|tsx|js)$', // 要进行test的文件正则
};

```

2、package.json 启动配置,

```
"scripts": {
    "test": "jest",
    "test:w": "jest --watchAll",
    "test:c": "jest --coverage", //  生成覆盖率文件
  },
```

3、运行yarn test :c

```bash
------------|---------|----------|---------|---------|-------------------
File        | % Stmts | % Branch | % Funcs | % Lines | Uncovered Line #s
------------|---------|----------|---------|---------|-------------------
All files   |     100 |    89.47 |     100 |     100 |
 01demo.js  |     100 |       50 |     100 |     100 | 2
 03es6.js   |     100 |       50 |     100 |     100 | 2
 04async.js |     100 |      100 |     100 |     100 |
 05hook.js  |     100 |      100 |     100 |     100 |
------------|---------|----------|---------|---------|-------------------

Test Suites: 6 passed, 6 total
Tests:       28 passed, 28 total
Snapshots:   0 total
Time:        9.378 s, estimated 11 s
Ran all test suites.
✨  Done in 10.55s.
```

- %stmts是语句覆盖率（statement coverage）：是不是每个语句都执行了？

- %Branch分支覆盖率（branch coverage）：是不是每个if代码块都执行了？

- %Funcs函数覆盖率（function coverage）：是不是每个函数都调用了？

- %Lines行覆盖率（line coverage）：是不是每一行都执行了？

  

  专业术语里，把describe包含的块叫做suite，把it/test包含的块叫做specification，也简称为spec，在一个suite里面可以包含多个数量的spec，但是也要注意结构化语义化。
