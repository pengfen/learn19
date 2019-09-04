问题 服务器默认 php7.0  由于框架需要安装php7.2 后启动项目出现 权限限制

解决 修改 /etc/php/7.2/fpm/pool.d/www.conf 文件

user = 
group = 

listen.owner =
listen.group = 

查看 /etc/php/7.2/fpm/pool.d/www.conf 的内容