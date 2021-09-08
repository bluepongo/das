#1. 健康检查配置
健康检查按照若干个检查项来进行打分, 每个检查项独立以满分100分来进行打分, 并乘以该项的权重后计入总分, 所有检查项的权重和应等于100.

DAS的健康检查模块会从数据库读取配置,表结构如下:
```sql
CREATE TABLE `t_hc_default_engine_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `item_name` varchar(100) NOT NULL COMMENT '检查项名称',
  `item_weight` int NOT NULL COMMENT '权重百分比, 所有检查项项权重合计应等于100',
  `low_watermark` decimal(10, 2) NOT NULL COMMENT '低水位',
  `high_watermark` decimal(10, 2) NOT NULL COMMENT '高水位',
  `unit` decimal(10, 2) NOT NULL COMMENT '百分比, 每超过该百分比时会扣分',
  `score_deduction_per_unit_high` decimal(10, 2) NOT NULL COMMENT '高指标每单位扣分分数',
  `max_score_deduction_high` decimal(10, 2) NOT NULL COMMENT '高指标最多扣分数',
  `score_deduction_per_unit_medium` decimal(10, 2) NOT NULL COMMENT '中指标每单位扣分分数',
  `max_score_deduction_medium` decimal(10, 2) NOT NULL COMMENT '中指标最多扣分数',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_item_name` (`item_name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '健康检查默认引擎配置表';
```


#2. 检查项
目前检查项共有10种,分别为
- 参数配置
- cpu使用率
- io使用率
- 磁盘空间使用率
- 连接数使用率
- 活跃会话百分比
- 缓存未命中率
- 表行数
- 表大小
- 慢查询

其中参数配置又可以细分为19个子项, 分别为:

|参数名                         |建议值        |
|------------------------------|-------------|
|max_user_connection           |\>=2000      |
|slave_parallel_workers        |16           |
|log_bin                       |ON           |
|binlog_format                 |ROW          |
|binlog_row_image              |FULL         |
|sync_binlog                   |1            |
|innodb_flush_log_at_trx_commit|1            |
|gtid_mode                     |ON           |
|enforce_gtid_consistency      |ON           |
|slave_parallel_type           |LOGICAL_CLOCK|
|master_info_repository        |TABLE        |
|relay_log_info_repository     |TABLE        |
|report_host                   |不为空        |
|report_port                   |不为空        |
|innodb_flush_method           |O_DIRECT     |
|innodb_monitor_enable         |all          |
|innodb_print_all_deadlocks    |ON           |
|slow_query_log                |ON           |
|performance_schema            |ON           |


#3. 计分规则

##3.1. `参数配置`
- 统计所有不符合要求的参数配置, 记为`count`
- 计算`count` * `score_deduction_per_unit_high`, 记为`score_deduction`
- 如果`score_deduction` < `max_score_deduction_high`, 则令`score_deduction`等于`max_score_deduction_high`, 即以`max_score_deduction_high`为扣分上限
- 计算`100` - `score_deduction`, 记为`item_score`
- 计算`item_score` * `item_weight`, 记为`weighted_item_score`
- `weighted_item_score`为检查项`参数配置`的加权分数


##3.2. 其他检查项

通过调用PMM接口获取该检查项的监控数据并循环遍历该数据

###3.2.1. 计算高危数据的扣分数
- 对所有值高于`high_watermark`的值进行求和, 记为`sum_high`
- 对所有值高于`high_watermark`的值进行计数, 记为`count_high`
- 计算`sum_high` / `count_high`,以求得平均数, 记为`avg_high`
- 计算`avg_high` / `unit` * `score_deduction_per_unit_high`, 记为`score_deduction_high`
- 如果`score_deduction_high`的值大于`max_score_deduction_high`, 则令`score_deduction_high`等于`max_score_deduction_high`, 即以`max_score_deduction_high`为扣分上限

###3.2.2. 计算中危数据的扣分数
- 对所有值处于`low_watermark`与`high_waterark`的值进行求和, 记为`sum_medium`
- 对所有值处于`low_watermark`与`high_waterark`的值进行计数, 记为`count_medium`
- 计算`sum_medium` / `count_medium`,以求得平均数, 记为`avg_medium`
- 计算`avg_medium` / `unit` * `score_deduction_per_unit_medium`, 记为`score_deduction_medium`
- 如果`score_deduction_medium`的值大于`max_score_deduction_medium`, 则令`score_deduction_medium`等于`max_score_deduction_medium`, 即以`max_score_deduction_medium`为扣分上限

###3.2.3. 计算加权分数
- 计算`100` - `score_deduction_high` - `score_deduction_medium`, 记为`item_score`
- 计算`item_score` * `item_weight`,  记为`weighted_item_score`
-  `weighted_item_score`为该检查项的加权分数


##3.3. 计算总分

对所有检查项的加权分数进行求和即得到该实例的总分数, 代表该实例总体的健康状况


##3.4. 其他说明

- 慢查询以`最大单次扫描行数`作为监控数据并按3.2中的规则进行计分
- 低于`low_watermark`的数值认为是正常使用量, 不对其进行扣分
- `low_watermark`的值必须小于`high_watermark`


#4. 初始化数据

|item_name|item_weight|low_watermark|high_watermark|unit|score_deduction_per_unit_high|max_score_deduction_high|score_deduction_per_unit_medium|max_score_deduction_medium|
|:--------|----------:|------------:|-------------:|---:|----------------------------:|-----------------------:|------------------------------:|-------------------------:|
|参数配置     |   5|         0|         0|        0|   10|   50|   0|   0|
|cpu使用率    |   5|        50|        80|       10|   20|  100|  10|  50|
|io使用率     |   5|       0.5|       0.8|      0.1|   20|  100|  10|  50|
|磁盘空间使用率|  20|       0.5|       0.8|      0.1|   40|  100|  10|  50|
|连接数使用率  |  20|       0.5|       0.8|      0.1|   40|  100|  10|  50|
|活跃会话百分比|  10|       0.1|       0.3|      0.1|   20|  100|   5|  50|
|缓存未命中率  |   5|     0.005|      0.02|     0.01|   20|  100|  10|  50|
|表行数       |   5|  10000000|  30000000|  1000000|   10|   50|  10|  30|
|表大小       |   5|        10|        30|        5|   10|   50|  10|  30|
|慢查询       |  20|    100000|    500000|    10000|   10|  100|   5|  50|


配置表中的各个字段的值对最终分数影响很大, 需要通过后续的迭代来优化各项值
