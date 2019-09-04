<?php

/**
  * 单例模式链接数据库
  *
  * 单例模式三大原则
  *
  * 构造函数需要标记为非 public(防止外部使用 new 操作符创建对象) 单例类不能在其他类中实例化 只能被其自身实例化
  * 拥有一个保存类的实例的静态成员变量 $_instance
  * 拥有一个访问这个实例的公共的静态方法
*/

class Db{

    static private $_instance;  // 实例
    static private $_connectSource; // 数据连接源

    // 数据库相关配置 (主机 用户名 密码 数据库名)
    private $_dbConfig = array(
    	'host' => '127.0.0.1',
    	'user' => 'root',
    	'password' => '',
    	'database' => 'test'
    	);

	private function __construct(){
		// 构造函数
	}

	static public function getInstance(){ // 获取实例
		if (!(self::$_instance instanceof self)) {
			self::$_instance = new self();
		}
		
		return self::$_instance;
	}

	public function connect(){ // 获取数据库连接
		if (!self::$_connectSource) {

			self::$_connectSource = mysql_connect($this->_dbConfig['host'], $this->_dbConfig['user'], $this->_dbConfig['password']);

			if (!self::$_connectSource) {
				die('mysql connect error'.mysql_error());
			}

			mysql_select_db($this->_dbConfig['database'], self::$_connectSource);

			mysql_query("set names UTF8", self::$_connectSource);
		}

		return self::$_connectSource;
	}

}

$connect = Db::getInstance();
var_dump($connect);