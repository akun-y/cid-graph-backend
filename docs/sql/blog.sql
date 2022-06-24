/*
 Navicat MySQL Data Transfer
 
 Source Database       : [cid_graph]
 
 Target Server Type    : MYSQL
 Target Server Version : 50639
 File Encoding         : 65001
 
 Date: 2022-06-24 16:52:35
 */
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for graph_data
-- ----------------------------
DROP TABLE IF EXISTS `graph_data`;
CREATE TABLE `graph_data` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `name` varchar(100) DEFAULT '' COMMENT 'cid名称',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `ipfs_cid` varchar(255) COMMENT 'ipfs cid',
  `size` int(10) unsigned DEFAULT '0' COMMENT '文件大小',
  `length` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `type` varchar(32) DEFAULT '' COMMENT '文件类型',
  `version` int(10) unsigned DEFAULT '0' COMMENT '版本',
  `publish_time` int(10) unsigned DEFAULT '0' COMMENT '发布时间',
  `copyright` varchar(255) DEFAULT '' COMMENT ' 版权',
  `metas` varchar(4096) DEFAULT '' COMMENT ' 元数据',
  `graph_author_id` int(10) unsigned DEFAULT '0' COMMENT 'graph作者',
  `owner_id` int(10) unsigned DEFAULT '0' COMMENT '所有者',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT = 'CID管理';
-- ----------------------------
-- Table structure for graph_auth
-- ----------------------------
DROP TABLE IF EXISTS `graph_auth`;
CREATE TABLE `graph_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(256) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8;
INSERT INTO `graph_auth` (`id`, `username`, `password`)
VALUES ('1', 'cid', 'cid1234');
-- ----------------------------
-- Table structure for graph_tag
-- ----------------------------
DROP TABLE IF EXISTS `graph_tag`;
CREATE TABLE `graph_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT = '标签管理';
-- ----------------------------
-- Table structure for graph_user
-- ----------------------------
DROP TABLE IF EXISTS `graph_user`;
CREATE TABLE `graph_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  `website` varchar(128) DEFAULT '' COMMENT 'website',
  `github` varchar(128) DEFAULT '' COMMENT 'github',
  `wallet` varchar(128) DEFAULT '' COMMENT 'wallet address',
  `linkin` varchar(128) DEFAULT '' COMMENT 'linkin',
  `email` varchar(128) DEFAULT '' COMMENT 'email',
  `cids` varchar(4096) DEFAULT '' COMMENT 'cids,ex:[cid1,cid2]',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT = '用户管理';