<?php
/**
 * Created by PhpStorm.
 * User: artnew
 * Date: 2018-07-20
 * Time: 10:00
 */

// 创建server对象 监听127.0.0.1:9501
$serv = new swoole_server("127.0.0.1", 9501);

// 监听连接进入事件
$serv->on("connect", function ($serv, $fd) {
    echo "client: connect";
});

// 监听数据接收事件
$serv->on("receive", function ($serv, $fd, $from_id, $data) {
    $serv->send($fd, "server: ".$data);
});

// 监听连接关闭事件
$serv->on("close", function ($serv, $fd) {
    echo "client close ";
});

// 启动服务器
$serv->start();