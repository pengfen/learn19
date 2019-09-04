# 交易 转账

## 转账流程
1. 点击submit
2. 地址通过get请求发给了服务器
3. web3.js 创造了一个转账(transcation)
* nonce
* to
* value
* gas
* gaslimit
* 密码学相关 有privatekey生成
4. 后台server发送transcation到rinkeby网络
5. 等待转账
6. 后台服务器把成功信息反馈给前端

## 转账信息
* nonce 汇款方提交了多少次转账
* to 钱要被汇入哪个地址
* value 要汇多少钱(单位 wei)
* gasPrice 手续费一个单位的汽油 汇款方打算给多少钱(单位 wei)
* startGas/gasLimit 汽油的最大使用量
* v r s 密码学相关数据 由私钥生成 相当于数字签名

## 等待流程
1. transcation提交给一个node
2. node可能在某个时间内收到很多个transcation
3. node对transcation排序 挖矿
4. node算好nonce就会广播给全网
5. 最先算出nonce的node会得到挖矿奖励