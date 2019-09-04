# web3 转账

1. 创建项目 trans ---> 删除不必要的文件 目录
2. npm init
3. npm install --save web3 ganache-cli
4. 新建trans.js文件
5. 修改js版本 file ---> setings ---> language & ... ---> javaScrit ---> ECMAScript 6
6. 编写代码
```bash
const Web3 = require('web3');
const ganache = require('ganache-cli');
const web3 = new Web3(ganache.provider());

send = async ()=>{
    const accounts = await web3.eth.getAccounts();
    let account1 = await web3.eth.getBalance(accounts[0]);
    let account2 = await web3.eth.getBalance(accounts[1]);

    console.log(accounts[0] + 'account0' + account1 + 'wei');
    console.log(accounts[1] + 'account2' + account2 + 'wei');

    await web3.eth.sendTransaction({
        from:accounts[0],
        to:accounts[1],
        value: 1000000
    });
    account1 = await web3.eth.getBalance(accounts[0]);
    account2 = await web3.eth.getBalance(accounts[1]);

    console.log('account1: ' + account1 + "wei");
    console.log('account2: ' + account2 + 'wei');
}
send();
```
7. renkeby 网络测试
```bash
// rinkeby 环境测试
const Web3 = require('web3');
// npm install --save truffle-hdwallet-provider@0.0.3
const HDWalletProvider = require("truffle-hdwallet-provider");
const mnemonic = "dutch stone surface matrix desk odor slush another local aspect two "; // 12个单词 助记词
const provider = new HDWalletProvider(mnemonic, "https://rinkeby.infura.io/v3/c19bba97105c4d88ae2d934040860742");
const web3 = new Web3(provider);

send = async ()=>{
    const accounts = await web3.eth.getAccounts();
    //console.log(accounts[1]);
    let account_bal1 = await web3.eth.getBalance(accounts[0]);
    let addb = '0x9EE4Efd27979399F40C4bfE25f9E575E88F13cdE';
    let account_bal2 = await web3.eth.getBalance(addb);

    console.log("account_bal1: " + account_bal1 + 'wei');
    console.log("account_bal2: " + account_bal2 + 'wei');

    const str = 'welcome to solidity world';
    let data = Buffer.from(str).toString('hex');
    data = '0x' + data;

    await web3.eth.sendTransaction({
        from:accounts[0],
        to:addb,
        value: '3000000000000000',
        data: data
    });

    account_bal1 = await web3.eth.getBalance(accounts[0]);
    account_bal2 = await web3.eth.getBalance(addb);

    console.log("account_bal1: " + account_bal1 + 'wei');
    console.log("account_bal2: " + account_bal2 + 'wei');
};

send();
```