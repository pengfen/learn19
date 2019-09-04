package main

import (
	"fmt"
	"os"
	"time"
)

/**
大写字母开头的变量是可导出的 是公用变量(public) 小写字母开头的不可导出 是私有变量(private)
大写字母开头的函数是可导出的 是公用函数(public) 小写字母开头的不可导出 是私有函数(private)

var 关键字定义变量 注意变量类型放在变量名后
编译器可推测变量类型 没有char 只有rune 原生支持复数类型

var variableName type // 定义一个名称为variableName 类型为type的变量

语法  var identifier type
var a int
var b string
var c bool
var d int = 8
var e string = "welcome"
var v1 [10]int // 数组
var v2 []int // 数组切片
var v3 *int // 指针
var v4 map[string]int // map, key为string类型 value为int类型

// 定义多个变量
var (
    a int // 默认为0
    b string // 默认为""
    c bool // 默认false
    d int = 8
    e string = "welcome"
)

// 变量初始化
var v1 int = 10
var v2 = 10 // 编译器自动推导数据类型
v3 := 10 // 编译器自动推导数据类型

// 定义三个变量 并分别初始化相应的值
var vname1, vname2, vname3 = v1, v2, v3

注意 (:=左侧的变量不能被声明过) var i int
i := 2 // 会报错 no new variables on left side of :=

匿名变量
func GetName() (firstName, lastName, nickName string) {
    return "cao", "peng", "ricky"
}

_, _, nickName := GetName() // 只获取nickName

_(下划线) 是特殊的变量名 任何赋予它的值都会被丢弃

变量的作用域
在函数内部声明的变量叫做局部变量 生命周期仅限于函数内部
在函数外部声明的变量叫做全部变量 生命周期作用于整个包 如果是大写的 则作用于整个程序


常量
1. 常量使用const修饰 代表永远是只读的 不能修改
2. const 只能修饰boolean number (int 相关类型 浮点类型complex)和string
3. 语法 const identifier [type] = value 其中type可以省略

const b string = "welcome to "
const b = "welcome"
const Pi = 3.1414926
const a = 9/3
const c = getValue() // 非法

4. 比较优雅的写法
const(a=0 b=1 c=2)

5. 更加专业的写法
const(
a=iota
b   // 1
c   //2
)

预定义常量 true false iota
iota 可被认为可被编译器修改的常量 在第一个const关键字出现时重置为0
*/

func test_path() { // 测试当前环境变量
	var goos string = os.Getenv("GOOS")
	fmt.Printf("The operating system is: %s", goos)
	fmt.Println()

	var path string = os.Getenv("PATH")
	fmt.Printf("Path is %s", path)
}

/**
定义两个常量 Man=1 和 Female=2 获取当前时间秒数 如果参数Female整除 则在终端打印female 否则打印man

常量 常写在方法外 以便所有方法都可使用
const (
    Man = 1
    Female = 2
)
*/
func test_const() {
	const Man int64 = 1
	const Female int64 = 2

	Second := time.Now().Unix() // 获取当前时间戳
	if (Second % Female == 0) {
		fmt.Print(Female)
	} else {
		fmt.Print(Man)
	}

	time.Sleep(1000 * time.Millisecond) // 睡眼一秒
}

var g = "G"
func n() {
	fmt.Println(g)
}

func m() {
	g := "o"
	fmt.Println(g)
}

var p string
func f1() {
	p := "o"
	fmt.Println(p) // o
	f2() // G
}
func f2()  {
	fmt.Println(p)
}

const (
	c0 = iota // 1
	c1 = iota // 2
	c2 = iota // 4
)
const (
	d1 = 1 << iota // 1 << 0
	d2 = 1 << iota // 1 << 1
	d3 = 1 << iota // 1 << 2
)
const (
	e1 = iota * 42 // 0 * 42
	e2 = iota * 42 // 1 * 42
	e3 = iota * 42 // 2 * 42
)
const (
	h1 = iota // 0
	h2        // h2 = iota ---> 1
	h3        // h3 = iota ---> 2
)

func main() {
	// test_path() // 测试当前环境变量

	// test_const() // 测试常量

	// 测试全局变量
	// n() // G 全局变量
	// m() // o 局部变量

	p = "G"
	fmt.Println(p) // G
	f1()

	// 测试常量 (iota)
	fmt.Println(c0, c1, c2) // 0 1 2
	fmt.Println(d1, d2, d3) // 1 2 4
	fmt.Println(e1, e2, e3) // 0 42 84
	fmt.Println(h1, h2, h3) // 0 1 2
}