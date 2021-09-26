insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('db_config', 5, 0, 0, 0, 10, 50, 0, 0);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('avg_backup_failed_ratio', 5, 0.1, 0.2, 0.1, 20, 100, 10, 50);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('statistics_failed_ratio', 5, 0.1, 0.2, 0.1, 20, 100, 10, 50);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('cpu_usage', 5, 0.5, 0.8, 0.1, 20, 100, 10, 50);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('io_util', 5, 0.5, 0.8, 0.1, 20, 100, 10, 50);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('disk_capacity_usage', 20, 0.5, 0.8, 0.1, 40, 100, 10, 50);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('connection_usage', 10, 0.5, 0.8, 0.1, 40, 100, 10, 50);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('average_active_session_percents', 10, 0.1, 0.3, 0.1, 20, 100, 10, 50);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('cache_miss_ratio', 5, 0.005, 0.02, 0.01, 20, 100, 10, 50);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('table_rows', 5, 10000000, 30000000, 1000000, 10, 50, 10, 30);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('table_size', 5, 10, 30, 5, 10, 50, 10, 30);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('slow_query_rows_examined', 20, 100000, 500000, 100000, 10, 100, 5, 50);
