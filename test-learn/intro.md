# 测试

## 性能测试的常规流程
* 确定需求并分析业务场景
* 设计测试方案(模板)　测试脚本
* 准备测试数据并执行测试脚本
* 监控资源消耗
* 瓶颈定位和性能调优验证
* 输出报告(模板)和跟踪

## 流程通常不能如你所愿的规范
* 测试环境和实际环境的配置不一致
* 没有明确目标的需求

## 测试技能
* 开发语言
* 操作系统
* 数据库
* 测试工具
* 网络知识
* 业务知识

## 性能测试的目的
* 发现性能瓶颈

## 性能测试的分类
* 性能测试是一个非常广泛的概念　包括很多方面的测试　也可称之为非功能测试
* 自动化测试属于功能测试的范围　由于其测试方法要求测试人员拥有一定的代码能力　所以被单独分成一个测试模块
* 测试范围
* 负载测试　通过逐步加压的方法　达到既定的性能阈值的目标　阈值的设定应是小于等于某个值　如cpu使用率小于等于80%
* 压力测试　通过逐步加压的方法　使得系统的某些资源达到饱和　甚至失效的状态　简单粗暴的解释就是什么条件能把系统压崩溃
* 并发测试　在同一时间内　多个虚拟用户同时访问同一模块　同一功能　通常的测试方法是设置集合点
* 容量测试　通常是指数据库层面的　目标是获取数据库的最佳容量的能力　又称之为容量预估　具体测试方法为在一定的并发用户　不同的基础数据量下　观察数据库的处理能力　即获取数据库的各项性能指标
* 可靠性测试　又称之为稳定性测试或疲劳测试　是指系统在高压情况下　长时间的运行系统是否稳定　如cpu使用率在80%以上 7X24小时运行　系统是否稳定
* 异常测试　又称之为失败测试　是指系统架构方面的测试　如在负载均衡架构中　要测试岩机　节点挂掉等情况系统的反映


## 性能测试的工作流程
* 需求分析 ---> 性能指标制定 ---> 脚本开发 ---> 场景设置 ---> 监控部署 ---> 测试执行 ---> 性能分析 ---> 性能调优 ---> 测试报告
* 性能调优 ---> 测试执行　---> 性能分析 ---> 性能调优　迭代进行

## 系统常见分层
* 显示层(View)  web android ios H5
* 逻辑控制层(Controller) Api
* 数据存储层(Model) mysql mongodb redis ...

## 性能指标
* 事务 　从客户端发起的一个或多个请求(这些请求组成一个完整的操作) 到客户端接收到从服务器返回的响应
* TPS(Transaction PerSecond) 每秒系统处理的事务数
* 请求响应时间  从客户端发起的一个请求开始 到客户端接收到从服务器返回的响应　整个过程所耗费的时间
* 事务响应时间  事务可能是由一个或多个请求组成的 事务响应时间主要是针对于用户的角度而言 如转账
* 并发 没有严格意义上的并发　并发总有先后　无论差距是


## 并发
* 没有严格意义上的并发 并发总有先后　无论差距是1毫秒或者是1微秒 总有一个时间差　所以并发讲的是一个时间范围内　比如1秒内
* 多用户在系统上进行同一操作　比如双十一时　大家都针对同一种商品进行秒杀
* 多用户在系统上进行不同操作　比如双十一时　大家针对不同商品进行秒杀　或者是大家有进行其他不同的操作　比如商品浏览

## 并发用户数
* 同一单位时间内对系统发起请求的用户数量

## 吞吐量
* 一次性测试过程中网络上传输的数据量的总和

## 吞吐率
* 单位时间内网络上传输的数据量
* 吞吐率 = 吞吐量 / 传输时间

## 点击率
* 每秒钟用户向服务器提交的请求数　这个指标是web应用程序特有的一个指标　可以想象为每秒钟用户总共在页面上进行多少次点击动作　但是需要注意的是一次鼠标单击的操作后　客户端有可能向服务器发送了多次请求

## 资源使用率
* 对不同的系统资源的使用情况 如cpu 内存 io

## 需求分析
* 分析的目的
* 明确测试指标
* 明确测试场景

## 新系统
* 同行业比较
* 业务预期

## 老系统
* 对比以往的用户使用行为以及用户量

## 测试测试工具
* 常用工具  LoadRunner JMeter
* 对比纬度  LoadRunner JMeter
* 量级     重          轻
* 易用性   易          易
* 是否开源　否          是
* 语言支持  C/java1.5  java
* 是否收费  是         否