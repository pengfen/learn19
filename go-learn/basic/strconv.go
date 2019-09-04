package main

import (
	"fmt"
	"strconv"
)

/*
17. strconv.Itoa(i int) 把一个整数i转成字符串
18. strconv.Atoi(str string) (int, error) 把一个字符串转成整数
练习6 写一个函数分别演示Itoa, Atoi的用法
*/
func test_prog5()  {
	str := strconv.Itoa(1000)
	fmt.Println("Itoa: ", str)

	number, err := strconv.Atoi("abc")
	if err != nil {
		fmt.Println("不能转换整型", err)
		return
	}
	fmt.Println("number: ", number)
}

func test_str()  {
	var str string
	fmt.Scanf("%s", &str)

	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换失败", err)
		return
	}
	fmt.Println(number)
}

func test_str1()  {
	var str string
	fmt.Scanf("%s", &str)

	var result = 0
	for i := 0; i < len(str); i ++ {
		num := int(str[i] - '0')
		result += (num * num * num)
	}

	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("%s不能转换整型", str)
		fmt.Println()
		return
	}

	if result == number {
		fmt.Printf("%d 是水仙花数", number)
	} else {
		fmt.Printf("%d 不是水仙花数", number)
	}
}

func main()  {
	test_prog5()

	test_str()

	test_str1()
}