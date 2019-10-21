<?php
    // smarty的引入和配置化
    require ('./Smarty/Smarty.class.php');
    $smarty = new Smarty();
    
    //五配置 两方法
    $smarty->left_delimiter = "{";  //左定界符
    $smarty->right_delimiter = "}"; //右定界符
    $smarty->template_dir = "tpl";  //html模板地址
    $smarty->compile_dir = "template_c"; //模板编译生成的文件
    $smarty->cache_dir = "cache"; //缓存

    //下面是开启缓存的另外2个配置
    $smarty->caching = true; //开始缓存
    $smarty->cache_lifetime = 120; // 缓存时间

    $smarty->assign("articletitle", "qwertt");  
    
    $arr = array("title"=>array("title"=>"二维数组"), "author"=>"小明");
    $smarty->assign("arr", $arr);

    $smarty->assign("time", time());  

    $smarty->assign("url", "https://www.runoob.com/php/func-date-time.html");

    $smarty->assign("article", "哈哈哈
    哈哈哈
    哈哈哈");

    $smarty->assign("ifkey", "guoguo");  

    $smarty->assign("score", 90);  

    $list = array(
        array("title"=>"1", "author"=>"guoguo1", "age"=>18),
        array("title"=>"1", "author"=>"guoguo2", "age"=>18)
    );
    $smarty->assign("list", $list);

    function test($param) {
        $p1 = $param['p1'];
        $p2 = $param['p2'];
        return "传入的参数p1=".$p1."传入的参数p2=".$p2;
    }
    $smarty->registerPlugin("function", "f_test", "test"); //第一个是注册的类型，注册的函数名，原函数名

    //页面显示
    $smarty->display("test.tpl");
?>