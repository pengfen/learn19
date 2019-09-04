package main

import (
"fmt"
"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("连接redis失败", err)
		return
	}

	defer c.Close()

    /*
    127.0.0.1:6379> get abc
	"200"
	127.0.0.1:6379> get abc
	(nil)
    */
	_, err = c.Do("expire", "abc", 10) // 十秒过期
	if err != nil {
		fmt.Println(err)
		return
	}
}