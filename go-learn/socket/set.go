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

	_, err = c.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

    /*
    ricky@ricky:~$ redis-cli 
	127.0.0.1:6379> get abc
	"100"
    */
	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("获取abc失败", err)
		return
	}

	fmt.Println(r)
}