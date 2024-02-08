-- 用户表
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id，全局唯一',
  `username` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名，可以重复',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '用户密码，采用argon2算法加密',
  `avatar_url` varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像url',
  `background_image_url` varchar(255) NOT NULL DEFAULT '' COMMENT '用户个人页顶部大图url',
  `signature` varchar(255) NOT NULL DEFAULT 'Happy~' COMMENT '个人简介',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'gorm维护，创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'gorm维护，更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'gorm维护，删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_delete_at` (`deleted_at`) USING BTREE COMMENT 'gorm删除时间索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';