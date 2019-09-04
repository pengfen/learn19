## nodejs 安装和配置

1. 下载地址
[当前版本](https://nodejs.org/en/download)

2. 版本
LTS 版本 Long-term Support 版本  长期支持版(稳定版)
Current 版本  Latest Features 版本  最新版本 新特性会在该版本中最先加入

3. 使用 node =v 确定是否安装成功

## REPL 介绍
1. REPL 全称  Read-Eval-Print-Loop (交互式解释器)
* R 读取 - 读取用户输入 解析输入了 JavaScript 数据结构并存储在内存中
* E 执行 - 执行输入的数据结构
* P 打印 - 输出结果
* L 循环 - 循环操作以上步骤直到用户两次按下 ctrl+c 按钮退出

2. 在REPL中编写程序 node命令进入REPL环境


https://nodejs.org/en/download/ 下载node.js

安装 (点击下一步)

node.js command prompt # 进入到nodejs命令窗口

在 F:\nodejs目录下编写js文件
var http = require("http"); 
http.createServer(function(request, response) { 
response.writeHead(200, {"Content-Type": "text/plain"}); 
response.write("test nodjs"); 
response.end(); 
}).listen(8899); 
console.log("nodejs start listen 8899 port!");

C:\Users\Administrator>f:

F:\>cd nodejs

F:\nodejs>node test.js
nodejs start listen 8899 port!


linux安装nodejs
1.下载nodejs
https://nodejs.org/en/download/

2. 解压
tar -xvf node-v8... -C ~/app/

3. 使用软链接生成命令
sudo ln -s /home/ricky/app/node-v8.../bin/node /usr/bin/node
sudo ln -s /home/ricky/app/node-v8.../bin/npm /usr/bin/npm

node -v
npm -v