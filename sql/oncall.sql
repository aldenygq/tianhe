CREATE DATABASE `ops` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;

CREATE TABLE IF NOT EXISTS `ops`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `ctime` int NOT NULL default 0 comment '创建时间',
   `email` VARCHAR(256) NOT NULL default '' comment '邮箱',
   `mtime` int NOT NULL default 0 comment '更新时间',
   `cn_name` VARCHAR(256)  NOT NULL default '' comment '中文名称',
   `en_name` VARCHAR(256)  NOT NULL default '' comment '英文名称',
   `password` VARCHAR(256)  NOT NULL default '' comment '密码',
   `mobile` VARCHAR(16)  NOT NULL default '' comment '手机号码',
   `status` int  NOT NULL default 1 comment '用户状态',
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
