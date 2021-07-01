-- 工作流
CREATE TABLE `workflow`
(
    `id`               BIGINT(20) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    `name`             VARCHAR(32)         NOT NULL DEFAULT '' COMMENT '名称',
    `type`             TINYINT(3)          NOT NULL DEFAULT 1 COMMENT '1 test; 2 dev; 3 release',
    `developer_symbol` VARCHAR(32)                  DEFAULT NULL COMMENT '开发者标识',
    `deployer_symbol`  VARCHAR(32)                  DEFAULT NULL COMMENT '部署者标识',
    `manager_symbol`   VARCHAR(32)                  DEFAULT NULL COMMENT '管理员标识',
    `create_at`        INT(10)             NOT NULL DEFAULT 0 COMMENT '创建时间',
    `delete_at`        INT(10)             NOT NULL DEFAULT 0 COMMENT '删除时间',
    `update_at`        INT(10)             NOT NULL DEFAULT 0 COMMENT '删除时间',
    `project_id`       BIGINT(20) unsigned NOT NULL COMMENT '项目ID',
    `before`           longtext COMMENT '部署前置指令集',
    `after`            longtext COMMENT '部署后置指令集',
    KEY `type` (`type`) USING BTREE,
    KEY `developer_symbol` (`developer_symbol`) USING BTREE,
    KEY `deployer_symbol` (`deployer_symbol`) USING BTREE,
    KEY `manager_symbol` (`manager_symbol`) USING BTREE,
    CONSTRAINT `project_id` FOREIGN KEY (`project_id`) REFERENCES `project` (`id`) ON UPDATE CASCADE
) ENGINE = InnoDB
  AUTO_INCREMENT = 10000
  DEFAULT CHARSET = utf8mb4 COMMENT '工作流';
