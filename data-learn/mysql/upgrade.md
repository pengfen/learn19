# windows mysql升级

* 启动mysql --- cmd --- net start mysql

* 升级mysql至5.7(解决utfmb4)
1. 备份phpstudy下MySQL目录
2. 解决下载的mysql安装包 https://dev.mysql.com/downloads/mysql/5.7.html#downloads
3. 修改my.ini (my-default.ini)
skip-grant-tables #安全模式
basedir="D:/phpstudy/PHPTutorial/MySQL/"
datadir="D:/phpstudy/PHPTutorial/MySQL/data/"
4. 配置环境变量
5. mysqld --intialize #初始化
6. mysqld -install
7. net start mysql
8. mysql -uroot -p
9. use mysql
10. update mysql.user set authentication_string=password('新密码') where user='root';
11：flush privileges;
12. 注释安全模式 #skip-grant-tables

* net start mysql
* net stop mysql

# linux maridb 参考 dba/install.md