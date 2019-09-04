package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

/*
http 编程

go 原生支持http import("net/http")
go 的http服务性能和nginx比较接近
几行代码就可以实现一个web服务

http.go
http_client.go

http常见请求方法
get请求
post请求
put请求
delete请求
head请求
http_head.go

http常见状态码
http.StatusContinue = 100
http.StatusOK = 200
http.StatusFound = 302
http.StatusBadRequest = 400
http.StatusUnauthorized = 401
http.StatusForbidden = 403
http.StatusNotFound = 404
http.StatusInternalServerError = 500
*/

func say(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm() // 解析参数 默认是不会解析的 对于POST则解析响应包的主体(request body)
	// 注意 如果没有调用ParseForm方法 下面无法获取表单的数据
	fmt.Println(r.Form) // 输出到服务端
	fmt.Println("path", r.URL.Path)
	fmt.Println("schema", r.URL.Scheme)
	fmt.Println(r.Form["url_long"]) // 获取url中的参数 ?url_long=2
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "welcome to go web world") // 写在浏览器
}

func test_http()  { // http://localhost:9090/
	http.HandleFunc("/", say)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func test_http1()  {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	defer conn.Close()
	msg := "GET /HTTP/1.1\r\n"
	msg += "Host:www.baidu.com\r\n"
	msg += "Connection:close\r\n"
	msg += "\r\n\r\n"

	_, err = io.WriteString(conn, msg)
	if err != nil {
		fmt.Println("write string failed,", err)
		return
	}
	buf := make([]byte, 4096)
	for {
		count, err := conn.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[0:count]))
	}
}

// 登录页 ---> login.gtpl
func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method: ", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl") // 注意模板导入 html/template
		t.Execute(w, nil)
	} else {
		// 处理请求数据 执行登录的逻辑判断
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])

		// 为空处理
		// if len(r.Form["username"][0]) == 0 {}

		// 数字
		//getint, err := strconv.Atoi(r.Form.Get("age"))
		//if err != nil {} // 数字转化出错
		//if getint > 100 {} // 判断数值范围

		// 正则处理
		//if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {return false}
		//
		//// 中文
		//if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", r.Form.Get("realname")); !m {return false}
		//
		//
		//// 英文
		//if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {return false}
		//
		//// 电子邮件地址
		//if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
		//	fmt.Println("no")
		//} else {
		//	fmt.Println("yes")
		//}
		//
		//// 手机号码
		//if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
		//	return false
		//}
		//
		//// 下拉菜单
		//<select name="fruit">
		//<option value="apple">apple</option>
		//<option value="pear">pear</option>
		//<option value="banane">banane</option>
		//</select>
		//
		//slice := []string{"apple", "pear", "banane"}
		//
		//for _, v := range slice {
		//	if v == r.Form.Get("fruit") {
		//		return true
		//	}
		//}
		//
		//return false
		//
		//// 单选
		//<input type="radio" name="gender" value="1">男
		//<input type="radio" name="gender" value="2">女
		//
		//slice := []int{1,2}
		//
		//for _, v := range slice {
		//	if v == r.Form.Get("gender") {
		//		return true
		//	}
		//}
		//
		//return false
		//
		//// 复选框
		//<input type="checkbox" name="interest" value="football">足球
		//<input type="checkbox" name="interest" value="basketball">篮球
		//<input type="checkbox" name="interest" value="tennis">网球
		//
		//slice := []string{"football", "basketball", "tennis"}
		//a := Slice_diff(r.Form["interest"], slice)
		//if a == nil {
		//	return true
		//}
		//
		//return false
		//
		//// 日期和时间
		//t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
		//fmt.Printf("%s\n", t.Local())
		//
		//// 身份证号码
		//// 验证15位身份证 15位全是数字
		//if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); !m {
		//	return false
		//}
		//
		//// 验证18位身份证 18位前17位为数字 最后一位是校验位 数字或X
		//if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
		//	return false
		//}
	}
}

func test_login()  {
	http.HandleFunc("/", say)
	http.HandleFunc("/login", login) // 设置访问路由
	err := http.ListenAndServe(":9090", nil) // 设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

type MyMux struct {

}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path == "/" {
		say(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func test_http3()  {
	/*
	1. 首先调用Http.HandleFunc(调用DefaultServerMux的HandleFunc  调用DefaultServerMux的Handle  往DefaultServeMux的map[string]muxEntry中增加对应的handler和路由规则)
	2. 调用http.ListenAndServe(":9090", nil) (实例化Server  调用Server的ListenAndServe()  调用net.Listen("tcp", addr)监听端口
	启动一个for循环 在循环体中Accept请求  对每个请求实例化一个Conn 并且开启一个goroutine为这个请求进行服务c.serve()
	读取每个请求的内容w, err := c.readRequest()  判断handler是否为空 如果没有设置handler handler就设置为DefaultServeMux
	调用handler的ServeHttp  根据request选择handler 并且进入到这个handler的ServeHTTP mux.handler(r).ServeHTTP(w, r)
	)
	选择handler
	a. 判断是否有路由能满足这个request(循环遍历ServerMux的muxEntry)
	b. 如果有路由满足 调用这个路由handler的ServeHttp
	c. 如果没有路由满足 调用NotFoundHandler的ServeHttp
	*/
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}

func test_http4()  {
	resp, err := http.Post("http://www.baidu.com", "application/x-www-form-urlencoded", strings.NewReader("id=1"))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func test_http5()  {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to here"))
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}

/*
命令行运行
go run web.go
http://localhost:9090/
*/
func main()  {
	// test_http()

	// test_http1()

	// test_login()

	// test_http3()

	test_http4()

	test_http5()
}
