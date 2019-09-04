# post 请求　

## postman
* json 格式　　设置Headers   Content-Type application/json
* body raw
```
{
	"name":"",
	"value";"",
}
```
* php接受json   file_get_contents("php://input");

## jmeter设置Headers
* Add ---> Config Element ---> HTTP Header Manager