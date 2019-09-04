# nodejsjs 介绍

## node.js
1. nodejs 是一个开发平台
2. 使用编程语言 JS
3. nodejs 平台是基于 chrome V8 JavaScript 引擎构建
4. 基于 nodejs 可以开发控制台程序(命令行程序 CLI程序) 桌面应用程序(GUI) Web应用程序

## 特点
1. 事件驱动(当事件被触发时 执行传递过去的回调函数)
2. 非阻塞 I/O 模型(当执行I/O操作时 不会阻塞线程)
3. 单线程
4. 拥有最大的开源库生态系统 npm

require 加载需要的node.js原生模块
var http = require('http'); // 服务器创建
dns = require('dns'); // DNS 查询
fs = require('fs'); // 文件操作
url = require('url'); // url 处理
querystring = require('querystring'); // 字符串处理

http 模块负责HTTP服务器的创建
dns 模块主要负责解析当前DNS域名 返回DNS服务器IP地址
url 模块处理文件url路径
querystring 模块处理前端传回的字符串解析