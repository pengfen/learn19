# idea 开发智能合约

## nodejs 开发环境搭建
* nodejs 版本大于8 下载地址 https://nodejs.org/en/download
* 注意nodejs版本 node10时使用idea创建项目会出现 please specify npm or yarn package
* 配置npm镜像仓库 npm --registry https://registry.npm.taobao.org info underscore
* 编辑~/.npmrc  registry = https://registry.npm.taobao.org

## 智能合约编译
* 源代码 ---> solidity编译器 ---> abi/bytecode ---> 部署到某个网络
* truffle(智能合约创建 本地测试 部署) ---> rinkeby(网上大多都是基于truffle)
> truffle问题 1. api不稳定 2. bug比较多 3. 一些功能缺失
* 源代码类型   字节码类型      执行环境     调用者
* .java源文件 .class字节码   jvm执行      java代码调用
* .sol源文件  bytecode字节码 区块链环境执行 javascript代码调用

## solidity开发环境搭建
1. solidity编译器,编译环境
2. mocha 抹茶测试环境
3. 部署智能合约的脚本 到指定网络
4. npm init

* 使用idea ---> 安装idea ---> 安装插件nodejs intellij-solidity ---> 创建项目 ---> node.js and NPM ---> 选择node(注意版本) --> 删除不必要的文件 ---> 创建相关文件
* contracts 合约存放目录
* test 合约测试目录
* package.js 
* compile.js 编译
* deploy.js 部署

* 编写合约代码
```bash
pragma solidity ^0.4.17;

contract Welcome {
    string public message;
    function Welcome(string _message) public{
        message = _message;
    }
    function setMessage(string _message) public {
        message = _message;
    }
    function getMessage() public view returns(string) {
        return message;
    }
}
```

* 编译 cd wel 进入项目目录 ---> npm init ---> 安装solc(npm install --save solc)
* 编写编译代码 complie.js
```bash
// 编译智能合约的脚本
const path = require('path');
const fs = require("fs");
const solc = require("solc");

const srcpath = path.resolve(__dirname, "contracts","Welcome.sol");
// console.log(srcpath); // /home/ricky/IdeaProjects/wel/contracts/Welcome.sol
const source = fs.readFileSync(srcpath, 'utf-8'); // 源代码
const result = solc.compile(source, 1);
//console.log(result);
module.exports = result.contracts[':Welcome'];
```