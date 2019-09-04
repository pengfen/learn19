package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 文章控制器
type RedisController struct {
	beego.Controller
}

// 读写redis库
// 读SET/写GET
// 批量读MSET/写MGET
// 批量写入HMSET/读取HMGET
func (r *RedisController) Write() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		beego.Info("连接失败")
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superWang", "EX", "5") // 按秒计算
	checkErr(err)
	beego.Info("添加成功")

	username, err := redis.String(c.Do("GET", "mykey"))
	checkErr(err)
	beego.Info(username)
	time.Sleep(10 * time.Second)
	username, err = redis.String(c.Do("GET", "mykey"))
	checkErr(err)
	beego.Info(username)
	r.Ctx.WriteString("添加成功")
}

// 检测值是否存在
func (r *RedisController) Check() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		beego.Info("连接失败")
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superWang")
	checkErr(err)

	is_key_exit, err := redis.Bool(c.Do("EXISTS", "my"))
	checkErr(err)
	beego.Info(is_key_exit)
	r.Ctx.WriteString("添加成功")
}

// 删除key
func (r *RedisController) Delete() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		beego.Info("连接失败")
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superWang")
	checkErr(err)

	username, err := redis.String(c.Do("GET", "mykey"))
	checkErr(err)
	beego.Info(username)

	_, err = c.Do("DEL", "mykey")
	checkErr(err)

	username, err = redis.String(c.Do("GET", "mykey"))
	checkErr(err)
	beego.Info(username)
	r.Ctx.WriteString("删除成功")
}

// 处理json
func (r *RedisController) Json() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		beego.Info("连接失败")
		return
	}
	defer c.Close()

	key := "profile"
	imap := map[string]string{"username":"ricky", "number":"888"}
	value, _ := json.Marshal(imap)

	n, err := c.Do("SETNX", key, value)
	checkErr(err)
	if n == int64(1) {
		beego.Info("success")
	}

	var imapGet map[string]string
	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	beego.Info(imapGet["username"])
	beego.Info(imapGet["number"])
	r.Ctx.WriteString("处理json")
}