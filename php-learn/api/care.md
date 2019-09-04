# app接口注意事项

## postman使用
* 选择POST请求 ---> 输入请求接口路由 ---> 选择Headers ---> 输入请求参数 ---> 发送请求 ---> 添加至集合 ---> 分享

* 错误返回值 {"code":1011, "msg":""}
* 1011 ...
* 1012 ...

* 正确返回值 {"code":200, "date":{...}}

## api 返回数据
* ["code" => 1011, "msg" => $code[1011]] => {"code" : 1011, "msg" : "" } // 单条记录
* ["article" => [[], [], []]] => {"article" : [{}, {}, {}]} // 多条记录

* 如果value为空 设置 ''等
* 如果单条记录为空(如某作者不存在) 设置null

* 单条记录使用单数user
* 多条记录使用复数users
* 增删改使用token判断用户是否有效


* json_encode 无效
* var_dump(json_last_error());

* //错误码对照

0 JSON_ERROR_NONE  没有错误发生

1 JSON_ERROR_DEPTH  到达了最大堆栈深度

2 JSON_ERROR_STATE_MISMATCH 无效或异常的 JSON

3 JSON_ERROR_CTRL_CHAR 	控制字符错误，可能是编码不对

4 JSON_ERROR_SYNTAX 语法错误

5 JSON_ERROR_UTF8 	异常的 UTF-8 字符，也许是因为不正确的编码。

6 JSON_ERROR_RECURSION

7 JSON_ERROR_INF_OR_NAN

8 JSON_ERROR_UNSUPPORTED_TYPE 指定的类型，值无法编码。
 
## 权限问题
* 当前用户是否拥有查看当前订单权限
* 用户id 是否等于 当前订单所属用户id