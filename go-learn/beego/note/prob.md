## bee run 出现错误 no space left on device
* 解决方案一 
* go build -o wel .
* ./wel

* 解决方案二
* 编辑/etc/sysctl.conf  vim /etc/sysctl.conf
* 添加  fs.inotify.max_user_watches = 65536
* 重启配置  sudo sysctl -p /etc/sysctl.conf
* sudo reboot
* bee runs