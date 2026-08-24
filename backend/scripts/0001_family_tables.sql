-- 家属端小程序所需新增表
-- 说明：family_member 表已由 gentity 生成，无需重复创建。

-- 家属账号表：一个手机号对应一个账号，可绑定多位老人（通过 family_member.phone 关联）
CREATE TABLE IF NOT EXISTS `family_account` (
  `id`          BIGINT       NOT NULL AUTO_INCREMENT COMMENT '主键',
  `phone`       VARCHAR(20)  NOT NULL COMMENT '家属手机号（登录账号）',
  `pass`        VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '登录密码（md5）',
  `openid`      VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '微信 openid（充值支付前置）',
  `create_time` DATETIME     DEFAULT NULL COMMENT '创建时间',
  `update_time` DATETIME     DEFAULT NULL COMMENT '更新时间',
  `del_flag`    TINYINT      NOT NULL DEFAULT 0 COMMENT '逻辑删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='家属账号表';

-- 家属充值订单表：充值直接加款到对应老人 elder.balance
CREATE TABLE IF NOT EXISTS `family_recharge` (
  `id`          BIGINT       NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_no`    VARCHAR(64)  NOT NULL COMMENT '商户订单号',
  `phone`       VARCHAR(20)  NOT NULL DEFAULT '' COMMENT '家属手机号',
  `elder_id`    BIGINT       NOT NULL COMMENT '充值到哪位老人',
  `amount`      BIGINT       NOT NULL DEFAULT 0 COMMENT '充值金额（单位：分）',
  `status`      TINYINT      NOT NULL DEFAULT 0 COMMENT '0-待支付 1-已支付 2-已关闭',
  `prepay_id`   VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '微信预支付 id',
  `create_time` DATETIME     DEFAULT NULL COMMENT '创建时间',
  `update_time` DATETIME     DEFAULT NULL COMMENT '更新时间',
  `del_flag`    TINYINT      NOT NULL DEFAULT 0 COMMENT '逻辑删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_order_no` (`order_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='家属充值订单表';
