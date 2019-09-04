# mysql 安装
* sudo apt-get install mysql-server mysql-client libmysqlclient-dev

## 解决mysql5.7登录不了
* sudo cat /etc/mysql/debian.cnf
* mysql -u账户 -p密码
* use mysql
* update user set authentication_string=password("123456") where user="root"
* flush privileges;
* select user, host, plugin from user where user = "root"
* update user set plugin = "mysql_native_password" where user = "root"
* flush privileges;

## 相关目录
* /usr/bin/mysql mysql命令
* /etc/mysql mysql相关配置修改目录
* /var/lib/mysql 数据库(sudo su - cd /var/lib/mysql)
* /var/log/mysql 日志目录