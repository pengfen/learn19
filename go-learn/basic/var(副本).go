/*
打印当前系统常量
*/
package main
import (
"fmt"
"os"
"time"
)

func test_path() { // 测试当前环境变量
	var goos string = os.Getenv("GOOS")
	fmt.Printf("The operating system is: %s\n", goos)

	var path string = os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}

/*
定义两个常量Man=1和Female=2 获取当前时间秒数 如果参被Female整除 则在终端打印female否则打印man
*/
/* 常量常写在方法外 以便所有方法都可使用
const (
	Man = 1
	Female = 2
)
*/

func test_const() {
	const Man int64 = 1
	const Female int64 = 2

	Second := time.Now().Unix() // 获取当前时间戳
	if (Second % Man == 0) {
		fmt.Print("female")
	} else {
		fmt.Print("male")
	}

	fmt.Println(1000 * time.Millisecond) // 1s
	time.Sleep(1000 * time.Millisecond)
}

func test_type() {
	var a = 100
	//var b chan int = make(chan int, 1)
	b := make(chan int, 1)
	fmt.Println("a=", a)
	fmt.Println("b=", b) // b= 0xc4200ac000 地址

	handle_type(a)
	fmt.Println("a=", a)
	handle_type2(&a)
	fmt.Println("a=", a)
}

func handle_type(a int) {
	a = 10
	return
}

func handle_type2(a *int) {
	*a = 10
}

/*
写一个程序 交换两个整数的值
a=3;b=4 交换之后 a=4;b=3
*/
func test_change() {
	var a int = 3
	var b int = 4

	fmt.Printf("a=%d;b=%d\n", a, b);

	handle_change(a, b)

	// first := 100
	// second := 200
	// handle_change1(&first, &second)
	// fmt.Println("first=", first)
	// fmt.Println("second=", second)
}

func handle_change(a int, b int) {
	var c int = a
	a = b
	b = c
	fmt.Printf("a=%d;b=%d\n", a, b);
}

// func handle_change1(e *int,f *int) {
// 	tmp := *e
// 	*e = f
// 	*f = tmp
// 	return
// }


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
	fmt.Println(p)
	f2()
}

func f2() {
	fmt.Println(p)
}

const (
	c0 = iota // 0
	c1 = iota // 1
	c2 = iota // 2
)

const (
	d1 = 1 << iota // 1 << 0
	d2 = 1 << iota // 1 << 1
	d3 = 1 << iota // 1 << 2
)

const (
	e1 = iota * 42
	e2 = iota * 42
	e3 = iota * 42 
)

// 可简写
const (
	h1 = iota
	h2
	h3
)

func main() {
	//test_path() // 测试当前环境变量

	//test_const() // 测试常量

	//test_type() // 测试值

	//test_change() // 测试交换二个值

	// n() // G 全局变量
	// m() // o 局部变量
	// n() // G 全局变量

	// p = "G"
	// fmt.Print(p)
	// f1()

	fmt.Println(c0, c1, c2) // 0 1 2

	fmt.Println(d1, d2, d3) // 1 2 4

	fmt.Println(e1, e2, e3) // 0 42 84

	fmt.Println(h1, h2, h3) // 0 1 2
}