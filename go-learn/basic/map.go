package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

func test_map()  {
	// map 声明一个key是字符串 值为int的字典 这种声明需要在使用之前使用make初始化
	// var numbers map[string] int
	// 使用make
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["ten"] = 10
	numbers["three"] = 3
	fmt.Println("第三个数字是: ", numbers["three"]) // 第三个数字是:  3

	// 初始化一个字典
	rating := map[string]float32 {"C":5, "Go":4.5, "Python":4.5, "C++":2}
	// map有两个返回值 第二个返回值 如里不存在key ok为false 否则为true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("存在", csharpRating)
	} else {
		fmt.Println("不存在", csharpRating) // 不存在 0
	}
	delete(rating, "C")

	m := make(map[string]string)
	m["wel"] = "ricky"
	m1 := m
	m1["wel"] = "peng"
}

func test_map1()  {
	m := map[string]string {
		"name": "ricky",
		"course": "golang",
		"site": "resource",
		"quality": "notbad",
	}
	m2 := make(map[string]int)
	var m3 map[string]int

	fmt.Println("n, m2, m3: ", m, m2, m3) // map[name:ricky course:golang site:resource quality:notbad] map[] map[]

	fmt.Println("Traversing map m")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(`m["course"] = `, courseName) // m["course"] =  golang
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key 'cause' does not exist") // key 'cause' does not exist
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Printf("m[%q] before delete: %q, %v", "name", name, ok) // m["name"] before delete: "ricky", true

	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("m[%q] after delete: %q, %v", "name", name, ok) // truem["name"] after delete: "", false
}

// PersonInfo 是一个包含个人详细信息的类型
type PersonInfo struct {
	ID string
	Name string
	Address string
}

func test_map2()  {
	var personDB map[string] PersonInfo // personDB map变量名 string键类型 PersonInfo 存放的值类型
	personDB = make(map[string] PersonInfo)

	// 向这个map昊添加数据 注意对象信息使用{}
	personDB["1234"] = PersonInfo{"12345", "ricky", "上海松江"}
	personDB["1"] = PersonInfo{"1", "peng", "上海徐汇"}

	// 从这个map查找键为1234的信息
	person, ok := personDB["1234"]

	// ok 是一个返回的bool类型 返回true表示找到了对应的数据
	if ok {
		fmt.Println("个人信息", person.Name, " ID为1234") // 个人信息 ricky  ID为1234
	} else {
		fmt.Println("没找到个人信息")
	}

	// 删除元素
	delete(personDB, "1234")
}

func test_map3()  {
	var a map[string]string = map[string]string {"key":"value"}

	a["abc"] = "efg"
	a["abc"] = "efg"
	a["abc1"] = "efg"
	fmt.Println(a)
}

func test_map4()  {
	a := make(map[string]map[string]string, 100)
	a["key1"] = make(map[string]string)
	a["key1"]["key2"] = "abc"
	a["key1"]["key3"] = "abc"
	a["key1"]["key4"] = "abc"
	a["key1"]["key5"] = "abc"
	fmt.Println(a)
}

func test_map5()  {
	a := make(map[string]map[string]string, 100)
	modify1(a)
	fmt.Println(a)
}

func modify1(a map[string]map[string]string)  {
	_, ok := a["zhangsan"]
	if !ok {
		a["zhangsan"] = make(map[string]string)
	}

	a["zhangsan"]["passwd"] = "123456"
	a["zhangsan"]["nickname"] = "pangpang"
	return
}

func test_map6()  {
	a := make(map[string]map[string]string, 100)
	a["key1"] = make(map[string]string)
	a["key1"]["key2"] = "abc"
	a["key1"]["key3"] = "abc"
	a["key1"]["key4"] = "abc"
	a["key1"]["key5"] = "abc"

	a["key2"] = make(map[string]string)
	a["key2"]["key2"] = "abc"
	a["key2"]["key3"] = "abc"

	trans(a)
	delete(a, "key1")
	fmt.Println()
	trans(a)
	fmt.Println(len(a))
}

func trans(a map[string]map[string]string)  {
	for k, v := range a {
		fmt.Println(k)
		for k1, v1 := range v {
			fmt.Println(k1, v1)
		}
	}
}

func test_map7()  {
	var a []map[int]int
	a = make([]map[int]int, 5)
	if a[0] == nil {
		a[0] = make(map[int]int)
	}
	a[0][10] = 10
	fmt.Println(a)
}

func test_map8()  {
	var a map[int]int
	a = make(map[int]int, 5)

	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	var keys []int
	for k, _ := range a {
		keys = append(keys, k)
		// fmt.Println(k, v)
	}

	sort.Ints(keys)
	for _, v := range keys {
		fmt.Println(v, a[v])
	}
}

func test_map9()  {
	var a map[string]int
	var b map[int]string

	a = make(map[string]int, 5)
	b = make(map[int]string, 5)

	a["abc"] = 101
	a["efg"] = 10

	for k, v := range a {
		b[v] = k
	}
	fmt.Println(b)
}

var lock sync.Mutex
var rwLock sync.RWMutex

func test_map10()  {
	var a map[int]int
	a = make(map[int]int, 5)

	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	for i := 0; i < 2; i ++ {
		go func(b map[int]int) {
			lock.Lock()
			b[8] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()
	time.Sleep(time.Second)
}

func test_map11()  {
	var a map[int]int
	a = make(map[int]int, 5)
	var count int32

	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	for i := 0; i < 2; i ++ {
		go func(b map[int]int) {
			//rwLock.Lock()
			lock.Lock()
			b[8] = rand.Intn(100)
			time.Sleep(10 * time.Millisecond)
			lock.Unlock()
			//rwLock.Unlock()
		}(a)
	}

	for i := 0; i < 100; i ++ {
		go func(b map[int]int) {
			for {
				lock.Lock()
				// rwLock.RLock()
				time.Sleep(time.Millisecond)
				// fmt.Println(a)
				// rwLock.RUnlock()
				lock.Unlock()
				atomic.AddInt32(&count, 1)
			}
		}(a)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count))
}

func test_map12()  {
	//var x map[string]float32
	//x["ricky"] = 63.2
	//fmt.Printf("%v", x) // 会报错 因为map使用前必须make分配内存 否则会报nil map错误

	// 正常声明
	var x map[string]float32
	x = make(map[string]float32)
	x["ricky"] = 63.2
	fmt.Printf("%v", x)

	// 或者直接使用简短声明符号
	y := make(map[string]float32)
	y["ric"] = 76.5
	fmt.Printf("%v", y)
}

func main()  {
	// test_map()

	//test_map1()

	test_map2()

	test_map3()

	test_map4()

	test_map5()

	test_map6()

	test_map7()

	// 排序
	test_map8()

	test_map9()

	test_map10()

	test_map11()

	test_map12()
}
