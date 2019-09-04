package main

import (
	"fmt"
	"strconv"
)

func isSXH(n int) bool {
	// 打印出100~999中所有的水仙花数
	// 水仙花数 是指一个三位数 其各位数字立方和等于该数本身
	var i, j, k int
	i = n % 10 // 获取个位数
	j = (n / 10) % 10
	k = (n / 100) % 10
	sum := i * i * i + j * j * j + k * k * k
	return sum == n
}

func test_sxh() {
	//var n int
	//var m int
	//fmt.Scanf("%d, %d", &n, &m)

	for i := 100; i < 200; i ++ {
		if isSXH(i) == true {
			fmt.Println(i, "是水仙花数")
		}
	}
}

func text_sxh1() {
	var str string
	fmt.Scanf("%s", &str)

	var result = 0
	for i := 0; i < len(str); i ++ {
		num := int(str[i] - '0')
		result += (num * num * num)
	}

	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("此字符%s不能转换整型", str)
		return
	}

	if result == number {
		fmt.Printf("%d是水仙花数", number)
	} else {
		fmt.Printf("%d不是水仙花数", number)
	}
}

func main()  {
	// test_sxh()
	text_sxh1()
}
