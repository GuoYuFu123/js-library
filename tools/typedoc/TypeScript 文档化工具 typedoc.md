

### TypeScript 文档化工具: typedoc

#### 一、简介

目前咱们的项目基本是ts来开发，那么通过typedoc这个神器，就能自动化的帮助我们去生成文档，其实它就是一个ts的文档化工具

#### 二、与jsDoc的区别

TypeDoc 使用的是 `JavaDoc` 而不是 `JSDoc`，由于 TypeScript 更像 Java 而不像 JavaScript 所以使用了更为简单的 JavaDoc 这样一样，Jsdoc里面的很多标签也就没有什么意义了。

#### 三、diff

```javascript
// 1. 在JSDoc里面的写法
/**
* @param {string} name 姓名
* @param {number} age 年龄
*/


// 2. 在TypeDoc里面的写法
/**
* @param name 姓名
* @param age 年龄
*/
```

#### 四、TypeDoc的标签列表

| 名称        | 作用       | 备注                                                         |
| ----------- | ---------- | ------------------------------------------------------------ |
| @param      | 参数描述   | 仅供类、接口、方法注释时使用。同一个注释块可同时出现多个`param`描述。 |
| @return     | 返回描述   | 仅供方法注释时使用。除`void`方法外其它所有方法必须有一个`return`描述。 |
| @throws     | 异常描述   | 零到多个。                                                   |
| @exception  | 异常描述   | 零到多个。                                                   |
| @author     | 作者       | 类和接口注释中必须有。可有零到多个。                         |
| @version    | 版本描述   | 类和接口注释中必须有。零或一个。                             |
| @see        | 参考描述   | 可有零到多个。                                               |
| @since      | 起始版本   | 只有一个。                                                   |
| @serial     | 序列化描述 | 或`@serialField`或`@serialData`，可有多个                    |
| @deprecated | 废除标志   | 最多一个。                                                   |

#### 五、开始使用

##### 1、安装

```bas
yarn add typedoc -D
```

##### 2、使用生成

1）直接通过命令行生成，配置项可单独写文件也可作为参数【package, json】

```json
{
  "scripts": {
    "docs": "typedoc",
    "docs:shell":"typedoc --out docs --target es6 --theme minimal"
  },
}
```

【typedoc.json】[官网直达](https://typedoc.org/guides/installation/)

```json
{
  "name": "guoguo-ts-axios",
  "entryPoints": ["src"],
  "out": "docs",
  "theme": "default",
  "readme": "../../README.md"
}
```

2)  通过插件生成

Webpack ：https://www.npmjs.com/package/typedoc-webpack-plugin

Gulp ：https://www.npmjs.org/package/gulp-typedoc/

Grunt ：https://www.npmjs.org/package/grunt-typedoc

#### 六、配置项参数 【[官网直达](https://typedoc.org/guides/options/)】

| 参数                | 类型    | 说明                                                         |
| ------------------- | ------- | ------------------------------------------------------------ |
| out                 | string  | 输出目录                                                     |
| module              | string  | 模块引入方式，可以是 commonjs, amd, system, umd              |
| target              | string  | ES3(默认), ES5, ES6                                          |
| name                | string  | 项目标题                                                     |
| theme               | string  | 皮肤可以是 `default` or `minimal` or 一个路径，[更多资料](http://typedoc.org/guides/themes/) |
| readme              | string  | readme文件，markdown文件的相对地址                           |
| includeDeclarations | boolean | 是否包含 `.d.ts` 文件，如果你的项目是javascript写的，可以使用声明文件的方式来支持TypeScript并生成文档 |
| excludeExternals    | boolean | 是否排除外部引入的模块                                       |
| excludePrivate      | boolean | 是否排除 `private` 修饰的相关字段方法                        |
| excludeProtected    | boolean | 是否排除 `protected` 修饰的相关字段方法                      |
| hideGenerator       | boolean | 隐藏页底的全局链接                                           |
| version             | boolean | 显示 typedoc 版本                                            |
| help                | boolean | 显示帮助信息                                                 |
| gaID                | string  | 如果有 `Google Analytics` 的跟踪ID，可以设置                 |
| exclude             | string  | 排除文件                                                     |
| includes            | string  | 包含文件，应该是一个文件夹的名字，会将下面所有的md文件包含进来（需要使用 `[[include:document.md]]` 引入） |
| media               | string  | 包含媒体，应该是一个文件夹的名字，会包含文件夹下的图片等各种媒体文件（需要使用 `![logo](media://logo.png)` 引入） |

#### 七、更多

还支持生成md文件等，更多请移步[官网](https://typedoc.org/)学习