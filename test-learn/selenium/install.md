# 安装selenium

* sudo apt-get install python3-pip   安装pip
* pip3 install selenium              安装selenium
* pip3 show selenium                 显示是否成功

* 查看 chrome 浏览版本 (版本 68.0.3440.84（正式版本） （64 位）)
* https://blog.csdn.net/huilan_same/article/details/51896672 查看对应版本的selenium驱动
* http://chromedriver.storage.googleapis.com/index.html 下载对应chromedriver
* sudo mv chromedriver /usr/lib/python3/chromedriver
* sudo ln -s /usr/lib/python3/chromedriver /usr/bin/chromedriver