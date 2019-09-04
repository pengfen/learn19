package main

import "fmt"

func bsort(a []int)  {
	for i := 0; i < len(a); i ++ {
		for j := 1; j < len(a) - i; j ++ {
			if a[j] < a[j - 1] {
				a[j], a[j -1] = a[j - 1], a[j]
			}
		}
	}
}

func BubbleSort(values []int)  {
	flag := true

	for i := 0; i < len(values) - 1; i ++ {
		flag = true

		for j := 0; j < len(values) - i - 1; j ++ {
			values[j], values[j + 1] = values[j + 1], values[j]
			flag = false
		}
		if flag == true {
			break
		}
	}
}

func main() {
	//b := [...]int{8, 7, 5, 4, 3, 10, 15}
	//bsort(b[:])
	//fmt.Println(b) // [3 4 5 7 8 10 15]

	values := []int{5, 4, 3, 2, 1}
	BubbleSort(values)
	fmt.Println(values) // [1 2 3 4 5]
}
