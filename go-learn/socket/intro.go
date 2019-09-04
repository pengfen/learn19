package main // main包 go程序入口包

import "net/http" // http基础封装和访问

func main() { // 程序执行的入口函数main()
	// 文件服务器将当前目录作为根目录
	http.Handle("/", http.FileServer(http.Dir(".")))
	// 默认端口8080
	http.ListenAndServe(":8080", nil)
}