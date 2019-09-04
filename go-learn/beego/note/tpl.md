# TplName
* c.Data["errmsg"] = "ricky"
* c.TplName = "login.html"
* 指定视图文件 同时可以给这个视图传递一些数据

# Redirect
* c.Redirect("/login", 302)
* 跳转 不有传递数据 速度快
* 第一个参数是 url 地址
* 第二个参数是 HTTP 的状态码

* 1XX 请求已经被接收 需要继续发送请求 100
* 2XX 请求成功 200
* 3XX 请求资被转移 请求被转接 302
* 4XX 请求失败 404
* 5XX 服务器错误 500