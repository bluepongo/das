CREATE TABLE `t_meta_app_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `app_name` varchar(100) NOT NULL COMMENT '应用系统名称',
  `level` tinyint(4) NOT NULL COMMENT '系统等级: 1-A, 2-B, 3-C',
  `owner_id` int(11) DEFAULT NULL COMMENT '应用系统主要负责人ID',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_app_name` (`app_name`),
  KEY `idx02_level` (`level`),
  KEY `idx03_owner_id` (`owner_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用系统信息表';

CREATE TABLE `t_meta_app_db_map` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `app_id` int(11) NOT NULL COMMENT '应用系统ID',
  `db_id` int(11) NOT NULL COMMENT '数据库ID',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_app_id_db_id` (`app_id`,`db_id`),
  KEY `idx02_db_id` (`db_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用系统-数据库映射表';

CREATE TABLE `t_meta_db_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `db_name` varchar(100) NOT NULL COMMENT '数据库名称',
  `cluster_id` int(11) NOT NULL COMMENT '数据库集群ID',
  `cluster_type` tinyint(4) NOT NULL COMMENT '集群类型: 1-单库, 2-分库分表',
  `owner_id` int(11) DEFAULT NULL COMMENT '数据库主要负责人ID',
  `env_id` int(11) NOT NULL COMMENT '环境: 1-online, 2-rel, 3-sit, 4-uat, 5-dev',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_db_name_cluster_id_cluster_type` (`db_name`,`cluster_id`,`cluster_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='数据库信息表';

CREATE TABLE `t_meta_env_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `env_name` varchar(100) NOT NULL COMMENT '环境名称',
  `del_flag` tinyint(4) DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_env_name` (`env_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='环境信息表';

CREATE TABLE `t_meta_middleware_cluster_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `cluster_name` varchar(100) NOT NULL COMMENT '中间件集群名称',
  `owner_id` int(11) DEFAULT NULL COMMENT '中间件主要负责人ID',
  `env_id` int(11) NOT NULL COMMENT '环境: 1-online, 2-rel, 3-sit, 4-uat, 5-dev',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_cluster_name` (`cluster_name`),
  KEY `idx02_owner_id` (`owner_id`),
  KEY `idx03_env_id` (`env_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='中间件集群信息表';

CREATE TABLE `t_meta_middleware_server_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `cluster_id` int(11) NOT NULL COMMENT '中间件集群ID',
  `server_name` varchar(100) NOT NULL COMMENT '中间件服务名称',
  `middleware_role` tinyint(4) NOT NULL COMMENT '中间件角色: 1-rw, 2-ro, 3-das',
  `host_ip` varchar(100) NOT NULL COMMENT '中间件服务器IP',
  `port_num` int(11) NOT NULL COMMENT '中间件端口',
  `del_flag` tinyint(4) DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_host_ip_port_num` (`host_ip`,`port_num`),
  KEY `idx02_cluster_id_middleware_role_env_id` (`cluster_id`,`middleware_role`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='中间件服务器信息表';

CREATE TABLE `t_meta_monitor_system_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `system_name` varchar(100) NOT NULL COMMENT '监控系统名称',
  `system_type` tinyint(4) NOT NULL COMMENT '监控系统类型: 1-pmm1.x, 2-pmm2.x',
  `host_ip` varchar(100) NOT NULL COMMENT '监控系统服务器IP',
  `port_num` int(11) NOT NULL COMMENT '监控系统服务器端口',
  `port_num_slow` int(11) NOT NULL COMMENT '监控系统服务器慢查询日志端口',
  `base_url` varchar(200) NOT NULL COMMENT '监控系统API入口地址',
  `env_id` tinyint(4) NOT NULL COMMENT '环境: 1-online, 2-rel, 3-sit, 4-uat, 5-dev',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_system_name` (`system_name`),
  UNIQUE KEY `idx02_host_ip_port_num` (`host_ip`,`port_num`),
  KEY `idx03_env_id` (`env_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='监控系统信息表';

CREATE TABLE `t_meta_mysql_cluster_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `cluster_name` varchar(100) NOT NULL COMMENT '集群名称',
  `middleware_cluster_id` int(11) DEFAULT NULL COMMENT '中间件集群ID',
  `monitor_system_id` int(11) DEFAULT NULL COMMENT '监控系统ID',
  `owner_id` int(11) DEFAULT NULL COMMENT '数据库集群主要负责人ID',
  `env_id` int(11) NOT NULL COMMENT '环境: 1-online, 2-rel, 3-sit, 4-uat, 5-dev',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_cluster_name` (`cluster_name`),
  KEY `idx03_monitor_system_id` (`monitor_system_id`),
  KEY `idx04_owner_id` (`owner_id`),
  key `idx05_env_id` (`env_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='MySQL集群信息表';

CREATE TABLE `t_meta_mysql_server_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `cluster_id` int(11) NOT NULL COMMENT '集群ID',
  `server_name` varchar(100) NOT NULL COMMENT '数据库实例名称',
  `host_ip` varchar(100) NOT NULL COMMENT '服务器IP',
  `port_num` int(11) NOT NULL COMMENT '端口',
  `deployment_type` tinyint(4) NOT NULL COMMENT '部署方式: 1-容器, 2-物理机, 3-虚拟机',
  `version` varchar(100) DEFAULT NULL COMMENT '版本, 示例: 5.7.21',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  KEY `idx01_cluster_id` (`cluster_id`),
  UNIQUE KEY `idx02_host_ip_port_num` (`host_ip`,`port_num`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='MySQL服务器信息表';

CREATE TABLE `t_meta_user_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_name` varchar(100) NOT NULL COMMENT '姓名',
  `department_name` varchar(100) NOT NULL COMMENT '部门/团队名称',
  `employee_id` varchar(100) NOT NULL COMMENT '工号',
  `account_name` varchar(100) NOT NULL COMMENT '账号名称',
  `email` varchar(100) NOT NULL COMMENT '邮箱',
  `telephone` varchar(100) DEFAULT NULL COMMENT '固定电话',
  `mobile` varchar(100) DEFAULT NULL COMMENT '手机号码',
  `role` tinyint(4) NOT NULL DEFAULT '3' COMMENT '角色: 1-admin, 2-dba, 3-developer',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_employee_id` (`employee_id`),
  UNIQUE KEY `idx02_account_name` (`account_name`),
  UNIQUE KEY `idx03_user_name` (`email`),
  KEY `idx04_user_name` (`user_name`),
  KEY `idx05_user_name` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';