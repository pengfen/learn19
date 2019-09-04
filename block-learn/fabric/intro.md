# 介绍

* hyperledger-fabric.readthedocs.io fabric官网
* hyperledger.github.io介绍

## 环境安装
* 安装 cURL  curl --version   sudo apt-get install curl
* 安装 docker
* sudo apt-get install docker docker-engine ---> docker --version
* sudo apt-get install docker-compose  ---> docker-compose --version
* 安装go  sudo apt-get install golang-go  ---> go version
* 安装nvm  git clone https://github.com/creationix/nvm.git ~/.nvm && cd ~/.nvm && git checkout `git describe --abbrev=0 --tags`
* vim .bashrc ---> source ~/.nvm/nvm.sh
* nvm --version
* nvm use 使用指定版本node
* nvm install 6.9.5  安装指定版本node	

## Hyperledger Fabric Sample 下载安装
* 拉取源码
* git clone https://github.com/hyperledger/fabric-samples.git (注意版本问题)可手动下载指定版本 
* 坑(默认下地1.2) 安装第一个应用时报错 删除镜像(sudo docker images -q -a   sudo docker rmi -f 指定image)和相关目录 ---> 重新下载包 
* cd fabric-samples	
* 下载二进制文件 curl -sSL https://goo.gl/byy2Qj | bash -s 1.0.5 (会生成一个bin目录)
* 安装 docker images (https://goo.gl/byy2Qj 复制到init.sh文件中)
* 修改init.sh文件
* 所有 docker 前面加上 sudo
* 注释掉 echo "===> Downloading platform binaries" ...
* sudo chmod 777 init.sh
* ./init.sh

* 会安装相关镜像
* fabric-peer
* fabric-orderer
* fabric-couchdb
* fabric-ccenv
* fabric-javaenv
* fabric-kafka
* fabric-zookeeper
* fabric-tools
* fabric-ca