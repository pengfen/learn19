<?php
// 启动服务器 php udp_server.php
// netcat -u 127.0.0.1 9502 输入测试内容

// 创建 server对象 监听127.0.0.1:9502端口 类型SWOOLE_SOCK_UDP
$serv = new swoole_server("127.0.0.1", 9502, SWOOLE_PROCESS, SWOOLE_SOCK_UDP);

// 监听数据接收事件
// $clientInfo 客户端相关信息
$serv->on("Packet", function($serv, $data, $clientInfo) {
    $serv->sendto($clientInfo["address"], $clientInfo["port"], "server ".$data);
    //var_dump($clientInfo);
});

// 启动服务器
$serv->start();