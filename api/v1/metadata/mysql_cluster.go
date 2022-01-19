package metadata

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pingcap/errors"

	"github.com/gin-gonic/gin"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"

	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/pkg/message"
	msgmeta "github.com/romberli/das/pkg/message/metadata"
	"github.com/romberli/das/pkg/resp"
)

const (
	mysqlClusterIDJSON          = "id"
	mysqlClusterEnvIDJSON       = "env_id"
	mysqlClusterClusterNameJSON = "name"
	mysqlClusterUserIDJSON      = "user_id"

	mysqlClusterClusterNameStruct         = "ClusterName"
	mysqlClusterMiddlewareClusterIDStruct = "MiddlewareClusterID"
	mysqlClusterMonitorSystemIDStruct     = "MonitorSystemID"
	mysqlClusterEnvIDStruct               = "EnvID"

	mysqlClusterMySQLServersStruct = "MySQLServers"
	mysqlClusterDBsStruct          = "DBs"
	mysqlClusterUsersStruct        = "Users"
)

// @Tags mysql cluster
// @Summary get all mysql clusters
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"middleware_cluster_id":1,"monitor_system_id":1,"env_id":1,"del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","last_update_time":"2021-02-23T20:57:24.603009+08:00","id":1,"cluster_name":"cluster_name_init","owner_id":1},{"monitor_system_id":1,"owner_id":1,"env_id":1,"create_time":"2021-02-23T04:14:23.707238+08:00","last_update_time":"2021-02-23T04:14:23.707238+08:00","id":2,"cluster_name":"newTest","middleware_cluster_id":1,"del_flag":0}]}"
// @Router /api/v1/metadata/mysql-cluster [get]
func GetMySQLCluster(c *gin.Context) {
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLClusterAll, err)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	fmt.Println("ok")
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLClusterAll, jsonBytes).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLClusterAll)
}

// @Tags mysql cluster
// @Summary get mysql cluster by env id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"owner_id":1,"del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","id":1,"monitor_system_id":1,"env_id":1,"last_update_time":"2021-02-23T20:57:24.603009+08:00","cluster_name":"cluster_name_init","middleware_cluster_id":1}]}"
// @Router /api/v1/metadata/mysql-cluster/:id [get]
func GetMySQLClusterByEnv(c *gin.Context) {
	// get param
	envIDStr := c.Param(mysqlClusterEnvIDJSON)
	if envIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterEnvIDJSON)
		return
	}
	envID, err := strconv.Atoi(envIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return

	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entity
	err = s.GetByEnv(envID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLClusterByEnv, err, envID)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLClusterByEnv, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLClusterByEnv, envID)
}

// @Tags mysql cluster
// @Summary get mysql cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"owner_id":1,"del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","id":1,"monitor_system_id":1,"env_id":1,"last_update_time":"2021-02-23T20:57:24.603009+08:00","cluster_name":"cluster_name_init","middleware_cluster_id":1}]}"
// @Router /api/v1/metadata/mysql-cluster/:id [get]
func GetMySQLClusterByID(c *gin.Context) {
	// get param
	idStr := c.Param(mysqlClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return

	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entity
	err = s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLClusterByID, id, err)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLClusterByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLClusterByID, id)
}

// @Tags mysql cluster
// @Summary get mysql cluster by name
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"owner_id":1,"del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","id":1,"monitor_system_id":1,"env_id":1,"last_update_time":"2021-02-23T20:57:24.603009+08:00","cluster_name":"cluster_name_init","middleware_cluster_id":1}]}"
// @Router /api/v1/metadata/mysql-cluster/cluster-name/:name [get]
func GetMySQLClusterByName(c *gin.Context) {
	// get param
	clusterName := c.Param(mysqlClusterClusterNameJSON)
	if clusterName == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterClusterNameJSON)
		return
	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entity
	err := s.GetByName(clusterName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLClusterByName, err, clusterName)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLClusterByName, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLClusterByName, clusterName)
}

// @Tags mysql cluster
// @Summary get mysql servers by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"owner_id":1,"del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","id":1,"monitor_system_id":1,"env_id":1,"last_update_time":"2021-02-23T20:57:24.603009+08:00","cluster_name":"cluster_name_init","middleware_cluster_id":1}]}"
// @Router /api/v1/metadata/mysql-cluster/mysql-server/:id [get]
func GetMySQLServersByID(c *gin.Context) {
	// get params
	idStr := c.Param(mysqlClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterIDJSON)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entity
	err = s.GetMySQLServersByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLServers, id, err)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(mysqlClusterMySQLServersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)

	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLServers, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLServers, id)
}

// @Tags mysql cluster
// @Summary get master servers by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"owner_id":1,"del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","id":1,"monitor_system_id":1,"env_id":1,"last_update_time":"2021-02-23T20:57:24.603009+08:00","cluster_name":"cluster_name_init","middleware_cluster_id":1}]}"
// @Router /api/v1/metadata/mysql-cluster/master-server/:id [get]
func GetMasterServersByID(c *gin.Context) {
	// get params
	idStr := c.Param(mysqlClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterIDJSON)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entity
	err = s.GetMasterServersByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMasterServers, id, err)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(mysqlClusterMySQLServersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMasterServers, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMasterServers, id)
}

// @Tags mysql cluster
// @Summary get dbs by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/mysql-cluster/db/:id [get]
func GetDBsByMySQLCLusterID(c *gin.Context) {
	// get params
	idStr := c.Param(mysqlClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterIDJSON)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entity
	err = s.GetDBsByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBs, id, err)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(mysqlClusterDBsStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)

	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDBs, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBs, id)
}

// @Tags mysql cluster
// @Summary get mysql cluster users
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/mysql-cluster/user/:id [get]
func GetUsersByMySQLClusterID(c *gin.Context) {
	// get params
	idStr := c.Param(mysqlClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterIDJSON)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entity
	err = s.GetUsersByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAppUsers, id, err)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(mysqlClusterUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAppUsers, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAppUsers, id)
}

// @Tags mysql cluster
// @Summary add user map
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"user_id": 1}]}"
// @Router /api/v1/metadata/mysql-cluster/add-user/:id [post]
func MySQLClusterAddUser(c *gin.Context) {
	// get params
	idStr := c.Param(mysqlClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}

	dataMap := make(map[string]int)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataMySQLClusterAddUser, id, errors.Trace(err))
		return
	}
	userID, userIDExists := dataMap[mysqlClusterUserIDJSON]
	if !userIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterUserIDJSON)
		return
	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// update entities
	err = s.AddUser(id, userID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataMySQLClusterAddUser, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(mysqlClusterUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, errors.Trace(err))
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataMySQLClusterAddUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataMySQLClusterAddUser, id, userID)
}

// @Tags mysql cluster
// @Summary delete user map
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"user_id": 1}]}"
// @Router /api/v1/metadata/mysql-cluster/delete-user/:id [post]
func MySQLClusterDeleteUser(c *gin.Context) {
	// get params
	idStr := c.Param(mysqlClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	dataMap := make(map[string]int)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataMySQLClusterDeleteUser, errors.Trace(err), id)
		return
	}
	userID, userIDExists := dataMap[mysqlClusterUserIDJSON]
	if !userIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterUserIDJSON)
		return
	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// update entities
	err = s.DeleteUser(id, userID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataMySQLClusterDeleteUser, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(mysqlClusterUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataMySQLClusterDeleteUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataMySQLClusterDeleteUser, id, userID)
}

// @Tags mysql cluster
// @Summary get app users
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/mysql-cluster/app-user/:id [get]
func GetAppUsersByMySQLCLusterID(c *gin.Context) {
	// get params
	idStr := c.Param(mysqlClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterIDJSON)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entity
	err = s.GetAppUsersByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAppUsers, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(mysqlClusterUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAppUsers, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAppUsers, id)
}

// @Tags mysql cluster
// @Summary get db users
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/mysql-cluster/db-user/:id [get]
func GetDBUsersByMySQLCLusterID(c *gin.Context) {
	// get params
	idStr := c.Param(mysqlClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterIDJSON)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entity
	err = s.GetDBUsersByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBUsers, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(mysqlClusterUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDBUsers, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBUsers, id)
}

// @Tags mysql cluster
// @Summary get all users
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/mysql-cluster/all-user/:id [get]
func GetAllUsersByMySQLCLusterID(c *gin.Context) {
	// get params
	idStr := c.Param(mysqlClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterIDJSON)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// get entity
	err = s.GetAllUsersByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAllUsers, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(mysqlClusterUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAllUsers, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAllUsers, id)
}

// @Tags mysql cluster
// @Summary add a new mysql cluster
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"cluster_name":"api_test","monitor_system_id":0,"owner_group":"","del_flag":0,"create_time":"2021-02-24T02:33:50.936279+08:00","last_update_time":"2021-02-24T02:33:50.936279+08:00","middleware_cluster_id":0,"owner_id":0,"env_id":0,"id":154}]}"
// @Router /api/v1/metadata/mysql-cluster [post]
func AddMySQLCluster(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MySQLClusterInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	if _, ok := fields[mysqlClusterClusterNameStruct]; !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterClusterNameStruct)
		return
	}
	if _, ok := fields[mysqlClusterEnvIDStruct]; !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterEnvIDStruct)
		return
	}

	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddMySQLCluster,
			fields[mysqlClusterClusterNameStruct],
			fields[mysqlClusterEnvIDStruct],
			err)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddMySQLCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddMySQLCluster,
		fields[mysqlClusterClusterNameStruct],
		fields[mysqlClusterEnvIDStruct],
	)
}

// @Tags mysql cluster
// @Summary update mysql cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":154,"middleware_cluster_id":0,"owner_id":0,"env_id":0,"create_time":"2021-02-24T02:33:50.936279+08:00","cluster_name":"api_test","monitor_system_id":0,"owner_group":"","del_flag":1,"last_update_time":"2021-02-24T02:33:50.936279+08:00"}]}"
// @Router /api/v1/metadata/mysql-cluster/:id [post]
func UpdateMySQLClusterByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(mysqlClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return

	}
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MySQLClusterInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, clusterNameExists := fields[mysqlClusterClusterNameStruct]
	_, middlewareClusterIDExists := fields[mysqlClusterMiddlewareClusterIDStruct]
	_, monitorSystemIDExists := fields[mysqlClusterMonitorSystemIDStruct]
	// _, ownerIDExists := fields[mysqlClusterUserIDStruct]
	_, envIDExists := fields[mysqlClusterEnvIDStruct]
	_, delFlagExists := fields[envDelFlagStruct]
	if !clusterNameExists &&
		!middlewareClusterIDExists &&
		!monitorSystemIDExists &&
		// !ownerIDExists &&
		!envIDExists &&
		!delFlagExists {
		resp.ResponseNOK(
			c, message.ErrFieldNotExists,
			fmt.Sprintf("%s, %s, %s, %s and %s",
				mysqlClusterClusterNameStruct,
				mysqlClusterMiddlewareClusterIDStruct,
				mysqlClusterMonitorSystemIDStruct,
				// mysqlClusterOwnerIDStruct,
				mysqlClusterEnvIDStruct,
				envDelFlagStruct))
		return
	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateMySQLCluster, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// resp
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateMySQLCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.DebugMetadataUpdateMySQLCluster, fields[mysqlClusterClusterNameStruct])
}

func DeleteMySQLClusterByID(c *gin.Context) {
	var fields map[string]interface{}

	// get param
	idStr := c.Param(mysqlClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlClusterIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return

	}
	// init service
	s := metadata.NewMySQLClusterServiceWithDefault()
	// insert into middleware
	err = s.Delete(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteMySQLCluster,
			err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, errors.Trace(err))
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteMySQLCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteMySQLCluster, fields[mysqlClusterClusterNameStruct])
}
