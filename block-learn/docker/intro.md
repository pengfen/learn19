# docker介绍

*               传统虚拟化                                        容器虚拟化
* 创建速度        很慢                                            非常快
* 性能影响        通过对于硬件层的模拟 增加了系统调用链路的环节 有性能损耗  共享kernel 几乎没有性能损耗
* 资源消耗        很大                                            很小 一台机器可以轻松创建多个Container
* 操作系统覆盖     支持linux windows mac等                         仅仅kernel所支持的OS

* CGroups 限制容器的资源使用
* Namespace 机制 实现窗口间的隔离
* chroot 文件系统的隔离

* Linux 内核提供的限制 记录和隔离进程组所使用的资源 由Googles的工程师提出 后台被整合到kernel
* 通过不同的子系统(blkio cpu cpuacct等) 来实现对不同资源使用控制和记录

* pid 容器有自己独立的进程表和1号进程
* net 容器有自己独立的network info
* ipc 在ipc通信时候 需要加入额外信息来标示进程
* mnt 每个容器有自己唯一的目录挂载
* utc 每个容器有独立的hostname和domain

* docker 文件系统
* advanced multi layer unification filesystem
* 可以实现多个不同目录的内容合并在一起
* 允许read-only和read-write目录并存
* docker使用aufs来实现分层的文件系统的管理
* 只读部分定义为Image 可写部分是container
* Image类似一个单链表系统 每个Image包含一个指向parent Image指针
* 没有parent image 的 image是base image

* 类似于Github的服务 用来分发Image
* 大量标准的Image 例如 Tutum/Ubuntu Tutum/Mysql

* docker search(repository name)
* 默认是在 docker hub上进行搜索

## 安装docker
* sudo apt-get remove docker docker-engine

* sudo apt-get install apt-transport-https ca-certificates curl python-software-properties software-properties-common
* curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -
* sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/debian wheezy stable"
* 以上会写入 /etc/apt/sources.list
* sudo apt-get update
* sudo apt-get install docker-ce
* sudo docker version
* sudo docker run hello-world

## 删除运行容器
* 查询正在运行的容器 docker ps -a
* 删除运行的容器  docker rm -f $(docker ps -aq)