-- MySQL dump 10.13  Distrib 8.0.26, for Win64 (x86_64)
--
-- Host: 123.56.5.53    Database: snow_gerocomium
-- ------------------------------------------------------
-- Server version	8.0.36

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
SET @MYSQLDUMP_TEMP_LOG_BIN = @@SESSION.SQL_LOG_BIN;
SET @@SESSION.SQL_LOG_BIN= 0;

--
-- GTID state at the beginning of the backup 
--

SET @@GLOBAL.GTID_PURGED=/*!80000 '+'*/ 'a0c872c6-8513-11ef-a440-00163e0e83db:1-2226338,
fbab4a2c-a8f6-11f0-8e6c-00163e32338b:1-5322203';

--
-- Table structure for table `accident`
--

DROP TABLE IF EXISTS `accident`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `accident` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `staff_id` bigint unsigned NOT NULL COMMENT '值班护工id',
  `occur_date` datetime NOT NULL COMMENT '发生时间',
  `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '事故描述',
  `picture` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '事故图片',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_accident_tenant` (`tenant_id`,`elder_id`,`occur_date`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accident`
--

LOCK TABLES `accident` WRITE;
/*!40000 ALTER TABLE `accident` DISABLE KEYS */;
INSERT INTO `accident` VALUES (1,1,1,5,'2022-12-13 00:00:00','摔倒','http://127.0.0.1:9001/upload/img/20230401/1642192054401458176_屏幕截图_20221210_093317.png',1,1,'2023-02-04 19:24:16',1,'2023-04-01 23:47:53'),(2,1,2,6,'2023-04-06 00:00:00','摔倒','http://127.0.0.1:9001/upload/img/20230401/1642190705605500928_屏幕截图_20221211_143319.png',1,1,'2023-04-01 22:40:27',1,'2023-04-01 23:46:58'),(3,1,4,1,'2023-04-25 00:00:00','摔倒','http://127.0.0.1:9001/upload/img/20230426/1651132340481650688_56.jpg',0,1,'2023-04-26 15:53:28',1,'2023-04-26 15:53:28');
/*!40000 ALTER TABLE `accident` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `active`
--

DROP TABLE IF EXISTS `active`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `active` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `type_id` bigint unsigned NOT NULL COMMENT '活动类别id',
  `theme` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '活动主题',
  `name` varchar(25) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '活动名称',
  `content` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '活动内容',
  `address` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '活动地点',
  `organizer` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '组织者姓名',
  `phone` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '组织者电话',
  `active_date` datetime NOT NULL COMMENT '活动日期',
  `active_picture` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '活动图片',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `active`
--

LOCK TABLES `active` WRITE;
/*!40000 ALTER TABLE `active` DISABLE KEYS */;
INSERT INTO `active` VALUES (1,1,1,'关爱老人','关爱老人文艺汇演','内容','礼堂','张三','13546574576','2023-01-10 09:48:24','url',0,1,'2023-01-10 09:48:37',1,'2023-01-10 13:27:32'),(5,1,1,'关爱老人','文艺汇演','文艺汇演内容','礼堂','张三','13546574657','2022-12-13 00:00:00','url',0,1,'2023-01-10 13:38:53',1,'2023-01-12 09:09:53');
/*!40000 ALTER TABLE `active` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `active_participant`
--

DROP TABLE IF EXISTS `active_participant`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `active_participant` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `active_id` bigint unsigned NOT NULL COMMENT '活动id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `active_participant`
--

LOCK TABLES `active_participant` WRITE;
/*!40000 ALTER TABLE `active_participant` DISABLE KEYS */;
INSERT INTO `active_participant` VALUES (1,1,1,1,1,'2023-01-10 12:57:34',1,'2023-01-10 12:57:37'),(3,1,5,1,1,'2023-01-12 09:09:53',1,'2023-01-12 09:09:53'),(4,1,5,2,1,'2023-01-12 09:09:53',1,'2023-01-12 09:09:53');
/*!40000 ALTER TABLE `active_participant` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `active_type`
--

DROP TABLE IF EXISTS `active_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `active_type` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '活动类型名称',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `active_type`
--

LOCK TABLES `active_type` WRITE;
/*!40000 ALTER TABLE `active_type` DISABLE KEYS */;
INSERT INTO `active_type` VALUES (1,1,'老年大学',0,1,'2023-01-04 00:44:56',1,'2023-01-04 00:47:11'),(2,1,'文艺演出',0,1,'2023-01-04 00:45:09',1,'2023-01-04 00:47:13');
/*!40000 ALTER TABLE `active_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `assessment`
--

DROP TABLE IF EXISTS `assessment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `assessment` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `elder_id` bigint unsigned DEFAULT NULL COMMENT '老人ID',
  `assessment_date` date NOT NULL COMMENT '评估日期',
  `assessment_type` tinyint NOT NULL COMMENT '评估类型：1-能力评估，2-健康评估',
  `scale_data` json DEFAULT NULL COMMENT '量表原始数据',
  `result_level` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '评估结果（护理等级）',
  `evaluator_id` bigint unsigned DEFAULT NULL COMMENT '评估人ID',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elder_id` (`elder_id`),
  KEY `idx_evaluator_id` (`evaluator_id`),
  CONSTRAINT `fk_assessments_elder` FOREIGN KEY (`elder_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_assessments_staff` FOREIGN KEY (`evaluator_id`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_assessments_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评估记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `assessment`
--

LOCK TABLES `assessment` WRITE;
/*!40000 ALTER TABLE `assessment` DISABLE KEYS */;
/*!40000 ALTER TABLE `assessment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `audit_log`
--

DROP TABLE IF EXISTS `audit_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `audit_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `table` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '被操作表名',
  `row_id` bigint unsigned NOT NULL COMMENT '被操作行主键id',
  `action` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '操作：create/update/delete',
  `operator_id` bigint unsigned NOT NULL COMMENT '操作员id',
  `operator_name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '操作员名称',
  `change_after` json DEFAULT NULL COMMENT '变更后整行快照(JSON)',
  `change_label` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '可读变更摘要(中文字段名)',
  `comment` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '业务备注',
  `create_time` datetime NOT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `idx_row` (`table`,`row_id`),
  KEY `idx_operator` (`operator_id`),
  KEY `idx_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='数据轨迹日志';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `audit_log`
--

LOCK TABLES `audit_log` WRITE;
/*!40000 ALTER TABLE `audit_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `audit_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `auth`
--

DROP TABLE IF EXISTS `auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `auth` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `parent_id` bigint unsigned NOT NULL COMMENT '父级id',
  `title` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '权限标题',
  `name` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '权限名称',
  `path` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '权限path',
  `icon` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '权限图标',
  `url` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '权限url',
  `type` varchar(5) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '权限类别（MENU/BTN）',
  `method` varchar(6) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '权限请求方式（GET/POST/PUT/DELETE）',
  `create_id` bigint unsigned DEFAULT NULL COMMENT '创建人id',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_id` bigint unsigned DEFAULT NULL COMMENT '修改人id',
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `auth`
--

LOCK TABLES `auth` WRITE;
/*!40000 ALTER TABLE `auth` DISABLE KEYS */;
INSERT INTO `auth` VALUES (1,0,'首页','Home','/home',NULL,'/home',NULL,NULL,NULL,NULL,NULL,NULL),(2,0,'营销管理','SolesManage','/soles',NULL,'/soles',NULL,NULL,NULL,NULL,NULL,NULL),(3,2,'咨询管理','CounselManage','/counsel',NULL,'/soles/counsel',NULL,NULL,NULL,NULL,NULL,NULL),(4,2,'意向客户','IntentionClient','/intention',NULL,'/soles/intention',NULL,NULL,NULL,NULL,NULL,NULL),(5,2,'预定管理','BookingManage','/booking',NULL,'/soles/booking',NULL,NULL,NULL,NULL,NULL,NULL),(6,0,'入住管理','CheckInManage','/check-in',NULL,'/check-in',NULL,NULL,NULL,NULL,NULL,NULL),(7,6,'床位全景','BedGlobal','/bed',NULL,'/check-in/bed',NULL,NULL,NULL,NULL,NULL,NULL),(8,6,'入住签约','EnterSign','/enter',NULL,'/check-in/enter',NULL,NULL,NULL,NULL,NULL,NULL),(9,6,'外出登记','LeaveSign','/leave',NULL,'/check-in/leave',NULL,NULL,NULL,NULL,NULL,NULL),(10,6,'来访登记','VisitSign','/visit',NULL,'/check-in/visit',NULL,NULL,NULL,NULL,NULL,NULL),(11,6,'事故登记','AccidentSign','/accident',NULL,'/check-in/accident',NULL,NULL,NULL,NULL,NULL,NULL),(12,6,'退住申请','CheckOut','/check-out',NULL,'/check-in/check-out',NULL,NULL,NULL,NULL,NULL,NULL),(13,0,'人员管理','PeopleManage','/people',NULL,'/people',NULL,NULL,NULL,NULL,NULL,NULL),(14,13,'长者档案','OldFile','/old',NULL,'/people/old',NULL,NULL,NULL,NULL,NULL,NULL),(15,13,'员工管理','StaffManage','/staff',NULL,'/people/staff',NULL,NULL,NULL,NULL,NULL,NULL),(16,13,'活动管理','ActivityManage','/activity',NULL,'/people/activity',NULL,NULL,NULL,NULL,NULL,NULL),(17,0,'服务管理','ServiceManage','/service',NULL,'/service',NULL,NULL,NULL,NULL,NULL,NULL),(18,17,'服务项目','ServiceProject','/project',NULL,'/service/project',NULL,NULL,NULL,NULL,NULL,NULL),(19,17,'护理等级','ServiceLevel','/level',NULL,'/service/level',NULL,NULL,NULL,NULL,NULL,NULL),(20,17,'服务预定','ServiceBook','/book',NULL,'/service/book',NULL,NULL,NULL,NULL,NULL,NULL),(21,0,'物资管理','ResourceManage','/resource',NULL,'/resource',NULL,NULL,NULL,NULL,NULL,NULL),(22,21,'物资信息','ResourceInfo','/info',NULL,'/resource/info',NULL,NULL,NULL,NULL,NULL,NULL),(23,21,'仓库设置','StorageSet','/set',NULL,'/resource/set',NULL,NULL,NULL,NULL,NULL,NULL),(24,21,'入库管理','StorageEnter','/enter',NULL,'/resource/enter',NULL,NULL,NULL,NULL,NULL,NULL),(25,21,'出库管理','StorageLeave','/leave',NULL,'/resource/leave',NULL,NULL,NULL,NULL,NULL,NULL),(26,21,'库存查询','StorageSearch','/search',NULL,'/resource/search',NULL,NULL,NULL,NULL,NULL,NULL),(27,0,'餐饮管理','FoodManage','/food',NULL,'/food',NULL,NULL,NULL,NULL,NULL,NULL),(28,27,'菜品管理','DishManage','/dish',NULL,'/food/dish',NULL,NULL,NULL,NULL,NULL,NULL),(29,27,'餐饮套餐','FoodList','/list',NULL,'/food/list',NULL,NULL,NULL,NULL,NULL,NULL),(30,27,'点餐','FoodOrder','/order',NULL,'/food/order',NULL,NULL,NULL,NULL,NULL,NULL),(31,0,'费用管理','FeeManage','/fee',NULL,'/fee',NULL,NULL,NULL,NULL,NULL,NULL),(32,31,'预存充值','StoredPay','/pay',NULL,'/fee/pay',NULL,NULL,NULL,NULL,NULL,NULL),(33,31,'消费记录','FeeRecord','/record',NULL,'/fee/record',NULL,NULL,NULL,NULL,NULL,NULL),(34,31,'退住费用审核','FeeAudit','/audit',NULL,'/fee/audit',NULL,NULL,NULL,NULL,NULL,NULL),(35,0,'基础数据配置','BaseDataSet','/base',NULL,'/base',NULL,NULL,NULL,NULL,NULL,NULL),(36,35,'营销','Marketing','/marketing',NULL,'/base/marketing',NULL,NULL,NULL,NULL,NULL,NULL),(37,36,'来源渠道','OriginChannel','/origin',NULL,'/base/marketing/origin',NULL,NULL,NULL,NULL,NULL,NULL),(38,36,'客户标签','ClientTag','/tag',NULL,'/base/marketing/tag',NULL,NULL,NULL,NULL,NULL,NULL),(39,35,'3','CheckIn','/check-in',NULL,'/base/check-in',NULL,NULL,NULL,NULL,NULL,NULL),(40,39,'房间类型','RoomType','/room',NULL,'/base/check-in/room',NULL,NULL,NULL,NULL,NULL,NULL),(41,39,'楼栋管理','BuildingManage','/building',NULL,'/base/check-in/building',NULL,NULL,NULL,NULL,NULL,NULL),(42,35,'活动','Activity','/activity',NULL,'/base/activity',NULL,NULL,NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `auth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `base_attachment`
--

DROP TABLE IF EXISTS `base_attachment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `base_attachment` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `name` varchar(225) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '上传后文件名',
  `real_name` varchar(225) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文件真实名称',
  `path` varchar(225) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文件绝对路径',
  `url` varchar(225) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'url相对路径',
  `suff` varchar(225) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文件后缀',
  `size` bigint NOT NULL COMMENT '文件大小 B',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '删除状态(Y/N)',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1450 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `base_attachment`
--

LOCK TABLES `base_attachment` WRITE;
/*!40000 ALTER TABLE `base_attachment` DISABLE KEYS */;
INSERT INTO `base_attachment` VALUES (1344,1,'1609459836239560704_logo.787fb61a.png','logo.787fb61a.png','D:\\图片\\testDir\\upload\\img\\20230101\\1609459836239560704_logo.787fb61a.png','http://111.231.218.141:9001/upload/img/20230101/1609459836239560704_logo.787fb61a.png','png',46258,0,1,'2023-01-01 16:01:42',1,'2023-01-01 16:01:42'),(1345,1,'1609460779576209408_00717pRCly4gs71qfmxq2j30k00b43yu.jpg','00717pRCly4gs71qfmxq2j30k00b43yu.jpg','D:\\图片\\testDir\\upload\\img\\20230101\\1609460779576209408_00717pRCly4gs71qfmxq2j30k00b43yu.jpg','http://111.231.218.141:9001/upload/img/20230101/1609460779576209408_00717pRCly4gs71qfmxq2j30k00b43yu.jpg','jpg',22580,0,1,'2023-01-01 16:05:27',1,'2023-01-01 16:05:27'),(1346,1,'1609460823654150144_9C0mso5xyxHg9fa86766b2a3a201423493bf199755a3.png','9C0mso5xyxHg9fa86766b2a3a201423493bf199755a3.png','D:\\图片\\testDir\\upload\\img\\20230101\\1609460823654150144_9C0mso5xyxHg9fa86766b2a3a201423493bf199755a3.png','http://111.231.218.141:9001/upload/img/20230101/1609460823654150144_9C0mso5xyxHg9fa86766b2a3a201423493bf199755a3.png','png',57975,0,1,'2023-01-01 16:05:37',1,'2023-01-01 16:05:37'),(1347,1,'1609460831073873920_00717pRCly4gs71qfmxq2j30k00b43yu.jpg','00717pRCly4gs71qfmxq2j30k00b43yu.jpg','D:\\图片\\testDir\\upload\\img\\20230101\\1609460831073873920_00717pRCly4gs71qfmxq2j30k00b43yu.jpg','http://111.231.218.141:9001/upload/img/20230101/1609460831073873920_00717pRCly4gs71qfmxq2j30k00b43yu.jpg','jpg',22580,0,1,'2023-01-01 16:05:39',1,'2023-01-01 16:05:39'),(1348,1,'1609460857183416320_1671628788072.jpg','1671628788072.jpg','D:\\图片\\testDir\\upload\\img\\20230101\\1609460857183416320_1671628788072.jpg','http://111.231.218.141:9001/upload/img/20230101/1609460857183416320_1671628788072.jpg','jpg',609022,0,1,'2023-01-01 16:05:45',1,'2023-01-01 16:05:45'),(1349,1,'1609484689793429504_j3m8y5.jpg','j3m8y5.jpg','D:\\图片\\testDir\\upload\\img\\20230101\\1609484689793429504_j3m8y5.jpg','http://111.231.218.141:9001/upload/img/20230101/1609484689793429504_j3m8y5.jpg','jpg',42421,0,1,'2023-01-01 17:40:28',1,'2023-01-01 17:40:28'),(1350,1,'1642094693574533120_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642094693574533120_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642094693574533120_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 17:20:58',1,'2023-04-01 17:20:58'),(1351,1,'1642103683448594432_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642103683448594432_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642103683448594432_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 17:56:42',1,'2023-04-01 17:56:42'),(1352,1,'1642103836792348672_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642103836792348672_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642103836792348672_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 17:57:18',1,'2023-04-01 17:57:18'),(1353,1,'1642104058717167616_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642104058717167616_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642104058717167616_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 17:58:11',1,'2023-04-01 17:58:11'),(1354,1,'1642104223133884416_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642104223133884416_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642104223133884416_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 17:58:50',1,'2023-04-01 17:58:50'),(1355,1,'1642104241681096704_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642104241681096704_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642104241681096704_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 17:58:55',1,'2023-04-01 17:58:55'),(1356,1,'1642104378440572928_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642104378440572928_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642104378440572928_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 17:59:27',1,'2023-04-01 17:59:27'),(1357,1,'1642105412160348160_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642105412160348160_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642105412160348160_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 18:03:34',1,'2023-04-01 18:03:34'),(1358,1,'1642122455173586944_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642122455173586944_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642122455173586944_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 19:11:17',1,'2023-04-01 19:11:17'),(1359,1,'1642122486433734656_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642122486433734656_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642122486433734656_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 19:11:25',1,'2023-04-01 19:11:25'),(1360,1,'1642123686482829312_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642123686482829312_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642123686482829312_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 19:16:11',1,'2023-04-01 19:16:11'),(1361,1,'1642124355688226816_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642124355688226816_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642124355688226816_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 19:18:50',1,'2023-04-01 19:18:50'),(1362,1,'1642124610609635328_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642124610609635328_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642124610609635328_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 19:19:51',1,'2023-04-01 19:19:51'),(1363,1,'1642124677416509440_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642124677416509440_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642124677416509440_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 19:20:07',1,'2023-04-01 19:20:07'),(1364,1,'1642124805292449792_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642124805292449792_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642124805292449792_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 19:20:37',1,'2023-04-01 19:20:37'),(1365,1,'1642124851090055168_屏幕截图_20221211_143319.png','屏幕截图_20221211_143319.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642124851090055168_屏幕截图_20221211_143319.png','http://111.231.218.141:9001/upload/img/20230401/1642124851090055168_屏幕截图_20221211_143319.png','png',48958,0,1,'2023-04-01 19:20:48',1,'2023-04-01 19:20:48'),(1366,1,'1642125027045302272_屏幕截图_20221211_143319.png','屏幕截图_20221211_143319.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642125027045302272_屏幕截图_20221211_143319.png','http://111.231.218.141:9001/upload/img/20230401/1642125027045302272_屏幕截图_20221211_143319.png','png',48958,0,1,'2023-04-01 19:21:30',1,'2023-04-01 19:21:30'),(1367,1,'1642136417256628224_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642136417256628224_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642136417256628224_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 20:06:46',1,'2023-04-01 20:06:46'),(1368,1,'1642138561145757696_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642138561145757696_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642138561145757696_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 20:15:17',1,'2023-04-01 20:15:17'),(1369,1,'1642139006618591232_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642139006618591232_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642139006618591232_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 20:17:03',1,'2023-04-01 20:17:03'),(1370,1,'1642139100327731200_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642139100327731200_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642139100327731200_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 20:17:26',1,'2023-04-01 20:17:26'),(1371,1,'1642139831931789312_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642139831931789312_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642139831931789312_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 20:20:20',1,'2023-04-01 20:20:20'),(1372,1,'1642139860281090048_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642139860281090048_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642139860281090048_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 20:20:27',1,'2023-04-01 20:20:27'),(1373,1,'1642139867486904320_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642139867486904320_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642139867486904320_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 20:20:29',1,'2023-04-01 20:20:29'),(1374,1,'1642139944292999168_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642139944292999168_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642139944292999168_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 20:20:47',1,'2023-04-01 20:20:47'),(1375,1,'1642139975494426624_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642139975494426624_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642139975494426624_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 20:20:54',1,'2023-04-01 20:20:54'),(1376,1,'1642139981748133888_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642139981748133888_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642139981748133888_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 20:20:56',1,'2023-04-01 20:20:56'),(1377,1,'1642142554819420160_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642142554819420160_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642142554819420160_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 20:31:09',1,'2023-04-01 20:31:09'),(1378,1,'1642142616249196544_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642142616249196544_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642142616249196544_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 20:31:24',1,'2023-04-01 20:31:24'),(1379,1,'1642142747358945280_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642142747358945280_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642142747358945280_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 20:31:55',1,'2023-04-01 20:31:55'),(1380,1,'1642142821682012160_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642142821682012160_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642142821682012160_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 20:32:13',1,'2023-04-01 20:32:13'),(1381,1,'1642143387682365440_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642143387682365440_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642143387682365440_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 20:34:28',1,'2023-04-01 20:34:28'),(1382,1,'1642143484574982144_屏幕截图_20221211_143319.png','屏幕截图_20221211_143319.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642143484574982144_屏幕截图_20221211_143319.png','http://111.231.218.141:9001/upload/img/20230401/1642143484574982144_屏幕截图_20221211_143319.png','png',48958,0,1,'2023-04-01 20:34:51',1,'2023-04-01 20:34:51'),(1383,1,'1642144914513879040_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642144914513879040_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642144914513879040_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 20:40:32',1,'2023-04-01 20:40:32'),(1384,1,'1642145684579704832_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642145684579704832_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642145684579704832_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 20:43:35',1,'2023-04-01 20:43:35'),(1385,1,'1642145743757139968_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642145743757139968_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642145743757139968_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 20:43:50',1,'2023-04-01 20:43:50'),(1386,1,'1642147290364796928_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642147290364796928_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642147290364796928_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 20:49:58',1,'2023-04-01 20:49:58'),(1387,1,'1642147310916886528_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642147310916886528_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642147310916886528_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 20:50:03',1,'2023-04-01 20:50:03'),(1388,1,'1642148311417774080_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642148311417774080_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642148311417774080_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 20:54:02',1,'2023-04-01 20:54:02'),(1389,1,'1642155063865532416_屏幕截图(2).png','屏幕截图(2).png','D:\\图片\\testDir\\upload\\img\\20230401\\1642155063865532416_屏幕截图(2).png','http://111.231.218.141:9001/upload/img/20230401/1642155063865532416_屏幕截图(2).png','png',582420,0,1,'2023-04-01 21:20:52',1,'2023-04-01 21:20:52'),(1390,1,'1642155114906017792_屏幕截图_20221211_143319.png','屏幕截图_20221211_143319.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642155114906017792_屏幕截图_20221211_143319.png','http://111.231.218.141:9001/upload/img/20230401/1642155114906017792_屏幕截图_20221211_143319.png','png',48958,0,1,'2023-04-01 21:21:04',1,'2023-04-01 21:21:04'),(1391,1,'1642155860556800000_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642155860556800000_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642155860556800000_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 21:24:02',1,'2023-04-01 21:24:02'),(1392,1,'1642155869842989056_屏幕截图_20221211_143319.png','屏幕截图_20221211_143319.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642155869842989056_屏幕截图_20221211_143319.png','http://111.231.218.141:9001/upload/img/20230401/1642155869842989056_屏幕截图_20221211_143319.png','png',48958,0,1,'2023-04-01 21:24:04',1,'2023-04-01 21:24:04'),(1393,1,'1642155920573095936_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642155920573095936_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642155920573095936_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 21:24:16',1,'2023-04-01 21:24:16'),(1394,1,'1642156602684366848_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642156602684366848_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642156602684366848_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 21:26:59',1,'2023-04-01 21:26:59'),(1395,1,'1642156746188283904_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642156746188283904_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642156746188283904_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 21:27:33',1,'2023-04-01 21:27:33'),(1396,1,'1642157272367915008_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642157272367915008_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642157272367915008_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 21:29:38',1,'2023-04-01 21:29:38'),(1397,1,'1642157281041735680_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642157281041735680_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642157281041735680_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 21:29:40',1,'2023-04-01 21:29:40'),(1398,1,'1642157288834752512_屏幕截图_20221211_143319.png','屏幕截图_20221211_143319.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642157288834752512_屏幕截图_20221211_143319.png','http://111.231.218.141:9001/upload/img/20230401/1642157288834752512_屏幕截图_20221211_143319.png','png',48958,0,1,'2023-04-01 21:29:42',1,'2023-04-01 21:29:42'),(1399,1,'1642157432485470208_Capture001.png','Capture001.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642157432485470208_Capture001.png','http://111.231.218.141:9001/upload/img/20230401/1642157432485470208_Capture001.png','png',1803433,0,1,'2023-04-01 21:30:16',1,'2023-04-01 21:30:16'),(1400,1,'1642160792039088128_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642160792039088128_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642160792039088128_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 21:43:37',1,'2023-04-01 21:43:37'),(1401,1,'1642160958531985408_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642160958531985408_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642160958531985408_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 21:44:17',1,'2023-04-01 21:44:17'),(1402,1,'1642161078090620928_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642161078090620928_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642161078090620928_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 21:44:46',1,'2023-04-01 21:44:46'),(1403,1,'1642161363273932800_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642161363273932800_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642161363273932800_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 21:45:54',1,'2023-04-01 21:45:54'),(1404,1,'1642161516030484480_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642161516030484480_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642161516030484480_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 21:46:30',1,'2023-04-01 21:46:30'),(1405,1,'1642161750961840128_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642161750961840128_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642161750961840128_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 21:47:26',1,'2023-04-01 21:47:26'),(1406,1,'1642163072796745728_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642163072796745728_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642163072796745728_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 21:52:41',1,'2023-04-01 21:52:41'),(1407,1,'1642164466605907968_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642164466605907968_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642164466605907968_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 21:58:13',1,'2023-04-01 21:58:13'),(1408,1,'1642165463172538368_屏幕截图_20221225_202505.png','屏幕截图_20221225_202505.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642165463172538368_屏幕截图_20221225_202505.png','http://111.231.218.141:9001/upload/img/20230401/1642165463172538368_屏幕截图_20221225_202505.png','png',5242,0,1,'2023-04-01 22:02:11',1,'2023-04-01 22:02:11'),(1409,1,'1642166386628255744_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642166386628255744_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642166386628255744_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 22:05:51',1,'2023-04-01 22:05:51'),(1410,1,'1642166991811796992_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642166991811796992_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642166991811796992_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 22:08:15',1,'2023-04-01 22:08:15'),(1411,1,'1642167029912854528_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642167029912854528_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642167029912854528_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 22:08:25',1,'2023-04-01 22:08:25'),(1412,1,'1642167164147359744_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642167164147359744_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642167164147359744_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 22:08:57',1,'2023-04-01 22:08:57'),(1413,1,'1642167403663089664_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642167403663089664_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642167403663089664_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 22:09:54',1,'2023-04-01 22:09:54'),(1414,1,'1642167820291694592_屏幕截图(2).png','屏幕截图(2).png','D:\\图片\\testDir\\upload\\img\\20230401\\1642167820291694592_屏幕截图(2).png','http://111.231.218.141:9001/upload/img/20230401/1642167820291694592_屏幕截图(2).png','png',582420,0,1,'2023-04-01 22:11:33',1,'2023-04-01 22:11:33'),(1415,1,'1642168144523976704_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642168144523976704_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642168144523976704_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 22:12:50',1,'2023-04-01 22:12:50'),(1416,1,'1642168250019110912_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642168250019110912_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642168250019110912_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 22:13:15',1,'2023-04-01 22:13:15'),(1417,1,'1642168642593382400_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642168642593382400_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642168642593382400_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 22:14:49',1,'2023-04-01 22:14:49'),(1418,1,'1642168866778931200_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642168866778931200_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642168866778931200_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 22:15:43',1,'2023-04-01 22:15:43'),(1419,1,'1642168924551274496_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642168924551274496_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642168924551274496_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 22:15:56',1,'2023-04-01 22:15:56'),(1420,1,'1642169039944966144_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642169039944966144_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642169039944966144_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 22:16:24',1,'2023-04-01 22:16:24'),(1421,1,'1642170915599966208_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642170915599966208_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642170915599966208_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 22:23:51',1,'2023-04-01 22:23:51'),(1422,1,'1642170981786083328_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642170981786083328_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642170981786083328_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 22:24:07',1,'2023-04-01 22:24:07'),(1423,1,'1642170993412694016_屏幕截图(2).png','屏幕截图(2).png','D:\\图片\\testDir\\upload\\img\\20230401\\1642170993412694016_屏幕截图(2).png','http://111.231.218.141:9001/upload/img/20230401/1642170993412694016_屏幕截图(2).png','png',582420,0,1,'2023-04-01 22:24:10',1,'2023-04-01 22:24:10'),(1424,1,'1642171012081541120_屏幕截图(1).png','屏幕截图(1).png','D:\\图片\\testDir\\upload\\img\\20230401\\1642171012081541120_屏幕截图(1).png','http://111.231.218.141:9001/upload/img/20230401/1642171012081541120_屏幕截图(1).png','png',685545,0,1,'2023-04-01 22:24:14',1,'2023-04-01 22:24:14'),(1425,1,'1642171648072245248_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642171648072245248_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642171648072245248_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 22:26:46',1,'2023-04-01 22:26:46'),(1426,1,'1642171720319131648_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642171720319131648_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642171720319131648_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 22:27:03',1,'2023-04-01 22:27:03'),(1427,1,'1642171886581342208_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642171886581342208_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642171886581342208_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 22:27:42',1,'2023-04-01 22:27:42'),(1428,1,'1642171893560664064_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642171893560664064_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642171893560664064_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 22:27:44',1,'2023-04-01 22:27:44'),(1429,1,'1642172012326576128_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642172012326576128_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230401/1642172012326576128_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-01 22:28:12',1,'2023-04-01 22:28:12'),(1430,1,'1642172081289322496_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642172081289322496_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642172081289322496_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 22:28:29',1,'2023-04-01 22:28:29'),(1431,1,'1642175009257447424_屏幕截图(2).png','屏幕截图(2).png','D:\\图片\\testDir\\upload\\img\\20230401\\1642175009257447424_屏幕截图(2).png','http://111.231.218.141:9001/upload/img/20230401/1642175009257447424_屏幕截图(2).png','png',582420,0,1,'2023-04-01 22:40:07',1,'2023-04-01 22:40:07'),(1432,1,'1642175028949704704_屏幕截图(2).png','屏幕截图(2).png','D:\\图片\\testDir\\upload\\img\\20230401\\1642175028949704704_屏幕截图(2).png','http://111.231.218.141:9001/upload/img/20230401/1642175028949704704_屏幕截图(2).png','png',582420,0,1,'2023-04-01 22:40:12',1,'2023-04-01 22:40:12'),(1433,1,'1642190705605500928_屏幕截图_20221211_143319.png','屏幕截图_20221211_143319.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642190705605500928_屏幕截图_20221211_143319.png','http://111.231.218.141:9001/upload/img/20230401/1642190705605500928_屏幕截图_20221211_143319.png','png',48958,0,1,'2023-04-01 23:42:29',1,'2023-04-01 23:42:29'),(1434,1,'1642192054401458176_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230401\\1642192054401458176_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230401/1642192054401458176_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-01 23:47:51',1,'2023-04-01 23:47:51'),(1435,1,'1642203499394588672_屏幕截图_20221210_093139.png','屏幕截图_20221210_093139.png','D:\\图片\\testDir\\upload\\img\\20230402\\1642203499394588672_屏幕截图_20221210_093139.png','http://111.231.218.141:9001/upload/img/20230402/1642203499394588672_屏幕截图_20221210_093139.png','png',26590,0,1,'2023-04-02 00:33:20',1,'2023-04-02 00:33:20'),(1436,1,'1642351453174554624_屏幕截图_20221210_093317.png','屏幕截图_20221210_093317.png','D:\\图片\\testDir\\upload\\img\\20230402\\1642351453174554624_屏幕截图_20221210_093317.png','http://111.231.218.141:9001/upload/img/20230402/1642351453174554624_屏幕截图_20221210_093317.png','png',71598,0,1,'2023-04-02 10:21:15',1,'2023-04-02 10:21:15'),(1437,1,'1642351466961231872_屏幕截图_20221211_143319.png','屏幕截图_20221211_143319.png','D:\\图片\\testDir\\upload\\img\\20230402\\1642351466961231872_屏幕截图_20221211_143319.png','http://111.231.218.141:9001/upload/img/20230402/1642351466961231872_屏幕截图_20221211_143319.png','png',48958,0,1,'2023-04-02 10:21:18',1,'2023-04-02 10:21:18'),(1438,1,'1642351477455380480_屏幕截图(2).png','屏幕截图(2).png','D:\\图片\\testDir\\upload\\img\\20230402\\1642351477455380480_屏幕截图(2).png','http://111.231.218.141:9001/upload/img/20230402/1642351477455380480_屏幕截图(2).png','png',582420,0,1,'2023-04-02 10:21:20',1,'2023-04-02 10:21:20'),(1439,1,'1642352088343175168_屏幕截图_20221211_143319.png','屏幕截图_20221211_143319.png','D:\\图片\\testDir\\upload\\img\\20230402\\1642352088343175168_屏幕截图_20221211_143319.png','http://111.231.218.141:9001/upload/img/20230402/1642352088343175168_屏幕截图_20221211_143319.png','png',48958,0,1,'2023-04-02 10:23:46',1,'2023-04-02 10:23:46'),(1440,1,'1642352470234554368_屏幕截图(2).png','屏幕截图(2).png','D:\\图片\\testDir\\upload\\img\\20230402\\1642352470234554368_屏幕截图(2).png','http://111.231.218.141:9001/upload/img/20230402/1642352470234554368_屏幕截图(2).png','png',582420,0,1,'2023-04-02 10:25:17',1,'2023-04-02 10:25:17'),(1441,1,'1643600038046306304_logo.png','logo.png','D:\\图片\\testDir\\upload\\img\\20230405\\1643600038046306304_logo.png','http://111.231.218.141:9001/upload/img/20230405/1643600038046306304_logo.png','png',228686,0,1,'2023-04-05 21:02:40',1,'2023-04-05 21:02:40'),(1442,1,'1651104776698556416_56.jpg','56.jpg','D:\\图片\\testDir\\upload\\img\\20230426\\1651104776698556416_56.jpg','http://111.231.218.141:9001/upload/img/20230426/1651104776698556416_56.jpg','jpg',653154,0,1,'2023-04-26 14:03:49',1,'2023-04-26 14:03:49'),(1443,1,'1651104796717969408_56.jpg','56.jpg','D:\\图片\\testDir\\upload\\img\\20230426\\1651104796717969408_56.jpg','http://111.231.218.141:9001/upload/img/20230426/1651104796717969408_56.jpg','jpg',653154,0,1,'2023-04-26 14:03:54',1,'2023-04-26 14:03:54'),(1444,1,'1651132340481650688_56.jpg','56.jpg','D:\\图片\\testDir\\upload\\img\\20230426\\1651132340481650688_56.jpg','http://111.231.218.141:9001/upload/img/20230426/1651132340481650688_56.jpg','jpg',653154,0,1,'2023-04-26 15:53:21',1,'2023-04-26 15:53:21'),(1445,1,'1781361739637297152_1.png','1.png','D:\\java1234-image\\upload\\img\\20240420\\1781361739637297152_1.png','http://127.0.0.1:9001/upload/img/20240420/1781361739637297152_1.png','png',118170,0,1,'2024-04-20 00:38:29',1,'2024-04-20 00:38:29'),(1446,1,'1781361825431785472_2.png','2.png','D:\\java1234-image\\upload\\img\\20240420\\1781361825431785472_2.png','http://127.0.0.1:9001/upload/img/20240420/1781361825431785472_2.png','png',91431,0,1,'2024-04-20 00:38:50',1,'2024-04-20 00:38:50'),(1447,1,'1781362049885769728_3.png','3.png','D:\\java1234-image\\upload\\img\\20240420\\1781362049885769728_3.png','http://127.0.0.1:9001/upload/img/20240420/1781362049885769728_3.png','png',116733,0,1,'2024-04-20 00:39:43',1,'2024-04-20 00:39:43'),(1448,1,'1781362175953965056_4.png','4.png','D:\\java1234-image\\upload\\img\\20240420\\1781362175953965056_4.png','http://127.0.0.1:9001/upload/img/20240420/1781362175953965056_4.png','png',110895,0,1,'2024-04-20 00:40:14',1,'2024-04-20 00:40:14'),(1449,1,'1781362261823950848_3.png','3.png','D:\\java1234-image\\upload\\img\\20240420\\1781362261823950848_3.png','http://127.0.0.1:9001/upload/img/20240420/1781362261823950848_3.png','png',116733,0,1,'2024-04-20 00:40:34',1,'2024-04-20 00:40:34');
/*!40000 ALTER TABLE `base_attachment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bed`
--

DROP TABLE IF EXISTS `bed`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bed` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `room_id` bigint unsigned NOT NULL COMMENT '房间id',
  `name` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '床位编号',
  `type_id` bigint unsigned DEFAULT NULL COMMENT '床型编号：1-普通床，2-护理床，3-加床(关联 material_type.id，kind=1)',
  `status` tinyint NOT NULL COMMENT '床位状态(空闲/预定/入住/退住审核/维修)',
  `price` bigint DEFAULT '0' COMMENT '床位费（分）',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=正常',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_room_id` (`room_id`),
  CONSTRAINT `fk_beds_room` FOREIGN KEY (`room_id`) REFERENCES `room` (`id`),
  CONSTRAINT `fk_beds_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='床位表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bed`
--

LOCK TABLES `bed` WRITE;
/*!40000 ALTER TABLE `bed` DISABLE KEYS */;
INSERT INTO `bed` VALUES (1,1,1,'爱心楼-1层-1房-1床',NULL,0,0,0,1,'2023-01-02 14:28:38',1,'2023-04-14 18:18:29'),(2,1,2,'爱心楼-1层-2房-1床',NULL,0,0,0,1,'2023-01-02 14:28:38',1,'2023-04-14 18:19:12'),(3,1,3,'爱心楼-1层-3房-1床',NULL,0,0,0,1,'2023-01-02 14:28:38',1,'2023-04-23 14:34:34'),(4,1,4,'爱心楼-1层-4房-1床',NULL,2,0,0,1,'2023-01-02 14:28:38',1,'2023-02-01 23:29:43'),(5,1,5,'爱心楼-1层-5房-1床',NULL,0,0,0,1,'2023-01-02 14:28:38',1,'2023-04-19 00:00:00'),(6,1,6,'楼栋2-1层-1房-1床',NULL,3,0,0,1,'2023-01-04 21:39:55',1,'2023-04-23 19:28:37'),(7,1,6,'楼栋2-1层-1房-2床',NULL,3,0,0,1,'2023-01-04 21:40:01',1,'2023-04-26 15:53:50'),(8,1,8,'测试床位1',NULL,1,0,0,1,'2023-04-05 02:38:17',1,'2023-04-24 14:50:51'),(9,1,8,'测试床位2',NULL,0,0,1,1,'2023-04-05 02:38:57',1,'2023-04-05 02:39:54'),(10,1,8,'测试床位21',NULL,3,0,0,1,'2023-04-05 02:40:15',1,'2023-04-23 19:27:18');
/*!40000 ALTER TABLE `bed` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bill`
--

DROP TABLE IF EXISTS `bill`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bill` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人ID',
  `bill_no` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账单编号',
  `bill_period` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账单周期',
  `total_amount` bigint NOT NULL DEFAULT '0' COMMENT '总金额（分）',
  `paid_amount` bigint NOT NULL DEFAULT '0' COMMENT '已付金额（分）',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-未支付，1-部分支付，2-已支付，3-逾期',
  `due_date` date DEFAULT NULL COMMENT '缴费截止日',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_bill_no` (`bill_no`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elder_id` (`elder_id`),
  KEY `idx_status` (`status`),
  KEY `idx_bill_period` (`bill_period`),
  CONSTRAINT `fk_bills_elder` FOREIGN KEY (`elder_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_bills_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='账单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bill`
--

LOCK TABLES `bill` WRITE;
/*!40000 ALTER TABLE `bill` DISABLE KEYS */;
/*!40000 ALTER TABLE `bill` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bill_item`
--

DROP TABLE IF EXISTS `bill_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bill_item` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `bill_id` bigint unsigned NOT NULL COMMENT '账单ID',
  `fee_item_id` bigint unsigned DEFAULT NULL COMMENT '费用项ID',
  `description` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '描述',
  `quantity` decimal(10,2) NOT NULL COMMENT '数量',
  `unit_price` bigint NOT NULL DEFAULT '0' COMMENT '单价（分）',
  `amount` bigint NOT NULL DEFAULT '0' COMMENT '小计（分）',
  `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_bill_id` (`bill_id`),
  KEY `idx_fee_item_id` (`fee_item_id`),
  CONSTRAINT `fk_bill_items_bill` FOREIGN KEY (`bill_id`) REFERENCES `bill` (`id`),
  CONSTRAINT `fk_bill_items_fee_item` FOREIGN KEY (`fee_item_id`) REFERENCES `fee_item` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='账单明细表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bill_item`
--

LOCK TABLES `bill_item` WRITE;
/*!40000 ALTER TABLE `bill_item` DISABLE KEYS */;
/*!40000 ALTER TABLE `bill_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `building`
--

DROP TABLE IF EXISTS `building`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `building` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '楼栋名称',
  `floor_num` int NOT NULL COMMENT '楼层数量',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `building`
--

LOCK TABLES `building` WRITE;
/*!40000 ALTER TABLE `building` DISABLE KEYS */;
INSERT INTO `building` VALUES (1,1,'楼栋1',2,0,1,'2023-01-02 14:19:07',1,'2023-01-04 21:01:15'),(2,1,'楼栋2',3,0,1,'2023-01-04 20:59:59',1,'2023-04-05 02:16:01'),(3,1,'楼栋3',5,1,1,'2023-01-04 21:00:02',1,'2023-04-05 02:33:49'),(4,1,'楼栋4',5,1,1,'2023-01-04 21:00:05',1,'2023-04-05 02:33:45'),(5,1,'楼栋5',5,1,1,'2023-01-04 21:00:08',1,'2023-04-05 02:33:38'),(6,1,'楼栋6',5,1,1,'2023-01-04 21:00:10',1,'2023-04-05 02:33:35'),(7,1,'楼栋7',5,1,1,'2023-01-04 21:00:13',1,'2023-04-05 02:33:05'),(8,1,'楼栋8',5,1,1,'2023-01-04 21:00:15',1,'2023-04-05 02:33:02'),(9,1,'楼栋9',5,1,1,'2023-01-04 21:00:18',1,'2023-04-05 02:32:59'),(10,1,'楼栋10',5,1,1,'2023-01-04 21:00:23',1,'2023-04-05 00:56:29'),(11,1,'楼栋11',10,1,1,'2023-04-05 01:03:10',1,'2023-04-05 01:07:42'),(12,1,'楼栋10',5,1,1,'2023-04-05 01:08:00',1,'2023-04-05 02:32:56'),(13,1,'测试楼栋',2,0,1,'2023-04-05 02:33:25',1,'2023-04-05 02:35:14'),(14,1,'胡图图',10,0,1,'2023-04-14 18:24:21',1,'2023-04-14 18:24:21');
/*!40000 ALTER TABLE `building` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `catering_set`
--

DROP TABLE IF EXISTS `catering_set`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `catering_set` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '餐饮套餐名称',
  `month_price` bigint NOT NULL COMMENT '月套餐费用',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `catering_set`
--

LOCK TABLES `catering_set` WRITE;
/*!40000 ALTER TABLE `catering_set` DISABLE KEYS */;
INSERT INTO `catering_set` VALUES (1,1,'颐养套餐',1200,0,1,'2023-01-13 16:22:34',1,'2023-01-13 16:24:34'),(2,1,'黄金套餐',1200,0,1,'2023-01-13 16:23:12',1,'2023-01-13 16:23:12'),(3,1,'测试',13,0,1,'2023-04-04 11:55:31',1,'2023-04-04 11:57:24'),(4,1,'测试1',120,0,1,'2023-04-04 14:56:11',1,'2023-04-04 14:57:28'),(5,1,'胡图图',100,0,1,'2023-04-14 18:12:09',1,'2023-04-14 18:12:09');
/*!40000 ALTER TABLE `catering_set` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `communication_record`
--

DROP TABLE IF EXISTS `communication_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `communication_record` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `record_date` datetime NOT NULL COMMENT '记录时间',
  `communication_record` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '沟通记录',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_com_tenant` (`tenant_id`,`elder_id`,`record_date`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `communication_record`
--

LOCK TABLES `communication_record` WRITE;
/*!40000 ALTER TABLE `communication_record` DISABLE KEYS */;
INSERT INTO `communication_record` VALUES (1,1,1,'2022-12-14 00:00:00','测试记录',0,1,'2023-01-05 23:12:59',1,'2023-01-05 23:13:57');
/*!40000 ALTER TABLE `communication_record` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `consult`
--

DROP TABLE IF EXISTS `consult`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `consult` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `source_id` bigint unsigned NOT NULL COMMENT '来源渠道id',
  `staff_id` bigint unsigned NOT NULL COMMENT '接待人id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '咨询人姓名',
  `phone` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '咨询人电话',
  `relation` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '与老人关系',
  `consult_date` datetime NOT NULL COMMENT '咨询日期',
  `consult_content` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '咨询内容',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=74 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `consult`
--

LOCK TABLES `consult` WRITE;
/*!40000 ALTER TABLE `consult` DISABLE KEYS */;
INSERT INTO `consult` VALUES (1,1,1,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',2,'2023-01-03 18:51:16',2,'2023-01-03 20:22:07'),(2,1,2,1,1,'张三','13546564658','父子','2022-10-11 00:00:00','养老院环境',2,'2023-01-03 18:53:42',2,'2023-01-03 18:53:42'),(4,1,3,1,1,'张三','13546564658','父子','2022-10-11 00:00:00','养老院环境',2,'2023-01-03 20:04:05',2,'2023-01-03 20:04:05'),(5,1,4,1,1,'张三','13546564658','父子','2022-10-11 00:00:00','养老院环境',2,'2023-01-03 20:04:37',2,'2023-01-03 20:04:37'),(6,1,8,1,1,'张三','13546564658','父子','2022-10-11 00:00:00','养老院环境',2,'2023-01-03 20:05:18',2,'2023-01-03 20:05:18'),(7,1,7,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',1,'2023-01-05 10:23:15',1,'2023-01-05 10:28:34'),(8,1,18,1,1,'管志鑫','18428167777','兄弟','2023-03-09 00:00:00','你好呀',4,'2023-03-09 17:02:54',4,'2023-03-09 17:02:54'),(9,1,19,1,1,'管志鑫','18428167777','兄弟','2023-03-09 00:00:00','你好呀',4,'2023-03-09 17:03:13',4,'2023-03-20 14:56:15'),(10,1,20,1,1,'管志鑫','18428167777','兄弟','2023-03-09 00:00:00','你好呀',4,'2023-03-09 17:03:19',4,'2023-03-09 17:03:19'),(11,1,21,1,1,'管志鑫','18428167777','兄弟','2023-03-09 00:00:00','你好呀',4,'2023-03-09 17:03:30',4,'2023-03-09 17:03:30'),(12,1,22,1,1,'管志鑫','18428167777','兄弟','2023-03-09 00:00:00','你好呀',4,'2023-03-09 17:03:34',4,'2023-03-09 17:03:34'),(13,1,23,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:05:25',4,'2023-03-09 17:05:25'),(14,1,24,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:05:32',4,'2023-03-09 17:05:32'),(15,1,25,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:05:34',4,'2023-03-09 17:05:34'),(16,1,26,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:05:36',4,'2023-03-09 17:05:36'),(17,1,27,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:05:37',4,'2023-03-09 17:05:37'),(18,1,28,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:05:38',4,'2023-03-09 17:05:38'),(19,1,29,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:05:40',4,'2023-03-09 17:05:40'),(20,1,30,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:05:43',4,'2023-03-09 17:05:43'),(21,1,31,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:05:45',4,'2023-03-09 17:05:45'),(22,1,32,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:05:47',4,'2023-03-09 17:05:47'),(23,1,33,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:05:49',4,'2023-03-09 17:05:49'),(24,1,34,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:05:54',4,'2023-03-09 17:05:54'),(25,1,35,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:05:55',4,'2023-03-09 17:05:55'),(26,1,36,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:05:56',4,'2023-03-09 17:05:56'),(27,1,37,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:05:58',4,'2023-03-09 17:05:58'),(28,1,38,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:05:59',4,'2023-03-09 17:05:59'),(29,1,39,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:06:01',4,'2023-03-09 17:06:01'),(30,1,40,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:06:02',4,'2023-03-09 17:06:02'),(31,1,41,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:06:04',4,'2023-03-09 17:06:04'),(32,1,42,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:06:11',4,'2023-03-09 17:06:11'),(33,1,43,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:06:13',4,'2023-03-09 17:06:13'),(34,1,44,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:06:14',4,'2023-03-09 17:06:14'),(35,1,45,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:06:16',4,'2023-03-09 17:06:16'),(36,1,46,1,1,'张三','13546564657','父子','2022-10-10 00:00:00','无',4,'2023-03-09 17:07:35',4,'2023-03-09 17:07:35'),(37,1,47,1,1,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-10 15:09:41',4,'2023-03-10 15:09:41'),(38,1,48,1,1,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-10 15:15:52',4,'2023-03-10 15:15:52'),(39,1,49,1,1,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-10 15:15:55',4,'2023-03-10 15:15:55'),(40,1,50,1,1,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-10 15:16:00',4,'2023-03-10 15:16:00'),(41,1,51,1,1,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-10 15:16:07',4,'2023-03-10 15:16:07'),(42,1,52,1,1,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-10 15:16:09',4,'2023-03-10 15:16:09'),(43,1,53,1,1,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-10 15:16:10',4,'2023-03-10 15:16:10'),(44,1,54,1,1,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-10 15:16:12',4,'2023-03-10 15:16:12'),(45,1,55,1,1,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-10 15:16:14',4,'2023-03-10 15:16:14'),(46,1,56,1,1,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-10 15:16:17',4,'2023-03-10 15:16:17'),(47,1,57,1,1,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-10 15:17:05',4,'2023-03-10 15:17:05'),(48,1,58,1,1,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-10 15:19:04',4,'2023-03-10 15:19:04'),(49,1,59,1,1,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-10 15:19:09',4,'2023-03-10 15:19:09'),(50,1,60,1,1,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-12 14:10:56',4,'2023-03-12 14:10:56'),(51,1,61,1,1,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-12 14:12:33',4,'2023-03-12 14:12:33'),(52,1,62,1,1,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-12 14:12:56',4,'2023-03-12 14:12:56'),(53,1,63,4,2,'李四','18469652132','父子','2002-02-12 00:00:00','你好',4,'2023-03-12 14:43:22',4,'2023-03-12 14:43:22'),(54,1,64,4,2,'李四','18469652132','父子','2002-02-12 00:00:00','你好',4,'2023-03-12 14:57:20',4,'2023-03-12 14:57:20'),(55,1,65,3,2,'李四','18469652132','父子','2002-02-12 00:00:00','你好',4,'2023-03-12 14:57:31',4,'2023-03-12 14:57:31'),(56,1,66,3,2,'李四','18469652132','父子','2002-02-12 00:00:00','你好',4,'2023-03-12 14:57:51',4,'2023-03-12 14:57:51'),(57,1,67,3,2,'刘心','18469652132','父子','2002-02-28 00:00:00','你好',4,'2023-03-12 15:00:24',4,'2023-03-12 15:00:24'),(58,1,68,3,2,'刘心','18469652132','父子','2002-02-28 00:00:00','你好',4,'2023-03-12 15:01:14',4,'2023-03-12 15:01:14'),(59,1,69,3,2,'刘楼','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-12 15:03:00',4,'2023-03-12 15:03:00'),(60,1,70,3,2,'刘楼','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-12 15:03:16',4,'2023-03-12 15:03:16'),(61,1,71,1,2,'刘楼','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-12 15:03:29',4,'2023-03-12 15:03:29'),(62,1,72,1,2,'刘楼','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-12 15:03:50',4,'2023-03-12 15:03:50'),(63,1,73,3,2,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-12 15:05:20',4,'2023-03-12 15:05:20'),(64,1,74,3,2,'李四','18469652132','父子','2002-02-03 00:00:00','你好',4,'2023-03-12 15:07:52',4,'2023-03-12 15:07:52'),(65,1,75,4,2,'qqqqq','18888888888','兄弟','2023-03-25 00:00:00','wwwwwwwwwwwwww',4,'2023-03-12 15:55:29',4,'2023-03-20 16:22:42'),(66,1,76,1,2,'fefef','13443334456','dsd','2023-02-23 00:00:00','fefefefe',4,'2023-03-12 15:59:16',4,'2023-03-12 15:59:16'),(67,1,77,3,2,'龙哥','18888998989','大武当','2023-03-24 00:00:00','dvvvv',4,'2023-03-20 14:45:31',4,'2023-03-20 14:45:31'),(68,1,78,4,2,'测试','13547584560','父子','2023-03-30 00:00:00','测试',1,'2023-03-30 17:40:25',1,'2023-03-30 17:40:52'),(69,1,82,3,2,'蔡徐坤','13267879999','123','2023-04-07 16:24:15','123',1,'2023-04-05 19:28:45',1,'2023-04-22 23:08:55'),(70,1,83,3,2,'全二涛','17666666666','本人','2027-03-10 00:00:00','好',1,'2023-04-05 19:29:00',1,'2023-04-05 19:30:13'),(71,1,84,4,2,'2323','13541990238','123','2023-04-08 00:00:00','13213',1,'2023-04-10 14:19:47',1,'2023-04-14 10:35:54'),(72,1,85,4,1,'asuka','18228011369','无关系','2023-03-30 00:00:00','EVA',1,'2023-04-14 15:43:00',1,'2023-04-14 15:43:00'),(73,1,86,4,1,'张晓莉','18228011369','无关系','2023-04-12 00:00:00','大飒飒',1,'2023-04-14 15:45:31',1,'2023-04-14 15:51:56');
/*!40000 ALTER TABLE `consult` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `consume`
--

DROP TABLE IF EXISTS `consume`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `consume` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `consume_type` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '消费类别',
  `consume_amount` bigint NOT NULL COMMENT '消费金额',
  `consume_date` datetime NOT NULL COMMENT '消费日期',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_consume_tenant` (`tenant_id`,`elder_id`,`consume_date`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `consume`
--

LOCK TABLES `consume` WRITE;
/*!40000 ALTER TABLE `consume` DISABLE KEYS */;
INSERT INTO `consume` VALUES (1,1,1,'护理',100,'2023-01-07 18:47:51',1,'2023-01-07 18:47:57',1,'2023-01-07 18:48:01'),(2,1,2,'点餐',100,'2023-01-07 18:47:51',1,'2023-01-07 18:47:57',1,'2023-01-07 18:48:01'),(3,1,2,'护理',240,'2022-12-13 00:00:00',1,'2023-01-12 23:44:33',1,'2023-01-12 23:44:33'),(4,1,2,'点餐',6,'2022-12-14 00:00:00',1,'2023-01-14 21:09:05',1,'2023-01-14 21:09:05'),(5,1,1,'护理',12,'2023-04-04 00:00:00',1,'2023-04-04 10:14:06',1,'2023-04-04 10:14:06'),(6,1,1,'点餐',3,'2023-04-04 00:00:00',1,'2023-04-04 16:35:29',1,'2023-04-04 16:35:29'),(7,1,1,'点餐',2,'2023-04-13 00:00:00',1,'2023-04-05 19:34:28',1,'2023-04-05 19:34:28'),(8,1,1,'护理',12,'2023-04-24 00:00:00',1,'2023-04-24 19:36:54',1,'2023-04-24 19:36:54');
/*!40000 ALTER TABLE `consume` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `contract`
--

DROP TABLE IF EXISTS `contract`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `contract` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `staff_id` bigint unsigned NOT NULL COMMENT '销售人员id',
  `sign_date` datetime NOT NULL COMMENT '合同签订日期',
  `start_date` datetime NOT NULL COMMENT '合同开始日期',
  `end_date` datetime NOT NULL COMMENT '合同结束日期',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `contract`
--

LOCK TABLES `contract` WRITE;
/*!40000 ALTER TABLE `contract` DISABLE KEYS */;
INSERT INTO `contract` VALUES (8,1,4,1,'2022-12-23 00:00:00','2022-12-23 00:00:00','2023-02-02 00:00:00',1,'2023-02-01 23:28:46',1,'2023-02-01 23:28:46'),(13,1,2,2,'2023-04-01 00:00:00','2023-03-31 00:00:00','2023-04-29 00:00:00',1,'2023-03-31 20:32:03',1,'2023-03-31 20:32:03'),(14,1,17,1,'2022-12-23 12:00:00','2023-03-31 12:00:00','2023-04-30 12:00:00',1,'2023-04-01 09:10:38',1,'2023-04-01 09:10:38'),(59,1,1,2,'2023-04-01 12:00:00','2023-02-01 12:00:00','2023-03-28 12:00:00',1,'2023-04-14 16:58:32',1,'2023-04-14 16:58:32'),(60,1,88,1,'2023-04-08 12:00:00','2023-03-30 12:00:00','2027-04-29 12:00:00',1,'2023-04-24 14:00:57',1,'2023-04-24 14:00:57'),(61,1,8,2,'2023-04-12 12:00:00','2023-05-12 12:00:00','2023-05-17 12:00:00',1,'2023-04-24 17:13:59',1,'2023-04-24 17:13:59');
/*!40000 ALTER TABLE `contract` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `deposit_info`
--

DROP TABLE IF EXISTS `deposit_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `deposit_info` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `deposit_id` bigint unsigned NOT NULL COMMENT '药品缴存id',
  `medicine_id` bigint unsigned NOT NULL COMMENT '药品id',
  `deposit_num` int NOT NULL COMMENT '缴存数量',
  `surplus_num` int NOT NULL COMMENT '剩余数量',
  `warn_num` int NOT NULL COMMENT '预警数量',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `deposit_info`
--

LOCK TABLES `deposit_info` WRITE;
/*!40000 ALTER TABLE `deposit_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `deposit_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `dishes`
--

DROP TABLE IF EXISTS `dishes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `dishes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `type_id` bigint unsigned NOT NULL COMMENT '菜品类别id',
  `name` varchar(15) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜品名称',
  `price` bigint NOT NULL COMMENT '菜品价格',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dishes`
--

LOCK TABLES `dishes` WRITE;
/*!40000 ALTER TABLE `dishes` DISABLE KEYS */;
INSERT INTO `dishes` VALUES (1,1,1,'粥',2,0,1,'2023-01-13 13:03:21',1,'2023-01-13 13:06:54'),(2,1,2,'粥',2,0,1,'2023-01-13 13:03:29',1,'2023-01-13 13:03:29'),(3,1,2,'包子',1,0,1,'2023-01-13 13:06:37',1,'2023-04-14 18:10:01'),(4,1,2,'烧麦',2,1,1,'2023-04-04 10:58:26',1,'2023-04-04 10:58:42'),(5,1,1,'牛爷爷',10,0,1,'2023-04-14 18:09:22',1,'2023-04-14 18:09:22');
/*!40000 ALTER TABLE `dishes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `dishes_type`
--

DROP TABLE IF EXISTS `dishes_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `dishes_type` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜品类别名称',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dishes_type`
--

LOCK TABLES `dishes_type` WRITE;
/*!40000 ALTER TABLE `dishes_type` DISABLE KEYS */;
INSERT INTO `dishes_type` VALUES (1,1,'早餐',0,1,'2023-01-13 13:01:25',1,'2023-01-13 13:02:57'),(2,1,'午餐',0,1,'2023-01-13 13:01:35',1,'2023-01-13 13:01:35'),(3,1,'晚餐',0,1,'2023-01-13 13:01:41',1,'2023-01-13 13:01:41'),(4,1,'曹氏0',1,1,'2023-04-04 11:05:33',1,'2023-04-04 11:07:19'),(5,1,'蔡徐坤',1,1,'2023-04-05 19:33:10',1,'2023-04-05 19:33:33'),(6,1,'胡图图',0,1,'2023-04-14 18:08:54',1,'2023-04-14 18:08:54');
/*!40000 ALTER TABLE `dishes_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `drug_deposit`
--

DROP TABLE IF EXISTS `drug_deposit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `drug_deposit` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `mode` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '缴存药品方式',
  `status` tinyint NOT NULL COMMENT '缴存状态',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `drug_deposit`
--

LOCK TABLES `drug_deposit` WRITE;
/*!40000 ALTER TABLE `drug_deposit` DISABLE KEYS */;
/*!40000 ALTER TABLE `drug_deposit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `elder`
--

DROP TABLE IF EXISTS `elder`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `elder` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `nursing_grade_id` bigint unsigned DEFAULT NULL COMMENT '护理等级id',
  `catering_set_id` bigint unsigned DEFAULT NULL COMMENT '餐饮套餐id',
  `bed_id` bigint unsigned DEFAULT NULL COMMENT '床位id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '老人姓名',
  `id_num` varchar(18) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '身份证号',
  `age` int NOT NULL COMMENT '年龄',
  `sex` varchar(2) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '性别(男/女)',
  `phone` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '老人电话',
  `address` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '地址',
  `balance` bigint NOT NULL COMMENT '余额',
  `status` tinyint NOT NULL COMMENT '入住状态',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=90 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `elder`
--

LOCK TABLES `elder` WRITE;
/*!40000 ALTER TABLE `elder` DISABLE KEYS */;
INSERT INTO `elder` VALUES (1,1,NULL,NULL,NULL,'张三','511901198309133448',70,'男','13546457890','四川南充仪陇',1987,5,1,'2023-01-05 10:23:15',1,'2023-04-24 19:36:54'),(2,1,NULL,NULL,NULL,'李四','511901198309133447',70,'男','13546457890','四川绵阳',4,5,2,'2023-01-03 18:53:42',1,'2023-04-14 17:35:41'),(3,1,NULL,NULL,NULL,'王五','511901198309133448',78,'女','13546457890','四川成都',0,6,2,'2023-01-03 20:04:05',2,'2023-01-03 20:04:05'),(4,1,1,1,6,'张三','511901198309133449',67,'男','13546574657','四川南充',3111,4,2,'2023-01-03 20:04:37',1,'2023-04-24 15:14:41'),(6,1,NULL,NULL,NULL,'李四','511901198309133490',80,'女','13546457890','四川泸州泸县',0,1,2,'2023-01-03 20:05:18',1,'2023-04-04 08:52:13'),(7,1,NULL,NULL,NULL,'张三','511901198309133451',67,'男','13546574657','四川南充营山',56,1,1,'2023-01-05 10:23:15',1,'2023-04-05 19:34:11'),(8,1,1,1,7,'王五','511901198309133452',67,'男','13546574657','四川凉山冕宁',200,4,1,'2023-01-05 18:48:04',1,'2023-04-26 15:53:50'),(9,1,NULL,NULL,NULL,'张三','511901198309133453',78,'男','13546574657','四川南充仪陇',0,1,1,'2023-01-06 19:41:38',1,'2023-04-05 20:38:35'),(16,1,NULL,NULL,NULL,'张三','511901198309133454',78,'男','13546574657','四川南充',0,0,1,'2023-01-06 19:51:48',1,'2023-02-04 22:43:07'),(17,1,1,1,4,'张三','511901198309133455',67,'男','13546574657','四川南充营山',0,3,1,'2023-02-01 23:29:43',1,'2023-04-01 09:10:37'),(18,1,NULL,NULL,NULL,'管志淼','51021252200523',120,'男','18428888888','中国某处',0,6,4,'2023-03-09 17:02:54',4,'2023-03-12 15:24:36'),(19,1,NULL,NULL,NULL,'管志淼','51021252200526',120,'男','18428888888','中国某处',0,0,4,'2023-03-09 17:03:13',4,'2023-03-20 14:56:15'),(20,1,NULL,NULL,NULL,'管志淼','51021252200529',120,'男','18428888888','中国某处',0,6,4,'2023-03-09 17:03:19',4,'2023-03-12 15:20:53'),(21,1,NULL,NULL,NULL,'管志淼','51021252200529999',120,'男','18428888888','中国某处',0,6,4,'2023-03-09 17:03:30',4,'2023-03-12 15:24:51'),(22,1,NULL,NULL,NULL,'管志淼','510212522005299998',120,'男','18428888888','中国某处',0,6,4,'2023-03-09 17:03:34',4,'2023-03-12 15:26:17'),(23,1,NULL,NULL,NULL,'李四','null',67,'男','13546457878','四川南充',0,6,4,'2023-03-09 17:05:25',4,'2023-03-12 15:26:32'),(24,1,NULL,NULL,NULL,'李四','1',67,'男','13546457878','四川南充',0,0,4,'2023-03-09 17:05:32',4,'2023-03-09 17:05:32'),(25,1,NULL,NULL,NULL,'李四','12',67,'男','13546457878','四川南充',0,0,4,'2023-03-09 17:05:34',4,'2023-03-09 17:05:34'),(26,1,NULL,NULL,NULL,'李四','123',67,'男','13546457878','四川南充',0,0,4,'2023-03-09 17:05:36',4,'2023-03-09 17:05:36'),(27,1,NULL,NULL,NULL,'李四','1234',67,'男','13546457878','四川南充',0,6,4,'2023-03-09 17:05:37',1,'2023-03-30 00:42:49'),(28,1,NULL,NULL,NULL,'李四','12345',67,'男','13546457878','四川南充',0,6,4,'2023-03-09 17:05:38',1,'2023-03-30 00:39:50'),(29,1,NULL,NULL,NULL,'李四','123456',67,'男','13546457878','四川南充',0,0,4,'2023-03-09 17:05:40',4,'2023-03-09 17:05:40'),(30,1,NULL,NULL,NULL,'李四','1234567',67,'男','13546457878','四川南充',0,6,4,'2023-03-09 17:05:43',4,'2023-03-12 15:30:33'),(31,1,NULL,NULL,NULL,'李四','12345678',67,'男','13546457878','四川南充',0,6,4,'2023-03-09 17:05:45',4,'2023-03-12 15:36:11'),(32,1,NULL,NULL,NULL,'李四','123456789',67,'男','13546457878','四川南充',0,6,4,'2023-03-09 17:05:47',4,'2023-03-12 15:36:14'),(33,1,NULL,NULL,NULL,'李四','1234567890',67,'男','13546457878','四川南充',0,6,4,'2023-03-09 17:05:49',4,'2023-03-12 15:36:17'),(34,1,NULL,NULL,NULL,'李四','12345678907',67,'男','13546457878','四川南充',0,0,4,'2023-03-09 17:05:54',4,'2023-03-09 17:05:54'),(35,1,NULL,NULL,NULL,'李四','123456786907',67,'男','13546457878','四川南充',0,0,4,'2023-03-09 17:05:55',4,'2023-03-09 17:05:55'),(36,1,NULL,NULL,NULL,'李四','1234567686907',67,'男','13546457878','四川南充',0,0,4,'2023-03-09 17:05:56',4,'2023-03-09 17:05:56'),(37,1,NULL,NULL,NULL,'李四','12345676686907',67,'男','13546457878','四川南充',0,0,4,'2023-03-09 17:05:58',4,'2023-03-09 17:05:58'),(38,1,NULL,NULL,NULL,'李四','123456676686907',67,'男','13546457878','四川南充',0,0,4,'2023-03-09 17:05:59',4,'2023-03-09 17:05:59'),(39,1,NULL,NULL,NULL,'李四','1234566766686907',67,'男','13546457878','四川南充',0,0,4,'2023-03-09 17:06:01',4,'2023-03-09 17:06:01'),(40,1,NULL,NULL,NULL,'李四','12345667666686907',67,'男','13546457878','四川南充',0,0,4,'2023-03-09 17:06:02',4,'2023-03-09 17:06:02'),(41,1,NULL,NULL,NULL,'李四','123456676666686907',67,'男','13546457878','四川南充',0,0,4,'2023-03-09 17:06:04',4,'2023-03-09 17:06:04'),(42,1,NULL,NULL,NULL,'李四','123456664686907',67,'男','13546457878','四川南充',0,0,4,'2023-03-09 17:06:11',4,'2023-03-09 17:06:11'),(43,1,NULL,NULL,NULL,'李四','12345664686907',67,'男','13546457878','四川南充',0,0,4,'2023-03-09 17:06:13',4,'2023-03-09 17:06:13'),(44,1,NULL,NULL,NULL,'李四','1234564686907',67,'男','13546457878','四川南充',0,0,4,'2023-03-09 17:06:14',4,'2023-03-09 17:06:14'),(45,1,NULL,NULL,NULL,'李四','123454686907',67,'男','13546457878','四川南充',0,0,4,'2023-03-09 17:06:16',4,'2023-03-09 17:06:16'),(46,1,NULL,NULL,NULL,'李四','1234546868907',67,'男','18428197749','四川南充',0,0,4,'2023-03-09 17:07:35',4,'2023-03-09 17:07:35'),(47,1,NULL,NULL,NULL,'李武','2386539',156,'男','13266669885','成都',0,0,4,'2023-03-10 15:09:41',4,'2023-03-10 15:09:41'),(48,1,NULL,NULL,NULL,'李武','1375720',156,'男','13266669885','成都',0,0,4,'2023-03-10 15:15:52',4,'2023-03-10 15:15:52'),(49,1,NULL,NULL,NULL,'李武','274536',156,'男','13266669885','成都',0,6,4,'2023-03-10 15:15:55',1,'2023-04-10 14:20:35'),(50,1,NULL,NULL,NULL,'李武','800',156,'男','13266669885','成都',0,0,4,'2023-03-10 15:16:00',4,'2023-03-10 15:16:00'),(51,1,NULL,NULL,NULL,'李武','1558848',156,'男','13266669885','成都',0,0,4,'2023-03-10 15:16:07',4,'2023-03-10 15:16:07'),(52,1,NULL,NULL,NULL,'李武','2609280',156,'男','13266669885','成都',0,0,4,'2023-03-10 15:16:09',4,'2023-03-10 15:16:09'),(53,1,NULL,NULL,NULL,'李武','31152',156,'男','13266669885','成都',0,0,4,'2023-03-10 15:16:10',4,'2023-03-10 15:16:10'),(54,1,NULL,NULL,NULL,'李武','3221800',156,'男','13266669885','成都',0,0,4,'2023-03-10 15:16:12',4,'2023-03-10 15:16:12'),(55,1,NULL,NULL,NULL,'李武','2265666',156,'男','13266669885','成都',0,0,4,'2023-03-10 15:16:14',4,'2023-03-10 15:16:14'),(56,1,NULL,NULL,NULL,'李武','167580',156,'男','13266669885','成都',0,6,4,'2023-03-10 15:16:17',4,'2023-03-12 15:37:04'),(57,1,NULL,NULL,NULL,'李武','50400',156,'男','13266669885','成都',0,6,4,'2023-03-10 15:17:05',4,'2023-03-12 15:37:03'),(58,1,NULL,NULL,NULL,'李武','0',156,'男','13266669885','成都',0,0,4,'2023-03-10 15:19:04',4,'2023-03-10 15:19:04'),(59,1,NULL,NULL,NULL,'李武','331500',156,'男','13266669885','成都',0,0,4,'2023-03-10 15:19:09',4,'2023-03-10 15:19:09'),(60,1,NULL,NULL,NULL,'李武','27966',156,'0','13266669885','成都',0,6,4,'2023-03-12 14:10:56',4,'2023-03-12 15:19:06'),(61,1,NULL,NULL,NULL,'李武','1024850',156,'女','13266669885','成都',0,0,4,'2023-03-12 14:12:33',4,'2023-03-12 14:12:33'),(62,1,NULL,NULL,NULL,'李武','33824',156,'女','13266669885','成都',0,6,4,'2023-03-12 14:12:56',4,'2023-03-12 15:37:06'),(63,1,NULL,NULL,NULL,'李武','373230',156,'男','13266669885','成都',0,0,4,'2023-03-12 14:43:22',4,'2023-03-12 14:43:22'),(64,1,NULL,NULL,NULL,'李武','486963',156,'男','13266669885','成都',0,0,4,'2023-03-12 14:57:20',4,'2023-03-12 14:57:20'),(65,1,NULL,NULL,NULL,'李武','3337928',156,'男','13266669885','成都',0,0,4,'2023-03-12 14:57:31',4,'2023-03-12 14:57:31'),(66,1,NULL,NULL,NULL,'李武','1056951',156,'男','13266669885','成都',0,0,4,'2023-03-12 14:57:51',4,'2023-03-12 14:57:51'),(67,1,NULL,NULL,NULL,'李红','396676',123,'女','13266669885','成都',0,0,4,'2023-03-12 15:00:24',4,'2023-03-12 15:00:24'),(68,1,NULL,NULL,NULL,'李红','500472',123,'女','13266669885','成都',0,0,4,'2023-03-12 15:01:14',4,'2023-03-12 15:01:14'),(69,1,NULL,NULL,NULL,'刘ii','1931524',156,'女','13266669885','成都',0,6,4,'2023-03-12 15:03:00',4,'2023-03-12 15:37:01'),(70,1,NULL,NULL,NULL,'刘ii','5030688',156,'女','13266669885','成都',0,6,4,'2023-03-12 15:03:16',4,'2023-03-12 15:36:58'),(71,1,NULL,NULL,NULL,'刘ii','0xx',156,'男','13266669885','成都',0,6,4,'2023-03-12 15:03:29',4,'2023-03-12 15:36:49'),(72,1,NULL,NULL,NULL,'刘ii','3085936',156,'男','13266669885','成都',0,6,4,'2023-03-12 15:03:50',4,'2023-03-12 15:36:47'),(73,1,NULL,NULL,NULL,'李武','146466',156,'女','13266669885','成都',0,6,4,'2023-03-12 15:05:20',4,'2023-03-12 15:20:35'),(74,1,NULL,NULL,NULL,'李武','257424',156,'男','13266669885','成都',0,6,4,'2023-03-12 15:07:52',4,'2023-03-12 15:20:26'),(75,1,NULL,NULL,NULL,'xccccx','31232141421421421',99,'女','18888882222','打完的无多无多无多',0,6,4,'2023-03-12 15:55:29',1,'2023-03-30 17:41:02'),(76,1,NULL,NULL,NULL,'ndwd','21232323',111,'男','12344567889','defef',0,0,4,'2023-03-12 15:59:16',4,'2023-03-12 15:59:16'),(77,1,NULL,NULL,NULL,'龙爷','4242424243',123,'男','19203300444','的味道无多',0,0,4,'2023-03-20 14:45:31',4,'2023-03-20 14:45:31'),(78,1,NULL,NULL,NULL,'曹氏','512324220503276119',56,'男','13547584390','四川省南充市',0,6,1,'2023-03-30 17:40:25',1,'2023-03-30 17:40:52'),(81,1,NULL,NULL,NULL,'管志淼','511320208912303776',120,'男','18428888888','四川成都高新',0,6,1,'2023-03-31 20:32:03',1,'2023-03-31 20:32:03'),(82,1,NULL,NULL,NULL,'坤哥','510509200708092221',59,'女','13267879999','123456',0,0,1,'2023-04-05 19:28:45',1,'2023-04-22 23:08:55'),(83,1,NULL,NULL,NULL,'全大涛','511820199512230002',131,'男','17666666667','地址地址地址',0,6,1,'2023-04-05 19:29:00',1,'2023-04-05 19:30:26'),(84,1,NULL,NULL,NULL,'1234','511724200202156027',51,'男','13541990238','121312121',0,0,1,'2023-04-10 14:19:47',1,'2023-04-14 10:35:54'),(85,1,NULL,NULL,NULL,'asuka','510182233210186329',63,'男','18228011369','翻斗花园3619',0,0,1,'2023-04-14 15:43:00',1,'2023-04-14 15:43:00'),(86,1,NULL,NULL,NULL,'打到','510182233210186293',63,'男','18228011369','翻斗花园宋大',0,0,1,'2023-04-14 15:45:31',1,'2023-04-14 15:51:56'),(87,1,NULL,NULL,NULL,'胡图图','510182199003041847',96,'男','19126322173','翻斗花园3692',0,1,1,'2023-04-14 16:24:32',1,'2023-04-14 16:24:32'),(88,1,1,1,10,'胡图图','510182200112186491',96,'女','19126322368','翻斗大街翻斗花园',0,4,1,'2023-04-14 16:54:01',1,'2023-04-24 14:00:56'),(89,1,NULL,NULL,8,'王五','513433200412172725',55,'男','13881569055','成都市桂溪街道天益街',0,2,1,'2023-04-24 14:50:51',1,'2023-04-24 14:50:51');
/*!40000 ALTER TABLE `elder` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `elder_label`
--

DROP TABLE IF EXISTS `elder_label`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `elder_label` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `label_id` bigint unsigned NOT NULL COMMENT '标签id',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `elder_label`
--

LOCK TABLES `elder_label` WRITE;
/*!40000 ALTER TABLE `elder_label` DISABLE KEYS */;
INSERT INTO `elder_label` VALUES (1,1,6,1,1,'2023-01-05 13:15:25',1,'2023-01-05 13:15:29'),(2,1,6,2,1,'2023-01-05 13:15:25',1,'2023-01-05 13:15:29'),(4,1,7,1,1,'2023-01-05 21:27:42',1,'2023-01-05 21:27:42'),(5,1,7,2,1,'2023-01-05 21:27:42',1,'2023-01-05 21:27:42');
/*!40000 ALTER TABLE `elder_label` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `emergency_contact`
--

DROP TABLE IF EXISTS `emergency_contact`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `emergency_contact` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '紧急联系人姓名',
  `phone` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '紧急联系人电话',
  `email` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '紧急联系人邮箱',
  `relation` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '与老人关系',
  `status` tinyint NOT NULL COMMENT '是否接收消息（Y/N）',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=103 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `emergency_contact`
--

LOCK TABLES `emergency_contact` WRITE;
/*!40000 ALTER TABLE `emergency_contact` DISABLE KEYS */;
INSERT INTO `emergency_contact` VALUES (5,1,4,'张三','13547564867','3363937530@qq.com','父子',1,1,'2023-02-01 23:28:46',1,'2023-02-01 23:28:46'),(10,1,81,'王氏','13546547364','94659846@qq.com','父子',1,1,'2023-03-31 20:32:03',1,'2023-03-31 20:32:03'),(11,1,17,'王五','13547564867','3363937530@qq.com','父子',0,1,'2023-04-01 09:10:38',1,'2023-04-01 09:10:38'),(99,1,1,'王五','13547564867','3363937530@qq.com','父子',1,1,'2023-04-14 16:58:32',1,'2023-04-14 16:58:32'),(100,1,1,'王氏','13547584400','38274823@qq.com','父子',0,1,'2023-04-14 16:58:32',1,'2023-04-14 16:58:32'),(101,1,88,'牛爷爷','19120322017','m46kf2n@136.cnm','五个关系',1,1,'2023-04-24 14:00:56',1,'2023-04-24 14:00:56'),(102,1,8,'二涛','18740671234','2487283726@qq.com','fa',1,1,'2023-04-24 17:13:59',1,'2023-04-24 17:13:59');
/*!40000 ALTER TABLE `emergency_contact` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `family_account`
--

DROP TABLE IF EXISTS `family_account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `family_account` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `phone` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '家属手机号（登录账号）',
  `pass` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '登录密码（md5）',
  `openid` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '微信 openid（充值支付前置）',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '逻辑删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='家属账号表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `family_account`
--

LOCK TABLES `family_account` WRITE;
/*!40000 ALTER TABLE `family_account` DISABLE KEYS */;
/*!40000 ALTER TABLE `family_account` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `family_member`
--

DROP TABLE IF EXISTS `family_member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `family_member` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '家属姓名',
  `id_num` varchar(18) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '身份证号',
  `phone` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '家属电话',
  `email` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '家属邮箱',
  `address` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '地址',
  `relation` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '与老人关系',
  `status` tinyint NOT NULL COMMENT '是否接收消息（Y/N）',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `family_member`
--

LOCK TABLES `family_member` WRITE;
/*!40000 ALTER TABLE `family_member` DISABLE KEYS */;
/*!40000 ALTER TABLE `family_member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `family_recharge`
--

DROP TABLE IF EXISTS `family_recharge`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `family_recharge` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_no` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商户订单号',
  `phone` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '家属手机号',
  `elder_id` bigint NOT NULL COMMENT '充值到哪位老人',
  `amount` bigint NOT NULL DEFAULT '0' COMMENT '充值金额（单位：分）',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '0-待支付 1-已支付 2-已关闭',
  `prepay_id` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '微信预支付 id',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '逻辑删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_order_no` (`order_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='家属充值订单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `family_recharge`
--

LOCK TABLES `family_recharge` WRITE;
/*!40000 ALTER TABLE `family_recharge` DISABLE KEYS */;
/*!40000 ALTER TABLE `family_recharge` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `fee_item`
--

DROP TABLE IF EXISTS `fee_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `fee_item` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '费用项名称',
  `fee_type` tinyint NOT NULL COMMENT '类型：1-一次性，2-周期性，3-按次',
  `calculation_type` tinyint NOT NULL COMMENT '计算方式：1-固定金额，2-按天，3-按用量',
  `default_price` bigint NOT NULL DEFAULT '0' COMMENT '默认单价（分）',
  `billing_cycle` tinyint DEFAULT NULL COMMENT '周期性费用计费周期：1-月，2-季，3-年',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_tenant_id` (`tenant_id`),
  CONSTRAINT `fk_fee_items_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='费用项配置表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fee_item`
--

LOCK TABLES `fee_item` WRITE;
/*!40000 ALTER TABLE `fee_item` DISABLE KEYS */;
/*!40000 ALTER TABLE `fee_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `field_dict`
--

DROP TABLE IF EXISTS `field_dict`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `field_dict` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `table` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '表名',
  `field` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '字段名(英文列名)',
  `label` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '字段中文名',
  PRIMARY KEY (`id`),
  KEY `idx_field` (`table`,`field`)
) ENGINE=InnoDB AUTO_INCREMENT=794 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='字段中文名字典';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `field_dict`
--

LOCK TABLES `field_dict` WRITE;
/*!40000 ALTER TABLE `field_dict` DISABLE KEYS */;
INSERT INTO `field_dict` VALUES (1,'accident','id','编号'),(2,'accident','staff_id','值班护工'),(3,'accident','occur_date','发生时间'),(4,'accident','description','事故描述'),(5,'accident','picture','事故图片'),(6,'accident','elder_id','老人id'),(7,'accident','state','管理状态'),(8,'active','id','id'),(9,'active','tenant_id','租户id'),(10,'active','type_id','活动类别id'),(11,'active','theme','活动主题'),(12,'active','name','活动名称'),(13,'active','content','活动内容'),(14,'active','address','活动地点'),(15,'active','organizer','组织者姓名'),(16,'active','phone','组织者电话'),(17,'active','active_date','活动日期'),(18,'active','active_picture','活动图片'),(19,'active','state','管理状态'),(20,'active','create_id','创建人id'),(21,'active','create_time','创建时间'),(22,'active','update_id','修改人id'),(23,'active','update_time','修改时间'),(24,'active_participant','id','id'),(25,'active_participant','tenant_id','租户id'),(26,'active_participant','active_id','活动id'),(27,'active_participant','elder_id','老人id'),(28,'active_participant','create_id','创建人id'),(29,'active_participant','create_time','创建时间'),(30,'active_participant','update_id','修改人id'),(31,'active_participant','update_time','修改时间'),(32,'active_type','id','id'),(33,'active_type','tenant_id','租户id'),(34,'active_type','name','活动类型名称'),(35,'active_type','state','管理状态'),(36,'active_type','create_id','创建人id'),(37,'active_type','create_time','创建时间'),(38,'active_type','update_id','修改人id'),(39,'active_type','update_time','管理状态'),(40,'assessment','id','主键ID'),(41,'assessment','tenant_id','租户ID'),(42,'assessment','elder_id','老人ID'),(43,'assessment','assessment_date','评估日期'),(44,'assessment','assessment_type','评估类型'),(45,'assessment','scale_data','量表原始数据'),(46,'assessment','result_level','评估结果'),(47,'assessment','evaluator_id','评估人ID'),(48,'assessment','created_at','创建时间'),(49,'assessment','updated_at','更新时间'),(50,'assessment','deleted_at','软删除时间'),(51,'audit_log','id','主键id'),(52,'audit_log','tenant_id','租户id'),(53,'audit_log','table','被操作表名'),(54,'audit_log','row_id','被操作行主键id'),(55,'audit_log','action','操作'),(56,'audit_log','operator_id','操作员id'),(57,'audit_log','operator_name','操作员名称'),(58,'audit_log','change_after','变更快照'),(59,'audit_log','change_label','变更摘要'),(60,'audit_log','comment','业务备注'),(61,'audit_log','create_time','操作时间'),(62,'auth','id','id'),(63,'auth','parent_id','父级id'),(64,'auth','title','权限标题'),(65,'auth','name','权限名称'),(66,'auth','path','权限path'),(67,'auth','icon','权限图标'),(68,'auth','url','权限url'),(69,'auth','type','权限类别'),(70,'auth','method','权限请求方式'),(71,'auth','create_id','创建人id'),(72,'auth','create_time','创建时间'),(73,'auth','update_id','修改人id'),(74,'auth','update_time','修改时间'),(75,'base_attachment','id','id'),(76,'base_attachment','tenant_id','租户id'),(77,'base_attachment','name','上传后文件名'),(78,'base_attachment','real_name','文件真实名称'),(79,'base_attachment','path','文件绝对路径'),(80,'base_attachment','url','url相对路径'),(81,'base_attachment','suff','文件后缀'),(82,'base_attachment','size','文件大小 B'),(83,'base_attachment','state','管理状态'),(84,'base_attachment','create_id','创建人id'),(85,'base_attachment','create_time','创建时间'),(86,'base_attachment','update_id','修改人id'),(87,'base_attachment','update_time','修改时间'),(88,'bed','id','id'),(89,'bed','tenant_id','租户id'),(90,'bed','room_id','房间id'),(91,'bed','name','床位编号'),(92,'bed','type_id','床型编号'),(93,'bed','status','床位状态'),(94,'bed','price','床位费（分）'),(95,'bed','state','管理状态'),(96,'bed','create_id','创建人id'),(97,'bed','create_time','创建时间'),(98,'bed','update_id','修改人id'),(99,'bed','update_time','修改时间'),(100,'bill','id','主键ID'),(101,'bill','tenant_id','租户ID'),(102,'bill','elder_id','老人ID'),(103,'bill','bill_no','账单编号'),(104,'bill','bill_period','账单周期'),(105,'bill','total_amount','总金额（分）'),(106,'bill','paid_amount','已付金额（分）'),(107,'bill','status','账单支付状态'),(108,'bill','due_date','缴费截止日'),(109,'bill','create_id','创建人id'),(110,'bill','create_time','创建时间'),(111,'bill','update_id','修改人id'),(112,'bill','update_time','修改时间'),(113,'bill_item','id','主键ID'),(114,'bill_item','bill_id','账单ID'),(115,'bill_item','fee_item_id','费用项ID'),(116,'bill_item','description','描述'),(117,'bill_item','quantity','数量'),(118,'bill_item','unit_price','单价（分）'),(119,'bill_item','amount','小计（分）'),(120,'bill_item','created_time','创建时间'),(121,'building','id','id'),(122,'building','tenant_id','租户id'),(123,'building','name','楼栋名称'),(124,'building','floor_num','楼层数量'),(125,'building','state','管理状态'),(126,'building','create_id','创建人id'),(127,'building','create_time','创建时间'),(128,'building','update_id','修改人id'),(129,'building','update_time','修改时间'),(130,'catering_set','id','id'),(131,'catering_set','tenant_id','租户id'),(132,'catering_set','name','餐饮套餐名称'),(133,'catering_set','month_price','月套餐费用'),(134,'catering_set','state','管理状态'),(135,'catering_set','create_id','创建人id'),(136,'catering_set','create_time','创建时间'),(137,'catering_set','update_id','修改人id'),(138,'catering_set','update_time','修改时间'),(139,'communication_record','id','id'),(140,'communication_record','tenant_id','租户id'),(141,'communication_record','elder_id','老人id'),(142,'communication_record','record_date','记录时间'),(143,'communication_record','communication_record','沟通记录'),(144,'communication_record','state','管理状态'),(145,'communication_record','create_id','创建人id'),(146,'communication_record','create_time','创建时间'),(147,'communication_record','update_id','修改人id'),(148,'communication_record','update_time','修改时间'),(149,'consult','id','id'),(150,'consult','tenant_id','租户id'),(151,'consult','elder_id','老人id'),(152,'consult','source_id','来源渠道id'),(153,'consult','staff_id','接待人id'),(154,'consult','name','咨询人姓名'),(155,'consult','phone','咨询人电话'),(156,'consult','relation','与老人关系'),(157,'consult','consult_date','咨询日期'),(158,'consult','consult_content','咨询内容'),(159,'consult','create_id','创建人id'),(160,'consult','create_time','创建时间'),(161,'consult','update_id','修改人id'),(162,'consult','update_time','修改时间'),(163,'consume','id','id'),(164,'consume','tenant_id','租户id'),(165,'consume','elder_id','老人id'),(166,'consume','consume_type','消费类别'),(167,'consume','consume_amount','消费金额'),(168,'consume','consume_date','消费日期'),(169,'consume','create_id','创建人id'),(170,'consume','create_time','创建时间'),(171,'consume','update_id','修改人id'),(172,'consume','update_time','修改时间'),(173,'contract','id','id'),(174,'contract','tenant_id','租户id'),(175,'contract','elder_id','老人id'),(176,'contract','staff_id','销售人员id'),(177,'contract','sign_date','合同签订日期'),(178,'contract','start_date','合同开始日期'),(179,'contract','end_date','合同结束日期'),(180,'contract','create_id','创建人id'),(181,'contract','create_time','创建时间'),(182,'contract','update_id','修改人id'),(183,'contract','update_time','修改时间'),(184,'deposit_info','id','id'),(185,'deposit_info','tenant_id','租户id'),(186,'deposit_info','deposit_id','药品缴存id'),(187,'deposit_info','medicine_id','药品id'),(188,'deposit_info','deposit_num','缴存数量'),(189,'deposit_info','surplus_num','剩余数量'),(190,'deposit_info','warn_num','预警数量'),(191,'deposit_info','create_id','创建人id'),(192,'deposit_info','create_time','创建时间'),(193,'deposit_info','update_id','修改人id'),(194,'deposit_info','update_time','修改时间'),(195,'dishes','id','id'),(196,'dishes','tenant_id','租户id'),(197,'dishes','type_id','菜品类别id'),(198,'dishes','name','菜品名称'),(199,'dishes','price','菜品价格'),(200,'dishes','state','管理状态'),(201,'dishes','create_id','创建人id'),(202,'dishes','create_time','创建时间'),(203,'dishes','update_id','修改人id'),(204,'dishes','update_time','修改时间'),(205,'dishes_type','id','id'),(206,'dishes_type','tenant_id','租户id'),(207,'dishes_type','name','菜品类别名称'),(208,'dishes_type','state','管理状态'),(209,'dishes_type','create_id','创建人id'),(210,'dishes_type','create_time','创建时间'),(211,'dishes_type','update_id','修改人id'),(212,'dishes_type','update_time','修改时间'),(213,'drug_deposit','id','id'),(214,'drug_deposit','tenant_id','租户id'),(215,'drug_deposit','elder_id','老人id'),(216,'drug_deposit','mode','缴存药品方式'),(217,'drug_deposit','status','缴存状态'),(218,'drug_deposit','state','管理状态'),(219,'drug_deposit','create_id','创建人id'),(220,'drug_deposit','create_time','创建时间'),(221,'drug_deposit','update_id','修改人id'),(222,'drug_deposit','update_time','修改时间'),(223,'elder','id','id'),(224,'elder','tenant_id','租户id'),(225,'elder','nursing_grade_id','护理等级id'),(226,'elder','catering_set_id','餐饮套餐id'),(227,'elder','bed_id','床位id'),(228,'elder','name','老人姓名'),(229,'elder','id_num','身份证号'),(230,'elder','age','年龄'),(231,'elder','sex','性别(男/女)'),(232,'elder','phone','老人电话'),(233,'elder','address','地址'),(234,'elder','balance','余额'),(235,'elder','status','入住状态'),(236,'elder','create_id','创建人id'),(237,'elder','create_time','创建时间'),(238,'elder','update_id','修改人id'),(239,'elder','update_time','修改时间'),(240,'elder_label','id','id'),(241,'elder_label','tenant_id','租户id'),(242,'elder_label','elder_id','老人id'),(243,'elder_label','label_id','标签id'),(244,'elder_label','create_id','创建人id'),(245,'elder_label','create_time','创建时间'),(246,'elder_label','update_id','修改人id'),(247,'elder_label','update_time','修改时间'),(248,'emergency_contact','id','id'),(249,'emergency_contact','tenant_id','租户id'),(250,'emergency_contact','elder_id','老人id'),(251,'emergency_contact','name','紧急联系人姓名'),(252,'emergency_contact','phone','紧急联系人电话'),(253,'emergency_contact','email','紧急联系人邮箱'),(254,'emergency_contact','relation','与老人关系'),(255,'emergency_contact','status','是否接收消息'),(256,'emergency_contact','create_id','创建人id'),(257,'emergency_contact','create_time','创建时间'),(258,'emergency_contact','update_id','修改人id'),(259,'emergency_contact','update_time','修改时间'),(260,'family_account','id','主键'),(261,'family_account','phone','家属登录账号'),(262,'family_account','pass','登录密码'),(263,'family_account','openid','微信 openid'),(264,'family_account','create_time','创建时间'),(265,'family_account','update_time','更新时间'),(266,'family_account','state','逻辑删除'),(267,'family_member','id','id'),(268,'family_member','tenant_id','租户id'),(269,'family_member','elder_id','老人id'),(270,'family_member','name','家属姓名'),(271,'family_member','id_num','身份证号'),(272,'family_member','phone','家属电话'),(273,'family_member','email','家属邮箱'),(274,'family_member','address','地址'),(275,'family_member','relation','与老人关系'),(276,'family_member','status','是否接收消息'),(277,'family_member','state','管理状态'),(278,'family_member','create_id','创建人id'),(279,'family_member','create_time','创建时间'),(280,'family_member','update_id','修改人id'),(281,'family_member','update_time','修改时间'),(282,'family_recharge','id','主键'),(283,'family_recharge','order_no','商户订单号'),(284,'family_recharge','phone','家属手机号'),(285,'family_recharge','elder_id','充值到老人id'),(286,'family_recharge','amount','充值金额（分）'),(287,'family_recharge','status','支付状态'),(288,'family_recharge','prepay_id','微信预支付 id'),(289,'family_recharge','create_time','创建时间'),(290,'family_recharge','update_time','更新时间'),(291,'family_recharge','state','逻辑删除'),(292,'fee_item','id','主键ID'),(293,'fee_item','tenant_id','租户ID'),(294,'fee_item','name','费用项名称'),(295,'fee_item','fee_type','类型'),(296,'fee_item','calculation_type','计算方式'),(297,'fee_item','default_price','默认单价（分）'),(298,'fee_item','billing_cycle','计费周期'),(299,'fee_item','state','管理状态'),(300,'fee_item','create_id','创建人id'),(301,'fee_item','create_time','创建时间'),(302,'fee_item','update_id','修改人id'),(303,'fee_item','update_time','修改时间'),(304,'floor','id','id'),(305,'floor','tenant_id','租户id'),(306,'floor','building_id','楼栋id'),(307,'floor','name','楼层名称'),(308,'floor','room_num','房间数量'),(309,'floor','state','管理状态'),(310,'floor','create_id','创建人id'),(311,'floor','create_time','创建时间'),(312,'floor','update_id','修改人id'),(313,'floor','update_time','修改时间'),(314,'health_data','id','id'),(315,'health_data','tenant_id','租户id'),(316,'health_data','elder_id','老人id'),(317,'health_data','height','身高'),(318,'health_data','weight','体重'),(319,'health_data','temperature','体温'),(320,'health_data','heart_rate','心率'),(321,'health_data','systolic_blood_pressure','收缩血压'),(322,'health_data','diastolic_blood_pressure','舒张血压'),(323,'health_data','fasting_blood_glucose','空腹血糖'),(324,'health_data','postprandial_blood_glucose','餐后血糖'),(325,'health_data','blood_oxygen_saturation','血氧饱和度'),(326,'health_data','cholesterol','总胆固醇'),(327,'health_data','uric_acid','尿酸'),(328,'health_data','left_eye','左眼'),(329,'health_data','right_eye','右眼'),(330,'health_data','left_ear','左耳'),(331,'health_data','right_ear','右耳'),(332,'health_data','muscle_percentage','肌肉率'),(333,'health_data','body_fat_percentage','体脂率'),(334,'health_data','waist_circumference','腰围'),(335,'health_data','hip_circumference','臀围'),(336,'health_data','moisture_content','水分率'),(337,'health_data','create_id','创建人id'),(338,'health_data','create_time','创建时间'),(339,'health_data','update_id','修改人id'),(340,'health_data','update_time','修改时间'),(341,'health_info','id','id'),(342,'health_info','tenant_id','租户id'),(343,'health_info','elder_id','老人id'),(344,'health_info','self_care','自理情况'),(345,'health_info','vision','视力'),(346,'health_info','hearing','听力'),(347,'health_info','hospital','主治医院'),(348,'health_info','doctor','主治医师'),(349,'health_info','phone','医院电话'),(350,'health_info','allergy_drug','过敏药物'),(351,'health_info','medical_history','病史'),(352,'health_info','major_disease','主要疾病'),(353,'health_info','create_id','创建人id'),(354,'health_info','create_time','创建时间'),(355,'health_info','update_id','修改人id'),(356,'health_info','update_time','修改时间'),(357,'label','id','id'),(358,'label','tenant_id','租户id'),(359,'label','type_id','类别id'),(360,'label','name','标签名称'),(361,'label','color','标签颜色'),(362,'label','state','管理状态'),(363,'label','create_id','创建人id'),(364,'label','create_time','创建时间'),(365,'label','update_id','修改人id'),(366,'label','update_time','修改时间'),(367,'label_type','id','id'),(368,'label_type','tenant_id','租户id'),(369,'label_type','name','分类名称'),(370,'label_type','state','管理状态'),(371,'label_type','create_id','创建人id'),(372,'label_type','create_time','创建时间'),(373,'label_type','update_id','修改人id'),(374,'label_type','update_time','修改时间'),(375,'material','id','id'),(376,'material','tenant_id','租户id'),(377,'material','type_id','物资类别id'),(378,'material','name','物资名称'),(379,'material','price','物资单价'),(380,'material','state','管理状态'),(381,'material','create_id','创建人id'),(382,'material','create_time','创建时间'),(383,'material','update_id','修改人id'),(384,'material','update_time','修改时间'),(385,'material_type','id','id'),(386,'material_type','tenant_id','租户id'),(387,'material_type','name','物资类别名称'),(388,'material_type','kind','分类用途'),(389,'material_type','state','管理状态'),(390,'material_type','create_id','创建人id'),(391,'material_type','create_time','创建时间'),(392,'material_type','update_id','修改人id'),(393,'material_type','update_time','修改时间'),(394,'medicine','id','id'),(395,'medicine','tenant_id','租户id'),(396,'medicine','name','药品名称'),(397,'medicine','type','药品类别'),(398,'medicine','specification','药品规格'),(399,'medicine','dosage_form','药品剂型'),(400,'medicine','manufacturer','生产厂家'),(401,'medicine','state','管理状态'),(402,'medicine','create_id','创建人id'),(403,'medicine','create_time','创建时间'),(404,'medicine','update_id','修改人id'),(405,'medicine','update_time','修改时间'),(406,'medicine_record','id','id'),(407,'medicine_record','tenant_id','租户id'),(408,'medicine_record','elder_id','老人id'),(409,'medicine_record','deposit_info_id','药品缴存信息id'),(410,'medicine_record','medicine_time','用药时间（早/中/晚）'),(411,'medicine_record','medicine_date','用药日期'),(412,'medicine_record','create_id','创建人id'),(413,'medicine_record','create_time','创建时间'),(414,'medicine_record','update_id','修改人id'),(415,'medicine_record','update_time','修改时间'),(416,'medicine_set','id','id'),(417,'medicine_set','tenant_id','租户id'),(418,'medicine_set','deposit_info_id','药品缴存信息id'),(419,'medicine_set','medicine_time','用药时间（餐前/餐后）'),(420,'medicine_set','day_frequency','天频率'),(421,'medicine_set','create_id','创建人id'),(422,'medicine_set','create_time','创建时间'),(423,'medicine_set','update_id','修改人id'),(424,'medicine_set','update_time','修改时间'),(425,'member','id','id'),(426,'member','user_id','全局用户id'),(427,'member','tenant_id','租户id'),(428,'member','role_id','角色编号(关联role)'),(429,'member','permissions','预留：细粒度权限'),(430,'member','status','状态：0在职 1离职'),(431,'member','create_id','创建人id'),(432,'member','create_time','创建时间'),(433,'member','update_id','修改人id'),(434,'member','update_time','修改时间'),(435,'member','state','管理状态'),(436,'nurse','id','id'),(437,'nurse','tenant_id','租户id'),(438,'nurse','elder_id','老人id'),(439,'nurse','staff_id','护理人员id'),(440,'nurse','nurse_date','护理时间'),(441,'nurse','complete_status','护理完成情况'),(442,'nurse','dine_status','进餐情况'),(443,'nurse','rest','休息'),(444,'nurse','take_medicine','服药'),(445,'nurse','active','活动'),(446,'nurse','create_id','创建人id'),(447,'nurse','create_time','创建时间'),(448,'nurse','update_id','修改人id'),(449,'nurse','update_time','修改时间'),(450,'nurse_grade','id','id'),(451,'nurse_grade','tenant_id','租户id'),(452,'nurse_grade','name','级别名称'),(453,'nurse_grade','type','护理类型'),(454,'nurse_grade','month_price','月护理费用'),(455,'nurse_grade','state','管理状态'),(456,'nurse_grade','create_id','创建人id'),(457,'nurse_grade','create_time','创建时间'),(458,'nurse_grade','update_id','修改人id'),(459,'nurse_grade','update_time','修改时间'),(460,'nurse_group','id','id'),(461,'nurse_group','tenant_id','租户id'),(462,'nurse_group','staff_id','护工小组组长id'),(463,'nurse_group','name','护工小组名称'),(464,'nurse_group','state','管理状态'),(465,'nurse_group','create_id','创建人id'),(466,'nurse_group','create_time','创建时间'),(467,'nurse_group','update_id','修改人id'),(468,'nurse_group','update_time','修改时间'),(469,'nurse_group_member','id','id'),(470,'nurse_group_member','tenant_id','租户id'),(471,'nurse_group_member','group_id','护工小组id'),(472,'nurse_group_member','staff_id','护工id'),(473,'nurse_group_member','create_id','创建人id'),(474,'nurse_group_member','create_time','创建时间'),(475,'nurse_group_member','update_id','修改人id'),(476,'nurse_group_member','update_time','修改时间'),(477,'nurse_item','id','id'),(478,'nurse_item','tenant_id','租户id'),(479,'nurse_item','grade_id','护理等级id'),(480,'nurse_item','service_id','服务项目id'),(481,'nurse_item','create_id','创建人id'),(482,'nurse_item','create_time','创建时间'),(483,'nurse_item','update_id','修改人id'),(484,'nurse_item','update_time','修改时间'),(485,'nurse_reserve','id','id'),(486,'nurse_reserve','tenant_id','租户id'),(487,'nurse_reserve','elder_id','老人id'),(488,'nurse_reserve','staff_id','服务人id'),(489,'nurse_reserve','service_name','服务项目名称'),(490,'nurse_reserve','need_date','所需时间'),(491,'nurse_reserve','service_price','服务费用'),(492,'nurse_reserve','charge_method','收费方式'),(493,'nurse_reserve','frequency','服务次数'),(494,'nurse_reserve','pay_amount','支付总额'),(495,'nurse_reserve','status','订单状态'),(496,'nurse_reserve','nurse_date','护理时间'),(497,'nurse_reserve','create_id','创建人id'),(498,'nurse_reserve','create_time','创建时间'),(499,'nurse_reserve','update_id','修改人id'),(500,'nurse_reserve','update_time','修改时间'),(501,'order','id','id'),(502,'order','tenant_id','租户id'),(503,'order','elder_id','老人id'),(504,'order','staff_id','送餐人id'),(505,'order','deliver_dishes_date','送餐时间'),(506,'order','dine_date','就餐时间'),(507,'order','dine_type','就餐方式'),(508,'order','pay_amount','支付总额'),(509,'order','status','订单状态'),(510,'order','create_id','创建人id'),(511,'order','create_time','创建时间'),(512,'order','update_id','修改人id'),(513,'order','update_time','修改时间'),(514,'order_dishes','id','id'),(515,'order_dishes','tenant_id','租户id'),(516,'order_dishes','order_id','订餐id'),(517,'order_dishes','dishes_name','菜品名称'),(518,'order_dishes','dishes_price','菜品价格'),(519,'order_dishes','order_num','菜品份数'),(520,'order_dishes','status','套餐标记'),(521,'order_dishes','total_amount','菜品总额'),(522,'order_dishes','really_amount','实际总额'),(523,'order_dishes','create_id','创建人id'),(524,'order_dishes','create_time','创建时间'),(525,'order_dishes','update_id','修改人id'),(526,'order_dishes','update_time','修改时间'),(527,'outbound_material','id','id'),(528,'outbound_material','tenant_id','租户id'),(529,'outbound_material','outbound_record_id','出库登记id'),(530,'outbound_material','warehouse_material_id','入库物资id'),(531,'outbound_material','material_id','物资id'),(532,'outbound_material','outbound_num','出库数量'),(533,'outbound_material','create_id','创建人id'),(534,'outbound_material','create_time','创建时间'),(535,'outbound_material','update_id','修改人id'),(536,'outbound_material','update_time','修改时间'),(537,'outbound_record','id','id'),(538,'outbound_record','tenant_id','租户id'),(539,'outbound_record','warehouse_id','仓库id'),(540,'outbound_record','staff_id','经办人id'),(541,'outbound_record','recipient_id','领用人id'),(542,'outbound_record','recipient_type','领用人类型'),(543,'outbound_record','material_use','物资去向'),(544,'outbound_record','outbound_date','出库时间'),(545,'outbound_record','status','出库状态'),(546,'outbound_record','state','管理状态'),(547,'outbound_record','create_id','创建人id'),(548,'outbound_record','create_time','创建时间'),(549,'outbound_record','update_id','修改人id'),(550,'outbound_record','update_time','修改时间'),(551,'outward','id','id'),(552,'outward','tenant_id','租户id'),(553,'outward','elder_id','老人id'),(554,'outward','chaperone_name','陪同人姓名'),(555,'outward','chaperone_phone','陪同人电话'),(556,'outward','chaperone_type','陪同人类型（家属/护工）'),(557,'outward','outward_date','外出时间'),(558,'outward','plan_return_date','计划返回时间'),(559,'outward','real_return_date','实际返回时间'),(560,'outward','state','管理状态'),(561,'outward','create_id','创建人id'),(562,'outward','create_time','创建时间'),(563,'outward','update_id','修改人id'),(564,'outward','update_time','修改时间'),(565,'reserve','id','id'),(566,'reserve','tenant_id','租户id'),(567,'reserve','elder_id','老人id'),(568,'reserve','staff_id','销售人员id'),(569,'reserve','name','交款人姓名'),(570,'reserve','phone','交款人电话'),(571,'reserve','due_date','预定到期时间'),(572,'reserve','deposit','定金'),(573,'reserve','status','退款状态（N/Y）'),(574,'reserve','create_id','创建人id'),(575,'reserve','create_time','创建时间'),(576,'reserve','update_id','修改人id'),(577,'reserve','update_time','修改时间'),(578,'retreat','id','id'),(579,'retreat','tenant_id','租户id'),(580,'retreat','elder_id','老人id'),(581,'retreat','retreat_form','退住形式'),(582,'retreat','evaluate','对老人评价'),(583,'retreat','retreat_cause','退住原因'),(584,'retreat','create_id','创建人id'),(585,'retreat','create_time','创建时间'),(586,'retreat','update_id','修改人id'),(587,'retreat','update_time','修改时间'),(588,'retreat_apply','id','id'),(589,'retreat_apply','tenant_id','租户id'),(590,'retreat_apply','elder_id','老人id'),(591,'retreat_apply','status','业务状态(退住申请流程状态)'),(592,'retreat_apply','state','管理状态：-1=删除，0=禁用，1=正常'),(593,'retreat_apply','create_id','创建人id'),(594,'retreat_apply','create_time','创建时间'),(595,'retreat_apply','update_id','修改人id'),(596,'retreat_apply','update_time','修改时间'),(597,'role','id','id'),(598,'role','tenant_id','租户ID（平台角色tenant_id=0）'),(599,'role','name','角色名称'),(600,'role','permissions','权限列表'),(601,'role','is_system','是否系统预置：0-否，1-是'),(602,'role','create_id','创建人id'),(603,'role','create_time','创建时间'),(604,'role','update_id','修改人id'),(605,'role','update_time','修改时间'),(606,'role_auth','id','id'),(607,'role_auth','role_id','角色id'),(608,'role_auth','auth_id','权限id'),(609,'role_auth','create_id','创建人id'),(610,'role_auth','create_time','创建时间'),(611,'role_auth','update_id','修改人id'),(612,'role_auth','update_time','修改时间'),(613,'room','id','id'),(614,'room','tenant_id','租户id'),(615,'room','type_id','房间类型id'),(616,'room','build_id','楼栋id'),(617,'room','floor_id','楼层id'),(618,'room','name','房间名称'),(619,'room','bed_num','床位数量'),(620,'room','type','房间类型：1-单人间，2-双人间，3-多人间'),(621,'room','status','状态：0-空闲，1-部分占用，2-全满，3-维修'),(622,'room','price','房间基准价（分）'),(623,'room','state','管理状态'),(624,'room','create_id','创建人id'),(625,'room','create_time','创建时间'),(626,'room','update_id','修改人id'),(627,'room','update_time','修改时间'),(628,'room_material','id','id'),(629,'room_material','tenant_id','租户id'),(630,'room_material','room_id','房间id'),(631,'room_material','material_type_id','设施(物资分类)编号，kind=0'),(632,'room_material','state','删除状态(Y/N)'),(633,'room_material','create_id','创建人id'),(634,'room_material','create_time','创建时间'),(635,'room_material','update_id','修改人id'),(636,'room_material','update_time','修改时间'),(637,'room_transfer','id','主键ID'),(638,'room_transfer','tenant_id','租户ID'),(639,'room_transfer','elder_id','老人ID'),(640,'room_transfer','from_bed_id','原床位ID'),(641,'room_transfer','to_bed_id','新床位ID'),(642,'room_transfer','transfer_date','转房日期'),(643,'room_transfer','reason','原因'),(644,'room_transfer','create_id','创建人id'),(645,'room_transfer','create_time','创建时间'),(646,'room_transfer','update_id','修改人id'),(647,'room_transfer','update_time','修改时间'),(648,'room_type','id','id'),(649,'room_type','tenant_id','租户id'),(650,'room_type','name','房间类型名称'),(651,'room_type','month_price','月房间费用'),(652,'room_type','state','管理状态'),(653,'room_type','create_id','创建人id'),(654,'room_type','create_time','创建时间'),(655,'room_type','update_id','修改人id'),(656,'room_type','update_time','修改时间'),(657,'service_item','id','id'),(658,'service_item','tenant_id','租户id'),(659,'service_item','type_id','服务项目类别id'),(660,'service_item','name','服务名称'),(661,'service_item','charge_method','收费方式'),(662,'service_item','price','服务费用'),(663,'service_item','need_date','所需时间(分)'),(664,'service_item','state','管理状态'),(665,'service_item','create_id','创建人id'),(666,'service_item','create_time','创建时间'),(667,'service_item','update_id','修改人id'),(668,'service_item','update_time','修改时间'),(669,'service_type','id','id'),(670,'service_type','tenant_id','租户id'),(671,'service_type','name','服务项目名称'),(672,'service_type','state','管理状态'),(673,'service_type','create_id','创建人id'),(674,'service_type','create_time','创建时间'),(675,'service_type','update_id','修改人id'),(676,'service_type','update_time','修改时间'),(677,'set_dishes','id','id'),(678,'set_dishes','tenant_id','租户id'),(679,'set_dishes','set_id','餐饮套餐id'),(680,'set_dishes','dishes_id','菜品食物id'),(681,'set_dishes','create_id','创建人id'),(682,'set_dishes','create_time','创建时间'),(683,'set_dishes','update_id','修改人id'),(684,'set_dishes','update_time','修改时间'),(685,'source','id','id'),(686,'source','tenant_id','租户id'),(687,'source','name','来源渠道名称'),(688,'source','state','管理状态'),(689,'source','create_id','创建人id'),(690,'source','create_time','创建时间'),(691,'source','update_id','修改人id'),(692,'source','update_time','修改时间'),(693,'staff','id','id'),(694,'staff','tenant_id','租户id'),(695,'staff','role_id','角色id'),(696,'staff','name','姓名'),(697,'staff','id_num','身份证号'),(698,'staff','age','年龄'),(699,'staff','sex','性别(男/女)'),(700,'staff','phone','电话'),(701,'staff','email','邮箱'),(702,'staff','pass','密码'),(703,'staff','avator','头像'),(704,'staff','address','地址'),(705,'staff','status','状态：0-离职，1-在职'),(706,'staff','create_id','创建人id'),(707,'staff','create_time','创建时间'),(708,'staff','update_id','修改人id'),(709,'staff','update_time','修改时间'),(710,'tenant','id','id'),(711,'tenant','name','企业名称'),(712,'tenant','logo','企业logo'),(713,'tenant','contact_name','联系人姓名'),(714,'tenant','contact_phone','联系电话'),(715,'tenant','plan','套餐'),(716,'tenant','status','状态：0试用中 1正式 2锁定'),(717,'tenant','expire_time','套餐到期时间'),(718,'tenant','create_id','创建人id'),(719,'tenant','create_time','创建时间'),(720,'tenant','update_id','修改人id'),(721,'tenant','update_time','修改时间'),(722,'tenant','state','管理状态'),(723,'user','id','id'),(724,'user','union_id','微信UnionID（全局唯一）'),(725,'user','openid','微信OpenID（兜底匹配）'),(726,'user','phone','手机号（账号密码登录）'),(727,'user','pass','密码(md5)'),(728,'user','name','姓名'),(729,'user','avator','头像'),(730,'user','create_time','创建时间'),(731,'user','update_time','更新时间'),(732,'user','state','管理状态'),(733,'visit','id','id'),(734,'visit','tenant_id','租户id'),(735,'visit','elder_id','老人id'),(736,'visit','name','来访者姓名'),(737,'visit','id_card','身份证号（加密）'),(738,'visit','phone','来访者电话'),(739,'visit','relation','与老人关系'),(740,'visit','visit_time','来访时间'),(741,'visit','leave_time','离开时间'),(742,'visit','visit_num','来访者人数'),(743,'visit','status','来访状态'),(744,'visit','state','删除状态'),(745,'visit','create_id','创建人id'),(746,'visit','create_time','创建时间'),(747,'visit','update_id','修改人id'),(748,'visit','update_time','修改时间'),(749,'visit_plan','id','id'),(750,'visit_plan','tenant_id','租户id'),(751,'visit_plan','elder_id','老人id'),(752,'visit_plan','title','回访计划标题'),(753,'visit_plan','plan_date','计划回访时间'),(754,'visit_plan','content','回访计划内容'),(755,'visit_plan','complete_date','计划完成时间'),(756,'visit_plan','state','管理状态'),(757,'visit_plan','create_id','创建人id'),(758,'visit_plan','create_time','创建时间'),(759,'visit_plan','update_id','修改人id'),(760,'visit_plan','update_time','修改时间'),(761,'warehouse','id','id'),(762,'warehouse','tenant_id','租户id'),(763,'warehouse','staff_id','仓库管理员id'),(764,'warehouse','name','仓库名称'),(765,'warehouse','state','管理状态'),(766,'warehouse','create_id','创建人id'),(767,'warehouse','create_time','创建时间'),(768,'warehouse','update_id','修改人id'),(769,'warehouse','update_time','修改时间'),(770,'warehouse_material','id','id'),(771,'warehouse_material','tenant_id','租户id'),(772,'warehouse_material','warehouse_record_id','入库登记id'),(773,'warehouse_material','material_id','物资id'),(774,'warehouse_material','warehouse_num','入库数量'),(775,'warehouse_material','inventory','库存量'),(776,'warehouse_material','product_date','生产日期'),(777,'warehouse_material','expire_date','有效期'),(778,'warehouse_material','create_id','创建人id'),(779,'warehouse_material','create_time','创建时间'),(780,'warehouse_material','update_id','修改人id'),(781,'warehouse_material','update_time','修改时间'),(782,'warehouse_record','id','id'),(783,'warehouse_record','tenant_id','租户id'),(784,'warehouse_record','warehouse_id','仓库id'),(785,'warehouse_record','staff_id','经办人id'),(786,'warehouse_record','source','物资来源'),(787,'warehouse_record','warehouse_date','入库时间'),(788,'warehouse_record','status','入库状态'),(789,'warehouse_record','state','管理状态'),(790,'warehouse_record','create_id','创建人id'),(791,'warehouse_record','create_time','创建时间'),(792,'warehouse_record','update_id','修改人id'),(793,'warehouse_record','update_time','修改时间');
/*!40000 ALTER TABLE `field_dict` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `floor`
--

DROP TABLE IF EXISTS `floor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `floor` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `building_id` bigint unsigned NOT NULL COMMENT '楼栋id',
  `name` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '楼层名称',
  `room_num` int NOT NULL COMMENT '房间数量',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `floor`
--

LOCK TABLES `floor` WRITE;
/*!40000 ALTER TABLE `floor` DISABLE KEYS */;
INSERT INTO `floor` VALUES (1,1,1,'楼栋1-1层',10,0,1,'2023-01-02 14:22:54',1,'2023-01-02 14:22:57'),(2,1,1,'楼栋1-2层',10,0,1,'2023-01-02 14:22:54',1,'2023-01-04 22:23:22'),(3,1,1,'楼栋1-3层',10,0,1,'2023-01-02 14:22:54',1,'2023-01-02 14:22:57'),(4,1,1,'楼栋1-4层',10,0,1,'2023-01-02 14:22:54',1,'2023-01-02 14:22:57'),(5,1,1,'楼栋1-5层',10,0,1,'2023-01-02 14:22:54',1,'2023-01-02 14:22:57'),(6,1,2,'楼栋2-1层',2,0,1,'2023-01-04 21:22:36',1,'2023-01-04 22:23:49'),(7,1,2,'楼栋2-2层',1,0,1,'2023-01-04 21:22:47',1,'2023-01-04 22:22:40'),(8,1,2,'楼栋2-3层',12,1,1,'2023-04-05 02:16:26',1,'2023-04-05 02:18:48'),(9,1,2,'楼栋2-3层',12,0,1,'2023-04-05 02:19:05',1,'2023-04-05 02:19:05'),(10,1,13,'测试楼层1',2,0,1,'2023-04-05 02:34:54',1,'2023-04-05 02:37:33'),(11,1,13,'测试楼层2',1,0,1,'2023-04-05 02:35:26',1,'2023-04-05 02:35:26');
/*!40000 ALTER TABLE `floor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `health_data`
--

DROP TABLE IF EXISTS `health_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `health_data` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `height` int NOT NULL COMMENT '身高',
  `weight` double NOT NULL COMMENT '体重',
  `temperature` double NOT NULL COMMENT '体温',
  `heart_rate` int NOT NULL COMMENT '心率',
  `systolic_blood_pressure` int NOT NULL COMMENT '收缩血压',
  `diastolic_blood_pressure` int NOT NULL COMMENT '舒张血压',
  `fasting_blood_glucose` int NOT NULL COMMENT '空腹血糖',
  `postprandial_blood_glucose` int NOT NULL COMMENT '餐后血糖',
  `blood_oxygen_saturation` int NOT NULL COMMENT '血氧饱和度',
  `cholesterol` int NOT NULL COMMENT '总胆固醇',
  `uric_acid` int NOT NULL COMMENT '尿酸',
  `left_eye` double NOT NULL COMMENT '左眼',
  `right_eye` double NOT NULL COMMENT '右眼',
  `left_ear` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '左耳',
  `right_ear` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '右耳',
  `muscle_percentage` int NOT NULL COMMENT '肌肉率',
  `body_fat_percentage` int NOT NULL COMMENT '体脂率',
  `waist_circumference` int NOT NULL COMMENT '腰围',
  `hip_circumference` int NOT NULL COMMENT '臀围',
  `moisture_content` int NOT NULL COMMENT '水分率',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elder_id` (`elder_id`),
  KEY `idx_record_time` (`create_time`),
  KEY `idx_recorder_id` (`create_id`),
  CONSTRAINT `fk_health_data_elderly` FOREIGN KEY (`elder_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_health_data_recorder` FOREIGN KEY (`create_id`) REFERENCES `staff` (`id`),
  CONSTRAINT `fk_health_data_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='生命体征记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `health_data`
--

LOCK TABLES `health_data` WRITE;
/*!40000 ALTER TABLE `health_data` DISABLE KEYS */;
/*!40000 ALTER TABLE `health_data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `health_info`
--

DROP TABLE IF EXISTS `health_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `health_info` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `self_care` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '自理情况',
  `vision` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '视力',
  `hearing` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '听力',
  `hospital` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '主治医院',
  `doctor` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '主治医师',
  `phone` varchar(11) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '医院电话',
  `allergy_drug` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '过敏药物',
  `medical_history` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '病史',
  `major_disease` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '主要疾病',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='健康信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `health_info`
--

LOCK TABLES `health_info` WRITE;
/*!40000 ALTER TABLE `health_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `health_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `label`
--

DROP TABLE IF EXISTS `label`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `label` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `type_id` bigint unsigned NOT NULL COMMENT '类别id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标签名称',
  `color` varchar(15) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标签颜色',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `label`
--

LOCK TABLES `label` WRITE;
/*!40000 ALTER TABLE `label` DISABLE KEYS */;
INSERT INTO `label` VALUES (1,1,1,'跑步','rgb',0,1,'2023-01-03 21:14:18',1,'2023-01-03 22:53:35'),(2,1,1,'看书','rgb',0,1,'2023-01-03 21:14:18',1,'2023-01-03 21:14:22'),(3,1,2,'标签1','rgb',1,1,'2023-01-03 21:14:18',1,'2023-01-03 22:52:41'),(4,1,2,'标签2','rgb',0,1,'2023-01-03 21:14:18',1,'2023-01-03 21:14:22'),(5,1,2,'标签3','rgb',0,1,'2023-01-03 21:14:18',1,'2023-01-03 21:14:22'),(6,1,2,'标签4','rgb',0,1,'2023-01-03 22:50:38',1,'2023-01-03 22:50:38'),(7,1,2,'标签5','rgb',0,1,'2023-01-03 22:50:45',1,'2023-01-03 22:50:45'),(8,1,2,'标签6','rgb',0,1,'2023-01-03 22:50:47',1,'2023-01-03 22:50:47'),(9,1,2,'标签7','rgb',0,1,'2023-01-03 22:50:50',1,'2023-01-03 22:50:50'),(10,1,2,'标签8','rgb',0,1,'2023-01-03 22:50:53',1,'2023-01-03 22:50:53'),(11,1,2,'标签9','rgb',0,1,'2023-01-03 22:50:55',1,'2023-01-03 22:50:55'),(12,1,2,'标签10','rgb',0,1,'2023-01-03 22:51:00',1,'2023-01-03 22:51:00'),(13,1,1,'运动','rgb',0,1,'2023-01-03 22:52:02',1,'2023-01-03 22:52:02'),(14,1,2,'标签11','rgb',0,1,'2023-01-03 22:52:57',1,'2023-01-03 22:52:57');
/*!40000 ALTER TABLE `label` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `label_type`
--

DROP TABLE IF EXISTS `label_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `label_type` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分类名称',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `label_type`
--

LOCK TABLES `label_type` WRITE;
/*!40000 ALTER TABLE `label_type` DISABLE KEYS */;
INSERT INTO `label_type` VALUES (1,1,'分类1',0,1,'2023-01-03 21:13:23',1,'2023-01-03 22:48:39'),(2,1,'分类2',0,1,'2023-01-03 21:13:23',1,'2023-01-03 21:13:26'),(3,1,'分类3',0,1,'2023-01-03 22:22:00',1,'2023-01-03 22:22:00'),(4,1,'分类4',0,1,'2023-01-03 22:47:10',1,'2023-01-03 22:47:10'),(5,1,'分类5',0,1,'2023-01-03 22:47:13',1,'2023-01-03 22:47:13'),(6,1,'分类6',0,1,'2023-01-03 22:47:17',1,'2023-01-03 22:47:17'),(7,1,'分类7',0,1,'2023-01-03 22:47:20',1,'2023-01-03 22:47:20'),(8,1,'分类8',0,1,'2023-01-03 22:47:23',1,'2023-01-03 22:47:23'),(9,1,'分类9',0,1,'2023-01-03 22:47:27',1,'2023-01-03 22:47:27'),(10,1,'分类10',0,1,'2023-01-03 22:47:31',1,'2023-01-03 22:47:31'),(11,1,'分类11',0,1,'2023-01-03 22:48:46',1,'2023-01-03 22:48:46');
/*!40000 ALTER TABLE `label_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `material`
--

DROP TABLE IF EXISTS `material`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `material` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `type_id` bigint unsigned NOT NULL COMMENT '物资类别id',
  `name` varchar(15) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '物资名称',
  `price` bigint NOT NULL COMMENT '物资单价',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `material`
--

LOCK TABLES `material` WRITE;
/*!40000 ALTER TABLE `material` DISABLE KEYS */;
INSERT INTO `material` VALUES (1,1,1,'勺子',10,0,1,'2023-01-15 11:06:24',1,'2023-01-15 11:08:58'),(2,1,2,'当归',100,0,1,'2023-01-15 11:06:36',1,'2023-01-15 11:07:51'),(3,1,1,'盘子',10,0,1,'2023-01-15 11:07:03',1,'2023-01-15 11:09:29');
/*!40000 ALTER TABLE `material` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `material_type`
--

DROP TABLE IF EXISTS `material_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `material_type` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '物资类别名称',
  `kind` tinyint NOT NULL DEFAULT '99' COMMENT '分类用途：1=床型，99=设施/其他',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `material_type`
--

LOCK TABLES `material_type` WRITE;
/*!40000 ALTER TABLE `material_type` DISABLE KEYS */;
INSERT INTO `material_type` VALUES (1,1,'餐具',99,0,1,'2023-01-15 11:03:17',1,'2023-01-15 11:05:31'),(2,1,'药品',99,0,1,'2023-01-15 11:03:57',1,'2023-01-15 11:05:24');
/*!40000 ALTER TABLE `material_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `medicine`
--

DROP TABLE IF EXISTS `medicine`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `medicine` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `name` varchar(15) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '药品名称',
  `type` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '药品类别',
  `specification` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '药品规格',
  `dosage_form` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '药品剂型',
  `manufacturer` varchar(25) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '生产厂家',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `medicine`
--

LOCK TABLES `medicine` WRITE;
/*!40000 ALTER TABLE `medicine` DISABLE KEYS */;
/*!40000 ALTER TABLE `medicine` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `medicine_record`
--

DROP TABLE IF EXISTS `medicine_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `medicine_record` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `deposit_info_id` bigint unsigned NOT NULL COMMENT '药品缴存信息id',
  `medicine_time` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用药时间（早/中/晚）',
  `medicine_date` datetime NOT NULL COMMENT '用药日期',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `medicine_record`
--

LOCK TABLES `medicine_record` WRITE;
/*!40000 ALTER TABLE `medicine_record` DISABLE KEYS */;
/*!40000 ALTER TABLE `medicine_record` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `medicine_set`
--

DROP TABLE IF EXISTS `medicine_set`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `medicine_set` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `deposit_info_id` bigint unsigned NOT NULL COMMENT '药品缴存信息id',
  `medicine_time` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用药时间（餐前/餐后）',
  `day_frequency` int NOT NULL COMMENT '天频率',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `medicine_set`
--

LOCK TABLES `medicine_set` WRITE;
/*!40000 ALTER TABLE `medicine_set` DISABLE KEYS */;
/*!40000 ALTER TABLE `medicine_set` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `member`
--

DROP TABLE IF EXISTS `member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `member` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint unsigned NOT NULL COMMENT '全局用户id',
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户id',
  `role_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '角色编号(关联role)',
  `permissions` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '预留：细粒度权限',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0在职 1离职',
  `create_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_user_tenant` (`user_id`,`tenant_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='成员关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `member`
--

LOCK TABLES `member` WRITE;
/*!40000 ALTER TABLE `member` DISABLE KEYS */;
INSERT INTO `member` VALUES (1,1,1,1,'',0,0,'2022-12-31 12:34:43',0,'2022-12-31 12:34:43',0);
/*!40000 ALTER TABLE `member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `nurse`
--

DROP TABLE IF EXISTS `nurse`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `nurse` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `staff_id` bigint unsigned NOT NULL COMMENT '护理人员id',
  `nurse_date` datetime NOT NULL COMMENT '护理时间',
  `complete_status` tinyint NOT NULL COMMENT '护理完成情况',
  `dine_status` tinyint NOT NULL COMMENT '进餐情况',
  `rest` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '休息',
  `take_medicine` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '服药',
  `active` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '活动',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `nurse`
--

LOCK TABLES `nurse` WRITE;
/*!40000 ALTER TABLE `nurse` DISABLE KEYS */;
/*!40000 ALTER TABLE `nurse` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `nurse_grade`
--

DROP TABLE IF EXISTS `nurse_grade`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `nurse_grade` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '级别名称',
  `type` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '护理类型',
  `month_price` bigint NOT NULL COMMENT '月护理费用',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `nurse_grade`
--

LOCK TABLES `nurse_grade` WRITE;
/*!40000 ALTER TABLE `nurse_grade` DISABLE KEYS */;
INSERT INTO `nurse_grade` VALUES (1,1,'一级护理','自理',1000,0,1,'2023-01-12 09:13:25',1,'2023-04-04 00:25:17'),(2,1,'二级护理','自理',1000,0,1,'2023-01-12 09:14:16',1,'2023-01-12 09:15:01'),(3,1,'三级护理','自理',1000,0,1,'2023-02-01 12:15:12',1,'2023-02-01 12:15:12'),(4,1,'四级护理','自理',1200,1,1,'2023-04-04 00:17:10',1,'2023-04-04 00:17:27'),(5,1,'测试','介护',1200,0,1,'2023-04-04 15:12:40',1,'2023-04-04 15:12:54'),(6,1,'胡图图','自理',1000,0,1,'2023-04-14 17:53:27',1,'2023-04-14 17:53:27');
/*!40000 ALTER TABLE `nurse_grade` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `nurse_group`
--

DROP TABLE IF EXISTS `nurse_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `nurse_group` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `staff_id` bigint unsigned NOT NULL COMMENT '护工小组组长id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '护工小组名称',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `nurse_group`
--

LOCK TABLES `nurse_group` WRITE;
/*!40000 ALTER TABLE `nurse_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `nurse_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `nurse_group_member`
--

DROP TABLE IF EXISTS `nurse_group_member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `nurse_group_member` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `group_id` bigint unsigned NOT NULL COMMENT '护工小组id',
  `staff_id` bigint unsigned NOT NULL COMMENT '护工id',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `nurse_group_member`
--

LOCK TABLES `nurse_group_member` WRITE;
/*!40000 ALTER TABLE `nurse_group_member` DISABLE KEYS */;
/*!40000 ALTER TABLE `nurse_group_member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `nurse_item`
--

DROP TABLE IF EXISTS `nurse_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `nurse_item` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `grade_id` bigint unsigned NOT NULL COMMENT '护理等级id',
  `service_id` bigint unsigned NOT NULL COMMENT '服务项目id',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `nurse_item`
--

LOCK TABLES `nurse_item` WRITE;
/*!40000 ALTER TABLE `nurse_item` DISABLE KEYS */;
INSERT INTO `nurse_item` VALUES (3,1,2,1,1,'2023-01-12 09:14:16',1,'2023-01-12 09:14:16'),(6,1,3,1,1,'2023-02-01 12:15:12',1,'2023-02-01 12:15:12'),(7,1,3,2,1,'2023-02-01 12:15:12',1,'2023-02-01 12:15:12'),(22,1,4,16,1,'2023-04-04 00:17:10',1,'2023-04-04 00:17:10'),(28,1,1,16,1,'2023-04-04 00:25:17',1,'2023-04-04 00:25:17'),(29,1,1,14,1,'2023-04-04 00:25:17',1,'2023-04-04 00:25:17'),(30,1,1,11,1,'2023-04-04 00:25:17',1,'2023-04-04 00:25:17'),(31,1,1,9,1,'2023-04-04 00:25:17',1,'2023-04-04 00:25:17'),(32,1,1,8,1,'2023-04-04 00:25:17',1,'2023-04-04 00:25:17'),(43,1,5,11,1,'2023-04-04 15:12:54',1,'2023-04-04 15:12:54'),(44,1,5,9,1,'2023-04-04 15:12:54',1,'2023-04-04 15:12:54'),(45,1,5,8,1,'2023-04-04 15:12:54',1,'2023-04-04 15:12:54'),(46,1,5,7,1,'2023-04-04 15:12:54',1,'2023-04-04 15:12:54'),(47,1,5,6,1,'2023-04-04 15:12:54',1,'2023-04-04 15:12:54'),(48,1,5,5,1,'2023-04-04 15:12:54',1,'2023-04-04 15:12:54'),(49,1,5,4,1,'2023-04-04 15:12:54',1,'2023-04-04 15:12:54'),(50,1,5,3,1,'2023-04-04 15:12:54',1,'2023-04-04 15:12:54'),(51,1,5,16,1,'2023-04-04 15:12:54',1,'2023-04-04 15:12:54'),(52,1,5,14,1,'2023-04-04 15:12:54',1,'2023-04-04 15:12:54'),(53,1,6,16,1,'2023-04-14 17:53:27',1,'2023-04-14 17:53:27');
/*!40000 ALTER TABLE `nurse_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `nurse_reserve`
--

DROP TABLE IF EXISTS `nurse_reserve`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `nurse_reserve` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `staff_id` bigint unsigned DEFAULT NULL COMMENT '服务人id',
  `service_name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '服务项目名称',
  `need_date` int NOT NULL COMMENT '所需时间',
  `service_price` bigint NOT NULL COMMENT '服务费用',
  `charge_method` varchar(3) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '收费方式',
  `frequency` int NOT NULL COMMENT '服务次数',
  `pay_amount` bigint NOT NULL COMMENT '支付总额',
  `status` tinyint NOT NULL COMMENT '订单状态',
  `nurse_date` datetime DEFAULT NULL COMMENT '护理时间',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_nurse_tenant` (`tenant_id`,`elder_id`,`create_time`) /*!80000 INVISIBLE */,
  KEY `idx_nurse_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `nurse_reserve`
--

LOCK TABLES `nurse_reserve` WRITE;
/*!40000 ALTER TABLE `nurse_reserve` DISABLE KEYS */;
INSERT INTO `nurse_reserve` VALUES (1,1,2,5,'血糖监测',30,120,'按次',2,240,1,'2022-12-13 00:00:00',1,'2023-01-12 23:30:13',1,'2023-01-12 23:44:33'),(4,1,1,6,'视力检测',12,12,'按次',1,12,1,'2023-04-04 00:00:00',1,'2023-04-04 10:11:12',1,'2023-04-04 10:14:06'),(5,1,1,NULL,'视力检测',12,12,'按次',1,12,0,NULL,1,'2023-04-07 17:18:18',1,'2023-04-07 17:18:18'),(6,1,1,1,'视力检测',12,12,'按次',1,12,1,'2023-04-24 00:00:00',1,'2023-04-14 17:55:10',1,'2023-04-24 19:36:54'),(7,1,1,NULL,'视力检测',12,12,'按次',4,48,0,NULL,1,'2023-04-14 17:55:24',1,'2023-04-14 17:55:24');
/*!40000 ALTER TABLE `nurse_reserve` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order`
--

DROP TABLE IF EXISTS `order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `staff_id` bigint unsigned DEFAULT NULL COMMENT '送餐人id',
  `deliver_dishes_date` datetime DEFAULT NULL COMMENT '送餐时间',
  `dine_date` datetime NOT NULL COMMENT '就餐时间',
  `dine_type` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '就餐方式',
  `pay_amount` bigint NOT NULL COMMENT '支付总额',
  `status` tinyint NOT NULL COMMENT '订单状态',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_order_tenant` (`tenant_id`,`elder_id`,`create_time`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order`
--

LOCK TABLES `order` WRITE;
/*!40000 ALTER TABLE `order` DISABLE KEYS */;
INSERT INTO `order` VALUES (1,1,2,5,'2022-12-13 00:00:00','2022-12-14 00:00:00','送餐',6,1,1,'2023-01-14 20:51:41',1,'2023-01-14 21:09:05'),(2,1,1,6,'2023-04-05 00:00:00','2023-04-04 00:00:00','送餐',3,1,1,'2023-04-04 16:34:45',1,'2023-04-04 16:35:29'),(3,1,1,6,'2023-04-12 00:00:00','2023-04-13 00:00:00','堂食',2,1,1,'2023-04-05 19:34:09',1,'2023-04-05 19:34:28'),(4,1,1,NULL,NULL,'2023-04-07 17:23:55','送餐',0,0,1,'2023-04-07 17:24:16',1,'2023-04-07 17:24:16'),(5,1,1,NULL,NULL,'2023-04-22 00:00:00','送餐',10,0,1,'2023-04-14 18:14:39',1,'2023-04-14 18:14:39');
/*!40000 ALTER TABLE `order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_dishes`
--

DROP TABLE IF EXISTS `order_dishes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_dishes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `order_id` bigint unsigned NOT NULL COMMENT '订餐id',
  `dishes_name` varchar(15) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜品名称',
  `dishes_price` bigint NOT NULL COMMENT '菜品价格',
  `order_num` int NOT NULL COMMENT '菜品份数',
  `status` tinyint NOT NULL COMMENT '套餐标记',
  `total_amount` bigint NOT NULL COMMENT '菜品总额',
  `really_amount` bigint NOT NULL COMMENT '实际总额',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_order_dishes` (`order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_dishes`
--

LOCK TABLES `order_dishes` WRITE;
/*!40000 ALTER TABLE `order_dishes` DISABLE KEYS */;
INSERT INTO `order_dishes` VALUES (1,1,1,'粥',2,4,1,8,6,1,'2023-01-14 20:51:41',1,'2023-01-14 20:51:41'),(2,1,2,'粥',2,1,1,2,0,1,'2023-04-04 16:34:45',1,'2023-04-04 16:34:45'),(3,1,2,'粥',2,1,0,2,2,1,'2023-04-04 16:34:45',1,'2023-04-04 16:34:45'),(4,1,2,'包子',1,2,1,2,1,1,'2023-04-04 16:34:45',1,'2023-04-04 16:34:45'),(5,1,3,'粥',2,1,0,2,2,1,'2023-04-05 19:34:09',1,'2023-04-05 19:34:09'),(6,1,4,'包子',1,1,1,1,0,1,'2023-04-07 17:24:16',1,'2023-04-07 17:24:16'),(7,1,5,'牛爷爷',10,1,0,10,10,1,'2023-04-14 18:14:39',1,'2023-04-14 18:14:39');
/*!40000 ALTER TABLE `order_dishes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `outbound_material`
--

DROP TABLE IF EXISTS `outbound_material`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `outbound_material` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `outbound_record_id` bigint unsigned NOT NULL COMMENT '出库登记id',
  `warehouse_material_id` bigint unsigned NOT NULL COMMENT '入库物资id',
  `material_id` bigint unsigned NOT NULL COMMENT '物资id',
  `outbound_num` int NOT NULL COMMENT '出库数量',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `outbound_material`
--

LOCK TABLES `outbound_material` WRITE;
/*!40000 ALTER TABLE `outbound_material` DISABLE KEYS */;
INSERT INTO `outbound_material` VALUES (4,1,7,1,1,2,1,'2023-01-31 21:32:29',1,'2023-01-31 21:32:29'),(5,1,7,2,2,1,1,'2023-01-31 21:32:30',1,'2023-01-31 21:32:30'),(6,1,9,3,1,1,1,'2023-01-31 21:34:36',1,'2023-01-31 21:34:36'),(7,1,9,4,2,1,1,'2023-01-31 21:34:36',1,'2023-01-31 21:34:36'),(8,1,10,3,1,1,1,'2023-01-31 22:18:00',1,'2023-01-31 22:18:00');
/*!40000 ALTER TABLE `outbound_material` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `outbound_record`
--

DROP TABLE IF EXISTS `outbound_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `outbound_record` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `warehouse_id` bigint unsigned NOT NULL COMMENT '仓库id',
  `staff_id` bigint unsigned NOT NULL COMMENT '经办人id',
  `recipient_id` bigint unsigned NOT NULL COMMENT '领用人id',
  `recipient_type` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '领用人类型',
  `material_use` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '物资去向',
  `outbound_date` datetime NOT NULL COMMENT '出库时间',
  `status` tinyint NOT NULL COMMENT '出库状态',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `outbound_record`
--

LOCK TABLES `outbound_record` WRITE;
/*!40000 ALTER TABLE `outbound_record` DISABLE KEYS */;
INSERT INTO `outbound_record` VALUES (7,1,1,6,2,'老人','使用','2022-12-14 00:00:00',2,0,1,'2023-01-31 21:32:29',1,'2023-01-31 21:44:09'),(9,1,1,6,2,'员工','使用','2022-12-14 00:00:00',0,1,1,'2023-01-31 21:34:36',1,'2023-01-31 21:43:15'),(10,1,1,6,2,'员工','使用','2022-12-14 00:00:00',-1,0,1,'2023-01-31 22:18:00',1,'2023-01-31 22:18:44');
/*!40000 ALTER TABLE `outbound_record` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `outward`
--

DROP TABLE IF EXISTS `outward`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `outward` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `chaperone_name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '陪同人姓名',
  `chaperone_phone` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '陪同人电话',
  `chaperone_type` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '陪同人类型（家属/护工）',
  `outward_date` datetime NOT NULL COMMENT '外出时间',
  `plan_return_date` datetime NOT NULL COMMENT '计划返回时间',
  `real_return_date` datetime DEFAULT NULL COMMENT '实际返回时间',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_outward_tenant` (`tenant_id`,`elder_id`,`outward_date`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `outward`
--

LOCK TABLES `outward` WRITE;
/*!40000 ALTER TABLE `outward` DISABLE KEYS */;
INSERT INTO `outward` VALUES (1,1,1,'张三','13546473658','家属','2022-12-12 00:00:00','2022-12-14 00:00:00','2022-12-13 00:00:00',1,1,'2023-02-03 15:02:36',1,'2023-02-03 21:32:02'),(2,1,17,'张三','13547584403','护工','2023-04-01 00:00:00','2023-04-26 00:00:00','2023-04-26 00:00:00',0,1,'2023-04-01 10:00:08',1,'2023-04-26 15:52:03'),(3,1,2,'张三','13547584403','护工','2023-04-01 00:00:00','2023-05-04 00:00:00',NULL,1,1,'2023-04-01 10:25:20',1,'2023-04-01 11:29:20'),(4,1,4,'张三','13547584403','护工','2023-04-01 00:00:00','2023-05-02 00:00:00','2023-04-01 00:00:00',0,1,'2023-04-01 10:28:40',1,'2023-04-01 10:57:56'),(5,1,1,'文帝','13547584400','护工','2023-04-05 00:00:00','2023-04-27 00:00:00','2023-04-25 00:00:00',0,1,'2023-04-14 17:04:19',1,'2023-04-25 09:28:24');
/*!40000 ALTER TABLE `outward` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reserve`
--

DROP TABLE IF EXISTS `reserve`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reserve` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `staff_id` bigint unsigned NOT NULL COMMENT '销售人员id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '交款人姓名',
  `phone` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '交款人电话',
  `due_date` datetime NOT NULL COMMENT '预定到期时间',
  `deposit` bigint NOT NULL COMMENT '定金',
  `status` tinyint NOT NULL COMMENT '退款状态（N/Y）',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reserve`
--

LOCK TABLES `reserve` WRITE;
/*!40000 ALTER TABLE `reserve` DISABLE KEYS */;
INSERT INTO `reserve` VALUES (25,1,6,2,'曹氏','13547584480','2023-04-29 00:00:00',1000,1,1,'2023-04-04 08:52:13',1,'2023-04-23 14:34:34'),(26,1,7,2,'张三','17666666666','2028-04-21 00:00:00',1,1,1,'2023-04-05 19:34:11',1,'2023-04-14 10:43:56'),(29,1,89,1,'小红','13881469052','2023-05-02 00:00:00',500,0,1,'2023-04-24 14:50:51',1,'2023-04-24 14:50:51');
/*!40000 ALTER TABLE `reserve` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `retreat`
--

DROP TABLE IF EXISTS `retreat`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `retreat` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `retreat_form` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '退住形式',
  `evaluate` int NOT NULL COMMENT '对老人评价',
  `retreat_cause` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '退住原因',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `retreat`
--

LOCK TABLES `retreat` WRITE;
/*!40000 ALTER TABLE `retreat` DISABLE KEYS */;
/*!40000 ALTER TABLE `retreat` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `retreat_apply`
--

DROP TABLE IF EXISTS `retreat_apply`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `retreat_apply` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `status` tinyint NOT NULL COMMENT '业务状态(退住申请流程状态)',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=正常',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `retreat_apply`
--

LOCK TABLES `retreat_apply` WRITE;
/*!40000 ALTER TABLE `retreat_apply` DISABLE KEYS */;
INSERT INTO `retreat_apply` VALUES (3,1,1,1,-1,1,'2023-02-04 22:46:29',1,'2023-02-04 22:46:29'),(4,1,1,1,-1,1,'2023-02-13 17:42:29',1,'2023-02-13 17:42:45'),(5,1,1,1,-1,1,'2023-02-13 17:42:57',1,'2023-02-13 17:43:23'),(6,1,1,1,2,1,'2023-04-02 01:03:34',1,'2023-04-14 18:18:29'),(7,1,1,2,-1,1,'2023-04-04 19:31:00',1,'2023-04-04 19:37:08'),(8,1,1,2,-1,1,'2023-04-05 16:58:31',1,'2023-04-05 17:10:25'),(9,1,1,2,2,1,'2023-04-14 17:35:41',1,'2023-04-14 18:19:12'),(10,1,1,88,0,1,'2023-04-23 19:27:18',1,'2023-04-23 19:27:18'),(11,1,1,4,0,1,'2023-04-23 19:28:37',1,'2023-04-23 19:28:37'),(12,1,1,8,0,1,'2023-04-26 15:53:50',1,'2023-04-26 15:53:50');
/*!40000 ALTER TABLE `retreat_apply` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint NOT NULL COMMENT '租户ID（平台角色tenant_id=0）',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名称',
  `permissions` json DEFAULT NULL COMMENT '权限列表',
  `is_system` tinyint NOT NULL DEFAULT '0' COMMENT '是否系统预置：0-否，1-是',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tenant_id` (`tenant_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,0,'超级管理员',NULL,0,1,'2022-12-31 22:40:22',1,'2022-12-31 22:40:29'),(2,0,'销售管理员',NULL,0,1,'2023-01-02 18:23:58',1,'2023-01-02 18:24:01'),(3,0,'入住管理员',NULL,0,1,'2023-01-02 18:26:12',1,'2023-01-02 18:26:15'),(4,0,'人事管理员',NULL,0,1,'2023-01-02 18:26:12',1,'2023-01-02 18:26:15'),(5,0,'服务管理员',NULL,0,1,'2023-01-02 18:26:12',1,'2023-01-02 18:26:15'),(6,0,'仓库管理员',NULL,0,1,'2023-01-02 18:26:12',1,'2023-01-02 18:26:15'),(7,0,'餐饮管理员',NULL,0,1,'2023-01-02 18:26:12',1,'2023-01-02 18:26:15'),(8,0,'财务管理员',NULL,0,1,'2023-01-02 18:26:12',1,'2023-01-02 18:26:15');
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_auth`
--

DROP TABLE IF EXISTS `role_auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role_auth` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `role_id` bigint unsigned NOT NULL COMMENT '角色id',
  `auth_id` bigint unsigned NOT NULL COMMENT '权限id',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=83 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_auth`
--

LOCK TABLES `role_auth` WRITE;
/*!40000 ALTER TABLE `role_auth` DISABLE KEYS */;
INSERT INTO `role_auth` VALUES (1,1,1,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(2,1,2,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(3,1,3,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(4,1,4,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(5,1,5,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(6,1,6,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(7,1,7,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(8,1,8,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(9,1,9,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(10,1,10,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(11,1,11,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(12,1,12,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(13,1,13,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(14,1,14,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(15,1,15,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(16,1,16,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(17,1,17,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(18,1,18,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(19,1,19,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(20,1,20,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(21,1,21,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(22,1,22,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(23,1,23,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(24,1,24,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(25,1,25,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(26,1,26,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(27,1,27,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(28,1,28,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(29,1,29,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(30,1,30,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(31,1,31,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(32,1,32,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(33,1,33,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(34,1,34,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(35,1,35,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(36,1,36,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(37,1,37,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(38,1,38,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(39,1,39,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(40,1,40,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(41,1,41,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(42,1,42,1,'2023-01-03 14:07:52',1,'2023-01-03 14:07:52'),(43,2,1,1,'2023-01-03 14:10:09',1,'2023-01-03 14:10:09'),(44,2,2,1,'2023-01-03 14:10:09',1,'2023-01-03 14:10:09'),(45,2,3,1,'2023-01-03 14:10:09',1,'2023-01-03 14:10:09'),(46,2,4,1,'2023-01-03 14:10:09',1,'2023-01-03 14:10:09'),(47,2,5,1,'2023-01-03 14:10:43',1,'2023-01-03 14:10:43'),(48,3,1,1,'2023-01-03 14:11:34',1,'2023-01-03 14:11:34'),(49,3,6,1,'2023-01-03 14:12:05',1,'2023-01-03 14:12:05'),(50,3,7,1,'2023-01-03 14:12:05',1,'2023-01-03 14:12:05'),(51,3,8,1,'2023-01-03 14:12:05',1,'2023-01-03 14:12:05'),(52,3,9,1,'2023-01-03 14:12:05',1,'2023-01-03 14:12:05'),(53,3,10,1,'2023-01-03 14:12:05',1,'2023-01-03 14:12:05'),(54,3,11,1,'2023-01-03 14:12:05',1,'2023-01-03 14:12:05'),(55,3,12,1,'2023-01-03 14:12:05',1,'2023-01-03 14:12:05'),(56,4,1,1,'2023-01-03 14:12:52',1,'2023-01-03 14:12:52'),(57,4,13,1,'2023-01-03 14:13:27',1,'2023-01-03 14:13:27'),(58,4,14,1,'2023-01-03 14:13:27',1,'2023-01-03 14:13:27'),(59,4,15,1,'2023-01-03 14:13:27',1,'2023-01-03 14:13:27'),(60,4,16,1,'2023-01-03 14:13:27',1,'2023-01-03 14:13:27'),(61,5,1,1,'2023-01-03 14:13:51',1,'2023-01-03 14:13:51'),(62,5,17,1,'2023-01-03 14:14:09',1,'2023-01-03 14:14:09'),(63,5,18,1,'2023-01-03 14:14:09',1,'2023-01-03 14:14:09'),(64,5,19,1,'2023-01-03 14:14:09',1,'2023-01-03 14:14:09'),(65,5,20,1,'2023-01-03 14:14:09',1,'2023-01-03 14:14:09'),(66,6,1,1,'2023-01-03 14:14:52',1,'2023-01-03 14:14:52'),(67,6,21,1,'2023-01-03 14:15:17',1,'2023-01-03 14:15:17'),(68,6,22,1,'2023-01-03 14:15:17',1,'2023-01-03 14:15:17'),(69,6,23,1,'2023-01-03 14:15:17',1,'2023-01-03 14:15:17'),(70,6,24,1,'2023-01-03 14:15:17',1,'2023-01-03 14:15:17'),(71,6,25,1,'2023-01-03 14:15:17',1,'2023-01-03 14:15:17'),(72,6,26,1,'2023-01-03 14:15:17',1,'2023-01-03 14:15:17'),(73,7,1,1,'2023-01-03 14:15:42',1,'2023-01-03 14:15:42'),(74,7,27,1,'2023-01-03 14:16:08',1,'2023-01-03 14:16:08'),(75,7,28,1,'2023-01-03 14:16:08',1,'2023-01-03 14:16:08'),(76,7,29,1,'2023-01-03 14:16:08',1,'2023-01-03 14:16:08'),(77,7,30,1,'2023-01-03 14:16:08',1,'2023-01-03 14:16:08'),(78,8,1,1,'2023-01-03 14:17:14',1,'2023-01-03 14:17:14'),(79,8,31,1,'2023-01-03 14:17:32',1,'2023-01-03 14:17:32'),(80,8,32,1,'2023-01-03 14:17:32',1,'2023-01-03 14:17:32'),(81,8,33,1,'2023-01-03 14:17:32',1,'2023-01-03 14:17:32'),(82,8,34,1,'2023-01-03 14:17:32',1,'2023-01-03 14:17:32');
/*!40000 ALTER TABLE `role_auth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `room`
--

DROP TABLE IF EXISTS `room`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `room` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `type_id` bigint unsigned NOT NULL COMMENT '房间类型id',
  `build_id` bigint unsigned NOT NULL COMMENT '楼栋id',
  `floor_id` bigint unsigned NOT NULL COMMENT '楼层id',
  `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '房间名称',
  `bed_num` int NOT NULL COMMENT '床位数量',
  `type` tinyint DEFAULT '1' COMMENT '房间类型：1-单人间，2-双人间，3-多人间',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-空闲，1-部分占用，2-全满，3-维修',
  `price` bigint DEFAULT '0' COMMENT '房间基准价（分）',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `room`
--

LOCK TABLES `room` WRITE;
/*!40000 ALTER TABLE `room` DISABLE KEYS */;
INSERT INTO `room` VALUES (1,1,1,0,1,'爱心楼-1层-1房',1,1,0,0,0,1,'2023-01-02 14:25:33',1,'2023-01-02 14:25:38'),(2,1,1,0,1,'爱心楼-1层-2房',1,1,0,0,0,1,'2023-01-02 14:25:33',1,'2023-01-02 14:25:38'),(3,1,1,0,1,'爱心楼-1层-3房',1,1,0,0,0,1,'2023-01-02 14:25:33',1,'2023-01-02 14:25:38'),(4,1,1,0,1,'爱心楼-1层-4房',1,1,0,0,0,1,'2023-01-02 14:25:33',1,'2023-01-02 14:25:38'),(5,1,1,0,1,'爱心楼-1层-5房',1,1,0,0,0,1,'2023-01-02 14:25:33',1,'2023-04-05 00:33:40'),(6,1,1,0,6,'楼栋2-1层-1房',2,1,0,0,0,1,'2023-01-04 21:37:02',1,'2023-01-04 22:24:35'),(7,1,1,0,6,'楼栋2-1层-2房',1,1,0,0,0,1,'2023-01-04 21:38:22',1,'2023-01-04 22:23:49'),(8,1,2,0,10,'测试房间1',2,1,0,0,0,1,'2023-04-05 02:37:04',1,'2023-04-05 02:38:47'),(9,1,1,0,10,'测试房间2',1,1,0,0,0,1,'2023-04-05 02:37:47',1,'2023-04-05 02:37:47');
/*!40000 ALTER TABLE `room` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `room_material`
--

DROP TABLE IF EXISTS `room_material`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `room_material` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `room_id` bigint unsigned NOT NULL COMMENT '房间id',
  `material_type_id` bigint unsigned NOT NULL COMMENT '设施(物资分类)编号，kind=0',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态(Y/N)',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_room` (`room_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `room_material`
--

LOCK TABLES `room_material` WRITE;
/*!40000 ALTER TABLE `room_material` DISABLE KEYS */;
/*!40000 ALTER TABLE `room_material` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `room_transfer`
--

DROP TABLE IF EXISTS `room_transfer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `room_transfer` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人ID',
  `from_bed_id` bigint unsigned NOT NULL COMMENT '原床位ID',
  `to_bed_id` bigint unsigned NOT NULL COMMENT '新床位ID',
  `transfer_date` date NOT NULL COMMENT '转房日期',
  `reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '原因',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elder_id` (`elder_id`),
  KEY `idx_from_bed_id` (`from_bed_id`),
  KEY `idx_to_bed_id` (`to_bed_id`),
  CONSTRAINT `fk_room_transfers_elderly` FOREIGN KEY (`elder_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_room_transfers_from_bed` FOREIGN KEY (`from_bed_id`) REFERENCES `bed` (`id`),
  CONSTRAINT `fk_room_transfers_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`),
  CONSTRAINT `fk_room_transfers_to_bed` FOREIGN KEY (`to_bed_id`) REFERENCES `bed` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='转房记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `room_transfer`
--

LOCK TABLES `room_transfer` WRITE;
/*!40000 ALTER TABLE `room_transfer` DISABLE KEYS */;
/*!40000 ALTER TABLE `room_transfer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `room_type`
--

DROP TABLE IF EXISTS `room_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `room_type` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '房间类型名称',
  `month_price` bigint NOT NULL COMMENT '月房间费用',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `room_type`
--

LOCK TABLES `room_type` WRITE;
/*!40000 ALTER TABLE `room_type` DISABLE KEYS */;
INSERT INTO `room_type` VALUES (1,1,'单人间',600,0,1,'2023-01-02 14:20:09',1,'2023-01-04 11:47:33'),(2,1,'双人间',500,0,1,'2023-01-04 11:46:31',1,'2023-01-04 11:46:31'),(3,1,'四人间',120,1,1,'2023-04-04 20:07:36',1,'2023-04-04 20:08:08'),(4,1,'胡图图',1000,0,1,'2023-04-14 18:22:33',1,'2023-04-14 18:22:33'),(5,1,'打答复',100,0,1,'2023-04-14 18:24:51',1,'2023-04-14 18:24:51');
/*!40000 ALTER TABLE `room_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_item`
--

DROP TABLE IF EXISTS `service_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_item` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `type_id` bigint unsigned NOT NULL COMMENT '服务项目类别id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '服务名称',
  `charge_method` varchar(3) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '收费方式',
  `price` bigint NOT NULL COMMENT '服务费用',
  `need_date` int NOT NULL COMMENT '所需时间(分)',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_item`
--

LOCK TABLES `service_item` WRITE;
/*!40000 ALTER TABLE `service_item` DISABLE KEYS */;
INSERT INTO `service_item` VALUES (1,1,1,'服务1','按月',50,5,1,1,'2023-01-10 16:32:27',1,'2023-01-10 17:04:04'),(2,1,2,'心电检测','按次',120,30,0,1,'2023-01-10 16:32:36',1,'2023-01-10 16:32:36'),(3,1,1,'服务2','按月',120,30,0,1,'2023-01-10 16:33:13',1,'2023-01-10 16:33:13'),(4,1,1,'服务3','按月',120,30,0,1,'2023-01-10 16:46:18',1,'2023-01-10 16:46:18'),(5,1,1,'服务4','按月',120,30,0,1,'2023-01-10 16:46:22',1,'2023-01-10 16:46:22'),(6,1,1,'服务5','按月',120,30,0,1,'2023-01-10 16:46:24',1,'2023-01-10 16:46:24'),(7,1,1,'服务6','按月',120,30,0,1,'2023-01-10 16:46:30',1,'2023-01-10 16:46:30'),(8,1,1,'服务7','按月',120,30,0,1,'2023-01-10 16:46:34',1,'2023-01-10 16:46:34'),(9,1,1,'服务8','按月',120,30,0,1,'2023-01-10 16:46:36',1,'2023-01-10 16:46:36'),(10,1,1,'服务9','按月',120,30,1,1,'2023-01-10 16:46:39',1,'2023-01-10 17:00:52'),(11,1,1,'服务10','按月',120,30,1,1,'2023-01-10 16:46:43',1,'2023-04-14 17:52:08'),(12,1,1,'服务11','按月',120,30,1,1,'2023-01-10 16:46:56',1,'2023-04-03 15:58:06'),(13,1,1,'血糖监测','按次',120,30,0,1,'2023-01-10 16:47:23',1,'2023-04-03 17:16:23'),(14,1,2,'翻身','按月',120,30,0,1,'2023-01-12 13:13:47',1,'2023-01-12 13:14:37'),(15,1,2,'血脂检测','按次',12,24,0,1,'2023-04-03 17:20:49',1,'2023-04-03 17:20:49'),(16,1,1,'血压监测','按月',12,12,0,1,'2023-04-03 17:21:32',1,'2023-04-14 17:51:31'),(17,1,1,'视力检测','按次',12,12,0,1,'2023-04-03 17:24:09',1,'2023-04-03 17:24:59'),(18,1,2,'胡图图1','按月',244,1,0,1,'2023-04-14 17:50:43',1,'2023-04-14 17:51:22');
/*!40000 ALTER TABLE `service_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_type`
--

DROP TABLE IF EXISTS `service_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_type` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '服务项目名称',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_type`
--

LOCK TABLES `service_type` WRITE;
/*!40000 ALTER TABLE `service_type` DISABLE KEYS */;
INSERT INTO `service_type` VALUES (1,1,'饮食服务',0,1,'2023-01-10 16:27:58',1,'2023-04-26 13:23:43'),(2,1,'生活起居服务',0,1,'2023-01-10 16:30:36',1,'2023-04-26 13:23:54'),(3,1,'医疗护理服务',0,1,'2023-01-10 16:43:04',1,'2023-04-26 13:24:05'),(4,1,'康复训练服务',0,1,'2023-01-10 16:43:06',1,'2023-04-26 13:24:16'),(5,1,'休闲娱乐服务',0,1,'2023-01-10 16:43:09',1,'2023-04-26 13:24:29'),(6,1,'心理疏导服务',0,1,'2023-01-10 16:43:17',1,'2023-04-26 13:24:38'),(7,1,'社交与交流服务',0,1,'2023-01-10 16:43:19',1,'2023-04-26 13:24:48'),(8,1,'服务类型8',1,1,'2023-01-10 16:43:22',1,'2023-04-03 15:56:06'),(9,1,'服务类型9',1,1,'2023-01-10 16:43:25',1,'2023-04-03 15:56:01'),(10,1,'服务类型10',1,1,'2023-01-10 16:43:28',1,'2023-01-10 17:04:24'),(11,1,'服务类型11',1,1,'2023-01-10 16:43:55',1,'2023-04-03 15:53:36'),(12,1,'管理服务',0,1,'2023-04-03 16:34:33',1,'2023-04-26 13:24:57'),(13,1,'安全服务',0,1,'2023-04-14 17:49:45',1,'2023-04-26 13:25:05');
/*!40000 ALTER TABLE `service_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `set_dishes`
--

DROP TABLE IF EXISTS `set_dishes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `set_dishes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `set_id` bigint unsigned NOT NULL COMMENT '餐饮套餐id',
  `dishes_id` bigint unsigned NOT NULL COMMENT '菜品食物id',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `set_dishes`
--

LOCK TABLES `set_dishes` WRITE;
/*!40000 ALTER TABLE `set_dishes` DISABLE KEYS */;
INSERT INTO `set_dishes` VALUES (3,1,2,2,1,'2023-01-13 16:23:12',1,'2023-01-13 16:23:12'),(5,1,1,1,1,'2023-01-13 16:24:22',1,'2023-01-13 16:24:22'),(6,1,1,3,1,'2023-01-13 16:24:22',1,'2023-01-13 16:24:22'),(9,1,3,3,1,'2023-04-04 11:57:24',1,'2023-04-04 11:57:24'),(10,1,3,2,1,'2023-04-04 11:57:24',1,'2023-04-04 11:57:24'),(11,1,3,1,1,'2023-04-04 11:57:24',1,'2023-04-04 11:57:24'),(17,1,4,2,1,'2023-04-04 14:57:28',1,'2023-04-04 14:57:28'),(18,1,4,1,1,'2023-04-04 14:57:28',1,'2023-04-04 14:57:28'),(19,1,4,3,1,'2023-04-04 14:57:28',1,'2023-04-04 14:57:28'),(20,1,5,5,1,'2023-04-14 18:12:09',1,'2023-04-14 18:12:09');
/*!40000 ALTER TABLE `set_dishes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `source`
--

DROP TABLE IF EXISTS `source`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `source` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '来源渠道名称',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `source`
--

LOCK TABLES `source` WRITE;
/*!40000 ALTER TABLE `source` DISABLE KEYS */;
INSERT INTO `source` VALUES (1,1,'传单海报',0,1,'2023-01-03 00:18:44',1,'2023-01-04 00:08:57'),(3,1,'老客户',0,1,'2023-01-04 00:07:19',1,'2023-01-04 00:07:19'),(4,1,'电视广播',0,1,'2023-01-04 00:07:46',1,'2023-01-04 00:07:46'),(5,1,'测试',1,1,'2023-04-04 19:54:50',1,'2023-04-04 19:55:24'),(6,1,'胡图图',0,1,'2023-04-14 18:20:59',1,'2023-04-14 18:20:59');
/*!40000 ALTER TABLE `source` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `staff`
--

DROP TABLE IF EXISTS `staff`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `staff` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `role_id` bigint unsigned NOT NULL COMMENT '角色id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '姓名',
  `id_num` varchar(18) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '身份证号',
  `age` int NOT NULL COMMENT '年龄',
  `sex` varchar(2) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '性别(男/女)',
  `phone` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '电话',
  `email` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '邮箱',
  `pass` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '头像',
  `address` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '地址',
  `status` tinyint NOT NULL COMMENT '状态：0-离职，1-在职',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `staff`
--

LOCK TABLES `staff` WRITE;
/*!40000 ALTER TABLE `staff` DISABLE KEYS */;
INSERT INTO `staff` VALUES (1,1,1,'超级管理员','1',29,'男','13547584400','756932656@qq.com','7217ac017b4fb2352ec9e65576c5c0b1','http://127.0.0.1:9001/upload/img/20230405/1643600038046306304_logo.png','1',0,0,'2022-12-31 12:34:43',1,'2023-04-26 15:48:02'),(2,1,2,'小管','230182198012251659',31,'女','13547584401','2175728501@qq.com','7217ac017b4fb2352ec9e65576c5c0b1','http://127.0.0.1:9001/upload/img/20240420/1781361739637297152_1.png','黑龙江哈尔滨',0,0,'2022-12-31 12:34:43',1,'2024-04-20 00:38:31'),(3,1,4,'小妹','230182198012251659',30,'女','13547584402','2927803979@qq.com','7217ac017b4fb2352ec9e65576c5c0b1','http://127.0.0.1:9001/upload/img/20240420/1781361825431785472_2.png','黑龙江哈尔滨',0,0,'2022-12-31 12:34:43',1,'2024-04-20 00:39:35'),(4,1,5,'小轩','230182198012251659',30,'女','18428167423','2710123337@qq.com','7217ac017b4fb2352ec9e65576c5c0b1','http://127.0.0.1:9001/upload/img/20240420/1781362049885769728_3.png','黑龙江哈尔滨',0,0,'2022-12-31 12:34:43',1,'2024-04-20 00:40:04'),(5,1,3,'张三','230182198012251659',34,'男','13547584403','emperorwen@qq.com','acbc6a79c54b20e9110ff7dc7a7db261','url','四川南充',1,1,'2023-01-09 11:02:55',1,'2023-01-09 11:02:55'),(6,1,7,'张三','230182198012251659',34,'男','13547584403','6745756876@qq.com','acbc6a79c54b20e9110ff7dc7a7db261','http://127.0.0.1:9001/upload/img/20240420/1781362261823950848_3.png','湖北省武汉市',0,1,'2023-01-09 11:08:45',1,'2024-04-20 00:40:36'),(7,1,8,'王五','230182198012251659',34,'男','13547584404','6745756877@qq.com','acbc6a79c54b20e9110ff7dc7a7db261','http://127.0.0.1:9001/upload/img/20240420/1781362175953965056_4.png','湖北省武汉市',0,1,'2023-01-09 11:08:45',1,'2024-04-20 00:40:22'),(8,1,4,'曹氏','230182198012251659',23,'男','13547584407','376598236@qq.com','c16a1f4716e9533b1b0e03d2315c7cf3','http://127.0.0.1:9001/upload/img/20230402/1642352470234554368_屏幕截图(2).png','湖北省武汉市',1,1,'2023-04-02 10:25:19',1,'2023-04-02 10:25:56');
/*!40000 ALTER TABLE `staff` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tenant`
--

DROP TABLE IF EXISTS `tenant`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tenant` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '企业名称',
  `logo` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '企业logo',
  `contact_name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '联系人姓名',
  `contact_phone` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '联系电话',
  `plan` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '套餐',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0试用中 1正式 2锁定',
  `expire_time` datetime DEFAULT NULL COMMENT '套餐到期时间',
  `create_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='租户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tenant`
--

LOCK TABLES `tenant` WRITE;
/*!40000 ALTER TABLE `tenant` DISABLE KEYS */;
INSERT INTO `tenant` VALUES (1,'平台租户','','平台','','',1,NULL,0,'2023-01-01 00:00:00',0,'2023-01-01 00:00:00',0);
/*!40000 ALTER TABLE `tenant` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `union_id` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '微信UnionID（全局唯一）',
  `openid` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '微信OpenID（兜底匹配）',
  `phone` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号（账号密码登录）',
  `pass` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码(md5)',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_union_id` (`union_id`),
  UNIQUE KEY `uk_phone` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='全局用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'','','13547584400','7217ac017b4fb2352ec9e65576c5c0b1','超级管理员','http://127.0.0.1:9001/upload/img/20230405/1643600038046306304_logo.png','2022-12-31 12:34:43','2023-04-26 15:48:02',0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `visit`
--

DROP TABLE IF EXISTS `visit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `visit` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '来访者姓名',
  `id_card` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '身份证号（加密）',
  `phone` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '来访者电话',
  `relation` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '与老人关系',
  `visit_time` datetime NOT NULL COMMENT '来访时间',
  `leave_time` datetime DEFAULT NULL COMMENT '离开时间',
  `visit_num` int NOT NULL COMMENT '来访者人数',
  `status` tinyint NOT NULL COMMENT '来访状态',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '删除状态',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_elder_id` (`elder_id`) /*!80000 INVISIBLE */,
  KEY `idx_visit_time` (`visit_time`),
  CONSTRAINT `fk_visitors_elder` FOREIGN KEY (`elder_id`) REFERENCES `elder` (`id`),
  CONSTRAINT `fk_visitors_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='访客记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `visit`
--

LOCK TABLES `visit` WRITE;
/*!40000 ALTER TABLE `visit` DISABLE KEYS */;
INSERT INTO `visit` VALUES (1,1,1,'王伍',NULL,'13647563648','侄子','2022-12-13 00:00:00','2022-12-14 00:00:00',4,1,0,1,'2023-02-04 15:42:16',1,'2023-02-04 15:54:38'),(2,1,2,'王氏0',NULL,'13547584491','侄子','2023-03-31 00:00:00','2023-04-01 00:00:00',11,1,1,1,'2023-04-01 14:19:44',1,'2023-04-01 14:33:13'),(3,1,1,'测试',NULL,'13547563980','父子','2023-04-07 00:00:00',NULL,8,0,0,1,'2023-04-07 17:09:02',1,'2023-04-07 17:09:02'),(4,1,1,'胡图图',NULL,'19120362147','无关系','2023-04-10 00:00:00',NULL,6,0,0,1,'2023-04-14 17:19:26',1,'2023-04-14 17:19:26');
/*!40000 ALTER TABLE `visit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `visit_plan`
--

DROP TABLE IF EXISTS `visit_plan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `visit_plan` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `elder_id` bigint unsigned NOT NULL COMMENT '老人id',
  `title` varchar(25) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '回访计划标题',
  `plan_date` datetime NOT NULL COMMENT '计划回访时间',
  `content` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '回访计划内容',
  `complete_date` datetime DEFAULT NULL COMMENT '计划完成时间',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `visit_plan`
--

LOCK TABLES `visit_plan` WRITE;
/*!40000 ALTER TABLE `visit_plan` DISABLE KEYS */;
INSERT INTO `visit_plan` VALUES (1,1,1,'测试','2023-01-02 15:40:41',NULL,NULL,0,1,'2023-01-02 15:40:54',1,'2023-01-05 22:35:05'),(2,1,1,'测试标题','2022-12-13 00:00:00',NULL,NULL,0,1,'2023-01-05 22:29:36',1,'2023-01-05 22:29:36');
/*!40000 ALTER TABLE `visit_plan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `warehouse`
--

DROP TABLE IF EXISTS `warehouse`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `warehouse` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `staff_id` bigint unsigned NOT NULL COMMENT '仓库管理员id',
  `name` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '仓库名称',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `warehouse`
--

LOCK TABLES `warehouse` WRITE;
/*!40000 ALTER TABLE `warehouse` DISABLE KEYS */;
INSERT INTO `warehouse` VALUES (1,1,6,'仓库1',0,1,'2023-01-15 22:31:36',1,'2023-01-15 22:32:59'),(2,1,6,'仓库2',0,1,'2023-01-15 22:31:42',1,'2023-01-15 22:33:51');
/*!40000 ALTER TABLE `warehouse` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `warehouse_material`
--

DROP TABLE IF EXISTS `warehouse_material`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `warehouse_material` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `warehouse_record_id` bigint unsigned NOT NULL COMMENT '入库登记id',
  `material_id` bigint unsigned NOT NULL COMMENT '物资id',
  `warehouse_num` int NOT NULL COMMENT '入库数量',
  `inventory` int NOT NULL COMMENT '库存量',
  `product_date` datetime NOT NULL COMMENT '生产日期',
  `expire_date` datetime NOT NULL COMMENT '有效期',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `warehouse_material`
--

LOCK TABLES `warehouse_material` WRITE;
/*!40000 ALTER TABLE `warehouse_material` DISABLE KEYS */;
INSERT INTO `warehouse_material` VALUES (1,1,1,1,2,0,'2022-12-12 00:00:00','2022-12-13 00:00:00',1,'2023-01-30 11:06:01',1,'2023-01-31 21:32:30'),(2,1,1,2,1,0,'2022-12-12 00:00:00','2022-12-13 00:00:00',1,'2023-01-30 11:06:01',1,'2023-01-31 21:32:30'),(3,1,2,1,2,2,'2022-12-12 00:00:00','2022-12-13 00:00:00',1,'2023-01-30 11:08:22',1,'2023-01-31 22:18:44'),(4,1,2,2,1,1,'2022-12-12 00:00:00','2022-12-13 00:00:00',1,'2023-01-30 11:08:22',1,'2023-01-31 21:43:15');
/*!40000 ALTER TABLE `warehouse_material` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `warehouse_record`
--

DROP TABLE IF EXISTS `warehouse_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `warehouse_record` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tenant_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '租户id',
  `warehouse_id` bigint unsigned NOT NULL COMMENT '仓库id',
  `staff_id` bigint unsigned NOT NULL COMMENT '经办人id',
  `source` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '物资来源',
  `warehouse_date` datetime NOT NULL COMMENT '入库时间',
  `status` tinyint NOT NULL COMMENT '入库状态',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '管理状态：-1=删除，0=禁用，1=可用',
  `create_id` bigint unsigned NOT NULL COMMENT '创建人id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_id` bigint unsigned NOT NULL COMMENT '修改人id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `warehouse_record`
--

LOCK TABLES `warehouse_record` WRITE;
/*!40000 ALTER TABLE `warehouse_record` DISABLE KEYS */;
INSERT INTO `warehouse_record` VALUES (1,1,1,6,'购买','2022-12-13 00:00:00',2,0,1,'2023-01-30 11:06:01',1,'2023-01-30 11:33:27'),(2,1,2,6,'购买','2022-12-13 00:00:00',2,0,1,'2023-01-30 11:08:22',1,'2023-01-30 11:08:22');
/*!40000 ALTER TABLE `warehouse_record` ENABLE KEYS */;
UNLOCK TABLES;
SET @@SESSION.SQL_LOG_BIN = @MYSQLDUMP_TEMP_LOG_BIN;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-08-29 11:36:55
