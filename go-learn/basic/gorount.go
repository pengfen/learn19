package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

var ch chan int

func test_channel()  {
	ch <- 1
	fmt.Println("ch 1")
	ch <- 1
	fmt.Println("ch 2")
	ch <- 1
	fmt.Println("come to end goroutine 1")
}

func test_gorout()  {
	ch = make(chan int, 0) // 等价于 ch = make(chan int)都是不带缓冲的channel
	ch = make(chan int, 2) // 是带缓冲的channel
	go test_channel()
	time.Sleep(2 * time.Second)
	fmt.Println("running end!")
	<-ch

	time.Sleep(time.Second)
}

func test_gorout1()  {
	// 协程
	go func() {
		for i := 1; i < 100; i ++ {
			if i == 10 {
				// 主动出让cpu　使用的话　需要导入runtime包
				runtime.Gosched()
			}
			fmt.Println("rountine 1: " + strconv.Itoa(i))
		}
	}()

	// 协程
	go func() {
		for i := 100; i < 200; i ++ {
			fmt.Println("rountine 2: " + strconv.Itoa(i))
		}
	}()
	time.Sleep(time.Second)
}

func test_gorunt2()  {
	ch = make(chan int)
	// 协程
	go func() {
		for i := 1; i < 100; i ++ {
			if i == 10 {
				// 遇到阻塞
				<- ch
			}
			fmt.Println("rountine 1: " + strconv.Itoa(i))
		}
	}()

	// 协程
	go func() {
		for i := 100; i < 200; i ++ {
			fmt.Println("rountine 2: " + strconv.Itoa(i))
		}
		ch <- 1
	}()

	time.Sleep(time.Second)
}

func main() {
	test_gorout()

	test_gorout1()

	test_gorunt2()
}