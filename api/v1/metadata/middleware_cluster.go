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
	middlewareClusterIDJSON          = "id"
	middlewareClusterClusterNameJSON = "cluster_name"
	middlewareClusterEnvIDJSON       = "env_id"
	middlewareClusterUserIDJSON      = "user_id"
	middlewareClusterDelFlagJSON     = "del_flag"

	middlewareClusterIDStruct                = "ID"
	middlewareClusterClusterNameStruct       = "ClusterName"
	middlewareClusterEnvIDStruct             = "EnvID"
	middlewareClusterMiddlewareServersStruct = "MiddlewareServers"
	middlewareClusterUsersStruct             = "Users"
	middlewareClusterDelFlagStruct           = "DelFlag"
)

// @Tags	middleware cluster
// @Summary	get all middleware clusters
// @Accept	application/json
// @Param	token body string true "token"
// @Produce	application/json
// @Success	200 {string} string {"middleware_clusters":[{"id":1,"cluster_name":"middleware-cluster-1","env_id":1,"del_flag":0,"create_time":"2021-11-09T18:06:57.917596+08:00","last_update_time":"2021-11-18T15:39:52.927116+08:00"}]}
// @Router	/api/v1/metadata/middleware-cluster [get]
func GetMiddlewareCluster(c *gin.Context) {
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareClusterAll, err)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareClusterAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareClusterAll)
}

// @Tags	middleware cluster
// @Summary	get middleware cluster by env_id
// @Accept	application/json
// @Param	token 	body 	string 	true 	"token"
// @Param	env_id	body	int		true	"env id"
// @Produce	application/json
// @Success	200 {string} string {"middleware_clusters":[{"del_flag":0,"create_time":"2021-11-09T18:06:57.917596+08:00","last_update_time":"2021-11-18T15:39:52.927116+08:00","id":1,"cluster_name":"middleware-cluster-1","env_id":1}]}
// @Router	/api/v1/metadata/middleware-cluster/env [get]
func GetMiddlewareClusterByEnv(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	envID, err := jsonparser.GetInt(data, middlewareClusterEnvIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), middlewareClusterEnvIDJSON)
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entity
	err = s.GetByEnv(int(envID))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareClusterByEnv, err)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareClusterByEnv, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareClusterByEnv, envID)
}

// @Tags	middleware cluster
// @Summary	get middleware cluster by id
// @Accept	application/json
// @Param	token 	body 	string 	true 	"token"
// @Param	id		body	int		true	"middleware cluster id"
// @Produce	application/json
// @Success	200 {string} string {"middleware_clusters":[{"del_flag":0,"create_time":"2021-11-09T18:06:57.917596+08:00","last_update_time":"2021-11-18T15:39:52.927116+08:00","id":1,"cluster_name":"middleware-cluster-1","env_id":1}]}
// @Router	/api/v1/metadata/middleware-cluster/get [get]
func GetMiddlewareClusterByID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, middlewareClusterIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), middlewareClusterIDJSON)
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entity
	err = s.GetByID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareClusterByID, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareClusterByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareClusterByID, id)
}

// @Tags	middleware cluster
// @Summary	get middleware cluster by name
// @Accept	application/json
// @Param	token 			body 	string 	true 	"token"
// @Param	cluster_name	body	string	true	"middleware cluster name"
// @Produce	application/json
// @Success	200 {string} string {"middleware_clusters":[{"cluster_name":"middleware-cluster-1","env_id":1,"del_flag":0,"create_time":"2021-11-09T18:06:57.917596+08:00","last_update_time":"2021-11-18T15:39:52.927116+08:00","id":1}]}
// @Router	/api/v1/metadata/middleware-cluster/cluster-name [get]
func GetMiddlewareClusterByName(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	middlewareClusterName, err := jsonparser.GetString(data, middlewareClusterClusterNameJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), middlewareClusterClusterNameJSON)
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entity
	err = s.GetByName(middlewareClusterName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareClusterByName, err, middlewareClusterName)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareClusterByName, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareClusterByName, middlewareClusterName)
}

// @Tags	middleware cluster
// @Summary	get middleware servers by cluster id
// @Accept	application/json
// @Param	token 	body 	string 	true 	"token"
// @Param	id		body	int		true	"middleware cluster id"
// @Produce	application/json
// @Success	200 {string} string {"middleware_servers":[{"id":1,"cluster_id":1,"server_name":"update_middeware_server","middleware_role":1,"port_num":33061,"del_flag":0,"create_time":"2021-11-17T14:47:10.521279+08:00","last_update_time":"2022-03-02T10:18:52.564191+08:00","host_ip":"192.168.10.219"}]}
// @Router	/api/vi/metadata/middleware-server [get]
func GetMiddlewareServers(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, middlewareClusterIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), middlewareClusterIDJSON)
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// get entity
	err = s.GetMiddlewareServersByID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareServers, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(middlewareClusterMiddlewareServersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareServers, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareServers, id)

}

// @Tags	middleware cluster
// @Summary	get middleware servers by cluster id
// @Accept	application/json
// @Param	token 	body 	string 	true 	"token"
// @Param	id		body	int		true	"middleware cluster id"
// @Produce	application/json
// @Success	200 {string} string {"users":[{"id":1,"email":"allinemailtest@163.com","role":3,"del_flag":0,"last_update_time":"2021-11-22T13:46:20.430926+08:00","create_time":"2021-10-25T09:21:50.364327+08:00","user_name":"zhangsan","department_name":"arch","employee_id":"100001","account_name":"zs001","telephone":"01012345678","mobile":"13012345678"}]}
// @Router	/api/vi/metadata/middleware-cluster/users [get]
func GetUsersByMiddlewareClusterID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, middlewareClusterIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), middlewareClusterIDJSON)
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()

	// get entity
	err = s.GetUsersByMiddlewareClusterID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetUsersByMiddlewareClusterID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(middlewareClusterUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetUsersByMiddlewareClusterID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetUsersByMiddlewareClusterID, id)
}

// @Tags	middleware cluster
// @Summary	add a new middleware cluster
// @Accept	application/json
// @Param	token 			body 	string 	true 	"token"
// @Param	cluster_name	body	string	true	"middleware cluster name"
// @Param	env_id			body	int    	true 	"env id"
// @Produce	application/json
// @Success	200 {string} string {"middleware_clusters":[{"id":65,"cluster_name":"new_middleware_cluster","env_id":1,"del_flag":0,"create_time":"2022-03-02T10:39:06.206145+08:00","last_update_time":"2022-03-02T10:39:06.206145+08:00"}]}
// @Router	/api/v1/metadata/middleware-cluster [post]
func AddMiddlewareCluster(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err)
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MiddlewareClusterInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, ok := fields[middlewareClusterClusterNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterClusterNameJSON)
		return
	}
	_, ok = fields[middlewareClusterEnvIDStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterClusterNameJSON)
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddMiddlewareCluster, err, fields[middlewareClusterClusterNameStruct])
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddMiddlewareCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddMiddlewareCluster, fields[middlewareClusterClusterNameStruct])
}

// @Tags	middleware cluster
// @Summary	update middleware cluster by id
// @Accept	application/json
// @Param	token 			body 	string 	true 	"token"
// @Param	id				body	int		true	"middleware cluster id"
// @Param	cluster_name	body	string	false	"middleware cluster name"
// @Param	env_id			body	int		false	"env id"
// @Param	del_flag		body	int		false	"delete flag"
// @Produce	application/json
// @Success	200 {string} string {"middleware_clusters":[{"last_update_time":"2021-11-18T15:39:52.927116+08:00","id":1,"cluster_name":"update_middleware_cluster","env_id":1,"del_flag":0,"create_time":"2021-11-09T18:06:57.917596+08:00"}]}
// @Router	/api/v1/metadata/middleware-cluster/update [post]
func UpdateMiddlewareClusterByID(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.MiddlewareClusterInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	idInterface, idExists := fields[middlewareClusterIDStruct]
	if !idExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, middlewareClusterIDJSON)
		return
	}
	id, ok := idInterface.(int)
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, middlewareClusterIDJSON)
		return
	}
	_, middlewareClusterNameExists := fields[middlewareClusterClusterNameStruct]
	_, middlewareClusterEnvIDExists := fields[middlewareClusterEnvIDStruct]
	_, middlewareClusterDelFlagExists := fields[middlewareClusterDelFlagStruct]
	if !middlewareClusterNameExists && !middlewareClusterEnvIDExists && !middlewareClusterDelFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s %s and %s", middlewareClusterClusterNameJSON, middlewareClusterEnvIDJSON, middlewareClusterDelFlagJSON))
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateMiddlewareCluster, err)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateMiddlewareCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataUpdateMiddlewareCluster, id)
}

// @Tags	middleware cluster
// @Summary	delete middleware cluster by id
// @Accept	application/json
// @Param	token 	body	string 	true 	"token"
// @Param	id		body	int		true	"middleware cluster id"
// @Produce	application/json
// @Success	200 {string} string {"middleware_clusters":[{"id":65,"cluster_name":"new_middleware_cluster","env_id":1,"del_flag":0,"create_time":"2022-03-02T10:39:06.206145+08:00","last_update_time":"2022-03-02T10:39:06.206145+08:00"}]}
// @Router	/api/v1/metadata/app/delete [post]
func DeleteMiddlewareClusterByID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, middlewareClusterIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), middlewareClusterIDJSON)
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// update entities
	err = s.Delete(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteMiddlewareCluster, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteMiddlewareCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteMiddlewareCluster, id)
}

// @Tags	middleware cluster
// @Summary	add user map
// @Accept	application/json
// @Param	token 	body 	string 	true 	"token"
// @Param	id		body	int		true	"middleware cluster id"
// @Param	user_id	body	int		true	"user id"
// @Produce	application/json
// @Success	200 {string} string {"users":[{"account_name":"zs001","email":"allinemailtest@163.com","mobile":"13012345678","del_flag":0,"create_time":"2021-10-25T09:21:50.364327+08:00","user_name":"zhangsan","department_name":"arch","employee_id":"100001","telephone":"01012345678","role":3,"last_update_time":"2021-11-22T13:46:20.430926+08:00","id":1}]}
// @Router	/api/v1/metadata/middleware-cluster/add-user [post]
func MiddlewareClusterAddUser(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}

	id, err := jsonparser.GetInt(data, middlewareClusterIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), middlewareClusterIDJSON)
		return
	}
	userID, err := jsonparser.GetInt(data, middlewareClusterUserIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), middlewareClusterUserIDJSON)
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// update entities
	err = s.AddUser(int(id), int(userID))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataMiddlewareClusterAddUser, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(middlewareClusterUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataMiddlewareClusterAddUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataMiddlewareClusterAddUser, id, userID)
}

// @Tags	middleware cluster
// @Summary	delete user map
// @Accept	application/json
// @Param	token 	body 	string 	true 	"token"
// @Param	id		body	int		true	"middleware cluster id"
// @Param	user_id	body	int		true	"user id"
// @Produce	application/json
// @Success	200 {string} string {"users":[]}
// @Router	/api/v1/metadata/middleware-cluster/delete-user [post]
func MiddlewareClusterDeleteUser(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, middlewareClusterIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), middlewareClusterIDJSON)
		return
	}
	userID, err := jsonparser.GetInt(data, middlewareClusterUserIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), middlewareClusterUserIDJSON)
		return
	}
	// init service
	s := metadata.NewMiddlewareClusterServiceWithDefault()
	// update entities
	err = s.DeleteUser(int(id), int(userID))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataMiddlewareClusterDeleteUser, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(middlewareClusterUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataMiddlewareClusterDeleteUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataMiddlewareClusterDeleteUser, id, userID)
}
