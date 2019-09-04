/**
request 对象 和 response 对象

request 对象类型 <http.IncomingMessage>, 继承自stream.Readable
request 对象常用成员
request.headers
request.rawHeaders
request.httpVersion
request.method
request.url
*/

var http = require('http');

http.createServer(function (req, res) {
	// 1. 获取所有请求报文头
	// req.headers 返回的是一个对象 这个对象中包含了所有的请求报文头
	console.log(req.headers);

	// req.rawHeaders 返回的是一个数组 数组中保存的都是请求报文头的字符串
	console.log(req.rawHeaders);

	// 2. httpVersion 获取请求客户端所使用的http版本
	console.log(req.httpVersion);

	// 3. method 获取客户端请求使用的方法 
	console.log(req.method);

	// 4. url 获取这次请求的路径 (不包含主机名称 端口号 协议)
	console.log(req.url);

	res.end("game over");
}).listen(3000, function() {
	console.log('http://localhost:3000');
});


// 1. res.write()
res.write("welcome");

// 2. 每个请求都必须要调用一个方法 res.end(); 结束响应(请求)
// 告诉服务器该响应的报文头 报文体等等全部已经响应完毕了 可以考虑本次响应结束
// res.end() 要响应数据的话 数据必须是 string 类型或是 Buffer 类型

// 3. 通过 res.setHader() 来设置响应报文头
// rea.setHeader() 要放在 res.write() 和 res.end() 之前设置
res.setHeader("Content-Type", 'text/plain; charset=utf-8');

// 4. 设置 http 响应状态码
// res.statusCode 设置 http 响应状态码
// res.statusMessage 设置 http 响应状态码对应的消息
res.statusCode = 404;
res.statusMessage = 'Not Found';

// 5. res.writeHead() 
// 直接向客户端响应 (写入) http响应报文头
res.writeHead(404, 'Not Found', {
	'Content-Type':'text/plain; charset=utf-8'
});