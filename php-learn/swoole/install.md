# swoole 安装

## 安装phpsize
* sudo apt-get install -y php-dev
* sudo apt-get install -y php7.2-dev (安装具体版本)

## 安装swoole
* 下载swoole https://github.com/swoole/swoole-src/releases
* 解压 tar -zxvf swoole... -C /...
* cd /... 进入相关目录
* phpize (phpize7.2 注意版本)
* php7.0 ---> /usr/lib/php/20151012 /usr/include/php/20151012
* php7.2 ---> /usr/lib/php/20170718 /usr/include/php/20170718
* ./configure 检测配置文件
* make 编译
* sudo make install 安装

* php -m
* php -i | grep php.ini

* vi /etc/php/cli/php.ini
* extension=swoole.so