# Hyperledge Fabric Samples 建立第一个网络

## Generate Network Artifacts
* cd first-network
* sudo ./byfn.sh -m down  #关闭服务
* sudo ./byfn.sh -m generate

* 加密生成器 使用该cryptogen工具为各种网络实体生成加密材料（x509证书和签名密钥）。这些证书代表身份，它们允许在我们的实体进行通信和交易时进行签名/验证身份验证
* 配置事务生成器 该configtxgen工用于创建四个配置工件
* orderer genesis block
* 频道 configuration transaction
* 每个Peer Org 两个-一个 anchor peer transactions

* 生成初始区块
* ../bin/cryptogen generate --config=./crypto-config.yaml
* export FABRIC_CFG_PATH=$PWD
* sudo chmod -R 777 channel-artifacts
* ../bin/configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
* 生成应用通道的配置信息
* export CHANNEL_NAME=mychannel
* ../bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME
* 生成锚节点配置更新文件
* ../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
* ../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP
* vim docker-compose-cli.yaml
* # command: /bin/bash ...
* CHANNEL_NAME=$CHANNEL_NAME TIMEOUT=600 sudo docker-compose -f docker-compose-cli.yaml up -d

## 创建和加入通道
* 进入docker容器  sudo docker exec -it cli bash
* 创建通道  export CHANNEL_NAME=mychannel
* peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls $CORE_PEER_TLS_ENABLED --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
* 加入通道 peer channel join -b mychannel.block
* 安装链码 peer chaincode install -n mycc -v 1.0 -p github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02
* 实例化   peer chaincode instantiate -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc -v 1.0 -c '{"Args":["init","a", "100", "b","200"]}' -P "OR('Org1MSP.member','Org2MSP.member')“
* 查询 peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a"]}'
* 转账 peer chaincode invoke -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem	-C $CHANNEL_NAME -n mycc -c '{"Args":["invoke","a","b","10"]}’
* 查看A的余额 peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a"]}'