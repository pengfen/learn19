# NVM

* 安装nvm  wget -qO- https://raw.githubusercontent.com/creationix/nvm/v0.33.0/install.sh | bash
* 添加环境变量
```bash
export NVM_DIR="/home/ricky/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && . "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion
source ~/.nvm/nvm.sh
```
* 使用环境变量生效  source .bashrc
* 安装node nvm install 6.9.5

## node
* node
* npm install remix-ide -g  安装remix-ide
* which remix-ide

## node进程重启
* pm2 start s1.js --name IM --watch   启动nodejs程序
* pm2 save 保存pm2启动信息
* pm2 startup 设置开机启动
* pm2 安装
* sudo apt-get install -y nodejs
* sudo apt-get install -y npm
* sudo npm install pm2 -g
* sudo pm2 start s1.js --name IM --watch