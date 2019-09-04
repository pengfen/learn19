package main

import "fmt"

// rune 与 byte 相似
// byte 等同于int8 常用来处理ascii字符
// rune 等同于int32 常用来处理unicode或utf-8字符
func process1(str string) bool  {
	t := []rune(str)
	length := len(t)

	for i, _ := range t {
		if i == length / 2 {
			break
		}

		last := length - i - 1
		if t[i] != t[last] {
			return false
		}
	}
	return true
}

func main()  {
	var str string // 信言不美 美言不信
	fmt.Scanf("%s", &str)

	if process1(str) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
