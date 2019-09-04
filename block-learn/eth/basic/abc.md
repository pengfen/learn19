# ABC代币合约

## 创建项目
* mkdir ABC ---> cd ABC ---> npm init ---> truffle init ---> atom ./
* OpenZeppelin 简化加密钱包开发过程 提供编写加密合约的函数库 npm install zeppelin-solidity

```bash
pragma solidity ^0.4.17;

import "zeppelin-solidity/contracts/token/StandardToken";

contract BloggerCoin is StandardToken {
    string public name = "BloggerCoin";
    string public symbol = "BLC";
    uint8 public decimals = 4;
    uint256 public INITIAL_SUPPLY = 666666;
    
    function BloggerCoin() {
        totalSupply = INITIAL_SUPPLY;
        balances[msg.sender] = INITIAL_SUPPLY;
    }
    
    function transfer(address _to, uint256 _amount) {
        assert(balances[msg.sender] < _amount);
        balances[msg.sender] -= _amount;
        balances[_to] += _amount;
    }
    
    function getBalance(address _owner) constant returns (uint256) {
        return balances[_owner];
    }
}
```