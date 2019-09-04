package main

import "fmt"

// PersonInfo是一个包含个人详细信息的类型
type PersonInfo struct {
	ID string
	Name string
	Address string
}

func test_map() {
    m := map[string]string{
    	"name": "ricky",
    	"course": "golang",
    	"site": "resource",
    	"quality": "notbad"
    }

    m2 := make(map[string]int)

    var m3 map[string]int

    fmt.Println("m, m2, m3:")
    fmt.Println(m, m2, m3)

    fmt.Println("Traversing map m")
    for k, v := range m {
    	fmt.Println(k, v)
    }

    fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(`m["course"] =`, courseName)
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key 'cause' does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Printf("m[%q] before delete: %q, %v\n",
		"name", name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("m[%q] after delete: %q, %v\n",
		"name", name, ok)
}

func main() {
	test_map()

	var personDB map[string] PersonInfo // personDB map变量名  string键类型  PersonInfo 存放的值类型
	personDB = make(map[string] PersonInfo) // 使用make创新map 如何初始存储能力使用 make(map[string] PersonInfo, 100)

	// 往这个msp里添加数据 注意对象信息使用{}
	personDB["1234"] = PersonInfo{"12345", "ricky", "上海松江"}
	personDB["1"] = PersonInfo{"1", "peng", "上海徐汇"}

	// 从这个map查找键为1234的信息
	person, ok := personDB["1234"]

	// ok是一个返回的bool类型 返回true表示找到了对应的数据
	if ok {
		fmt.Println("个人信息", person.Name, " ID为1234")
	} else {
		fmt.Println("没找到个人信息")
	}

	// 删除元素
	delete(personDB, "1234")
}