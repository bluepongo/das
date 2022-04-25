package metadata

import (
	"fmt"
	"github.com/buger/jsonparser"
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
	monitorSystemIDJSON          = "id"
	monitorSystemNameJSON        = "system_name"
	monitorSystemTypeJSON        = "system_type"
	monitorSystemHostIPJSON      = "host_ip"
	monitorSystemPortNumJSON     = "port_num"
	monitorSystemPortNumSlowJSON = "port_num_slow"
	monitorSystemBaseUrlJSON     = "base_url"
	monitorSystemEnvIDJSON       = "env_id"
	monitorSystemDelFlagJSON     = "del_flag"

	monitorSystemIDStruct          = "ID"
	monitorSystemNameStruct        = "MonitorSystemName"
	monitorSystemTypeStruct        = "MonitorSystemType"
	monitorSystemHostIPStruct      = "MonitorSystemHostIP"
	monitorSystemPortNumStruct     = "MonitorSystemPortNum"
	monitorSystemPortNumSlowStruct = "MonitorSystemPortNumSlow"
	monitorSystemBaseUrlStruct     = "BaseURL"
	monitorSystemEnvIDStruct       = "EnvID"
	monitorSystemDelFlagStruct     = "DelFlag"
)

// @Tags    monitor system
// @Summary get all monitor systems
// @Accept	application/json
// @Param	token body string true "token"
// @Produce application/json
// @Success 200 {string} string "{"monitor_systems": [{"id": 1, "system_name": "pmm", "system_type": 1, "host_ip": "127.0.0.1", "port_num": 3306, "port_num_slow": 3307, "base_url": "http://127.0.0.1/prometheus/api/v1/", "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router  /api/v1/metadata/monitor-system [get]
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

// @Tags    monitor system
// @Summary get monitor system by env_id
// @Accept	application/json
// @Param	token  body string true "token"
// @Param	env_id body int    true "env id"
// @Produce application/json
// @Success 200 {string} string "{"monitor_systems": [{"id": 1, "system_name": "pmm", "system_type": 1, "host_ip": "127.0.0.1", "port_num": 3306, "port_num_slow": 3307, "base_url": "http://127.0.0.1/prometheus/api/v1/", "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router  /api/v1/metadata/monitor-system/env [get]
func GetMonitorSystemByEnv(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	envID, err := jsonparser.GetInt(data, monitorSystemEnvIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), monitorSystemEnvIDJSON)
		return
	}
	// init service
	s := metadata.NewMonitorSystemServiceWithDefault()
	// get entity
	err = s.GetByEnv(int(envID))
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

// @Tags    monitor system
// @Summary get monitor system by id
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id    body int    true "monitor system id"
// @Produce application/json
// @Success 200 {string} string "{"monitor_systems": [{"id": 1, "system_name": "pmm", "system_type": 1, "host_ip": "127.0.0.1", "port_num": 3306, "port_num_slow": 3307, "base_url": "http://127.0.0.1/prometheus/api/v1/", "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router  /api/v1/metadata/monitor-system/get [get]
func GetMonitorSystemByID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, monitorSystemIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), monitorSystemIDJSON)
		return
	}
	// init service
	s := metadata.NewMonitorSystemServiceWithDefault()
	// get entity
	err = s.GetByID(int(id))
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

// @Tags    monitor system
// @Summary get monitor system by host info
// @Accept	application/json
// @Param	token    body string true "token"
// @Param	host_ip  body string true "host ip"
// @Param 	port_num body int    true "port num"
// @Produce application/json
// @Success 200 {string} string "{"monitor_systems": [{"port_num_slow":9000,"base_url":"/prometheus","create_time":"2021-09-02T09:06:30.736111+08:00","last_update_time":"2021-11-18T16:16:18.702104+08:00","host_ip":"192.168.10.219","port_num":80,"env_id":1,"del_flag":0,"id":1,"system_name":"pmm2","system_type":2},{"base_url":"/prometheus","env_id":1,"del_flag":0,"create_time":"2021-09-02T15:11:19.558733+08:00","id":2,"port_num":80,"port_num_slow":33061,"last_update_time":"2021-11-10T10:01:52.717786+08:00","system_name":"pmm1","system_type":1,"host_ip":"192.168.10.220"}]}"
// @Router  /api/v1/metadata/monitor-system/host-info [get]
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

// @Tags    monitor system
// @Summary add a new monitor system
// @Accept	application/json
// @Param	token         body string true "token"
// @Param	system_name	  body string true	"system name"
// @Param 	system_type   body int    true	"system type"
// @Param 	host_ip       body string true	"host ip"
// @Param 	port_num      body int    true	"port num"
// @Param 	port_num_slow body int    true	"port num slow"
// @Param 	base_url      body string true	"base url"
// @Param 	env_id        body int    true	"env id"
// @Produce application/json
// @Success 200 {string} string "{"monitor_systems": [{"id":40,"system_name":"new_monitor_system","system_type":2,"host_ip":"192.168.10.219","port_num":8080,"port_num_slow":9000,"create_time":"2022-03-02T12:06:38.622752+08:00","env_id":1,"del_flag":0,"last_update_time":"2022-03-02T12:06:38.622752+08:00","base_url":"/prometheus"}]}"
// @Router  /api/v1/metadata/monitor-system [post]
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
			monitorSystemNameJSON, monitorSystemTypeJSON, monitorSystemHostIPJSON, monitorSystemPortNumJSON, monitorSystemPortNumSlowJSON,
			monitorSystemBaseUrlJSON, monitorSystemEnvIDJSON))
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

// @Tags    monitor system
// @Summary update monitor system by id
// @Accept	application/json
// @Param	token         body string true "token"
// @Param	id		      body int	  true	"monitor system id"
// @Param	system_name	  body string false	"system name"
// @Param 	system_type   body int    false	"system type"
// @Param 	host_ip       body string false	"host ip"
// @Param 	port_num      body int    false	"port num"
// @Param 	port_num_slow body int    false	"port num slow"
// @Param 	base_url      body string false	"base url"
// @Param 	env_id        body int    false	"env id"
// @Param 	del_flag      body int    false	"delete flag"
// @Produce application/json
// @Success 200 {string} string "{"monitor_systems": [{"system_name":"update_monitor_system","host_ip":"192.168.10.219","port_num_slow":9000,"env_id":1,"id":1,"system_type":2,"port_num":3300,"base_url":"/prometheus","del_flag":0,"create_time":"2021-09-02T09:06:30.736111+08:00","last_update_time":"2021-11-18T16:16:18.702104+08:00"}]}"
// @Router  /api/v1/metadata/monitor-system/update [post]
func UpdateMonitorSystemByID(c *gin.Context) {
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
	idInterface, idExists := fields[monitorSystemIDStruct]
	if !idExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, monitorSystemIDJSON)
		return
	}
	id, ok := idInterface.(int)
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, monitorSystemIDJSON)
		return
	}
	_, systemNameExists := fields[monitorSystemNameStruct]
	_, systemTypeExists := fields[monitorSystemTypeStruct]
	_, hostIpExists := fields[monitorSystemHostIPStruct]
	_, portNumExists := fields[monitorSystemPortNumStruct]
	_, portNumSlowExists := fields[monitorSystemPortNumSlowStruct]
	_, baseUrlExists := fields[monitorSystemBaseUrlStruct]
	_, envIdExists := fields[monitorSystemEnvIDStruct]
	_, delFlagExists := fields[monitorSystemDelFlagStruct]
	if !systemNameExists && !systemTypeExists && !hostIpExists && !portNumExists && !portNumSlowExists && !baseUrlExists && !envIdExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s and %s and %s and %s and %s and %s and %s",
			monitorSystemNameJSON, monitorSystemTypeJSON, monitorSystemHostIPJSON, monitorSystemPortNumJSON,
			monitorSystemPortNumSlowJSON, monitorSystemBaseUrlJSON, monitorSystemEnvIDJSON, monitorSystemDelFlagJSON))
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

// @Tags    monitor system
// @Summary delete monitor system by id
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id    body int	  true "monitor system id"
// @Produce application/json
// @Success 200 {string} string "{"monitor_systems": [{"id":40,"system_type":2,"port_num_slow":9000,"env_id":1,"create_time":"2022-03-02T12:06:38.622752+08:00","system_name":"new_monitor_system","host_ip":"192.168.10.219","port_num":8080,"base_url":"/prometheus","del_flag":0,"last_update_time":"2022-03-02T12:06:38.622752+08:00"}]}"
// @Router  /api/v1/metadata/monitor-system/delete [post]
func DeleteMonitorSystemByID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, monitorSystemIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), monitorSystemIDJSON)
		return
	}
	// init service
	s := metadata.NewMonitorSystemServiceWithDefault()
	// update entity
	err = s.Delete(int(id))
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
