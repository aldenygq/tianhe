CREATE DATABASE `ops` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;

CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `ctime` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `email` varchar(256) NOT NULL DEFAULT '' COMMENT '邮箱',
  `mtime` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `en_name` varchar(256) NOT NULL DEFAULT '' COMMENT '英文名称',
  `password` varchar(256) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(16) NOT NULL DEFAULT '' COMMENT '手机号码',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '用户状态',
  `create_type` varchar(16) NOT NULL DEFAULT '' COMMENT '创建类型',
  `creator` varchar(256) NOT NULL DEFAULT '' COMMENT '创建人',
  `expire_time` int(11) NOT NULL DEFAULT '0' COMMENT 'token有效期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
