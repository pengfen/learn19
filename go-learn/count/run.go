package main

import (
	"flag"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

/**
命令行操作
go build run.go
./run

go run run.go --total=1000 --filePath=
./run --total=1000 --filePath=
*/

type resource struct {
	url string
	target string
	start int
	end int
}

func ruleResource() []resource {
	var res []resource
	r1 := resource{ // 首页
		url: "http://",
		target: "",
		start: 0,
		end: 0,
	}

	r2 := resource{ // 列表页
		url: "http",
		target: "{$id}",
		start: 1,
		end: 21,
	}

	r3 := resource{ // 详情页
	    url: "http://",
	    target: "{$id}",
	    start: 1,
	    end: 12924,
	}

	res = append(append(append(res, r1), r2), r3)
	return res
}

func buildUrl( res []resource ) []string {
	var list []string

	for _, resItem := range res {
		if len(resItem.target) == 0 {
			list = append( list, resItem.url )
		} else {
			for i := resItem.start; i <= resItem.end; i ++ {
				urlStr := strings.Replace( resItem.url, resItem.target, strconv.Itoa(i), -1)
				list = append(list, urlStr)
			}
		}
	}

	return list
}

func checkLog( current, refer, ua string ) string {
	u := url.Values{}
	u.Set( "time", "1")
	u.Set("url", current)
	u.Set("refer", refer)
	u.Set("ua", ua)
	paramsStr := u.Encode()

	logTemplate := "127.0.0.1 {$paramsStr} {$ua} sss"
	log := strings.Replace( logTemplate, "{$paramsStr}", paramsStr, -1)
	log = strings.Replace( log, "{$ua}", ua, -1)
	return log
}

func main()  {
	total := flag.Int("total", 100, "how many rows by created")
	filePath := flag.String("filePath", "/var/log/nginx/dig.log", "log file path")
	flag.Parse()

	// total 地址值 *total 默认值 100
	fmt.Println( *total, *filePath)

	// 需要构造出真实的网站url集合
	// list := []string
	res := ruleResource()
	list := buildUrl(res)
	fmt.Println(list)

	// 按照要求 生成 $total 行日志内容
	for i := 0; i <= *total; i ++ {
		//logStr := ...

		//ioutil.WriteFile( *filePath, []byte(logStr), 0644)
	}
	fmt.Println("done")
	fmt.Println()
}
