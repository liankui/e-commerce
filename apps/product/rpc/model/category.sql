CREATE TABLE `category`
(
    `id`          smallint(6) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类id',
    `parentid`    smallint(6) NOT NULL DEFAULT 0 COMMENT '父类别id当id=0时说明是根节点,一级类别',
    `name`        varchar(50) NOT NULL DEFAULT '' COMMENT '类别名称',
    `status`      tinyint(4) NOT NULL DEFAULT 1 COMMENT '类别状态1-正常,2-已废弃',
    `create_time` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品类别表';
