package main

import "fmt"

// make 用于内建类型(map slice channel) 的内存分配

func test_make()  {
	s1 := new([]int)
	fmt.Println(s1) // &[]

	s2 := make([]int, 10)
	fmt.Println(s2) // [0 0 0 0 0 0 0 0 0 0]

	*s1 = make([]int, 5)
	(*s1)[0] = 100
	s2[0] = 100
	fmt.Println(s1) // &[100 0 0 0 0]
	fmt.Println(s2) // [100 0 0 0 0 0 0 0 0 0]
	return
}

func main()  {
	test_make()
}