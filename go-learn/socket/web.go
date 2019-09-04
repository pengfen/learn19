package main

import (
    "fmt"
    "net/http"
    "html/template"
    "strings"
    "log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  // 解析参数 默认是不会解析的
	fmt.Println(r.Form) // 服务端打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"]) // 获取url中的参数 ?url_long=2

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "welcome to go world") // 输出到客户端的
}

// 登录页 ---> login.gtpl
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl") // 注意模板导入 html/template
		t.Execute(w, nil)
	} else {
		// 处理请求数据 执行登录的逻辑判断
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])

		// 为空处理
		if len(r.Form["username"][0]) == 0 {}

		// 数字
		getint, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {} // 数字转化出错
		if getint > 100 {} // 判断数值范围

		// 正则处理
		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {return false}

		// 中文
		if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", r.Form.Get("realname")); !m {return false}

		// 英文
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {return false}

		// 电子邮件地址
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
		    fmt.Println("no")
		} else {
		    fmt.Println("yes")
		}

		// 手机号码
		if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
		    return false
		}

		// 下拉菜单
		<select name="fruit">
		    <option value="apple">apple</option>
		    <option value="pear">pear</option>
		    <option value="banane">banane</option>
		</select>

		slice := []string{"apple", "pear", "banane"}

		for _, v := range slice {
			if v == r.Form.Get("fruit") {
				return true
			}
		}

		return false

		// 单选
		<input type="radio" name="gender" value="1">男
		<input type="radio" name="gender" value="2">女

		slice := []int{1,2}

		for _, v := range slice {
			if v == r.Form.Get("gender") {
				return true
			}
		}

		return false

		// 复选框
		<input type="checkbox" name="interest" value="football">足球
		<input type="checkbox" name="interest" value="basketball">篮球
		<input type="checkbox" name="interest" value="tennis">网球

		slice := []string{"football", "basketball", "tennis"}
		a := Slice_diff(r.Form["interest"], slice)
		if a == nil {
			return true
		}

		return false

		// 日期和时间
		t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
		fmt.Printf("%s\n", t.Local())

		// 身份证号码
		// 验证15位身份证 15位全是数字
		if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); !m {
			return false
		}

		// 验证18位身份证 18位前17位为数字 最后一位是校验位 数字或X
		if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
			return false
		}
	}
}

/*
cmd 命令行运行
go run web.go
http://localhost:9090/
*/
func main() {
	http.HandleFunc("/", sayhelloName) // 设置访问的路由
	http.HandleFunc("/login", login) // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}