## 数据驱动

* 从数据文件中读取测试数据　驱动测试过程的一种测试方法
* 数据驱动可以理解为更高级的参数化
* 测试数据与测试代码分离
* 数据控制过程
* 好处　减少测试代码量　降低脚本开发和维护的成本　便于用例的修改和维护(不用修改代码)
* 要求　较强的代码能力　较强的分层架构设计思维　对开发框架要有一定的了解

## 数据驱动的使用场景
* 复杂的业务流程
* 根据业务场景分流
* 符合条件的并发场景

## Jmeter中的数据驱动
* 控制方式　参数化　逻辑控制器
* id  name      sex  age
* 1   zhangsan  0    20
* 2   lisi      1    25
* 3   wangwu    0    22
* 4   maliu     0    21
```
create table user(
	`id` int unsigned not null auto_increment primary key,
	`name` varchar(255) not null default '' comment '名字',
	`sex` tinyint unsigned not null default 0 comment '性别',
	`age` int unsigned not null default 0 comment '年龄'
	) engine=innodb default charset=utf8 comment='测试用户表';
	insert into user(name, sex, age) values('zhangsan', 0, 20),('lisi', 1, 25),('wangwa', 0, 22),('maliu', 0, 21);
```

* 添加JDBC Connection Configuration  Add ---> Config Element ---> JDBC Connection Configuration
* 添加JDBC Request  Add ---> Sampler ---> JDBC Request
* 添加Debug Sampler Add ---> Sampler ---> Debug Sampler
* 添加View Results Tree  Add ---> Listener ---> View Results Tree
* Database URL: jdbc:mysql://127.0.0.1:3306/test?serverTimezone=UTC&autoReconnect=true
* JDBC Driver class: com.mysql.jdbc.Driver
* variable Name for created pool: users
* 注意　JDBC Request  variable Name of Pool declared in JDBC Connection Configuration: users 与上面的一样

## JDBC Request 操作
* select tele_id, tele_name from fri_users;
* variable names: tele_id
* Result variable name: teleArray
* 可在 View Results Tree ---> Debug Sampler ---> Response data面板中查看
