<?php
/**
 * Created by PhpStorm.
 * User: artnew
 * Date: 2018-08-23
 * Time: 9:42
 */

$source = isset($_POST["source"])?$_POST["source"]:"";
$type = isset($_POST["type"])?$_POST["type"]:1;

if ($source) {
    if ($type == 1) {
        require_once './yudao.php';
        //调用翻译
        $info = translate($source,"zh-CHS","EN");
        $translate = $info["translation"];
        echo $translate[0];
    } else {
        require_once './baidu.php';
        $info = translate($source,"zh","en");
        echo $info["trans_result"][0]["dst"];
    }

} else {
    echo "无翻译内容";
}