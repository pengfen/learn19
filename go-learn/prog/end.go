package main

import "fmt"

// 完数 一个数恰好等于它的因子之和  6 = 1 = 2 + 3

func perfect(n int) bool  {
	var sum int = 0
	for i := 1; i < n; i ++ {
		if n % i == 0 {
			sum += i
		}
	}

	return n == sum
}

func process(n int)  {
	for i := 0; i < n + 1; i ++ {
		if perfect(i) {
			fmt.Println(i)
		}
	}
}

func main()  {
	var n int
	fmt.Scanf("%d", &n)
	process(n)
}