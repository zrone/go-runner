-- 用户
CREATE TABLE `manager`
(
    `id`        BIGINT(20) unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '用户ID',
    `role`   TINYINT(4)          NOT NULL DEFAULT 0 COMMENT '角色类型 1：sa; 2 user',
    `username`  VARCHAR(32)         NOT NULL DEFAULT '' COMMENT '账号',
    `account`   VARCHAR(32)         NOT NULL DEFAULT '' COMMENT '账号名',
    `salt`      VARCHAR(64)         NOT NULL DEFAULT '' COMMENT '加密盐',
    `create_at` INT(10)             NOT NULL DEFAULT 0 comment '创建时间',
    `delete_at` INT(10)             NOT NULL DEFAULT 0 comment '删除时间',
    `update_at` INT(10)             NOT NULL DEFAULT 0 comment '删除时间',
    KEY `username` (`username`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 10000
  DEFAULT CHARSET = utf8mb4 COMMENT '用户';
