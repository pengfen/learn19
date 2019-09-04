package main

import "fmt"

func reverse(str string) string {
	var result string
	strlen := len(str)
	for i := 0; i < strlen; i++ {
		result = result + fmt.Sprintf("%c", str[strlen - i - 1])
	}
	return result
}

func reverse1(str string) string {
	var result []byte
	tmp := []byte(str)
	length := len(str)

	for i := 0; i < length; i++ {
		result = append(result, tmp[length - i - 1])
	}
	return string(result)
}

func test_string() {
	var str1 = "welcome"
	str2 := "wel"

	str3 := fmt.Sprintf("%s %s", str1, str2)
	n := len(str3)

	fmt.Println(str3)

	fmt.Printf("len(str3)=%d\n", n) // 11

	substr := str3[0:5]
	fmt.Println(substr) // welco 前五个字符

	substr = str3[6:]
	fmt.Println(substr) // e wel

	result := reverse(str3) // lew emoclew 反转
	fmt.Println(result)

	result = reverse1(result) // welcome wel 反转
	fmt.Println(result)
}

func test_for() {
    str := "welcome to go world"
    n := len(str)

    for i := 0; i < n; i ++ {
    	ch := str[i] // 依据下标取字符串中的字符 类型为byte
    	fmt.Println(i, ch)
    }	
}

func main() {
	// test_string()

	test_for() // 测试字符串遍历
}