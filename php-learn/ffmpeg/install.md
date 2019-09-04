1. 下载ffmpeg 安装包 http://ffmpeg.org/
2. 解压文件到指定目录
3. 检测环境 ./configure --enable-shared --prefix=/usr/local/ffmpeg --disable-yasm
4. 编译 make
5. 安装 sudo make install
6. 测试转码 /usr/bin/ffmpeg -i aaa.avi -c:v libx264 -strict -2 bbb.mp4
