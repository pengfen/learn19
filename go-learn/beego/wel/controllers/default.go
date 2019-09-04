package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//// 1. 创建 ORM 对象
	//o := orm.NewOrm();
	//// 2. 有一个要插入数据的结构体对象
	//user := models.User{}
	//// 3. 对结构体赋值
	//user.Name = "ricky"
	//user.Pwd = "ricky"
	//// 4. 插入
	//_, err := o.Insert(&user)
	//if err != nil {
	//	beego.Info("插入失败", err)
	//	return
	//}

	//// 1. 创建ORM对象
	//o := orm.NewOrm()
	//// 2. 需要更新的结构体对象
	//user := models.User{}
	//// 3. 查到需要更新的数据
	//user.Id = 1
	//err := o.Read(&user)
	//
	//// 4. 给数据重新赋值
	//if err == nil {
	//	user.Name = "ricky1"
	//	user.Pwd = "ricky1"
	//	// 5. 更新
	//	_, err = o.Update(&user)
	//	if err != nil {
	//		beego.Info("更新失败", err)
	//		return
	//	}
	//}

	// 查询数据
	// 1. 创建ORM对象
	//o := orm.NewOrm()
	//// 2. 查询对象
	//user := models.User{}
	//// 3. 指定查询对象字段值
	//user.Id = 1
	//// 4. 查询
	//err := o.Read(&user)
	//if err != nil {
	//	beego.Info("查询失败", err)
	//	return
	//}
	//
	//beego.Info("查询成功", user) // 查询成功 {1 ricky1 ricky1}


	//// 1. 创建ORM对象
	//o := orm.NewOrm()
	//// 2. 删除的对象
	//user := models.User{}
	//// 3. 指定删除哪一条数据
	//user.Id = 1
	//// 4. 删除
	//_, err := o.Delete(&user)
	//if err != nil {
	//	beego.Info("删除错误", err)
	//	return
	//}

	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.TplName = "user/register.html"
}
