package main

import "fmt" // import 用于导入外部代码 fmt包用于格式化输出数据

/*
idea 支持 go (直接使用goLand)

1. 安装go插件
settings ---> plugins ---> Install JetBrains Plugin... ---> go

2. 创建新项目
New ---> Project ---> go ---> 选择 go SDK(/home/ricky/app/go) ---> chat (默认目录 ~/IdeaProjects/chat)

3. 设置字体大小
settings ---> Color Scheme Font ---> 勾选use clolor ... Size:18 ---> Apply

4. 测试项目
chat目录下创建src目录 ---> src目录下创建wel.go ---> 编写wel.go ---> 运行wel.go

https://play.golang.org/ 浏览器编辑并运行
*/
func main() {
	// func Println(a ...interface{}) (n int, err error) ---> Fprintln(os.Stdout, a...)
	fmt.Println("welcome to go worold")
}
