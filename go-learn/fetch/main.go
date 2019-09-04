package main

import (
	"fetch/engine"
	"fetch/zhenai/parser"
)

// 注意转码 gbk ---> utf8
// 安装gopm ---> go get -v -u github.com/gpmgo/gopm
// gopm get -g -v golang.org/x/text
// 判断网页编码
// gopm get -g -v golang.org/x/net/html
func main()  {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
