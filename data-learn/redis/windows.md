# windows
* https://github.com/MSOpenTech/redis/releases

* 下载 zip包

* 解压 --- 复制到 e:/files/redis

* cmd --- e: --- cd files/redis

* E:\files\redis>redis-server.exe redis.windows.conf

* //连接本地的 Redis 服务
* $redis = new Redis();
* $redis->connect('127.0.0.1', 6379);
* var_dump($redis); // object(Redis)#1 (1) { ["socket"]=> resource(3) of type (Redis Socket Buffer) }
* // 如果redis连接不上 会出现object(Redis)#1 (0) { }

* 打开客户端
* redis-cli.exe

127.0.0.1:6379> keys *
(empty list or set)
127.0.0.1:6379> set name apeng
OK
127.0.0.1:6379> get name
"apeng"
127.0.0.1:6379> select 1 #选择一号数据库
OK
127.0.0.1:6379[1]> keys *
1) "x"
127.0.0.1:6379[1]> get x
"6"