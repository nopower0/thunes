create table user
(
    add_time    timestamp                         not null default current_timestamp,
    update_time timestamp                         not null default current_timestamp on update current_timestamp,
    uid         bigint primary key auto_increment not null comment 'user id',
    username    varchar(100)                      not null default '' comment 'username',
    password    varchar(64)                       not null default '' comment 'password'
) engine InnoDB
  charset utf8mb4;

alter table user
    add index idx_user_add_time (add_time),
    add index idx_user_update_time (update_time),
    add constraint unique uniq_user_username (username);