# ORM

* O object 对象   M mapping 映射   R relation 关系
* 1. 能够通过结构体对象来操作对应的数据库表
* 2. 能通过ORM生成结构体相对应的数据库表
* 安装ORM  go get github.com/astaxie/beego/orm

```bash
// 设置数据库基本信息
orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test1?charset=utf8")
// 映射 model 数据
orm.RegisterModel(new(Book))
// 生成表
orm.RunSyncdb("default", false, true)
```

## ORM 连接
* 入口文件中导入模型  	import _ "wel/models"
* 建立对应model
```bash
package models

// go get github.com/astaxie/beego/orm

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    Id int
    Name string
    Pwd string
}

func init() {
    // 设置数据库基本信息
    // 别名 驱动名 数据源(用户名:密码@tcp(IP:端口号)/数据库名?字符集)
    // create database beego charset=utf8
    orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/beego?charset=utf8")
    // 映射数据表
    orm.RegisterModel(new(User))
    //
    orm.RunSyncdb("default", false, true)
}
```

## 插入数据
```bash
// 1. 创建 ORM 对象
o := orm.NewOrm();
// 2. 有一个要插入数据的结构体对象
user := models.User{}
// 3. 对结构体赋值
user.Name = "ricky"
user.Pwd = "ricky"
// 4. 插入
_, err := o.Insert(&user)
if err != nil {
	beego.Info("插入失败", err)
	return
}
```

## 更新数据
```bash
// 1. 创建ORM对象
o := orm.NewOrm()
// 2. 需要更新的结构体对象
user := models.User{}
// 3. 查到需要更新的数据
user.Id = 1
err := o.Read(&user)

// 4. 给数据重新赋值
if err == nil {
    user.Name = "ricky1"
    user.Pwd = "ricky1"
    // 5. 更新
    _, err = o.Update(&user)
    if err != nil {
        beego.Info("更新失败", err)
        return
    }
}
```

## 删除数据
```bash
// 1. 创建ORM对象
o := orm.NewOrm()
// 2. 删除的对象
user := models.User{}
// 3. 指定删除哪一条数据
user.Id = 1
// 4. 删除
_, err := o.Delete(&user)
if err != nil {
    beego.Info("删除错误", err)
    return
}
```

## 查询数据
```bash
// 查询数据
// 1. 创建ORM对象
o := orm.NewOrm()
// 2. 查询对象
user := models.User{}
// 3. 指定查询对象字段值
user.Id = 1
// 4. 查询
err := o.Read(&user)
if err != nil {
    beego.Info("查询失败", err)
    return
}

beego.Info("查询成功", user) // 查询成功 {1 ricky1 ricky1}
```