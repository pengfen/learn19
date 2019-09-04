package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIp string
}

type Serverslice struct {
	Servers []Server
}

func test_json()  {
	var s Serverslice

	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s) // {[{Shanghai_VPN 127.0.0.1} {Beijing_VPN 127.0.0.2}]}
}

type Studentj struct {
	Name string `json:"student_name"`
	Age int
}

func test_json1()  {
	// 对数组类型的json编码
	x := [5]int{1, 2, 3, 4, 5}
	s, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))

	// 对map类型进行json编码
	m := make(map[string]float64)
	m["ricky"] = 100.4
	s2, err2 := json.Marshal(m)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(string(s2))

	// 对象进行json编码
	student := Studentj{"ricky", 20}
	s3, err3 := json.Marshal(student)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(string(s3))

	// 对s3进行json解码
	var s4 interface{}
	json.Unmarshal([]byte(s3), &s4)
	fmt.Printf("%v", s4)
}

func test_json2()  {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte("ricky"))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x", Result)
}

func main() {
	test_json()

	test_json1()

	test_json2()
}