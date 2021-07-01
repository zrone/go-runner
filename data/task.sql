-- 任务链
CREATE TABLE `task`
(
    `id`          BIGINT(20) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    `status`      TINYINT(3)          NOT NULL DEFAULT 10 COMMENT '10 AUDIT；20 UNCHECK 30 PENDING；40 RUNNING；50 SUCCESS；60 FAILED 70 CANCELED',
    `create_at`   INT(10)             NOT NULL DEFAULT 0 COMMENT '创建时间',
    `start_at`    INT(10)             NOT NULL DEFAULT 0 COMMENT '开始时间',
    `retry_count` INT(10)             NOT NULL DEFAULT 0 COMMENT '重试次数',
    `retry_at`    INT(10)             NOT NULL DEFAULT 0 COMMENT '重试时间',
    `process`     INT(10)             NOT NULL DEFAULT 1 COMMENT '全部流程',
    `stage`       INT(10)             NOT NULL DEFAULT 1 COMMENT '当前进度',
    `relation_id` BIGINT(20) unsigned NOT NULL COMMENT '关联ID',
    CONSTRAINT `tk_relation_id` FOREIGN KEY (`relation_id`) REFERENCES `relation` (`id`) ON UPDATE CASCADE
) ENGINE = InnoDB
  AUTO_INCREMENT = 10000
  DEFAULT CHARSET = utf8mb4 COMMENT '任务链';
