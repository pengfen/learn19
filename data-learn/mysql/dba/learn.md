# mysql DBA 学习路线

## linux基础
* linux内核和发行版本 (uname -a 查看版本)
* glibc的版本 (cd /lib)
* mysql-server 安装 (sudo apt-get install mysql-server mysql-client libmysqlclient-dev)

## 解决mysql5.7登录不了
* sudo cat /etc/mysql/debian.cnf (maridb)
* mysql -u账户 -p密码
* use mysql
* update user set authentication_string=password("123456") where user="root"
* flush privileges;
* select user, host, plugin from user where user = "root"
* update user set plugin = "mysql_native_password" where user = "root"
* flush privileges;

## 脚本语言
* shell
* python
* perl

## 数据库基础理论
* 关系型数据库基础理论
* MVCC/ACID/范式设计/索引设计
* 分布式数据库基础理论
* 2PC/CAP/base/paxos算法
* MVCC - mvcc.md
* ACID - acid.md
* 2PC/CAP/base/paxos - algor.md

## MySQL 相关书籍
* 高性能 MySQL
* MySQL技术内幕 InnoDB存储引擎
* 高可用MySQL

## MySQL 相关技能
* 高可用  为什么要高可用　传统主从复制 MHA
* 备份    备份原理 逻辑备份　物理备份　备份工具
* 监控    监控指示 LOAD/CPU/TPS/QPS/IOW/AIT/CONN ZABBIX监控系统
* 高性能  深入理解体系结构 SQL优化

## DBA等级
* 初级 熟练掌握日常操作及数据库开发工作 - 操作方向
* 中级 深入学习MySQL原理　掌握日常MySQ运维及优化　成为独挡一面的DBA - 优化方向
* 高级 数据库架构师　精通MySQL优化　熟悉MySQL大规模运维　指导公司数据库整体架构 - 架构方向
* 资深 深入学习MySQL内核原理　可以快速定位问题　并可以根据需求做一些定制 - 源码方向

## 相关工作
* 运维DBA 负责数据库日常维护　监控数据库状态　搭建高可用集群　备份恢复　数据迁移
* 开发DBA　负责数据库设计　负责应用程序的建表　负责表的索引创建 SQL审计　SQL优化
* 内核DBA 负责数据库内核级开发 处理内核级的相关　跟踪最新数据库内核