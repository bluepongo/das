package metadata

import (
	"encoding/json"
	"fmt"
	"strconv"

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
	middlewareClusterIDJSON          = "id"
	middlewareClusterClusterNameJSON = "cluster_name"
	middlewareClusterEnvIDJSON       = "env_id"
	middlewareClusterUserIDJSON      = "user_id"

	middlewareClusterClusterNameStruct       = "ClusterName"
	middlewareClusterEnvIDStruct             = "EnvID"
	middlewareClusterMiddlewareServersStruct = "MiddlewareServers"
	middlewareClusterUsersStruct             = "Users"
	middlewareClusterClusterIDStruct         = "ClusterID"
	middlewareClusterUserIDStruct            = "UserID"
)

// @Tags middleware cluster
// @Summary get all middleware clusters
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":13,"cluster_name":"test001","env_id":1,"del_flag":0,"create_time":"2021-04-09T10:55:43.920406+08:00","last_update_time":"2021-04-09T10:55:43.920406+08:00"}]}"
// @Router /api/v1/metadata/middleware-cluster [get]
func GetMiddlewareCluster(c *gin.Context) {
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareClusterAll, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareClusterAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareClusterAll)
}

// @Tags middleware cluster
// @Summary get middleware cluster by env
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"del_flag":0,"create_time":"2021-04-09T10:55:43.920406+08:00","last_update_time":"2021-04-09T10:55:43.920406+08:00","id":13,"cluster_name":"test001","env_id":1}]}"
// @Router /api/v1/metadata/middleware-cluster/env/:env_id [get]
func GetMiddlewareClusterByEnv(c *gin.Context) {
	// get param
	envIDStr := c.Param(middlewareClusterEnvIDJSON)
	if envIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterIDJSON)
		return
	}
	envID, err := strconv.Atoi(envIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entity
	err = s.GetByEnv(envID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareClusterByEnv, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareClusterByEnv, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareClusterByEnv, envID)
}

// @Tags middleware cluster
// @Summary get middleware cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":13,"cluster_name":"test001","env_id":1,"del_flag":0,"create_time":"2021-04-09T10:55:43.920406+08:00","last_update_time":"2021-04-09T10:55:43.920406+08:00"}]}"
// @Router /api/v1/metadata/middleware-cluster/get/:id [get]
func GetMiddlewareClusterByID(c *gin.Context) {
	// get param
	idStr := c.Param(middlewareClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entity
	err = s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareClusterByID, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareClusterByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareClusterByID, id)
}

// @Tags middleware cluster
// @Summary get middleware cluster by name
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":13,"cluster_name":"test001", "env_id":1,"del_flag":0,"create_time":"2021-04-09T10:55:43.920406+08:00","last_update_time":"2021-04-09T10:55:43.920406+08:00"}]}"
// @Router /api/v1/metadata/middleware-cluster/cluster-name/:cluster_name [get]
func GetMiddlewareClusterByName(c *gin.Context) {
	// get params
	middlewareClusterName := c.Param(middlewareClusterClusterNameJSON)
	if middlewareClusterName == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterClusterNameJSON)
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entity
	err := s.GetByName(middlewareClusterName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareClusterByName, middlewareClusterName, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareClusterByName, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareClusterByName, middlewareClusterName)
}

// @Tags application
// @Summary get middleware servers by cluster id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [1,2]}"
// @Router /api/vi/metadata/middleware-server/:id [get]
func GetMiddlewareServers(c *gin.Context) {
	// get params
	idStr := c.Param(middlewareClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterIDJSON)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entity
	err = s.GetMiddlewareServersByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareServers, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(middlewareClusterMiddlewareServersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareServers, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareServers, id)

}

// @Tags application
// @Summary get middleware servers by cluster id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"users":[{"id":1,"department_name":"arch","email":"allinemailtest@163.com","mobile":"13012345678","last_update_time":"2021-11-22T13:46:20.430926+08:00","create_time":"2021-10-25T09:21:50.364327+08:00","user_name":"zhangsan","employee_id":"100001","account_name":"zs001","telephone":"01012345678","role":3,"del_flag":0}]}]}"
// @Router /api/vi/metadata/middleware-cluster/users/:id [get]
func GetUsersByMiddlewareClusterID(c *gin.Context) {
	// get params
	idStr := c.Param(middlewareClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterIDJSON)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()

	// get entity
	err = s.GetUsersByMiddlewareClusterID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetUsersByMiddlewareClusterID, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(middlewareClusterUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetUsersByMiddlewareClusterID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetUsersByMiddlewareClusterID, id)
}

// @Tags middleware cluster
// @Summary add a new middleware cluster
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"del_flag":0,"create_time":"2021-04-09T16:02:25.541701+08:00","last_update_time":"2021-04-09T16:02:25.541701+08:00","id":14,"cluster_name":"rest_test","env_id":1}]}"
// @Router /api/v1/metadata/middleware-cluster [post]
func AddMiddlewareCluster(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MiddlewareClusterInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, ok := fields[middlewareClusterClusterNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterClusterNameStruct)
		return
	}
	_, ok = fields[middlewareClusterEnvIDStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterEnvIDStruct)
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddMiddlewareCluster, fields[middlewareClusterClusterNameStruct], err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddMiddlewareCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddMiddlewareCluster, fields[middlewareClusterClusterNameStruct])
}

// @Tags middleware cluster
// @Summary update middleware cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id":13,"cluster_name":"new_test","env_id":1,"del_flag":1,"create_time":"2021-04-09T10:55:43.920406+08:00","last_update_time":"2021-04-09T10:55:43.920406+08:00"}]}"
// @Router /api/v1/metadata/middleware-cluster/update/:id [post]
func UpdateMiddlewareClusterByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(middlewareClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MiddlewareClusterInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, middlewareClusterNameExists := fields[middlewareClusterClusterNameStruct]
	_, middlewareClusterEnvIDExists := fields[middlewareClusterEnvIDStruct]
	_, delFlagExists := fields[envDelFlagStruct]
	if !middlewareClusterNameExists && !middlewareClusterEnvIDExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s %s and %s", middlewareClusterClusterNameStruct, middlewareClusterEnvIDStruct, envDelFlagStruct))
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateMiddlewareCluster, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, id, err.Error())
		return
	}
	// resp
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateMiddlewareCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataUpdateMiddlewareCluster, id)
}

// @Tags application
// @Summary delete middleware cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": []}"
// @Router /api/v1/metadata/app/delete/:id [post]
func DeleteMiddlewareClusterByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(middlewareClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// update entities
	err = s.Delete(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteMiddlewareCluster, fields[middlewareClusterClusterNameStruct], err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteMiddlewareCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteMiddlewareCluster, fields[middlewareClusterClusterNameStruct])
}

// @Tags application
// @Summary add user map
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"users":[{"del_flag":0,"id":1,"user_name":"zhangsan","department_name":"arch","account_name":"zs001","email":"allinemailtest@163.com","mobile":"13012345678","role":3,"create_time":"2021-10-25T09:21:50.364327+08:00","employee_id":"100001","telephone":"01012345678","last_update_time":"2021-11-22T13:46:20.430926+08:00"}]}]}"
// @Router /api/v1/metadata/middleware-cluster/add-user/:id [post]
func MiddlewareClusterAddUser(c *gin.Context) {
	// get params
	idStr := c.Param(middlewareClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}

	dataMap := make(map[string]int)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataMiddlewareClusterAddUser, id, err.Error())
		return
	}
	userID, userIDExists := dataMap[middlewareClusterUserIDJSON]
	if !userIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterUserIDJSON)
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// update entities
	err = s.AddUser(id, userID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataMiddlewareClusterAddUser, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(middlewareClusterUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataMiddlewareClusterAddUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataMiddlewareClusterAddUser, id, userID)
}

// @Tags application
// @Summary delete user map
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": {"users":[]}}"
// @Router /api/v1/metadata/middleware-cluster/delete-user/:id [post]
func MiddlewareClusterDeleteUser(c *gin.Context) {
	// get params
	idStr := c.Param(middlewareClusterIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	dataMap := make(map[string]int)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataMiddlewareClusterDeleteUser, id, err.Error())
		return
	}
	userID, userIDExists := dataMap[middlewareClusterUserIDJSON]
	if !userIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterUserIDJSON)
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// update entities
	err = s.DeleteUser(id, userID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataMiddlewareClusterDeleteUser, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(middlewareClusterUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataMiddlewareClusterDeleteUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataMiddlewareClusterDeleteUser, id, userID)
}
