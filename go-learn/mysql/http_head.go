package main

import (
"fmt"
"net/http"
)

var url = []string {
	"http://www.baidu.com",
	"http://google.com",
	"http://taobao.com",
}

func main() {
	for _, v := range url {
		resp, err := http.Head(v)
		if err != nil {
			fmt.Printf("head %s 失败, err: %v\n", v, err)
			continue
		}

		fmt.Printf("头部成功, 状态: %v\n", resp.Status)
	}
}