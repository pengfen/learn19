package main

import (
	"database/sql"
	"fmt"
	"time"
	_"github.com/go-sql-driver/mysql"
)

/*
create table `userinfo` (
`uid` int not null auto_increment primary key,
`username` varchar(64) not null default '' comment '用户名',
`departname` varchar(64) not null default '' comment '部门名',
`create_time` int(11) not null default 0 comment '创建时间',
) engine=innodb default charset=utf8 comment='用户信息表';

create table `userdetail` (
`uid` int not null default 0 primary key,
`intro` text null comment '用户介绍',
`profile` text null comment '用户详情'
) engine=innodb default charset=utf8 comment='用户详情表';
*/

func checkErr1(err error)  {
	if err != nil {
		panic(err)
	}
}
func main()  {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/beego?charset=utf8")
	checkErr1(err)

	// 添加数据
	stmt, err := db.Prepare("insert userinfo set username=?, departname=?, create_time=?")
	checkErr1(err)

	res, err := stmt.Exec("astaxie", "研发部门", time.Now().Unix())
	checkErr1(err)

	id, err := res.LastInsertId()
	checkErr1(err)

	fmt.Println(id) // 3


	// 更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=? ")
	checkErr1(err)

	res, err = stmt.Exec("asupdate", id)
	checkErr1(err)

	affect, err := res.RowsAffected()
	checkErr1(err)

	fmt.Println(affect) // 1


	// 查询数据
	rows, err := db.Query("select * from userinfo")
	checkErr1(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var create_time string
		err = rows.Scan(&uid, &username, &department, &create_time)
		checkErr1(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(create_time)
	}

	// 删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr1(err)

	res, err = stmt.Exec(id)
	checkErr1(err)

	affect, err = res.RowsAffected()
	checkErr1(err)

	fmt.Println(affect) // 1
	db.Close()
}
