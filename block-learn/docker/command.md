# Docker 的常用命令

* 查看docker版本   sudo docker version
* 查看镜像的列表    sudo docker images
* 查看正在运行的镜像 sudo docker ps
* 查看容器的详细信息 sudo docker inspect ...
* docker inspect 后面的参数是docker ps命令返回结果中的 container id 字段的值
* 运行镜像 docker start -it -p 8880:80 nginx
* docker start (start后面的参数为容器编号或者容器名称 在执行完docker create命令后 可通过docker ps -a 命令获取)
* 创建容器并且运行容器  sudo docker run -d -p 8880:80 nginx
* -d 守护进程运行
* -p 端口映射 后面的是docker容器内的端口 前面的是宿主服务器的端口
* nginx 是镜像的名字
* 查询本镜像 sudo docker images
* 停止正在运行的镜像实例(容器) sudo docker stop ...
* 进入容器 sudo docker exec -it ... /bin/bash
* -it后面的就是容器实例的编号
* 导出docker镜像 sudo docker save -o 导出的文件路径 镜像名称 : 版本号
* 导出docker镜像 sudo docker load --input