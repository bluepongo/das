package query

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/romberli/das/internal/app/query"
	"github.com/romberli/das/pkg/message"
	msgquery "github.com/romberli/das/pkg/message/query"
	"github.com/romberli/das/pkg/resp"
	util "github.com/romberli/das/pkg/util/query"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
)

const (
	mysqlClusterIDJSON = "mysql_cluster_id"
	mysqlServerIDJSON  = "mysql_server_id"
	dbIDJSON           = "db_id"
	sqlIDJSON          = "sql_id"
)

// @Tags query
// @Summary get slow queries by mysql server id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": []}"
// @Router /api/v1/query/cluster/:mysqlClusterID [get]
func GetByMySQLClusterID(c *gin.Context) {

}

// @Tags query
// @Summary get slow queries by mysql server id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": []}"
// @Router /api/v1/query/server/:mysqlServerID [get]
func GetByMySQLServerID(c *gin.Context) {
	// get data
	mysqlServerIDStr := c.Param(mysqlServerIDJSON)
	if mysqlServerIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlServerIDJSON)
		return
	}
	mysqlServerID, err := strconv.Atoi(mysqlServerIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err)
		return
	}

	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	dataMap := make(map[string]string)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}

	// get config
	config, err := util.GetConfig(dataMap)

	// init service
	service := query.NewServiceWithDefault(config)
	err = service.GetByMySQLServerID(mysqlServerID)
	if err != nil {
		resp.ResponseNOK(c, msgquery.ErrQueryGetByMySQLServerID, mysqlServerID, err.Error())
		return
	}

	// marshal
	jsonBytes, err := service.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
	}
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgquery.DebugQueryGetByMySQLServerID, mysqlServerID, jsonStr).Error())

	// response
	resp.ResponseOK(c, jsonStr, msgquery.InfoQueryGetByMySQLServerID, mysqlServerID)
}

// @Tags query
// @Summary get slow queries by db id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": []}"
// @Router /api/v1/query/db/:dbID [get]
func GetByDBID(c *gin.Context) {

}

// @Tags query
// @Summary get slow query by query id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": []}"
// @Router /api/v1/query/:sqlID [get]
func GetBySQLID(c *gin.Context) {

}
