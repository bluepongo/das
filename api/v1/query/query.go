package query

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pingcap/errors"
	"github.com/romberli/das/internal/app/query"
	"github.com/romberli/das/pkg/message"
	msgquery "github.com/romberli/das/pkg/message/query"
	"github.com/romberli/das/pkg/resp"
	utilquery "github.com/romberli/das/pkg/util/query"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
)

const (
	mysqlClusterIDJSON     = "mysql_cluster_id"
	mysqlServerIDJSON      = "mysql_server_id"
	mysqlServerHostIPJSON  = "host_ip"
	mysqlServerPortNumJSON = "port_num"
	dbIDJSON               = "db_id"
	sqlIDJSON              = "sql_id"
)

// @Tags query
// @Summary get slow queries by mysql server id
// @Accept  application/json
// @Param	mysql_cluster_id 	path int	true "mysql cluster id"
// @Param	start_time			body string true "start time"
// @Param	end_time			body string true "end time"
// @Param	limit				body int	true "limit"
// @Param	offset				body int	true "offset"
// @Produce application/json
// @Success 200 {string} string "{"queries":[{"sql_id":"F9A57DD5A41825CA","fingerprint":"select sleep(?)","example":"select sleep(3)","db_name":"","exec_count":1,"total_exec_time":3,"avg_exec_time":3,"rows_examined_max":0}]}"
// @Router	/api/v1/query/cluster/:mysql_cluster_id [get]
func GetByMySQLClusterID(c *gin.Context) {
	// get data
	mysqlClusterIDStr := c.Param(mysqlClusterIDJSON)
	if mysqlClusterIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterIDJSON)
		return
	}
	mysqlClusterID, err := strconv.Atoi(mysqlClusterIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}

	var rd *utilquery.Range
	// bind json
	err = c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	config, err := rd.GetConfig()
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	// init server
	service := query.NewServiceWithDefault(config)
	err = service.GetByMySQLClusterID(mysqlClusterID)
	if err != nil {
		resp.ResponseNOK(c, msgquery.ErrQueryGetByMySQLClusterID, err, mysqlClusterID)
		return
	}
	// marshal
	jsonBytes, err := service.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgquery.DebugQueryGetByMySQLClusterID, mysqlClusterID, jsonStr).Error())

	// response
	resp.ResponseOK(c, jsonStr, msgquery.InfoQueryGetByMySQLClusterID, mysqlClusterID)
}

// @Tags query
// @Summary get slow queries by mysql server id
// @Accept  application/json
// @Param	mysql_server_id	path int	true "mysql server id"
// @Param	start_time		body string true "start time"
// @Param	end_time		body string true "end time"
// @Param	limit			body int	true "limit"
// @Param	offset			body int	true "offset"
// @Produce  application/json
// @Success 200 {string} string "{"queries":[{"sql_id":"F9A57DD5A41825CA","fingerprint":"select sleep(?)","example":"select sleep(3)","db_name":"","exec_count":1,"total_exec_time":3,"avg_exec_time":3,"rows_examined_max":0}]}"
// @Router /api/v1/query/server/:mysql_server_id [get]
func GetByMySQLServerID(c *gin.Context) {
	// get data
	mysqlServerIDStr := c.Param(mysqlServerIDJSON)
	if mysqlServerIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlServerIDJSON)
		return
	}
	mysqlServerID, err := strconv.Atoi(mysqlServerIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	var rd *utilquery.Range
	// bind json
	err = c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	config, err := rd.GetConfig()
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	// init service
	service := query.NewServiceWithDefault(config)
	err = service.GetByMySQLServerID(mysqlServerID)
	if err != nil {
		resp.ResponseNOK(c, msgquery.ErrQueryGetByMySQLServerID, err, mysqlServerID)
		return
	}

	// marshal
	jsonBytes, err := service.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgquery.DebugQueryGetByMySQLServerID, mysqlServerID, jsonStr).Error())

	// response
	resp.ResponseOK(c, jsonStr, msgquery.InfoQueryGetByMySQLServerID, mysqlServerID)
}

// @Tags query
// @Summary get slow queries by mysql server host ip and port number
// @Accept  application/json
// @Param	host_ip		body 	string	true "mysql server host ip"
// @Param	port_num	body	int		true "mysql server port number"
// @Param	start_time	body	string	true "start time"
// @Param	end_time	body	string	true "end time"
// @Param	limit		body	int		true "limit"
// @Param	offset		body	int		true "offset"
// @Produce  application/json
// @Success 200 {string} string "{"queries":[{"sql_id":"F9A57DD5A41825CA","fingerprint":"select sleep(?)","example":"select sleep(3)","db_name":"","exec_count":1,"total_exec_time":3,"avg_exec_time":3,"rows_examined_max":0}]}"
// @Router /api/v1/query/host-info [get]
func GetByHostInfo(c *gin.Context) {
	// get data

	var rd *utilquery.HostInfoRange
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	config, err := rd.GetConfig()
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	// init service
	service := query.NewServiceWithDefault(config)
	err = service.GetByHostInfo(rd.GetHostIP(), rd.GetPortNum())
	if err != nil {
		resp.ResponseNOK(c, msgquery.ErrQueryGetByHostInfo, err, rd.GetHostIP(), rd.GetPortNum())
		return
	}

	// marshal
	jsonBytes, err := service.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgquery.DebugQueryGetByHostInfo, rd.GetHostIP(), rd.GetPortNum(), jsonStr).Error())

	// response
	resp.ResponseOK(c, jsonStr, msgquery.InfoQueryGetByMySQLServerID, rd.GetHostIP(), rd.GetPortNum())
}

// @Tags query
// @Summary get slow queries by db id
// @Accept  application/json
// @Param	db_id			path int	true "db id"
// @Param	mysql_server_id	body int	true "mysql server id"
// @Param	start_time		body string true "start time"
// @Param	end_time		body string true "end time"
// @Param	limit			body int	true "limit"
// @Param	offset			body int	true "offset"
// @Produce application/json
// @Success 200 {string} string "{"queries":[{"sql_id":"F9A57DD5A41825CA","fingerprint":"select sleep(?)","example":"select sleep(3)","db_name":"","exec_count":1,"total_exec_time":3,"avg_exec_time":3,"rows_examined_max":0}]}"
// @Router	/api/v1/query/db/:db_id [get]
func GetByDBID(c *gin.Context) {
	// get data
	dbIDStr := c.Param(dbIDJSON)
	if dbIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
		return
	}
	dbID, err := strconv.Atoi(dbIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}

	var rd *utilquery.ServerRange
	// bind json
	err = c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	config, err := rd.GetConfig()
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	// init service
	service := query.NewServiceWithDefault(config)
	err = service.GetByDBID(rd.GetMySQLServerID(), dbID)
	if err != nil {
		resp.ResponseNOK(c, msgquery.DebugQueryGetByDBID, err, dbID)
		return
	}

	// marshal
	jsonBytes, err := service.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
	}
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgquery.DebugQueryGetByDBID, dbID, jsonStr).Error())

	// response
	resp.ResponseOK(c, jsonStr, msgquery.DebugQueryGetByDBID, dbID)
}

// @Tags query
// @Summary get slow query by query id
// @Accept  application/json
// @Param	sql_id			path int	true "sql id"
// @Param	mysql_server_id	body int	true "mysql server id"
// @Param	start_time		body string true "start time"
// @Param	end_time		body string true "end time"
// @Param	limit			body int	true "limit"
// @Param	offset			body int	true "offset"
// @Produce application/json
// @Success 200 {string} string "{"queries":[{"sql_id":"F9A57DD5A41825CA","fingerprint":"select sleep(?)","example":"select sleep(3)","db_name":"","exec_count":1,"total_exec_time":3,"avg_exec_time":3,"rows_examined_max":0}]}"
// @Router	/api/v1/query/:sql_id [get]
func GetBySQLID(c *gin.Context) {
	// get data
	sqlIDStr := c.Param(sqlIDJSON)
	if sqlIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, sqlIDJSON)
		return
	}
	var rd *utilquery.ServerRange
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	config, err := rd.GetConfig()
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	// init service
	service := query.NewServiceWithDefault(config)
	err = service.GetBySQLID(rd.GetMySQLServerID(), sqlIDStr)
	if err != nil {
		resp.ResponseNOK(c, msgquery.DebugQueryGetBySQLID, err, rd.GetMySQLServerID(), sqlIDStr)
		return
	}

	// marshal
	jsonBytes, err := service.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
	}
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgquery.DebugQueryGetBySQLID, rd.GetMySQLServerID(), sqlIDStr).Error())

	// response
	resp.ResponseOK(c, jsonStr, msgquery.InfoQueryGetBySQLID, sqlIDStr)
}
