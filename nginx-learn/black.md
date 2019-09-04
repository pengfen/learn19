# nginx 服务器上配置黑名单

* /etc/nginx 目录创建黑名单文件 black.conf
* deny IP
* 单个网站屏蔽IP  include black.conf 放到网址对应的server {}语句块中
* 多个网站屏蔽IP inclue black.conf 放到nginx.conf http {} 语句块中