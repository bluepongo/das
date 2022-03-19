package metadata

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pingcap/errors"
	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/pkg/message"
	msgmeta "github.com/romberli/das/pkg/message/metadata"
	"github.com/romberli/das/pkg/resp"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
)

// @Tags	Tables
// @Summary
// @Accept	application/json
// @Param	id path int true "db id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router /api/v1/metadata/table/db/:id
func GetTablesByDBID(c *gin.Context) {
	// get param
	dbIDStr := c.Param(mysqlClusterIDJSON)
	if dbIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterIDJSON)
		return
	}
	dbID, err := strconv.Atoi(dbIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return

	}
	ds := metadata.NewDBServiceWithDefault()

	err = ds.GetByID(dbID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBByID, dbID, err)
		return
	}
	dbName := ds.GetDBs()[constant.ZeroInt].GetDBName()

	err = ds.GetMySQLClusterByID(dbID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLClusterByDBID, dbID, err)
		return
	}
	mysqlCluster := ds.GetMySQLCluster()

	masterServers, err := mysqlCluster.GetMasterServers()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMasterServers, mysqlCluster.Identity(), err)
		return
	}

	hostIP := masterServers[constant.ZeroInt].GetHostIP()
	portNum := masterServers[constant.ZeroInt].GetPortNum()
	dbAddr := fmt.Sprintf("%s:%d", hostIP, portNum)
	dbUser := global.DASMySQLPool.Config.DBUser
	dbPass := global.DASMySQLPool.Config.DBPass
	conn, err := mysql.NewConn(dbAddr, dbName, dbUser, dbPass)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataTableCreateApplicationMySQLConn, dbAddr, dbName, err)
		return
	}
	tableRepo := metadata.NewTableRepo(conn)

	// init service
	ts := metadata.NewTableService(tableRepo)
	// get entity
	err = ts.GetByDBName(dbName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetTablesByDBID, dbID, err)
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
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetTablesByDBID, dbID)
}

// @Tags	Tables
// @Summary
// @Accept	application/json
// @Param	db_id		body int	true "db id"
// @Param	table_name	body string	true "table name"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router /api/v1/metadata/table/statistic/db-table
func GetStatisticsByDBIDAndTableName(c *gin.Context) {
	// get params
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}

	dataMap := make(map[string]interface{})
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	dbIDInterface, dbIDExists := dataMap[mysqlClusterUserIDJSON]
	if !dbIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterUserIDJSON)
		return
	}
	dbID := dbIDInterface.(int)
	tableNameInterface, tableNameExists := dataMap[mysqlClusterUserIDJSON]
	if !tableNameExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterUserIDJSON)
		return
	}
	tableName := tableNameInterface.(string)

	ds := metadata.NewDBServiceWithDefault()

	err = ds.GetByID(dbID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBByID, dbID, err)
		return
	}
	dbName := ds.GetDBs()[constant.ZeroInt].GetDBName()

	err = ds.GetMySQLClusterByID(dbID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLClusterByDBID, dbID, err)
		return
	}
	mysqlCluster := ds.GetMySQLCluster()

	masterServers, err := mysqlCluster.GetMasterServers()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMasterServers, mysqlCluster.Identity(), err)
		return
	}

	hostIP := masterServers[constant.ZeroInt].GetHostIP()
	portNum := masterServers[constant.ZeroInt].GetPortNum()
	dbAddr := fmt.Sprintf("%s:%d", hostIP, portNum)
	dbUser := global.DASMySQLPool.Config.DBUser
	dbPass := global.DASMySQLPool.Config.DBPass
	conn, err := mysql.NewConn(dbAddr, dbName, dbUser, dbPass)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataTableCreateApplicationMySQLConn, dbAddr, dbName, err)
		return
	}
	tableRepo := metadata.NewTableRepo(conn)

	// init service
	ts := metadata.NewTableService(tableRepo)
	// get entity
	err = ts.GetStatisticsByDBNameAndTableName(dbName, tableName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetStatisticsByDBIDAndTableName, err, dbID, tableName)
	}
	// marshal service
	jsonBytes, err := ts.MarshalWithFields()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetStatisticsByDBIDAndTableName, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetStatisticsByDBIDAndTableName, dbID, tableName)
}

// @Tags	Tables
// @Summary
// @Accept	application/json
// @Param
// @Produce	application/json
// @Success	200 {string} string ""
// @Router /api/v1/metadata/table/statistic/host-info-db-table
func GetStatisticsByHostInfoAndDBNameAndTableName(c *gin.Context) {

}

// @Tags	Tables
// @Summary
// @Accept	application/json
// TODO: @Param
// @Produce	application/json
// @Success	200 {string} string ""
// @Router /api/v1/metadata/table/analyze/db
func AnalyzeTableByDBIDAndTableName(c *gin.Context) {

}

// @Tags	Tables
// @Summary
// @Accept	application/json
// TODO: @Param
// @Produce	application/json
// @Success	200 {string} string ""
// @Router /api/v1/metadata/table/analyze/host-info
func AnalyzeTableByHostInfoAndDBNameAndTableName(c *gin.Context) {

}
