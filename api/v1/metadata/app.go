package metadata

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pingcap/errors"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/pkg/message"
	msgmeta "github.com/romberli/das/pkg/message/metadata"
	"github.com/romberli/das/pkg/resp"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
)

const (
	appIDJSON      = "id"
	appAppNameJSON = "app_name"
	appDBIDJSON    = "db_id"
	appUserIDJSON  = "user_id"

	appAppNameStruct = "AppName"
	appLevelStruct   = "Level"
	appDelFlagStruct = "DelFlag"
	appDBsStruct     = "DBs"
	appUsersStruct   = "Users"
)

// @Tags application
// @Summary get all applications
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8}]}"
// @Router /api/v1/metadata/app [get]
func GetApp(c *gin.Context) {
	// init service
	s := metadata.NewAppServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAppAll, err)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAppAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAppAll)
}

// @Tags application
// @Summary get application by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8}]}"
// @Router /api/v1/metadata/app/:id [get]
func GetAppByID(c *gin.Context) {
	// get param
	idStr := c.Param(appIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// get entity
	err = s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAppByID, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAppByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAppByID, id)
}

// @Tags application
// @Summary get application by system name
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "app_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8}]}"
// @Router /api/v1/metadata/app/app-name/:name [get]
func GetAppByName(c *gin.Context) {
	// get params
	appName := c.Param(appAppNameJSON)
	if appName == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appAppNameJSON)
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// get entity
	err := s.GetAppByName(appName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAppByName, err, appName)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAppByName, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAppByName, appName)
}

// @Tags application
// @Summary get dbs
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [1, 2]}"
// @Router /api/vi/metadata/app/db/:id [get]
func GetDBsByAppID(c *gin.Context) {
	// get param
	idStr := c.Param(appIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// get entity
	err = s.GetDBsByAppID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBsByAppID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(appDBsStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDBsByAppID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBsByAppID, id)
}

// @Tags application
// @Summary add a new application
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8}]}"
// @Router /api/v1/metadata/app [post]
func AddApp(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.AppInfo{}, constant.DefaultJSONTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, ok := fields[appAppNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appAppNameStruct)
		return
	}
	_, ok = fields[appLevelStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appLevelStruct)
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddApp, err, fields[appAppNameStruct])
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddApp, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddApp, fields[appAppNameStruct])
}

// @Tags application
// @Summary update application by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8}]}"
// @Router /api/v1/metadata/app/:id [post]
func UpdateAppByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(appIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.AppInfo{}, constant.DefaultJSONTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, AppNameExists := fields[appAppNameStruct]
	_, levelExists := fields[appLevelStruct]
	_, delFlagExists := fields[appDelFlagStruct]
	if !AppNameExists && !delFlagExists && !levelExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s", appAppNameStruct, appDelFlagStruct))
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// update entities
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateApp, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateApp, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataUpdateApp, fields[appAppNameStruct])
}

// @Tags application
// @Summary delete app by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8}]}"
// @Router /api/v1/metadata/app/delete/:id [post]
func DeleteAppByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(appIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// update entities
	err = s.Delete(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteApp, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteApp, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteApp, fields[appAppNameStruct])
}

// @Tags application
// @Summary add database map
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8}]}"
// @Router /api/v1/metadata/app/add-db/:id [post]
func AppAddDB(c *gin.Context) {
	// get params
	idStr := c.Param(appIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
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
		resp.ResponseNOK(c, msgmeta.ErrMetadataAppAddDB, errors.Trace(err), id)
		return
	}
	dbID, dbIDExists := dataMap[appDBIDJSON]
	if !dbIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appDBIDJSON)
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// update entities
	err = s.AddDB(id, dbID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAppAddDB, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(appDBsStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAppAddDB, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAppAddDB, id, dbID)
}

// @Tags application
// @Summary delete database map
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [1]}"
// @Router /api/v1/metadata/app/delete-db/:id [post]
func AppDeleteDB(c *gin.Context) {
	// get params
	idStr := c.Param(appIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
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
		resp.ResponseNOK(c, msgmeta.ErrMetadataAppDeleteDB, errors.Trace(err), id)
		return
	}
	dbID, dbIDExists := dataMap[appDBIDJSON]
	if !dbIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appDBIDJSON)
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// update entities
	err = s.DeleteDB(id, dbID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAppDeleteDB, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(appDBsStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAppDeleteDB, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAppDeleteDB, id, dbID)
}

// @Tags application
// @Summary get users
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [1, 2]}"
// @Router /api/vi/metadata/app/user/:id [get]
func GetUsersByAppID(c *gin.Context) {
	// get param
	idStr := c.Param(appIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// get entity
	err = s.GetUsersByAppID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetUsersByAppID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(appUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetUsersByAppID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetUsersByAppID, id)
}

// @Tags application
// @Summary add user map
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 66, "system_name": "kkk", "del_flag": 0, "create_time": "2021-01-21T10:00:00+08:00", "last_update_time": "2021-01-21T10:00:00+08:00", "level": 8,"owner_id": 8}]}"
// @Router /api/v1/metadata/app/add-user/:id [post]
func AppAddUser(c *gin.Context) {
	// get params
	idStr := c.Param(appIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
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
		resp.ResponseNOK(c, msgmeta.ErrMetadataAppAddUser, errors.Trace(err), id)
		return
	}
	userID, userIDExists := dataMap[appUserIDJSON]
	if !userIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appUserIDJSON)
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// update entities
	err = s.AddUser(id, userID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAppAddUser, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(appUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAppAddUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAppAddUser, id, userID)
}

// @Tags application
// @Summary delete user map
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [1]}"
// @Router /api/v1/metadata/app/delete-user/:id [post]
func AppDeleteUser(c *gin.Context) {
	// get params
	idStr := c.Param(appIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
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
		resp.ResponseNOK(c, msgmeta.ErrMetadataAppDeleteUser, errors.Trace(err), id)
		return
	}
	userID, userIDExists := dataMap[appUserIDJSON]
	if !userIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appUserIDJSON)
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// update entities
	err = s.DeleteUser(id, userID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAppDeleteUser, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(appUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAppDeleteUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAppDeleteUser, id, userID)
}
