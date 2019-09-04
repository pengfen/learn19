基础类型

string 字符串 'welcome'
bool 布尔 true false
int 整型 300, 0, -200
uint 非负整数 30
fixed/unfixed 小数 3.14 -3.14数据
address 地址

## 代码
```bash
pragma solidity ^0.4.17; # 注意编译也选择这版本

contract First {
    uint data;
    
    function setData(uint x) public{
        data = x;
    }
    
    function getData() constant public returns (uint) {
        return data;
    }
}
```

* 浏览器开发 http://localhost:8080
* Run ---> deploy ---> 0x5a12c4cf557ff4a6fa47bc737669c2a310b60f68 (https://etherscan.io/ ---> More ---> Rinkeby 查看合约相关交易信息)
* setData 设置值(红色需要花费)   getData 取值(蓝色不需要花费)

* idea开发 (ganache环境)
* npm install --save ganache-cli mocha web3 solc@0.4.17
* 编译文件 complie.js
* 测试用例 First.test.js
* 部署文件 deploy.js
