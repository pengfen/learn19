#!/bin/bash

sslocal -s 18.1.1.1 -p 443 -k "bamcdx123." -l 1080 -t 300 -m aes-256-cfb
# sudo apt-get install shadowsocks
# sslocal -s server_ip -p server_port -k "keyword" -l local_port -t 300 -m aes-256-cfb
# server_ip 服务器IP
# server_port 服务器端口
# keyword 密码
# local_post 1080
# 设置本地代理  控制中心 ---> 系统代理 ---> Socks代理 127.0.0.1 1080