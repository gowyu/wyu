/*
 Navicat Premium Data Transfer

 Source Server         : Mac
 Source Server Type    : MySQL
 Source Server Version : 50725
 Source Host           : localhost:3306
 Source Schema         : wYu

 Target Server Type    : MySQL
 Target Server Version : 50725
 File Encoding         : 65001

 Date: 30/12/2019 15:30:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for test_test
-- ----------------------------
DROP TABLE IF EXISTS `test_test`;
CREATE TABLE `test_test` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `test_id` int(11) DEFAULT NULL,
  `content` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `test_id` (`test_id`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for tests
-- ----------------------------
DROP TABLE IF EXISTS `tests`;
CREATE TABLE `tests` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'AutoIdentity',
  `name` varchar(50) COLLATE utf8_bin DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of tests
-- ----------------------------
BEGIN;
INSERT INTO `tests` VALUES (1, 'name_test_1');
INSERT INTO `tests` VALUES (2, 'name_test_2');
INSERT INTO `tests` VALUES (3, 'name_test_3');
INSERT INTO `tests` VALUES (4, 'name_test_4');
INSERT INTO `tests` VALUES (5, 'name_test_5');
INSERT INTO `tests` VALUES (6, 'name_test_6');
INSERT INTO `tests` VALUES (7, 'name_test_7');
INSERT INTO `tests` VALUES (8, 'update success');
INSERT INTO `tests` VALUES (10, 'Add_Test_Two');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
