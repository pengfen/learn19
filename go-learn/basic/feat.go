package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test_goroute()  { // 测试并发
	//go handler_goroute(100)
	//go handler_goroute1(300, 300)

	for i := 0; i < 10; i ++ {
		go handler_goroute(i)
	}

	// fmt.Println("time: ", time.Second) // 1s

	time.Sleep(time.Second) // 睡眠一秒
}

func handler_goroute1(a int, b int)  { // 并发测试执行函数
	sum := a + b
	fmt.Println(sum)
}

func handler_goroute(a int)  { // 并发测试执行函数
	fmt.Println(a)
}

func test_goroute1()  {
	var pipe chan int
	pipe = make(chan int, 1)
	go Add(100, 300, pipe)

	sum := <- pipe
	fmt.Println("sum = ", sum) // 400
}

func Add(a int, b int, c chan int)  {
	sum := a + b
	c <- sum
}

func test_goroute2()  {
	sum := Add1(100, 300)
	sub := Sub(100, 300)

	fmt.Println("sum = ", sum)
	fmt.Println("sub = ", sub)
}

func test_goroute3()  {
	message := make( chan string)
	go func(){ // 注意 协程中不能使用打印， 只能使用管道来存储数据
		message <- "welcome to go world"
	}()
	str := <- message
	fmt.Println(str)
}

func Add1(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

// chan  FIFO  first In first out 先进先出
func test_pro_con()  { // 测试管理生产者 --- 消费者
	c := make(chan int, 5)
	go produce(c)
	go consumer(c)
	time.Sleep(time.Second)
}

// 生产者
func produce(p chan <- int) { // 将生产者的数据放入管道
	for i := 0; i < 10; i ++ {
		p <- i
		fmt.Println("发送方: ", i)
	}
}

// 消费者
func consumer(c <- chan int)  { // 从管道中取出数据
	for i := 0; i < 10; i ++ {
		v := <- c
		fmt.Println("接收方: ", v)
	}
}

func test_pro_con1() {
	// producer(生产者) --- channel(通过通道传输数据) ---> customer(消费者)

	// 创建一个字符串类型的通道
	channel := make(chan string)
	// 创建producer()函数的并发goroutine
	go producer("cat", channel)
	go producer("dog", channel)

	// 数据消费函数
	customer(channel)
}

// 数据生产者
func producer(header string, channel chan <- string)  {
	// 无限循环 不停地生产数据
	for {
		// 将随机数和字符串格式化为字符串发送给通道
		channel <- fmt.Sprintf("%s: %v", header, rand.Int())

		// 等待一秒
		time.Sleep(time.Second)
	}
}

// 数据消费者
func customer(channel <- chan string)  {
	// 不停地消费数据
	for {
		// 从通道中取出数据 此处会阻塞直接通道中返回数据
		message := <- channel

		// 打印数据
		fmt.Println(message)
	}
}
func test_chan()  { // 测试管道
	ch := make(chan int, 1) // 带有缓冲区的通道

	ch <- 1 // 向通道中写入数据
	go func() {
		v := <- ch // 从通道中读出数据
		fmt.Println(v) // 1
	}()

	fmt.Println("主线程") // 打印主线程 然后休眠一秒
	time.Sleep(time.Second)
}

func test_pip() { // 测试管道
    pipe := make(chan int, 3)
    pipe <- 1
    fmt.Println(pipe) // 地址 0xc420018100

    pipe <- 2
    fmt.Println(pipe) // 地址 0xc420018100

    pipe <- 3
    fmt.Println(pipe) // 地址 0xc420018100

    var t1 int
    t1 =<- pipe

    pipe <- 4
    fmt.Println(t1) // 1
    fmt.Println(len(pipe)) // 3
}

func calc(a int, b int) (int, int) {
	sum := a + b
	avg := (a + b) / 2
	return sum, avg
}
func test_more() {
	a := 100
	b := 200

	c, d := calc(a, b)
	fmt.Println(c) // 300
	fmt.Println(d) // 150
}

func main()  {
	// 特性 --- 并发
	// test_goroute() // 并发打印的数字并不是顺序的
	// test_goroute1()
	// test_goroute2()
	test_goroute3()

	// 测试生产者 --- 消费者
	// test_pro_con()
	// test_pro_con1()

	// 特性 --- 管道
	// test_chan()

	// 测试管道
	// test_pip()

	// 没底多返回值
	// test_more()
}
