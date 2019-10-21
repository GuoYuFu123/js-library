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

    //页面显示
    $smarty->display("area.tpl");
?>