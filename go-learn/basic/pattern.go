package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func test_pattern()  {
	// reg := regexp.MustCompile("\\w+") 正则表达式中的\需要转义
	reg := regexp.MustCompile(`^z.*1$`)

	result := reg.FindAllString("ricky", -1)
	fmt.Printf("%v", result)

	reg1 := regexp.MustCompile(`^z(.*)l$`)

	result1 := reg1.FindAllString("zhangsand", -1)
	fmt.Printf("%v\n", result1)

	reg2 := regexp.MustCompile(`^z(.{1})(.{1})(.*)l$`)

	result2 := reg2.FindAllStringSubmatch("zhangsanl", -1)
	fmt.Printf("%v\n", result2)
}

func test_pattern1()  {
	url := "https://movie.douban.com/subject/24751763/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	sHtml, _ := ioutil.ReadAll(resp.Body)

	reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*)</span>`)
	result := reg.FindAllStringSubmatch(string(sHtml), -1)

	fmt.Println(result[0][1])

	reg1 := regexp.MustCompile(`<strong\s*class="ll\s*rating_num"\s*property="v:average">(.*)</strong>`)
	result1 := reg1.FindAllStringSubmatch(string(sHtml), -1)

	fmt.Println(result1[0][1])
}

func main() {
	test_pattern()

	test_pattern1()
}
