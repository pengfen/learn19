package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

// 注意转码 gbk ---> utf8
// 安装gopm ---> go get -v -u github.com/gpmgo/gopm
// gopm get -g -v golang.org/x/text
// 判断网页编码
// gopm get -g -v golang.org/x/net/html
func main_v1()  {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	// 如果网页编码为gbk 则转成utf8
	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	//all, err := ioutil.ReadAll(utf8Reader)

	//e := determineEncoding(resp.Body)
	//utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s/n", all)
	printCityList(all)
}

func printCityList(contens []byte)  {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	//matches := re.FindAll(contens, -1)
	// 匹配()子元素
	matches := re.FindAllSubmatch(contens, -1)
	for _, m := range matches {
		//fmt.Printf("%s\n", m)
		fmt.Printf("City: %s, URL: %s", m[2], m[1])
		//for _, subMatch := range m {
		//	fmt.Printf("%s ", subMatch)
		//}
		fmt.Println()
	}

	// 一共多少个
	//fmt.Printf("Matches found: %d\n", len(matches))
}

// 获取网页编码
func determineEncoding(r io.Reader) encoding.Encoding  {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}