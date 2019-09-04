# yii安装

## 复制代码
* 创建项目(mkdir project_name) ---> 复制代码
* 赋于权限 chmod 755 init ---> ./init
* 修改数据库 common/config/main-local.php
* 配置域名 sudo vim /etc/hosts   cd /etc/nginx/conf.d
* 重启服务器 sudo service nginx restart
* 复制相关配置 api/web/index.php api/config/main-local.php params-local.php
* chmod -R 777 runtime