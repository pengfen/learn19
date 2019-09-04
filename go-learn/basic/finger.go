package main

import "fmt"

func modify(a int) {
	a = 10
	return
}

func modify1(a *int) {
	*a = 10
}

func modify2(p *int) {
	fmt.Println(p)
	*p = 1000000
	return
}

func test_finger() {
	var a int = 5
	fmt.Println(&a); // 获取一个变量的地址使用&

	/*
	写一个函数 传入一个int类型的指针 并在函数中修改所指向的值
	在main函数中调用这个函数 并把修改前后的值打印到终端
	*/
	// = 赋值  := 声明变量并赋值
	b := 5
	// chan 通道 make用于内存分配
	c := make(chan int, 1)

	fmt.Println("b=", b) // 5
	fmt.Println("c=", c) // 地址

	modify(b)
	fmt.Println("b=", b) // 5
	modify1(&b)
	fmt.Println("b=", b) // 10
}

func test_finger1() {
	var d int = 10
	fmt.Println(&d)

	var p *int
	p = &d

	fmt.Println("p地址:", &p)
	fmt.Println("p的值:", p)
	fmt.Println("p指针:", *p)

	fmt.Println(*p)
	*p = 100
	fmt.Println(d)

	var e int = 999
	p = &e
	*p = 5

	fmt.Println(d)
	fmt.Println(e)

	modify2(&d)
	fmt.Println(d)
}

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
	return
}

func swap1(a int, b int) (int, int) {
	return b, a
}

func test() {
	var a = 100
	fmt.Println(a)

	for i := 0; i < 100; i++ {
		var b = i * 2
		fmt.Println(b)
	}
}

func test2() {
	var a int8 = 100
	var b int16 = int16(a)

	fmt.Printf("a=%d b=%d\n", a, b)
}

func main() {
	test_finger() // 测试指针

	test_finger1()

	first := 100
	second := 200

	first, second = second, first
	fmt.Println("first=", first)
	fmt.Println("second=", second)

	test()

	test2()
}