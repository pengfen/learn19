package parser

import (
	"fetch/engine"
	"fetch/model"
	"regexp"
	"strconv"
)

//const ageRe = `<td><span class="label">年龄：</span>([d+])岁</td>`
//const marriageRe = `<td><span class="label"婚况: </span>[^<]+</td>`
// 进行预编译
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高: </span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入: </span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重: </span>(\d+)CM</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别: </span>(\d+)CM</td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座: </span>(\d+)CM</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label"婚况: </span>[^<]+</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历: </span>(\d+)CM</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业: </span>(\d+)CM</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯: </span>(\d+)CM</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件: </span>(\d+)CM</td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车: </span>(\d+)CM</td>`)

func ParseProfile (contents []byte) engine.ParseResult {
	profile := model.Profile{}
	//re := regexp.MustCompile(ageRe)
	//match := re.FindSubmatch(contents)
	//
	//if match != nil {
	//	age, err := strconv.Atoi(string(match[1]))
	//	if err != nil {
	//		profile.Age = age
	//	}
	//}

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Weight = weight
	}

	profile.Income = extractString(contents, incomeRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Car = extractString(contents, carRe)
	profile.Education = extractString(contents, educationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.House = extractString(contents, houseRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}