/**
文件读写 使用模块 var fs = require('fs');

写文件 fs.writeFile(file, data[, options], callback);
参数1 要写入的文件路径 必填
参数2 要写入的数据 必填
参数3 写入文件时的选项 比如文件编码
参数4 文件写入完毕后的回调函数 必填

注意事项
该操作采用异步执行
如果文件已经存在则替换掉
默认写入的文件编码为utf8
回调函数有1个参数 err 表示在写入文件的操作过程中是否出错了 出错err != null  否则 err === null
*/

// 实现文件写入操作
// 1. 加载文件操作模块 fs 模块
//var fs = require("fs");

// 2. 实现文件写入操作
// var msg = "welcome to nodejs world";

// // 调用 fs.writeFile() 进行文件写入
// // fs.writeFile(file, data[,options], callback)
// fs.writeFile('./wel.log', msg, 'utf8', function (err) {
// 	// 如果 err === null 表示写入文件成功 没有错误
// 	// 只要 err 里面不是 null 就表示写入文件失败
// 	if (err) {
// 		console.log("写文件: " + err);
// 	} else {
// 		console.log('ok');
// 	}
// })


// 实现文件读取操作
// 1. 加载 fs 模块
var fs = require("fs");

// 2. 调用 fs.readFile() 方法来读取文件
// fs.readFile(file[, options], callback)
fs.readFile('./wel.txt', function (err, data) {
	if (err) {
		throw err;
	}

	// data 参数的数据类型是一个 Buffer 对象 里面保存的就是一个一个的字节
	// 把 buffer 对象转换为字符串 调用 toString() 方法
	// console.log(data)
	console.log(data.toString('utf8'));
});


// 读写文件中的路径问题
var fs = require("fs");

fs.readFile("./wel.txt", 'utf8', function (err, data) {
	if (err) {
		throw err;
	}
	console.log(data);
});

// 解决在文件读取 ./ 相对路径的问题
// 解决  __dirname __filename
// __dirname 表示 当前正在执行 js 文件所在的目录
// __filename 表示 当前正在执行的 js 文件的完整路径 
// 注意 以上二者并不是全局的， 会放在 ()() 中执行

var fs = require('fs');

//var filename = __dirname + '\\' + 'wel.txt';
// 处理斜杠
var path = require("path");
var filename = path.join(__dirname, 'wel.txt');

fs.readFile(filename, 'utf8', function(err, data) {
	if (err) {
		throw err;
	}

	console.log(data);
})

// 创建一个文件夹

// 加载文件操作模块
var fs  = require("fs");

// 创建一个目录
fs.mkdir("test-mkdir", function (err) {
	if (err) {
		console.log('创建目录出错了, 详细信息如下',)
	} else {
		console.log("目录创建成功");
	}
})