# 链码结构

## Chaincode 接口
* 每个链码都需要实现 Chaincode 接口
* init 当链码实例化或者升级的时候 init方法会被调用
* invoke 当链码收到调用(invoke)或者查询的时候 invoke会被调用
```bash
type Chaincode interface {
	Init(stub ChaincodeStubInterface) pb.Response
	Invoke(stub ChaincodeStubInterface) pb.Response
}
```

## 链码结构
```bash
package main

// 引入必要的依赖包
import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// 声明一个结构体
type SimpleAsset struct {
}

// 为结构体添加 init 方法
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
}

// 为结构体添加 Invoke 方法
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
}

// 主函数 链码实例化时调用
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
```

## 依赖包
* fmt 包含了 Printf Errorf 等标准函数
* "github.com/hyperledger/fabric/core/chaincode/shim" shim包提供了链码与账本交互的中间层 通过shim.ChaincodeStubInterface提供的方法读取和修改账本的状态
* "github.com/hyperledger/fabric/protos/peer" init和invoke方法需要返回的peer.Response

## 链码开发的API
* GetStringArgs 提取调用链码交易中的参数 以字符串数组(string[])的形式返回
* GetFunctionAndParameters 提取调用链码交易中的参数 其中第一个参数作为被调用的函数名称
* GetState 按照给定的key查询返回对遗址的byte数组
* PutState 将给定的key和value写入账本
* DelState 在账本中移除给定的key和对应的value记录

## 运行链码
* 进入 fabric-samples/chaincode-docker-devmode 启动三个终端
* 终端一 开启网络
* sudo docker-compose -f docker-compose-simple.yaml up
* 终端二 编译链码
* docker exec -it chaincode bash
* cd sacc ---> go build
* CORE_PEER_ADDRESS=peer:7051 CORE_CHAINCODE_ID_NAME=mycc:0 ./sacc
* 终端三 使用锭码
* sudo docker exec -it cli bash
* peer chaincode install -p chaincodedev/chaincode/sacc -n mycc -v 0
* 链码实例化 peer chaincode instantiate -n mycc -v 0 -c '{"Args":["a","10"]}' -C myc
* 修改a的值为20 peer chaincode invoke -n mycc -c '{"Args":["set","a","20"]}' -C myc
* 查询a的值 peer chaincode query -n mycc -c '{"Args":["query","a"]}' -C myc