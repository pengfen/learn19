package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func test_basic()  {
	m := add
	fmt.Println(m) // 0x4821f0
	sum := m(200, 300)
	fmt.Println(sum) // 500
}

/*
func add(a, b int) int {
    return a + b
}
*/
// 命名返回值的名字 返回c
func add(a, b int) (c int)  {
	c = a + b
	return
}

func test_change1(x, y string) (string, string) {
	return y, x
}

func add3(a, b int) (c int) {
	c = a + b
	return
}

type op_func func(int, int) int

func operator(op op_func, a, b int) int  {
	return op(a, b)
}

func test_oper()  {
	var a, b int
	add3(a, b)

	var e op_func
	e = add3
	fmt.Println(add3) // 0x4820f0
	fmt.Println(e) // 0x4820f0

	sum := operator(e, 100, 200) // 300
	fmt.Println(sum)
}

func concat(a string, arg ...string) (result string) {
	result = a
	for i := 0; i < len(arg); i ++ {
		result += arg[i]
	}
	return
}

func test_eval(a, b int, op string) int {
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

func eval_much(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a / b, nil
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("未知操作: %s", op)
	}
}

func div(a, b int) (q, r int)  {
	q = a / b
	r = a % b
	return
}

func test_div()  {
	//fmt.Println(test_eval(3, 4, "*"))
	//q, r := div(13, 3)
	//fmt.Println(q, r)

	fmt.Println("错误处理")
	if result, err := eval_much(3, 4, "x"); err != nil {
		fmt.Println("错误: ", err)
	} else {
		fmt.Println(result)
	}

	q, r := div(13, 3)
	fmt.Printf("13 / 3 = %d mod %d", q, r)

	fmt.Println("pow(3, 4) = ", apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println("1 + 2 + ... + 5 = ", sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println("a, b交换后: ", a, b)
}

func apply(op func(int, int) int , a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println("调用函数 %s 参数" + "(%d, %d)", opName, a, b)
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

func main()  {
	// test_basic()

	// c, d := test_change1("ricky", "caopeng") // 交换两个值
	// fmt.Println(c, d) // caopeng ricky

	//test_oper() // 测试加减操作
	//
	//res := concat("welcome", "to", "world") // welcometoworld
	//fmt.Println(res)

	//test_eval(3, 5, "+")
	//eval_much(3, 5, "-")

	test_div()
}
