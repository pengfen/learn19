# https协议配置

## 生成https证书
* https://freessl.cn/
* 输入域名 res.com
* 以下需要在域名管理后台配置
* 验证域名 res.com
* TXT记录 _dnsauth
* 记录值 2019011803575732lzrn8163ljl21l6izb1aqppd412aldng6g227njrzgqonq8x

## 配置https证书
* 将https私钥复制到 /opt/
* res.conf中配置

## 使用腾迅云证书配置https
* 腾迅云管理后台搜索ssl
* 申请免费证书
* 下载证书上传至服务器
* test.conf 配置证书 (注意 阿里云管理后台需要开放443端口)
listen 80;
listen 443;

server_name test.com;

ssl on;
ssl_certificate /etc/nginx/test_cret/1_test.com_bundle.crt;
ssl_certificate_key /etc/nginx/test_cret/2_test.com.key;
ssl_session_timeout 5m;
ssl_protocols TLSv1 TLSv1.1 TLSv1.2; #按照这个协议配置
ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:HIGH:!aNULL:!MD5:!RC4:!DHE; #按照这个套件配置
ssl_prefer_server_ciphers on;

root /project/dir;
index index.php

# http自动跳转至https
if ($server_post = 80) {
	return 301 https://$host$request_uri;
}


## 注意
* http 可以访问https接口
* https 不可以访问http接口