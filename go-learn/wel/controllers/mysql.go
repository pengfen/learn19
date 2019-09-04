package controllers

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
    _"github.com/go-sql-driver/mysql"
	"time"
)

// 文章控制器
type MysqlController struct {
	beego.Controller
}

/**
 * 处理错误
 */
func checkErr(err error)  {
	if err != nil {
		// panic(err)
		fmt.Println("执行失败", err)
		return
	}
}

// 处理添加操作
func (c *MysqlController) Add() {
	// 打开数据库 beego
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/beego?charset=utf8")
	checkErr(err) // 调用checkErr函数

	// 添加数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?, departname=?, create_time=?")
	checkErr(err)

	// Unix() 当前时间转换成秒
	res, err := stmt.Exec("test", "people", time.Now().Unix())
	checkErr(err)
	// res, err := Db.Exec("insert into userinfo(username, departname, create_time) values(?,?,?)", "test", "people", time.Now().Unix())

	// 记录添加ID
	id, err := res.LastInsertId()
	checkErr(err)

	beego.Info(id)
	db.Close()
	c.Ctx.WriteString("添加成功")
}

// 处理更新操作
func (c *MysqlController) Update() {
	// 打开数据库 beego
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/beego?charset=utf8")
	checkErr(err) // 调用checkErr函数

	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err := stmt.Exec("man", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	beego.Info(affect)
	db.Close()
	c.Ctx.WriteString("修改成功")
}

// 查询数据
func (c *MysqlController) Select() {
	// 打开数据库 beego
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/beego?charset=utf8")
	checkErr(err) // 调用checkErr函数

	// 查询数据
	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var create_time int
		err = rows.Scan(&uid, &username, &department, &create_time)
		checkErr(err)

		beego.Info(uid)
		beego.Info(username)
		beego.Info(department)
		beego.Info(create_time)
	}
	db.Close()
	c.Ctx.WriteString("查询成功")
}

// 删除数据
func (c *MysqlController) Delete() {
	// 打开数据库 beego
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/beego?charset=utf8")
	checkErr(err) // 调用checkErr函数

	// 删除数据
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err := stmt.Exec(1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	beego.Info(affect)
	db.Close()
	c.Ctx.WriteString("删除成功")
}