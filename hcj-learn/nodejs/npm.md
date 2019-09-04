# NPM (Node Package Manager Node包管理器)

* NPM 包管理库 存储了大量的JavaScript代码库
* NPM 客户端 npm命令行工具 使用JavaScript开发的基于nodejs的命令行工具 也是node的一个包

## NPM 安装
* npm 会随着nodejs自动安装 安装完毕nodejs后会自动安装npm
* 查看当前npm版本 npm -v
* 更新npm  npm install npm@latest -g

## NPM 使用
* https://www.npmjs.com 网站找到需要包
* 在项目根目录下 执行npm install 包名称 安装
* 在nodejs代码中通过require('包名') 加载该模块

## NPM 常用命令介绍
* install 安装包  npm install 包名
* uninstall 卸载包  npm uninstall 包名
* version  查看当前npm版本  npm version 或 npm -v
* init 创建一个package.json文件  npm init
* 注意 当使用npm init -y 的时候 如果当前文件夹(目录)的名字比较怪(有大写，中文等)就会影响npm init -y 生成操作
* --save npm5以下需要 是将包的相关信息写入package.json

## package.json 文件
* package.json 文件是一个包说明文件(项目描述文件) 用来管理组织一个包(一个项目)
* package.json 文件是一个json格式的文件
* 位于当前项目的根目录下

* package 内容
* name 包名字
* version 包版本
* description 包描述
* author 包作者
* main 包的入口js文件 从main字段指定那个js文件开始执行
* dependencies 当前包依赖的其他包