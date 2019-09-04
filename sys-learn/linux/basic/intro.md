API Application Program Interface 程序员面对的编程接口
ABI Application Binary Interface 程序应用者面对运行程序的接口

CLI
    bash zsh sh csh tcsh ksh

    查看所用的shell的类型
    echo $SHELL

远程连接
    ssh协议 secure shell
    ss -tnl 查年系统是否监听于tcp协议的22号端口
    ip addr list 或 ip addr show 或 ifconfig 查看IP地址
    查看防火墙状态 iptables -L -n

    Linux的哲学思想之一 一切皆文件
    表现之一 硬件设备也通过文件表示
    物理终端 /dev/console
    虚拟终端 /dev/tty
    串行终端 /dev/ttys
    伪终端 /dev/pts

    注意　在启动设备之后 在其上关联一个用户接口程序　即可实现与用户交互

账户
root 登录用户
model 当前主机的主机名 非完整格式　此处的完整格式为 model.magedu.com
~ 用户当前所在的目录(current directory) 也称为工作目录(working directory) 相对路径
# 命令提示符 # 管理员账号为root(拥有最高权限) $ 普通用户 非root用户(不具有管理权限　不能执行系统管理类操作)
注意　建议使用非管理账号登录　执行管理操作临时切换至管理员　操作完成即退回

几个基础命令m
    tty 查看当前的终端设备
    ifconfig 或 ip addr list 查看活动接口的ip地址
    echo 回显
    ping 探测网络的目标主机与当前主机之间的连通性 Ctrl + c 终止命令执行

    关机命令
    poweroff reboot halt

Stallman
    自由含义 自由学习和修改 自由使用　自由分发 自由创建衍生版

Linux 的哲学思想
    一切皆文件
        把几乎所有资源统统抽象为文件形式　包括硬件设备　甚至通信接口等 open()  read() write() close() delete() create()
    由众多功能单一的程序组成　一个程序只做一件事　并且做好  组合小程序完成复杂任务
    尽量避免跟用户交互　目标　易于以编程的方式实现自动化任务
    使用文本文件保存配置信息

文件系统　层级结构　有索引
    文件的路径表示
        绝对路径　从根开始表示的路径
        相对路径  从当前位置开始表示出的路径

    文件名使用法则
        严格区分字符大小写 file1 File1
        目录也是文件　在同一路径下　两个文件不能同名
        支持使用除/以外的任意字符
        最长不能超过255个字符

    用户有家目录 home
        用户的起始目录　普通用户管理文件的位置

    程序的组成部分　二进制程序文件　库文件　配置文件　帮助文件
    二进制 库文件　可执行文件　
    库文件　不能独立执行　只能被调用时执行
    配置文件　帮助文件　可被查看其内容的文件

BSD
    如果二次发布的产品中包含源代码　则在源代码中必须带有原来的代码中的BSD协定
    如果二次发布产品是二进制格式的库或程序　则需要在发布的文档或版本声明中说明包含原来的代码中的BSD协定
    不可以用开源代码的作者或者组织　以及原来的产品的名字做市场推广

命令的语法通用格式
COMMAND OPTIONS ARGUMENTS
发起一个命令　请求内核将某个二进制程序运行为一个进程

命令本身是一个可执行的程序文件　二进制格式的文件　有可能会调用共享库文件
多数系统程序文件都存放在 /bin /sbin /usr/bin /usr/sbin /usr/local/bin /usr/local/sbin
普通命令 /bin /usr/bin /usr/local/bin
管理命令 /sbin /usr/sbin /usr/local/sbin
共享库 /lib /lib64 /usr/lib /usr/lib64 /usr/local/lib /usr/local/lib64
32bits的库 /lib /usr/lib /usr/local/lib
64bits的库 /lib64 /usr/lib64 /usr/local/lib64