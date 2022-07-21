CREATE TABLE `product`
(
    `id`          bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品id',
    `cateid`      smallint(6) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类别Id',
    `name`        varchar(100)   NOT NULL DEFAULT '' COMMENT '商品名称',
    `subtitle`    varchar(200)            DEFAULT NULL DEFAULT '' COMMENT '商品副标题',
    `images`      text COMMENT '图片地址,json格式,扩展用',
    `detail`      text COMMENT '商品详情',
    `price`       decimal(20, 2) NOT NULL DEFAULT 0 COMMENT '价格,单位-元保留两位小数',
    `stock`       int(11) NOT NULL DEFAULT 0 COMMENT '库存数量',
    `status`      int(6) NOT NULL DEFAULT 1 COMMENT '商品状态.1-在售 2-下架 3-删除',
    `create_time` timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY           `ix_cateid` (`cateid`),
    KEY           `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';
