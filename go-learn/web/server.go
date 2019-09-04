package main

import (
	"fmt"
	"net"
)

/*
socket 编程

服务端的处理流程
监听端口
接收客户端的连接
创建goroutine处理该连接

客户端的处理流程
建立与服务端的连接
进行数据收发
关闭连接
*/

func main()  {
	fmt.Println("start server ...")
	listen, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil {
		fmt.Println("listen failed, err: ", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn)  {
	defer conn.Close()

	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)

		if err != nil {
			fmt.Println("read err: ", err)
			return
		}
		fmt.Println("read: ", string(buf))
	}
}