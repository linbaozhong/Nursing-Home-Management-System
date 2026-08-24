-- ============================================================================
-- 床位床型 / 房间设施 / 楼层回写 所需表结构变更
-- 用途：让“房间→床位(床型)”与“房间→设施(物资分类)”与后端已存在的物资体系打通。
--
-- 约定：
--   material_type.kind   : 1 = 床型分类项（如 标准床/智能床/电动护理床）
--                         99 = 设施/其他分类项（如 餐桌/轮椅/空调/电视）
--   bed.bed_type_id      : 关联 material_type.id，指向 kind=1 的床型项
--   room_material        : 房间-设施 关联表（room.id <-> material_type.id，kind=99 的设施项）
--
-- 说明：
--   本脚本是数据库变更的“权威定义”，执行后请用 gentity 重新生成对应 .gen.go。
--   全部 .gen.go 由工具从以下表结构生成，勿手改。
-- ============================================================================

-- 1) 物资分类表：新增 kind 列（区分床型/设施）
ALTER TABLE `material_type`
  ADD COLUMN `kind` tinyint(4) NOT NULL DEFAULT 99 COMMENT '分类用途：1=床型，99=设施/其他' AFTER `name`;

-- 2) 床位表：新增 bed_type_id（关联 material_type 的床型项）
ALTER TABLE `bed`
  ADD COLUMN `bed_type_id` bigint(20) DEFAULT NULL COMMENT '床型编号(关联 material_type.id，kind=1)' AFTER `name`;

-- 3) 新建 房间-设施 关联表
DROP TABLE IF EXISTS `room_material`;
CREATE TABLE `room_material` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `tenant_id` bigint(20) NOT NULL DEFAULT 1 COMMENT '租户编号',
  `room_id` bigint(20) NOT NULL COMMENT '房间编号',
  `material_type_id` bigint(20) NOT NULL COMMENT '设施(物资分类)编号，kind=99',
  `del_flag` varchar(2) NOT NULL DEFAULT 'N' COMMENT '删除状态(Y/N)',
  `create_id` bigint(20) NOT NULL COMMENT '创建人编号',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint(20) NOT NULL COMMENT '修改人编号',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_room` (`room_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ============================================================================
-- 补充默认数据（方便联调；可按需调整）
-- ============================================================================
-- 床型分类项（kind=1）
-- INSERT INTO `material_type`(`tenant_id`,`name`,`kind`,`del_flag`,`create_id`,`create_time`,`update_id`,`update_time`)
-- VALUES (1,'标准床',1,'N',1,NOW(),1,NOW()),
--        (1,'智能床',1,'N',1,NOW(),1,NOW()),
--        (1,'电动护理床',1,'N',1,NOW(),1,NOW());
-- 设施分类项（kind=99）
-- INSERT INTO `material_type`(`tenant_id`,`name`,`kind`,`del_flag`,`create_id`,`create_time`,`update_id`,`update_time`)
-- VALUES (1,'餐桌',99,'N',1,NOW(),1,NOW()),
--        (1,'轮椅',99,'N',1,NOW(),1,NOW()),
--        (1,'空调',99,'N',1,NOW(),1,NOW()),
--        (1,'电视',99,'N',1,NOW(),1,NOW());
