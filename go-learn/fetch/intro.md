# 爬虫项目
* 有一定的复杂性
* 可以灵活调整项目的复杂性
* 平衡语言/爬虫之间的比重

# 网络爬虫分类
* 通用爬虫 如baidu google
* 聚焦爬虫 从互联网获取结构化数据

* 不使用现成爬虫库/框架
* 使用ElasticSearch作为数据存储
* 使用Go语言标准模板库实现http数据展示部分

# 使用css选择器
* 网页 console 控制台输出 $("#cityList")
* $("#cityList>dd>a")

# 使用xpath

# 使用正则表达式

# 解析器 Parser
* 输入 utf-8 编码的文本
* 输出 Request{URL 对应Parser}列表 Item列表
seed         parser
     \      /
      engine
         |  \ 
      任务队列 fetcher
* seed --- request ---> engine
* engine --- URL ---> Fetcher --- text ---> engine
* engine --- text ---> Parser --- requests, items ---> engine