package metadata

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/pingcap/errors"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/pkg/message"
	msgmeta "github.com/romberli/das/pkg/message/metadata"
	"github.com/romberli/das/pkg/resp"
	utilmeta "github.com/romberli/das/pkg/util/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
)

const (
	dbIDJSON          = "id"
	dbDBNameJSON      = "db_name"
	dbClusterIDJSON   = "cluster_id"
	dbClusterTypeJSON = "cluster_type"
	dbEnvIDJSON       = "env_id"
	dbAppIDJSON       = "app_id"
	dbUserIDJSON      = "user_id"
	dbDelFlagJSON     = "del_flag"

	dbIDStruct          = "ID"
	dbDBNameStruct      = "DBName"
	dbClusterIDStruct   = "ClusterID"
	dbClusterTypeStruct = "ClusterType"
	dbEnvIDStruct       = "EnvID"
	dbDelFlagStruct     = "DelFlag"

	dbMySQLClusterStruct = "MySQLCluster"
	dbAppsStruct         = "Apps"
	dbUsersStruct        = "Users"
)

// @Tags    database
// @Summary get all databases
// @Accept	application/json
// @Param	token body string true "token"
// @Produce application/json
// @Success 200 {string} string "{"dbs": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router  /api/v1/metadata/db [get]
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

// @Tags    database
// @Summary get database by env_id
// @Accept	application/json
// @Param	token  body string true "token"
// @Param	env_id body int    true "env id"
// @Produce application/json
// @Success 200 {string} string "{"dbs": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router  /api/v1/metadata/db/env [get]
func GetDBByEnv(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	envID, err := jsonparser.GetInt(data, dbEnvIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), dbEnvIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetByEnv(int(envID))
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

// @Tags    database
// @Summary get database by id
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id    body int    true "db id"
// @Produce application/json
// @Success 200 {string} string "{"dbs": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router  /api/v1/metadata/db/get [get]
func GetDBByID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, dbIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), dbIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetByID(int(id))
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

// @Tags    database
// @Summary get database by db name and cluster info
// @Accept	application/json
// @Param	token        body string true "token"
// @Param	db_name	     body string true "db name"
// @Param 	cluster_id   body int    true "cluster id"
// @Param 	cluster_type body int	 true "cluster type"
// @Produce application/json
// @Success 200 {string} string "{"dbs": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router  /api/v1/metadata/db/name-and-cluster-info [get]
func GetDBByNameAndClusterInfo(c *gin.Context) {
	var rd *utilmeta.NameAndClusterInfo
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetDBByNameAndClusterInfo(rd.GetDBName(), rd.GetClusterID(), rd.GetClusterType())
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBByNameAndClusterInfo, err, rd.GetDBName(), rd.GetClusterID(), rd.GetClusterType())
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
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBByNameAndClusterInfo, rd.GetDBName(), rd.GetClusterID(), rd.GetClusterType())
}

// @Tags    database
// @Summary get database by db name and host info
// @Accept	application/json
// @Param	token    body string true "token"
// @Param	db_name	 body string true "db name"
// @Param 	host_ip  body string true "host_ip"
// @Param 	port_num body int	 true "port_num"
// @Produce application/json
// @Success 200 {string} string "{"dbs": [{"id": 1, "db_name": "pmm_test", "cluster_id": 1, "cluster_type": 1, "env_id": 1, "del_flag": 0, "create_time": "2022-03-14T09:59:21.379851+08:00", "last_update_time": "2022-03-14T09:59:21.379851+08:00"}]}"
// @Router  /api/v1/metadata/db/name-and-host-info [get]
func GetDBByNameAndHostInfo(c *gin.Context) {
	var rd *utilmeta.NameAndHostInfo
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetDBByNameAndHostInfo(rd.GetDBName(), rd.GetHostIP(), rd.GetPortNum())
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBByNameAndHostInfo, err, rd.GetDBName(), rd.GetHostIP(), rd.GetPortNum())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDBByNameAndHostInfo, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBByNameAndHostInfo, rd.GetDBName(), rd.GetHostIP(), rd.GetPortNum())
}

// @Tags    database
// @Summary get databases by host info
// @Accept	application/json
// @Param	token    body string true "token"
// @Param 	host_ip  body string true "host_ip"
// @Param 	port_num body int	 true "port_num"
// @Produce application/json
// @Success 200 {string} string "{"dbs": [{"id": 1, "db_name": "pmm_test", "cluster_id": 1, "cluster_type": 1, "env_id": 1, "del_flag": 0, "create_time": "2022-03-14T09:59:21.379851+08:00", "last_update_time": "2022-03-14T09:59:21.379851+08:00"}]}"
// @Router  /api/v1/metadata/db/host-info [get]
func GetDBsByHostInfo(c *gin.Context) {
	var rd *utilmeta.HostInfo
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetDBsByHostInfo(rd.GetHostIP(), rd.GetPortNum())
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBsByHostInfo, err, rd.GetHostIP(), rd.GetPortNum())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDBsByHostInfo, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBsByHostInfo, rd.GetHostIP(), rd.GetPortNum())
}

// @Tags    database
// @Summary get apps by id
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id    body int    true "db id"
// @Produce application/json
// @Success 200 {string} string "{"apps": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8, "owner_group": "k"}]}"
// @Router  /api/v1/metadata/db/app [get]
func GetAppsByDBID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, dbIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), dbIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetAppsByDBID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAppsByDBID, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAppsByDBID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAppsByDBID, id)
}

// @Tags    database
// @Summary get mysql cluster by id
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id    body int    true "db id"
// @Produce application/json
// @Success 200 {string} string "{"mysql_cluster": [{"middleware_cluster_id":1,"monitor_system_id":1,"env_id":1,"del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","last_update_time":"2021-02-23T20:57:24.603009+08:00","id":1,"cluster_name":"cluster_name_init"},{"monitor_system_id":1, "env_id":1,"create_time":"2021-02-23T04:14:23.707238+08:00","last_update_time":"2021-02-23T04:14:23.707238+08:00","id":2,"cluster_name":"newTest","middleware_cluster_id":1,"del_flag":0}]}"
// @Router  /api/v1/metadata/db/mysql-cluster [get]
func GetMySQLClusterByDBID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, dbIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), dbIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetMySQLClusterByID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLClusterByDBID, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLClusterByDBID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLClusterByDBID, id)
}

// @Tags    database
// @Summary get app users
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id    body int    true "db id"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"department_name": "dn","account_name": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router  /api/v1/metadata/db/app-user [get]
func GetAppUsersByDBID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, dbIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), dbIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetAppUsersByDBID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAppUsersByDBID, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAppUsersByDBID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAppUsersByDBID, id)
}

// @Tags    database
// @Summary get db users
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id    body int    true "db id"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"department_name": "dn","account_name": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router  /api/v1/metadata/db/db-user [get]
func GetUsersByDBID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, dbIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), dbIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetUsersByDBID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetUsersByDBID, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetUsersByDBID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetUsersByDBID, id)
}

// @Tags    database
// @Summary get all users
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id    body int    true "db id"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"department_name": "dn","account_name": "da", "mobile": "m", "del_flag": 0,"last_update_time": "2021-01-21T13:00:00+08:00","user_name": "un","create_time": "2021-01-21T13:00:00+08:00","employee_id": 1,"email": "e","telephone": "t","role": 1, "id": 1}]}"
// @Router  /api/v1/metadata/db/all-user [get]
func GetAllUsersByDBID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, dbIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), dbIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// get entity
	err = s.GetAllUsersByDBID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAllUsersByDBID, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAllUsersByDBID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAllUsersByDBID, id)
}

// @Tags    database
// @Summary add a new database
// @Accept	application/json
// @Param	token        body string true "token"
// @Param	db_name	     body string true	"db name"
// @Param 	cluster_id   body int    true	"cluster id"
// @Param 	cluster_type body int	 false	"cluster type"
// @Param 	env_id       body int    true	"env id"
// @Produce application/json
// @Success 200 {string} string "{"dbs": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router  /api/v1/metadata/db [post]
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
			dbDBNameJSON, dbClusterIDJSON, dbClusterTypeJSON, dbEnvIDJSON))
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

// @Tags    database
// @Summary update database by id
// @Accept	application/json
// @Param	token        body string true "token"
// @Param	id		     body int	 true	"db id"
// @Param	db_name	     body string false	"db name"
// @Param 	cluster_id   body int    false	"cluster id"
// @Param 	cluster_type body int	 false	"cluster type"
// @Param 	env_id       body int    false	"env id"
// @Param 	del_flag     body int    false	"delete flag"
// @Produce application/json
// @Success 200 {string} string "{"dbs": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router  /api/v1/metadata/db/update [post]
func UpdateDBByID(c *gin.Context) {
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
	idInterface, idExists := fields[dbIDStruct]
	if !idExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
		return
	}
	id, ok := idInterface.(int)
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, dbIDJSON)
		return
	}
	_, dbNameExists := fields[dbDBNameStruct]
	_, clusterIDExists := fields[dbClusterIDStruct]
	_, clusterTypeExists := fields[dbClusterTypeStruct]
	_, envIDExists := fields[dbEnvIDStruct]
	_, delFlagExists := fields[dbDelFlagStruct]
	if !dbNameExists && !clusterIDExists && !clusterTypeExists && !envIDExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists,
			fmt.Sprintf("%s, %s, %s, %s, %s",
				dbDBNameJSON, dbClusterIDJSON, dbClusterTypeJSON, dbEnvIDJSON, dbDelFlagJSON))
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

// @Tags    database
// @Summary delete database by id
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id    body int	  true "db id"
// @Produce application/json
// @Success 200 {string} string "{"dbs": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router  /api/v1/metadata/db/delete [post]
func DeleteDBByID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, dbIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), dbIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// update entity
	err = s.Delete(int(id))
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

// @Tags    database
// @Summary add application map
// @Param	token  body string true "token"
// @Param	id     body int	   true "db id"
// @Param	app_id body int	   true "app id"
// @Produce application/json
// @Success 200 {string} string "{"apps": [{"create_time":"2021-11-10T18:39:12.395612+08:00","last_update_time":"2021-12-21T09:15:47.688546+08:00","id":1,"app_name":"app1","level":1,"del_flag":0},{"last_update_time":"2021-12-21T09:15:47.688546+08:00","id":3,"app_name":"app3","level":3,"del_flag":0,"create_time":"2021-11-02T18:02:34.153234+08:00"}]}"
// @Router  /api/v1/metadata/db/add-app [post]
func DBAddApp(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, dbIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), dbIDJSON)
		return
	}
	appID, err := jsonparser.GetInt(data, dbAppIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), dbAppIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// update entities
	err = s.AddApp(int(id), int(appID))
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

// @Tags    database
// @Summary delete application map
// @Param	token  body string true "token"
// @Param	id     body int	   true "db id"
// @Param	app_id body int	   true "app id"
// @Produce application/json
// @Success 200 {string} string "{"apps": [{"last_update_time":"2021-12-21T09:15:47.688546+08:00","id":1,"app_name":"app1","level":1,"del_flag":0,"create_time":"2021-11-10T18:39:12.395612+08:00"}]}"
// @Router  /api/v1/metadata/db/delete-app [post]
func DBDeleteApp(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, dbIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), dbIDJSON)
		return
	}
	appID, err := jsonparser.GetInt(data, dbAppIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), dbAppIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// update entities
	err = s.DeleteApp(int(id), int(appID))
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

// @Tags    database
// @Summary add user map
// @Param	token   body string true "token"
// @Param	id      body int    true "db id"
// @Param	user_id body int    true "user id"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"employee_id":"","telephone":"","create_time":"2022-03-01T17:53:21.046511+08:00","last_update_time":"2022-03-01T17:53:21.046511+08:00","mobile":"","role":3,"id":2,"user_name":"test","department_name":"","account_name":"aaaa","email":"qqqq","del_flag":0},{"role":3,"create_time":"2022-01-25T12:21:05.19953+08:00","user_name":"test1","employee_id":"","account_name":"aaa","email":"aaa","telephone":"","mobile":"","last_update_time":"2022-01-25T12:21:05.19953+08:00","id":3,"department_name":"","del_flag":0}]}"
// @Router  /api/v1/metadata/db/add-user [post]
func DBAddUser(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, dbIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), dbIDJSON)
		return
	}
	userID, err := jsonparser.GetInt(data, dbUserIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), dbUserIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// update entities
	err = s.DBAddUser(int(id), int(userID))
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

// @Tags    database
// @Summary delete user map
// @Param	id      path int    true "db id"
// @Param	token   body string true "token"
// @Param	user_id body int    true "user id"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"id":2,"employee_id":"","role":3,"del_flag":0,"create_time":"2022-03-01T17:53:21.046511+08:00","last_update_time":"2022-03-01T17:53:21.046511+08:00","user_name":"test","department_name":"","account_name":"aaaa","email":"qqqq","telephone":"","mobile":""}]}"
// @Router  /api/v1/metadata/db/delete-user/:id [post]
func DBDeleteUser(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, dbIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), dbIDJSON)
		return
	}
	userID, err := jsonparser.GetInt(data, dbUserIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), dbUserIDJSON)
		return
	}
	// init service
	s := metadata.NewDBServiceWithDefault()
	// update entities
	err = s.DBDeleteUser(int(id), int(userID))
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
