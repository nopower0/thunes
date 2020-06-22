create table wallet
(
    add_time    timestamp          not null default current_timestamp,
    update_time timestamp          not null default current_timestamp on update current_timestamp,
    uid         bigint primary key not null comment 'user id',
    sgd         bigint             not null default 0 comment 'SGD balance'
) engine InnoDB
  charset utf8mb4;

alter table wallet
    add index idx_wallet_add_time (add_time),
    add index idx_wallet_update_time (update_time);