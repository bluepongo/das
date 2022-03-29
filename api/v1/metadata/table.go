package metadata

import (
	"fmt"
	"strconv"

	"github.com/romberli/das/config"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/pingcap/errors"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/internal/app/privilege"
	"github.com/romberli/das/pkg/message"
	msgmeta "github.com/romberli/das/pkg/message/metadata"
	"github.com/romberli/das/pkg/resp"
	utilmeta "github.com/romberli/das/pkg/util/metadata"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
)

const (
	tableNameJSON        = "table_name"
	tableHostIPJSON      = "host_ip"
	tablePortNumJSON     = "port_num"
	tableDBIDJSON        = "db_id"
	tableDBNameJSON      = "db_name"
	tableAccountNameJSON = "account_name"

	analyzeRespMessage = `{"message": "analyze table completed. db name: %s, table name: %s"}`
)

// @Tags	Tables
// @Summary get tables by db id
// @Accept	application/json
// @Param	id path int true "db id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router /api/v1/metadata/table/db/:db_id
func GetTablesByDBID(c *gin.Context) {
	// get param
	dbIDStr := c.Param(tableDBIDJSON)
	if dbIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, tableDBIDJSON)
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
	// 获取application 的
	// dbUser := global.DASMySQLPool.Config.DBUser
	// dbPass := global.DASMySQLPool.Config.DBPass
	dbUser := viper.GetString(config.DBApplicationMySQLUserKey)
	dbPass := viper.GetString(config.DBApplicationMySQLPassKey)
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
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetTablesByDBID, dbID)
}

// @Tags	Tables
// @Summary get table statistics by db id and table name
// @Accept	application/json
// @Param	db_id		body int	true "db id"
// @Param	table_name	body string	true "table name"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router /api/v1/metadata/table/statistic/db-table
func GetStatisticsByDBIDAndTableName(c *gin.Context) {
	var rd *utilmeta.GetStatisticsByDBIDAndTableName
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	dbID := rd.GetDBID()
	tableName := rd.GetTableName()
	// get host info, db name and table name
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
	dbUser := viper.GetString(config.DBApplicationMySQLUserKey)
	dbPass := viper.GetString(config.DBApplicationMySQLPassKey)
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
		return
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
// @Summary	get table statistics by host info and db name and table name
// @Accept	application/json
// @Param	host_ip		body string	true "host ip"
// @Param	port_num	body int	true "port num"
// @Param	db_name		body string	true "db name"
// @Param	table_name	body string	true "table name"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router /api/v1/metadata/table/statistic/host-info-db-table
func GetStatisticsByHostInfoAndDBNameAndTableName(c *gin.Context) {
	var rd *utilmeta.GetStatisticsByHostInfoAndDBNameAndTableName
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	hostIP := rd.GetHostIP()
	portNum := rd.GetPortNum()
	dbName := rd.GetDBName()
	tableName := rd.GetTableName()

	dbAddr := fmt.Sprintf("%s:%d", hostIP, portNum)
	dbUser := viper.GetString(config.DBApplicationMySQLUserKey)
	dbPass := viper.GetString(config.DBApplicationMySQLPassKey)
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
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetStatisticsByHostInfoAndDBNameAndTableName, err, hostIP, portNum, dbName, tableName)
		return
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
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetStatisticsByHostInfoAndDBNameAndTableName, hostIP, portNum, dbName, tableName)
}

// @Tags	Tables
// @Summary analyze table by db id and table name
// @Accept	application/json
// @Param	db_id			body int	true "db id"
// @Param	table_name		body string	true "table name"
// @Param	account_name	body string true "account name or employee id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router /api/v1/metadata/table/analyze/db
func AnalyzeTableByDBIDAndTableName(c *gin.Context) {
	var rd *utilmeta.AnalyzeTableByDBIDAndTableName
	// bind json
	err := c.ShouldBind(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	dbID := rd.GetDBID()
	tableName := rd.GetTableName()
	loginName := rd.GetLoginName()
	// check privilege
	us := metadata.NewUserServiceWithDefault()
	err = us.GetByAccountNameOrEmployeeID(loginName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetByAccountNameOrEmployeeID, loginName, err)
		return
	}
	privilegeService := privilege.NewService(us.GetUsers()[constant.ZeroInt])
	err = privilegeService.CheckDBByID(dbID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAnalyzeTableByDBIDAndTableName, err, dbID, tableName)
		return
	}
	// get host info, db name and table name
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
	dbUser := viper.GetString(config.DBApplicationMySQLUserKey)
	dbPass := viper.GetString(config.DBApplicationMySQLPassKey)
	conn, err := mysql.NewConn(dbAddr, dbName, dbUser, dbPass)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataTableCreateApplicationMySQLConn, dbAddr, dbName, err)
		return
	}
	tableRepo := metadata.NewTableRepo(conn)
	// init service
	ts := metadata.NewTableService(tableRepo)
	// get entity
	err = ts.AnalyzeTableByDBNameAndTableName(dbName, tableName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAnalyzeTableByDBIDAndTableName, err, dbID, tableName)
		return
	}
	// response
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAnalyzeTableByHostInfoAndDBNameAndTableName).Error())
	resp.ResponseOK(c, fmt.Sprintf(analyzeRespMessage, dbName, tableName), msgmeta.InfoMetadataAnalyzeTableByHostInfoAndDBNameAndTableName, hostIP, portNum, dbName, tableName)
}

// @Tags	Tables
// @Summary analyze table by host info and db name and table name
// @Accept	application/json
// @Param	host_ip			body string	true "host ip"
// @Param	port_num		body int	true "port num"
// @Param	db_name			body string	true "db name"
// @Param	table_name		body string	true "table name"
// @Param	account_name	body string true "account name or employee id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router /api/v1/metadata/table/analyze/host-info
func AnalyzeTableByHostInfoAndDBNameAndTableName(c *gin.Context) {
	var rd *utilmeta.AnalyzeTableByHostInfoAndDBNameAndTableName
	// bind json
	err := c.ShouldBind(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	hostIP := rd.GetHostIP()
	portNum := rd.GetPortNum()
	dbName := rd.GetDBName()
	tableName := rd.GetTableName()
	loginName := rd.GetLoginName()
	// check privilege
	us := metadata.NewUserServiceWithDefault()
	err = us.GetByAccountNameOrEmployeeID(loginName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetByAccountNameOrEmployeeID, loginName, err)
		return
	}
	privilegeService := privilege.NewService(us.GetUsers()[constant.ZeroInt])
	err = privilegeService.CheckDBByNameAndHostInfo(dbName, hostIP, portNum)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAnalyzeTableByHostInfoAndDBNameAndTableName, err, hostIP, portNum, dbName, tableName)
		return
	}
	dbAddr := fmt.Sprintf("%s:%d", hostIP, portNum)
	dbUser := viper.GetString(config.DBApplicationMySQLUserKey)
	dbPass := viper.GetString(config.DBApplicationMySQLPassKey)
	conn, err := mysql.NewConn(dbAddr, dbName, dbUser, dbPass)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataTableCreateApplicationMySQLConn, dbAddr, dbName, err)
		return
	}
	tableRepo := metadata.NewTableRepo(conn)

	// init service
	ts := metadata.NewTableService(tableRepo)
	// get entity
	err = ts.AnalyzeTableByDBNameAndTableName(dbName, tableName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAnalyzeTableByHostInfoAndDBNameAndTableName, err, hostIP, portNum, dbName, tableName)
	}
	// response
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAnalyzeTableByHostInfoAndDBNameAndTableName).Error())
	resp.ResponseOK(c, fmt.Sprintf(analyzeRespMessage, dbName, tableName), msgmeta.InfoMetadataAnalyzeTableByHostInfoAndDBNameAndTableName, hostIP, portNum, dbName, tableName)
}
