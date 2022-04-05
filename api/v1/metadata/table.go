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
// @Param	id			body int	true "db id"
// @Param	login_name	body string true "account name or employee id"
// @Produce	application/json
// @Success	200 {string} string ""
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
// @Param	db_id		body int	true "db id"
// @Param	table_name	body string	true "table name"
// @Param	login_name	body string true "account name or employee id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router /api/v1/metadata/table/statistic/db-table [get]
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
// @Param	host_ip		body string	true "host ip"
// @Param	port_num	body int	true "port num"
// @Param	db_name		body string	true "db name"
// @Param	table_name	body string	true "table name"
// @Param	login_name	body string true "account name or employee id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router /api/v1/metadata/table/statistic/host-info-db-table [get]
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
// @Param	db_id		body int	true "db id"
// @Param	table_name	body string	true "table name"
// @Param	login_name	body string true "account name or employee id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router /api/v1/metadata/table/analyze/db-table [get]
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
// @Param	host_ip		body string	true "host ip"
// @Param	port_num	body int	true "port num"
// @Param	db_name		body string	true "db name"
// @Param	table_name	body string	true "table name"
// @Param	login_name	body string true "account name or employee id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router /api/v1/metadata/table/analyze/host-info-db-table [get]
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
