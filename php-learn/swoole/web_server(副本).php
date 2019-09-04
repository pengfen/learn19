<?php

$http = new swoole_http_server("127.0.0.1", 9501);

//$http->set(
//    [
//        'enable_static_handler' => true,
//        'document_root' => "/var/www/wordpress/swoole",
//    ]
//);

$http->on("request", function ($request, $response) {
    var_dump($request->get, $request->post);
    $response->header("Content-Type", "text/html; charset=utf-8");
    $response->end("<h1>Hello Swoole. #".rand(1000, 9999)."</h1>");
});

$http->start();

// 注意服务器端口不对外开放
// 一使用iptables规则对外开放9501端口

// 二使用代理
// php web_server.php
// cd /etc/nginx/cond.f
// sudo vim dev.res.com
// location /websocket { proxy_pass http:127.0.0.1:9501;}
// sudo nginx -r reload
// http://dev.res.com/websocket (相当于dev.res.com:9501)