# 代币合约

## 创建项目
* mkdir token ---> cd token ---> truffle init ---> atom ./

* 编写代码

```bash
pragma solidity ^0.4.4;

contract Token {
  uint256 INITIAL_SUPPLY = 666666; // 代币总量
  mapping (address => uint256) balances;

  function Token() {
    balances[msg.sender] = INITIAL_SUPPLY;
  }

  // 转账到一个指定的地址
  function transfer(address _to, uint256 _amount) {
    assert(balances[msg.sender]) > _amount);
    balances[msg.sender] -= _amount;
    balances[_to] += _amount;
  }

  // 查看指定地址的余额
  function balanceOf(address _owner) constant returns (uint256) {
    return balances[_owner];
  }
}
```

* 部署运行
> truffle develop ---> compile ---> 编写部署代码

```bash
var Token = artifacts.require("./Token.sol");

module.exports = function(deployer) {
  deployer.deploy(Token);
};
```

> migrate ---> let contract; ---> Token.deployed().then(instance => contract = instance) ---> contract.balanceOf(""); 注意地址默认是第一个 0x627306090abab3a6e1400e9345bc60c78a8bef57
truffle(develop)> contract.balanceOf("0x627306090abab3a6e1400e9345bc60c78a8bef57");
BigNumber { s: 1, e: 5, c: [ 666666 ] }
> contract.transfer("0xf17f52151ebef6c7334fad080c5704d77216b732");
> contract.balanceOf("0xf17f52151ebef6c7334fad080c5704d77216b732");
> contract.balanceOf("0x627306090abab3a6e1400e9345bc60c78a8bef57");
