// 创建一个简单的 http 服务器程序

// 1. 加载 http 模块
var http = require('http');

// 2. 创建一个 http 服务对象
var server = http.createServer();

// 3. 监听用户的请求事件 (request事件)
// request 对象包含了用户请求报文中的所有内容 通过request对象可以获取所有用户提交过来的数据
// response 对象用来向用户响应一些数据 当服务器要向客户端响应数据的时候必须使用response 对象
server.on("request", function (req, res) {
	// 解决乱码 服务器通过设置 http 响应报文头告诉浏览器使用相应的编码来解析网页
	// text/plain 纯文本
	// text/html html
	res.setHeader("Content-Type", 'text/html; charset=utf-8');
	res.write("welcome to nodejs world, <h1>美好</h1>");
	// 对于每一个请求 服务器必须结束响应 否则客户端(浏览器) 会一直等符服务器响应结束
	res.end();
});

// 4. 启动服务
server.listen(3000, function () {
	console.log("服务器启动 请访问 http://localhost:3000")
});