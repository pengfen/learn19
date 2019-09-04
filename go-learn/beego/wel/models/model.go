package models

// go get github.com/astaxie/beego/orm

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    "time"
)

// 用户表
type User struct {
    Id int
    Name string
    Pwd string
}

// 文章结构体
type Article struct {
    Id int
    Aname string
    Atime time.Time
    Acount int
    Acontent string
    Aimg string

    //ArticleType*ArticleType `orm:"rel(fk)"`
    //User []*User `orm:"reverse(many)"`
}

//类型表
type ArticleType struct {
    Id int
    Tname string
    //Article []*Article `orm:"reverse(many)"`
}

func init() {
    // 设置数据库基本信息
    // 别名 驱动名 数据源(用户名:密码@tcp(IP:端口号)/数据库名?字符集)
    // create database beego charset=utf8
    orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/beego?charset=utf8")
    // 映射数据表
    orm.RegisterModel(new(User), new(Article), new(ArticleType))
    //
    orm.RunSyncdb("default", false, true)
}