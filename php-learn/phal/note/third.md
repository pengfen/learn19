api ---
    Appapi
        Api
        Domain
        Model
    Config
    Library
        UCloud

第三方类库
将相关类库放在Library目录下
编写UCloud_Lite类

init.php中进行初始化UCloud_Lite类
DI()->ucloud = new UCloud_Lite();

控制器进行调用
DI()->ucloud->upfile($_FILES['file']);