package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

/**
1. 运行chatroom_server.go ---> 查看控制台
2. 运行chatroom_client.go ---> 查看控制台 ---> 查看服务端控制台

改进一
chatroom_client.go 添加控制台发送消息

改进二 服务器端进行消息转发
注意　客户端控制台消息格式　127.0.0.1:53938#welcome

改进三 解决以下问题
1. 如果发送的消息里面还有# 那就出现问题
2. 目前错误的处理　都是抛出异常　实际加上日志
3. 链接失败退出时　未清除map
4. 客户端 如果输入exit
5. 客户端 list
*/

/************ 改进三 start ************/
const (
	LOG_DIRECTORY = "./test.log" // 记录错误日志的路径
)
var logFile *os.File
var logger *log.Logger
/************ 改进三 end ************/

/************ 改进二 start ************/
var onlineConns = make(map[string]net.Conn) // 存储客户端连接映射　key为链接ip:port value为连接对象conn
var messageQueue = make(chan string, 1000) // 消息队列　带缓冲的buf

var quitChan = make(chan bool)
/************ 改进二 end ************/

func CheckError(err error)  {
	if err != nil {
		//fmt.Println("Error: %s", err.Error())
		//os.Exit(1)
		panic(err) // 抛出异常
	}
}

// 负责接收信息的协程
func ProcessInfo(conn net.Conn)  {
	buf := make([]byte, 1024)

	/************ 改进三 start ************/
	// defer conn.Close()
	// 协程退出时　说明客户端断开链接　所以要将当前链接从onlineConns删除掉
	defer func(conn net.Conn) {
		addr := fmt.Sprintf("%s")
		delete(onlineConns, addr)
		conn.Close()

		for i := range onlineConns {
			fmt.Println("new online conns: " + i)
		}
	}(conn) // 采用匿名函数的方式 调用defer
	/************ 改进三 end ************/

	for {
		numOfBytes, err := conn.Read(buf)
		if err != nil {
			break
		}

		// 如果接收字节数不为0 说明有消息发过来
		if numOfBytes != 0 {
			/************ 改进一 start ************/
			// 同时返回连接方的iP
			//remoteAddr := conn.RemoteAddr()
			//fmt.Print(remoteAddr) // 127.0.0.1:47480
			/************ 改进一 end ************/
			// 注意buf有可能会有上次的垃圾数据 string(buf[0:numOfBytes]) 每次只读取发送长度的数据
			//fmt.Printf("Has received this message: %s", string(buf))

			/************ 改进二 start ************/
			/* 结尾buf[0:numOfBytes]的原因是 numOfBytes是指接收的字节数　如果只用string(buf)
			   可能会导致接收字符串有其他之前接收的字符
			*/
			message := string(buf[0:numOfBytes])

			// 将消息放入到消息队列
			messageQueue <- message
			/************ 改进二 end ************/
		}
		//fmt.Println()
	}
}

/************ 改进二 start ************/

// 消费者协程
func ConsumeMessage()  {
	for {
		select {
			case message := <- messageQueue:
				// 对消息进行解析
				doProcessMessage(message)
			case <- quitChan:
				break
		}
	}
}

// 消息解析函数
func doProcessMessage(message string)  {
	contents := strings.Split(message, "#")
	if len(contents) > 1 {
		addr := contents[0]
		//sendMessage := contents[1]
		sendMessage := strings.Join(contents[1:], "#") // 防止消息体也含有 #

		addr = strings.Trim(addr, " ")

		// 通过addr查看是否有链接对象
		if conn, ok := onlineConns[addr]; ok {
			_, err := conn.Write([]byte(sendMessage))
			if err != nil {
				fmt.Println("online conns send failure!")
			}
		}
	} else { /************ 改进三 start ************/
		// 说明客户端调用list命令　查看系统当前链接IP
		contents := strings.Split(message, "*")
		if strings.ToUpper(contents[1]) == "LIST" {
			var ips string = ""
			for i := range onlineConns {
				ips = ips + "|" + i
			}
			if conn, ok := onlineConns[contents[0]]; ok {
				_, err := conn.Write([]byte(ips))
				if err != nil {
					fmt.Println("online conns send failure!")
				}
			}
		}
	}
	/************ 改进三 end ************/
}
/************ 改进二 end ************/

func main() {
	/************ 改进三 start ************/
	// 打开日志文件
	logFile, err := os.OpenFile(LOG_DIRECTORY, os.O_RDWR|os.O_CREATE, 0)
	if err != nil {
		fmt.Println("log file create failure!")
		os.Exit(-1)
	}

	defer logFile.Close()

	//利用go自带的log 将打开文件对象生成日志文件对象
	logger = log.New(logFile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)

	/************ 改进三 end ************/

	// 开启监听socket 监听在本地8080端口
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	CheckError(err)
	defer listen_socket.Close()

	fmt.Println("Server is waiting ...")

	/************ 改进三 start ************/
	logger.Println("I am writing the logs ...")
	/************ 改进三 end ************/


	/************ 改进二 start ************/
	go ConsumeMessage()
	/************ 改进二 end ************/

	for {
		conn, err := listen_socket.Accept() // 服务器端监听
		CheckError(err)

		/************ 改进二 start ************/
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		onlineConns[addr] = conn
		for i := range onlineConns { // 打印已连接客户端
			fmt.Println(i) // 127.0.0.1:53938 ---> 127.0.0.1:53938#welcome
		}
		/************ 改进二 end ************/

		// 如果有客户端链接　则打开一个协程处理
		go ProcessInfo(conn)
	}
}