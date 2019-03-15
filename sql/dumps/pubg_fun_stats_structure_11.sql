/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 100313
Source Host           : localhost:3307
Source Database       : pubg_fun_stats

Target Server Type    : MYSQL
Target Server Version : 100313
File Encoding         : 65001

Date: 2019-03-11 14:51:16
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for foo
-- ----------------------------
DROP TABLE IF EXISTS `foo`;
CREATE TABLE `foo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `item` int(255) DEFAULT NULL,
  `itemval` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=921459 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for kills
-- ----------------------------
DROP TABLE IF EXISTS `kills`;
CREATE TABLE `kills` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `attack_id` bigint(20) NOT NULL,
  `killer` varchar(255) NOT NULL,
  `victim` varchar(255) NOT NULL,
  `distance` double NOT NULL,
  `damage_causer_name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `attack_id_index` (`attack_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=46041 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for matches
-- ----------------------------
DROP TABLE IF EXISTS `matches`;
CREATE TABLE `matches` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `match_id` varchar(255) NOT NULL,
  `shard_id` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE current_timestamp(),
  `duration` int(11) NOT NULL DEFAULT current_timestamp(6),
  `game_mode` varchar(255) NOT NULL,
  `map_name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `match_id_index` (`match_id`)
) ENGINE=InnoDB AUTO_INCREMENT=658 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for participants
-- ----------------------------
DROP TABLE IF EXISTS `participants`;
CREATE TABLE `participants` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `match_id` varchar(128) NOT NULL,
  `participant_id` varchar(128) NOT NULL,
  `roster_id` varchar(128) NOT NULL,
  `dbnos` int(11) NOT NULL,
  `assists` int(11) NOT NULL,
  `boosts` int(11) NOT NULL,
  `damage_dealt` double(11,0) NOT NULL,
  `death_type` varchar(20) NOT NULL,
  `headshot_kills` int(11) NOT NULL,
  `heals` int(11) NOT NULL,
  `kill_place` int(11) NOT NULL,
  `kill_streaks` int(11) NOT NULL,
  `kills` int(11) NOT NULL,
  `longest_kill` int(11) NOT NULL,
  `name` varchar(32) NOT NULL,
  `player_id` varchar(64) NOT NULL,
  `revives` int(11) NOT NULL,
  `ride_distance` double NOT NULL,
  `road_kills` int(11) NOT NULL,
  `swim_distance` double NOT NULL,
  `team_kills` int(11) NOT NULL,
  `time_survived` double NOT NULL,
  `vehicle_destroys` int(11) NOT NULL,
  `walk_distance` double NOT NULL,
  `win_place` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `participant_id_index` (`participant_id`)
) ENGINE=InnoDB AUTO_INCREMENT=46392 DEFAULT CHARSET=utf8;

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
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for rosters
-- ----------------------------
DROP TABLE IF EXISTS `rosters`;
CREATE TABLE `rosters` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `match_id` varchar(255) NOT NULL,
  `roster_id` varchar(255) NOT NULL,
  `shard_id` varchar(255) NOT NULL,
  `rank` int(11) NOT NULL,
  `team_id` int(11) NOT NULL,
  `won` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `roster_id_index` (`roster_id`)
) ENGINE=InnoDB AUTO_INCREMENT=30575 DEFAULT CHARSET=utf8;
SET FOREIGN_KEY_CHECKS=1;
