<?php

echo "welcome to php world";  // 打印第一个 php 程序

phpinfo(); // 显示 php 信息函数


# 脚本注释 不常用
echo time();
echo "<br>";

/*
多行注释
功能： 实现时间输出
作者： 小艳艳
时间： 2015年4月1日
*/
echo time();
echo "<br>";

//输出时间  单行注释
echo time();
echo "<br>";

/*
啊，我是注释
不能嵌套 
*/


//注释
echo time();

echo time(); //输出时间  注释

// 1. 新建php
// 文件后缀 .php
// 必须通过apache（服务器）运行
// 2. php嵌套在html中
// 文件后缀 .php

// 3.  注释
// //  单行注释
// #   单行注释
// /*
//      多行注释
// */
// 4. php 指令分隔符  
// ;
// 5. 空白符的使用
// 注释前面有空行
// 逻辑代码段 空行
// 缩进四个空格 
?>
<!doctype html>
<html>
	<head>
		<meta charset="utf-8">
		<title>HTML中嵌套PHP</title>
		<style>
			body{
				// 输出颜色值
				background:rgb(<?php echo rand(0,255);?>, <?php echo rand(0,255);?>, <?php echo rand(0,255);?>);
			}
		</style>
	</head>
	<body>
		<h1>HTML嵌套PHP</h1>
		<?php
		    echo 1+1;
		?>
		<hr>
		<?php
			echo date("Y-m-d H:i:s");
			
			echo "<h3>艳艳我爱你</h3>";
		?>
		<hr>
		<img src="../images/bg02.jpg" width="<?php echo rand(100,200);?>">


		<h1>PHP标记</h1>
		<hr>
		<?php
			echo "我是标准风格， 我爱小艳艳";
		?>
		<hr>
		<script language="php">
		    echo "我是长风格，我爱小艳艳";
		</script>
		<hr>
		<?
		    echo "我是短风格,我爱小艳艳";
		?>
		<hr>
		<%
		    echo "我是asp风格，我爱小艳艳";
		%>
		<hr>
		
		<?php
		    echo time();
		?>
	</body>
</html>