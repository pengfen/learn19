package main

import "fmt"

/*
数组是值类型
[10]int 和 [20]int是不同类型
调用 func f(arr [10]int) 会拷贝 数组
在go语言中一般不直接使用数组
*/

func test_array()  {
	// 数组 定义方式  var arr [n]type
	// n 表示数组的长度 type表示存储元素的类型
	var arr [10]int // 声明一个int类型的数组
	arr[0] = 42 // 数组下标是从0开始的
	arr[1] = 13 // 赋值操作
	fmt.Printf("第一个元素是 %d", arr[0]) // 获取数据
	fmt.Println()
	fmt.Sprintf("最后一个元素是 %d", arr[9]) // 没赋值不会打印

	x := [3]int{1, 2, 3} // 声明一个长度为3的int数组 [1 2 3]
	y := [10]int{1, 2, 3} // 声明一个长度为10的int数组 其中前三个元素初始化为1 2 3 其它默认为0  [1 2 3 0 0 0 0 0 0 0]
	z := [...]int{4, 5, 6} // go会自动根据元素个数计算长度
	fmt.Println(x, y, z) // [1 2 3] [1 2 3 0 0 0 0 0 0 0] [4 5 6]

	// 声明一个二维数组 该数组以两个数组作为元素 其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}} // [[1 2 3 4] [5 6 7 8]]
	// 内部元素和外部的一样 声明简化 直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}} // [[1 2 3 4] [5 6 7 8]]
	fmt.Println(doubleArray, easyArray) // [[1 2 3 4] [5 6 7 8]]    [[1 2 3 4] [5 6 7 8]]
}

func printArray(arr [5]int)  {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func test_array1()  {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println("array definitions")
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	fmt.Println("printArray(arr1)")
	printArray(arr1)

	fmt.Println("printArray(arr3)")
	printArray(arr3)

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)

	sum := 0
	// 可通过_省略变量
	// 不仅 range 任何地方都可通过_省略变量
	for _, v := range arr3 {
		sum += v
	}
}

func test_array2()  {
	var a [10]int
	a[0] = 100
	fmt.Println(a)

	for i := 0; i < len(a); i ++ {
		fmt.Println(a[i])
	}
	for index, val := range a {
		fmt.Printf("a[%d]=%d", index, val)
	}
}

func test_array3()  {
	var a [10]int
	b := a

	b[0] = 100
	fmt.Println(a)
}

func test_array4(arr *[5]int)  {
	(*arr)[0] = 1000
}

func test_array5()  {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	var a1 = [5]int{1, 2, 3, 4, 5}
	var a2 = [...]int{38, 283, 48, 348, 387, 484}
	var a3 = [...]int{1: 100, 3: 200}
	var a4 = [...]string{1: "welcome", 3: "world"}

	fmt.Println(a) // [1 2 3 4 5]
	fmt.Println(a1) // [1 2 3 4 5]
	fmt.Println(a2) // [38 283 48 348 387 484]
	fmt.Println(a3) // [0 100 0 200]
	fmt.Println(a4) // [ welcome  world]
}

func test_array6(n int)  {
	var a []uint64
	a = make([]uint64, n)

	a[0] = 1
	a[1] = 1

	for i := 2; i < n; i ++ {
		a[i] = a[i - 1] + a[i - 2]
	}

	for _, v := range a {
		fmt.Println(v)
	}
}
func main()  {
	// test_array()

	// test_array1()

	// test_array2()

	// test_array3()

	//var a [5]int
	//test_array4(&a)
	//fmt.Println(a)

	// test_array5()

	test_array6(10)
}