/*
Navicat MySQL Data Transfer

Source Server         : gin-vue-admin
Source Server Version : 50650
Source Host           : 159.75.37.58:3306
Source Database       : gin-vue-admin

Target Server Type    : MYSQL
Target Server Version : 50650
File Encoding         : 65001

Date: 2023-02-20 10:27:42
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  `v6` varchar(25) DEFAULT NULL,
  `v7` varchar(25) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=169 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('1', 'p', '888', '/base/login', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2', 'p', '888', '/user/admin_register', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('3', 'p', '888', '/api/createApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('4', 'p', '888', '/api/getApiList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('5', 'p', '888', '/api/getApiById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('6', 'p', '888', '/api/deleteApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('7', 'p', '888', '/api/updateApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('8', 'p', '888', '/api/getAllApis', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('9', 'p', '888', '/api/deleteApisByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('10', 'p', '888', '/authority/copyAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('11', 'p', '888', '/authority/updateAuthority', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('12', 'p', '888', '/authority/createAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('13', 'p', '888', '/authority/deleteAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('14', 'p', '888', '/authority/getAuthorityList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('15', 'p', '888', '/authority/setDataAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('16', 'p', '888', '/menu/getMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('17', 'p', '888', '/menu/getMenuList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('18', 'p', '888', '/menu/addBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('19', 'p', '888', '/menu/getBaseMenuTree', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('20', 'p', '888', '/menu/addMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('21', 'p', '888', '/menu/getMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('22', 'p', '888', '/menu/deleteBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('23', 'p', '888', '/menu/updateBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('24', 'p', '888', '/menu/getBaseMenuById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('25', 'p', '888', '/user/getUserInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('26', 'p', '888', '/user/setUserInfo', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('27', 'p', '888', '/user/setSelfInfo', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('28', 'p', '888', '/user/getUserList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('29', 'p', '888', '/user/deleteUser', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('30', 'p', '888', '/user/changePassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('31', 'p', '888', '/user/setUserAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('32', 'p', '888', '/user/setUserAuthorities', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('33', 'p', '888', '/user/resetPassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('34', 'p', '888', '/fileUploadAndDownload/findFile', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('35', 'p', '888', '/fileUploadAndDownload/breakpointContinueFinish', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('36', 'p', '888', '/fileUploadAndDownload/breakpointContinue', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('37', 'p', '888', '/fileUploadAndDownload/removeChunk', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('38', 'p', '888', '/fileUploadAndDownload/upload', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('39', 'p', '888', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('40', 'p', '888', '/fileUploadAndDownload/editFileName', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('41', 'p', '888', '/fileUploadAndDownload/getFileList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('42', 'p', '888', '/casbin/updateCasbin', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('43', 'p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('44', 'p', '888', '/jwt/jsonInBlacklist', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('45', 'p', '888', '/system/getSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('46', 'p', '888', '/system/setSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('47', 'p', '888', '/system/getServerInfo', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('48', 'p', '888', '/customer/customer', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('49', 'p', '888', '/customer/customer', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('50', 'p', '888', '/customer/customer', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('51', 'p', '888', '/customer/customer', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('52', 'p', '888', '/customer/customerList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('53', 'p', '888', '/autoCode/getDB', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('54', 'p', '888', '/autoCode/getMeta', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('55', 'p', '888', '/autoCode/preview', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('56', 'p', '888', '/autoCode/getTables', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('57', 'p', '888', '/autoCode/getColumn', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('58', 'p', '888', '/autoCode/rollback', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('59', 'p', '888', '/autoCode/createTemp', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('60', 'p', '888', '/autoCode/delSysHistory', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('61', 'p', '888', '/autoCode/getSysHistory', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('62', 'p', '888', '/autoCode/createPackage', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('63', 'p', '888', '/autoCode/getPackage', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('64', 'p', '888', '/autoCode/delPackage', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('65', 'p', '888', '/autoCode/createPlug', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('66', 'p', '888', '/autoCode/installPlugin', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('67', 'p', '888', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('68', 'p', '888', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('69', 'p', '888', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('70', 'p', '888', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('71', 'p', '888', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('72', 'p', '888', '/sysDictionary/findSysDictionary', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('73', 'p', '888', '/sysDictionary/updateSysDictionary', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('74', 'p', '888', '/sysDictionary/getSysDictionaryList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('75', 'p', '888', '/sysDictionary/createSysDictionary', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('76', 'p', '888', '/sysDictionary/deleteSysDictionary', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('77', 'p', '888', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('78', 'p', '888', '/sysOperationRecord/updateSysOperationRecord', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('79', 'p', '888', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('80', 'p', '888', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('81', 'p', '888', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('82', 'p', '888', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('83', 'p', '888', '/email/emailTest', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('84', 'p', '888', '/simpleUploader/upload', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('85', 'p', '888', '/simpleUploader/checkFileMd5', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('86', 'p', '888', '/simpleUploader/mergeFileMd5', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('87', 'p', '888', '/authorityBtn/setAuthorityBtn', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('88', 'p', '888', '/authorityBtn/getAuthorityBtn', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('89', 'p', '888', '/authorityBtn/canRemoveAuthorityBtn', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('90', 'p', '8881', '/base/login', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('91', 'p', '8881', '/user/admin_register', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('92', 'p', '8881', '/api/createApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('93', 'p', '8881', '/api/getApiList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('94', 'p', '8881', '/api/getApiById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('95', 'p', '8881', '/api/deleteApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('96', 'p', '8881', '/api/updateApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('97', 'p', '8881', '/api/getAllApis', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('98', 'p', '8881', '/authority/createAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('99', 'p', '8881', '/authority/deleteAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('100', 'p', '8881', '/authority/getAuthorityList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('101', 'p', '8881', '/authority/setDataAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('102', 'p', '8881', '/menu/getMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('103', 'p', '8881', '/menu/getMenuList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('104', 'p', '8881', '/menu/addBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('105', 'p', '8881', '/menu/getBaseMenuTree', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('106', 'p', '8881', '/menu/addMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('107', 'p', '8881', '/menu/getMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('108', 'p', '8881', '/menu/deleteBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('109', 'p', '8881', '/menu/updateBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('110', 'p', '8881', '/menu/getBaseMenuById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('111', 'p', '8881', '/user/changePassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('112', 'p', '8881', '/user/getUserList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('113', 'p', '8881', '/user/setUserAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('114', 'p', '8881', '/fileUploadAndDownload/upload', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('115', 'p', '8881', '/fileUploadAndDownload/getFileList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('116', 'p', '8881', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('117', 'p', '8881', '/fileUploadAndDownload/editFileName', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('118', 'p', '8881', '/casbin/updateCasbin', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('119', 'p', '8881', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('120', 'p', '8881', '/jwt/jsonInBlacklist', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('121', 'p', '8881', '/system/getSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('122', 'p', '8881', '/system/setSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('123', 'p', '8881', '/customer/customer', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('124', 'p', '8881', '/customer/customer', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('125', 'p', '8881', '/customer/customer', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('126', 'p', '8881', '/customer/customer', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('127', 'p', '8881', '/customer/customerList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('128', 'p', '8881', '/user/getUserInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('129', 'p', '9528', '/base/login', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('130', 'p', '9528', '/user/admin_register', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('131', 'p', '9528', '/api/createApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('132', 'p', '9528', '/api/getApiList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('133', 'p', '9528', '/api/getApiById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('134', 'p', '9528', '/api/deleteApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('135', 'p', '9528', '/api/updateApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('136', 'p', '9528', '/api/getAllApis', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('137', 'p', '9528', '/authority/createAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('138', 'p', '9528', '/authority/deleteAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('139', 'p', '9528', '/authority/getAuthorityList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('140', 'p', '9528', '/authority/setDataAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('141', 'p', '9528', '/menu/getMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('142', 'p', '9528', '/menu/getMenuList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('143', 'p', '9528', '/menu/addBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('144', 'p', '9528', '/menu/getBaseMenuTree', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('145', 'p', '9528', '/menu/addMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('146', 'p', '9528', '/menu/getMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('147', 'p', '9528', '/menu/deleteBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('148', 'p', '9528', '/menu/updateBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('149', 'p', '9528', '/menu/getBaseMenuById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('150', 'p', '9528', '/user/changePassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('151', 'p', '9528', '/user/getUserList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('152', 'p', '9528', '/user/setUserAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('153', 'p', '9528', '/fileUploadAndDownload/upload', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('154', 'p', '9528', '/fileUploadAndDownload/getFileList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('155', 'p', '9528', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('156', 'p', '9528', '/fileUploadAndDownload/editFileName', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('157', 'p', '9528', '/casbin/updateCasbin', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('158', 'p', '9528', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('159', 'p', '9528', '/jwt/jsonInBlacklist', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('160', 'p', '9528', '/system/getSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('161', 'p', '9528', '/system/setSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('162', 'p', '9528', '/customer/customer', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('163', 'p', '9528', '/customer/customer', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('164', 'p', '9528', '/customer/customer', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('165', 'p', '9528', '/customer/customer', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('166', 'p', '9528', '/customer/customerList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('167', 'p', '9528', '/autoCode/createTemp', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('168', 'p', '9528', '/user/getUserInfo', 'GET', '', '', '', '', '');

-- ----------------------------
-- Table structure for exa_customers
-- ----------------------------
DROP TABLE IF EXISTS `exa_customers`;
CREATE TABLE `exa_customers` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_name` varchar(191) DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint(20) unsigned DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` bigint(20) unsigned DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_exa_customers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of exa_customers
-- ----------------------------

-- ----------------------------
-- Table structure for exa_file_chunks
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_chunks`;
CREATE TABLE `exa_file_chunks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `exa_file_id` bigint(20) unsigned DEFAULT NULL,
  `file_chunk_number` bigint(20) DEFAULT NULL,
  `file_chunk_path` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_chunks_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of exa_file_chunks
-- ----------------------------

-- ----------------------------
-- Table structure for exa_file_upload_and_downloads
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
CREATE TABLE `exa_file_upload_and_downloads` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of exa_file_upload_and_downloads
-- ----------------------------
INSERT INTO `exa_file_upload_and_downloads` VALUES ('1', '2023-02-20 10:25:35.533', '2023-02-20 10:25:35.533', null, '10.png', 'https://qmplusimg.henrongyi.top/gvalogo.png', 'png', '158787308910.png');
INSERT INTO `exa_file_upload_and_downloads` VALUES ('2', '2023-02-20 10:25:35.533', '2023-02-20 10:25:35.533', null, 'logo.png', 'https://qmplusimg.henrongyi.top/1576554439myAvatar.png', 'png', '1587973709logo.png');

-- ----------------------------
-- Table structure for exa_files
-- ----------------------------
DROP TABLE IF EXISTS `exa_files`;
CREATE TABLE `exa_files` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `file_name` varchar(191) DEFAULT NULL,
  `file_md5` varchar(191) DEFAULT NULL,
  `file_path` varchar(191) DEFAULT NULL,
  `chunk_total` bigint(20) DEFAULT NULL,
  `is_finish` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_files_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of exa_files
-- ----------------------------

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `jwt` text COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=90 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES ('1', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/base/login', '用户登录(必选)', 'base', 'POST');
INSERT INTO `sys_apis` VALUES ('2', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/jwt/jsonInBlacklist', 'jwt加入黑名单(退出，必选)', 'jwt', 'POST');
INSERT INTO `sys_apis` VALUES ('3', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/user/deleteUser', '删除用户', '系统用户', 'DELETE');
INSERT INTO `sys_apis` VALUES ('4', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/user/admin_register', '用户注册', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES ('5', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/user/getUserList', '获取用户列表', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES ('6', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/user/setUserInfo', '设置用户信息', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES ('7', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/user/setSelfInfo', '设置自身信息(必选)', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES ('8', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/user/getUserInfo', '获取自身信息(必选)', '系统用户', 'GET');
INSERT INTO `sys_apis` VALUES ('9', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/user/setUserAuthorities', '设置权限组', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES ('10', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/user/changePassword', '修改密码（建议选择)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES ('11', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/user/setUserAuthority', '修改用户角色(必选)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES ('12', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/user/resetPassword', '重置用户密码', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES ('13', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/api/createApi', '创建api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES ('14', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/api/deleteApi', '删除Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES ('15', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/api/updateApi', '更新Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES ('16', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/api/getApiList', '获取api列表', 'api', 'POST');
INSERT INTO `sys_apis` VALUES ('17', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/api/getAllApis', '获取所有api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES ('18', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/api/getApiById', '获取api详细信息', 'api', 'POST');
INSERT INTO `sys_apis` VALUES ('19', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/api/deleteApisByIds', '批量删除api', 'api', 'DELETE');
INSERT INTO `sys_apis` VALUES ('20', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/authority/copyAuthority', '拷贝角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES ('21', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/authority/createAuthority', '创建角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES ('22', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/authority/deleteAuthority', '删除角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES ('23', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/authority/updateAuthority', '更新角色信息', '角色', 'PUT');
INSERT INTO `sys_apis` VALUES ('24', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/authority/getAuthorityList', '获取角色列表', '角色', 'POST');
INSERT INTO `sys_apis` VALUES ('25', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/authority/setDataAuthority', '设置角色资源权限', '角色', 'POST');
INSERT INTO `sys_apis` VALUES ('26', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES ('27', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES ('28', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/menu/addBaseMenu', '新增菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES ('29', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/menu/getMenu', '获取菜单树(必选)', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES ('30', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/menu/deleteBaseMenu', '删除菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES ('31', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/menu/updateBaseMenu', '更新菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES ('32', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/menu/getBaseMenuById', '根据id获取菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES ('33', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/menu/getMenuList', '分页获取基础menu列表', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES ('34', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/menu/getBaseMenuTree', '获取用户动态路由', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES ('35', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/menu/getMenuAuthority', '获取指定角色menu', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES ('36', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/menu/addMenuAuthority', '增加menu和角色关联关系', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES ('37', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/fileUploadAndDownload/findFile', '寻找目标文件（秒传）', '分片上传', 'GET');
INSERT INTO `sys_apis` VALUES ('38', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/fileUploadAndDownload/breakpointContinue', '断点续传', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES ('39', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/fileUploadAndDownload/breakpointContinueFinish', '断点续传完成', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES ('40', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/fileUploadAndDownload/removeChunk', '上传完成移除文件', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES ('41', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/fileUploadAndDownload/upload', '文件上传示例', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES ('42', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/fileUploadAndDownload/deleteFile', '删除文件', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES ('43', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/fileUploadAndDownload/editFileName', '文件名或者备注编辑', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES ('44', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/fileUploadAndDownload/getFileList', '获取上传文件列表', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES ('45', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/system/getServerInfo', '获取服务器信息', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES ('46', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/system/getSystemConfig', '获取配置文件内容', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES ('47', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/system/setSystemConfig', '设置配置文件内容', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES ('48', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/customer/customer', '更新客户', '客户', 'PUT');
INSERT INTO `sys_apis` VALUES ('49', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/customer/customer', '创建客户', '客户', 'POST');
INSERT INTO `sys_apis` VALUES ('50', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/customer/customer', '删除客户', '客户', 'DELETE');
INSERT INTO `sys_apis` VALUES ('51', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/customer/customer', '获取单一客户', '客户', 'GET');
INSERT INTO `sys_apis` VALUES ('52', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/customer/customerList', '获取客户列表', '客户', 'GET');
INSERT INTO `sys_apis` VALUES ('53', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/autoCode/getDB', '获取所有数据库', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES ('54', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/autoCode/getTables', '获取数据库表', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES ('55', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/autoCode/createTemp', '自动化代码', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES ('56', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/autoCode/preview', '预览自动化代码', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES ('57', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/autoCode/getColumn', '获取所选table的所有字段', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES ('58', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/autoCode/createPlug', '自动创建插件包', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES ('59', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/autoCode/installPlugin', '安装插件', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES ('60', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/autoCode/createPackage', '生成包(package)', '包（pkg）生成器', 'POST');
INSERT INTO `sys_apis` VALUES ('61', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/autoCode/getPackage', '获取所有包(package)', '包（pkg）生成器', 'POST');
INSERT INTO `sys_apis` VALUES ('62', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/autoCode/delPackage', '删除包(package)', '包（pkg）生成器', 'POST');
INSERT INTO `sys_apis` VALUES ('63', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/autoCode/getMeta', '获取meta信息', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES ('64', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/autoCode/rollback', '回滚自动生成代码', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES ('65', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/autoCode/getSysHistory', '查询回滚记录', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES ('66', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/autoCode/delSysHistory', '删除回滚记录', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES ('67', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/sysDictionaryDetail/updateSysDictionaryDetail', '更新字典内容', '系统字典详情', 'PUT');
INSERT INTO `sys_apis` VALUES ('68', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/sysDictionaryDetail/createSysDictionaryDetail', '新增字典内容', '系统字典详情', 'POST');
INSERT INTO `sys_apis` VALUES ('69', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/sysDictionaryDetail/deleteSysDictionaryDetail', '删除字典内容', '系统字典详情', 'DELETE');
INSERT INTO `sys_apis` VALUES ('70', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/sysDictionaryDetail/findSysDictionaryDetail', '根据ID获取字典内容', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES ('71', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/sysDictionaryDetail/getSysDictionaryDetailList', '获取字典内容列表', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES ('72', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/sysDictionary/createSysDictionary', '新增字典', '系统字典', 'POST');
INSERT INTO `sys_apis` VALUES ('73', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/sysDictionary/deleteSysDictionary', '删除字典', '系统字典', 'DELETE');
INSERT INTO `sys_apis` VALUES ('74', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/sysDictionary/updateSysDictionary', '更新字典', '系统字典', 'PUT');
INSERT INTO `sys_apis` VALUES ('75', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/sysDictionary/findSysDictionary', '根据ID获取字典', '系统字典', 'GET');
INSERT INTO `sys_apis` VALUES ('76', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/sysDictionary/getSysDictionaryList', '获取字典列表', '系统字典', 'GET');
INSERT INTO `sys_apis` VALUES ('77', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', '操作记录', 'POST');
INSERT INTO `sys_apis` VALUES ('78', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', '操作记录', 'GET');
INSERT INTO `sys_apis` VALUES ('79', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', '操作记录', 'GET');
INSERT INTO `sys_apis` VALUES ('80', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', '操作记录', 'DELETE');
INSERT INTO `sys_apis` VALUES ('81', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', '操作记录', 'DELETE');
INSERT INTO `sys_apis` VALUES ('82', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/simpleUploader/upload', '插件版分片上传', '断点续传(插件版)', 'POST');
INSERT INTO `sys_apis` VALUES ('83', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/simpleUploader/checkFileMd5', '文件完整度验证', '断点续传(插件版)', 'GET');
INSERT INTO `sys_apis` VALUES ('84', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/simpleUploader/mergeFileMd5', '上传完成合并文件', '断点续传(插件版)', 'GET');
INSERT INTO `sys_apis` VALUES ('85', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/email/emailTest', '发送测试邮件', 'email', 'POST');
INSERT INTO `sys_apis` VALUES ('86', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/email/emailSend', '发送邮件示例', 'email', 'POST');
INSERT INTO `sys_apis` VALUES ('87', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/authorityBtn/setAuthorityBtn', '设置按钮权限', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES ('88', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/authorityBtn/getAuthorityBtn', '获取已有按钮权限', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES ('89', '2023-02-20 10:25:11.149', '2023-02-20 10:25:11.149', null, '/authorityBtn/canRemoveAuthorityBtn', '删除按钮', '按钮权限', 'POST');

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `authority_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint(20) unsigned DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`),
  UNIQUE KEY `authority_id` (`authority_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9529 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
INSERT INTO `sys_authorities` VALUES ('2023-02-20 10:25:13.476', '2023-02-20 10:25:29.131', null, '888', '普通用户', '0', 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2023-02-20 10:25:13.476', '2023-02-20 10:25:34.201', null, '8881', '普通用户子角色', '888', 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2023-02-20 10:25:13.476', '2023-02-20 10:25:30.957', null, '9528', '测试角色', '0', 'dashboard');

-- ----------------------------
-- Table structure for sys_authority_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_btns`;
CREATE TABLE `sys_authority_btns` (
  `authority_id` bigint(20) unsigned DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint(20) unsigned DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint(20) unsigned DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_authority_btns
-- ----------------------------

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus` (
  `sys_base_menu_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `sys_authority_authority_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
INSERT INTO `sys_authority_menus` VALUES ('1', '888');
INSERT INTO `sys_authority_menus` VALUES ('1', '8881');
INSERT INTO `sys_authority_menus` VALUES ('1', '9528');
INSERT INTO `sys_authority_menus` VALUES ('2', '888');
INSERT INTO `sys_authority_menus` VALUES ('2', '8881');
INSERT INTO `sys_authority_menus` VALUES ('2', '9528');
INSERT INTO `sys_authority_menus` VALUES ('3', '888');
INSERT INTO `sys_authority_menus` VALUES ('4', '888');
INSERT INTO `sys_authority_menus` VALUES ('4', '8881');
INSERT INTO `sys_authority_menus` VALUES ('5', '888');
INSERT INTO `sys_authority_menus` VALUES ('5', '8881');
INSERT INTO `sys_authority_menus` VALUES ('6', '888');
INSERT INTO `sys_authority_menus` VALUES ('6', '8881');
INSERT INTO `sys_authority_menus` VALUES ('7', '888');
INSERT INTO `sys_authority_menus` VALUES ('7', '8881');
INSERT INTO `sys_authority_menus` VALUES ('8', '888');
INSERT INTO `sys_authority_menus` VALUES ('8', '8881');
INSERT INTO `sys_authority_menus` VALUES ('8', '9528');
INSERT INTO `sys_authority_menus` VALUES ('9', '888');
INSERT INTO `sys_authority_menus` VALUES ('9', '8881');
INSERT INTO `sys_authority_menus` VALUES ('10', '888');
INSERT INTO `sys_authority_menus` VALUES ('10', '8881');
INSERT INTO `sys_authority_menus` VALUES ('11', '888');
INSERT INTO `sys_authority_menus` VALUES ('11', '8881');
INSERT INTO `sys_authority_menus` VALUES ('12', '888');
INSERT INTO `sys_authority_menus` VALUES ('13', '888');
INSERT INTO `sys_authority_menus` VALUES ('13', '8881');
INSERT INTO `sys_authority_menus` VALUES ('14', '888');
INSERT INTO `sys_authority_menus` VALUES ('14', '8881');
INSERT INTO `sys_authority_menus` VALUES ('15', '888');
INSERT INTO `sys_authority_menus` VALUES ('15', '8881');
INSERT INTO `sys_authority_menus` VALUES ('16', '888');
INSERT INTO `sys_authority_menus` VALUES ('16', '8881');
INSERT INTO `sys_authority_menus` VALUES ('17', '888');
INSERT INTO `sys_authority_menus` VALUES ('17', '8881');
INSERT INTO `sys_authority_menus` VALUES ('18', '888');
INSERT INTO `sys_authority_menus` VALUES ('19', '888');
INSERT INTO `sys_authority_menus` VALUES ('20', '888');
INSERT INTO `sys_authority_menus` VALUES ('21', '888');
INSERT INTO `sys_authority_menus` VALUES ('22', '888');
INSERT INTO `sys_authority_menus` VALUES ('23', '888');
INSERT INTO `sys_authority_menus` VALUES ('24', '888');
INSERT INTO `sys_authority_menus` VALUES ('25', '888');
INSERT INTO `sys_authority_menus` VALUES ('26', '888');
INSERT INTO `sys_authority_menus` VALUES ('27', '888');
INSERT INTO `sys_authority_menus` VALUES ('28', '888');
INSERT INTO `sys_authority_menus` VALUES ('29', '888');

-- ----------------------------
-- Table structure for sys_auto_code_histories
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_code_histories`;
CREATE TABLE `sys_auto_code_histories` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `package` varchar(191) DEFAULT NULL,
  `business_db` varchar(191) DEFAULT NULL,
  `table_name` varchar(191) DEFAULT NULL,
  `request_meta` text,
  `auto_code_path` text,
  `injection_meta` text,
  `struct_name` varchar(191) DEFAULT NULL,
  `struct_cn_name` varchar(191) DEFAULT NULL,
  `api_ids` varchar(191) DEFAULT NULL,
  `flag` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_code_histories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_auto_code_histories
-- ----------------------------

-- ----------------------------
-- Table structure for sys_auto_codes
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_codes`;
CREATE TABLE `sys_auto_codes` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `package_name` varchar(191) DEFAULT NULL COMMENT '包名',
  `label` varchar(191) DEFAULT NULL COMMENT '展示名',
  `desc` varchar(191) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_codes_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_auto_codes
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menu_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_btns`;
CREATE TABLE `sys_base_menu_btns` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) DEFAULT NULL,
  `sys_base_menu_id` bigint(20) unsigned DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_btns_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_base_menu_btns
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sys_base_menu_id` bigint(20) unsigned DEFAULT NULL,
  `type` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_base_menu_parameters
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` bigint(20) unsigned DEFAULT NULL,
  `parent_id` varchar(191) DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) DEFAULT NULL COMMENT '附加属性',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
INSERT INTO `sys_base_menus` VALUES ('1', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '0', 'dashboard', 'dashboard', '0', 'view/dashboard/index.vue', '1', '', '0', '0', '仪表盘', 'odometer', '0');
INSERT INTO `sys_base_menus` VALUES ('2', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '0', 'about', 'about', '0', 'view/about/index.vue', '9', '', '0', '0', '关于我们', 'info-filled', '0');
INSERT INTO `sys_base_menus` VALUES ('3', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '0', 'admin', 'superAdmin', '0', 'view/superAdmin/index.vue', '3', '', '0', '0', '超级管理员', 'user', '0');
INSERT INTO `sys_base_menus` VALUES ('4', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '3', 'authority', 'authority', '0', 'view/superAdmin/authority/authority.vue', '1', '', '0', '0', '角色管理', 'avatar', '0');
INSERT INTO `sys_base_menus` VALUES ('5', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '3', 'menu', 'menu', '0', 'view/superAdmin/menu/menu.vue', '2', '', '1', '0', '菜单管理', 'tickets', '0');
INSERT INTO `sys_base_menus` VALUES ('6', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '3', 'api', 'api', '0', 'view/superAdmin/api/api.vue', '3', '', '1', '0', 'api管理', 'platform', '0');
INSERT INTO `sys_base_menus` VALUES ('7', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '3', 'user', 'user', '0', 'view/superAdmin/user/user.vue', '4', '', '0', '0', '用户管理', 'coordinate', '0');
INSERT INTO `sys_base_menus` VALUES ('8', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '3', 'dictionary', 'dictionary', '0', 'view/superAdmin/dictionary/sysDictionary.vue', '5', '', '0', '0', '字典管理', 'notebook', '0');
INSERT INTO `sys_base_menus` VALUES ('9', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '3', 'dictionaryDetail/:id', 'dictionaryDetail', '1', 'view/superAdmin/dictionary/sysDictionaryDetail.vue', '1', 'dictionary', '0', '0', '字典详情-${id}', 'list', '0');
INSERT INTO `sys_base_menus` VALUES ('10', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '3', 'operation', 'operation', '0', 'view/superAdmin/operation/sysOperationRecord.vue', '6', '', '0', '0', '操作历史', 'pie-chart', '0');
INSERT INTO `sys_base_menus` VALUES ('11', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '0', 'person', 'person', '1', 'view/person/person.vue', '4', '', '0', '0', '个人信息', 'message', '0');
INSERT INTO `sys_base_menus` VALUES ('12', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '0', 'example', 'example', '0', 'view/example/index.vue', '7', '', '0', '0', '示例文件', 'management', '0');
INSERT INTO `sys_base_menus` VALUES ('13', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '12', 'upload', 'upload', '0', 'view/example/upload/upload.vue', '5', '', '0', '0', '媒体库（上传下载）', 'upload', '0');
INSERT INTO `sys_base_menus` VALUES ('14', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '12', 'breakpoint', 'breakpoint', '0', 'view/example/breakpoint/breakpoint.vue', '6', '', '0', '0', '断点续传', 'upload-filled', '0');
INSERT INTO `sys_base_menus` VALUES ('15', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '12', 'customer', 'customer', '0', 'view/example/customer/customer.vue', '7', '', '0', '0', '客户列表（资源示例）', 'avatar', '0');
INSERT INTO `sys_base_menus` VALUES ('16', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '0', 'systemTools', 'systemTools', '0', 'view/systemTools/index.vue', '5', '', '0', '0', '系统工具', 'tools', '0');
INSERT INTO `sys_base_menus` VALUES ('17', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '16', 'autoCode', 'autoCode', '0', 'view/systemTools/autoCode/index.vue', '1', '', '1', '0', '代码生成器', 'cpu', '0');
INSERT INTO `sys_base_menus` VALUES ('18', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '16', 'formCreate', 'formCreate', '0', 'view/systemTools/formCreate/index.vue', '2', '', '1', '0', '表单生成器', 'magic-stick', '0');
INSERT INTO `sys_base_menus` VALUES ('19', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '16', 'system', 'system', '0', 'view/systemTools/system/system.vue', '3', '', '0', '0', '系统配置', 'operation', '0');
INSERT INTO `sys_base_menus` VALUES ('20', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '16', 'autoCodeAdmin', 'autoCodeAdmin', '0', 'view/systemTools/autoCodeAdmin/index.vue', '1', '', '0', '0', '自动化代码管理', 'magic-stick', '0');
INSERT INTO `sys_base_menus` VALUES ('21', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '16', 'autoCodeEdit/:id', 'autoCodeEdit', '1', 'view/systemTools/autoCode/index.vue', '0', '', '0', '0', '自动化代码-${id}', 'magic-stick', '0');
INSERT INTO `sys_base_menus` VALUES ('22', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '16', 'autoPkg', 'autoPkg', '0', 'view/systemTools/autoPkg/autoPkg.vue', '0', '', '0', '0', '自动化package', 'folder', '0');
INSERT INTO `sys_base_menus` VALUES ('23', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '0', 'https://www.gin-vue-admin.com', 'https://www.gin-vue-admin.com', '0', '/', '0', '', '0', '0', '官方网站', 'home-filled', '0');
INSERT INTO `sys_base_menus` VALUES ('24', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '0', 'state', 'state', '0', 'view/system/state.vue', '8', '', '0', '0', '服务器状态', 'cloudy', '0');
INSERT INTO `sys_base_menus` VALUES ('25', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '0', 'plugin', 'plugin', '0', 'view/routerHolder.vue', '6', '', '0', '0', '插件系统', 'cherry', '0');
INSERT INTO `sys_base_menus` VALUES ('26', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '25', 'https://plugin.gin-vue-admin.com/', 'https://plugin.gin-vue-admin.com/', '0', 'https://plugin.gin-vue-admin.com/', '0', '', '0', '0', '插件市场', 'shop', '0');
INSERT INTO `sys_base_menus` VALUES ('27', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '25', 'installPlugin', 'installPlugin', '0', 'view/systemTools/installPlugin/index.vue', '1', '', '0', '0', '插件安装', 'box', '0');
INSERT INTO `sys_base_menus` VALUES ('28', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '25', 'autoPlug', 'autoPlug', '0', 'view/systemTools/autoPlug/autoPlug.vue', '2', '', '0', '0', '插件模板', 'folder', '0');
INSERT INTO `sys_base_menus` VALUES ('29', '2023-02-20 10:25:28.353', '2023-02-20 10:25:28.353', null, '0', '25', 'plugin-email', 'plugin-email', '0', 'plugin/email/view/index.vue', '3', '', '0', '0', '邮件插件', 'message', '0');

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
INSERT INTO `sys_data_authority_id` VALUES ('888', '888');
INSERT INTO `sys_data_authority_id` VALUES ('888', '8881');
INSERT INTO `sys_data_authority_id` VALUES ('888', '9528');
INSERT INTO `sys_data_authority_id` VALUES ('9528', '8881');
INSERT INTO `sys_data_authority_id` VALUES ('9528', '9528');

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
INSERT INTO `sys_dictionaries` VALUES ('1', '2023-02-20 10:25:17.796', '2023-02-20 10:25:18.571', null, '性别', 'gender', '1', '性别字典');
INSERT INTO `sys_dictionaries` VALUES ('2', '2023-02-20 10:25:17.796', '2023-02-20 10:25:19.562', null, '数据库int类型', 'int', '1', 'int类型对应的数据库类型');
INSERT INTO `sys_dictionaries` VALUES ('3', '2023-02-20 10:25:17.796', '2023-02-20 10:25:20.449', null, '数据库时间日期类型', 'time.Time', '1', '数据库时间日期类型');
INSERT INTO `sys_dictionaries` VALUES ('4', '2023-02-20 10:25:17.796', '2023-02-20 10:25:21.319', null, '数据库浮点型', 'float64', '1', '数据库浮点型');
INSERT INTO `sys_dictionaries` VALUES ('5', '2023-02-20 10:25:17.796', '2023-02-20 10:25:22.281', null, '数据库字符串', 'string', '1', '数据库字符串');
INSERT INTO `sys_dictionaries` VALUES ('6', '2023-02-20 10:25:17.796', '2023-02-20 10:25:23.664', null, '数据库bool类型', 'bool', '1', '数据库bool类型');

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `label` varchar(191) DEFAULT NULL COMMENT '展示值',
  `value` bigint(20) DEFAULT NULL COMMENT '字典值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint(20) unsigned DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
INSERT INTO `sys_dictionary_details` VALUES ('1', '2023-02-20 10:25:18.751', '2023-02-20 10:25:18.751', null, '男', '1', '1', '1', '1');
INSERT INTO `sys_dictionary_details` VALUES ('2', '2023-02-20 10:25:18.751', '2023-02-20 10:25:18.751', null, '女', '2', '1', '2', '1');
INSERT INTO `sys_dictionary_details` VALUES ('3', '2023-02-20 10:25:19.787', '2023-02-20 10:25:19.787', null, 'smallint', '1', '1', '1', '2');
INSERT INTO `sys_dictionary_details` VALUES ('4', '2023-02-20 10:25:19.787', '2023-02-20 10:25:19.787', null, 'mediumint', '2', '1', '2', '2');
INSERT INTO `sys_dictionary_details` VALUES ('5', '2023-02-20 10:25:19.787', '2023-02-20 10:25:19.787', null, 'int', '3', '1', '3', '2');
INSERT INTO `sys_dictionary_details` VALUES ('6', '2023-02-20 10:25:19.787', '2023-02-20 10:25:19.787', null, 'bigint', '4', '1', '4', '2');
INSERT INTO `sys_dictionary_details` VALUES ('7', '2023-02-20 10:25:20.620', '2023-02-20 10:25:20.620', null, 'date', '0', '1', '0', '3');
INSERT INTO `sys_dictionary_details` VALUES ('8', '2023-02-20 10:25:20.620', '2023-02-20 10:25:20.620', null, 'time', '1', '1', '1', '3');
INSERT INTO `sys_dictionary_details` VALUES ('9', '2023-02-20 10:25:20.620', '2023-02-20 10:25:20.620', null, 'year', '2', '1', '2', '3');
INSERT INTO `sys_dictionary_details` VALUES ('10', '2023-02-20 10:25:20.620', '2023-02-20 10:25:20.620', null, 'datetime', '3', '1', '3', '3');
INSERT INTO `sys_dictionary_details` VALUES ('11', '2023-02-20 10:25:20.620', '2023-02-20 10:25:20.620', null, 'timestamp', '5', '1', '5', '3');
INSERT INTO `sys_dictionary_details` VALUES ('12', '2023-02-20 10:25:21.477', '2023-02-20 10:25:21.477', null, 'float', '0', '1', '0', '4');
INSERT INTO `sys_dictionary_details` VALUES ('13', '2023-02-20 10:25:21.477', '2023-02-20 10:25:21.477', null, 'double', '1', '1', '1', '4');
INSERT INTO `sys_dictionary_details` VALUES ('14', '2023-02-20 10:25:21.477', '2023-02-20 10:25:21.477', null, 'decimal', '2', '1', '2', '4');
INSERT INTO `sys_dictionary_details` VALUES ('15', '2023-02-20 10:25:22.417', '2023-02-20 10:25:22.417', null, 'char', '0', '1', '0', '5');
INSERT INTO `sys_dictionary_details` VALUES ('16', '2023-02-20 10:25:22.417', '2023-02-20 10:25:22.417', null, 'varchar', '1', '1', '1', '5');
INSERT INTO `sys_dictionary_details` VALUES ('17', '2023-02-20 10:25:22.417', '2023-02-20 10:25:22.417', null, 'tinyblob', '2', '1', '2', '5');
INSERT INTO `sys_dictionary_details` VALUES ('18', '2023-02-20 10:25:22.417', '2023-02-20 10:25:22.417', null, 'tinytext', '3', '1', '3', '5');
INSERT INTO `sys_dictionary_details` VALUES ('19', '2023-02-20 10:25:22.417', '2023-02-20 10:25:22.417', null, 'text', '4', '1', '4', '5');
INSERT INTO `sys_dictionary_details` VALUES ('20', '2023-02-20 10:25:22.417', '2023-02-20 10:25:22.417', null, 'blob', '5', '1', '5', '5');
INSERT INTO `sys_dictionary_details` VALUES ('21', '2023-02-20 10:25:22.417', '2023-02-20 10:25:22.417', null, 'mediumblob', '6', '1', '6', '5');
INSERT INTO `sys_dictionary_details` VALUES ('22', '2023-02-20 10:25:22.417', '2023-02-20 10:25:22.417', null, 'mediumtext', '7', '1', '7', '5');
INSERT INTO `sys_dictionary_details` VALUES ('23', '2023-02-20 10:25:22.417', '2023-02-20 10:25:22.417', null, 'longblob', '8', '1', '8', '5');
INSERT INTO `sys_dictionary_details` VALUES ('24', '2023-02-20 10:25:22.417', '2023-02-20 10:25:22.417', null, 'longtext', '9', '1', '9', '5');
INSERT INTO `sys_dictionary_details` VALUES ('25', '2023-02-20 10:25:23.885', '2023-02-20 10:25:23.885', null, 'tinyint', '0', '1', '0', '6');

-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `ip` varchar(191) DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) DEFAULT NULL COMMENT '请求路径',
  `status` bigint(20) DEFAULT NULL COMMENT '请求状态',
  `latency` bigint(20) DEFAULT NULL COMMENT '延迟',
  `agent` varchar(191) DEFAULT NULL COMMENT '代理',
  `error_message` varchar(191) DEFAULT NULL COMMENT '错误信息',
  `body` text COMMENT '请求Body',
  `resp` text COMMENT '响应Body',
  `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_operation_records_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_operation_records
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority` (
  `sys_user_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `sys_authority_authority_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_user_authority
-- ----------------------------
INSERT INTO `sys_user_authority` VALUES ('1', '888');
INSERT INTO `sys_user_authority` VALUES ('1', '8881');
INSERT INTO `sys_user_authority` VALUES ('1', '9528');
INSERT INTO `sys_user_authority` VALUES ('2', '888');

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(191) DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) DEFAULT '系统用户' COMMENT '用户昵称',
  `side_mode` varchar(191) DEFAULT 'dark' COMMENT '用户侧边主题',
  `header_img` varchar(191) DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `base_color` varchar(191) DEFAULT '#fff' COMMENT '基础颜色',
  `active_color` varchar(191) DEFAULT '#1890ff' COMMENT '活跃颜色',
  `authority_id` bigint(20) unsigned DEFAULT '888' COMMENT '用户角色ID',
  `phone` varchar(191) DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint(20) DEFAULT '1' COMMENT '用户是否被冻结 1正常 2冻结',
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`),
  KEY `idx_sys_users_uuid` (`uuid`),
  KEY `idx_sys_users_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES ('1', '2023-02-20 10:25:25.373', '2023-02-20 10:25:25.881', null, '5d7d0ccf-a2f3-47f8-b273-8c826917e0a8', 'admin', '$2a$10$T5fpZpBzZ62LJAlDUTdCYuVTc0BcINbBi9p/IAz7SfW1hhxp0ADw2', 'Mr.奇淼', 'dark', 'https://qmplusimg.henrongyi.top/gva_header.jpg', '#fff', '#1890ff', '888', '17611111111', '333333333@qq.com', '1');
INSERT INTO `sys_users` VALUES ('2', '2023-02-20 10:25:25.373', '2023-02-20 10:25:27.051', null, '3c485dda-4872-40d2-a0c5-9e6a491ad110', 'a303176530', '$2a$10$HAVsj7axOt65ZfimDO/LlOP4V39Y9WgQ3moSO.iPVi8EzC/DWUxyG', '用户1', 'dark', 'https:///qmplusimg.henrongyi.top/1572075907logo.png', '#fff', '#1890ff', '9528', '17611111111', '333333333@qq.com', '1');
