-- 工作流 bind 用户
CREATE TABLE `relation`
(
    `id`          BIGINT(20) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    `symbol_type` TINYINT(3)          NOT NULL DEFAULT 1 COMMENT '1 developer; 2 deployer; 3 manager',
    `symbol`      VARCHAR(32)                  DEFAULT NULL COMMENT '标识',
    `workflow_id` BIGINT(20) unsigned NOT NULL COMMENT '工作流ID',
    `manager_id`  BIGINT(20) unsigned NOT NULL COMMENT '用户ID',
    KEY `af_symbol` (`symbol`) USING BTREE,
    CONSTRAINT `af_workflow_id` FOREIGN KEY (`workflow_id`) REFERENCES `workflow` (`id`) ON UPDATE CASCADE,
    CONSTRAINT `af_manager_id` FOREIGN KEY (`manager_id`) REFERENCES `manager` (`id`) ON UPDATE CASCADE
) ENGINE = InnoDB
  AUTO_INCREMENT = 10000
  DEFAULT CHARSET = utf8mb4 COMMENT '工作流';
