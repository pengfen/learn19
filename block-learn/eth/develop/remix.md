# remix 开发 solidity

https://remix.ethereum.org

# 选择编译版本
* compile ---> select new compiler version ---> 勾选Auto compile
* 注意编译版本与代码版本一致 param solidity ^0.4.25;

# 代码
```bash
pragma solidity ^0.4.25; // 声明编译器的版本 版本标识符 与nodejs版本标记一致

// Run ---> Environment 选择 JavaScript VM ---> Deploy ---> setMessage 设置信息 ---> getMessage 获取
contract Wel { // 类似java class声明 Wel智能合约
    string public message; // string数据类型 message类的成员变量 在整个智能合约生命周期都可以访问 pulib 是访问修饰符 是storage类型的变量 成员变量和局部变量
    
    // 函数以function开头
    function setMessage(string newMessage) public {
        message = newMessage;
    }
    
    function getMessage() public view returns(string) {
        return message;
    }
    // getMessage 函数名
    // public view 函数类型
    // returns (string) 返回值

    // public 公有 任何人(拥有以及坊账户的)都可以调用
    // private 私有 只有智能合约内部可以使用
    // view/constant 函数不会修改任何contract的数据
    // pure 函数不使用任何智能合约的变量
    // payable 调用函数需要付钱 钱付给了智能合约的账户
    // returns 返回值 函数声明中使用
}
```

## setMessage getMessage
* setMessage
* does change blockchain 需要成本
* 函数调用 send transaction
* 函数不能返回数据 因为函数需要花时间执行
* 需要十几秒才能执行完毕 返回值是transaction的hash
* 需要花费

* getMessage
* doesn't change blockchain 不需要成本
* 函数调用call
* 不修改智能合约的数据
* 函数可以返回数据
* 立刻执行
* 免费

# 搭建本地solidity开发环境
* 安装 nodejs
* 安装 remix-ide  npm install remix-ide -g
* ln -s /home/ricky/app/node10/bin/remix-ide /usr/bin/remix-ide
* remix-ide
* http://localhost:8080

# remix工作原理
* 虚拟的以太坊网络 javascript vm
* 控制台 编译节点 run节点 environment javascript vm
* account 虚拟账户5个
* 创建智能合约 智能合约在网络上的地址
* 调用getMessage setMessage
  > set的红色(send) 发送transaction 请求  至少花费十几秒
  > get的蓝色(call) 调用 瞬间(免费)

* injected web3 ---> 会调用 metamask 中的 Rinkeby网

## 获取更多的 rinkeby网以太币用于测试
* https://faucet.rinkeby.io/
* 转发账户地址
* 信息地址填入rinkeby,io的输入框中