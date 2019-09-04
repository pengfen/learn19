package main

import (
	"fmt"
	"math"
)

/**
 判断 101 ~ 200 之间有多少素数 并输出所有素数
 定义 在大于1的自然数中 除了1和它本身以外不再有其他函数的数称为素数

判断素数的方法 用一个数分别去除2到sqrt(这个数) 如果能被整除 则表明此数不是素数 反之是素数
对正整数n 如果用2到根号n之间的所有整数去除 均无法整除 则n为质数
*/

func main()  {
	//var res bool
	//for i := 101; i <= 200; i ++ {
	//	res = isPrime(i)
	//	if (res) {
	//		fmt.Println(i)
	//	}
	//}

	/*
	var n int 是声明 但是没有赋值 它作用于main函数中
	n是值 &n是n的地址
	Scanf是一个函数 如果传入的是n 它是将n复制一份 传入Scanf进行处理 没有将main的n赋值成功 只是改变了赋值后的n
	传入&n地址 它指向n的存储地址 通过Scanf处理 可以真正改变n的值
	*/
	//var n int
	//var m int
	//fmt.Scanf("%d%d%s", &n, &m) // 控制台输入 101 300
	//for i := n; i < m; i ++ {
	//	if isPrime1(i) == true {
	//		fmt.Printf("%d\n", i)
	//		continue
	//	}
	//}

	// 判断输入的数是否是不是素数
	var n int
	fmt.Scanf("%d", &n)

	var flag bool = true
	for i := 2; i < n; i ++ {
		if n % i == 0 {
			flag = false
			break
		}
	}

	if flag == true {
		fmt.Printf("n[%d] 是素数", n)
		fmt.Println()
	} else {
		fmt.Printf("n[%d] 不是素数", n)
		fmt.Println()
	}
}

func isPrime(value int) bool  {
	if value <= 3 {
		return value >= 2
	}

	if (value % 2 == 0 || value % 3 == 0) {
		return false
	}

	for i := 5; i * i <= value; i += 6 {
		if value % i == 0 || value % (i + 2) == 0 {
			return false
		}
	}
	return true
}

// 判断输入的数是否为素数
func isPrime1(n int) bool{
	for i := 2; i <= int(math.Sqrt(float64(n))); i ++ {
		if n % i == 0 {
			return false
		}
	}
	return true;
}