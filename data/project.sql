-- 项目
CREATE TABLE `project`
(
    `id`         BIGINT(20) unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '项目ID',
    `name`       VARCHAR(191) NOT NULL DEFAULT '' COMMENT '项目名称',
    `symbol`     TINYINT(3) NOT NULL DEFAULT 1 COMMENT '1 github; 2 gitee; 3 gitlab',
    `manager_id` BIGINT(20) unsigned NOT NULL COMMENT '创建用户ID',
    `create_at`  INT(10) NOT NULL DEFAULT 0 comment '创建时间',
    `delete_at`  INT(10) NOT NULL DEFAULT 0 comment '删除时间',
    `update_at`  INT(10) NOT NULL DEFAULT 0 comment '删除时间',
    KEY          `name` (`name`) USING BTREE,
    KEY          `symbol` (`symbol`) USING BTREE,
    CONSTRAINT `manager_id` FOREIGN KEY (`manager_id`) REFERENCES `manager` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT '项目';


-- 主动发布
-- 自动发布