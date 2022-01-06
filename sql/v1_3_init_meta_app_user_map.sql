CREATE TABLE `t_meta_app_user_map` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `app_id` int(11) NOT NULL COMMENT '应用系统ID',
  `user_id` int(11) NOT NULL COMMENT '用户ID',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_app_id_user_id` (`app_id`, `user_id`),
  KEY `idx02_user_id` (`user_id`)
) ENGINE = Innodb DEFAULT CHARSET = utf8mb4 COMMENT = '应用系统-用户映射表';

insert into t_meta_app_user_map(app_id,user_id) select id,owner_id from t_meta_app_info;