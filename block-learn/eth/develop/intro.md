# Solidity介绍
* Solidity 是一种语法与javascript相似的高级语言 为Ethereum虚拟机(EVM)编译代码而设计
* Solidity 是静态类型的 支持继承 库和复杂的用户自定义类型以及其他功能
* 可用来创建投票 众筹 盲拍 多重签名钱包等的智能合约

## Solidity 环境安装

* 使用官网提供Wallet开发 注意部署经常失败
* testrpc + truffle 开发 注意需要nodejs
* 网页版 + MetaMask 开发

### [atom](https://atom.io/)
1. 安装atom (安装solidity插件)
file ---> settings ---> install ---> 搜索solidity 安装 autocomplete-solidity(代码补全) linter-solidity(代码错误检查)
安装 linter-solium language-ethereum(支持solidity代码高亮以及solidity代码片段)
2. 以管理员身份打开cmd (win10)
2. 安装node nodejs.cn下载node node -v
3. 安装节点仿真器 npm install -g ganache-cli
4. 安装solidity npm install -g solr
5. 安装web3 npm install -g web3
6. 安装truffle npm install -g truffle
7. 安装webpack npm install -g webpack
8. 启动testRPC ganache-cli

> 可使用命令安装
apm install linter
apm install linter-solidity
apm install linter-solium

### [remix](https://remix.ethereum.org/#optimize=true&version=soljson-v0.4.25+commit.59dbf8f1.js)

### 使用网页版 + MetaMask 开发

* 安装MetaMask
> 安装chrome插件MetaMask
> Accept
> 输入密码
> 复制账号的安全码 一共是12个单词 方便恢复账号

* 配置MetaMask的Test Net
> 从Main Ethereum Network 切换到 Ropsten Test Network
> 购买以太币

* 使用网页版Solidity编辑器 [browser-solidity](https://ethereum.github.io/browser-solidity)
> 编写代码
> 点击Create按钮后会弹出MetaMask界面

### 终端编写 solr命令

## hello world开发

### 创建项目
* cd project ---> mkdir wel --->  cd wel ---> truffle init

### 项目目录结构
* contracts truffle默认的合约文件存放地址
* migrations 存放发布的脚本文件
* test 存放测试应用和合约的测试文件
* truffle.js truffle-config.js truffle的配置文件

### 新建 wel 合约 (Wel.sol)
* cd contracts ---> truffle create contract Wel

```bash
pragma solidity ^0.4.4;

contract Wel {

  function sayWel() returns (string) {
    return ("welcome to solidity world");
  }
}
```

* 编译合约(注意在项目wel目录下) truffle compile