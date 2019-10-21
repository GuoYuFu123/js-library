<?php
    function smarty_block_test2($params, $content){
        // echo $params;
        $replace = $params["replace"];
        $maxnum = $params["maxnum"];
        if ($replace == "true") {
            $content = str_replace('，',",", $content);
            $content = str_replace('。',".", $content);
        }
        $content = substr($content, 0 ,$maxnum);
        return $content;
    }
?>