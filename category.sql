/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3306
 Source Schema         : bingo

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 13/12/2019 17:00:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bg_category
-- ----------------------------
DROP TABLE IF EXISTS `bg_category`;
CREATE TABLE `bg_category` (
  `id` mediumint(6) unsigned NOT NULL AUTO_INCREMENT,
  `cate_name` varchar(128) NOT NULL COMMENT '名称',
  `slug` varchar(128) NOT NULL COMMENT '缩略名',
  `counts` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文章数量',
  `parent_id` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '上级id',
  `intro` varchar(255) DEFAULT NULL COMMENT '介绍',
  `list_order` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(10) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `cover` varchar(255) DEFAULT '' COMMENT '封面',
  `template` varchar(64) NOT NULL COMMENT '模板',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`),
  UNIQUE KEY `cate_name` (`cate_name`),
  KEY `list_order` (`list_order`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='分类';

-- ----------------------------
-- Records of bg_category
-- ----------------------------
BEGIN;
INSERT INTO `bg_category` VALUES (1, 'teste', 'test', 0, 0, '', 0, 0, 1576227503, '', '', 1);
INSERT INTO `bg_category` VALUES (2, 'Flutter', 'flutter', 0, 0, '', 0, 1573004970, 0, '', '', 1);
INSERT INTO `bg_category` VALUES (3, '我的日记', 'my-diary', 0, 0, '', 0, 1573008562, 0, '', '', 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
