# php 安装

## windows 直接安装xampp等集成工具
* hosts 文件修改
* 选中文件 ---> 右击 ---> 属性 ---> 安全选项 高级 ---> 更改权限 ---> 添加 ---> 选择主体 ---> 高级 ---> 立即查找 ---> 确立

## deepin(ubuntu) 安装
* sudo apt-get install -y php php-fpm

## 安装php7.2
* 安装软件源拓展工具  sudo apt -y install software-properties-common apt-transport-https lsb-release ca-certificates
* 添加 Ondrej Sury 的 PHP PPA 源，需要按一次回车   sudo add-apt-repository ppa:ondrej/php 
* 更新软件源缓存：sudo apt update
* sudo apt install php7.2-fpm php7.2-mysql php7.2-curl php7.2-gd php7.2-mbstring php7.2-xml php7.2-xmlrpc php7.2-zip php7.2-opcache php7.2-redis -y 安装php7.2及相关依赖

* apt install php-mysql php-curl php-gd php-mbstring php-xml php-xmlrpc php-zip php redis php-opcache


# nginx 
sudo apt-get install nginx

## 源码安装
* 安装gcc g++依赖库
sudo apt-get install build-essential
sudo apt-get install libtool
* 安装pcre依赖库
sudo apt-get update
sudo apt-get install libpcre3 libpcre3-dev
* 安装zlib依赖度
sudo apt-get install zlib1g-dev
* 安装ssl依赖库
sudo apt-get install openssl

* 下载最新版本 wget http://nginx.org/download/nginx-1.13.6.tar.gz
* 解压 tar -zxvf nginx-1.13.6.tar.gz
* 进入解压目录 cd nginx-1.13.6
* 配置 ./configure --prefix=/usr/local/nginx 
* 编译 make
* 安装 sudo make install
* 启动 sudo /usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf  注意：-c 指定配置文件的路径，不加的话，nginx会自动加载默认路径的配置文件，可以通过-h查看帮助命令。
* 查看进程  ps -ef | grep nginx
* sudo ln -s /usr/local/nginx/sbin/nginx /usr/bin/nginx

## 卸载 删除 nginx
* 删除nginx，–purge包括配置文件  sudo apt-get --purge remove nginx
* 自动移除全部不使用的软件包 sudo apt-get autoremove
* 罗列出与nginx相关的软件 dpkg --get-selections | grep nginx
* 删除相关软件 sudo apt-get --purge remove nginx    sudo apt-get --purge remove nginx-common
* 查看nginx正在运行的进程 ps -ef | grep nginx

## 卸载 删除 apache
dpkg -l | grep apache #查找apache相关包
sudo apt-get --purge remove apache

## mysql 安装
* sudo apt-get install mysql-server
* sudo apt-get isntall mysql-client
* sudo apt-get install libmysqlclient-dev

## 解决mysql5.7登录不了
* sudo cat /etc/mysql/debian.cnf
* mysql -u账户 -p密码
* use mysql
* update user set authentication_string=password("123456") where user="root"
* flush privileges;
* select user, host, plugin from user where user = "root"
* update user set plugin = "mysql_native_password" where user = "root"
* flush privileges;
