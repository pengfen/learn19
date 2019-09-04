# centos 安装mysql

## 关闭防火墙
* service iptables status
* service iptables stop
* chkconfig iptables off
* vim /etc/selinux/config - SELINUX=disabled

## 配置sysctl.conf
* vim /etc/sysctl.conf

## 检查操作系统上是否安装了MySQL
* rpm -qa | grep mysql
* rpm -e mysql... #删除指定安装包
* yum remove mysql* #删除mysql相关安装包

## 下载mysql源码包
* https://dev.mysql.com/downloads/mysql/5.7.html#downloads

## 添加用户和组
* id mysql #查看当前用户和组
* groupadd mysql #添加用户组
* useradd -d /home/mysql -g mysql -m mysql
* passwd mysql #设置mysql用户密码
* userdel -r mysql #删除用户及组

## 配置MySQL环境变量
* .bashrc .bash_profile #配置环境变量或使用ln(ln -s /.. /usr/bin/mysql)

## 创建目录及授权
* mkdir -p /u01/my3306/data
* ...
* chown -R mysql:mysql /u01/my3306
* chmod -R 755 /u01/my3306

## 解压mysql5.6
* tar -zxvf mysql...

## 配置yum源　安装cmake
* yum install -y cmake gcc gcc-c++ ncurses-devel bison zlib libxml openssl

cmake \
-DCMAKE_INSTALL_PREFIX=/u01/my3306 \
-DINSTALL_DATADIR=/u01/my3306/data  \
-DDEFAULT_CHARSET=utf8 \
-DDEFAULT_COLLATION=utf8_general_ci \
-DEXTRA_CHARSETS=all \
-DWITH_SSL=yes \
-DWITH_EMBEDDED_SERVER=1 \
-DENABLED_LOCAL_INFILE=1 \
-DWITH_MYISAM_STORAGE_ENGINE=1 \
-DWITH_INNOBASE_STORAGE_ENGINE=1 \
-DWITH_ARCHIVE_STORAGE_ENGINE=1 \
-DWITH_BLACKHOLE_STORAGE_ENGINE=1 \
-DWITH_FEDERATED_STORAGE_ENGINE=1 \
-DWITH_PARTITION_STORAGE_ENGINE=1 \
-DMYSQL_UNIX_ADDR=/u01/my3306/run/mysql.sock \
-DMYSQL_TCP_PORT=3306 \
-DENABLED_LOCAL_INFILE=1 \
-DSYSCONFDIR=/etc \
-DWITH_READLINE=on

## 编译并安装
* make
* sudo make install

## MySQL参数配置

## 初始化MySQL脚本

## 启动MySQL

## 登录MySQL
