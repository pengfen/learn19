package main

/*
判断素数的方法 用一个数分别去除2到sqrt(这个数) 如果能被整除 则表明此数不是素数 反之是素数
对正整数n 如果用2到根号n之间的所有整数去除 均无法整除 则n为质数
质数大于等于2不能被它本身和1以外的数据整除
*/
import "fmt"

func main() {
	var res bool;
	for i := 101; i <= 200; i ++ {
		res = isPrime(i)
		if (res) {
			fmt.Println(i);
		}
	}

	var n int
	var m int
	fmt.Scanf("%d%d%s", &n, &m)
	for i := n; i < m; i++ {
		if isPrime1(i) == true {
			fmt.Printf("%d\n", i)
			continue
		}
	}
}

func isPrime(value int) bool {
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
	return true;
}

func isPrime1(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}
