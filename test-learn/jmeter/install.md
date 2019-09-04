# jmeter

## 安装
* 下载 http://jmeter.apache.org/download_jmeter.cgi　　下载二进制包 tgz
* tar -zxvf ... -C ~/app/
* cd ~/app/ap../bin
* ./jmeter.sh 启动

## 创建图标
* sudo vim ~/.local/share/applications/apache_jmeter.desktop
```
[Desktop Entry]
Encoding=UTF-8
Version=1.0
Type=Application
Name=Apache JMeter (3.2 r1790748)
Icon=apache_jmeter.png
Path=/home/ricky/app/apache-jmeter-5.0
Exec=java -jar /home/ricky/app/apache-jmeter-5.0/bin/ApacheJMeter.jar
StartupNotify=false
StartupWMClass=Apache JMeter
OnlyShowIn=Unity;
X-UnityGenerated=true
```