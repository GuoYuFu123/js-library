### PHP模板Smarty学习 

*created by fuguoyu 2019.10.21*

#### 一、项目搭建区域

##### 1、引入和实例化Smarty

```
    // smarty的引入和配置化
    require ('./Smarty/Smarty.class.php');
    $smarty = new Smarty();
```

##### 2、五配置两方法

```
    //五配置 两方法
    $smarty->left_delimiter = "{";  //左定界符
    $smarty->right_delimiter = "}"; //右定界符
    $smarty->template_dir = "tpl";  //html模板地址
    $smarty->compile_dir = "template_c"; //模板编译生成的文件
    $smarty->cache_dir = "cache"; //缓存

    //下面是开启缓存的另外2个配置
    $smarty->caching = true; //开始缓存
    $smarty->cache_lifetime = 120; // 缓存时间
```

##### 3、分配变量与页面显示

```
    $smarty->assign("articletitle", "qwertt");  
    //页面显示
    $smarty->display("test.tpl");
```



#### 二：语法区域

##### 1、注释与变量输出

```
1.注释
    {*  XXX *}
    
2.在smarty中输出赋值进来的变量
```

##### 2、变量调节器

```
1、首字母大写 capitalize
eg：{$articleTitle| capitalize}

2、字符串连接cat，通过:可以进行无限连接
eg：{$articleTitle|cat:" yesterday"}

3、日期格式化 date_format
eg:  {$yesterday|date_format}
	 {$yesterday|date_format:"%A, %B, %e, %Y, %H:%M:%S"}

4、为未赋值或者为空的变量制定默认值default——"%Y/%m/%d %H:%M:%S %B %e" 
eg：{$articleTitle|default:"no title"} //变量为空的时候，输出 no title

5、转码 escape 
eg: 用于html的转码，url转码，在没有转码的变量上转换单引号，十六进制转码，十六进制美化，或者javascript转码，默认是html转码
{url|escape:"url"}

6、小写lower大写upper
eg：将变量字符串小（大）写
{$articleTitle|lower}  {$articleTitle|upper}
 
7、所有的换行符被替换成<br /> nl2br功能和PHP中的nl2br()函数一样
eg：{$articleTitle|nl2br}

```

##### 3、条件判断

```
1、基本句式
{if $ifkey eq "guoguo"}
	guoguo
{elseif $ifkey eq "haha"}
	haha
{else}
	hehe
{/if}

2、条件修饰符很多
eg：eq（==）  neq（!=）  gt(>)  gte(>=)  lt(<)  lte(<=)

3、修饰词时必须和变量或者常量用空格隔开
```

##### 4、循环

```
1、section， sectionelse
	（1）功能多，参数多，但是个人感觉并不实用，是smarty用来做循环操作的函数之一
	（2）基本属性
		1)name 循环数组的每一项
		2)loop 循环的数组
		3)start 循环执行的初始位置，如果为负数，则循环从数组的尾部算起
		4)step 决定循环的步长，如果为负数，遍历数组从否像前进行遍历
		5)max 设置循环的最大执行次数
		6)show 决定是否显示该循环
	（3）code
	{section name=list loop=$list max=1}
        {$list[list].title}
        <br />
    {/section}
		
2、foreach【推荐使用】
只能处理简单的数组
code：
{foreach $list as $item}
    {$item.author}
	{foreachelse}
		当前没有文章
{/foreach}
```

##### 5、模板的引用（引入文件或者子组件）

```
1、include用法和php钟的include差不多
2、smarty的include还具备自定义的属性的功能【属性传值】
eg：{include file="header.tpl" title="网站头部" table_bgcolor="#e6e6e6" sitename="果果网"}
```

##### 6、类与对象赋值与使用

```
1、smarty中的register_object方法，在smarty3中已经废除
2、使用assign把一个类中的对象以变量的形式赋值到smarty的模板中使用
```

##### 7、内置函数与自定义函数

```
1、可以使用php的内置函数
    {* date("Y-m-d", $time) *}
    {'Y-m-d'|date: $time}
    <br />
    {* str_replace("q", "Q", $articletitle) *}
    {"q"|str_replace:"Q": $articletitle}

2、可以是自定义函数，并用registerPlugin注册到smarty模板中使用
第一个参数可以是function，modifier，block等，下面以函数为例
	1）定义函数
        function test($param) {
            $p1 = $param['p1'];
            $p2 = $param['p2'];
            return "传入的参数p1=".$p1."传入的参数p2=".$p2;
        }
        $smarty->registerPlugin("function", "f_test", "test"); //第一个是注册的类型，注册的函数名，原函数名
    2）调用函数 {函数名 参数1=XXX 参数2=XXX}
    	{* 调用f_test函数，属性会整合为参数数组即array("p1"=>"guoguo","p2"=>"最帅"...) *}
		{f_test p1="guoguo" p2="最帅"} 
    
```

##### 8、插件学习

```
1、什么是smarty的插件？
	Smarty的插件本质上是funcion函数
2、Smarty插件常用的类型
	（1）functions 函数插件
	（2）modifiers 修饰插件
	（3）block functions 区块函数插件
3、如何来制作、使用插件
	（1）使用registerPlugin方法注册写好的自定义函数
	（2）将写好的插件放入Smarty解压目录中的lib目录下的plugin目录里【注意插件命名格式】
	（3）php的内置函数，可以自动以修饰插件（变量调节器插件）的形式在模板里使用
```

