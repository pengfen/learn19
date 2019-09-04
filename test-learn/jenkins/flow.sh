#!/bin/bash

tomcat_home=/opt/soft/tomcat

### 检测 tomcat 是否存在
tomcat=`ps -ef | grep /home/ricky/app/tomcat8.5 | grep -v 'grep\|tail' | awk '{print $2}' `
echo ${tomcat}

if test -z ${tomcat}
then
    echo "tomcat 没有启动! "
else
	echo "shutdown tomcat!"
	sh ${tomcat_home}/bin/shutdown.sh
	sleep 2
fi

### 备份
mv ${tomcat_home}/webapps/hello.war ${tomcat_home}/hello.war-$(date "+%Y%m%d-%H%M%s")

### 部署
cp -r /root/.jenkins/workspace/testDeploy/target/hello.war ${tomcat_home}/webapps/hello.war

### 启动tomcat
sh ${tomcat_home}/bin/startup.sh

tomcat=`ps -ef | grep /home/ricky/app/tomcat8.5 | grep -v 'grep\|tail' | awk '{print $2}' `
echo ${tomcat}

if test -z ${tomcat}
then
    echo "tomcat 没有启动! "
else
	echo "部署成功"
fi