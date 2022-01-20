package metadata

import (
	"encoding/json"
	"fmt"
	"github.com/pingcap/errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"

	"github.com/romberli/das/internal/app/metadata"
	msgmeta "github.com/romberli/das/pkg/message/metadata"

	"github.com/romberli/das/pkg/message"
	"github.com/romberli/das/pkg/resp"
)

const (
	dbIDJSON     = "id"
	dbEnvIDJSON  = "env_id"
	dbAppIDJSON  = "app_id"
	dbUserIDJSON = "user_id"

	dbDBNameStruct      = "DBName"
	dbClusterIDStruct   = "ClusterID"
	dbClusterTypeStruct = "ClusterType"
	dbEnvIDStruct       = "EnvID"

	dbMySQLClusterStruct = "MySQLCluster"
	dbAppsStruct         = "Apps"
	dbUsersStruct        = "Users"
	dbAppOwnersStruct    = "AppOwners"
	dbDBOwnersStruct     = "DBOwners"
	dbAllOwnersStruct    = "AllOwners"
)

// @Tags database
// @Summary get all databases
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db [get]
func GetDB(c *gin.Context) {
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBAll, err)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDBAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBAll)
}

// @Tags database
// @Summary get database by env_id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db/env/:env_id [get]
func GetDBByEnv(c *gin.Context) {
	// get param
	envIDStr := c.Param(dbEnvIDJSON)
	if envIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbEnvIDJSON)
		return
	}
	envID, err := strconv.Atoi(envIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return

	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetByEnv(envID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBByEnv, err)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDBByEnv, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBByEnv, envID)
}

// @Tags database
// @Summary get database by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db/get/:id [get]
func GetDBByID(c *gin.Context) {
	// get param
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBByID, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDBByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBByID, id)
}

// @Tags database
// @Summary get database by db name and cluster info
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db/name-and-cluster-info [get]
func GetDBByNameAndClusterInfo(c *gin.Context) {
	var dbInfo *metadata.DBInfo
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	err = json.Unmarshal(data, dbInfo)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetByNameAndClusterInfo(dbInfo.DBName, dbInfo.ClusterID, dbInfo.ClusterType)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBByNameAndClusterInfo, err, dbInfo.DBName, dbInfo.ClusterID, dbInfo.ClusterType)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDBByNameAndClusterInfo, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBByNameAndClusterInfo, dbInfo.DBName, dbInfo.ClusterID, dbInfo.ClusterType)
}

// @Tags db
// @Summary get apps by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8, "owner_group": "k"}]}"
// @Router /api/v1/metadata/db/app/:id [get]
func GetAppsByDBID(c *gin.Context) {
	// get param
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetAppsByDBID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAppsByID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(dbAppsStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAppsByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAppsByID, id)

}

// @Tags db
// @Summary get mysql cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"middleware_cluster_id":1,"monitor_system_id":1,"env_id":1,"del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","last_update_time":"2021-02-23T20:57:24.603009+08:00","id":1,"cluster_name":"cluster_name_init"},{"monitor_system_id":1, "env_id":1,"create_time":"2021-02-23T04:14:23.707238+08:00","last_update_time":"2021-02-23T04:14:23.707238+08:00","id":2,"cluster_name":"newTest","middleware_cluster_id":1,"del_flag":0}]}"
// @Router /api/v1/metadata/db/mysql-cluster/:id [get]
func GetMySQLClusterByDBID(c *gin.Context) {
	// get param
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetMySQLClusterByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLClusterByID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(dbMySQLClusterStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLClusterByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLClusterByID, id)
}

// @Tags db
// @Summary get app owners
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/db/app-owner/:id [get]
func GetAppUsersByDBID(c *gin.Context) {
	// get param
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetAppUsersByDBID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAppOwners, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(dbAppOwnersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAppOwners, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAppOwners, id)
}

// @Tags db
// @Summary get db owners
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/db/db-owner/:id [get]
func GetUsersByDBID(c *gin.Context) {
	// get param
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetUsersByDBID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBOwners, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(dbDBOwnersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDBOwners, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBOwners, id)
}

// @Tags db
// @Summary get all owners
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"department_name": "dn","accountNameStruct = "AccountName"": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router /api/v1/metadata/db/all-owner/:id [get]
func GetAllUsersByDBID(c *gin.Context) {
	// get param
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetAllUsersByDBID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAllOwners, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(dbAllOwnersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAllOwners, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAllOwners, id)
}

// @Tags database
// @Summary add a new database
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db [post]
func AddDB(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.DBInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, dbNameExists := fields[dbDBNameStruct]
	_, clusterIDExists := fields[dbClusterIDStruct]
	_, clusterTypeExists := fields[dbClusterTypeStruct]
	_, envIDExists := fields[dbEnvIDStruct]
	if !dbNameExists || !clusterIDExists || !clusterTypeExists || !envIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s and %s and %s",
			dbDBNameStruct, dbClusterIDStruct, dbClusterTypeStruct, dbEnvIDStruct))
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddDB, err,
			fields[dbDBNameStruct], fields[dbClusterIDStruct], fields[dbClusterTypeStruct], fields[dbEnvIDStruct])
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddDB, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddDB, fields[dbDBNameStruct], fields[dbClusterIDStruct], fields[dbClusterTypeStruct], fields[dbEnvIDStruct])
}

// @Tags database
// @Summary update database by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db/update/:id [post]
func UpdateDBByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.DBInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, dbNameExists := fields[dbDBNameStruct]
	_, clusterIDExists := fields[dbClusterIDStruct]
	_, clusterTypeExists := fields[dbClusterTypeStruct]
	_, envIDExists := fields[dbEnvIDStruct]
	_, delFlagExists := fields[envDelFlagStruct]
	if !dbNameExists && !clusterIDExists && !clusterTypeExists && !envIDExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists,
			fmt.Sprintf("%s, %s, %s, %s, %s",
				dbDBNameStruct, dbClusterIDStruct, dbClusterTypeStruct, dbEnvIDStruct, envDelFlagStruct))
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateDB, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateDB, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataUpdateDB, id)
}

// @Tags database
// @Summary delete database by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db/delete/:id [post]
func DeleteDBByID(c *gin.Context) {
	// get params
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// update entity
	err = s.Delete(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteDB, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteDB, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteDB, id)
}

// @Tags database
// @Summary add application map
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [1, 2]}"
// @Router /api/v1/metadata/db/add-app/:id [post]
func DBAddApp(c *gin.Context) {
	// get params
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
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
		resp.ResponseNOK(c, msgmeta.ErrMetadataDBAddApp, errors.Trace(err), id)
		return
	}
	appID, appIDExists := dataMap[dbAppIDJSON]
	if !appIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbAppIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// update entities
	err = s.AddApp(id, appID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDBAddApp, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(dbAppsStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDBAddApp, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDBAddApp, id, appID)
}

// @Tags database
// @Summary delete application map
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [1]}"
// @Router /api/v1/metadata/db/delete-app/:id [post]
func DBDeleteApp(c *gin.Context) {
	// get params
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
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
		resp.ResponseNOK(c, msgmeta.ErrMetadataDBDeleteApp, errors.Trace(err), id)
		return
	}
	appID, appIDExists := dataMap[dbAppIDJSON]
	if !appIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbAppIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// update entities
	err = s.DeleteApp(id, appID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDBDeleteApp, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(dbAppsStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDBAddApp, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDBDeleteApp, id, appID)
}

// @Tags database
// @Summary add user map
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [1, 2]}"
// @Router /api/v1/metadata/db/add-user/:id [post]
func DBAddUser(c *gin.Context) {
	// get params
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
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
		resp.ResponseNOK(c, msgmeta.ErrMetadataDBAddUser, errors.Trace(err), id)
		return
	}
	userID, userIDExists := dataMap[dbUserIDJSON]
	if !userIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbUserIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// update entities
	err = s.DBAddUser(id, userID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDBAddUser, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(dbUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDBAddUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDBAddUser, id, userID)
}

// @Tags database
// @Summary delete user map
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [1]}"
// @Router /api/v1/metadata/db/delete-user/:id [post]
func DBDeleteUser(c *gin.Context) {
	// get params
	idStr := c.Param(dbIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
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
		resp.ResponseNOK(c, msgmeta.ErrMetadataDBDeleteUser, errors.Trace(err), id)
		return
	}
	userID, userIDExists := dataMap[dbUserIDJSON]
	if !userIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbUserIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// update entities
	err = s.DBDeleteUser(id, userID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDBDeleteUser, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(dbUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDBDeleteUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDBDeleteUser, id, userID)
}
