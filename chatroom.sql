/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : 127.0.0.1:3306
 Source Schema         : chatroom

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 23/04/2021 16:25:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for c_users
-- ----------------------------
DROP TABLE IF EXISTS `c_users`;
CREATE TABLE `c_users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT '用户账号',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT '用户密码',
  `mail` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT '联系邮箱',
  `login_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT '登录IP',
  `login_num` bigint unsigned DEFAULT '0' COMMENT '登录次数',
  `desc` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT '备注说明',
  `status` tinyint unsigned DEFAULT '1' COMMENT '状态(0禁用,1启用)',
  `head` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `is_save_head` tinyint DEFAULT '0' COMMENT '是否设置头像',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '最后登录时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE COMMENT '用户名唯一',
  KEY `index_system_user_username` (`username`) USING BTREE,
  KEY `index_system_user_status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10046 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='系统-用户';

-- ----------------------------
-- Records of c_users
-- ----------------------------
BEGIN;
INSERT INTO `c_users` VALUES (10044, '一叶之秋', '4297f44b13955235245b2497399d7a93', '18337180580@163.com', '', 0, '', 0, '{\"skin\":\"skin-3\",\"hat\":\"no-hat\",\"eyebrow\":\"no-eyebrows\",\"eye\":\"default\",\"mouth\":\"default\",\"faceExtras\":\"sweat\",\"item\":\"default\"}', 0, '2021-04-19 17:14:42', '2021-04-22 18:21:07');
INSERT INTO `c_users` VALUES (10045, '城南花已开', '4297f44b13955235245b2497399d7a93', '1921197829@qq.com', '', 0, '', 0, '{\"hat\":\"no-hat\",\"eyebrow\":\"no-eyebrows\",\"eye\":\"default\",\"mouth\":\"default\",\"faceExtras\":\"sweat\",\"item\":\"default\"}', 0, '2021-04-20 18:35:55', '2021-04-20 18:35:55');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
