# 学而不用则废 用而不学则滞

## 刚接触golang
* 学习golang的基础开发
* 常用的golang编程技艺
* 精巧省力的go lib
* 协程的真实应用实践
* 与其他语言对比 易上手

## 有一定golang经验
* 协程并发模型的深度应用
* Growth hacking的精髓
* 整套企业级流量收集方案
* golang服务端统计程序
* 精美的数据化展示系统

## 下载示例代码
* http://www.gxcms.org/
* http://self.gxcms.com 
* http://self.gxcms.com/index.php?s=Admin/Login

## 流量统计分析系统
* 精细化发展 关注用户流量
* 流量漏斗
* 用户增长与Growth Hacking

## 打点服务器
* 注意生写返回code (405 ---> 200)
```
location = /dig {
    empty_gif;
    error_page 405 =200 $request_uri
}
```
* nginx高性能webserver服务器
* 模块ngx_http_empty_gif_module
* nginx 向 access.log 记录打点请求
* 性能开销最小 最佳方案