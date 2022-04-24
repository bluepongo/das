package query

import (
	"github.com/gin-gonic/gin"
	"github.com/pingcap/errors"
	"github.com/romberli/das/internal/app/query"
	"github.com/romberli/das/pkg/message"
	msgquery "github.com/romberli/das/pkg/message/query"
	"github.com/romberli/das/pkg/resp"
	utilquery "github.com/romberli/das/pkg/util/query"
	"github.com/romberli/log"
)

// @Tags query
// @Summary get slow queries by mysql server id
// @Accept  application/json
// @Param	token	 			body string true "token"
// @Param	mysql_cluster_id 	body int	true "mysql cluster id"
// @Param	start_time			body string true "start time"
// @Param	end_time			body string true "end time"
// @Param	limit				body int	true "limit"
// @Param	offset				body int	true "offset"
// @Produce application/json
// @Success 200 {string} string "{"queries":[{"sql_id":"F9A57DD5A41825CA","fingerprint":"select sleep(?)","example":"select sleep(3)","db_name":"","exec_count":1,"total_exec_time":3,"avg_exec_time":3,"rows_examined_max":0}]}"
// @Router	/api/v1/query/cluster [get]
func GetByMySQLClusterID(c *gin.Context) {
	var rd *utilquery.MySQLClusterRange
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
	// init server
	service := query.NewServiceWithDefault(config)
	err = service.GetByMySQLClusterID(rd.GetMySQLClusterID())
	if err != nil {
		resp.ResponseNOK(c, msgquery.ErrQueryGetByMySQLClusterID, err, rd.GetMySQLClusterID())
		return
	}
	// marshal
	jsonBytes, err := service.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgquery.DebugQueryGetByMySQLClusterID, rd.GetMySQLClusterID(), jsonStr).Error())

	// response
	resp.ResponseOK(c, jsonStr, msgquery.InfoQueryGetByMySQLClusterID, rd.GetMySQLClusterID())
}

// @Tags query
// @Summary get slow queries by mysql server id
// @Accept  application/json
// @Param	token	 		body string true "token"
// @Param	mysql_server_id	body int	true "mysql server id"
// @Param	start_time		body string true "start time"
// @Param	end_time		body string true "end time"
// @Param	limit			body int	true "limit"
// @Param	offset			body int	true "offset"
// @Produce  application/json
// @Success 200 {string} string "{"queries":[{"sql_id":"F9A57DD5A41825CA","fingerprint":"select sleep(?)","example":"select sleep(3)","db_name":"","exec_count":1,"total_exec_time":3,"avg_exec_time":3,"rows_examined_max":0}]}"
// @Router /api/v1/query/server [get]
func GetByMySQLServerID(c *gin.Context) {
	var rd *utilquery.MySQLServerRange
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
	err = service.GetByMySQLServerID(rd.GetMySQLServerID())
	if err != nil {
		resp.ResponseNOK(c, msgquery.ErrQueryGetByMySQLServerID, err, rd.GetMySQLServerID())
		return
	}

	// marshal
	jsonBytes, err := service.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgquery.DebugQueryGetByMySQLServerID, rd.GetMySQLServerID(), jsonStr).Error())

	// response
	resp.ResponseOK(c, jsonStr, msgquery.InfoQueryGetByMySQLServerID, rd.GetMySQLServerID())
}

// @Tags query
// @Summary get slow queries by mysql server host ip and port number
// @Accept  application/json
// @Param	token	 	body string true "token"
// @Param	host_ip		body string	true "mysql server host ip"
// @Param	port_num	body int	true "mysql server port number"
// @Param	start_time	body string	true "start time"
// @Param	end_time	body string	true "end time"
// @Param	limit		body int	true "limit"
// @Param	offset		body int	true "offset"
// @Produce  application/json
// @Success 200 {string} string "{"queries":[{"sql_id":"F9A57DD5A41825CA","fingerprint":"select sleep(?)","example":"select sleep(3)","db_name":"","exec_count":1,"total_exec_time":3,"avg_exec_time":3,"rows_examined_max":0}]}"
// @Router /api/v1/query/host-info [get]
func GetByHostInfo(c *gin.Context) {
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
	resp.ResponseOK(c, jsonStr, msgquery.InfoQueryGetByHostInfo, rd.GetHostIP(), rd.GetPortNum())
}

// @Tags query
// @Summary get slow queries by db id
// @Accept  application/json
// @Param	token	 		body string true "token"
// @Param	db_id			body int	true "db id"
// @Param	start_time		body string true "start time"
// @Param	end_time		body string true "end time"
// @Param	limit			body int	true "limit"
// @Param	offset			body int	true "offset"
// @Produce application/json
// @Success 200 {string} string "{"queries":[{"sql_id":"F9A57DD5A41825CA","fingerprint":"select sleep(?)","example":"select sleep(3)","db_name":"","exec_count":1,"total_exec_time":3,"avg_exec_time":3,"rows_examined_max":0}]}"
// @Router	/api/v1/query/db [get]
func GetByDBID(c *gin.Context) {
	var rd *utilquery.DBRange
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
	err = service.GetByDBID(rd.GetDBID())
	if err != nil {
		resp.ResponseNOK(c, msgquery.ErrQueryGetByDBID, err, rd.GetDBID())
		return
	}

	// marshal
	jsonBytes, err := service.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
	}
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgquery.DebugQueryGetByDBID, rd.GetDBID(), jsonStr).Error())

	// response
	resp.ResponseOK(c, jsonStr, msgquery.InfoQueryGetByDBID, rd.GetDBID())
}

// @Tags query
// @Summary get slow query by query id
// @Accept  application/json
// @Param	token	 		body string true "token"
// @Param	mysql_server_id	body int	true "mysql server id"
// @Param	sql_id			body string	true "sql id"
// @Param	start_time		body string true "start time"
// @Param	end_time		body string true "end time"
// @Param	limit			body int	true "limit"
// @Param	offset			body int	true "offset"
// @Produce application/json
// @Success 200 {string} string "{"queries":[{"sql_id":"F9A57DD5A41825CA","fingerprint":"select sleep(?)","example":"select sleep(3)","db_name":"","exec_count":1,"total_exec_time":3,"avg_exec_time":3,"rows_examined_max":0}]}"
// @Router	/api/v1/query/sql [get]
func GetBySQLID(c *gin.Context) {
	var rd *utilquery.SQLIDRange
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
	err = service.GetBySQLID(rd.GetMySQLServerID(), rd.GetSQLID())
	if err != nil {
		resp.ResponseNOK(c, msgquery.DebugQueryGetBySQLID, err, rd.GetMySQLServerID(), rd.GetSQLID())
		return
	}

	// marshal
	jsonBytes, err := service.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
	}
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgquery.DebugQueryGetBySQLID, rd.GetMySQLServerID(), rd.GetSQLID()).Error())

	// response
	resp.ResponseOK(c, jsonStr, msgquery.InfoQueryGetBySQLID, rd.GetSQLID())
}
