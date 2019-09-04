package main

import (
"fmt"
"math"
)

func isPrime(n int) bool {
	/*
	判断101-200之间有多少个素数 并输出所有素数
	*/
	for i := 2; i < int(math.Sqrt(float64(n))); i ++ {
		if n % i == 0 {
			return false
		}
	}
	return true;
}

func main() {
	var n int
	var m int
	fmt.Scanf("%d%d", &n, &m) // 增加提示信息

	/*
	var n int是声明 但是没有赋值 它作用于main函数中
	n是值 &n是n的地址
	Scanf是一个函数 如果传入的是n 它是将n拷贝一份 传入Scanf进行处理 没有将main的n赋值成功 只是改变了赋值后的n
	传入&n地址 它指向n的存储地址 通过Scanf处理 可以真正改变n的值
	*/
	fmt.Printf("%d %d\n", n, m)

	for i := n; i < m; i ++ {
		if isPrime(i) == true {
			fmt.Printf("%d\n", i)
			continue;
		}
	}
}