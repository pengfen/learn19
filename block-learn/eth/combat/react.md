# react 开发

* npm install -g create-react-app
* sudo ln -s /home/ricky/app/node10/bin/create-react-app /usr/bin/create-react-app
* create-react-app lottery-react
* npm run start 运行项目
* npm install --save semantic-ui-react 安装ui库

## react 生命周期
* Component renders
* componentDidMount 被调用
* 准备数据 调用call获取区块链的数据
* 修改数据(setState) ui自动更新

## 代码
```bash
pragma solidity ^0.4.17;

contract Lottery {
    address public manager;
    address[] public players;
    address public winner;

    // 构造函数 注意版本 高版本使用 constructor() {}
    function Lottery() public {
        manager = msg.sender;
    }

    // 管理者 ---> 部署合约账号
    function getManager() public view returns (address) {
        return manager;
    }

    // 投注彩票
    function enter() public payable {
        require(msg.value == 1 ether);
        players.push(msg.sender);
    }

    // 返回所有的投注彩票的人
    function getAllPlayers() public view returns (address[]) {
        return players;
    }

    // 合约余额
    function getBalance() public view returns(uint) {
        return this.balance;
    }

    // 玩家数量
    function getPlayersCount() public view returns(uint) {
        return players.length;
    }

    // 随机种子
    function random() private view returns (uint) {
        return uint(keccak256(block.difficulty, now, players));
    }

    // 中奖者
    function pickWinner() public onlyManagerCanCall{
        uint index = random() % players.length;
        //address winner = players[index];
        winner = players[index];
        players = new address[](0); // 重复开奖
        winner.transfer(this.balance);
        return winner;
    }

    // 用户退款
    function refund() public onlyManagerCanCall {
        for(uint i = 0; i < players.length; i ++) {
            players[i].transfer(1 ether);
        }
        players = new address[](0);
    }

    modifier onlyManagerCanCall() {
        require(msg.sender == manager);
        _;
    }
}
```

* 编译
```bash
// 编译的智能合约的脚本
const path = require('path');
const fs = require('fs');
const solc = require('solc'); // 注意版本 npm install --save solc@0.4.17

const srcpath = path.resolve(__dirname, 'contracts', 'Lottery.sol');
const source = fs.readFileSync(srcpath, 'utf-8');

const result = solc.compile(source, 1);
module.exports = result.contracts[':Lottery'];
```

* 部署
```bash
// 部署智能合约
const Web3 = require('web3');
const {interface, bytecode} = require('./compile');
// npm install --save truffle-hdwallet-provider@0.0.3
const HDWalletProvider = require("truffle-hdwallet-provider");
const mnemonic = "dutch stone surface matrix desk odor slush another local aspect two screen";
const provider = new HDWalletProvider(mnemonic, "https://rinkeby.infura.io/v3/c19bba97105c4d88ae2d934040860742");
const web3 = new Web3(provider);

deploy = async ()=>{
    const accounts = await web3.eth.getAccounts();

    const result = await new web3.eth.Contract(JSON.parse(interface))
        .deploy({data:bytecode})
        .send({
            from:accounts[0],
            gas:'3000000'
        });
    console.log('address: ' + result.options.address);
    console.log('----------------------');
    console.log(interface);
};

deploy();
```

* web3
```bash
// 使用web3 v0.2 的provider注入到web3 v1.0的provider

import Web3 from 'web3';
const web3 = new Web3(window.web3.currentProvider);
export default web3;
```

* 调用智能合约
```bash
// 获取rinkeby区块链智能合约

import web3 from './web3';

const address = '0x5bCf16eb0ac46E42864E9D037b41e13b6d79E58D';

// ctrl + shift + J 多行转一行
const abi = [{
    "constant": true,
    "inputs": [],
    "name": "getBalance",
    "outputs": [{"name": "", "type": "uint256"}],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
}, {
    "constant": true,
    "inputs": [],
    "name": "manager",
    "outputs": [{"name": "", "type": "address"}],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
}, {
    "constant": false,
    "inputs": [],
    "name": "refund",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
}, {
    "constant": false,
    "inputs": [],
    "name": "pickWinner",
    "outputs": [{"name": "", "type": "address"}],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
}, {
    "constant": true,
    "inputs": [],
    "name": "getPlayersCount",
    "outputs": [{"name": "", "type": "uint256"}],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
}, {
    "constant": true,
    "inputs": [],
    "name": "getManager",
    "outputs": [{"name": "", "type": "address"}],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
}, {
    "constant": false,
    "inputs": [],
    "name": "enter",
    "outputs": [],
    "payable": true,
    "stateMutability": "payable",
    "type": "function"
}, {
    "constant": true,
    "inputs": [],
    "name": "getAllPlayers",
    "outputs": [{"name": "", "type": "address[]"}],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
}, {
    "constant": true,
    "inputs": [{"name": "", "type": "uint256"}],
    "name": "players",
    "outputs": [{"name": "", "type": "address"}],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
}, {"inputs": [], "payable": false, "stateMutability": "nonpayable", "type": "constructor"}];


const lottery = new web3.eth.Contract(abi, address);

export default lottery;
```

* js 操作
```bash

// constructor(props) {
//     super(props);
//     this.state = {manager: ''};
// }

state = {
    manager: '',
    playersCount:0, // 玩家数量
    balance:0, // 合约余额
    loading:false, // 默认不loading
    showbutton: 'none'
};

async componentDidMount() {
    //console.log(lottery);
    const address = await lottery.methods.getManager().call();
    console.log('address: ' + address);
    this.setState({manager: address})
    // const playersCount = await lottery.methods.getPlayersCount().call();
    // this.setState({playersCount:playersCount});
    // const balance = await lottery.methods.getBalance().call();
    // console.log(balance);
    // this.setState({balance: balance});
    //this.setState({balance: web3.utils.fromWei(balance, 'ether')});

    const accounts = await web3.eth.getAccounts();
    if (accounts[0] === address) {
        this.setState({showbutton:'inline'})
    }
}

// enter = async ()=>{
//     this.setState({loading:true});
//     // 获取账户
//     const accounts = await web3.eth.getAccounts();
//     //
//     await lottery.methods.enter().send({
//         from:accounts[0],
//         value:'1000000000000000000'
//     });
//     this.setState({loading:false});
// };

pickWinner = async ()=>{
    this.setState({loading:true});
    // 获取当前账户
    const accounts = await web3.eth.getAccounts();
    await lottery.methods.pickWinner().send({
        from:accounts[0]
    });
    this.setState({loading:false})
    window.location.reload(true);
}
```