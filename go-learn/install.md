# go 安装

# GVM
* curl -s https://raw.github.com/moovweb/gvm/master/binscripts/gvm-installer
* gvm install go1.0.3
* gvm use go1.0.3

# apt-get 安装
* sudo add-apt-repository ppa:gophers/go
* sudo apt-get update
* sudo apt-get install golang-stable

# 二进制包安装
* tar -zxvf go... -C ~/app
* vim .bashrc
```
export GOROOT=/home/ricky/app/go
export GOPATH=$HOME/go
export PATH=$PATH:/home/ricky/go/bin
```
* source .bashrc
* go version 