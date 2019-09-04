package main

import "fmt"
import "time"

func handle_goroute(a int) { // 并发测试执行函数
	fmt.Println(a)
}

func handle_goroute2(a int, b int) { // 并发测试执行函数
	sum := a + b
	fmt.Println(sum)
}

func test_goroute() { // 测试并发
	// go handle_goroute(100)
	// go handle_goroute2(300, 300)

	for i := 0; i < 100; i++ {
		go handle_goroute(i)
	}

    fmt.Println("time: ", time.Second); // 1s

	time.Sleep(time.Second) // 睡眠一秒
}

func test_chan() { // 测试管道
    ch := make(chan int, 1) // 带有缓冲区的通道

	ch <- 1 // 向通道中写入数据
	go func() {
		v := <- ch // 从通道中读出数据
		fmt.Println(v)
	}()

	fmt.Println("2") // 先打印主线程 然后休眠一秒
	time.Sleep(time.Second)
}

func test_pro_con() { // 测试管理生产者-消费者
	c := make(chan int, 5)
	go produce(c)
	go consumer(c)
	time.Sleep(time.Second)
}

// 生产者
func produce(p chan <- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("发送方: ", i)
	}
}

// 消费者
func consumer(c <- chan int) {
	for i := 0; i < 10; i++ {
		v := <- c
		fmt.Println("接收方: ", v)
	}
}

func test_pip() { // 测试管理
	pipe := make(chan int, 3)
	pipe <- 1
	fmt.Println(pipe) // 地址

	pipe <- 2
	fmt.Println(pipe)

	pipe <- 3
	fmt.Println(pipe)
	

	var t1 int
	t1 =<- pipe


	pipe <- 4

    fmt.Println(t1) // 1
	fmt.Println(len(pipe)) // 3
}

func calc(a int, b int) (int, int) { // 测试多返回值 
	sum := a + b
	avg := (a + b) / 2
	return sum, avg
}

func test_more() { // 测试多返回值 
	a := 100
	b := 200

	c, d := calc(a, b)
	fmt.Println(c)
	fmt.Println(d)
}

func main() {
	// 特性 --- 并发
	//test_goroute()

	// 特性 --- 管道
	// test_chan()

	// 测试生产者-消费者
	//test_pro_con()

	// 特性 --- 管道
	//test_pip()

	test_more()
}