package main

import (
    "fmt"
    "time"
)

func test() {
	time.Sleep(time.Millisecond * 100)
}

func test_time() {
	curr := time.Now()
	fmt.Printf("%2d/%2d/%2d %2d:%2d:%2d", curr.Year(), curr.Month(), curr.Day(), curr.Hour(), curr.Minute(), curr.Second())

	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05"))

	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()

	fmt.Printf("cost:%d \n", (end - start) / 1000)
}

func test_mirco() {
		//记录开始时间
	//start := time.Nanoseconds()
	start := time.Now()
 
	//计算过程
	sum := 0
	for i := 0; i <= 100000000; i++{
		sum += i
	}
 
	//记录结束时间
	//end := time.Nanoseconds()
 
	//输出执行时间 单位为微秒
	//fmt.Println((end - start) / 1000000000)
 
	//输出执行结果
	//fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func main() {
	test_time()

	test_mirco()
}