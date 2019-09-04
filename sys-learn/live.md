# 直播项目迁移 (从A主机迁移到B主机)

## 迁移代码
* zip -r live.zip live/*  (将代码压缩至包中)
* mysqldump -u -p live > live.sql
* 将压缩包及sql迁移到B主机
* unzip live.zip
* mysql -h -u -p
* create database live charset=utf8
* use live
* source /var/www/live/live.sql

## 修改配置文件
* config.php

## 安装redis
* sudo apt-get install -y redis-server
* sudo apt-get install php-redis
* php -m
* sudo /etc/init.d/php7.0-fpm restart
* sudo vim /etc/redis/redis.conf
* 修改 requirepass password
* sudo /etc/init.d/redis-server restart

## 启动js (socket)
* sudo apt-get install -y nodejs npm
* node -v
* npm -v
* sudo npm install pm2 -g
* sudo pm2 start server.js --name IM --watch