package metadata

import (
	"fmt"

	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/pingcap/errors"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"

	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/pkg/message"
	msgmeta "github.com/romberli/das/pkg/message/metadata"
	"github.com/romberli/das/pkg/resp"
)

const (
	mysqlServerIDJSON             = "id"
	mysqlServerClusterIDJSON      = "cluster_id"
	mysqlServerServerNameJSON     = "server_name"
	mysqlServerHostIPJSON         = "host_ip"
	mysqlServerPortNumJSON        = "port_num"
	mysqlServerDeploymentTypeJSON = "deployment_type"

	mysqlServerIDStruct             = "ID"
	mysqlServerClusterIDStruct      = "ClusterID"
	mysqlServerServerNameStruct     = "ServerName"
	mysqlServerHostIPStruct         = "HostIP"
	mysqlServerPortNumStruct        = "PortNum"
	mysqlServerDeploymentTypeStruct = "DeploymentType"
	mysqlServerVersionStruct        = "Version"
	mysqlServerMySQLClusterStruct   = "MySQLCluster"

	isMasterResponse = `{"is_master": "%t"}`
)

// @Tags	mysql server
// @Summary	get all mysql servers
// @Accept	application/json
// @Produce	application/json
// @Param	token	body string	true "token"
// @Success	200 {string} string "{"mysql_servers":[{"port_num":3306,"create_time":"2021-09-02T11:16:06.561525+08:00","last_update_time":"2022-03-01T08:19:09.779365+08:00","cluster_id":1,"server_name":"192-168-10-219","service_name":"192-168-10-219:3306","host_ip":"192.168.10.219","id":1,"deployment_type":1,"version":"5.7","del_flag":0}]}"
// @Router	/api/v1/metadata/mysql-server [get]
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

// @Tags	mysql server
// @Summary get mysql servers by cluster id
// @Accept	application/json
// @Param	token		body string	true "token"
// @Param	cluster_id	body int	true "mysql cluster id"
// @Produce	application/json
// @Success	200 {string} string "{"mysql_servers":[{"port_num":3306,"create_time":"2021-09-02T11:16:06.561525+08:00","last_update_time":"2022-03-01T08:19:09.779365+08:00","cluster_id":1,"server_name":"192-168-10-219","service_name":"192-168-10-219:3306","host_ip":"192.168.10.219","id":1,"deployment_type":1,"version":"5.7","del_flag":0}]}"
// @Router	/api/v1/metadata/mysql-server/cluster-id [get]
func GetMySQLServerByClusterID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	clusterID, err := jsonparser.GetInt(data, mysqlServerClusterIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), mysqlServerClusterIDJSON)
		return
	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// get entity
	err = s.GetByClusterID(int(clusterID))
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

// @Tags	mysql server
// @Summary	get mysql server by id
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	id		body int	true "mysql server id"
// @Produce	application/json
// @Success	200 {string} string "{"mysql_servers":[{"port_num":3306,"create_time":"2021-09-02T11:16:06.561525+08:00","last_update_time":"2022-03-01T08:19:09.779365+08:00","cluster_id":1,"server_name":"192-168-10-219","service_name":"192-168-10-219:3306","host_ip":"192.168.10.219","id":1,"deployment_type":1,"version":"5.7","del_flag":0}]}"
// @Router	/api/v1/metadata/mysql-server/get [get]
func GetMySQLServerByID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, mysqlServerIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), mysqlServerIDJSON)
		return
	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// get entity
	err = s.GetByID(int(id))
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

// @Tags	mysql server
// @Summary	get mysql servers by host info
// @Accept	application/json
// @Param	token		body string	true "token"
// @Param	host_ip		body string	true "host ip"
// @Param	port_num	body int	true "host port number"
// @Produce	application/json
// @Success	200 {string} string "{"mysql_servers":[{"port_num":3306,"create_time":"2021-09-02T11:16:06.561525+08:00","last_update_time":"2022-03-01T08:19:09.779365+08:00","cluster_id":1,"server_name":"192-168-10-219","service_name":"192-168-10-219:3306","host_ip":"192.168.10.219","id":1,"deployment_type":1,"version":"5.7","del_flag":0}]}"
// @Router	/api/v1/metadata/mysql-server/host-info [get]
func GetMySQLServerByHostInfo(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	hostIP, err := jsonparser.GetString(data, mysqlServerHostIPJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), mysqlServerHostIPJSON)
		return
	}
	portNum, err := jsonparser.GetInt(data, mysqlServerPortNumJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), mysqlServerPortNumJSON)
		return
	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// get entity
	err = s.GetByHostInfo(hostIP, int(portNum))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLServerByHostInfo, err, hostIP, portNum)
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
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLServerByHostInfo, hostIP, portNum)
}

// @Tags	mysql server
// @Summary	check if mysql server is a master node
// @Accept	application/json
// @Param	token		body string	true "token"
// @Param	host_ip		body string	true "host ip"
// @Param	port_num	body int	true "host port number"
// @Produce	application/json
// @Success	200 {string} string "{"host_ip":"192.168.1.2","port_num":"3306"}"
// @Router	/api/v1/metadata/mysql-server/is-master/host-info [get]
func IsMaster(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	hostIP, err := jsonparser.GetString(data, mysqlServerHostIPJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), mysqlServerHostIPJSON)
		return
	}
	portNum, err := jsonparser.GetInt(data, mysqlServerPortNumJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), mysqlServerPortNumJSON)
		return
	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// get entity
	isMaster, err := s.IsMaster(hostIP, int(portNum))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataIsMaster, err, hostIP, portNum)
		return
	}
	// response
	jsonStr := fmt.Sprintf(isMasterResponse, isMaster)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataIsMaster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataIsMaster, hostIP, portNum)
}

// @Tags	mysql server
// @Summary	get mysql cluster by id
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	id		body int	true "mysql server id"
// @Produce	application/json
// @Success	200 {string} string "{"mysql_servers":[{"server_name":"test","service_name":"test","host_ip":"192.168.1.1","port_num":3306,"del_flag":0,"create_time":"2022-03-02T01:26:32.107625+08:00","last_update_time":"2022-03-02T01:26:32.107625+08:00","id":26,"cluster_id":1,"deployment_type":1,"version":""}]}"
// @Router	/api/v1/metadata/mysql-server/mysql-cluster [get]
func GetMySQLClusterByMySQLServerID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, mysqlServerIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), mysqlServerIDJSON)
		return
	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// get entity
	err = s.GetMySQLClusterByID(int(id))
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

// @Tags	mysql server
// @Summary	add a new mysql server
// @Accept	application/json
// @Param	token			body string	true  "token"
// @Param	cluster_id		body int	true  "mysql cluster id"
// @Param	server_name		body string	true  "mysql server name"
// @Param	service_name	body string	false "mysql server service name"
// @Param	host_ip			body string	true  "mysql server host ip"
// @Param	port_num		body int	true  "mysql server port num"
// @Param	deployment_type	body int	true  "mysql deployment type"
// @Produce	application/json
// @Success	200 {string} string "{"mysql_servers":[{"deployment_type":1,"version":"","id":26,"server_name":"test","service_name":"test","port_num":3306,"del_flag":0,"create_time":"2022-03-02T01:26:32.107625+08:00","last_update_time":"2022-03-02T01:26:32.107625+08:00","cluster_id":97,"host_ip":"192.168.1.1"}]}"
// @Router	/api/v1/metadata/mysql-server [post]
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
				mysqlServerClusterIDJSON,
				mysqlServerServerNameJSON,
				mysqlServerHostIPJSON,
				mysqlServerPortNumJSON,
				mysqlServerDeploymentTypeJSON))
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

// @Tags	mysql server
// @Summary	update mysql server by id
// @Accept	application/json
// @Param	token			body string	true  "token"
// @Param	id				body int	true  "mysql server id"
// @Param	cluster_id		body int	false "mysql cluster id"
// @Param	server_name		body string	false "mysql server name"
// @Param	service_name	body string	false "mysql server service name"
// @Param	host_ip			body string	false "mysql server host ip"
// @Param	port_num		body int	false "mysql server port num"
// @Param	deployment_type	body int	false "mysql deployment type"
// @Produce	application/json
// @Success	200 {string} string "{"mysql_servers":[{"deployment_type":1,"version":"","id":26,"server_name":"test","service_name":"test","port_num":3306,"del_flag":0,"create_time":"2022-03-02T01:26:32.107625+08:00","last_update_time":"2022-03-02T01:26:32.107625+08:00","cluster_id":97,"host_ip":"192.168.1.1"}]}"
// @Router	/api/v1/metadata/mysql-server [post]
func UpdateMySQLServerByID(c *gin.Context) {
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
	idInterface, idExists := fields[mysqlServerIDStruct]
	if !idExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mysqlServerIDJSON)
		return
	}
	id, ok := idInterface.(int)
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, mysqlServerIDJSON)
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
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	id		body int	true "mysql server id"
// @Produce	application/json
// @Success	200 {string} string "{"mysql_servers":[{"id":1,"port_num":3306,"create_time":"2021-09-02T11:16:06.561525+08:00","last_update_time":"2022-03-02T01:14:14.13647+08:00","deployment_type":1,"version":"5.7","del_flag":0,"cluster_id":1,"server_name":"192-168-10-219","service_name":"192-168-10-219:3306","host_ip":"192.168.10.219"}]}"
// @Router /api/v1/metadata/mysql-server [get]
func DeleteMySQLServerByID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, mysqlServerIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), mysqlServerIDJSON)
		return
	}
	// init service
	s := metadata.NewMySQLServerServiceWithDefault()
	// insert into middleware
	err = s.Delete(int(id))
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
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteMySQLServer, id)
}
