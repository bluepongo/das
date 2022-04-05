package metadata

import (
	"fmt"
	"strconv"

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
	middlewareServerIDJSON        = "id"
	middlewareServerClusterIDJSON = "cluster_id"

	middlewareServerClusterIDStruct      = "ClusterID"
	middlewareServerNameStruct           = "ServerName"
	middlewareServerMiddlewareRoleStruct = "MiddlewareRole"
	middlewareServerHostIPStruct         = "HostIP"
	middlewareServerPortNumStruct        = "PortNum"
)

// @Tags	middleware server
// @Summary	get all middleware servers
// @Accept	application/json
// @Produce	application/json
// @Success	200 {string} string {"middleware_servers":[{"last_update_time":"2021-11-18T15:54:10.599097+08:00","cluster_id":1,"server_name":"middleware-server-1","port_num":33061,"del_flag":0,"create_time":"2021-11-17T14:47:10.521279+08:00","id":1,"middleware_role":1,"host_ip":"192.168.10.219"}]}
// @Router	/api/v1/metadata/middleware-server [get]
func GetMiddlewareServer(c *gin.Context) {
	// init service
	s := metadata.NewMiddlewareServerServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareServerAll, err)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareServerAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareServerAll)
}

// @Tags     middleware server
// @Summary	get middleware servers by cluster id
// @Accept	application/json
// @Param	cluster_id	path	int	true	"middleware cluster id"
// @Produce	application/json
// @Success 200 {string} string {"middleware_servers":[{"id":1,"cluster_id":1,"server_name":"middleware-server-1","host_ip":"192.168.10.219","port_num":33061,"del_flag":0,"create_time":"2021-11-17T14:47:10.521279+08:00","last_update_time":"2021-11-18T15:54:10.599097+08:00","middleware_role":1}]}
// @Router	/api/v1/metadata/middleware-server/cluster-id/:cluster_id [get]
func GetMiddlewareServerByClusterID(c *gin.Context) {
	// get param
	clusterIDStr := c.Param(middlewareServerClusterIDJSON)
	if clusterIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerClusterIDJSON)
		return
	}
	clusterID, err := strconv.Atoi(clusterIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewMiddlewareServerServiceWithDefault()
	// get entity
	err = s.GetByClusterID(clusterID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareSeverByClusterID, err)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareSeverByClusterID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareSeverByClusterID)
}

// @Tags	middleware server
// @Summary	get middleware server by id
// @Accept	application/json
// @Param	id	path	int	true	"middleware server id"
// @Produce	application/json
// @Success	200 {string} string {"middleware_servers":[{"last_update_time":"2021-11-18T15:54:10.599097+08:00","id":1,"server_name":"middleware-server-1","middleware_role":1,"port_num":33061,"cluster_id":1,"host_ip":"192.168.10.219","del_flag":0,"create_time":"2021-11-17T14:47:10.521279+08:00"}]}
// @Router	/api/v1/metadata/middleware-server/get/:id [get]
func GetMiddlewareServerByID(c *gin.Context) {
	// get param
	idStr := c.Param(middlewareServerIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewMiddlewareServerServiceWithDefault()
	// get entity
	err = s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareServerByID, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareServerByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareServerByID, id)
}

// @Tags	middleware server
// @Summary	get middleware server by host info
// @Accept	application/json
// @Param	host_ip		body	string	true	"host ip"
// @Param	port_num	body	int		true	"port number"
// @Produce	application/json
// @Success	200 {string} string {"middleware_servers":[{"middleware_role":1,"host_ip":"192.168.10.219","del_flag":0,"last_update_time":"2021-11-18T15:54:10.599097+08:00","id":1,"server_name":"middleware-server-1","cluster_id":1,"port_num":33061,"create_time":"2021-11-17T14:47:10.521279+08:00"}]}
// @Router	/api/v1/metadata/middleware-server/host-info [get]
func GetMiddlewareServerByHostInfo(c *gin.Context) {
	var rd *utilmeta.HostInfo
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}

	// init service
	s := metadata.NewMiddlewareServerServiceWithDefault()
	// get entity
	err = s.GetByHostInfo(rd.GetHostIP(), rd.GetPortNum())
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareServerByHostInfo, rd.GetHostIP(), rd.GetPortNum())
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareServerByHostInfo, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareServerByHostInfo, rd.GetHostIP(), rd.GetPortNum())
}

// @Tags	middleware server
// @Summary	add a new middleware server
// @Accept	application/json
// @Param	cluster_id		body	int		true	"middleware cluster id"
// @Param	server_name		body	string	true	"middleware server name"
// @Param	middleware_role	body	int		true	"middleware role"
// @Param	host_ip			body	string	true	"host ip"
// @Param	port_num		body	int		true	"port number"
// @Produce	application/json
// @Success	200 {string} string {"middleware_servers":[{"host_ip":"192.168.10.219","del_flag":0,"create_time":"2022-03-02T10:18:28.021994+08:00","last_update_time":"2022-03-02T10:18:28.021994+08:00","cluster_id":1,"server_name":"new_middleware_server","middleware_role":1,"id":42,"port_num":33062}]}
// @Router	/api/v1/metadata/middleware-server [post]
func AddMiddlewareServer(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MiddlewareServerInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, ok := fields[middlewareServerClusterIDStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerClusterIDStruct)
		return
	}
	_, ok = fields[middlewareServerNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerNameStruct)
		return
	}
	_, ok = fields[middlewareServerMiddlewareRoleStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerMiddlewareRoleStruct)
		return
	}
	_, ok = fields[middlewareServerHostIPStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerHostIPStruct)
		return
	}
	_, ok = fields[middlewareServerPortNumStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerPortNumStruct)
		return
	}
	// init service
	s := metadata.NewMiddlewareServerServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddMiddlewareServer, err, fields[middlewareServerNameStruct])
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddMiddlewareServer, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddMiddlewareServer, fields[middlewareServerNameStruct])
}

// @Tags	middleware server
// @Summary	update middleware server by id
// @Accept	application/json
// @Param	cluster_id		body	int		true	"middleware cluster id"
// @Param	server_name		body	string	false	"middleware server name"
// @Param	middleware_role	body	int		false	"middleware role"
// @Param	host_ip			body	string	false	"host ip"
// @Param	port_num		body	int		false	"port number"
// @Param	del_flag		body	int		false	"delete flag"
// @Produce	application/json
// @Success	200 {string} string {"middleware_servers":[{"del_flag":0,"server_name":"update_middeware_server","host_ip":"192.168.10.219","port_num":33061,"middleware_role":1,"create_time":"2021-11-17T14:47:10.521279+08:00","last_update_time":"2021-11-18T15:54:10.599097+08:00","id":1,"cluster_id":1}]}
// @Router	/api/v1/metadata/middleware-server/update/:id [post]
func UpdateMiddlewareServerByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(middlewareServerIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerIDJSON)
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
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MiddlewareServerInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, middlewareServerClusterIDExists := fields[middlewareServerClusterIDStruct]
	_, middlewareServerNameExists := fields[middlewareServerNameStruct]
	_, middlewareServerMiddlewareRoleExists := fields[middlewareServerMiddlewareRoleStruct]
	_, middlewareServerHostIPExists := fields[middlewareServerHostIPStruct]
	_, middlewareServerPortNumExists := fields[middlewareServerPortNumStruct]
	_, delFlagExists := fields[envDelFlagStruct]
	if !middlewareServerClusterIDExists && !middlewareServerNameExists && !middlewareServerMiddlewareRoleExists && !middlewareServerHostIPExists && !middlewareServerPortNumExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s", middlewareServerNameStruct, envDelFlagStruct))
		return
	}
	// init service
	s := metadata.NewMiddlewareServerServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateMiddlewareServer, err)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err, id)
		return
	}
	// resp
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateMiddlewareServer, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataUpdateMiddlewareServer, id)
}

// @Tags	middleware server
// @Summary	delete middleware server by id
// @Accept	application/json
// @Param	id	path	int	true	"middleware server id"
// @Produce	application/json
// @Success	200 {string} string {"middleware_servers":[{"server_name":"new_middleware_server","host_ip":"192.168.10.219","port_num":33062,"create_time":"2022-03-02T10:18:28.021994+08:00","last_update_time":"2022-03-02T10:18:28.021994+08:00","id":42,"cluster_id":1,"middleware_role":1,"del_flag":0}]}
// @Router	/api/v1/metadata/middleware-server/delete/:id [post]
func DeleteMiddlewareServerByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(middlewareServerIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareServerIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewMiddlewareServerServiceWithDefault()
	// update entities
	err = s.Delete(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteMiddlewareServer, err, fields[middlewareClusterClusterNameStruct])
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteMiddlewareServer, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteMiddlewareServer, fields[middlewareClusterClusterNameStruct])
}
