# 注册功能

# 注册界面
* router.go 添加路由 beego.Router("/", &controllers.MainController{})
* default.go c.TplName = "register.html"

# 注册数据添加入库
* 配置路由 beego.Router("/register", &controllers.UserController{}, "post:Register")

```bash
func (c *UserController) Register()  {
	// 1. 获取表单数据
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")

	// 2. 对数据进行校验
	if userName == "" || pwd == "" {
		beego.Info("数据不能为空")
		c.Redirect("/register", 302)
		return
	}
	// 3. 插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Pwd = pwd
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据失败")
		c.Redirect("/register", 302)
		return
	}

	// 4. 返回登录页
	c.Ctx.WriteString("注册成功")
}
```