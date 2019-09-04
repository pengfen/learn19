# jenkins 介绍

* Jenkins之前叫做Hudson 是基于Java开发的一种持续集成工具 用于监控执行重复的工作 包括自动执行构建 测试 交付或部署软件相关的各种任务

* 软件构建自动化
* 构建可持续的自动化检查
* 构建可持续的自动化测试
* 生成后续过程的自动化

## jenkins 的优点
* 所有CI产品中在安装和配置上最简单的
* 基于Web访问 用户界面非常友好 直观和灵活 在许多情况下 还提供了AJAX的即时反馈
* jenkies 是基于Java开发的 但它不仅限于构建基于Java的软件
* jenkies拥有大量的插件 可以直接通过web界面进行安装和管理

## jenkins 安装
* jdk和tomcat
* jenkins war包
* war包放到webapps下 并启动进行设置
* https://jenkins.io/download/

## jdk 安装
* java -version

## Tomcat 安装
* 下载tomcat  https://tomcat.apache.org/download-80.cgi
* tar -zxvf apache-tomcat-8.5.35.tar.gz -C  ~/app/
* cd app ---> mv apache-tomcat-8.5.35 tomcat8.5
* cd tomcat8.5 ---> bin/startup.sh
* localhost:8080

* localhost:8000/jenkins
* cat /home/ricky/.jenkins/secrets/initialAdminPassword
* tail -f /opt/soft/tomcat8.5/logs/catalina.out
* /jenkins/pluginManager/advanced
* 重启 shutdown.sh  startup.sh
* 安装推荐插件 admin 6543210.

## 权限
* 系统管理 ---> 插件管理 ---> 已安装 ---> 搜索 role
* 安装 role-based 插件 ---> 全局安全配置 ---> 授权策略 ---> Role-Based Strategy
* Manage and Assign Roles 管理和分配权限