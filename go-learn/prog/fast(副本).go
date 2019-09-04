package main
 
import "fmt"
 
func qsort(a []int, left, right int) {
    if left >= right {
        return
    }
 
    val := a[left]
    k := left
    //确定val所在的位置
    for i := left + 1; i <= right; i++ {
        if a[i] < val {
            a[k] = a[i]
            a[i] = a[k+1]
            k++
        }
    }
 
    a[k] = val
    qsort(a, left, k-1)
    qsort(a, k+1, right)
}

func quickSort(values []int, left, right int) {
    temp := values[left]
    p := left
    i, j := left, right

    for i <= j {
        for j >= p && values[j] >= temp {
            j--
        }

        if j >= p {
            values[p] = values[j]
            p = j
        }

        if values[i] <= temp && i <= p {
            i ++
        }

        if i <= p {
            values[p] = values[i]
            p = i
        }
    }
    values[p] = temp
    if p - left > 1 {
        quickSort(values, left, p - 1)
    }
    if right - p > 1 {
        quickSort(values, p + 1, right)
    }
}

func QuickSort(values []int) {
    quickSort(values, 0, len(values) - 1)
}
 
func main() {
    b := [...]int{8, 7, 5, 4, 3, 10, 15}
    qsort(b[:], 0, len(b)-1)
    fmt.Println(b)
}