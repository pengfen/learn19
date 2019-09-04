# web3的常见api

## web3.version

## web3.currentProvider
```bash
const assert = require("assert");
const ganache = require("ganache-cli");
// 约定的规范 如果变量以大写字母开头 就是一个构造方法(构造函数)
const Web3 = require("web3");
const web3 = new Web3(ganache.provider()); // 把ganache测试网络的卡插入到web3里面

describe("测试智能合约", ()=>{
    it("测试web3的版本", ()=>{
        console.log(web3.version);
    });
    it("测试web3的网络", ()=>{
        console.log(web3.currentProvider);
    });
});
```

## web3.utils.toHex 转十六进制数字
```bash
var str = "abcABC";
var obj = {abc: 'ABC'};
var bignumber = new BigNumber('12345678901234567890');

var hstr = web3.utils.toHex(str); // 0x616263414243
var hobj = web3.utils.toHex(obj); // 0x7b22616263223a22414243227d
var hbg = web3.toHex(bignumber);
```

## web3.utils.toAscii 转Ascii码
```bash
var str = web3.utils.toAscii("0x657468657265756d000000000000000000000000000000000000000000000000"); // ethereum
console.log(str);
```

## web3.utils.fromAscii 转十六进制数
```bash
var str = web3.utils.fromAscii('ethereum');
console.log(str); // "0x657468657265756d"

var str2 = web3.utils.fromAscii('ethereum', 32);
console.log(str2); // "0x657468657265756d000000000000000000000000000000000000000000000000"
```

## web3.utils.fromWei
```bash
var value = web3.fromWei('21000000000000', 'finney');
console.log(value); // "0.021"
```

## web3.utils.toWei
```bash
var value = web3.toWei('1', 'ether');
console.log(value); // "1000000000000000000"
```

## web3.eth.getAccounts 获取测试网中十个账户
```bash
// then异步获取数据后执行
web3.eth.getAccounts().then((accounts)=>{
	console.log(accounts);
});
[ '0xcCFC7C9546F1804a8eb2b4f869aD4a4B5006A266',
  '0xe9FA8FDbe0607984F911b69d13618f16656A9A65',
  '0xE4a09C0133cDf5c7519d8e559dd1fD2EA8E22244',
  '0x8ab0c1D5b3d5fb49dC8B6Fa98C6fbEFCF2B126b2',
  '0xAeEEc66559C3688d3767aAA20e5a701B55Fdc53a',
  '0x5EeD0a281EB93e83eaB3559F12812Ff49243aAE9',
  '0xe787631D3C9F6d1fE59d6Ecf08969Ff74D182251',
  '0x3b331537e32dE00f8e0890c734ac89B7956b4b1b',
  '0xFf270A57CCfAd08d10537b868ae8B5c1Fe40E545',
  '0xe119Ab91ba935a544D4e67246A29538023e1Bf0F' ]

```

## web3.eth.getBalance 获取账号中余额
```bash
var balance = web3.eth.getBalance("0xcCFC7C9546F1804a8eb2b4f869aD4a4B5006A266");
console.log(balance); // instanceof BigNumber
console.log(balance.toString(10)); // '1000000000000'
console.log(balance.toNumber()); // 1000000000000
```

## es6的async和await
```bash
it("测试web3的api", async ()=>{
    const accounts = await web3.eth.getAccounts();
    console.log(accounts);
    const money = await web3.eth.getBalance(accounts[0]);
    console.log(web3.utils.fromWei(money, 'ether'));
});
```

## 测试智能合约
```bash
const assert = require("assert");
const ganache = require("ganache-cli");
// 约定的规范 如果变量以大写字母开头 就是一个构造方法(构造函数)
const {interface,bytecode} = require("../compile")
const Web3 = require("web3");
const web3 = new Web3(ganache.provider()); // 把ganache测试网络的卡插入到web3里面

describe("测试智能合约", ()=>{
    it("测试web3的版本", ()=>{
        console.log(web3.version);
    });

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
