<?php

// 同步TCP客户端
$client = new swoole_client(SWOOLE_SOCK_TCP);

// connect/send/recv会等待IO完成后再返回 同步阻塞操作并不消耗CPU资源 IO操作未完成当前进程会自动转入sleep模式 当IO完成后操作系统会唤醒当前进程 继续向下执行代码
// 连接服务器
if (!$client->connect("127.0.0.1", 9501, 0.5)) {
    die("connect failed. ");
}

// $client->send 发送少量数据时 都是立即返回的 发送大量数据时 socket缓存区可能会塞满 send操作会阻塞
// 向服务器发送数据
if (!$client->send("welcome to swoole world!")) {
    die("send failed. ");
}

// recv 操作会阻塞等待服务器返回数据 recv耗时等于服务器处理时间+网络传输耗时之合
// 从服务器接收数据
$data = $client->recv();
if (!$data) {
    die("recv failed. ");
}

echo $data;
// 关闭连接
$client->close();