# 扩展安装

## ubuntu redis扩展安装
---------------------------------------------------------------
apt-get install -y php-redis
php -m
sudo /etc/init.d/php7.0-fpm restart

## linux redis扩展安装
---------------------------------------------------------------

git clone https://github.com/phpredis/phpredis.git
cd phpredis
phpize
./configure
make
sudo make install
php -i | grep php.ini
cd /etc/php/7.0/cli
sudo vim php.ini
extension=redis.so
sudo nginx -s reload
php -m


注意事项
服务器上执行脚本使用cli/php.ini
页面访问使用 fpm/php.ini (使用phpinfo测试)
cd /etc/php/7.0/fpm
sudo vim php.ini
extension=redis.so

sudo nginx -s reload
sudo /etc/init.d/php7.0-fpm restart


## linux redis扩展安装 (方法二)
---------------------------------------------------------------
http://pecl.php.net/package/redis
下载扩展包 注意版本问题
redis-3.0.0.tgz (用于php7及以上 php5的会出现以下错误)
错误：ext/standard/php_smart_string.h：没有那个文件或目录

安装
cd /opt/php-ext 上传redis-2.2.8.tgz包
tar -zxvf redis-2.2.8.tgz #解压
cd redis-2.2.8.tgz #进入安装目录
/usr/local/php/bin/phpize #用phpize生成configure配置文件
./configure --with-php-config=/usr/local/php/bin/php-config  #配置
make  #编译
make install  #安装

[root@peng4 redis-2.2.8]# make install #安装完成之后，出现下面的安装路径
Installing shared extensions:     /usr/local/php/lib/php/extensions/no-debug-non-zts-20131226/

配置php支持
vi /usr/local/php/etc/php.ini  #编辑配置文件，在最后一行添加以下内容
添加
extension="redis.so"
:wq! #保存退出

lnmp restart #重启服务

测试 (使用code中代码测试)



## window redis扩展安装
---------------------------------------------------------------

nts (Non Thread Safe 非线程安全)
ts (Thread Safe 线程安全)

windows redis 扩展下载 http://pecl.php.net/package/redis

phpinfo(); 查看 php 版本 查看版本对应的位数

下载对应的位数(x86/X64) 扩展包

解压

将 php_redis.dll 复制到 php安装目录/ext

在 php.ini 文件中添加 extension=php_redis.dll

http://pecl.php.net/package/mongo/1.6.12/windows mongo 扩展

本系统不支持 Redis (需要安装redis扩展)