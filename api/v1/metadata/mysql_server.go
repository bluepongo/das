package metadata

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pingcap/errors"
	utilmeta "github.com/romberli/das/pkg/util/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"

	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/pkg/message"
	msgmeta "github.com/romberli/das/pkg/message/metadata"
	"github.com/romberli/das/pkg/resp"
)

const (
	mysqlServerIDJSON        = "id"
	mysqlServerClusterIDJSON = "cluster_id"
	mysqlServerHostIPJSON    = "host_ip"
	mysqlServerPortNumJSON   = "port_num"

	mysqlServerClusterIDStruct      = "ClusterID"
	mysqlServerServerNameStruct     = "ServerName"
	mysqlServerHostIPStruct         = "HostIP"
	mysqlServerPortNumStruct        = "PortNum"
	mysqlServerDeploymentTypeStruct = "DeploymentType"
	mysqlServerVersionStruct        = "Version"

	mysqlServerMySQLClusterStruct = "MySQLCluster"

	isMasterResponse = `{"is_master": "%t"}`
)

// @Tags mysql server
// @Summary get all mysql servers
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"cluster_id":1,"deployment_type":1,"host_ip":"host_ip_init","port_num":3306,"version":"1.1.1","del_flag":0,"create_time":"2021-02-23T23:43:37.236228+08:00","last_update_time":"2021-02-23T23:43:37.236228+08:00","id":1}]}"
// @Router /api/v1/metadata/mysql-server [get]
func GetMySQLServer(c *gin.Context) {
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLServerAll, err)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLServerAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLServerAll)
}

// @Tags mysql server
// @Summary get mysql servers by cluster id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"cluster_id":1,"deployment_type":1,"host_ip":"host_ip_init","port_num":3306,"version":"1.1.1","del_flag":0,"create_time":"2021-02-23T23:43:37.236228+08:00","last_update_time":"2021-02-23T23:43:37.236228+08:00","id":1}]}"
// @Router /api/v1/metadata/mysql-server/cluster-id/:cluster_id [get]
func GetMySQLServerByClusterID(c *gin.Context) {
	// get param
	clusterIDStr := c.Param(mysqlServerClusterIDJSON)
	if clusterIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlServerIDJSON)
		return
	}
	clusterID, err := strconv.Atoi(clusterIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return

	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// get entity
	err = s.GetByClusterID(clusterID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLServerByClusterID, err, clusterID)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLServerByClusterID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLServerByClusterID, clusterID)
}

// @Tags mysql server
// @Summary get mysql server by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"port_num":3306,"del_flag":0,"version":"1.1.1","create_time":"2021-02-23T23:43:37.236228+08:00","last_update_time":"2021-02-23T23:43:37.236228+08:00","id":1,"cluster_id":1,"host_ip":"host_ip_init","deployment_type":1}]}"
// @Router /api/v1/metadata/mysql-server/get/:id [get]
func GetMySQLServerByID(c *gin.Context) {
	// get param
	idStr := c.Param(mysqlServerIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlServerIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return

	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// get entity
	err = s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLServerByID, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLServerByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLServerByID, id)
}

// @Tags mysql server
// @Summary get mysql servers by host info
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"cluster_id":1,"deployment_type":1,"host_ip":"host_ip_init","port_num":3306,"version":"1.1.1","del_flag":0,"create_time":"2021-02-23T23:43:37.236228+08:00","last_update_time":"2021-02-23T23:43:37.236228+08:00","id":1}]}"
// @Router /api/v1/metadata/mysql-server/host-info [get]
func GetMySQLServerByHostInfo(c *gin.Context) {
	var rd *utilmeta.HostInfo
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}

	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// get entity
	err = s.GetByHostInfo(rd.GetHostIP(), rd.GetPortNum())
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLServerByHostInfo, err, rd.GetHostIP(), rd.GetPortNum())
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLServerByHostInfo, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLServerByHostInfo, rd.GetHostIP(), rd.GetPortNum())
}

// @Tags mysql server
// @Summary check if mysql server is a master node
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": {"is_master": "true"}}"
// @Router /api/v1/metadata/mysql-server/is-master/host-info [get]
func IsMaster(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	dataMap := make(map[string]string)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}

	hostIP, hostIPExists := dataMap[mysqlServerHostIPJSON]
	if !hostIPExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlServerHostIPJSON)
		return
	}
	portNumStr, portNumExists := dataMap[mysqlServerPortNumJSON]
	if !portNumExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlServerPortNumJSON)
		return
	}
	portNum, err := strconv.Atoi(portNumStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// get entity
	isMaster, err := s.IsMaster(hostIP, portNum)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataIsMaster, err, hostIP, portNum)
		return
	}
	// response
	jsonStr := fmt.Sprintf(isMasterResponse, isMaster)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataIsMaster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataIsMaster, hostIP, portNum)
}

// @Tags mysql server
// @Summary get mysql cluster by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"owner_id":1,"del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","id":1,"monitor_system_id":1,"env_id":1,"last_update_time":"2021-02-23T20:57:24.603009+08:00","cluster_name":"cluster_name_init","middleware_cluster_id":1}]}"
// @Router /api/v1/metadata/mysql-server/mysql-cluster/:id [get]
func GetMySQLClusterByMySQLServerID(c *gin.Context) {
	// get param
	idStr := c.Param(mysqlServerIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlServerIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return

	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	log.Debug("==========================")
	// get entity
	err = s.GetMySQLClusterByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLClusterByServerID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(mysqlServerMySQLClusterStruct)
	// jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLClusterByServerID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLClusterByServerID, id)
}

// @Tags mysql server
// @Summary add a new mysql server
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"create_time":"2021-02-24T02:47:19.589172+08:00","del_flag":0,"last_update_time":"2021-02-24T02:47:19.589172+08:00","id":93,"cluster_id":0,"host_ip":"192.168.1.1","port_num":3306,"deployment_type":0,"version":"5.7.35"}]}"
// @Router /api/v1/metadata/mysql-server [post]
func AddMySQLServer(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MySQLServerInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, clusterIDExists := fields[mysqlServerClusterIDStruct]
	_, serverNameExists := fields[mysqlServerServerNameStruct]
	_, hostIPExists := fields[mysqlServerHostIPStruct]
	_, portNumExists := fields[mysqlServerPortNumStruct]
	_, deploymentTypeExists := fields[mysqlServerDeploymentTypeStruct]
	if !clusterIDExists || !serverNameExists || !hostIPExists || !portNumExists || !deploymentTypeExists {
		resp.ResponseNOK(
			c, message.ErrFieldNotExists,
			fmt.Sprintf(
				"%s and %s and %s and %s and %s",
				mysqlServerClusterIDStruct,
				mysqlServerServerNameStruct,
				mysqlServerHostIPStruct,
				mysqlServerPortNumStruct,
				mysqlServerDeploymentTypeStruct))
		return
	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddMySQLServer,
			fields[mysqlServerServerNameStruct],
			fields[mysqlServerClusterIDStruct],
			fields[mysqlServerHostIPStruct],
			fields[mysqlServerPortNumStruct],
			fields[mysqlServerDeploymentTypeStruct],
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddMySQLServer, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddMySQLServer,
		fields[mysqlServerServerNameStruct],
		fields[mysqlServerClusterIDStruct],
		fields[mysqlServerHostIPStruct],
		fields[mysqlServerPortNumStruct],
		fields[mysqlServerDeploymentTypeStruct],
	)
}

// @Tags mysql server
// @Summary update mysql server by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"last_update_time":"2021-02-24T02:47:19.589172+08:00","id":93,"cluster_id":0,"host_ip":"192.168.1.1","version":"5.7.35","del_flag":1,"create_time":"2021-02-24T02:47:19.589172+08:00","port_num":3306,"deployment_type":0}]}"
// @Router /api/v1/metadata/mysql-server/:id [post]
func UpdateMySQLServerByID(c *gin.Context) {
	var fields map[string]interface{}
	// get param
	idStr := c.Param(mysqlServerIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlServerIDJSON)
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MySQLServerInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, clusterIDExists := fields[mysqlServerClusterIDStruct]
	_, serverNameExists := fields[mysqlServerServerNameStruct]
	_, hostIPExists := fields[mysqlServerHostIPStruct]
	_, portNumExists := fields[mysqlServerPortNumStruct]
	_, deploymentTypeExists := fields[mysqlServerDeploymentTypeStruct]
	_, versionExists := fields[mysqlServerVersionStruct]
	_, delFlagExists := fields[envDelFlagStruct]
	if !clusterIDExists &&
		!serverNameExists &&
		!hostIPExists &&
		!portNumExists &&
		!deploymentTypeExists &&
		!versionExists &&
		!delFlagExists {
		resp.ResponseNOK(
			c, message.ErrFieldNotExists,
			fmt.Sprintf("%s, %s, %s, %s, %s, %s and %s",
				fields[mysqlServerClusterIDStruct],
				fields[mysqlServerServerNameStruct],
				fields[mysqlServerHostIPStruct],
				fields[mysqlServerPortNumStruct],
				fields[mysqlServerDeploymentTypeStruct],
				fields[mysqlServerVersionStruct],
				fields[envDelFlagStruct]))
		return
	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateMySQLServer, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateMySQLServer, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.DebugMetadataUpdateMySQLServer, fields[mysqlServerServerNameStruct])
}

// @Tags mysql server
// @Summary get mysql servers by host info
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"last_update_time":"2021-02-24T02:47:19.589172+08:00","id":93,"cluster_id":0,"host_ip":"192.168.1.1","version":"5.7.35","del_flag":1,"create_time":"2021-02-24T02:47:19.589172+08:00","port_num":3306,"deployment_type":0}]}"
// @Router /api/v1/metadata/mysql-server/:id [get]
func DeleteMySQLServerByID(c *gin.Context) {
	var fields map[string]interface{}

	// get param
	idStr := c.Param(mysqlServerIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlServerIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return

	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// insert into middleware
	err = s.Delete(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteMySQLServer,
			err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteMySQLServer, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteMySQLServer, fields[mysqlServerServerNameStruct])
}
