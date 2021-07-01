-- 自动化部署
CREATE TABLE `internal_deploy`
(
    `id`     BIGINT(20) unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '用户ID',
    `symbol` VARCHAR(191)        NOT NULL DEFAULT '' COMMENT '项目标识',
    `secret` VARCHAR(191)        NOT NULL DEFAULT '' COMMENT '加密盐',
    `path`   VARCHAR(191)        NOT NULL DEFAULT '' COMMENT '项目目录',
    `auth`   VARCHAR(255)        DEFAULT NULL COMMENT '认证信息',
    KEY `symbol` (`symbol`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 10000
  DEFAULT CHARSET = utf8mb4 COMMENT '自动化部署';


-- ssh
-- user pwd
-- publickey
