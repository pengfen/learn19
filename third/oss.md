```
/* 阿里云上传图片 */
//  $files ---> $_FILES["file"]
public function upload($files, $type=0) {
    $postdata = fopen($files['tmp_name'], "r");
    /* Get file extension */
    $extension = substr($files['name'], strrpos($files['name'], '.'));

    /* Generate unique name */
    $suffix = uniqid() . $extension;
    $filename = "/var/www/11show/api/public/upload/" . $suffix;
    //$filename = "/home/ricky/app/11show/api/public/upload/" . $suffix;

    /* Open a file for writing */
    $fp = fopen($filename, "w");
    /* Read the data 1 KB at a time
      and write to the file */
    while ($data = fread($postdata, 1024))
        fwrite($fp, $data);
    /* Close the streams */
    fclose($fp);
    fclose($postdata);
    @unlink($files['tmp_name']);

    // 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
    $accessKeyId = "";
    $accessKeySecret = "";
    // Endpoint以杭州为例，其它Region请按实际情况填写。
    $endpoint = "oss-ap-southeast-5.aliyuncs.com";
    // 存储空间名称
    $bucket= "";
    // 文件名称
    //$object = $_FILES["name"];
    $dir = date("Ym", time());
    $object = $dir.'/'.$suffix;
    // <yourLocalFile>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt
    //$filePath = $_FILES["tmp_name"];
    // $filePath = $filename;

    $url = '';

    try{
        $ossClient = new OssClient($accessKeyId, $accessKeySecret, $endpoint);
        $res = $ossClient->uploadFile($bucket, $object, $filename);
        // 注意目录权限问题
        // 图片缩放 url?x-oss-process=image/resize,w_200
        $url = $res["oss-request-url"];
        @unlink($filename);
    } catch(OssException $e) {
    }
    // 1,http://pitx7ct6j.bkt.clouddn.com/20190220/5c6cfd6748805.jpg,1,20190220/5c6cfd6748805.jpg
    if ($type) return "1,".$url.",1,".$object;
    return $url;
}
```