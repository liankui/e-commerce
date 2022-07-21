CREATE TABLE `product_operation`
(
    `id`          bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品运营id',
    `product_id`  int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品id',
    `status`      tinyint(4) NOT NULL DEFAULT 0 COMMENT '运营商品状态 0-下线 1-上线',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY           `ix_productid` (`product_id`),
    KEY           `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品运营表';
