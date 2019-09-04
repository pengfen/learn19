## mysql 安装
* sudo apt-get install mysql-server
* sudo apt-get isntall mysql-client
* sudo apt-get install libmysqlclient-dev

## 解决mysql5.7登录不了
* sudo cat /etc/mysql/debian.cnf
* mysql -u账户 -p密码
* use mysql
* update user set authentication_string=password("123456") where user="root"
* flush privileges;
* select user, host, plugin from user where user = "root"
* update user set plugin = "mysql_native_password" where user = "root"
* flush privileges;

## redis 安装
* sudo apt-get install -y redis-server
* redis-cli
* sudo apt-get install -y php-redis