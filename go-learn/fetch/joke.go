package main

import "fmt"

// 本脚本爬取笑话列表

func SpiderPape(i int)  {
	// 明确爬取url
	// https://www.pengfu.com/xiaohua_1.html
	url := ""
}
func DoWorkJoke(start, end int)  {
	fmt.Printf("准备爬取第%d页到%d页的网址\n", start, end)

	for i := start; i <= end; i++ {
		// 定义一个函数 爬主页面
		SpiderPape(i)
	}
}

func main()  {
	var start, end int
	fmt.Printf("请输入起始页(>=1):")
	fmt.Scan(&start)

	fmt.Printf("请输入终止页(>=起始页):")
	fmt.Scan(&end)

	DoWorkJoke(start, end)
}
