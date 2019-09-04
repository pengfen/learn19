jmeter 性能压测

线程组 Http请求 查看结果树 聚合报告

线程组
	Http请求 ---> 取样器
	查看结果树 ---> 监听器
	聚合报告 ---> 监听器

ps -ef | grep php 查看php进程
netstat -anp | grep 80 查看端口号

聚合报告
平均值 X毫秒
90% X毫秒
95% X毫秒
99% X毫秒
最小值 
最大值 X毫秒
异常
吞吐量

发现容量问题
server 端并发线程数上不去

pstree -p 5240 | wc -l 查看线程数并统计
top -H

发现容量问题
响应时间变长 TPS上不去

单Web容器上限
线程数量 4核cpu 8G内存单进程调度线程数 800-1000 以上后即花费巨大的时间在cpu调度上
等待队列长度 队列做缓冲池用 但也不能无限长 消耗内存 出队入队也耗cpu

MySql 数据库 QPS 容量问题
主键查询 千万级别数据 = 1-10毫秒
唯一索引查询 千万级别数据 = 10-100毫秒
非唯一索引查询 千万级别数据 = 100-1000毫秒
无索引 百万条数据 = 1000 毫秒+

非插入更新删除操作 同查询
插入操作 1W ～ 10W tps (依赖配置优化)

http://self.libang.com/v1/company/industry
并发 500
使用redis缓存 并发1100