package main

import (
    "fmt"
)

// 爬取网页内容
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	// 读取网页body内容
	buf := make([]byte, 1024*4)
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

func DoWork(start, end int) {
	fmt.Printf("正在爬取%d到%d的页面", start, end)
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E9%87%91%E5%BA%B8&ie=utf-8&pn=" + strconv.Itoa((i - 1) * 50)
		fmt.Println("url = ", url)

		// 爬
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("HttpGet err = ", err)
			continue
		}

		// 把内容写入到文件
		fileName := strconv.Itoa(i) + ".html"
		f, err1 := os.Create(fileName)
		if err1 != nil {
			fmt.Println("os.Create err1 = ", err1)
			continue
		}
		f.WriteString(result) // 写内容
		f.Close() // 关闭文件
	}

	
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页(>=1):")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页(>=超始页):")
	fmt.Scan(&end)
	DoWork(start, end)
}