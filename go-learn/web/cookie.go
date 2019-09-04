package main

import (
	"net/http"
	"time"
)

type Cookie struct {
	Name string
	Value string
	Path string
	Domain string
	Expires time.Time
	RawExpires string

	// MaxAge=0 没有指定MaxAge属性
	// MaxAge<0 删除cookie 等价于MaxAge=0
	// MaxAge>0 以称为单位的有效值
	MaxAge int
	Secure bool
	HttpOnly bool
	Raw string
	Unparsed []string
}

func sayWel(w http.ResponseWriter, req *http.Request)  {
	w.Write([]byte("welcome"))
}

func ReadCookieServer(w http.ResponseWriter,  req *http.Request) {
	cookie, err := req.Cookie("test_cookie_name")
	if err == nil {
		cookievalue := cookie.Value
		w.Write([]byte("cookie的值是: " + cookievalue))
	} else {
		w.Write([]byte("<b>读取出现错误: " + err.Error()))
	}
}

func WriteCookieServer(w http.ResponseWriter, req *http.Request)  {
	cookie := http.Cookie{Name: "test_cookie_name", Value: "test_cookie_name", Path:"/", MaxAge:86400}
	http.SetCookie(w, &cookie)
	w.Write([]byte("<b>设置cookie成功"))
}

func DeleteCookieServer(w http.ResponseWriter, req *http.Request)  {
	cookie := http.Cookie{Name: "test_cookie_name", Path:"/", MaxAge: -1}
	http.SetCookie(w, &cookie)
	w.Write([]byte("<b>删除cookie成功"))
}
func main() {
	http.HandleFunc("/", sayWel)
		http.HandleFunc("/readcookie", ReadCookieServer) // http://localhost:3000/readcookie
	http.HandleFunc("/writecookie", WriteCookieServer) // http://localhost:3000/writecookie
	http.HandleFunc("/deletecookie", DeleteCookieServer) // http://localhost:3000/deletecookie
	http.ListenAndServe(":3000", nil)
}