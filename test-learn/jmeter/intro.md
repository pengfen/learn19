# jmeter

## 分布式压测的配置
* 代理机(agent)配置
* 确保防火墙是关闭的
* Agent机上需要安装JDK Jmeter 并且配置好环境变量

## Jmeter组成
* 取样器 进行脚本逻辑控制
* 线程组 场景设置
* 监视器 监控我们的脚本运行　取得性能指标

## JMeter脚本的两种录制方式
* 使用badboy进行录制
* 使用代理方式进行录制

## Jmeter实现多并发
* 线程组　负载发生器　用多线程或多进程的方式来模拟用户的使用行为　JMeter是以线程的方式来进行模拟用户的并发访问的
* 测试计划 ---> Add ---> Threads(Users) ---> Thread Group

## Jmeter实现逻辑分支控制 (if.md)
* 逻辑控制器　用来控制测试脚本的逻辑判断　也可以理解为如何控制脚本的运行
* 例如　如果控制器　就是当满足什么样的条件后执行哪一步操作
* Thread Group ---> Add ---> Logic Controller ---> If Controller

## Jmeter相关操作
* 自定义变量  把IP地址　用一个变量去代替　环境变化时　改一下变量值即可
* Add ---> Config Element ---> User Defined Variables
* 函数助手　__CSVREAD
* csv data set config (Add ---> Config Element ---> CSV Data Set Config)


## 关联 (当上文有一些变量的值在下文当中被使用)
* loadrunner 与 jmeter 关联的不同之处
* 在loadrunner中关联函数是写在要获取变量值的页面的前面
* 在jmeter中关联函数是要写在要获取变量值的页面的后面
* 在loadrunner中关联函数是注册函数］
* 在jmeter中使用正则表达式提取器来进行关联

## Jmeter实现配置管理
* 配置元件 用来提供一些配置相关的信息　如HTTP请求头 COOKIE管理　提供参数化管理　还可以进行用户自定义变量等
* Add ---> Config Element ---> JDBC Connection Configuration 
* Variable Name for created pool: 变量名必填
* Database Connection Configuration
* Database URL: jdbc:mysql://localhost:3306/mysql?serverTimezone=UTC&useSSL=true
* (注意 useSSL=true) 解决控制警告提示
* JDBC Driver class: com.mysql.jdbc.Driver
* Username:root
* Password:root