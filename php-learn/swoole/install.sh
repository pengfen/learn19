#!/usr/bash

mkdir -p ~/build && \ # 创建安装目录
cd ~/build && \ # 进入安装目录
rm -rf ./swoole-src && \
curl -o ./tmp/swoole.tar.gz https://github.com/swoole/swoole-src/archive/master.tar.gz -L && \
tar zxvf ./tmp/swoole.tar.gz && \
mv swoole-src* swoole-src && \
cd swoole-src && \
phpize && \
./configur && \
--enable-coroutine \
--enable-openssl \
--enable-http2 \
--enable-async-redis \
--enable-sockets \
--enable-mysqlnd && \
make clean && make && sudo make install