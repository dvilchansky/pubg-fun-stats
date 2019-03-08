/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 100313
Source Host           : localhost:3307
Source Database       : pubg_fun_stats

Target Server Type    : MYSQL
Target Server Version : 100313
File Encoding         : 65001

Date: 2019-03-08 22:31:39
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for matches
-- ----------------------------
DROP TABLE IF EXISTS `matches`;
CREATE TABLE `matches` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `match_id` varchar(255) NOT NULL,
  `match_game_id` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `match_id_index` (`match_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for players
-- ----------------------------
DROP TABLE IF EXISTS `players`;
CREATE TABLE `players` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `player_id` varchar(255) NOT NULL,
  `player_name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `player_id_index` (`player_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
SET FOREIGN_KEY_CHECKS=1;
