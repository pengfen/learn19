package main

import "fmt" // import 用于导入外部代码 fmt包用于格式化输出数据
import "math"
import "math/cmplx" // 提供复数处理的许多函数

/*
idea 支持 go (直接使用goLand)

1. 安装go插件
settings ---> plugins ---> Install JetBrains Plugin... ---> go

2. 创建新项目
New ---> Project ---> go ---> 选择 go SDK(/home/ricky/app/go) ---> chat (默认目录 ~/IdeaProjects/chat)

3. 设置字体大小
settings ---> Color Scheme Font ---> 勾选use clolor ... Size:18 ---> Apply

4. 测试项目
chat目录下创建src目录 ---> src目录下创建wel.go ---> 编写wel.go ---> 运行wel.go

https://play.golang.org/ 浏览器编辑并运行
*/

var (
	aa = 3
	ss = "kkk"
	bb = true
)

/*
需求 编写代码 对于给定一个数字n 求出所有两两相加等于n的组合

如 对于a=2 所有组合如下
0 + 2 = 2
1 + 1 = 2
2 + 0 = 2
*/
func test_conb(n int) {
	for a := 0; a <= n; a++ {
        // type int int
		var b int = n - a // b declared and not used 注意var定义的值得使用，不然报错
        // func Printf(format string, a ...interface{}) (n int, err error) ---> Fprintf(os.stdout, format, a...)
		fmt.Printf("%d + %d = %d\n", a, b, n)
	}
}

func variableInitialValue() {
    // 初始化默认值
    var a int
    var s string
    fmt.Printf("%d %f\n", a, s) // 0 ""

    var a, b int = 3, 4
    var s string = "abc"
    fmt.Println(a, b, s) // 3 4 abc
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s) // 3 4 true def
}

func variableShorter() {
    a, b, c, s := 3, 4, true, "def"
    b = 5 // 更改b的值 
    fmt.Println(a, b, c, s) // 3 5 true def
}

func euler() {
    c := 3 + 4i
    // func Abs(x complex128) float64 { return math.Hypot(real(x), imag(x)) }
    fmt.Println(cmplx.Abs(c)) // 5

    // python 实现方式
    // import cmath
    // cmath.exp(1j * cmath.pi) + 1
	fmt.Printf("%.3f\n", cmplx.Exp(1i * math.Pi) + 1) // 0.000 * 0.000i
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b)) // 勾股定理
}

func calcTriangle(a, b int) int {
	var c int
    // 强制类型转换 注意int float64
	c = int(math.Sqrt(float64(a * a + b * b))) // 5
	return c
}

// 常量
func consts() {
    // 常量的定义 filename = "abc.txt"
    const (
    	filename = "abc.txt"
    	a, b = 3, 4
    )
    var c int
    // const数值可作为各种类型使用/
    // const a, b = 3, 4
    // var c int = int(math.Sqrt(a * a + b * b)) ---> 注意 const可作为各种类型 所以不用强制转换
    c = int(math.Sqrt(a * a + b * b))
    fmt.Println(filename, c) // 5
}

// 枚举类型
// 普通枚举类型 自增枚举类型iota
func enums() {
    const (
    	cpp = iota
    	_
    	python
    	golang
    	javascript
    )

    const (
    	b = 1 << (10 * iota) // 1
    	kb // 1024
    	mb // 1048576 
    	gb // 1073741824
    	tb // 1099511627776
    	pb // 1125899906842624
    )

    fmt.Println(cpp, javascript, python, golang) // 0 4 2 3
    fmt.Println(b, kb, mb, gb, tb, pb) // 
}

func main() {
    // func Println(a ...interface{}) (n int, err error) ---> Fprintln(os.Stdout, a...)
	fmt.Println("welcome to go worold");

    // 测试两两相加为n
	//test_conb(5)

    // 使用var关键字
    // var a, b, c bool
    // var s1, s2 string = "wel", "welcome"
    // 可放在函数内 或直接放在包内
    // 使用var()集中定义变量
    // 让编译器自动决定类型
    // var a, b, i, s1, s2 = true, false, 3, "wel", "welcome"
    // 使用:=定义变量
    // a, b, i, s1, s2 := true, false, 3, "wel", "welcome"
    // 只能在函数内使用
	//variableInitialValue()

	//variableTypeDeduction()

	//variableShorter()
	//fmt.Println(aa, ss, bb)

	//euler()

	//triangle()

	consts()

	enums()
}
