package main

import (
	"fmt"
	"strconv"
	"time"
)

func sample(message chan string)  {
	message <- "goroutine 1"
	message <- "goroutine 2"
	message <- "goroutine 3"
	message <- "goroutine 4"
}

func sample2(message chan string)  {
	time.Sleep( 2 * time.Second)
	str := <- message
	str = str + "I'm goroutine"
	message <- str
	close(message)
}

func test_range()  {
	var message = make( chan string, 3 )
	go sample( message )
	go sample2( message )

	time.Sleep( 3 * time.Second )
	for str := range message { // 需要关闭chan ---> close(message) 否则会出错 all goroutines are asleep - deadlock
		fmt.Println( str )
	}
	fmt.Println(" welcome to go world")
}

func sample3(ch chan string)  {
	for i := 0; i < 5; i ++ {
		ch <- "num: " + strconv.Itoa(i)
	}
}

func sample4(ch chan int)  {
	for i := 0; i < 10; i ++ {
		ch <- i
	}
}

func test_range1()  {
	ch1 := make( chan string, 3 )
	ch2 := make( chan int, 5 )

	for i := 0; i < 10; i ++ {
		go sample3(ch1)
		go sample4(ch2)
	}

	for i := 0; i < 100; i ++ {
		select {
			case str, ok := <- ch1:
				if !ok {
					fmt.Println("ch1 failed")
				}
				fmt.Println(str)

			case in, ok := <- ch2:
				if !ok {
					fmt.Println("ch2 failed")
				}
				fmt.Println(in)
		}
	}
	time.Sleep(20 * time.Second)

}

func main() {
	// test_range()
	test_range1()
}
