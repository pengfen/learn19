package main

import (
"fmt"
"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}
	defer c.Close()
}