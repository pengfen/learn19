# 配置路由  域名/v1/controller/action.json

* nginx配置文件中配置重写   /v1/controller/action.json ---> /v1/?s=controller.action
``` 
if ( !-f $request_filename )
{
    rewrite ^/v1/(.*)/(.*).json /v1/?s=$1.$2;
}
```

* 域名指向 项目路径/public/
public
   --- v1
        --- index.php
        --- init.php (注意需要修改API_ROOT路径)

# 不使用字符路由替换
```
<?php
namespace App\Common;

// ./src/app/Common/Request.php
// ?s=wel.wel ---> ?r=wel/wel
class Request extends \PhalApi\Request {

    public function getService() {
        // 优先返回自定义格式的接口服务名称
        $service = $this->get('r');
        if (!empty($service)) {
            $namespace = count(explode('/', $service)) == 2 ? 'App.' : '';
            return $namespace . str_replace('/', '.', $service);
        }

        return parent::getService();
    }
}

// config/di.php 注册相关服务
$di->request = new App\Common\Request();
```