# 链码 开发环境

## 链码存放
* 将写好的链码文件复制到 fabric-samples/chaincode

## 开启3个终端
* 终端一 启动网络
* cd chaincode-docker-devmode
* sudo docker-compose -f docker-compose-simple.yaml up
* docker启动出错 sudo docker rm -f $(sudo docker ps -aq)
* 终端二 编译且启动链码
* docker exec -it chaincode bash
* cd chaincode002
* go build
* CORE_PEER_ADDRESS=peer:7051 CORE_CHAINCODE_ID_NAME=mycc:0 ./chaincode002
* 终端三 操作链码
* sudo docker exec -it cli bash
* 安装链码 peer chaincode install -p chaincodedev/chaincode/chaincode002 -n mycc -v 0
* 实例化链码 peer chaincode instantiate -n mycc -v 0 -c '{"Args":["str","helloworld"]}' -C myc
* 查询链码 peer chaincode query -n mycc -c '{"Args":["get","str"]}' -C myc
* 修改链码 peer chaincode invoke -n mycc -c '{"Args":["set","str","wocme"]}' -C myc
* 查询链码 peer chaincode query -n mycc -c '{"Args":["get","str"]}' -C myc