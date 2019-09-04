// 运行代码 node wel.js

// 输出一个三角形

// 循环有多少行
for (var i = 0; i < 10; i ++) {
	// 每行要输出多少个 *
	for (var j = 0; j <= i; j ++) {
		// 每行输出的 * 不能换行
		process.stdout.write("*");
	}
	// 执行
	console.log()
}

// process 模块在使用的时候无需通过 require() 函数来加载譔模块 可以直接使用 因为是全局模块