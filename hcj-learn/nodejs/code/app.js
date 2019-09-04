/**
模拟 Apache 服务器处理静态资源
--- app.js
--- public
*/

// 加载 http 模块
var http = require('http');
var path = require("path");
var fs = require("fs");
// npm install mime
var mime = require("mime");

// 创建服务
http.createServer(function (req, res) {

	// 1. 获取用户请求的路径
	// req.url
	// /css/index.css
	// /images/index.png

	// 2. 获取 public 目录的完整路径
	var publicDir = path.join(__dirname, 'public');

	// 3. 根据 public 的路径和用户请求的路径 最终计算出用户请求的静态资源的完整路径
	var filename = path.join(publicDir, req.url);

	// 4. 根据文件的完整路径去读取该文件 如果读取到了 则把文件返回给用户 如果读取不到 则返回 404
	fs.readFile(filename, function(err, data) {
		if (err) {
			res.end("文件不存在");
		} else {

			// 通过第三方模块 mime 来判断不同的资源对应的 Content-Type 类型
			res.setHeader("Content-Type", mime.getType(filename));
			// 如果找到了用户要读取的文件 直接把该文件返回给用户
			res.end(data);
		}
	})

}).listen(3000, function() {
	console.log("http://localhost:3000");
});