package main

import (
"fmt"
"html/template"
"io"
"net/http"
)

var myTemplate *template.template

type Result struct {
	output string
}

func (p *Result) Write(b []byte) (n int, err error) {
	fmt.Println("called by template")
	p.output += string(b)
	return len(b), nil
}

type Person struct {
	Name string
	Title string
	Age int
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle wel")

	var arr []Person
	p := Person{Name: "ricky", Age: 10, Title: "网站"}

	arr = append(arr, p)

	resultWriter := &Result{}
	io.WriteString(resultWriter, "welcome to go world")
	err := myTemplate.Execute(w, arr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("模板数据:", resultWriter.output)

	// myTemplate.Execute(w, p)
	// myTemplate.Execute(os.Stdout, p)
	// file, err := os.OpenFile("./test.log", os.O_CREATE|os.O_WRONLY, 0755)
	// if err != nil {
	// 	fmt.Println("打开文件错误", err)
	// 	return
	// }
}

func initTemplate(filename string) (err error) {
	myTemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("解析错误", err)
		return
	}
	return
}

func main() {
	initTemplate("./tem_http.html")
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("端口无效")
	}
}