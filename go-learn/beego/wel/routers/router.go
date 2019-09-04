package routers

import (
	"wel/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //beego.Router("/register", &controllers.MainController{})
    beego.Router("/register", &controllers.UserController{}, "post:Register")
    // 注意 当实现了自定义的get请求方法 请求将不会访问默认方法
	beego.Router("/login", &controllers.UserController{}, "get:Login;post:HandleLogin")
    beego.Router("/index", &controllers.ArticleController{}, "get:Index")
    beego.Router("/article/add", &controllers.ArticleController{}, "get:Add;post:HandlerAdd")
	beego.Router("/article/detail", &controllers.ArticleController{}, "get:Detail")
	beego.Router("/article/update", &controllers.ArticleController{}, "get:Update;post:HandlerUpd")
    beego.Router("/article/delete", &controllers.ArticleController{}, "get:HandleDelete")

    // mysql 操作
    beego.Router("/mysql/add", &controllers.MysqlController{}, "get:Add")
    beego.Router("/mysql/update", &controllers.MysqlController{}, "get:Update")
    beego.Router("/mysql/select", &controllers.MysqlController{}, "get:Select")
    beego.Router("/mysql/delete", &controllers.MysqlController{}, "get:Delete")

    // redis 操作
    beego.Router("/redis/write", &controllers.RedisController{}, "get:Write")
    beego.Router("/redis/check", &controllers.RedisController{}, "get:Check")
    beego.Router("/redis/delete", &controllers.RedisController{}, "get:Delete")
    beego.Router("/redis/json", &controllers.RedisController{}, "get:Json")
}
