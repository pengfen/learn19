# php 安装

## windows 直接安装xampp等集成工具
* hosts 文件修改
* 选中文件 ---> 右击 ---> 属性 ---> 安全选项 高级 ---> 更改权限 ---> 添加 ---> 选择主体 ---> 高级 ---> 立即查找 ---> 确立

## deepin(ubuntu) 安装
* sudo apt-get install -y php

## 安装php7.2
* 安装软件源拓展工具  apt -y install software-properties-common apt-transport-https lsb-release ca-certificates
* 添加 Ondrej Sury 的 PHP PPA 源，需要按一次回车   add-apt-repository ppa:ondrej/php 
* 更新软件源缓存：apt update
* apt install php7.2-fpm php7.2-mysql php7.2-curl php7.2-gd php7.2-mbstring php7.2-xml php7.2-xmlrpc php7.2-zip php7.2-opcache -y 安装php7.2及相关依赖