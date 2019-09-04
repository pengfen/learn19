package main

import (
	"bufio"
	"fmt"
	"os"
)

// 输入一行字符 分别统计出其中英文字母 空格 数字和其它字符的个数

func count(str string) (worldCount, spaceCount, numberCount, otherCount int)  {
	t := []rune(str)
	for _, v := range t {
		switch {
		case v >= 'a' && v <= 'z':
			worldCount ++
		case v >= 'A' && v <= 'Z':
			worldCount ++
		case v == ' ':
			spaceCount ++
		case v >= '0' && v <= '9':
			numberCount++
		default:
			otherCount++

		}
	}
	return
}

func main()  {
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()

	if err != nil {
		fmt.Println("read from console err : " , err)
		return
	}
	wc, sc, nc, oc := count(string(result))
	fmt.Printf("World count:%d space count: %d number count: %d others count: %d", wc, sc, nc, oc)
}
