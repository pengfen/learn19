package main

import "fmt"
import "add"

/*
add包中有两个变量 Name和Age 访问这两个变量

注意变量首字母大写才可以访问
package add

var Name string = "welcome to go world"
var Age int = 10

*/
func test_imp() {
	fmt.Println("名字: ", add.Name)
	fmt.Println("年龄: ", add.Age)
}

/*
使用包别名来访问包中的函数 import "add" ---> import a "add"
*/
func test_alias() {
	fmt.Println("名字: ", a.Name)
	fmt.Println("年龄: ", a.Age)
}

/*
init函数自动
*/
func init() {
	fmt.Println("initialized")
}

var (
    aa = 3
    ss = 'kkk'
    bb = true
)

func variableZeroValue() {
	// var a int
	// var s string
	// fmt.Printf("%d %q\m", a, s)
}

func main() {
	// 测试包导入
	//test_imp()

	//test_alias()

	//variableZeroValue()
}