package main

import (
"fmt"
"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}

	defer c.Close()

	_, err = c.Do("HSet", "books", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

    /*
    127.0.0.1:6379> hget books abc
    "100"
    */
	r, err := redis.Int(c.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println("获取abc失败", err)
		return
	}

	fmt.Println(r)
}