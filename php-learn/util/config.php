<?php
/**
 * 数据库配置文件
 */

// 开发环境

return $config = [
    "host" => "localhost",
    "user" => "root",
    "password" => "root",
    "dbname" => "",
    "charset" => "UTF8",
    "lang" => "en",
    "author_flag" => true, // false 只查询作者用户 true 查询所有用户
    "domain" => "localhost",
];


// 测试环境 dev
/*
return $config = [
    "host" => "localhost",
    "user" => "root",
    "password" => "root",
    "dbname" => "",
    "charset" => "UTF8",
    "lang" => "en",
    "author_flag" => true, // false 只查询作者用户 true 查询所有用户
    "domain" => "localhost",
];
*/


// 正式环境 online
/*
return $config = [
    "host" => "localhost",
    "user" => "root",
    "password" => "root",
    "dbname" => "wordpress",
    "charset" => "UTF8",
    "lang" => "en",
    "author_flag" => true, // false 只查询作者用户 true 查询所有用户
    "domain" => "localhost",
];
*/