package main

import "bufio"
import "flag"
import "fmt"
import "io"
import "os"
import "strconv"

var infile *string = flag.String("i", "unsorted.dat", "文件包含用于排序的值")
var outfile *string = flag.String("o", "sorted.dat", "接收排序值的文件")
var algorithm *string = flag.String("a", "qsort", "排序算法")

func readValue(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("输入源打开失败", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("错误")
			return
		}

		str := string(line) // 转换字符数组为字符串

		value, err1 := strconv.Atoi(str) // 

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("创建输出文件失败", outfile)
		return
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
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

/*
go build sorter
go install sorter
cd bin
./sorter -i unsorted.dat -o sorted.dat -a qsort
*/
func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile=", *infile, "outfile=", *outfile, "algorithm=", "algorithm")
	}
	values, err := readValue(*infile)

	if err == nil {
		//fmt.Println("读入的值:", values)
		t1 := time.Now()
		switch *algorithm {
			case "qsort":
			 	QuickSort(values)
			case "bubblesort":
			 	BubbleSort(values)
			default:
			 	fmt.Println("Sorting algorithm", *algorithm, "is either unknown orunsupported.")
			}

		t2 := time.Now()
		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")
		writeValues(values, *outfile) 
	} else {
		fmt.Println(err)
	}
}