# ubuntu 上安装redis
* sudo apt-get install -y redis-server
* http://redis.io/download 下载 redis相关包

* linux上创建redis
* mkdir redis
* cd redis
* make PREFIX=/usr/local/redis install
* cd /usr/local/redis
* cd bin
* ./redis-server

## redis 相关命令
redis-benchmark --- redis性能测试工具
redis-check-aof --- 检查aof日志工具 如果日志损坏能检查出来
redis-check-dump --- 检查rdb日志工具
redis-cli --- 连接用的客户端
redis-server --- redis服务区进程

## 添加对外访问端口
* iptables -I INPUT -p tcp --dport 6379 -j ACCEPT
* -I 添加一条规则
* INPUT 进入规则
* -p protocol协议 -p tcp tcp协议
* --dport destination port 目标端口
* -j 制定规则 

* vi redis.conf (配置参与redis.conf)
* bind 127.0.0.1
* bind 192.168.233.130
* 修改保护模式就不用配置bind
* protected-mode yes --- protected-mode no

* daemonize yes #使用redis-server启动时会在后台运行
* src/redis-server redis.conf #启动redis(设置daemonize后会在后台运行机制)
* ps -ef | grep redis #查看redis相关进程


# window 上安装redis

1. 下载
https://github.com/MicrosoftArchive/redis/releases/tag ---> 选择版本 ---> 下载 .msi包

2. 安装

3. 进入安装目录
redis-server.exe redis.windows-service.conf