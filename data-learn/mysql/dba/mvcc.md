# MVCC 多版本并发控制

* mysql的innodb采用的是行锁，而且采用了多版本并发控制来提高读操作的性能。
* 多版本并发控制呢 其实就是在每一行记录的后面增加两个隐藏列，记录创建版本号和删除版本号，
* 而每一个事务在启动的时候，都有一个唯一的递增的版本号。 

## 增删改查时版本号
1、在插入操作时 ： 记录的创建版本号就是事务版本号。 

比如我插入一条记录, 事务id 假设是1 ，那么记录如下：也就是说，创建版本号就是事务版本号。

id　name　　	create version　　	delete version　　
1	test　　	1

2、在更新操作的时候，采用的是先标记旧的那行记录为已删除，并且删除版本号是事务版本号，然后插入一行新的记录的方式。 

比如，针对上面那行记录，事务Id为2 要把name字段更新

update table set name= 'new_value' where id=1;

id    name       create version　delete version　　
1     test　　	 1	             2　　　　　　　　
1     new_value  2	 

3、删除操作的时候，就把事务版本号作为删除版本号。比如

delete from table where id=1; 

id　name       create version  delete version　　
1	new_value  2	           3　　

4、查询操作： 

从上面的描述可以看到，在查询时要符合以下两个条件的记录才能被事务查询出来： 

1) 删除版本号 大于 当前事务版本号，就是说删除操作是在当前事务启动之后做的。 

2) 创建版本号 小于或者等于 当前事务版本号 ，就是说记录创建是在事务中（等于的情况）或者事务启动之前。

这样就保证了各个事务互不影响。从这里也可以体会到一种提高系统性能的思路，就是： 

通过版本号来减少锁的争用。

另外，只有read-committed和 repeatable-read 两种事务隔离级别才能使用mvcc

read-uncommited由于是读到未提交的，所以不存在版本的问题

而serializable 则会对所有读取的行加锁。 