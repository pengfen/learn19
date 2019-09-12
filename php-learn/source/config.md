# 配置文件概述

TP配置文件种类
Config类 - 扩展内容说明

配置分类
惯例配置 核心框架内置的配置文件(thinkphp/convention.php) 无需更改
应用配置 每个应用的全局配置文件(项目根目录下的config目录下的文件)
模块配置 每个模块的配置文件(相同的配置参数会覆盖应用配置) 有部分配置参数模块配置是无效的 因为已经使用过(index/config/database.php)
动态配置 主要是指在控制器或者行为中进行(动态)更改配置 该配置方式只在当次请求有效 因为不会保存到配置文件中

ArrayAccess   提供像访问数组一样访问对象的能力的接口
offsetExists  检查偏移位置是否存在
offsetGet     获取一个偏移位置的值
offsetSet     设置一个偏移位置的值
offsetUnset   删除一个偏移位置的值

Yaconf
一个高性能的配置管理扩展 yaconf.so
安装yaconf
unzip 解压
phpize
./configure -php-config=/
make
make install
php -i | grep php.ini
extension=yaconf
yaconf.directory=项目目录

Config核心类
ArrayAccess (offsetSet offsetExists offsetUnset offsetGet)
set/get
加载配置文件:load() 
loadFile php yaml格式文件处理  parse处理 不同的文件类型 ini json xml 设计模式(工厂模式)