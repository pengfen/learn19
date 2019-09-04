# mysql操作

## 安装mysql
* go get github.com/go-sql-driver/mysql

## 数据表 userinfo
```bash
create table `userinfo` (
	`uid` int not null auto_increment primary key,
	`username` varchar(64) not null default '' comment '用户名',
	`departname` varchar(64) not null default '' comment '部门名',
	`create_time` int not null default 0 comment '创建时间'
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT="用户信息表";
```