package main

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

func test_string()  {
	var str1 = "welcome"
	str2 := "wel"

	str3 := fmt.Sprintf("%s %s", str1, str2)
	n := len(str3)
	fmt.Println(str3) // welcome wel
	fmt.Printf("len(str3)=%d", n) // 11
	fmt.Println()

	substr := str3[0:5]
	fmt.Println(substr) // 前五个字符 welco

	substr = str3[6:]
	fmt.Println(substr) // e wel

	result := reverse(str3)
	fmt.Println(result) // 反转 lew emoclew

	result = reverse1(result)
	fmt.Println(result) // 反转 welcome wel
}

func reverse(str string) string  {
	var result string
	strlen := len(str)
	for i := 0; i < strlen; i ++ {
		result = result + fmt.Sprintf("%c", str[strlen - i - 1])
	}
	return result
}

func reverse1(str string) string  {
	var result []byte
	tmp := []byte(str)
	length := len(str)

	for i := 0; i < length; i ++ {
		result = append(result, tmp[length - i -1])
	}
	return string(result)
}

func test_string1()  {
	s := "Yes我爱慕网!" // UTF-8
	fmt.Println(s)

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
	    fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeLastRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}

func test_string2()  {
	s := "welcome to go world"
	s1 := s[0:5]
	s2 := s[0:]

	fmt.Println(s1)
	fmt.Println(s2)
}

func test_string3()  {
	s := "welcome to go world"
	s1 := []rune(s)

	s1[0] = 200
	s1[1] = 123
	s1[2] = 64

	str := string(s1)
	fmt.Println(str)
}

func test_copy()  {
	var a []int = []int{1, 2, 3, 4, 5, 6}
	b := make([]int, 1)

	copy(b, a)
	fmt.Println(b)
}

func test_float()  {
	var a = [...]float64{2.3, 0.8, 28.2, 0.6}
	sort.Float64s(a[:])

	fmt.Println(a)
}

func test_strings()  {
	var a = [...]string{"abc", "efg", "b", "A", "eeee"}
	sort.Strings(a[:])

	fmt.Println(a)
}

func test_ints()  {
	var a = [...]int{1, 8, 38, 2, 348, 484}
	sort.Ints(a[:])

	fmt.Println(a)
}

func test_search()  {
	var a = [...]int{1, 8, 38, 2, 348, 484}
	sort.Ints(a[:])
	index := sort.SearchInts(a[:], 348)
	fmt.Println(index)
}

func main()  {
	test_string() // 测试字符串遍历

	test_string1()

	test_string2()

	test_string3()

	test_copy()

	// 排序
	test_float()

	test_strings()

	test_ints()

	test_search()
}