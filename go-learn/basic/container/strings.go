package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func reverse(str string) string {
	var result string
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result = result + fmt.Sprintf("%c", str[strLen - i - 1])
	}
	return result
}

func reverse1(str string) string {
	var result []type
	tmp := []byte(str)
	length := len(str)
	for i := 0; i < length; i++ {
		result = append(result, tmp[length - i - 1])
	}
	return string(result)
}

func urlProcess(url string) string {
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}

	return url
}

func pathProcess(path string) string {
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}

	return path
}

func main() {
	var (
		url string
		path string
	)

	fmt.Scanf("%s%s", &url, &path)
	url = urlProcess(url)
	fmt.Println(url)

	path = pathProcess(path)
	fmt.Println(path)

	var s  string = "http://www.res.com"
	var pre string = "http://"
	if (strings.HasPrefix(s, pre)) {
		fmt.Println("包括")
	} else {
		fmt.Println("不包括")
	}

	var suf string = "/"
	if (strings.HasSuffix(s, suf)) {
		fmt.Println("包括")
	} else {
		fmt.Println("不包括")
	}

	var index string = "w"
	fmt.Println(strings.Index(s, index)) // 7
	fmt.Println(strings.LastIndex(s, index)) // 9

	fmt.Println(strings.Replace(s, index, "e", 1)) //
	fmt.Println(strings.Count(s, index))
	fmt.Println(s + strings.Repeat("a", 2))


	var str1 = "welcome"
	str2 := "world"
	str3 := fmt.Sprintf("%s %s", str1, str2)
	n := len(str3)

	fmt.Println(str3)

	fmt.Printf("len(str3)=%d\n", n)

	substr := str3[0:5]
	fmt.Println(substr)

	substr = str3[6:]
	fmt.Println(substr)

	result := reverse(str3)
	fmt.Println(result)

	result = reverse1(result)
	fmt.Println(result)



	str := "welcome to go world \n"
	result := strings.Replace(str, "to", "too", 1)
	fmt.Println("Replace: ", result)

	count := strings.Count(str, "1")
	fmt.Println("count:", count)

	result = strings.Repeat(str, 3)
	fmt.Println("repeat:", result)

	result = strings.ToLower(str)
	fmt.Println("lower:", result)

	result = strings.ToUpper(str)
	fmt.Println("upper:", result)

	result = strings.TrimSpace(str)
	fmt.Println("trimSpace:", result)

	result = strings.Trim(str, "\n\r")
	fmt.Println("trim:", result)

	result = strings.TrimLeft(str, "\n\r")
	fmt.Println("trimLeft:", result)

	result = strings.TrimRight(str, "\n\r")
	fmt.Println("trimRight:", result)

	splitResult := strings.Fields(str)
	for i := 0; i < len(splitResult); i++ {
		fmt.Println(splitResult[i])
	}

	splitResult = strings.Split(str, "1")
	for i := 0; i < len(splitResult); i++ {
		fmt.Println(splitResult[i])
	}

	str2 := strings.Join(splitResult, "1")
	fmt.Println("join:", str2)

	str2 = strconv.Itoa(1000)
	fmt.Println("Itoa:", str2)

	number, err := strconv.Atoi("abc")
	if err != nil {
		fmt.Println("不能转换整型", err)
		return
	}

	fmt.Println("number: ", number)


	s := "Yes我爱慕课网!" // UTF-8
	fmt.Println(s)

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:",
		utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}