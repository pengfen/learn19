# php 相关问题

## 问题 服务器默认 php7.0  由于框架需要安装php7.2 后启动项目出现 权限限制
* 解决 修改 /etc/php/7.2/fpm/pool.d/www.conf 文件

* user = 
* group = 

* listen.owner =
* listen.group = 

* 查看 /etc/nginx/nginx.conf  user www-data; user 使用什么 www.conf 文件中使用什么

## 500 错误
* 使用postman调用接口时什么都没有
* 将代码抽出到单独文件中，使用php运行文件 查看具体错误信息
* 根据错误信息修改