package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/**
1. 客户端发一条数据给服务端

改进一 添加控制台发送消息

改进二 服务器端进行消息转发
注意　客户端控制台消息格式　127.0.0.1:53938#welcome

改进三
*/

func CheckError1(err error)  {
	if err != nil {
		panic(err)
	}
}

/************ 改进一 start ************/
func MessageSend(conn net.Conn)  {
	var input string
	for {
		// 接收系统标准输入
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		input = string(data)

		// 如果客户端输入exit 表示要结束连接
		if strings.ToUpper(input) == "EXIT" {
			conn.Close()
			break
		}

		_, err := conn.Write([]byte(input))
		if err != nil {
			conn.Close()
			fmt.Println("client connect failure: " + err.Error())
			break
		}
	}
}
/************ 改进一 end ************/

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	CheckError1(err)
	defer conn.Close()

	// conn.Write([]byte("welcome to here")) // 直接向服务端写入数据

	/************ 改进一 start ************/
	// 开启消息发送协程
	go MessageSend(conn)

	// 主协程负责接收消息
	buf := make([]byte, 1024)
	for {
		numOfBytes, err := conn.Read(buf)
		// CheckError1(err)
		/************ 改进三 start ************/
		// 如果客户端输入exit结束　不抛出异常　而结出提示消息
		if err != nil {
			fmt.Println("您已退出")
			os.Exit(0)
		}
		/************ 改进三 end ************/

		/*
		结尾buf[0:numOfBytes]的原因是 numOfBytes是指接收的字节数 如果只用string(buf)
		可能会导致接收字符串中有其他之前接收的字符
		*/
		fmt.Println("receive server message content: " + string(buf[0:numOfBytes]))
	}
	fmt.Println("Client program end!")
	/************ 改进一 end ************/
	
	//fmt.Println("sent the message!")
}
