# mysql性能

## 主要内容
* MySQL 数据库介绍
* MySQL 数据库监控指标
* MySQL 慢查询工作原理及操作
* SQL的分析与调优方法
* MySQL 索引的概念及作用
* MySQL 索引的工作原理与设计规范
* MySQL 存储引擎
* MySQL 实时监控
* MySQL 集群监控方案
* MySQL 性能测试的用例准备
* 使用Jmeter开发MySQL性能测试脚本
* 执行测试

## MySQL分支
* 主流分支 
* MySQL之父widenius创建　目标在于替换现有MySQL
* 兼容MySQL 对于开发者来说感知不到变化
* MariaDB is free and open source software

## MySQL 重点监控指标
* QPS queries per seconds 每秒钟查询数量
* show global status like 'Question%';
* Queries / seconds
* TPS Tranaction per seconds
* TPS = (Com_commit + Com_rollback) / seconds
* show global status like 'Com_commit';
* show global status like 'Com_rollback';
* 线程连接数
* show global status like 'Max_used_connections';
* show global status like 'Threads%';
* 最大连接数
* show variables like 'max_connections';
* Query Cache
* 查询缓存用于缓存select查询结果
* 当下次接收到相同查询请求时 不再执行实际查询处理而直接返回结果
* 适用于大量查询　很少改变表中数据
* 开启 Query Cache ---> 修改my.cnf
* 将query_cache_size设置为具体的大小 具体大小是多少取决于查询的实际情况　但最好设置为1024的倍数 参考值32M
* 增加一行 query_cache_type=0/1/2
* 如果设置为1 将会缓存所有的结果　除非你的select语句使用SQL_NO_CACHE禁用了查询缓存
* 如果设置为2 则只缓存在select语句中通过SQL_CACHE指定需要缓存的查询
* Query Cache 命中率
* show status like 'Qcache%';
* Query_cache_hits = (Qcache_hits / (Qcache_hits + Qcache_inserts)) * 100%;
* 锁定状态
* show global status like '%lock%';
* Table_locks_waited / Table_locks_immediate 值越大代表表锁造成的阻塞越严重
* innodb_row_lock_waits innodb行锁 太大可能是间隙锁造成的
* 主从延迟
* 查询主从延时时间 show slave status

## MySQL 慢查询
* 慢查询定义
* 执行速度超过定义的时间的查询
* 不同的系统定义不同的慢查询指标
* 开启慢查询
* 编辑/etc/my.cnf 在[mysqld]域中添加
* 开启慢查询 slow_query_log = 1
* 慢查询日志路径 slow_query_log_file=/data/mysql/slow.log
* 慢查询的时长　long_query_time = 1
* 末使用索引的查询也被记录到慢查询日志中
* log_queries_not_using_indexes = 1

## 慢查询日志分析
* mysqldumpslow命令
* -s 是表示按照何种方式排序
* -t 是top n的意思　即为返回前面多少条的数据
* -g 后边可以写一个正则匹配模式　大小写不敏感的
* 慢查询分析举例
* 得到返回记录集最多的10个SQL  mysqldumpslow -s r -t 10 slow.log
* 得到访问次数最多的10个SQL    mysqldumpslow -s c -t 10 slow.log
* 等到按照时间排序的前10条里面含有左连接的查询语句  mysqldumpslow -s t-t 10 -g "left join" slow.log

## explain执行计划
* explain select语句
* explain select * from users;
* explain返回结果分析
* id select识别符　代表语句的执行顺序 一般在select嵌套查询时会不同
* id列数字越大越先执行　如果说数字一样大　那么就从上往下依次执行
* id列为null的就表示这是一个结果集　不需要使用它来进行查询
* select_type simple 表示不需要union操作或者不包含子查询的简单select查询　有连接查询时　外层的查询为simple 且只有一个
* primary 一个需要union操作或者含有子查询的select 位于最外层的单位查询的select_type即为primary 且只有一个
* union union连接的两个select查询　第一个查询是dervied派生表　除了第一个表外　第二个以后的表select_type都是union
* dependent union 与union一样 出现在union或union all语句中 但是这个查询要受到外部查询的影响
* union result 包含union的结果集 在union和union all语句中 因为它不需要参与查询　所以id字段为null
* subquery 除了from字句中包含的子查询外　其他地方出现的子查询都可能是subquery
* dependent subquery 与dependent union 类似  表示这个subquery的查询要受到外部表查询的影响
* derived from字句中出现的子查询　也叫做派生表 其他数据库中可能叫做内联视图或嵌套select
* table 显示的查询表名　如果查询使用了别名　那么这里显示的是别名
* 如果不涉及对数据表的操作　那么这显示为null
* 如果显示为尖括号括起来的<derived N> 就表示这个是临时表
* type 依次从好到差 system const eq_ref ref fulltext ref_or_null unique_subquery index_subquery range index_merge index ALL
* 除了all之外 其他的type都可以使用到索引 除了index_merge之外 其他的type只可以用到一个索引
* system 表中只有一行数据或者是空表　且只能用于myisam和memory表 如果是innodb引擎表 type列在这个情况通常都是all或者index
* const 使用唯一索引或者主键　返回记录一定是1行记录的等值where条件时 通常type是const 其他数据库也叫做唯一索引扫描
* eq_ref 出现在要连接过个表的查询计划中　驱动表只返回一行数据　且这行数据是第二个表的主键或者唯一索引　且必须为not null 唯一索引和主键是多列时　只有所有的列都用作比较时才会出现eq_ref
* ref 不像eq_red那样要求连接顺序　也没有主键和唯一索引的要求　只要使用相等条件检索时就可能出现　常见与辅助索引的等值查找　或者多列主键　唯一索引中　使用第一个列之外的列作为等值查找也会出现　总之　返回数据不唯一的等值查找就可能出现
* fulltext 全文索引检索　要注意　全文索引的优先级很高　　若全文索引和普通索引同时存在时 mysql不管代价　优先选择使用全文索引
* ref_or_null 与ref方法类似　只是增加了null值的比较　实际用的不多
* unique_subquery 用于where中的in形式子查询　子查询返回不重复值唯一值
* index_subquery 用于in形式子查询使用到了辅助索引或者in常数列表　子查询可能返回重复值　可以使用索引将子查询去重
* range 索引范围扫描　常见于使用>,<,is null, between, in, like 等运算符的查询中
* index_merge 表示查询使用了两个以上的索引　最后取交集或者并集　常见and or的条件使用了不同的索引　官方排序这个在ref_or_null之后　但是实际上由于要读取所有索引　性能可能大部分时间都不如range
* index 索引全表扫描　把索引从头到尾扫一遍　常见于使用索引列就可以处理不需要读取数据文件的查询　可以使用索引排序或者分组的查询
* all 这个就是全表扫描数据文件　然后再在server层进行过滤返回符合要求的记录
* possible_keys 查询可能使用到的索引都会在这里列出来
* key 查询真正使用到的索引　select type为index_merge时　这里可能出现两个以上的索引　其他的select_type这里只会出现一个
* key_len 用于处理查询的索引长度　如果是单列索引　那就整个索引长度算进去　如果是多列索引　那么查询不一定都能使用到所有的列　具体使用到了多少个列的索引　这里就会计算进去　没有使用到的列　这里不会计算进去
* ref 如果是使用的常数等值查询　这里会显示const  如果是连接查询　被驱动表的执行计划这里会显示驱动表的关联字段 如果是条件使用了表达式或者函数　或者条件列发生了内部隐匿转换　这里可能显示为func
* rows 这里执行计划中估算的扫描行数　不是精确值
* extra 常见返回
* distinct 在select部分使用了distinct关键字
* no tables used 不带from字句的查询或者from dual查询
* using filesort 排序时无法使用到索引时　就会出现这个　常见于order by和group by语句中
* using index 查询时不需要回表查询　直接通过索引就可以获取查询的数据
* using intersect 表示使用and的各个索引的条件时　该信息表示是从处理结果获取交集
* using union 表示使用or连接各个使用索引的条件时　该信息表示处理结果获取并集
* using where 表示存储引擎返回的记录并不是所有的都满足查询条件　需要在server层进行过滤

## MySQL索引
* 类型　主键索引　全文索引　唯一索引　组合索引　普通索引
* 主键索引 它是一种特殊的唯一索引　不允许有空值
* 一般是在建表的时候同时创建主键索引
* 唯一索引　索引列的值必须唯一　但允许有空值
* 普通索引　最基本的索引　它没有任何限制
* 全文索引 fulltext是一种只适用于MyISAM表的一个索引类型
* 被索引列的数据类型只能是以下三种的组合char varchar text
* MySQL是通过match()和against()这两个函数来实现它的全文索引查询的功能
* 组合索引 也叫多列索引　就是在多列上同时创建索引　使用多列的组合值唯一　创建组合索引的好处是比分别创建多个单列索引的查询速度要快很多
* 组合索引 创建遵循 最左前缀 规则
* 如三列 id, name, age创建组合索引　则相当于分别创建了id, name, age  id, name   id这三个索引

## 索引创建规范
* 索引　可以提高查询效率但也会降低插入和更新的速度并占用磁盘空间
* 在插入与更新数据时　要重写索引文件
* 单张表中索引数量不超过５个
* 单个索引中的字段数不超过5个
* 不使用更新频繁地列作为主键
* 合理创建组合索引(避免冗余)
* 不在低基数列上建立索引 例如 '性别'
* 不在索引列进行数学运算和函数运算　会使索引失效
* 不使用%前导的查询　如like "%xxx" 无法使用索引
* 不使用反向查询　如not in / not like 无法使用索引　导致全表扫描
* 选择越小的数据类型越好　因为通常越小的数据类型通常在磁盘　内存　cpu　缓存中 占用的空间很少　处理起来更快
* 在经常需要排序order by 分组group by 和distinct列上加索引(单独order by用不了索引　索引考虑加where或加limit)
* 在表与表的连接条件上加上索引　可以加快连接查询的速度
* 使用短索引　如果你的一个字段是char32或者int32 在创建索引的时候指定前缀长度比如前10个字符(前提是多数值是唯一的)那么短索引可以提高查询速度　并且可以减少磁盘的空间　也可以减少I/O操作

## MySQL存储引擎
* MyISAM
* 优点　读的性能比InnoDB高很多
* 索引与数据分开　使用了压缩　从而提高了内存使用率
* 缺点　不支持事务
* 写入数据时　直接锁表
* InnoDB
* 优点
* 支持事务 支持外键 运行行级锁
* 缺点
* 不支持fulltext索引(全文索引)
* 行级锁并不绝对　当不确定扫描范围时　锁全表
* 索引与数据是紧密捆绑的　没有使用压缩导致体积庞大

## 实时监控orzdba

## 天免LEPUS简介
* 为所有数据库管理者　互联网企业数据库监控而设计
* 无需部署Agent 轻松监控1000+数据库实例 完善灵活的告警配置　详细的性能分析指标
