package main

import "errors"
import "fmt"
import "reflect"
import "runtime"

type op_func func(int, int) int

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 { // 只支持两个非负数字的加法
		err = errors.New("加数不能为负数")
		return
	}

	return a + b, nil // 支持多重返回值
}

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func MyPrintf(args ...interface{}) {
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

func test_closure() {
    var j int = 5

    a := func()(func()) {
    	var i int = 10
    	return func() {
    		fmt.Printf("i, j: %d, %d\n", i, j)
    	}
    }()

    a() // 10 5

    j *= 2

    a() // 10 10
}

// 三角形图形
func Print(n int) {
	for i := 0; i < n + 1; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("A")
		}
		fmt.Println()
	}
}

/*
func add(a, b int) int {
	return a + b
}*/
// 命名返回值的名字 返回c
func add(a, b int) (c int) {
	c = a + b
	return
}

func test_for() {
	str := "welcome to go world"

	for index, val := range str {
		fmt.Printf("index[%d] val[%c] len[%d]\n", index, val, len([]byte(string(val))))
	}

	// 定义局部变量
	var a int = 100
	var b int = 200
	var ret int

	// 调用函数并返回最大值 
	ret = max(a, b)
	fmt.Printf("最大值是: %d\n", ret)
}

func max(num1, num2 int) int {
	// if num1 > num2 {
	// 	return num1
	// }
	// return num2
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

func test_change(x, y string) (string, string) {
    return y, x	
}

func test_oper() {
	var a, b int
	add(a, b)

	var e op_func
	e = add
	fmt.Println(add)
	fmt.Println(e)

	sum := operator(e, 100, 200) // 300
	fmt.Println(sum)
}

func sub(a, b int) int {
	return a - b
}

func operator(op op_func, a, b int) int {
	return op(a, b)
}

// 多参数
func arg(a int, arg ...int) int {
	var sum int = a
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}

	return sum
}

func concat(a string, arg ...string) (result string) {
	result = a
	for i := 0; i < len(arg); i++ {
		result += arg[i]
	}

	return
}

var (
	result = func(a1 int, b1 int) int {
		return a1 + b1
	}
)

func test(a, b int) int {
	result := func(a1 int, b1 int) int {
		return a1 + b1
	}

	return result(a, b)
}

func modify(a int) {
	a = 100
}

func add1(a int) int { // 实现参数+1的操作
	a = a + 1
	return a
}

func test_value() {
	// 传值只修改copy
	x := 3

	fmt.Println("x = ", x) // 3

	x1 := add1(x)

	fmt.Println("x+1 = ", x1) // 4
	fmt.Println("x = ", x) // 3
}

func add2(a *int) int {
	*a = *a + 1
	return *a
}

func test_pointer() {
    x := 3

    fmt.Println("x = ", x) // 3

    x1 := add2(&x)
    
    fmt.Println("x+1 = ", x1) // 4
    fmt.Println("x = ", x) // 4
}

func eval(a, b int, op string) int {
	switch op {
		case "+":
			return a + b
		case "-":
		    return a - b
		case "*":
		    return a * b
		case "/":
			return a / b
		default:
			panic("未知操作" + op)
	}
}

// 多返回值
func eval_much(a, b int, op string) (int, error) {
	switch op {
		case "+":
			return a + b, nil
		case "-":
			return a - b, nil
		case "*":
			return a * b, nil
		case "/":
			q, _ := div(a, b)
			return q, nil
		default:
			return 0, fmt.Errorf("未知操作: %s", op)
	}
}

func div(a, b int) (q, r int) {
	//return a / b, a % b
	q = a / b
	r = a % b
	return
	// 函数返回多个值时可以起名字
	// 仅用于非常简单的函数
	// 对于调用者而言没有区别
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("调用函数 %s 参数" + "(%d, %d)\n", opName, a, b)

	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

func test_div() {
	// fmt.Println(eval(3, 4, "*"))
	// q, r := div(13, 3)
	// fmt.Println(q, r)

	fmt.Println("错误处理")
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println(result)
	}

	q, r := div(13, 3)
	fmt.Printf("13 / 3 = %d mod %d\n", q, r)

	fmt.Println("pow(3, 4) =", apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println("1 + 2 + ... + 5 = ", sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println("a, b交换后:", a, b)
}

func main() {
	a := 8
	fmt.Println(a) // 8
	modify(a)
	fmt.Println(a) // 8

	// 测试不定参数
	// myfunc(2, 3, 4)
	// myfunc(1, 3, 5, 13)

	// 
	// var v1 int = 1
	// var v2 int64 = 234
	// var v3 string = "welcome"
	// var v4 float32 = 1.234
	//MyPrintf(v1, v2, v3, v4)

	//test_closure() // 测试闭包

	//Print(6) // 三角形图形

	m := add
	fmt.Println(m) // 地址0x
	sum := m(200, 300)
	fmt.Println(sum) // 500

	// test_for()

	c, d := test_change("ricky", "caopeng") // 交换两个值
	fmt.Println(c, d)

	//test_oper() // 测试加减操作

	//sum := arg(10, 2, 3, 4, 5)
	//fmt.Println(sum) // 24

	res := concat("welcome", "to", "world")
	fmt.Println(res)

	fmt.Println(result(100, 200)) // 300
	fmt.Println(test(100, 200)) // 300

	var i int = 0
	// defer 释放资源 - 注意是在主线程执行完后再执行defer 
	//注意后进先出 先打印second 再打印0
	defer fmt.Println(i) // 0
	defer fmt.Println("second") // second

	i = 10
	fmt.Println(i) // 10

	test_value() // 测试传值

	test_pointer() // 测试传指针
}
