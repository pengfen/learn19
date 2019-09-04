<?php

//异步客户端是非阻塞的。可以用于编写高并发的程序

// 创建异步TCP客户端
$client = new swoole_client(SWOOLE_SOCK_TCP, SWOOLE_SOCK_ASYNC);

// 注册连接成功回调
$client->on("connect", function ($cli) {
    $cli->send("register success");
});

// 注册数据接收回调
$client->on("receive", function ($cli, $data) {
    echo "receive: ".$data;
});

// 注册连接失败回调
$client->on("error", function ($cli) {
    echo "connction falied";
});

// 注册连接关闭回调
$client->on("close", function () {
    echo "connection close";
});

// 发起连接
$client->connect('127.0.0.1', 9501, 0.5);