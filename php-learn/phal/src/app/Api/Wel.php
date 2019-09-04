<?php 
namespace App\Api;

use PhalApi\Api;
use PhalApi\Exception\BadRequestException;

class Wel extends Api {

	// http://self.phal.com/?s=wel.wel
	// 配置支持?r=wel/wel链接
	// common 下添加 Request.php 文件
	// config/di.php 注册  $di->request = new App\Common\Request();
	public function wel() {
		// 添加纪录埋点
        // \PhalApi\DI()->tracer->mark();

        // 添加纪录埋点，并指定节点标识
        // \PhalApi\DI()->tracer->mark('DO_SOMETHING');

        // 自定义调试信息
        // \PhalApi\DI()->response->setDebug('x', $x);
		return ["title" => "welcome to phalapi world"];
	}

	// http://self.phal.com/?s=wel.fail
	// 中文默认unicode
	// config/di.php $di->response = new \PhalApi\Response\JsonResponse(JSON_UNESCAPED_UNICODE); 
	public function fail() {
		throw new BadRequestException('签名失败', 1);
	}

	// 定义规则  http://docs.phalapi.net/#/v2.0/api
	public function getRules() {
		return [
			// 获取HTTP请求方法，判断是POST还是GET
			// source post $_POST   get $_GET   cookie $_COOKIE   server $_SERVER   request $_REQUEST    header $_SERVER['HTTP_X']
			'method' => ['name' => 'REQUEST_METHOD', 'source' => 'server'],
			// 获取COOKIE中的标识
			'is_new_user' => ['name' => 'is_new_user', 'source' => 'cookie'],
			// 获取HTTP头部中的编码，判断是否为utf-8
			'charset' => ['name' => 'Accept-Charset', 'source' => 'header'],					
			'login' => [ // 操作方法
			    "username" => ["name" => 'username']
			],
		];
	}

    // http://self.phal.com/?s=wel.login&username=ricky
	public function login() {
		return ["username" => $this->username];
	}
}