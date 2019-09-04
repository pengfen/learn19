package main

import "fmt"

type person struct {
	name string
	age int
}

func test_struct()  {
	var p person
	p.name = "ricky"
	p.age = 18
	fmt.Printf("名字: %s", p.name) // 名字: ricky

	// 按照顺序提供初始化值
	//p := person{"ricky", "20"}

	// 通过 field:value 的方式初始化
	// p := person{age:24, name:"ricky"}
}

// 比较两个人的年龄 返回年龄大的那个人 并且返回年龄差
// struct 也是传值
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age { // 比较p1和p2这两个人的年龄
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func test_struct1()  {
	var tom person

	// 赋值初始化
	tom.name, tom.age = "ricky", 18

	bob := person{age:25, name:"bob"}

	// 按struct定义顺序初始化值
	paul := person{"Paul", 43}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("%s和%s, %s大%d岁", tom.name, bob.name, tb_Older.name, tb_diff) // ricky和bob, bob大7岁
	fmt.Println()
	fmt.Printf("%s和%s, %s大%d岁", tom.name, paul.name, tp_Older.name, tp_diff) // ricky和Paul, Paul大25岁
	fmt.Println()
	fmt.Printf("%s和%s, %s大%d岁", bob.name, paul.name, bp_Older.name, bp_diff) // bob和Paul, Paul大18岁
}

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human // 匿名字段 默认Student就包含了Human的所有字段
	speciality string
}

func test_struct2()  {
	// 初始化一个学生
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}

	// 相应的字段
	fmt.Println("名字: ", mark.name)
	fmt.Println("年龄: ", mark.age)
	fmt.Println("体重: ", mark.weight)
	fmt.Println("speciality ", mark.speciality)

	// 修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("speciality ", mark.speciality)

	// 修改年龄信息
	mark.age = 46
	fmt.Println("年龄: ", mark.age)

	// 修改体重信息
	mark.weight += 60
	fmt.Println("体重: ", mark.weight)
}

type Skills []string

type Stu struct {
	Human // 匿名字段 struct
	Skills // 匿名字段 自定义的类型 string slice
	int // 内置类型作为匿名字段
	speciality string
}

func test_struct3()  {
	// 初始化学生 Ricky
	ricky := Stu{Human:Human{"Ricky", 35, 100}, speciality:"Biology"}
	// 访问对应字段
	fmt.Println("名称: ", ricky.name)
	fmt.Println("年龄: ", ricky.age)
	fmt.Println("体重: ", ricky.weight)
	fmt.Println("学科; ", ricky.speciality)

	// 修改技能字段
	ricky.Skills = []string{"anatomy"}
	fmt.Println("技能: ", ricky.Skills)
	ricky.Skills = append(ricky.Skills, "physics", "golang")
	fmt.Println("技能: ", ricky.Skills)
	// 修改匿名内置类型字段
	ricky.int = 3
	fmt.Println("int: ", ricky.int)
}

type Employee struct {
	Human
	speciality string
	phone string // 雇员的phone字段
}

func test_struct4()  {
	Bob := Employee{Human{"Bob", 34, 100}, "designer", "135-8559-3461"}
	fmt.Println("电话: ", Bob.phone)
}

func main()  {
	// test_struct()

	// test_struct1()

	// test_struct2() // 匿名字段

	test_struct3() // 匿名字段

	test_struct4()
}
