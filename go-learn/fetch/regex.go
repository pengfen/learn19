package main

import (
	"fmt"
	"regexp"
)

const text = "My email is caopeng8787@163.com"

func main()  {
	// .+ 匹配一个至多个字符
	//re := regexp.MustCompile(".+@.+\\..+")
	//re := regexp.MustCompile(`.+@.+\..+`)
	re := regexp.MustCompile(`[a-zA-Z0-9]@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := re.FindString(text)
	// 匹配多条记录
	//match := re.FindAllString(text, -1)
	// 正则中使用() ([a-zA-Z0-9])@([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)
	//for _, m := range match {
	//	fmt.Println(m)
	//}
	fmt.Println(match)
}
