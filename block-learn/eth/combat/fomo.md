# 区块链彩票项目 
* 字段
1. manager 创建智能合约的人的地址
2. players 数组 记录所有的参与彩票投注的人
* 函数
1. manager的开奖函数 pickWinner
2. player的投注函数 enter

```
pragma solidity ^0.4.17;

contract Lottery {
    address public manager;
    address[] public players;
    
    function Lottery() public {
        manager = msg.sender;
    }
    
    // deploy ---> getManager, manager 合约地址
    function getManager() public view returns (address) {
        return manager;
    }
}
```

## 项目 ---> 创建项目lottery
* npm init
* npm install --save web3 ganache-cli mocha
