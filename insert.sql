CREATE TABLE `increase` (
  `id` int(11) unsigned NOT NULL,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '',
  `value` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

INSERT INTO `test`.`increase`(`id`, `name`, `value`) VALUES (1, 'jxc', 1);

#需要多少个表就建立多少个表
CREATE TABLE `short_url_0` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(500) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT 'url',
  `url_md5` char(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `createtime` int(11) unsigned NOT NULL DEFAULT '0',
  `short_url` varchar(10) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_short_url` (`short_url`) USING BTREE,
  KEY `idx_url_md5` (`url_md5`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;