package main

import (
	"fmt"
)

// slice 本身没有数据 是对底层array的一个view

func test_slice()  {
	// 定义一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 切片 [first:last] [:] 所有元素  [:5] 0~4 [5:] 5-最后
	// make 创建数组切片 make([]int, 5) make([]int, 5, 10)
	// 基于数组创建一个数组切片
	var mySlice []int = myArray[:5]

	fmt.Println("数组元素: ")
	for _, v := range myArray {
		fmt.Print(v, " ") // 1 2 3 4 5 6 7 8 9 10
	}
	fmt.Println()

	fmt.Println("切片元素: ")
	for _, v := range mySlice {
		fmt.Print(v, " ") // 1 2 3 4 5
	}
	fmt.Println()
}

func test_slice1()  {
	mySlice := make([]int, 5, 10)

	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println("切片长度: ", len(mySlice))
	fmt.Println("切片容量: ", cap(mySlice))
	for _, v := range mySlice {
		fmt.Print(v, " ") // 0 0 0 0 0
	}
	fmt.Println()

	// mySlice = append(mySlice, 1, 2, 3) // 向mySlice追加元素
	mySlice2 := []int{8, 9, 10}
	// 给mySlice后面添加另一个数组切片
	mySlice = append(mySlice, mySlice2...)

	for _, v := range mySlice {
		fmt.Print(v, " ") // 0 0 0 0 0 8 9 10
	}
}

func test_slice2() {
	// slice 切片 (动态数组)
	// slice 并不是真正的动态数组 而是一个引用类型
	// 和声明array一样 只是少了长度
	var fslice []int
	fmt.Println(fslice) // []

	// 声明slice 并初始化数据
	slice := []byte {'a', 'b', 'c', 'd'}
	fmt.Println(slice) // [97 98 99 100]

	// 声明一个含有10个元素类型为byte的数组
	var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个含有byte的slice
	var a, b []byte

	// a指向数组的第三个元素开始 并到第五个元素
	a = ar[2:5]
	// b是数组ar的另一个slice
	b = ar[3:5]
	fmt.Println(a, b) // [99 100 101] [100 101]

	// 声明一个数组
	var array = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个slice
	var aSlice, bSlice []byte
	aSlice = array[:3] // 等价于 array[0:3]
	fmt.Println(aSlice) // [97 98 99] a b c
	aSlice = array[5:] // 等价于 array[5:10]
	fmt.Println(aSlice) // [102 103 104 105 106] f g h i j
	aSlice = array[:] // 等价于 array[0:10]
	fmt.Println(aSlice) // [97 98 99 100 101 102 103 104 105 106]

	aSlice = array[3:7] // len = 4, cap = 7
	fmt.Println(aSlice) // [100 101 102 103]
	bSlice = aSlice[1:3]
	fmt.Println(bSlice) // [101 102]
	bSlice = aSlice[:3] // aSlice[0:3]
	fmt.Println(bSlice) // [100 101 102]
	bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展
	fmt.Println(bSlice) // [100 101 102 103 104]
	bSlice = aSlice[:]
	fmt.Println(bSlice) // [100 101 102 103]

	// slice内置函数
	// len 获取slice的长度
	// cap 获取slice最大容量
	// append 向slice里面追加一个或者多个元素
	// copy 从源slice的src 复制元素到目标dst 并返回复制个数
}

func test_slice3()  {
	var slice []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}

	slice = arr[:]
	slice = slice[1:]
	slice = slice[:len(slice) - 1]

	fmt.Println(slice) // [2 3 4]
	fmt.Println(len(slice)) // 3
	fmt.Println(cap(slice)) // 4

	slice = slice[0:1]
	fmt.Println(len(slice)) // 1
	fmt.Println(cap(slice)) // 4
}

func test_slice4()  {
	var s1 slice
	s1 = make1(s1, 10)
	s1.ptr[0] = 100
	modify(s1)
	fmt.Println(s1.ptr)
}

type slice struct {
	ptr *[100]int
	len int
	cap int
}

func make1(s slice, cap int) slice {
	s.ptr = new ([100]int)
	s.cap = cap
	s.len = 0
	return s
}

func modify(s slice)  {
	s.ptr[0] = 1000
}

func test_slice5()  {
	var a [5]int = [...]int{1, 2, 3, 4, 5}
	s := a[1:]
	fmt.Printf("before len[%d] cap[%d]", len(s), cap(s))
	s[1] = 100
	fmt.Printf("s=%p a[1]=%p", s, &a[1])
	fmt.Println("before a:", a)

	s = append(s, 10)
	s = append(s, 10)
	fmt.Printf("after len[%d] cap[%d]", len(s), cap(s))
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)

	s[1] = 1000
	fmt.Println("after a: ", a)
	fmt.Println(s)
	fmt.Printf("s=%p a[1]=%p", s, &a[1])
}

func test_slice6()  {
	x := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var y []int
	y = x[1:3]
	fmt.Printf("%v", y) // [2 3]
	fmt.Println()

	y = x[:3] // 左边不写 代表从0开始
	fmt.Printf("%v", y) // [1 2 3]
	fmt.Println()

	y = x[7:] // 右边不写代表一直到最后
	fmt.Printf("%v", y) // [8 9 10]
}

func test_slice7()  {
	x := make([]int, 3, 5)
	x = append(x, 4, 5, 6, 7, 8)

	fmt.Printf("%v", x) // [0 0 0 4 5 6 7 8]
	fmt.Println()
	fmt.Printf("%v", len(x)) // 8   代表现在slice中元素的个数
	fmt.Println()
	fmt.Printf("%v", cap(x)) // 10  代表现在slice底层的分配的空间是多大 当append加入的元素个数超过此值时 底层会自动扩充
}

func main() {
	// test_slice()

	// test_slice1()

	// test_slice2()

	// test_slice3()

	// stest_slice4()

	// test_slice5()

	test_slice6()

	// test_slice7()
}
