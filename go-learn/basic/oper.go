/*
需求 求二个数的加减乘除
*/
package main

import "fmt"

func main() {
    var a int = 21
    var b int = 10
    var c int

    //c = a + b
    fmt.Printf("加 %d\n", a + b)
    c = a - b
    fmt.Printf("减 %d\n", c)
    c = a * b
    fmt.Printf("乘 %d\n", c)
    c = a / b
    fmt.Printf("除 %d\n", c)
    c = a % b
    fmt.Printf("余 %d\n", c)
    a--
    fmt.Printf("递减 %d\n", a)
}