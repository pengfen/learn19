package main

import "fmt"

/*
数组是值类型
[10]int和[20]int是不同类型
调用func f(arr [10]int)会拷贝 数组
在go语言中一般不直接使用数组
*/

func test_array() {
	// 先定义一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 切片 [first:last]  [:] 所有元素 [:5] 0-4 [5:] 5-最后
	// make创建数组切片 make([]int, 5)  make([]int, 5, 10)
	// 基于数组创建一个数组切片
	var mySlice []int = myArray[:5]

	fmt.Println("数组元素:")
	for _, v := range myArray {
		fmt.Print(v, " ") // 1 2 3 4 5 6 7 8 9 10
	}

	fmt.Println("切片元素:")
	for _, v := range mySlice {
		fmt.Print(v, " ") // 1 2 3 4 5 
	}

	fmt.Println()
}

func test_array2() {
	mySlice := make([]int, 5, 10)

	for _, v := range mySlice {
		fmt.Print(v, " ") // 0 0 0 0 0
	}

	fmt.Println("切片长度:", len(mySlice)) // 5
	fmt.Println("切片容量:", cap(mySlice)) // 10

	//mySlice = append(mySlice, 1, 2, 3) // 向mySlice追加元素
	mySlice2 := []int{8, 9, 10}
	// 给mySlice后面添加另一个数组切片
	mySlice = append(mySlice, mySlice2...) // 注意 需要加...
}

func test_array3() {
    // var array [n]type n表示数组长度 type表示存储元素类型
    var arr [10]int // 声明了一个int类型的数组
    arr[0] = 42 // 数组下标是从0开始的
    arr[1] = 13 // 赋值操作
    fmt.Printf("第一个元素:%d\n", arr[0]) // 获取数据 返回42
    fmt.Printf("最后一个元素:%d\n", arr[9]) // 最后一个元素未赋值 默认返回0 

    x := [3]int{1, 2, 3} // 声明一个长度为3的int数组
    y := [10]int{1, 2, 3} // 声明一个长度为10的int数组
    z := [...]int{4, 5, 6} // go会自动根据元素个数计算长度
    fmt.Println(x, y, z)

    // 声明一个二维数组 该数组以两个数组作为元素 其中每个数组中又有4个int类型的元素
    doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
    // 内部元素和外部的一样 声明简化 直接忽略内部的类型
    easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
    fmt.Println(doubleArray, easyArray)

    // slice 切片 (动态数组)
    // slice 并不是真正的动态数组 而是一个引用类型 
    // 和声明array一样 只是少了长度
    var fslice []int
    fmt.Println(fslice)

    // 声明slice 并初始化数据
    slice := []byte {'a', 'b', 'c', 'd'}
    fmt.Println(slice)

    // 声明一个含有10个元素类型为byte的数组
    var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
    // 声明两个含有byte的slice
    var a, b []byte

    // a指向数组的第3个元素开始 并到第五个元素
    a = ar[2:5]

    // b是数组ar的另一个slice
    b = ar[3:5]
    fmt.Print(a, b)

    // 声明一个数组
    var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
    // 声明两个slice
    var aSlice, bSlice []byte
    aSlice = array[:3] // 等价于 array[0:3]
    fmt.Println(aSlice)
    aSlice = array[5:] // 等价于 array[5:10]
    fmt.Println(aSlice)
    aSlice = array[:] // 等价于 array[0:10]
    fmt.Println(aSlice)

    aSlice = array[3:7] // len = 4, cap = 7
    fmt.Println(aSlice)
    bSlice = aSlice[1:3]
    fmt.Println(bSlice)
    bSlice = aSlice[:3] // aSlice[0:3]
    fmt.Println(bSlice)
    bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展
    fmt.Println(bSlice)
    bSlice = aSlice[:]
    fmt.Println(bSlice)

    // slice内置函数
    // len 获取slice的长度
    // cap 获取slice最大容量
    // append 向slice里面追加一个或者多个元素
    // copy 函数copy从源slice的src中复制元素到目标dst 并返回复制个数

    // map
    // 声明一个key是字符串 值为int的字典 这种声明需要在使用之前使用make初始化
    //var numbers map[string] int
    // 使用make
    numbers := make(map[string]int)
    numbers["one"] = 1 
    numbers["ten"] = 10
    numbers["three"] = 3 

    fmt.Println("第三个数字是: ", numbers["three"])

    // 初始化一个字典
    rating := map[string]float32 {"C":5, "Go":4.5, "PYthon":4.5, "C++":2}
    // map有两个返回值 第二个返回值 如果不存在key ok为false 否则为true
    csharpRating, ok := rating["c#"]
    if ok {
    	fmt.Println("存在", csharpRating)
    } else {
    	fmt.Println("不存在", csharpRating)
    }

    delete(rating, "C")

    m := make(map[string]string)
    m["wel"] = "ricky"
    m1 := m
    m1["wel"] = "peng"
}

func printArray(arr [5]int) {
    arr[0] = 100
    for i, v := range arr {
        fmt.Println(i, v)
    }
}

func test_array4 {
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
    // 不仅range 任何地方都可能过_省略变量
    for _, v := range numbers {
        sum += v
    }
}

func main() {
	//test_array()

	//test_array2()

	test_array3()
}