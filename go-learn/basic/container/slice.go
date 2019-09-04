package main

import "fmt"

// slice本身没有数据 是对底层array的一个view
// arr的值变为[0 1 10 3 4 5 6 7]

func updateSlice(s []int) {
	s[0] = 100
}

func test_slice() {
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

    fmt.Println("arr[2:6] =", arr[2:6])
    fmt.Println("arr[:6] =", arr[:6])

    s1 := arr[2:]
    fmt.Println("s1 =", s1)
    s2 := arr[:]
    fmt.Println("s2 =", s2)

    fmt.Println("After updateSlice(s1)")
    updateSlice(s1)

    fmt.Println(s1)
    fmt.Println(arr)

    fmt.Println("After updateSlice(s2)")
    updateSlice(s2)
    fmt.Println(s2)
    fmt.Println(arr)

    fmt.Println("Reslice")
    fmt.Println(s2)
    s2 = s2[:5]
    fmt.Println(s2)
    s2 = s2[2:]
    fmt.Println(s2)

    fmt.Println("Extending slice")
    arr[0], arr[2] = 0, 2
    fmt.Println("arr =", arr)
    s1 = arr[2:6]
    s2 = s1[3:5]
    fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
    fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))

    s3 := append(s2, 10)
    s4 := append(s2, 11)
    s5 := append(s2, 12)
    fmt.Println("s3, s4, s5 =", s3, s4, s5)
    fmt.Println("arr =", arr)

    fmt.Println("uncomment to see sliceOps demo")
}

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func test_slice1() {
	fmt.Println("创建切片")
	var s []int

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("复制切片")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("从切片中删除元素")
	s2 = append(s2[:1], s2[4:]...)
	printSlice(s2)

	fmt.Println("Propping fron front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Propping from back")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]

	fmt.Println(tail)
	printSlice(s2)
}

func main() {
	test_slice()

	test_slice1()
}