-- 成员
CREATE TABLE `member`
(
    `id`   TINYINT(4)  NOT NULL AUTO_INCREMENT COMMENT '成员ID',
    `name` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '名称'
) ENGINE = InnoDB
  AUTO_INCREMENT = 10000
  DEFAULT CHARSET = utf8mb4 COMMENT '成员';


-- data
INSERT INTO `role` (`name`)
values ("developer"), ("deployer"), ("manager");

# 工作流
# workflow

# test 主动发布
# developer -> deployer -> test

# dev 自动发布
# release

# release、notification 主动发布
# developer -> deployer -> manager -> deployment -> release
