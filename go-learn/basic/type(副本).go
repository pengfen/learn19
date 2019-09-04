package main

import "fmt"
import "math/rand"
import "time"

func test_bool() {
	var a bool = true
	var b bool = false

	fmt.Println(!a)
	fmt.Println(a && b)
	fmt.Println(a || b)
}

func test_int() {
	var n int16 = 34
	var m int32

	m = int32(n)

	fmt.Printf("m=%d n=%d\n", m, n)
}


func init() { // 随机种子
	rand.Seed(time.Now().UnixNano())
}	

func test_rand() {
	for i := 0; i < 10; i ++ { // 十个整数
		a := rand.Int()
		fmt.Println(a)
	}

	for i := 0; i < 10; i ++ { // 十个小于一百的整数
		a := rand.Intn(100)
		fmt.Println(a)
	}

	for i := 0; i < 10; i++ { // 十个浮点数
		a := rand.Float32()
		fmt.Println(a)
	}
}

func test_guess() {
	var n int
	n = rand.Intn(100)

	for {
		var input int
		fmt.Scanf("%d\n", &input)
		flag := false

		switch {
			case input == n:
				fmt.Println("你猎对了")
				flag = true
			case input > n:
				fmt.Println("大了")
			case input < n:
				fmt.Println("小了")
		}

		if flag {
			break
		}
	}
}

func test_byte() {
	var str = "welcome to go world\n"
	var str1 = `
	你好,
	好不好
	`

	var b byte = 'c'
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(b)
	fmt.Printf("%c\n", b)
}

func test_format() {
	var a int = 100
	var b bool
	c := 'a'

	fmt.Printf("%+v\n", a) // 100
	fmt.Printf("%#v\n", b) // false
	fmt.Printf("%T\n", c) // int32
	fmt.Printf("90%%\n") // 90%
	fmt.Printf("%t\n", b) // false
	fmt.Printf("%b\n", 100) // 1100100
	fmt.Printf("%f\n", 199.22) // 199.220000
	fmt.Printf("%q\n", "welcome to go world")
	fmt.Printf("%x\n", 398) // 18e
	fmt.Printf("%p\n", &a) // 地址 )x

	str := fmt.Sprintf("a=%d", a) // a=100
	fmt.Printf("%q\n", str)
}

func main() {
	//test_bool() // 测试布尔型

	//test_int()

	//test_rand()

	//test_guess() // 猜数字

	//test_byte()

	test_format()
}