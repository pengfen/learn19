# 链码介绍

## 什么是链码
* chaincode 链码 一般是用户使用go语言编写的应用代码
* 链码被部署在 Fabric 网络节点上 运行在 Docker 容器中 并通过gRPC协议与相应的Peer节点进行交互 以操作分布式账本中的数据

## 链码生命周期
* Fabric提供了 package install instantiate 和 upgrade 四个命令管理链码的生命周期
* 通过 install 安装链码 通过 instantiate 实例化链码 然后可以通过 invoke query 调用链码和查询链码
* 如果需要升级链码 则需要先 install 安装新版本的链码 通过 upgrade 升级链码
* 在 install 安装链码前 可能通过 package 打包并签名生成打包文件 然后在通过 install 安装

## 链码相关操作
* 安装链码 peer chaincode install -n sacc -v 1.0 -p sacc
* 实例化链码 peer chaincode instantiate -n sacc -v 1.0 -c '{"Args":["join","0"]}' -P "OR ('Org1.')"
* 调用链码 peer chaincode invoke -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile
* 查询链码 peer chaincode query -C mychannel -n mycc -c '{"Args":["query","a"]}'
* 升级链码
* 安装新版本的链码 peer chaincode install -n mycc -v 1 -p path/to/my/chaincode/v1
* upgrade 升级链码 peer chaincode upgrade -n mycc -v 1 -c '{"Args":["d","e","f"]}' -C mychannel
* 打包链码和签名
* 打包链码 peer chaincode package -n mycc -p github.om/...
* 对打包文件进行签名 peer chaincode singpackage ccpack.out singedccpack.out