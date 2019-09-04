# go 常用命令
* go build compile packages and dependencies 最常用的go command之一 编译go文件
* 跨平台编译 env GOOS=linux GOARCH=amd64 go build
* install compile and install packages and dependencies 编译 与build最大的区别是编译后会将输出文件打包成库放在pkg下
* 常用于本地打包编译的命令 go install 
* get download and install packages and dependencies 用于获取go的第三方包 通常会默认从git repo上pull最新的版本
* 常用命令如  go get -u github.com/go-sql-driver/mysql(从github上获取mysql的driver并安装至本地)
* fmt go fmt (reformat) package sources 类似于C中的lint 统一代码风格和排版
* 常用命令如 go fmt 
* test test packages  运行当前包目录下的tests
* 常用命令如 go test或go test -v等
* Go的test一般以XXX_test.go为文件名
* XXX的部分一般为XXX_test.go所要测试的代码文件名 (注 Go并没有特别要求XXX的部分必须是要测试的文件名)
