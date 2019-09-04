package main

import (
"fmt"
"io/ioutil"
"net/http"
)

func main() {
	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("错误", err)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("获取数据错误", err)
		return
	}

	fmt.Println(string(data))
}