package main

import "fmt"

type Student struct {
	Name string
	Age int
	score float32
}

type treeNode struct {
	value int
	left, right *treeNode
}

// 使用自定义工厂函数
// 注意返回了局部变量的地址
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

// node 接收者 调用root.print()
// print(node treeNode) 调用print(root)
func (node treeNode) print() {
	fmt.Print(node.value)
}

// 只有使用指针才可以改变结构内容
// nil 指针也可以调用方法
func (node *TreeNode) setValue(value int) {
	node.Value = value
}

// 值接收者 指针接收者
// 要改变内容必须使用指针接收者
// 结构过大也考虑使用指针接收者
// 一致性 如有指针接收者 最好都是指针接收者
// 值接收者 是go语言特有
// 值/指针接收者均可接收值/指针

func test_object() {
	var root treeNode

	root = treeNode{value: 3}
	// 不论地址还是结构本身 一律使用.来访问成员
	root.left = &treeNode{}
	root.right = treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	nodes := []treeNode {
		{value:3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.left.right = createNode(2)
}

func main() {
	test_object()

	var stu Student

	stu.Age = 18
	stu.Name = "ricky"
	stu.score = 90

	var stu1 *Student = &Student{
		Age: 20,
		Name: "welcome",
	}

	var stu3 = Student{
		Age: 20,
		Name: "ricky",
	}
	fmt.Println(stu1.Name)
	fmt.Println(stu3)
	fmt.Printf("Name:%p\n", &stu.Name)
	fmt.Printf("Age:%p\n", &stu.Age)
	fmt.Printf("score:%p\n", &stu.score)
}