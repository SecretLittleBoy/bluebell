create table if not exists `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) collate=utf8mb4_unicode_ci NOT NULL,
    `password` varchar(64) collate=utf8mb4_unicode_ci NOT NULL,
    `email` varchar(64) collate=utf8mb4_unicode_ci NOT NULL,
    `gender` tinyint(4) NOT NULL default 0,
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) engine=innodb default charset=utf8mb4 collate=utf8mb4_unicode_ci;