# 与合约交互

## 进入 fabcar
cd fabric-samples/fabcar

## 准备环境
* 杀掉活跃的容器  sudo docker rm -f $(docker ps -aq)
* 清理缓存网络    sudo docker network prune
* 删除fabcar智能合约的底层链码图像 如果是第一次运行这个项目可以不执行
* sudo docker rmi dev-peer0.org1.example.com-fabcar-1.0-5c906e402ed29f20260ae42283216aa75549c571e2e380f3615826365d8268ba
* 安装客户端并启动网络 npm install ---> ./startFabric.sh node
* 注册管理员用户 node enrollAdmin.js
* 注册user1用户 node registerUser.js
* 查询账目 node query.js