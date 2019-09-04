# redis.conf参数：

* daemonize：是否以后台daemon方式运行
* pidfile：pid文件位置
* port：监听的端口号
* timeout：请求超时时间
* loglevel：log信息级别
* logfile：log文件位置
* databases：开启数据库的数量
* save ：保存快照的频率，第一个表示多长时间（秒级），第三个表示执行多少次写操作。在一定时间内执行一定数量的写操作时，自动保存快照。可设置多个条件。
* rdbcompression：是否使用压缩
* dbfilename：数据快照文件名（只是文件名，不包括目录）
* dir：数据快照的保存目录（这个是目录）
* appendonly：是否开启appendonlylog，开启的话每次写操作会记一条log，这会提高数据抗风险能力，但影响效率。
* appendfsync：appendonlylog如何同步到磁盘（三个选项，分别是每次写都强制调用fsync、每秒启用一次fsync、不调用fsync等待系统自己同步）
* slaveof ：主从配置，在redis-slave上配置master的ip port，即可。


例如，我们可以修改为如下方式：
引用
daemonize yes #守护进程模式
save 60 1000 #当时间间隔超过60秒，或存储超过1000条记录时，进行持久化。
maxmemory 256mb #分配256MB内存