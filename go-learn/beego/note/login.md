# 登录功能

## 登录界面
* 配置路由 beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")

```bash
func (c *UserController) ShowLogin() {
	c.TplName = "login.html"
}
```

## 处理登录逻辑
```bash
func (c *UserController) HandleLogin() {
	// c.Ctx.WriteString("这是登录后的POST请求")
	// 1. 拿到数据
	userName := c.GetString("username")
	pwd := c.GetString("password")
	// 2. 判断数据是否合法
	if userName == "" || pwd == "" {
		beego.Info("输入数据不合法")
		c.TplName = "login.html"
		return
	}
	// 3. 查询账号密码是否正确
	o := orm.NewOrm()
	user := models.User{}

	user.Name = userName
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("查询失败")
		c.TplName = "login.html"
		return
	}
	// 4. 跳转
	//c.Ctx.WriteString("欢迎你 登录成功")
	c.Redirect("/index", 302);
}
```