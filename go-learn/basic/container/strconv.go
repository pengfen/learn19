package main

import "fmt"
import "strconv"

func main() {
	var str string
	fmt.Scanf("%s", &str)

	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换失败", err)
		return
	}

	fmt.Println(number)



	var str string
	fmt.Scanf("%s", &str)

	var result = 0
	for i := 0; i < len(str); i ++ {
		num := int(str[i] - '0')
		result += (num * num *num)
	}

	number, err != strconv.Atoi(str)
	if err != nil {
		fmt.Printf("%s不能转换整型\n", str)
		return
	}

	if result == number {
		fmt.Printf("%d 是水仙花数\n", number)
	} else {
		fmt.Printf("%d 不是水仙花数\n", number)
	}
}