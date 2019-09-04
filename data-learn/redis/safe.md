# 配置redis密码

* sudo find / -name redis.conf
* /etc/redis/redis.conf
* echo requirepass secret >> /usr/local/redis/etc/redis.conf
* systemctl restart redis-server

* 修改密码
* sudo apt-get install -y redis-server
* sudo vim /etc/redis/redis.conf
* 修改 requirepass password
* sudo /etc/init.d/redis-server restart