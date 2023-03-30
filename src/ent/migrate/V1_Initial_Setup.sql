
drop table if exists qy_member;
create table qy_member
(
    id         int(11) unsigned not null primary key auto_increment,
    email      varchar(64)      not null default '' comment '邮箱',
    phone      varchar(11)      not null default '' comment '手机号码',
    account    varchar(32)      not null default '' comment '账号',
    avatar     varchar(32)      not null default '' comment '头像',
    nickname   varchar(255)     not null default '' comment '名称',
    state      varchar(32)      not null default 'active' comment '状态: un_active 待激活, active 已激活, disabled 停用',
    created_at datetime                  default current_timestamp comment '创建时间',
    updated_at datetime         null comment '改时间',
    deleted_at datetime         null comment '删除时间',
    unique IDX_email (email),
    index IDX_phone_email (phone, email)
)
    default charset utf8mb4
    collate utf8mb4_general_ci
    engine myisam
    comment '用户数据表';

drop table if exists qy_member_password;
create table qy_member_password
(
    id         int(11) unsigned not null primary key auto_increment,
    member_id  int(11) unsigned not null comment '用户id',
    passwd     varchar(32)      not null default '' comment '密码',
    uuid       varchar(32)      not null default '' comment 'uuid',
    created_at datetime                  default current_timestamp comment '创建时间',
    updated_at datetime         null comment '修改时间',
    deleted_at datetime         null comment '删除时间',
    unique IDX_member_id (member_id)

)
    collate utf8mb4_general_ci
    default charset utf8mb4
    engine myisam
    comment '用户授权验证'
;
drop table if exists qy_member_authorize_record;
create table qy_member_authorize_record
(
    id            int(11) unsigned not null primary key auto_increment,
    member_id     int(11) unsigned not null comment '用户id',
    channel       varchar(32)      not null default 'passwd' comment '登录渠道 passwd 密码登录',
    device        varchar(32)      not null default '' comment '登录设备',
    device_detail varchar(32)      not null default '' comment '登录设备详情',
    ipv4          varchar(64)      not null default '' comment '登录ipv4',
    ipv6          varchar(128)     not null default '' comment '登录ipv6',
    snapshot      varchar(256)              default null comment '快照',
    created_at    datetime                  default current_timestamp comment '时间时间',
    index IDX_member_id (member_id)
)
    collate utf8mb4_general_ci
    default charset utf8mb4
    engine myisam
    comment '用户登录记录'
;

drop table if exists qy_member_role_relation;

create table qy_member_role_relation
(
    id         int(11) unsigned not null primary key auto_increment,
    member_id  int(11) unsigned not null comment '用户id',
    role_id    int(11) unsigned not null comment '角色id',
    created_at datetime default current_timestamp comment '创建时间',
    deleted_at datetime         null comment '删除时间'
)
    collate utf8mb4_general_ci
    default charset utf8mb4
    engine innodb
    comment '关联权限'
;

drop table if exists qy_role;

create table qy_role
(
    id         int(11) unsigned not null primary key auto_increment,
    role_name  varchar(32)      not null default '' comment '名称',
    state      varchar(32)      not null default 'on' comment '状态: on 开启, off 关闭',
    created_at datetime                  default current_timestamp comment '创建时间',
    updated_at datetime         null comment '修改时间',
    deleted_at datetime         null comment '删除时间'
)
    collate utf8mb4_general_ci
    default charset utf8mb4
    engine myisam
    comment '角色'
;

drop table if exists qy_role_permission_relation;
create table qy_role_permission_relation
(
    id                  int(11) unsigned not null primary key auto_increment,
    role_id             int(11) unsigned not null default 0 comment '角色id',
    permission_group_id int(11) unsigned not null default 0 comment '权限id',
    created_at          datetime                  default current_timestamp comment '创建时间',
    updated_at          datetime         null comment '修改时间',
    deleted_at          datetime         null comment '删除时间',
    index IDX_role_permission (role_id, permission_group_id)
)
    collate utf8mb4_general_ci
    default charset utf8mb4
    engine myisam
    comment '角色关联权限'
;

drop table if exists qy_permission_group;
create table qy_permission_group
(
    id              int(11) unsigned not null primary key auto_increment,
    permission_name varchar(32)      not null default '' comment '路由名称',
    icon            varchar(200) comment '前端图标',
    sort            tinyint(8) comment '菜单排序',
    left_index      int(11) unsigned not null comment '左索引',
    right_index     int(11) unsigned not null comment '右索引',
    state           varchar(32)      not null default 'on' comment '状态: on 开启, off 关闭',
    created_at      datetime                  default current_timestamp comment '创建时间',
    updated_at      datetime         null comment '修改时间',
    deleted_at      datetime         null comment '删除时间'
)
    collate utf8mb4_general_ci
    default charset utf8mb4
    engine innodb
    comment '权限路由组'
;


drop table if exists qy_permission_router_relation;

create table qy_permission_router_relation
(
    id                  int(11) unsigned not null primary key auto_increment,
    permission_group_id int(11) unsigned not null default 0 comment '权限组id',
    router_id           int(11) unsigned not null default 0 comment '路由id',
    created_at          datetime                  default current_timestamp comment '创建时间',
    updated_at          datetime         null comment '修改时间',
    deleted_at          datetime         null comment '删除时间'
)
    collate utf8mb4_general_ci
    default charset utf8mb4
    engine myisam
    comment '权限路由关联'
;


drop table if exists qy_router;
create table qy_router
(
    id          int(11) unsigned not null primary key auto_increment,
    route_name  varchar(32)      not null default '' comment '路由名称',
    route       varchar(128)     not null default '' comment '路由',
    description varchar(32)      not null default '' comment '描述',
    created_at  datetime                  default current_timestamp comment '创建时间',
    updated_at  datetime         null comment '修改时间',
    deleted_at  datetime         null comment '删除时间',
    unique IDX_route (route)
)
    collate utf8mb4_general_ci
    default charset utf8mb4
    engine myisam
    comment '权限路由'
;

