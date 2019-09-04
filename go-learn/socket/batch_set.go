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

	_, err = c.Do("MSet", "abc", 200, "efg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}

    /*
    127.0.0.1:6379> mget abc efg
	1) "200"
	2) "300"
    */
	r, err := redis.Ints(c.Do("MGet", "abc", "efg"))
	if err != nil {
		fmt.Println("获取值失败", err)
		return
	}

	for _, v := range r {
		fmt.Println(v)
	}
}