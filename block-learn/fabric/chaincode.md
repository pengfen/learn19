# 链码开发

## 判断$GOPATH是否存在
* echo $GOPATH  如果输出为空 则需要配置$GOPATH
* vi ~/.bashrc 添加
* export GOPATH="安装目录/go" (可通过 which go ---> ls -la /usr/bin/go)
* export PATH="$PATH:$GOPATH/bin"
* source .bashrc

## 链码
* 新建链码文件  mkdir -p $GOPATH/src/sacc && cd $GOPATH/src/sacc
* touch sacc.go

## 添加链码代码
```bash
package main
import (
    "fmt"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/hyperledger/fabric/protos/peer"
)

type SimpleAsset struct {
}

func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
    args := stub.GetStringArgs()
    if len(arg) != 2 {
        return shim.Error("Incorrect arguments. Expecting a key and a value")
    }
    //
    err := stub.PutState(args[0], []byte(args[1]))
    if erf != nil {
        return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
    }
    return shim.Success(nil)
}
//
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    fn, args := stub.Get FunctionAndParameters();
    var result string
    var err error
    if fn == "set" {
        result, err = set(stub, args)
    } else {
        result, err = get(stub, args)
    }
    if err != nil {
        return shim.Error(arr.Error())
    }
    return shim.Success([]byte(result))
}
//
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    if len(args) != 2 {
        return "", fmt.Error("Incorrect argguments, Expecting a key and a value")
    }
    err := stub.PutState(args[0], []byte(args[1]))
    if err != nil {
        return "", fmt.Errorf("Failed to set asset: %s", args[0])
    }
    return args[1], nil
}
//
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    if len(args) != 1 {
        return "", fmt.Errorf("Incorrect arguments, Expecting a key")
    }
    value, err := stub.GetState(args[0])
    if err != nil {
        return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
    }
    if value == nil {
        return "", fmt.Errorf("Asset not found: %s", args[0])
    }
    return string(value), nil
}
//
func main() {
    if err := shim.Start(new(SimpleAsset)); err != nil {
        fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
    }
}
```

## 编译链码
go get -u --tags nopkcs11 github.com/hy...
go build --tags nopkcs11

## 部署
* cd chaincode-docker-devmode
* 打开三个终端
* 终端一 开启网络  sudo docker-composer -f docker-compose-simple.yaml up
* 终端二 编译链码 运行链码  sudo docker exec -it chaincode bash
* cd sacc   ---> go build  
* 运行链码 CORE_PEER_ADDRESS=peer:7051 CORE_CHAINCODE_ID_NAME=mycc:0 ./sacc
* 终端三 使用链码
* sudo docker exec -it cli bash
* peer chaincode install -p chaincodedev/chaincode/sacc -m mycc -v 0
* peer chaincode instantiate -n mycc -v 0 -c '{"Args":["a","10"]}' -C myc
* 修改a的值为20
* peer chaincode invoke -n mycc -c '{"Args":["set","a","20"]}' -C myc
* 查询a的值
* peer chaincode query -n mycc -c '{"Args":["query","a"]}' -C myc