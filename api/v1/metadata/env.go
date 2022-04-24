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
	envIDJSON      = "id"
	envEnvNameJSON = "env_name"

	envIDStruct      = "ID"
	envDelFlagStruct = "DelFlag"
	envEnvNameStruct = "EnvName"
)

// @Tags	environment
// @Summary	get all environments
// @Accept	application/json
// @Param	token body string true "token"
// @Produce application/json
// @Success	200 {string} string "{"envs": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router	/api/v1/metadata/env [get]
func GetEnv(c *gin.Context) {
	// init service
	s := metadata.NewEnvServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetEnvAll, err)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetEnvAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetEnvAll)
}

// @Tags	environment
// @Summary get environment by id
// @Accept	application/json
// @Param	token	body string true "token"
// @Param	id 		body int 	true "env id"
// @Produce application/json
// @Success	200 {string} string "{"envs": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router	/api/v1/metadata/env [get]
func GetEnvByID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, envIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), envIDJSON)
		return
	}
	// init service
	s := metadata.NewEnvServiceWithDefault()
	// get entity
	err = s.GetByID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetEnvByID, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetEnvByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetEnvByID, id)
}

// @Tags 	environment
// @Summary	get environment by Name
// @Accept	application/json
// @Param	token	 body string true "token"
// @Param	env_name body string true "env name"
// @Produce application/json
// @Success 200 {string} string "{"envs": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router	/api/v1/metadata/env/env-name [get]
func GetEnvByName(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	envName, err := jsonparser.GetString(data, envEnvNameJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), envEnvNameJSON)
		return
	}
	// init service
	s := metadata.NewEnvServiceWithDefault()
	// get entity
	err = s.GetEnvByName(envName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetEnvByName, err, envName)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetEnvByName, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetEnvByName, envName)
}

// @Tags	environment
// @Summary add a new environment
// @Accept	application/json
// @Param	token	 body string true "token"
// @Param	env_name body string true "env name"
// @Produce application/json
// @Success 200 {string} string "{"envs": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router	/api/v1/metadata/env [post]
func AddEnv(c *gin.Context) {
	var fields map[string]interface{}
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.EnvInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, ok := fields[envEnvNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, envEnvNameStruct)
		return
	}
	// init service
	s := metadata.NewEnvServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddEnv, err, fields[envEnvNameStruct])
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddEnv, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddEnv, fields[envEnvNameStruct])
}

// @Tags	environment
// @Summary	update environment by id
// @Accept	application/json
// @Param	token	 body 	string 	true 	"token"
// @Param	id		 body 	int		true	"env id"
// @Param 	env_name body 	string 	false	"env name"
// @Param 	del_flag body 	int		false	"delete flag"
// @Produce application/json
// @Success	200 {string} string "{"envs": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router	/api/v1/metadata/env/update [post]
func UpdateEnvByID(c *gin.Context) {
	var fields map[string]interface{}
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.EnvInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	idInterface, idExists := fields[envIDStruct]
	if !idExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, envIDStruct)
		return
	}
	id, ok := idInterface.(int)
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, envIDStruct)
		return
	}
	_, envNameExists := fields[envEnvNameStruct]
	_, delFlagExists := fields[envDelFlagStruct]
	if !envNameExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s", envEnvNameStruct, envDelFlagStruct))
		return
	}
	// init service
	s := metadata.NewEnvServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateEnv, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateEnv, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.DebugMetadataUpdateEnv, id)
}

// @Tags	environment
// @Summary delete environment by id
// @Accept	application/json
// @Param	token   body 	string 	true 	"token"
// @Param	id      body 	int 	true 	"env id"
// @Produce application/json
// @Success	200 {string} string "{"envs": [{"id": 1, "env_name": "online", "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router	/api/v1/metadata/env/delete [post]
func DeleteEnvByID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, envIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), envIDJSON)
		return
	}
	// init service
	s := metadata.NewEnvServiceWithDefault()
	// update entities
	err = s.Delete(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteEnvByID, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteEnvByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteEnvByID, id)
}
