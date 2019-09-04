# 表属性设置

```bash
// 文章结构体
type Article struct {
    Id int `orm:"pk;auto"`
    Aname string `orm:"size(20)"`
    Atime time.Time `orm:"auto_now"`
    Acount int `orm:"default(0);null"`
    Acontent string
    Aimg string
}
```

* 设置主键 `pk`
* 设置自增 `auto` 当Field类型int int32 int64 uint uint32 uint64 时 可以设置字段为自增键 当模型定义里没有主键时 符合上述类型且名称为Id的Field将被视为自增键
* 设置默认值 `default(11)`
* 设置长度 `orm:size(100)`
* 设置允许为空 `null` 数据库默认是非空 设置null之后就可以变为`ALLOW NULL`
* 设置唯一 `orm:"unique"
* 设置浮点数精度 `orm:"digits(12);decimals(4)"` 总共12位 四位是小数位
* 设置时间 `orm:"auto_now_add;type(datetime)"`  `orm:"auto_now;type(date)"`
* auto_now 每次 model 保存时都会对时间自动更新   auto_now_add 第一次保存时才设置时间
* 设置时间的格式 type