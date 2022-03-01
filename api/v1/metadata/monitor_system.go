package metadata

import (
	"fmt"
	"strconv"

	"github.com/pingcap/errors"

	"github.com/gin-gonic/gin"
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
	monitorSystemIDJSON    = "id"
	monitorSystemEnvIDJSON = "env_id"

	monitorSystemNameStruct        = "MonitorSystemName"
	monitorSystemTypeStruct        = "MonitorSystemType"
	monitorSystemHostIPStruct      = "MonitorSystemHostIP"
	monitorSystemPortNumStruct     = "MonitorSystemPortNum"
	monitorSystemPortNumSlowStruct = "MonitorSystemPortNumSlow"
	monitorSystemBaseUrlStruct     = "BaseURL"
	monitorSystemEnvIDStruct       = "EnvID"
)

// @Tags monitor system
// @Summary get all monitor systems
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "system_name": "pmm", "system_type": 1, "host_ip": "127.0.0.1", "port_num": 3306, "port_num_slow": 3307, "base_url": "http://127.0.0.1/prometheus/api/v1/", "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/monitor-system [get]
func GetMonitorSystem(c *gin.Context) {
	// init service
	s := metadata.NewMonitorSystemServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMonitorSystemAll, err)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMonitorSystemAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMonitorSystemAll)
}

// @Tags monitor system
// @Summary get monitor system by env_id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "system_name": "pmm", "system_type": 1, "host_ip": "127.0.0.1", "port_num": 3306, "port_num_slow": 3307, "base_url": "http://127.0.0.1/prometheus/api/v1/", "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/monitor-system/env/:env_id [get]
func GetMonitorSystemByEnv(c *gin.Context) {
	// get param
	envIDStr := c.Param(monitorSystemEnvIDJSON)
	if envIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, monitorSystemEnvIDJSON)
		return
	}
	envID, err := strconv.Atoi(envIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return

	}
	// init service
	s := metadata.NewMonitorSystemServiceWithDefault()
	// get entity
	err = s.GetByEnv(envID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMonitorSystemByEnv, err)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMonitorSystemByEnv, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMonitorSystemByEnv, envID)

}

// @Tags monitor system
// @Summary get monitor system by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "system_name": "pmm", "system_type": 1, "host_ip": "127.0.0.1", "port_num": 3306, "port_num_slow": 3307, "base_url": "http://127.0.0.1/prometheus/api/v1/", "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/monitor-system/get/:id [get]
func GetMonitorSystemByID(c *gin.Context) {
	// get param
	idStr := c.Param(monitorSystemIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, monitorSystemIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewMonitorSystemServiceWithDefault()
	// get entity
	err = s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMonitorSystemByID, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMonitorSystemByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMonitorSystemByID, id)
}

// @Tags monitor system
// @Summary get monitor system by host info
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "system_name": "pmm", "system_type": 1, "host_ip": "127.0.0.1", "port_num": 3306, "port_num_slow": 3307, "base_url": "http://127.0.0.1/prometheus/api/v1/", "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/monitor-system/host-info [get]
func GetMonitorSystemByHostInfo(c *gin.Context) {
	var rd *utilmeta.HostInfo
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}

	// init service
	s := metadata.NewMonitorSystemServiceWithDefault()
	// get entity
	err = s.GetByHostInfo(rd.GetHostIP(), rd.GetPortNum())
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMonitorSystemByHostInfo, err, rd.GetHostIP(), rd.GetPortNum())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMonitorSystemByHostInfo, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMonitorSystemByHostInfo, rd.GetHostIP(), rd.GetPortNum())
}

// @Tags monitor system
// @Summary add a new monitor system
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "system_name": "pmm", "system_type": 1, "host_ip": "127.0.0.1", "port_num": 3306, "port_num_slow": 3307, "base_url": "http://127.0.0.1/prometheus/api/v1/", "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/monitor-system [post]
func AddMonitorSystem(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MonitorSystemInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, systemNameExists := fields[monitorSystemNameStruct]
	_, systemTypeExists := fields[monitorSystemTypeStruct]
	_, hostIPExists := fields[monitorSystemHostIPStruct]
	_, portNumExists := fields[monitorSystemPortNumStruct]
	_, portNumSlowExists := fields[monitorSystemPortNumSlowStruct]
	_, baseUrlExists := fields[monitorSystemBaseUrlStruct]
	_, envIDExists := fields[monitorSystemEnvIDStruct]
	if !systemNameExists || !systemTypeExists || !hostIPExists || !portNumExists || !portNumSlowExists || !baseUrlExists || !envIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s and %s and %s and %s and %s and %s",
			monitorSystemNameStruct, monitorSystemTypeStruct, monitorSystemHostIPStruct, monitorSystemPortNumStruct, monitorSystemPortNumSlowStruct,
			monitorSystemBaseUrlStruct, monitorSystemEnvIDStruct))
		return
	}
	// init service
	s := metadata.NewMonitorSystemServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddMonitorSystem, err,
			fields[monitorSystemNameStruct], fields[monitorSystemTypeStruct], fields[monitorSystemHostIPStruct],
			fields[monitorSystemPortNumStruct], fields[monitorSystemPortNumSlowStruct], fields[monitorSystemBaseUrlStruct],
			fields[monitorSystemEnvIDStruct])
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddMonitorSystem, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddMonitorSystem, fields[monitorSystemNameStruct], fields[monitorSystemTypeStruct],
		fields[monitorSystemHostIPStruct], fields[monitorSystemPortNumStruct], fields[monitorSystemPortNumSlowStruct],
		fields[monitorSystemBaseUrlStruct], fields[monitorSystemEnvIDStruct])
}

// @Tags monitor system
// @Summary update monitor system by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "system_name": "pmm", "system_type": 1, "host_ip": "127.0.0.1", "port_num": 3306, "port_num_slow": 3307, "base_url": "http://127.0.0.1/prometheus/api/v1/", "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/monitor-system/update/:id [post]
func UpdateMonitorSystemByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(monitorSystemIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, monitorSystemIDJSON)
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MonitorSystemInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, systemNameExists := fields[monitorSystemNameStruct]
	_, systemTypeExists := fields[monitorSystemTypeStruct]
	_, hostIpExists := fields[monitorSystemHostIPStruct]
	_, portNumExists := fields[monitorSystemPortNumStruct]
	_, portNumSlowExists := fields[monitorSystemPortNumSlowStruct]
	_, baseUrlExists := fields[monitorSystemBaseUrlStruct]
	_, envIdExists := fields[monitorSystemEnvIDStruct]
	_, delFlagExists := fields[envDelFlagStruct]
	if !systemNameExists && !systemTypeExists && !hostIpExists && !portNumExists && !portNumSlowExists && !baseUrlExists && !envIdExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s and %s and %s and %s and %s and %s and %s",
			monitorSystemNameStruct, monitorSystemTypeStruct, monitorSystemHostIPStruct, monitorSystemPortNumStruct,
			monitorSystemPortNumSlowStruct, monitorSystemBaseUrlStruct, monitorSystemEnvIDStruct, envDelFlagStruct))
		return
	}
	// init service
	s := metadata.NewMonitorSystemServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateMonitorSystem, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateMonitorSystem, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataUpdateMonitorSystem, id)
}

// @Tags monitor system
// @Summary delete monitor system by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "system_name": "pmm", "system_type": 1, "host_ip": "127.0.0.1", "port_num": 3306, "port_num_slow": 3307, "base_url": "http://127.0.0.1/prometheus/api/v1/", "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/monitor-system/delete/:id [post]
func DeleteMonitorSystemByID(c *gin.Context) {
	// get params
	idStr := c.Param(monitorSystemIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, monitorSystemIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewMonitorSystemServiceWithDefault()
	// update entity
	err = s.Delete(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteMonitorSystem, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteMonitorSystem, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteMonitorSystem, id)
}
