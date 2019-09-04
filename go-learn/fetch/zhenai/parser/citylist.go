package parser

import (
	"fetch/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contens []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	//matches := re.FindAll(contens, -1)
	// 匹配()子元素
	matches := re.FindAllSubmatch(contens, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "City " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]), ParserFunc: ParseCity,
		})
		//fmt.Printf("%s\n", m)
		//fmt.Printf("City: %s, URL: %s", m[2], m[1])
		//for _, subMatch := range m {
		//	fmt.Printf("%s ", subMatch)
		//}
		//fmt.Println()
	}

	// 一共多少个
	//fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
