package main

import (
	"fmt"
	"time"
)

func test_select()  {
	ch := make(chan int)

	go func(ch chan int) {
		ch <- 1
	}(ch)

	time.Sleep(time.Second)

	select {
		case <- ch:
			fmt.Print("come to read ch!")
		default:
			fmt.Print("come to default!")
	}
}

func test_select1()  {
	ch := make(chan int)
	timeout := make(chan int, 1)

	go func() {
		time.Sleep(time.Second)
		timeout <- 1
	}()

	select {
		case <- ch:
			fmt.Println("come to read ch!")
		case <- timeout:
			fmt.Print("come to timeout!")
	}
	fmt.Print("end of code!")
}

func test_select2()  {
	ch := make(chan int)

	select {
		case <- ch:
			fmt.Print("come to read ch!")
		case <- time.After(time.Second):
			fmt.Print("come to timeout!")
	}
	fmt.Print("end of code!")
}

func main() {
	// test_select()

	// test_select1()

	test_select2()
}
