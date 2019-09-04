package main

import (
	"fmt"
	"strings"
)

/*
字符串 strings的使用
1. strings.HasPrefix(s string, prefix string) bool 判断字符串s是否以prefix开头
练习1 判断一个url是否以http://开头 如果不是 则加上http://

2. strings.HasSuffix(s string, suffix string) bool 判断字符串s是否以suffix结尾
练习2 判断一个路径是否以"/"结尾 如果不是 则加上/

3. strings.Index(s string, str string) int 判断str在s中首次出现的位置 如果没有出现 则返回-1
4. strings.LastIndex(s string, str string) int 判断str在s中最后出现的位置 如果没有出现 则返回-1
练习3 写一个函数返回一个字符串在另一个字符串的首次出现和最后出现位置
func StrIndex(str string, substr string) (int, int){}

5. strings.Replace(str string, old string, new string, n int) 字符串替换 int 字符串计数(替换次数)
6. strings.Count(str string, substr string) substr 重复count次str
7. strings.Repeat(str string, count int) 字符串连接 int 字符串次数
8. strings.ToLower(str string) string 转为小写
9. strings.ToUpper(str string) string 转为大写
练习4 写一个函数分别演示Replace Count Repeat ToLower ToUpper的用法

10. strings.TrimSpace(str string) 去掉字符串首尾空白字符
11. strings.Trim(str string, cut string) 去掉字符串首尾cut字符
12. strings.TrimLeft(str string, cut string) 去掉字符串首cut字符
13. strings.TrimRight(str string, cut string) 去年字符串首cut字符
14. strings.Field(str string) 返回str空格分隔的所有子串的slice
15. strings.Split(str string, split string) 返回str split分隔的所有子串的slice
16. strings.Join(s1 []string, sep string) 用sep把s1中的所有元素链接起来
练习5 写一个函数分别演示TrimSpace Trim TrimLeft TrimRight Field Split 以及Join的用法

17. strconv.Itoa(i int) 把一个整数i转成字符串
18. strconv.Atoi(str string) (int, error) 把一个字符串转成整数
练习6 写一个函数分别演示Itoa, Atoi的用法
*/

func test_prog() {
	//var (
	//	url string
	//	path string
	//)
	//
	//fmt.Scanf("%s%s", &url, &path)
	//url = urlProcess(url)
	//fmt.Println(url)

	var s string = "http://www.res.com"
	var pre string = "http://"
	if strings.HasPrefix(s, pre) {
		fmt.Println("包括")
	} else {
		fmt.Println("不包括")
	}
}

func urlProcess(url string) string {
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}
	return url
}

func test_prog1()  {
	//var (
	//	url string
	//	path string
	//)
	//
	//fmt.Scanf("%s %s", &url, &path)
	//path = pathProcess(path)
	//fmt.Println(path)

	var s string = "http://www.res.com"
	var suf string = "/"
	if (strings.HasSuffix(s, suf)) {
		fmt.Println("包括")
	} else {
		fmt.Println("不包括")
	}
}

func pathProcess(path string) string {
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}
	return path
}

func test_prog2()  {
	var s string = "http://www.res.com"
	var index string = "w"
	fmt.Println(strings.Index(s, index))
	fmt.Println(strings.LastIndex(s, index))
}

func test_prog3()  {
	var s string = "http://www.res.com"
	var index string = "w"

	fmt.Println(strings.Replace(s, index, "e", 1))
	fmt.Println(strings.Count(s, index))
	fmt.Println(s + strings.Repeat("a", 2))

	str := "welcome to go world"
	result := strings.Replace(str, "to", "too", 1)
	fmt.Println("Replace: ", result)

	count := strings.Count(str, "1")
	fmt.Println("count: ", count)

	result = strings.Repeat(str, 3)
	fmt.Println("repeat: ", result)

	result = strings.ToLower(str)
	fmt.Println("lower: ", result)

	result = strings.ToUpper(str)
	fmt.Println("upper: ", result)
}

func test_prog4()  {
	str := "welcome to go world"

	result := strings.TrimSpace(str)
	fmt.Println("trimSpace: ", result)

	result = strings.Trim(str, "\n\r")
	fmt.Println("trim: ", result)

	result = strings.TrimLeft(str, "\n\r")
	fmt.Println("trimLeft: ", result)

	result = strings.TrimRight(str, "\n\r")
	fmt.Println("trimRight: ", result)

	splitResult := strings.Fields(str)
	for i := 0; i < len(splitResult); i ++ {
		fmt.Println(splitResult[i])
	}

	splitResult = strings.Split(str, "1")
	for i := 0; i < len(splitResult); i ++ {
		fmt.Println(splitResult[i])
	}

	str2 := strings.Join(splitResult, "1")
	fmt.Println("join: ", str2)
}

func test_strs()  {
	f := Adder()
	fmt.Println(f(100))
	fmt.Println(f(1000))

	f1 := makeSuffix(".bmp")
	fmt.Println(f1("test"))
	fmt.Println(f1("pic"))

	f2 := makeSuffix(".jpg")
	fmt.Println(f2("test"))
	fmt.Println(f2("pic"))
}

func Adder() func(int) int {
	var x int
	f := func(d int) int {
		x += d
		return x
	}
	return f
}

func makeSuffix(suffix string) func(string) string {
	f := func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
	return f
}

func main()  {
	// test_prog() // 练习一

	// test_prog1() // 练习二

	// test_prog2() // 练习三

	// test_prog3() // 练习四

	// test_prog4() // 练习五

	test_strs()
}