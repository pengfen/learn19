package main

// 模板 替换 {{.字段名}}

import (
"fmt"
"os"
"text/template"
)

type Person struct {
	Name string
	age string
	Title string
}

func main() {
	t, err := template.ParseFiles("./tem.html")
	if err != nil {
		fmt.Println("解析错误:", err)
		return
	}

	p := Person{Name:"ricky", age:"31", Title: "测试网站"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("错误:", err.Error())
	}
}