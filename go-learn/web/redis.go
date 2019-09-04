package main

import (
	"fmt"
	"github.com/astaxie/goredis"
	"github.com/garyburd/redigo/redis"
)

/*
redis 安装

https://gitee.com/pengfen/learn18/blob/master/data-learn/redis/install.html

redis 是个开源的高性能的key-value的内存数据库 可以把它当成远程的数据结构

支持的value类型非常多 比如string list(链表) set(集合) hash表等

redis性能非常高 单机能够达到15w qps 通常适合做缓存


redis扩展安装
go get github.com/garyburd/redigo/redis
*/

func test_redis()  {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	checkErr("连接失败", err)
	fmt.Println("连接成功")
	defer c.Close()
}

func test_set()  {
	c, err := redis.Dial("tcp", "localhost:6379")
	checkErr("连接redis失败", err)

	defer c.Close()

	_, err = c.Do("Set", "abc", 100)
	checkErr("", err)

	// redis-cli ---> get abc
	r, err := redis.Int(c.Do("Get", "abc"))
	checkErr("获取abc失败", err)
	fmt.Println(r) // 100
}

/**
 * 处理错误
 */
func checkErr(str string, err error)  {
	if err != nil {
		// panic(err)
		fmt.Println(str, err)
		return
	}
}

func test_hash()  {
	c, err := redis.Dial("tcp", "localhost:6379")
	checkErr("连接失败", err)
	defer c.Close()

	_, err = c.Do("HSet", "books", "abc", 100)
	checkErr("", err)

	// hget books abc
	r, err := redis.Int(c.Do("HGet", "books", "abc"))
	checkErr("获取abc失败", err)
	fmt.Println(r) // 100
}

func test_batch()  {
	c, err := redis.Dial("tcp", "localhost:6379")
	checkErr("连接redis失败", err)
	defer c.Close()

	_, err = c.Do("MSet", "abc", 200, "efg", 300)
	checkErr("", err)

	// mget abc efg
	r, err := redis.Ints(c.Do("MGet", "abc", "efg"))
	checkErr("获取值失败", err)
	for _, v := range r {
		fmt.Println(v) // 200 300
	}
}

func test_expire()  {
	c, err := redis.Dial("tcp", "localhost:6739")
	checkErr("连接redis失败", err)
	defer c.Close()

	/*
	127.0.0.1:6379> get abc
	"200"
	127.0.0.1:6379> get abc
	(nil) #过期后为空
	*/
	_, err = c.Do("expire", "abc", 10)
	checkErr("", err)
}

func test_list()  {
	c, err := redis.Dial("tcp", "localhost:6379")
	checkErr("连接mysql失败", err)
	defer c.Close()

	_, err = c.Do("lpush", "book_list", "abc", "efg", 300)
	checkErr("", err)

	r, err := redis.String(c.Do("lpop", "book_list"))
	checkErr("获取abc失败", err)
	fmt.Println(r) // 300
}

func test_redis2()  {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	err := client.Set("test", []byte("welcome to here"))
	if err != nil {
		panic(err)
	}

	res, err := client.Get("test")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// test hmset
	f := make(map[string]interface{})
	f["name"] = "ricky"
	f["age"] = 12
	f["sex"] = "male"

	err = client.Hmset("test_hash", f)
	if err != nil {
		panic(err)
	}

	// test zset
	_, err = client.Zadd("test_zset", []byte("ricky"), 100)
	if err != nil {
		panic(err)
	}
}

func main()  {
	// test_redis()

	// test_set()

	// test_hash()

	// test_batch() // 批量set

	// test_expire()

	test_list()

	test_redis2()
}
