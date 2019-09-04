package main

import "fmt"

/*
条件语句注意点
条件语句不需要使用括号将条件包含起来()
无论语句体内有几条语句 花括号{}都是必须存在的
左花括号{必须与if或者else处于同一行
在if之后 条件语句之前 可以添加变量初始化语句 使用;间隔
在有返回值的函数中 不允许将最终的return语句包含在if ... else 结构中
*/
func test_if() {
	bool := true

	if bool {
		fmt.Print("这个值是true")
	} else {
		fmt.Print("这个值是false")
	}

	x := 5
	if x < 10 {
		fmt.Println("x小于10")
	} else {
		fmt.Println("x大于等于10")
	}

	if y := 10; y < 10 {
		fmt.Println("x小于10")
	} else {
		fmt.Println("x大于等于10")
	}

	// 多条件
	integer := 3
	if integer == 3 {
		fmt.Println("integer等于3")
	} else if integer < 3 {
		fmt.Println("integer小于3")
	} else {
		fmt.Println("integer大于3")
	}
}

func bounded(v int) int {
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
}

func branch() {
	const filename = "abc.txt" // abc.txt文件中写一些内容 会打印出来
	// io/ioutil
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

/*
注意点
左花括号{必须与switch处于同一行
条件表达式不限制为常量或整数
单个case中 可以出现多个结果选项
与c语言等规则相反 go不需要用break来明确退出一个case
只有在case中明确添加fallthrough关键字 才会继续执行紧跟的下一个case
可以不设定switch之后的条件表达式 等同于if ... else ...
*/
func test_switch() {
	var i = 0
	switch i {
	    case 0:
	    case 1:
	    	fmt.Println(i)
	    case 2:
	    	fmt.Println(i)
	    default:
	    	fmt.Println("def")
	}

	switch i {
	    case 0:
	    	fallthrough
	    case 1:
	    	fmt.Println(i)
	    case 2:
	    	fmt.Println(i)
	    default:
	    	fmt.Println("def")
	}

	switch i {
	    case 0, 1:
	    	fmt.Println("1")
	    case 2:
	    	fmt.Println(i)
	    default:
	    	fmt.Println("def")
	}

	switch { // 注意 无条件表达式
	    case i > 0 && i < 10:
	    	fmt.Println("i > 0 and i < 10")
	    case i > 10 && i < 20:
	    	fmt.Println("i > 10 and i < 20")
	    default:
	    	fmt.Println("def")
	}

    // fallthrough 强制执行后面的case代码
	integer := 6
	switch integer {
	    case 4:
	    	fmt.Println("小于等于4")
	    	fallthrough
	    case 5:
	    	fmt.Println("小于等于5")
	    	fallthrough
	    case 6:
	    	fmt.Println("小于等于6")
	    	fallthrough
	    case 7:
	    	fmt.Println("小于等于7")
	    	fallthrough
	    case 8:
	    	fmt.Println("小于等于8")
	    default:
	    	fmt.Println("默认值")
	}
}

func grade(score int) string {
	g := ""
	switch {
		case score < 0 || score > 100:
			panic(fmt.Sprintf("错误分数: %", score))
		case score < 60:
			g = "F"
	    case score < 80:
	    	g = "C"
	    case score < 90:
	    	g = "B"
	    case score <= 100:
	    	g = "A"
	}
	return g
}

/*
注意点
左花括号{必须与for处于同一行
不支持以逗号为间隔的多个赋值语句 必须使用平行赋值的方式来初始化多个变量
break与其它语言有区别
*/
func test_for() {
	str := "welcome to go world"

	for i, v := range str {
		fmt.Println(i, v)
	}

	for i, v := range str {
		if i > 2 {
			continue
		}
		if (i > 3) {
			break
		}
		fmt.Println(i, v)
	}

	sum := 0
	for index := 0; index < 10; index ++ {
		sum += index
	}
	fmt.Println("和等于", sum)

	for k, v := range map {
		fmt.Println("map's key:",k)
        fmt.Println("map's val:",v)
	}

	for _, v := range map(
		fmt.Println("map's val:", v)
	)
}

func convertToBin(n int) string {
	result := ""

	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
/*
用于for循环中迭代数组(array) 切片(slice) 通道(channel) 集合(map)的元素
*/
func test_range() {
	// 求切片slice的和
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("和:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("索引:", i)
		}
	}

	kvs := map[string]string{"a":"apple", "b":"banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func test_goto() {
	LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	i := 0
	HERE:
	    fmt.Print(i)
	    i ++
	    if i == 5 {
	    	return
	    }
	    goto HERE

	i = 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("", i)
		i ++;
	}

	for i := 0; i <= 7; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	j := 9
	Here:
	    fmt.Println(j)
	    j ++
	    goto Here // 跳转至Here处
}

func main() {
	//test_if() // 测试if, if ... else

	//test_switch()

	test_for()

	//test_range() // range

	//test_goto() // 测试goto
}