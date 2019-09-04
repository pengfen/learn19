package main

import (
"fmt"
"net/http"
)

func wel(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcome to go http world")
	fmt.Fprintf(w, "wel")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	fmt.Fprintf(w, "login ")
}

func history(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	fmt.Fprintf(w, "login ")
}

/*
1. 启动服务 go run http.go
2. 页面访问 http://localhost:8880/
*/
func main() {
	http.HandleFunc("/", wel)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http连接失败")
	}
}