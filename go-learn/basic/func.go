package main

import "fmt"

/*
函数是一等公民 参数 变量 返回值都可以是函数   不可变性 不能有状态 只有常量和函数   函数只能有一个参数

func funcName(input type1, input2 type2) (output1 type1, output2 type2) {return value1, value2 // 处理逻辑代码 --- 返回多个值}
关键字func用来声明一个函数funcName
函数可以有一个或者多个参数 每个参数后面带有类型 通过,分隔函数可以返回多个值
返回值声明了两个变量output1和output2 如果你不想声明也可以就保留两个类型声明
如果只有一个返回值且不声明返回值变量 那么你可以省略包括返回值的括号
如果没有返回值 就直接省略最后的返回信息
如果有返回值 那么必须在函数的外层添加return语句

声明语法 func 函数名(参数列表)[(返回值列表)]()
func add() {}   func add(a int, b int) {}   func add(a int, b int) int {}   func add(a int, b int) (int, int) {}   func add(a, b int)(int, int) {}

函数特点
a. 不支持重载 一个包不能有两个名字一样的函数
b. 函数是一等公民 函数也是一种类型 一个函数可以赋值给变量
c. 匿名函数
d. 多返回值

变值 func(arg ...int){}

panic 函数  中断原有的控制流程
recover 函数 把原有中断的流程恢复回来

main 函数
init 函数
*/

func test_func()  {
	x := 3
	y := 4
	z := 5

	max_xy := max(x, y) // 调用函数max(x, y)
	max_xz := max(x, z) // 调用函数max(x, z)

	fmt.Printf("max(%d, %d) = %d", x, y, max_xy) // max(3, 4) = 4
	fmt.Println()
	fmt.Printf("max(%d, %d) = %d", x, z, max_xz) // max(3, 5) = 5
	fmt.Println()
	fmt.Printf("max(%d, %d) = %d", y, z, max(y, z)) // max(4, 5) = 5
}

// 返回a, b中最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func test_func1()  {
	x := 3
	y := 4

	xPLUSy, xTIMESy := SumAndProduct(x, y)
	fmt.Printf("%d + %d = %d", x, y, xPLUSy) // 3 + 4 = 7
	fmt.Printf("%d * %d = %d", x, y, xTIMESy) // 3 * 4 = 12
}

// 简单的一个函数 实现了参数+1的操作
func add1(a int) int {
	a = a + 1 // 改变a值
	return a
}

func test_func2()  {
	x := 3
	fmt.Println("x = ", x) // x = 3

	x1 := add1(x)
	fmt.Println("x + 1 = ", x1) // x + 1 = 4
	fmt.Println("x = ", x) // x = 3
}

func test_func3()  {
	x := 3
	fmt.Println("x = ", x)

	// 传指针使得多个函数能操作同一个对象
	// 传指针比较轻量级(8bytes) 只是传内存地址
	x1 := add2(&x)
	fmt.Println("x + 1 = ", x1)
	fmt.Println("x = ", x)
}

// 实现了参数+1的操作
func add2(a *int) int {
	*a = *a + 1 // 修改a的值
	return *a
}

/*
defer 用途
1. 当函数返回时 执行defer语句 因此 可以用来做资源清理
2. 多个defer语句 按先进后出的方式执行
3. defer语句中的变量 在defer声明时就决定
*/
func test_defer()  {
	for i := 0; i < 5; i ++ {
		defer fmt.Printf("%d", i) // 43210
	}
}

func test_func4()  {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice) // [1 2 3 4 5 7]
	odd := filter(slice, isOdd)
	fmt.Println("Odd elements of slice are: ", odd) // [1 3 5 7]
	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are: ", even) // [2 4]
}

// 声明函数类型当做一个参数
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

type testInt func(int) bool // 声明函数类型

func isOdd(integer int) bool {
	if integer % 2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer % 2 == 0 {
		return true
	}
	return false
}

func test_arg(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func test_printf(args ...interface{})  {
	for _, arg := range args {
		switch arg.(type) {
			case int:
				fmt.Println(arg, "是整型")
			case string:
				fmt.Println(arg, "是字符串")
			case int64:
				fmt.Println(arg, "是int64类型")
			default:
				fmt.Println(arg, "是未知类型")
		}
	}
}

func test_closure()  {
	var j int = 5

	a := func()(func()) {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d", i, j) // i, j: 10, 5   i, j: 10, 10
		}
	}()
	a()

	j *= 2
	a()
}

// 三角形图形
func Print(n int){
	for i := 0; i < n + 1; i ++ {
		for j := 0; j < i; j ++ {
			fmt.Print("A")
		}
		fmt.Println()
	}
}

func test_panic() {
	defer func() {
		fmt.Println("After panic!") // panic之后defer里面的依然可以得到执行
	}()

	panic("I am panicing!")

	fmt.Println("After panic!")
}

func test_anony()  {
	type sum func(x, y int) int
	var f sum = func(x, y int) int {
		return x + y
	}

	fmt.Println(3, 4)
}

func main()  {
	// test_func()

	// test_func1() // 多个返回值

	// test_func2() // 传值

	// test_func3() // 传指针

	// test_defer() // 延迟

	// test_func4() // 函数作为值 类型

	// test_arg(2, 3, 4) // 测试不定参数
	// test_arg(1, 3, 5, 13) // 测试不定参数

	//var v1 int = 1
	//var v2 int64 = 234
	//var v3 string = "welcome"
	//var v4 float32 = 1.234
	//test_printf(v1, v2, v3, v4)

	// test_closure() // 测试闭包

	// Print(6) // 三角形图形

	test_panic()

	test_anony() // 匿名函数
}
