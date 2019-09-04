package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var flag bool = true
	for i := 2; i < n; i++ {
		if n % i == 0 {
			flag = false
			break
		}
 	}

 	if flag == true {
 		fmt.Printf("n[%d] 是素数\n", n)
 	} else {
 		fmt.Printf("n[%d] 不是素数\n", n)
 	}
}