# Cacti
* 监控工具
* 收集数据 根据数据绘图

# Nagios
* 监控工具
* 主机 服务/资源
* OK WARNGING CRITICAL UNKNOWN
* CPU 90%(CRITICAL) 80%(WARNGING) OK(UNKNOWN)
* 报警系统

## Nagios core
* 不做任何监控工作
* Plugins(scripts) check_nginx

### N种对象实现监控工作
* 主机 主机组
* 服务/资源 服务组
* 联系人 联系人组
* 时段
* 命令(应用到某个被监控对象 以实现具体的监控)

## Nagios 组成
* Nagios 通常由一个主程序(Nagios) 一个插件程序(Nagios-plugins)和四个可选的ADDON(NRPE NSCA NSClient++和NDOUtils)组成 Nagios的监控工作都是通过插件实现的 因此 Nagios和Nagios-plugins是服务器端工作所必须的组件
* NRPE 用来在监控的远程Linux/Unix主机上执行脚本插件以实现对这些主机资源的监控
* NSCA 用来让被监控的远程Linux/Unix主机主动将监控信息发送给Nagios服务器

## Nagios 安装
* Nagios 基本组件的运行依赖于httpd gcc gd
* yum -y install httpd gcc glibc glibc-common gd gd-devel php php-mysql mysql mysql-devel mysql-server
* 添加nagios运行所需要的用户和组
* groupadd -m nagcmd
* useradd -G nagcmd nagios
* passwd nagios

* 把apache加入到nagcmd组 以便于在通过web Interface操作 nagios时能够具有足够的权限
* usermod -a -G nagcmd apache

* 编译安装nagios