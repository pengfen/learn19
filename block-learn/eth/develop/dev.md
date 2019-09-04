# 智能合约完整开发流程

1. 创建项目 ---> npm init ---> 删除不相关文件 ---> 创建相关文件，目录
* 创建目录 contracts, test
* 创建文件 compile.js deploy.js

2. 安装相关依赖
* npm install --save solc
* npm install --save mocha 单元测试
* npm install --save ganache-cli 本地测试网络
* npm install --save web3 
* npm install --save truffle-hdwallet-provider@0.0.3

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

* 编译代码 complie.js
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

* 本地测试代码
```bash
const assert = require("assert");
const ganache = require("ganache-cli");
// 约定的规范 如果变量以大写字母开头 就是一个构造方法(构造函数)
// 编译
const {interface,bytecode} = require("../compile")
const Web3 = require("web3");
const web3 = new Web3(ganache.provider()); // 把ganache测试网络的卡插入到web3里面

describe("测试智能合约", ()=>{
    it("测试部署智能合约", async ()=>{
        const accounts = await web3.eth.getAccounts();
        const result = await new web3.eth.Contract(JSON.parse(interface))
            .deploy({
                data:bytecode,
                arguments:['abc']
            }).send({
                from:accounts[0],
                gas:1000000
            });
        console.log("address:" + result.options.address); // address:0x43d6F0Ef827bd53DAf0Be0755377058FE3bb75EB

        let message = await result.methods.getMessage().call();
        console.log(message); // abc
        assert.equal(message, 'abc');

        await result.methods.setMessage("ricky").send({
            from:accounts[0],
            gas:1000000
        });

        message = await result.methods.getMessage().call();
        console.log(message); // ricky
        assert.equal(message, 'ricky');
    });
});
```

* 部署合约到renkeby网络
```bash
// 部署智能合约到rankeby网络
const Web3 = require('web3');
const {interface, bytecode} = require('./compile');
// npm install --save truffle-hdwallet-provider@0.0.3
var HDWalletProvider = require("truffle-hdwallet-provider");
var mnemonic = "crouch fiction income edge cluster turtle plastic ozone mom predict goddess express"; // 12 word mnemonic
var provider = new HDWalletProvider(mnemonic, "https://rinkeby.infura.io/v3/c19bba97105c4d88ae2d934040860742");
const web3 = new Web3(provider);

const deploy = async()=>{
    //const accounts = await web3.eth.getAccounts();
    //console.log(accounts[0]);
    const result = await new web3.eth.Contract(JSON.parse(interface))
        .deploy({
            data:bytecode,
            arguments:['abc']
        }).send({
            //from:accounts[0],
            from:0x8dd4d248fa1ba7bc12a584791c1bdd8559bd9841,
            gas:'1000000'
        });
    console.log("address: " + result.options.address);
}
deploy();
```

* https://etherscan.io/ 