package metadata

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pingcap/errors"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/pkg/message"
	msgmeta "github.com/romberli/das/pkg/message/metadata"
	"github.com/romberli/das/pkg/resp"
	utilmeta "github.com/romberli/das/pkg/util/metadata"
	"github.com/romberli/log"
)

const (
	tableStatisticsStruct = "TableStatistics"
	indexStatisticsStruct = "IndexStatistics"
	createStatementStruct = "CreateStatement"

	analyzeDBIDRespMessage     = `{"message": "analyze table completed. db id: %d, table name: %s"}`
	analyzeHostInfoRespMessage = `{"message": "analyze table completed. host ip: %s, port num: %d, db name: %s, table name: %s"}`
)

// @Tags	Tables
// @Summary get tables by db id
// @Accept	application/json
// @Param	token 		body string true "token"
// @Param	id			body int	true "db id"
// @Param	login_name	body string true "account name or employee id"
// @Produce	application/json
// @Success	200 {string} string "{"tables": [{"db_name": "das","table_name": "t_meta_db_info"}]}"
// @Router /api/v1/metadata/table/db [get]
func GetTablesByDBID(c *gin.Context) {
	var rd *utilmeta.TablesByDBID
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	// init service
	ts := metadata.NewTableServiceWithDefault()
	// get entity
	err = ts.GetByDBID(rd.GetDBID(), rd.GetLoginName())
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetTablesByDBID, err, rd.GetDBID(), rd.GetLoginName())
		return
	}
	// marshal service
	jsonBytes, err := ts.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetTablesByDBID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetTablesByDBID, rd.GetDBID(), rd.GetLoginName())
}

// @Tags	Tables
// @Summary get table statistics by db id and table name
// @Accept	application/json
// @Param	token 		body string true "token"
// @Param	db_id		body int	true "db id"
// @Param	table_name	body string	true "table name"
// @Param	login_name	body string true "account name or employee id"
// @Produce	application/json
// @Success	200 {string} string "{"table_statistics": [{"size_mb": 0.015, "engine": "InnoDB", "char_set": "utf8mb4", "collation": "utf8mb4_general_ci", "create_time": "2022-03-04T14:24:11+08:00", "table_name": "t_meta_db_info", "size": 16384, "avg_row_length": 16384, "auto_increment": 2, "db_name": "das", "table_rows": 1}], "index_statistics": [{"is_nullable": false,"db_name": "das", "table_name": "t_meta_db_info", "index_name": "PRIMARY", "sequence":1, "column_name": "id", "cardinality": 1, "is_unique": true}, {"is_nullable": false, "db_name": "das", "table_name": "t_meta_db_info", "index_name": "idx01_db_name_cluster_id_cluster_type_env_id", "sequence": 1, "column_name": "db_name", "cardinality": 1, "is_unique": true}, {"column_name": "cluster_id", "cardinality": 1, "is_unique": true, "is_nullable": false, "db_name": "das", "table_name": "t_meta_db_info", "index_name": "idx01_db_name_cluster_id_cluster_type_env_id", "sequence": 2}, {"cardinality": 1, "is_unique": true, "is_nullable": false, "db_name": "das", "table_name": "t_meta_db_info", "index_name": "idx01_db_name_cluster_id_cluster_type_env_id", "sequence":3 ,"column_name": "cluster_type"}, {"sequence": 4,"column_name": "env_id", "cardinality": 1, "is_unique":true, "is_nullable": false, "db_name": "das", "table_name": "t_meta_db_info", "index_name": "idx01_db_name_cluster_id_cluster_type_env_id"}], "create_statement": "CREATE TABLE `t_meta_db_info` (\n  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',\n  `db_name` varchar(100) NOT NULL COMMENT '数据库名称',\n  `cluster_id` int(11) NOT NULL COMMENT '数据库集群ID',\n  `cluster_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '集群类型: 1-单库, 2-分库分表',\n  `env_id` int(11) NOT NULL COMMENT '环境: 1-online, 2-rel, 3-uat, 4-sit, 5-pt, 6-dev',\n  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',\n  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',\n  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',\n  PRIMARY KEY (`id`),\n  UNIQUE KEY `idx01_db_name_cluster_id_cluster_type_env_id` (`db_name`,`cluster_id`,`cluster_type`,`env_id`)\n) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='数据库信息表'"}"
// @Router /api/v1/metadata/table/statistic/db [get]
func GetStatisticsByDBIDAndTableName(c *gin.Context) {
	var rd *utilmeta.TableByDBIDAndTableName
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	// init service
	ts := metadata.NewTableServiceWithDefault()
	// get entity
	err = ts.GetStatisticsByDBIDAndTableName(rd.GetDBID(), rd.GetTableName(), rd.LoginName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetStatisticsByDBIDAndTableName, err, rd.GetDBID(), rd.GetTableName(), rd.LoginName)
		return
	}
	// marshal service
	jsonBytes, err := ts.MarshalWithFields(tableStatisticsStruct, indexStatisticsStruct, createStatementStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetStatisticsByDBIDAndTableName, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetStatisticsByDBIDAndTableName, rd.GetDBID(), rd.GetTableName(), rd.LoginName)
}

// @Tags	Tables
// @Summary	get table statistics by host info and db name and table name
// @Accept	application/json
// @Param	token 		body string true "token"
// @Param	host_ip		body string	true "host ip"
// @Param	port_num	body int	true "port num"
// @Param	db_name		body string	true "db name"
// @Param	table_name	body string	true "table name"
// @Param	login_name	body string true "account name or employee id"
// @Produce	application/json
// @Success	200 {string} string "{"table_statistics": [{"size_mb": 0.015, "engine": "InnoDB", "char_set": "utf8mb4", "collation": "utf8mb4_general_ci", "create_time": "2022-03-04T14:24:11+08:00", "table_name": "t_meta_db_info", "size": 16384, "avg_row_length": 16384, "auto_increment": 2, "db_name": "das", "table_rows": 1}], "index_statistics": [{"is_nullable": false,"db_name": "das", "table_name": "t_meta_db_info", "index_name": "PRIMARY", "sequence":1, "column_name": "id", "cardinality": 1, "is_unique": true}, {"is_nullable": false, "db_name": "das", "table_name": "t_meta_db_info", "index_name": "idx01_db_name_cluster_id_cluster_type_env_id", "sequence": 1, "column_name": "db_name", "cardinality": 1, "is_unique": true}, {"column_name": "cluster_id", "cardinality": 1, "is_unique": true, "is_nullable": false, "db_name": "das", "table_name": "t_meta_db_info", "index_name": "idx01_db_name_cluster_id_cluster_type_env_id", "sequence": 2}, {"cardinality": 1, "is_unique": true, "is_nullable": false, "db_name": "das", "table_name": "t_meta_db_info", "index_name": "idx01_db_name_cluster_id_cluster_type_env_id", "sequence":3 ,"column_name": "cluster_type"}, {"sequence": 4,"column_name": "env_id", "cardinality": 1, "is_unique":true, "is_nullable": false, "db_name": "das", "table_name": "t_meta_db_info", "index_name": "idx01_db_name_cluster_id_cluster_type_env_id"}], "create_statement": "CREATE TABLE `t_meta_db_info` (\n  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',\n  `db_name` varchar(100) NOT NULL COMMENT '数据库名称',\n  `cluster_id` int(11) NOT NULL COMMENT '数据库集群ID',\n  `cluster_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '集群类型: 1-单库, 2-分库分表',\n  `env_id` int(11) NOT NULL COMMENT '环境: 1-online, 2-rel, 3-uat, 4-sit, 5-pt, 6-dev',\n  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',\n  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',\n  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',\n  PRIMARY KEY (`id`),\n  UNIQUE KEY `idx01_db_name_cluster_id_cluster_type_env_id` (`db_name`,`cluster_id`,`cluster_type`,`env_id`)\n) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='数据库信息表'"}"
// @Router /api/v1/metadata/table/statistic/host-info-db [get]
func GetStatisticsByHostInfoAndDBNameAndTableName(c *gin.Context) {
	var rd *utilmeta.TableByHostInfoAndDBNameAndTableName
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	// init service
	ts := metadata.NewTableServiceWithDefault()
	// get entity
	err = ts.GetStatisticsByHostInfoAndDBNameAndTableName(rd.GetHostIP(), rd.GetPortNum(), rd.GetDBName(), rd.GetTableName(), rd.GetLoginName())
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetStatisticsByHostInfoAndDBNameAndTableName, err, rd.GetHostIP(), rd.GetPortNum(), rd.GetDBName(), rd.GetTableName(), rd.GetLoginName())
		return
	}
	// marshal service
	jsonBytes, err := ts.MarshalWithFields(tableStatisticsStruct, indexStatisticsStruct, createStatementStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetStatisticsByDBIDAndTableName, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetStatisticsByHostInfoAndDBNameAndTableName, rd.GetHostIP(), rd.GetPortNum(), rd.GetDBName(), rd.GetTableName(), rd.GetLoginName())
}

// @Tags	Tables
// @Summary analyze table by db id and table name
// @Accept	application/json
// @Param	token 		body string true "token"
// @Param	db_id		body int	true "db id"
// @Param	table_name	body string	true "table name"
// @Param	login_name	body string true "account name or employee id"
// @Produce	application/json
// @Success	200 {string} string "{"message": "analyze table completed. db id: 1, table name: t_meta_db_info"}"
// @Router /api/v1/metadata/table/analyze/db [post]
func AnalyzeTableByDBIDAndTableName(c *gin.Context) {
	var rd *utilmeta.TableByDBIDAndTableName
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	// init service
	ts := metadata.NewTableServiceWithDefault()
	// get entity
	err = ts.AnalyzeTableByDBIDAndTableName(rd.GetDBID(), rd.GetTableName(), rd.GetLoginName())
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAnalyzeTableByDBIDAndTableName, err, rd.GetDBID(), rd.GetTableName(), rd.GetLoginName())
		return
	}
	// response
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAnalyzeTableByHostInfoAndDBNameAndTableName).Error())
	resp.ResponseOK(c, fmt.Sprintf(analyzeDBIDRespMessage, rd.GetDBID(), rd.GetTableName()), msgmeta.ErrMetadataAnalyzeTableByDBIDAndTableName, rd.GetDBID(), rd.GetTableName(), rd.GetLoginName())
}

// @Tags	Tables
// @Summary analyze table by host info and db name and table name
// @Accept	application/json
// @Param	token 		body string true "token"
// @Param	host_ip		body string	true "host ip"
// @Param	port_num	body int	true "port num"
// @Param	db_name		body string	true "db name"
// @Param	table_name	body string	true "table name"
// @Param	login_name	body string true "account name or employee id"
// @Produce	application/json
// @Success	200 {string} string "{"message": "analyze table completed. host ip: 192.168.137.11, port num: 3306, db name: das, table name: t_meta_db_info"}"
// @Router /api/v1/metadata/table/analyze/host-info-db [post]
func AnalyzeTableByHostInfoAndDBNameAndTableName(c *gin.Context) {
	var rd *utilmeta.TableByHostInfoAndDBNameAndTableName
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	// init service
	ts := metadata.NewTableServiceWithDefault()
	// get entity
	err = ts.AnalyzeTableByHostInfoAndDBNameAndTableName(rd.GetHostIP(), rd.GetPortNum(), rd.GetDBName(), rd.GetTableName(), rd.GetLoginName())
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAnalyzeTableByHostInfoAndDBNameAndTableName, err, rd.GetHostIP(), rd.GetPortNum(), rd.GetDBName(), rd.GetTableName(), rd.GetLoginName())
	}
	// response
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAnalyzeTableByHostInfoAndDBNameAndTableName).Error())
	resp.ResponseOK(c, fmt.Sprintf(analyzeHostInfoRespMessage, rd.GetHostIP(), rd.GetPortNum(), rd.GetDBName(), rd.GetTableName()), msgmeta.InfoMetadataAnalyzeTableByHostInfoAndDBNameAndTableName, rd.GetHostIP(), rd.GetPortNum(), rd.GetDBName(), rd.GetTableName(), rd.GetLoginName())
}
