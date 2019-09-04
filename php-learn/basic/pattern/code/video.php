<?php

// 文章中添加视频 [embed]url[/embed]

$pattern = "/(?:\[)embed(?:\])(.*?)(?:\[)\/embed(?:\])/ism"; // 注意 ?:\[ 匹配中括号
preg_match_all($pattern, $content, $m); // 匹配所有的视频
if (isset($m[1]) && $m[1]) {
    foreach($m[1] as $key => $val) {
        $rep = "<video width='320' height='240' src='".$val."' loop='loop' x-webkit-airplay='allow' controls='controls'></video>";
        // 将匹配到的视频替换为video
        $content = str_replace($m[0][$key], $rep, $content);
    }
}