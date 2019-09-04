package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"wel/models"
)

// 用户控制器
type UserController struct {
	beego.Controller
}

// 注册功能
func (c *UserController) Register()  {
	// 1. 获取表单数据
	userName := c.GetString("username")
	pwd := c.GetString("password")

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

// 登录界面
func (c *UserController) Login() {
	c.TplName = "user/login.html"
}

// 处理登录逻辑
func (c *UserController) HandleLogin() {
	// c.Ctx.WriteString("这是登录后的POST请求")
	// 1. 拿到数据
	userName := c.GetString("username")
	pwd := c.GetString("password")
	// 2. 判断数据是否合法
	if userName == "" || pwd == "" {
		beego.Info("输入数据不合法")
		c.TplName = "user/login.html"
		return
	}
	// 3. 查询账号密码是否正确
	o := orm.NewOrm()
	user := models.User{}

	user.Name = userName
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("查询失败")
		c.TplName = "user/login.html"
		return
	}
	// 4. 跳转
	//c.Ctx.WriteString("欢迎你 登录成功")
	c.Redirect("/index", 302);
}