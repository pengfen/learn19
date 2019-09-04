# 事务的ACID及隔离级别
* 原子性(Atomicity):一个事务必须被视为一个不可分割的最小工作单元，整个事务中的所有操作要么全部提交成功，要么全部失败回滚，对于一个事务来说，不可能只执行其中的一部分操作，这就是事务的原子性
* 一致性(Consistency): 数据库总是从一个一致性的状态转换到另一个一致性的状态。
* 隔离性(Isolation)：一个事务所做的修改在最终提交以前，对其他事务是不可见的。
* 持久性(Durability):一旦事务提交，则其所做的修改不会永久保存到数据库。

## 隔离级别(隔离性有4种级别)
* read-uncommitted 读未提交
* read-committed 读提交
* repeatable-read 可重读
* serializable 可串行化

## 测试隔离级别前准备工作
* cards建表语句如下：
mysql> show create table cards \G;

CREATE TABLE `cards` (
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `name` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
   `balance` int(10) unsigned DEFAULT NULL COMMENT '用户余额',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='银行卡信息表';
insert into cards(name, balance) values('zhangsanfeng', 5000);

mysql> select * from cards;
 +----+--------------+---------+
 | id | name         | balance |
 +----+--------------+---------+
 |  1 | zhangsanfeng |    5000 |
 +----+--------------+---------+

mysql> show global variables like '%iso%'; # 查看当前MySQL系统隔离级别
 +---------------+-----------------+
 | Variable_name | Value           |
 +---------------+-----------------+
 | tx_isolation  | REPEATABLE-READ |
 +---------------+-----------------+

## 测试读未提交
 1 mysql> set tx_isolation='read-uncommitted';
 2 Query OK, 0 rows affected (0.00 sec)
 3 
 4 关闭自动提交，开启事务，模拟转账，将zhangsanfeng账户中的余额减去500
 5 
 6 mysql> set autocommit=0;
 7 Query OK, 0 rows affected (0.00 sec)
 8 mysql> start transaction;
 9 Query OK, 0 rows affected (0.00 sec)
10 mysql> update cards set balance=balance-500 where name='zhangsanfeng';
11 Query OK, 1 row affected (0.00 sec)
12 Rows matched: 1  Changed: 1  Warnings: 0
13 
14 mysql> select * from cards;
15 +----+--------------+---------+
16 | id | name         | balance |
17 +----+--------------+---------+
18 |  1 | zhangsanfeng |    4500 |
19 +----+--------------+---------+
20 1 row in set (0.00 sec)
复制代码
这时我并没有提交(commit),现在去第二个终端读取(这样主要是模拟现实中并发量很大的情况，这时没有提交事务，但是有个请求过来读取)

复制代码
 1 --->set tx_isolation='read-uncommitted';    #由于是两个会话终端，终端一中的设置是基于会话的，在终端二中不生效，所以终端二也要设置隔离级别
 2 Query OK, 0 rows affected (0.00 sec)
 3 
 4 --->select * from cards;
 5 +----+--------------+---------+
 6 | id | name         | balance |
 7 +----+--------------+---------+
 8 |  1 | zhangsanfeng |    4500 |
 9 +----+--------------+---------+
10 1 row in set (0.00 sec)
11 
12 --->
复制代码
可以看到，在终端二中，可以看到终端一提交后的最终结果，但是终端一并没有提交，所以这种隔离效果最差。
在实际情况中，这样就容易产生幻读，就是当一个用户在读取数据时，这时别的终端发起了一个事务修改了数据，导致用户
在前后看到的数据不一致。

同理在read-committed(读提交)情况下，也会发生幻读，只是产生幻读的时间点向后移了，当事务提交后，用户才会发现
读取的数据和之前读取的不同。原理较简单，就不演示了。

复制代码
 1 改变系统隔离级别，调整到(repeatable-read)
 2 终端一：
 3 mysql> set tx_isolation='repeatable-read';
 4 Query OK, 0 rows affected (0.00 sec)
 5 
 6 mysql> select * from cards;
 7 +----+--------------+---------+
 8 | id | name         | balance |
 9 +----+--------------+---------+
10 |  1 | zhangsanfeng |    5000 |
11 +----+--------------+---------+
12 1 row in set (0.00 sec)
13 
14 mysql> 
15 
16 终端二：
17 --->set tx_isolation='repeatable-read';
18 Query OK, 0 rows affected (0.00 sec)
19 
20 --->select * from cards;
21 +----+--------------+---------+
22 | id | name         | balance |
23 +----+--------------+---------+
24 |  1 | zhangsanfeng |    5000 |
25 +----+--------------+---------+
26 1 row in set (0.00 sec)
27 
28 --->
29 
30 
31 终端一开启了一个事务
32 mysql> set autocommit=0;
33 Query OK, 0 rows affected (0.00 sec)
34 
35 mysql> start transaction;
36 Query OK, 0 rows affected (0.00 sec)
37 
38 mysql> update cards set balance=balance-500 where name='zhangsanfeng';
39 Query OK, 1 row affected (0.00 sec)
40 Rows matched: 1  Changed: 1  Warnings: 0
41 
42 mysql> select * from cards;
43 +----+--------------+---------+
44 | id | name         | balance |
45 +----+--------------+---------+
46 |  1 | zhangsanfeng |    4500 |
47 +----+--------------+---------+
48 1 row in set (0.00 sec)
49 
50 mysql> 
51 
52 终端一没有提交事务，终端二查看结果如下：
53 --->select * from cards;
54 +----+--------------+---------+
55 | id | name         | balance |
56 +----+--------------+---------+
57 |  1 | zhangsanfeng |    5000 |
58 +----+--------------+---------+
59 1 row in set (0.01 sec)
60 
61 --->
62 
63 终端一提交后，终端二在读取结果如下：
64 mysql> commit;
65 Query OK, 0 rows affected (0.01 sec)
66 
67 mysql> 
68 
69 
70 --->select * from cards;
71 +----+--------------+---------+
72 | id | name         | balance |
73 +----+--------------+---------+
74 |  1 | zhangsanfeng |    4500 |
75 +----+--------------+---------+
76 1 row in set (0.00 sec)
77 
78 --->
复制代码
这样虽然也可能产生幻读，但是产生幻读的时间点是在另外一个事务结束后，也就是说，事务在未提交之前，其他用户是感受不到的。
所以相对而言，这种隔离效果较之前两种要好，起到了事务的隔离效果，又不影响并发，MySQL中默认就是该级别。

复制代码
 1 接下来测试serializable(串行化)
 2 终端一：
 3 mysql> set tx_isolation='serializable';
 4 Query OK, 0 rows affected (0.00 sec)
 5 
 6 mysql> set autocommit=0;
 7 Query OK, 0 rows affected (0.00 sec)
 8 
 9 mysql> start transaction;
10 Query OK, 0 rows affected (0.01 sec)
11 
12 mysql> update cards set balance=balance-500 where name='zhangsanfeng';
13 Query OK, 1 row affected (0.04 sec)
14 Rows matched: 1  Changed: 1  Warnings: 0
15 
16 mysql> select * from cards;
17 +----+--------------+---------+
18 | id | name         | balance |
19 +----+--------------+---------+
20 |  1 | zhangsanfeng |    4500 |
21 +----+--------------+---------+
22 1 row in set (0.00 sec)
23 
24 mysql> 
25 
26 终端二：
27 --->set tx_isolation='serializable';
28 Query OK, 0 rows affected (0.00 sec)
29 
30 --->select * from cards;
31 +----+--------------+---------+
32 | id | name         | balance |
33 +----+--------------+---------+
34 |  1 | zhangsanfeng |    5000 |
35 +----+--------------+---------+
36 1 row in set (0.00 sec)
37 --->update cards set balance=balance-500 where name='zhangsanfeng';    //这一步可以看到，光标一直卡者，数据的修改不能成功，就是应为上面的事务没有提交。
38 ERROR 1205 (HY000): Lock wait timeout exceeded; try restarting transaction
复制代码
可以看到串行化后，当事务没有结束时，该数据就一直处于锁定状态，其他用户不能操作。这种隔离级别最高，但是不支持并发，
所以在实际情况中一般不会用。

最后总结一下：
事务的隔离级别越高，事务越安全，但是并发能力越差。