// 根据用户的不同请求 服务器做出不同的响应

// 1. 加载 http 模块
var http = require('http');

// 2. 创建 http 服务
http.createServer(function (req, res) {

    // 获取用户请求的路径
    console.log(req.url);

    // 结束响应
    //res.end();

    res.setHeader("Content-Type", 'text/html; charset=utf-8');
    // 通过 req.url 获取用户请求的路径 根据不同的请求路径服务器做出不同的响应
    if (req.url === '/' || req.url === '/index') {
    	// res.write("welcome to here")
    	// res.end();
    	res.end("welcome to index");
    } else if (req.url === '/login') {
    	res.end("welcome to login");
    } else if (req.url === '/list') {
    	res.end("welcome to list");
    } else if (req.url === '/register') {
    	res.end("welcome to register");
    } else {
    	res.end("404, not Found")
    }
}).listen(3000, function() {
	console.log('http://localhost:3000')
});