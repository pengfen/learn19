create table `article` (
    `id` int not null auto_increment primary key,
    `aname` varchar(255) not null default '' comment '文章标题',
    `atime` datetime not null comment '创建时间',
    `account` int not null defautl 0 comment '',
    `acontent` varchar(255) not null default 0 comment '文章内容',
    `aimg` varchar(255) not null default '' comment '文章关联图'
) engine=innodb default charset=utf8 comment='文章表';

create table `article_type` (
    `id` int not null auto_increment primary key,
    `tname` varchar(255) not null default '' comment '文章分类名'
) engine=innodb default charset=utf8 comment='文章分类表';

create table `user` (
    `id` int not null auto_increment primary key,
    `name` varchar(255) not null default '' comment '用户名',
    `pwd` varchar(255) not null default '' comment '用户密码'
) engine=innodb default charset=utf8 comment='用户表';

create table `userinfo` (
`uid` int not null auto_increment primary key,
`username` varchar(64) not null default '' comment '用户名',
`departname` varchar(64) not null default '' comment '部门名',
`create_time` int(11) not null default 0 comment '创建时间',
) engine=innodb default charset=utf8 comment='用户信息表';

create table `userdetail` (
`uid` int not null default 0 primary key,
`intro` text null comment '用户介绍',
`profile` text null comment '用户详情'
) engine=innodb default charset=utf8 comment='用户详情表';