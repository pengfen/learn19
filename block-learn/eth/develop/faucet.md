# 开发以太币水龙头

* npm install express --save

```bash
const express = require("express");
const app = express();

// rinkeby 环境测试
const Web3 = require('web3');
// npm install --save truffle-hdwallet-provider@0.0.3
const HDWalletProvider = require("truffle-hdwallet-provider");
const mnemonic = "dutch stone surface matrix desk odor slush another local aspect two screen";
const provider = new HDWalletProvider(mnemonic, "https://rinkeby.infura.io/v3/c19bba97105c4d88ae2d934040860742");
const web3 = new Web3(provider);

app.get('/send/:address', async function (req, res) {
    const address = req.params.address;
    console.log(address);

    const accounts = await web3.eth.getAccounts();
    const trans = await web3.eth.sendTransaction({
        from:accounts[0],
        to:address,
        value: '10000000000000000'
    });
    res.send("转账成功: " + trans.id);
})

const server = app.listen(3000, function() {
    var host = server.address().address;
    var port = server.address().port;

    console.log("Example app listening at http://%s:%s", host, port);
});
```

* node webServer
* http://localhost:3000/send/0x9EE4Efd27979399F40C4bfE25f9E575E88F13cdE
* 查看账户余额变化