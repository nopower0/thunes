create table transfer_history
(
    add_time    timestamp                         not null default current_timestamp,
    update_time timestamp                         not null default current_timestamp on update current_timestamp,
    id          bigint primary key auto_increment not null comment 'user id',
    from_uid    bigint                            not null default 0 comment 'from UID',
    to_uid      bigint                            not null default 0 comment 'to UID',
    amount      bigint                            not null default 0 comment 'amount'
) engine InnoDB
  charset utf8mb4;

alter table transfer_history
    add index idx_transfer_history_add_time (add_time),
    add index idx_transfer_history_update_time (update_time);