package main

// 每一个test文件须import一个testing
import (
	"fmt"
	"testing"
)

/**
 * /d/go/main_test.go
 * cmd ---> D: ---> cd go
 * go test
 * test文件下的每一个test case均必须以Test开头并且符合TestXxx形式 否则go test会直接跳过测试不执行
 * test case 的入参为t *testing.T 或 b *testing.B
 * t.Errorf为打印错误信息 并且当前test case会被跳过
 *
 * Test注意要点
 * 不会保证多个TestXxx是顺序执行 但是通常会按顺序执行
 * t.Run("a1", func(t *testing.T) {fmt.Println("a1")})
 * 使用t.Run来执行subtests可以做到控制test输出以及test的顺序
 * 使用TestMain作为初始化test 并且使用m.Run()来调用其他tests可以完成一些需要初始化操作的testing 比如数据库连接 文件打开 REST服务登录等
 * 如果没有在TestMain中调用m.Run()则除了TestMain以外的其他tests都不会被执行
 */
func TestPrint(t *testing.T)  {
	// t.SkipNow() t.SkipNow()为跳过当前test 并且直接按PASS处理继续下一个test
	res := Print1to20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("Wrong result of Print1to20")
	}
}
