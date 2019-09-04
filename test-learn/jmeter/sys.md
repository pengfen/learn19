# linux服务器监控性能测试

## 测试目的
* 发现服务器的性能瓶颈　配置的不同能够承载的最大任务数不同　能够承载的压力也不同

## 测试范围及性能指标
* CPU
* 内存
* 磁盘
* 网络
* 版本

## 测试与生产的环境配置不同
* 通过多次压测来计算性能损耗
* 性能损耗的计算方式 多次压测后的性能预估

## 进程与线程
* 进程是具有一定独立功能的程序关于某个数据集合上的一次运行活动　是系统进行资源分配和调度的一个独立单位
* 线程是进程的一个实体　是cup调度和分派的基本单位　他是比进程更小的能够独立运行的基本单位　线程自己基本上不拥有系统资源　只拥有一点在运行中必不可少的资源　一个线程可以创建和撤销另一个线程
* 一个线程只能属于一个进程　而一个进程可以拥有多个线程
* 线程是进程工作的最小单位
* 一个进程会分配一个地址空间　进程与进程之间不共享地址空间　即不共享内存
* 同一个进程下的不同的多个线程　共享父进程的地址空间
* 线程在执行过程中　需要协作同步　不同进程的线程间要利用消息通信的办法实现同步
* 线程作为调度和分配的基本单位　进程作为拥有资源的基本单位

* 进程 - 优点
* 每个进程互相独立　不影响主程序的稳定性　子进程崩溃不影响其他进程
* 通过增加CPU 就可以扩充性能
* 进程 - 缺点
* 逻辑控制复杂　需要和主程序交互
* 线程 - 优点
* 程序逻辑和控制方式简单
* 所有线程可以直接共享内存和变量等
* 线程 - 缺点
* 每个线程与主程序共用地址空间　最大内存地址受限
* 线程之间的同步和加锁不易控制

## Linux帮助命令
* man Linux下的函数手册命令 可以查看所有命令的使用方法
* top 能够实时监控系统的运行状态　并且可以按照cpu及内存等进行排序 (top -h)　top -hv | -bcHiOSs -d secs -n max -u|U user -p pid(s) -o field -w [cols]
* -h 帮助
* -p 监控指定的进程　当监控多个进程时　进程ID以逗号分隔　这个选项只能在命令行下使用 top -p 1, 2
* M 按内存使用率排序 top 时　按M
* P 按CPU使用率排序
* z 彩色/黑白显示
* top中的load average 系统的运行队列的平均利用率　也可以认为是可运行进程的平均数　三个值分别表示在最后的一分钟　五分钟　十五分钟的平均负载值
* 在单核cpu中load average的值为1时表示满负荷状态　同理在多核cpu中满负载的load average的值为1Xcpu核数
* vmstat 可以监控操作系统的进程状态　内存　虚拟内存　磁盘IO CPU的信息　语法 vmstat [-a] [-n] [-S unit] [delay [count]]
* -S 使用指定单位显示　参数有k K m M 分别代表1000 1024 1000000 1048576字节(byte) 默认单位为K(1024 bytes)
* free 能够监控系统的内存使用状态　其中　total 总计物理内存的大小 Used 已使用多大 Free 可用有多少 shared 多个进程共享的内存总额　buffers/cached磁盘缓存的大小

## 实时监控 - cpu
* mpstat 最大的特点是 可以查看多核心cpu中每个计算核心的统计数据 语法 mpstat[-P {|ALL}] [internal [count]]
* -P {|ALL} 表示监控哪个CPU cpu在[0,cpu个数-1]中取值
* internal 相邻的两次采样的间隔时间
* count　采样的次数　count只能和delay一起使用
* 当没有参数时　mpstat则显示系统启动以后所有信息的平均值　有interval时 第一行的信息自系统启动以来的平均信息　从第二行开始　输出为前一个interval时间段的平均信息
* 如果mpstat命令不存在 sudo apt-get install -y sysstat 进行安装

## 实时监控 - 网络
* netstat参数说明
* -n 拒绝显示别名　能显示数字的全部转化成数字 (如 127.0.0.1 别名localhost)
* -l 仅列出有Listen(监听)的服务状态
* -p 显示建立相关链接的程序名
* -t(tcp) 显示tcp相关选项
* -u(udp) 仅显示udp相关选项
* -i 显示自动匹配接口的信息
* -c 每隔一个固定时间　执行该netstart命令
* netstat -ntlp

## 实时监控 - 磁盘
* iostat　是对系统磁盘IO操作进行监控　它的输出主要显示磁盘的读写操作的统计信息　同时给出cpu的使用情况
* iostat [-c|-d][-k|-m][-t][-V][-x][device[...]|ALL][-p[device|ALL]][interval|[count]]
* -x device 输出指定要统计的磁盘设备名称　默认为所有磁盘设备

## 万能命令 sar
* sar(System Activity Report 系统活动情况报告)　是目前Linux上最为全面的系统性能分析工具之一 可以从多方面对系统的活动进行报告
* sar的性能监控范围  文件的读写情况　系统调用的使用情况　磁盘I/O CPU效率　内存使用状况　进程活动及IPC有关的活动等
* sar [options][-A][-o file] t [n]
* sar语法说明　在命令行中 n和t两个参数组合起来定义采样间隔和次数 t为采样间隔　是必须有的参数　h为采样次数　是可选的　默认值是1 -o file表示将命令结果以二进制格式存放在文件中 file在此处不是关键字 是文件名 options为命令行选项
* -A 所有报告的总和
* -u CPU利用率
* -r 显示系统内存的使用情况
* -B 内存分页情况
* -b 缓冲区使用情况

## strace
* strace命令是一个集诊断　调试　统计与一体的工具　我们可以使用strace对应用的系统调用和信号传递的跟踪结果来对应用进行分析　以达到解决问题或者是了解应用工作过程的目的
* -p 跟踪指定的进程
* -f 跟踪由fork子进程系统调用
* -e expr 输出过滤器　通过表达式　可以过滤掉你不想要输出
* -o filename 默认strace将结果输出到stdout 通过-o可以将输出写入到filename文件中

## Linux监控工具
* nmon 是一种在Linux操作系统上广泛使用的监控与分析工具　nmon所记录的信息是比较全面的 它能在系统运行过程中实时地捕捉系统资源的使用情况　并且能输出结果到文件中 然后通过nmon_analyzer工具产生数据文件与图形化结果
* https://sourceforge.net/projects/nmon/
* -f 这是nmon必选参数 并且必须放在第一个　就是输出文件的意思　用该参数的话　nmon输出的文件名就是默认名称　hostname_date_time.nmon
* -F <filename> 这个参数和-f相同 只不过用户可以自己定义文件名称
* -s 采集数据频率　也就是保存数据的频率
* -c 采集数据次数
* -t 输出最消耗资源的进程数据
* -h 查看帮助

## nmon_analyser
* nmon analyser 的作用就是分析nmon数据采集后的结果　nmon analyser需要借助Excel的宏 WPS默认没有安装宏　需要下载插件
* SYS_SUMM 系统汇总页 包含cpu占有率变化情况磁盘IO的变化情况等信息
* AAA 关于操作系统以及nmon本身的一些信息
* CPUnn 显示执行时间内CPU占用情况
* CPU_ALL 所有CPU概述　显示所有CPU平均占用情况
* CPU_SUMM 每一个CPU在执行时间内的占用情况
* DGBUSY 磁盘组每个hdisk设备平均占用情况
* DGREAD 每个磁盘组的平均读情况
* DGSIZE 每个磁盘组的平均读写情况
* DGWRITE 每个磁盘组的平均写情况
* DGXFER 每个磁盘组的I/O每秒操作
* MEM 内存相关的主要信息　使用　空闲内存大小等
* NET 显示系统中每个网络适配器的数据传输速率(千字节/秒)
* PAGE 本sheet统计相关页信息的记录

## Linux定时任务 crontab
* linux系统是由cron这个系统服务来控制的 Linux系统上包含很多的计划性工作 使用者自己也可以设置计划任务 所以Linux系统提供了使用者控制计划任务的命令
* crontab的启动
* /sbin/service crond status 查看定时任务的服务是否启动
* start stop restart 启动服务　停止服务　重新启动服务
* reload 重新载入配置
* crontab的服务权限
* crontab的权限管理存储在cron.allow文件与cron.deny文件中 如果没有可创建在etc目录下
* cron.allow文件存储的是允许哪些用户使用crontab
* cron.deny文件存储的是不允许哪些用户使用crontab
* crontab的使用场景说明
* 当两个文件都不存在时　那么只允许root用户使用crontab
* 当cron.allow文件存在　而cron.deny文件不存在时　那么只允许cron.allow文件中的用户使用crontab
* 当cron.deny文件存在　而cron.allow文件不存在时　那么只要没有列在cron.deny文件中的用户都可以使用crontab
* 如果两个文件都存在　而一个用户在两个文件中都有　那么以cron.allow文件中的为准　只要cron.allow文件用有该用户　则该用户就可以使用crontab
* crontab的使用
* crontab -e 在编辑页面输入命令即可
* crontab 的编辑格式
* 基本格式 minute hour day month week command
*         分　时　日　月　周　命令
* crontab的时间单位说明
* 第１列表示分钟 00~59　每分钟用*或者*/1表示
* 第２列表示小时 00~23(0表示０点)
* 第３列表示日期01~31
* 第４列表示月份01~12
* 第５列标识星期0~6(0表示星期天)
* 第６列要运行的命令
* crontab的符号说明
* 代表取值范围内的所有值
* / 代表每的意思
* -代表从某个数字到某个数字
* ,分隔开几个不同的数字
