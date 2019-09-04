# get请求

* 测试计划 ---> Add ---> Threads(users) ---> Thread Group
* Thread Group ---> Add ---> Sampler ---> HTTP Request
* Thread Group ---> Add ---> Listener ---> View Results Tree (会消耗大量IO等资源　一般不开启)
* HTTP Request
* protocol: http或https  Server Name or IP:域名　Port Number:端口号
* Method: GET/POST  Path:/v1/contr/action
* Body Data(POST请求参数) {uid:1,token:""}

## 解决乱码问题
* 1. HTTP Request 页 Content encoding: UTF-8
* 2. 修改bin/jmeter.properties
* 搜索/ISO
* sampleresult.default.encoding=UTF-8
* 3. Add ---> Post Processors ---> BeanShell Post Processor
* prev.setDataEncoding("utf-8")

## 聚合报告
* Add ---> Listener ---> Aggregate Report

## 参数化
* Function Helper Dialog ---> 选择_CSVRead
* CSV fiel ...   ---> 文件路径
* CSV文件列号 ... ---> 0
* 生成　---> 将生成的字符串
* /v1/cont/action/参数化字符串