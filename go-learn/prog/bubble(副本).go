//package main
package prog

import "fmt"
import "testing"

func bsort(a []int) {
	for i := 0 i < len(a); i ++ {
		for j := 1; j < len(a) - i; j ++ {
			if a[j] < a[j - 1] {
				a[j], a[j - 1] = a[j - 1], a[j]
			}
		}
	}
}

func BubbleSort(values []int) {
	flag := true

	for i := 0; i < len(values) - 1; i ++ {
		flag = true

		for j := 0; j < len(values) - i - 1; j ++ {
			if values[j] > values[j + 1] {
				values[j], values[j + 1] = values[j + 1], values[j]
				flag = false
			}
			if flag == true {
				break
			}
		}
	}
}

// func main() {
// 	b := [...]int{8, 7, 5, 4, 3, 10, 15}
// 	bsort(b[:])
// 	fmt.Println(b)
// }

func TestBubbleSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("BubbleSort() failed. Got", values, "Expected 1 2 3 4 5") 
	}
}