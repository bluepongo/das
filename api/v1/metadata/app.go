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
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
)

const (
	appIDJSON      = "id"
	appAppNameJSON = "app_name"
	appDBIDJSON    = "db_id"
	appUserIDJSON  = "user_id"
	appLevelJSON   = "level"
	appDelFlagJSON = "del_flag"

	appAppIDStruct   = "ID"
	appAppNameStruct = "AppName"
	appLevelStruct   = "Level"
	appDelFlagStruct = "DelFlag"
	appDBsStruct     = "DBs"
	appUsersStruct   = "Users"
)

// @Tags 	application
// @Summary get all applications
// @Accept	application/json
// @Param	token body string true "token"
// @Produce application/json
// @Success 200 {string} string "{"apps": [ {"id": 2,"app_name": "app2","level": 2,"del_flag": 0,"create_time": "2021-10-25T09:46:04.275796+08:00","last_update_time": "2021-12-21T09:15:47.688546+08:00"}]}"
// @Router 	/api/v1/metadata/app [get]
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

// @Tags 	application
// @Summary get application by id
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id    body int 	  true "app id"
// @Produce application/json
// @Success 200 {string} string "{"apps": [ {"id": 2,"app_name": "app2","level": 2,"del_flag": 0,"create_time": "2021-10-25T09:46:04.275796+08:00","last_update_time": "2021-12-21T09:15:47.688546+08:00"}]}"
// @Router  /api/v1/metadata/app/get [get]
func GetAppByID(c *gin.Context) {
	// get param
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	appID, err := jsonparser.GetInt(data, appIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), appIDJSON)
		return
	}
	id := int(appID)
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

// @Tags 	application
// @Summary get application by system name
// @Accept	application/json
// @Param	token body string true "token"
// @Param	name  body string true "app name"
// @Produce application/json
// @Success 200 {string} string "{"apps": [ {"id": 2,"app_name": "app2","level": 2,"del_flag": 0,"create_time": "2021-10-25T09:46:04.275796+08:00","last_update_time": "2021-12-21T09:15:47.688546+08:00"}]}"
// @Router 	/api/v1/metadata/app/app-name [get]
func GetAppByName(c *gin.Context) {
	// get param
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	appName, err := jsonparser.GetString(data, appAppNameJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), appAppNameJSON)
		return
	}
	// init service
	s := metadata.NewAppServiceWithDefault()
	// get entity
	err = s.GetAppByName(appName)
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

// @Tags 	application
// @Summary get dbs
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id    body int    true "app id"
// @Produce application/json
// @Success 200 {string} string "{"dbs": [{"id": 1,"db_name": "db2","cluster_id": 3,"cluster_type": 1,"env_id": 1,"del_flag": 0,"create_time": "2022-01-04T15:08:33.418288+08:00","last_update_time": "2022-01-25T16:17:26.284761+08:00"},}]}"
// @Router 	/api/vi/metadata/app/db [get]
func GetDBsByAppID(c *gin.Context) {
	// get param
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	appID, err := jsonparser.GetInt(data, appIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), appIDJSON)
		return
	}
	id := int(appID)
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

// @Tags 	application
// @Summary add a new application
// @Accept	application/json
// @Param	token 	 body string true "token"
// @Param	app_name body string true "app name"
// @Param	level 	 body int    true "app level"
// @Produce application/json
// @Success 200 {string} string "{"apps": [ {"id": 2,"app_name": "app2","level": 2,"del_flag": 0,"create_time": "2021-10-25T09:46:04.275796+08:00","last_update_time": "2021-12-21T09:15:47.688546+08:00"}]}"
// @Router 	/api/v1/metadata/app [post]
func AddApp(c *gin.Context) {
	var fields map[string]interface{}

	// get param
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
		resp.ResponseNOK(c, message.ErrFieldNotExists, appAppNameJSON)
		return
	}
	_, ok = fields[appLevelStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appLevelJSON)
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

// @Tags 	application
// @Summary update application by id
// @Accept	application/json
// @Param	token    body string true   "token"
// @Param	id       body int    true   "app id"
// @Param	app_name body string false  "app name"
// @Param	level    body int    false  "app level"
// @Param 	del_flag body int    false  "delete flag"
// @Produce application/json
// @Success 200 {string} string "{"apps": [ {"id": 2,"app_name": "app2","level": 2,"del_flag": 0,"create_time": "2021-10-25T09:46:04.275796+08:00","last_update_time": "2021-12-21T09:15:47.688546+08:00"}]}"
// @Router  /api/v1/metadata/app [post]
func UpdateAppByID(c *gin.Context) {
	var fields map[string]interface{}

	// get param
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// appID, err := jsonparser.GetInt(data, appIDJSON)
	// if err != nil {
	// 	resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), appIDJSON)
	// 	return
	// }
	// id := int(appID)
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.AppInfo{}, constant.DefaultJSONTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	idInterface, IDExists := fields[appAppIDStruct]

	if !IDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
		return
	}

	id, ok := idInterface.(int)
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, appIDJSON)
		return
	}
	_, AppNameExists := fields[appAppNameStruct]
	_, levelExists := fields[appLevelStruct]
	_, delFlagExists := fields[appDelFlagStruct]
	if !AppNameExists && !delFlagExists && !levelExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf(" %s ,%s and %s", appAppNameJSON, appLevelJSON, appDelFlagJSON))
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

// @Tags 	 application
// @Summary  delete app by id
// @Accept   application/json
// @Param	token body string true "token"
// @Param	id    body int    true "app id"
// @Produce  application/json
// @Success 200 {string} string "{"apps": [ {"id": 2,"app_name": "app2","level": 2,"del_flag": 0,"create_time": "2021-10-25T09:46:04.275796+08:00","last_update_time": "2021-12-21T09:15:47.688546+08:00"}]}"
// @Router 	 /api/v1/metadata/app/delete [post]
func DeleteAppByID(c *gin.Context) {
	var fields map[string]interface{}

	// get param
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	appID, err := jsonparser.GetInt(data, appIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), appIDJSON)
		return
	}
	id := int(appID)
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

// @Tags 	application
// @Summary add database map
// @Accept	application/json
// @Param	token   body string true "token"
// @Param	id      body int    true "app id"
// @Param	db_id   body int    true "db id"
// @Produce application/json
// @Success 200 {string} string "{"dbs": [{"id": 1,"db_name": "db2","cluster_id": 3,"cluster_type": 1,"env_id": 1,"del_flag": 0,"create_time": "2022-01-04T15:08:33.418288+08:00","last_update_time": "2022-01-25T16:17:26.284761+08:00"},}]}"
// @Router 	/api/v1/metadata/app/add-db [post]
func AppAddDB(c *gin.Context) {
	// get param
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	appID, err := jsonparser.GetInt(data, appIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), appIDJSON)
		return
	}
	id := int(appID)

	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAppAddDB, errors.Trace(err), id)
		return
	}
	dbID64, err := jsonparser.GetInt(data, appDBIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), appDBIDJSON)
		return
	}
	dbID := int(dbID64)
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

// @Tags 	application
// @Summary delete database map
// @Accept	application/json
// @Param	token body string true  "token"
// @Param	id    body int true     "app id"
// @Param	db_id body int false 	"db id"
// @Produce application/json
// @Success 200 {string} string "{"dbs": [{"id": 1,"db_name": "db2","cluster_id": 3,"cluster_type": 1,"env_id": 1,"del_flag": 0,"create_time": "2022-01-04T15:08:33.418288+08:00","last_update_time": "2022-01-25T16:17:26.284761+08:00"},}]}"
// @Router  /api/v1/metadata/app/delete-db [post]
func AppDeleteDB(c *gin.Context) {
	// get param
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	appID, err := jsonparser.GetInt(data, appIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), appIDJSON)
		return
	}
	id := int(appID)

	dbID64, err := jsonparser.GetInt(data, appDBIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), appDBIDJSON)
		return
	}
	dbID := int(dbID64)
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

// @Tags 	application
// @Summary get users
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id    body int    true "app id"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"id": 18,"employee_id": "21213434","account_name": "kf-Tom","mobile": "18088888888","role": 2,"user_name": "Tom","department_name": "kf","email": "test@test.com.cn","telephone": "02188888888","del_flag": 0,"create_time": "2022-03-07T15:56:32.277857+08:00","last_update_time": "2022-03-07T15:56:32.277857+08:00"}]}"
// @Router  /api/vi/metadata/app/user [get]
func GetUsersByAppID(c *gin.Context) {
	// get param
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	appID, err := jsonparser.GetInt(data, appIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), appIDJSON)
		return
	}
	id := int(appID)
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

// @Tags 	application
// @Summary add user map
// @Accept	application/json
// @Param	token 	body string true  "token"
// @Param	id      body int    true  "app id"
// @Param	user_id body int    false "user id"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"id": 18,"employee_id": "21213434","account_name": "kf-Tom","mobile": "18088888888","role": 2,"user_name": "Tom","department_name": "kf","email": "test@test.com.cn","telephone": "02188888888","del_flag": 0,"create_time": "2022-03-07T15:56:32.277857+08:00","last_update_time": "2022-03-07T15:56:32.277857+08:00"}]}"
// @Router  /api/v1/metadata/app/add-user [post]
func AppAddUser(c *gin.Context) {
	// get param
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	appID, err := jsonparser.GetInt(data, appIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), appIDJSON)
		return
	}
	id := int(appID)

	userID64, err := jsonparser.GetInt(data, appUserIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), appUserIDJSON)
		return
	}
	userID := int(userID64)

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

// @Tags 	application
// @Summary delete user map
// @Accept	application/json
// @Param	token 	body string true "token"
// @Param	id      body int    true "app id"
// @Param	user_id body int    true "user id"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"id": 18,"employee_id": "21213434","account_name": "kf-Tom","mobile": "18088888888","role": 2,"user_name": "Tom","department_name": "kf","email": "test@test.com.cn","telephone": "02188888888","del_flag": 0,"create_time": "2022-03-07T15:56:32.277857+08:00","last_update_time": "2022-03-07T15:56:32.277857+08:00"}]}"
// @Router  /api/v1/metadata/app/delete-user [post]
func AppDeleteUser(c *gin.Context) {
	// get param
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	appID, err := jsonparser.GetInt(data, appIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), appIDJSON)
		return
	}
	id := int(appID)

	userID64, err := jsonparser.GetInt(data, appUserIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), appUserIDJSON)
		return
	}
	userID := int(userID64)
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
