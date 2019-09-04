package main

import (
"fmt"
_"github.com/go-sql-driver/mysql"
"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId int `db:"user_id"`
	Username string `db:"username"`
	Sex string `db:"sex"`
	Email string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City string `db:"city"`
	TelCode int `db:"telcode"`
}

var Db *sqlx.DB
func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go")
	if err != nil {
		fmt.Println("连接mysql失败", err)
		return
	}
	Db = database
}

func main() {
	r, err := Db.Exec("insert into person(username, sex, email) values(?,?,?)", "stu001", "man", "caopeng8787@163.com")
	if err != nil {
		fmt.Println("执行失败", err)
		return
	}

	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("执行失败", err)
	}

	fmt.Println("插入成功", id)
}