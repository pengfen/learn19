<?php

$ws_server = new swoole_websocket_server("127.0.0.1", 9502);

$redis = new Redis();
$redis->connect("127.0.0.1", 6379);

$ws_server->on('open', function($ws, $request) use ($redis) {
    $redis->sAdd("fd", $request->fd);
});

$ws_server->on('message', function($ws, $frame) use ($redis) {
    global $redis;
    $fds = $redis->sMembers("fd");
    foreach ($fds as $fd) {
        $ws->push($fd,$frame->fd.'--'.$frame->data);
    }
});

$ws_server->on("close", function ($ws, $fd) use ($redis) {
    $redis->sRem("fd", $fd);
});

$ws_server->start();