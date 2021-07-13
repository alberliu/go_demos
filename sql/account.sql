CREATE TABLE `account`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `user_id`     bigint(20) unsigned NOT NULL COMMENT '用户id',
    `code`        varchar(64) NOT NULL COMMENT '账户识别码',
    `balance`     bigint(20) NOT NULL COMMENT '余额',
    `status`      tinyint(4) NOT NULL COMMENT '状态',
    `create_time` datetime    NOT NULL COMMENT '创建时间',
    `update_time` datetime    NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id_code` (`user_id`,`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户账户';

CREATE TABLE `account_item`
(
    `id`                  bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `account_id`          bigint(20) unsigned NOT NULL COMMENT '账号id',
    `account_transfer_id` bigint(20) unsigned NOT NULL COMMENT '账户流水id',
    `amount`              int(20) NOT NULL COMMENT '金额',
    `balance`             int(20) NOT NULL COMMENT '余额快照',
    `remark`              varchar(256) NOT NULL COMMENT '账户标记',
    `create_time`         datetime     NOT NULL COMMENT '创建时间',
    `update_time`         datetime     NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY                   `idx_account_id` (`account_id`),
    KEY                   `idx_account_transfer_id` (`account_transfer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户账户明细';

CREATE TABLE `account_transfer`
(
    `id`              bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `from_account_id` bigint(20) unsigned NOT NULL COMMENT '出账账号',
    `to_account_id`   bigint(20) unsigned NOT NULL COMMENT '入账账号',
    `amount`          bigint(20) NOT NULL COMMENT '金额',
    `object_type`     varchar(64)  NOT NULL COMMENT '对象类型',
    `object_id`       bigint(20) unsigned NOT NULL COMMENT '对象id',
    `code`            varchar(64)  NOT NULL COMMENT '标记码',
    `remark`          varchar(256) NOT NULL COMMENT '标记',
    `status`          tinyint(4) NOT NULL COMMENT '状态',
    `create_time`     datetime     NOT NULL COMMENT '创建时间',
    `update_time`     datetime     NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY               `idx_from_account_id` (`from_account_id`) USING BTREE,
    KEY               `idx_to_account_id` (`to_account_id`) USING BTREE,
    KEY               `idx_object_type_object_id` (`object_type`,`object_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户账户交易流水';