package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
内建变量类型
bool string  (u)int (u)int8 (u)int16 (u)int32 uintptr   byte rune    float32 float64 complex64 complex128

boolean 布尔类型 true false 默认为false

整型 rune int8 int16 int32 int64     byte uint8 uint16 uint32 uint64
rune 是 int32 的别称
byte 是 uint8 的别称

浮点 float32 fload64
complex128(64位实数 + 64位虚数)  complex64(32位实数 + 32位虚数)

值类型和引用类型
1. 值类型 变量直接存储值 内存通常在栈中分配
2. 引用类型 变量存储的是一个地址 这个地址存储最终的值 内存通常在堆上分配 通过GC回收

bool 类型 只能存true和false
var a bool
var a bool = true
var a = true

相关操作符 ! && ||
var a bool = true
var b

v 默认格式 以符合 Go 语法的方式输出 不同类型的默认格式如下
t 布尔型
d 整型
f 浮点型
g 复数型
s 字符串
p 通道
p 指针
x 无符号整型

错误类型
error类型 专门用来处理错误信息
err := errors.New("emit macho dwarf: elf header corrupted")
if err != nil {
    fmt.Print(err)
}
*/
func init() { // 随机种子
	rand.Seed(time.Now().UnixNano())
}
func test_type() {
	var a = 100
	// var b chan int = make(chan int, 1)
	b := make(chan int, 1)
	fmt.Println("a=", a) // a= 100
	fmt.Println("b=", b) // b= 0xc42001a0e0

	handle_type(a)
	fmt.Println("a=", a) // a= 100
	handle_type2(&a)
	fmt.Println("a=", a) // a= 10
}

func handle_type(a int) {
	a = 10
	return
}

func handle_type2(a *int)  {
	*a = 10
}

func test_type1()  {
	var x int
	x = 2
	y := "welcome to go world"
	fmt.Printf("%d %s", x, y) // 2 welcome to go world
}

func test_type2() {
	//var x [10]int // 直接声明
	//x[0] = 1
	//x[1] = 2

	x := [10]int{1, 2, 3, 4, 5}
	fmt.Printf("%v", x) // 注意%v的运用 %v可以打印数组 map 字符串等各类值  [1 2 3 4 5 0 0 0 0 0]
}

func test_type3()  {
	var x complex64 = 2 + 4i
	fmt.Print(x * x) // (-12+16i)
}

/**
写一个程序 交换两个整数的值 a=3;b=4 交换之后 a=4;b=3
*/
func test_change() {
	var a int = 3
	var b int = 4

	fmt.Printf("a=%d;b=%d", a, b) // a=3;b=4
	handle_change(a, b) // 交换二个值

	first := 100
	second := 200
	handle_change1(&first, &second)
	fmt.Println("first=", first) // 200
	fmt.Println("second=", second) // 100
}

func handle_change1(e *int, f *int) {
	tmp := *e
	*e = *f
	*f = tmp
	return
}

func handle_change(a int, b int) {
	var c int = a
	a = b
	b = c
	fmt.Printf("a=%d;b=%d", a, b) // a=4;b=3
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
			case bool:
				fmt.Printf("参数%t是布尔型", i)
			case float64:
				fmt.Printf("参数%f是浮点型", i)
			case int, int64:
				fmt.Printf("参数%d是整型", i)
			case nil:
				fmt.Printf("参数%d是空类型", i)
			case string:
				fmt.Printf("参数%s是字符型", i)
			default:
				fmt.Printf("参数%d未知类型", i)
		}
	}
}

func test_bool()  {
	var a bool = true
	var b bool = false

	fmt.Println(!a) // false
	fmt.Println(a && b) // false
	fmt.Println(a || b) // true
}

func test_int() {
	var n int16 = 34
	var m int32

	m = int32(n)
	fmt.Printf("m=%d n=%d", m, n) // m=34 n=34
}

func test_rand() {
	for i := 0; i < 10; i ++ { // 随机生成十个整数
		a := rand.Int()
		fmt.Println(a)
	}

	for i := 0; i < 10; i ++ { // 随机生成十个小于一百的整数
		a := rand.Intn(100)
		fmt.Println(a)
	}

	for i := 0; i < 10; i ++ { // 随机生成十个浮点数
		a := rand.Float32()
		fmt.Println(a)
	}
}

func test_guess() {
	var n int
	n = rand.Intn(100) // 注意 要以当前时间为随机数种子

	for {
		var input int
		fmt.Scanf("%d", &input)
		fmt.Println()
		flag := false

		switch {
			case input == n:
				fmt.Println("你猜对了")
				flag = true
			case input > n:
				fmt.Println("你猜大了")
			case input < n:
				fmt.Println("你猜小了")
		}
		if flag {
			break
		}
	}
}

func test_byte() {
	var str = "welcome to go world"
	var str1 = `
    你好,
    好不好
    `

	var b byte = 'c'
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(b) // 99
	fmt.Printf("%c", b) // c
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

func test_type4()  {
	x := "ricky"

	for _, v := range x {
		fmt.Printf("%c", v)
		fmt.Println()
	}
}

func main()  {
	// test_type() // 测试值
	// test_change() // 测试交换二个值

	// test_type1()
	// test_type2()
	// test_type3()
	test_type4()

	// test_bool() // 测试布尔型
	// test_int() // 测试整型
	// test_rand() // 测试随机数
	// test_guess() // 猜数字
	// test_byte()
	// test_format()
}
