<?php

$redis = new redis();
$redis->connect('127.0.0.1', 6379);
$redis->select(2); // select 2

// keys * 
$redis->lpush('list', 'header1'); // lrange list 0 -1
$redis->lpush('list', 'header2'); // lrange list 0 -1

echo '列表长度是 : '.$redis->lsize('test');

$redis->rpush('list', 'tail1');
$redis->rpush('list', 'tail2');

echo '列表长度是 : '.$redis->llen('list'); // llen list