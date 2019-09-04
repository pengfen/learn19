package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
)

/*
go语言仅支持封装 不支持继承和多态(继承和多态使用接口来做)

封装 名字一般使用CamelCase

*/


type Rectangle struct {
	width, height float64
}

func test_object()  {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	fmt.Println("面积: ", area(r1)) // 24
	fmt.Println("面积: ", area(r2)) // 36
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (r Rectangle) a() float64 {
	return r.width * r.height
}

func (c Circle) a() float64 {
	return c.radius * c.radius * math.Pi
}

/*
method方法一样 如果接收者不一样 method就不一样
method里面可以访问接收者的字段
调用method访问 就像struct里面访问字段一样
*/
func test_object1()  {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	fmt.Println("面积: ", r1.a()) // 24
	fmt.Println("面积: ", r2.a()) // 36

	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println("面积: ", c1.a())
	fmt.Println("面积: ", c2.a())
}

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte // Color 作为 byte的别名

type Box struct { // 箱子 长 宽 高 颜色
	width, height, depth float64
	color Color
}

type BoxList []Box // 箱子列表

func (b Box) Volume() float64 { // 箱子容量
	return b.width * b.height * b.depth
}

func (b *Box) setColor(c Color)  { // 设置箱子颜色
	b.color = c
}

func (b1 BoxList) BiggestsColor() Color { // 返回箱子列表中的最大颜色
	v := 0.00
	k := Color(WHITE)
	for _, b := range b1 {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

func (b1 BoxList) PaintItBlack() { // 把箱子列表中的箱子颜色都更改为黑色
	for i, _ := range b1 {
		b1[i].setColor(BLACK)
	}
}

func (c Color) String() string  { // 返回具体的颜色
	strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func test_object2()  {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Println("长度 %d", len(boxes))
	fmt.Println("第一个箱子容量: ", boxes[0].Volume())
	fmt.Println("最后一个箱子颜色: ", boxes[len(boxes) - 1].color.String())
	fmt.Println("最大颜色: ", boxes.BiggestsColor().String())
	fmt.Println("改为黑色: ", boxes.PaintItBlack)
	fmt.Println("第二个箱子颜色: ", boxes[1].color.String())
	fmt.Printf("最大颜色: ", boxes.BiggestsColor().String())
}

type Stu1 struct {
	Name string
	Age int
	score float32
}

func test_object3()  {
	var stu Stu1
	stu.Age = 18
	stu.Name = "ricky"
	stu.score = 90

	var stu1 *Stu1 = &Stu1{
		Age: 20,
		Name: "welcome",
	}

	var stu2 = Stu1{
		Age: 20,
		Name: "ricky",
	}

	fmt.Println(stu1.Name)
	fmt.Println(stu2)
	fmt.Printf("Name:%p", &stu.Name)
	fmt.Printf("Age:%p", &stu.Age)
	fmt.Printf("score:%p", &stu.score)
}

type treeNode struct {
	value int
	left, right *treeNode
}

// 值接收者 指针接收者
// 要改变内容必须使用指针接收者
// 结构过大也考虑使用指针接收者
// 一致性 如有指针接收者 最好都是指针接收者
// 值接收者 是go语言持有
// 值/指针接收者均可接收值/指针
func test_object4()  {
	var root treeNode
	root = treeNode{value:3}
	// 不论地址还是结构本身 一律使用.来访问成员
	root.left = &treeNode{}
	//root.right = treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	nodes := []treeNode {
		{value:3},
		{},
		{6, nil, &root},
	}
	fmt.Print(nodes)
	root.left.right = createNode(2)
}

// 使用自定义工厂函数
// 注意返回了局部变量的地址
func createNode(value int) *treeNode  {
	return &treeNode{value: value}
}

// node 接收者 调用root.print()
// print(node treeNode) 调用print(root)
func (node treeNode) print()  {
	fmt.Print(node.value)
}

// 只有使用指针才可以改变结构内容
// nil 指针也可以调用方法
func (node *treeNode) setValue(value int) {
	node.value = value
}

type Stu2 struct {
	Name string
	Age int
	Score float32
	next *Stu2
}

func test_object5()  {
	var head *Stu2 = new(Stu2)
	head.Name = "hua"
	head.Age = 18
	head.Score = 100

	//insertTail(head)

	insertHead(&head)
	trans1(head)

	var newNode *Stu2 = new(Stu2)
	newNode.Name = "stu1000"
	newNode.Age = 18
	newNode.Score = 100
	addNode(head, newNode)
	trans1(head)
}

func delNode(p *Stu2)  {
	var prev *Stu2 = p
	for p != nil {
		if p.Name == "stu6" {
			prev.next = p.next
			break
		}

		prev = p
		p = p.next
	}
}

func addNode(p *Stu2, newNode *Stu2)  {
	for p != nil {
		if p.Name == "stu9" {
			newNode.next = p.next
			p.next = newNode
			break
		}
		p = p.next
	}
}

func trans1(p *Stu2)  {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
	fmt.Println()
}

func insertHead(p **Stu2)  {
	for i := 0; i < 10; i ++ {
		stu := Stu2{
			Name: fmt.Sprintf("stu%d", i),
			Age: rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		stu.next = *p
		*p = &stu
	}
}

func insertTail(p *Stu2)  {
	var tail = p
	for i := 0; i < 10; i ++ {
		stu := Stu2{
			Name: fmt.Sprintf("stu%d", i),
			Age: rand.Intn(100),
			Score: rand.Float32() * 100,
		}

		tail.next = &stu
		tail = &stu
	}
}

type Stu3 struct {
	Name string
	Age int
	Score float32
	left *Stu3
	right *Stu3
}

func trans2(root *Stu3) {
	if root == nil {
		return
	}
	fmt.Println(root)
	trans2(root.left)
	trans2(root.right)
}

func test_object6()  {
	var root *Stu3 = new(Stu3)
	root.Name = "stu01"
	root.Age = 18
	root.Score = 100

	var left1 *Stu3 = new(Stu3)
	left1.Name = "stu02"
	left1.Age = 18
	left1.Score = 100
	root.left = left1

	var right1 *Stu3 = new(Stu3)
	right1.Name = "stu04"
	right1.Age = 18
	right1.Score = 100
	root.right = right1

	var left2 *Stu3 = new(Stu3)
	left2.Name = "stu03"
	left2.Age = 18
	left2.Score = 100
	left1.left = left2

	trans2(root)
}

type Stu4 struct {
	Name string `json:"student_name"`
	Age int `json:"age"`
	Score int `json:"score"`
}

func test_object7()  {
	var stu Stu4 = Stu4{
		Name: "stu01",
		Age: 18,
		Score: 80,
	}

	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode stu failed, err:", err)
		return
	}

	fmt.Println(string(data))
}

func main()  {
	test_object()

	test_object1()

	test_object2()

	test_object3()

	test_object4()

	test_object5()

	test_object6()

	test_object7()
}
