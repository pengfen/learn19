package main

import "fmt"

func test_if()  {
	x := 10
	if x > 10 {
		fmt.Println("x 大于10")
	} else {
		fmt.Println("x 小于等于10") // x 小于等于10
	}

	// if 可在条件判断语句里面声明一个变量
	//if x := computedValue(); x > 10 {
	//	fmt.Println("x 大于10")
	//} else {
	//	fmt.Println("x 小于等于10")
	//}

	// 多条件
	integer := 3
	if integer == 3 {
		fmt.Println("等于3")
	} else if integer < 3 {
		fmt.Println("小于3")
	} else {
		fmt.Println("大于3")
	}
}

func test_goto()  {
	i := 0
	Here :
		if i < 5 {
			println(i)
			i ++
			goto Here
		}
}

func test_for()  {
	sum := 0
	for index := 0; index < 10; index ++ {
		sum += index
	}
	fmt.Println("计算总和: ", sum) // 45

	for ; sum < 100; {
		sum += sum
	}
	fmt.Println("计算总和: ", sum) // 180

	for sum < 105 {
		sum += sum
	}
	fmt.Println("计算总和: ", sum) // 180

	for x := 10; x > 0; x -- {
		if x == 5 {
			break
		}
		fmt.Print(x) // 109876
	}
	fmt.Println()

	for y := 10; y > 0; y -- {
		if y == 5 {
			continue
		}
		fmt.Print(y) // 1098764321
	}
	fmt.Println()

	// for 配合range可以用于读取slice和map的数据
	rating := map[string]float32 {"C":5, "Go":4.5, "Python":4.5, "C++":2}
	for k, v := range rating {
		fmt.Println("key: ", k)
		fmt.Println("val: ", v)
	}

	str := "welcome to go world"
	n := len(str)
	for i := 0; i < n; i ++ {
		ch := str[i] // 依据下标取字符串中的字符 类型为byte
		fmt.Println(i, ch)
	}
}

func test_switch() {
	i := 10
	switch i {
		case 1:
			fmt.Println("i = 1")
		case 2, 3, 4:
			fmt.Println("i = 2 或者 i = 3 或者 i = 4")
		case 10:
			fmt.Println("i = 10")
		default:
			fmt.Println("i 是一个整数")
	}

	integer := 6
	switch integer {
		case 4:
			fmt.Println("当前整数小于4")
		    fallthrough
		case 5:
			fmt.Println("当前整数小于5")
		    fallthrough
		case 6:
			fmt.Println("当前整数小于6")
		    fallthrough
		case 7:
			fmt.Println("当前整数小于7")
		    fallthrough
		default:
			fmt.Println("default case")
	}
}

func test_for1()  {
	str := "welcome to go world"

	for index, val := range str{
		fmt.Printf("index[%d] val[%c] len[%d]", index, val, len([]byte(string(val))))
		fmt.Println()
	}

	// 定义局部变量
	var a int = 100
	var b int = 200
	var ret int

	// 调用函数并返回最大值
	ret = max1(a, b)
	fmt.Printf("最大值是: %d", ret) // 200
}

func max1(num1, num2 int) int {
	//if num1 > num2 {
	//	return num1
	//}
	//return num2

	var result int
	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

func test_switch1()  {
	// fallthrough 会继续执行下一条
	x := 2
	switch x {
		case 1:
			fmt.Print("ricky 1")
		case 2:
			fallthrough
		case 3:
			fmt.Print("ricky 2")
		default:
			fmt.Print("ricky 3")
	}
}
func main()  {
	// test_if()

	// test_goto()

	// test_for()

	// test_for1()

	// test_switch()
	test_switch1()
}
