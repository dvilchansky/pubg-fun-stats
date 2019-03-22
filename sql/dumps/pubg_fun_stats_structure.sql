/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 100313
Source Host           : localhost:3307
Source Database       : pubg_fun_stats

Target Server Type    : MYSQL
Target Server Version : 100313
File Encoding         : 65001

Date: 2019-03-22 11:29:23
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for kill
-- ----------------------------
DROP TABLE IF EXISTS `kill`;
CREATE TABLE `kill` (
  `attack_id` bigint(20) NOT NULL,
  `killer` varchar(255) NOT NULL,
  `victim` varchar(255) NOT NULL,
  `distance` double NOT NULL,
  `damage_causer_name` varchar(255) NOT NULL,
  `match_id` varchar(255) NOT NULL,
  PRIMARY KEY (`attack_id`),
  UNIQUE KEY `attack_id_index` (`attack_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for match
-- ----------------------------
DROP TABLE IF EXISTS `match`;
CREATE TABLE `match` (
  `id` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE current_timestamp(),
  `duration` int(11) NOT NULL DEFAULT current_timestamp(6),
  `game_mode` varchar(255) NOT NULL,
  `map_name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `match_id_index` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for participant
-- ----------------------------
DROP TABLE IF EXISTS `participant`;
CREATE TABLE `participant` (
  `id` varchar(128) NOT NULL,
  `match_id` varchar(128) NOT NULL,
  `roster_id` varchar(128) NOT NULL,
  `damage_dealt` double(11,0) NOT NULL,
  `death_type` varchar(20) NOT NULL,
  `headshot_kills` int(11) NOT NULL,
  `kills` int(11) NOT NULL,
  `longest_kill` double(11,0) NOT NULL,
  `name` varchar(32) NOT NULL,
  `player_id` varchar(64) NOT NULL,
  `time_survived` double NOT NULL,
  `win_place` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `participant_id_index` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for player
-- ----------------------------
DROP TABLE IF EXISTS `player`;
CREATE TABLE `player` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `player_id` varchar(255) NOT NULL,
  `player_name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `player_id_index` (`player_id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for roster
-- ----------------------------
DROP TABLE IF EXISTS `roster`;
CREATE TABLE `roster` (
  `roster_id` varchar(255) NOT NULL,
  `match_id` varchar(255) NOT NULL,
  `shard_id` varchar(255) NOT NULL,
  `rank` int(11) NOT NULL,
  `team_id` int(11) NOT NULL,
  `won` varchar(255) NOT NULL,
  PRIMARY KEY (`roster_id`),
  UNIQUE KEY `roster_id_index` (`roster_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
SET FOREIGN_KEY_CHECKS=1;
