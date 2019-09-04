package main

import (
	"fmt"
	"os"
	"strconv"
	"net/http"
)

// 爬取网页内容
func HttpGet(url string) (result string, err error) {
	resp, err := http.Get(url)
	checkErr(err)
	defer resp.Body.Close()

	// 读取网页body内容
	buf := make([]byte, 1024*8)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 { // 读取结束 b或者出问题
			fmt.Println("resp.Body.Read err = ", err)
			break
		}
		result += string(buf[:n])
	}
	return
}

/**
 * @param start 启始页
 * @param end 结束页
 */
func DoWork(start, end int) {
	fmt.Printf("正在爬取%d到%d的页面", start, end)
	for i := start; i <= end; i ++ {
		url := "https://tieba.baidu.com/f?kw=%E9%87%91%E5%BA%B8&ie=utf-8&pn=" + strconv.Itoa((i - 1) * 50)
		fmt.Println("url = ", url)

		// 爬取具体内容
		result, err := HttpGet(url)
		checkErr(err)

		// 把内容写入到文件
		fileName := strconv.Itoa(i) + ".html"
		f, err := os.Create(fileName)
		checkErr(err)

		f.WriteString(result) // 写内容
		f.Close() // 关闭文件
	}
}

/**
 * 处理错误
 */
func checkErr(err error)  {
	if err != nil {
		// panic(err)
		fmt.Println("执行失败", err)
		return
	}
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页(>=1) ---> ")
	fmt.Scan(&start) // 从标准输入os.Stdin读取文本
	fmt.Printf("请输入终止页(>=起始页) ---> ")
	fmt.Scan(&end)
	DoWork(start, end)
}