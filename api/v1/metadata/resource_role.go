package metadata

import (
	"fmt"

	"github.com/buger/jsonparser"
	"github.com/pingcap/errors"

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
	resourceRoleIDJSON              = "id"
	resourceRoleRoleUUIDJSON        = "role_uuid"
	resourceRoleResourceGroupIDJSON = "resource_group_id"
	resourceRoleUserIDJSON          = "user_id"
	resourceRoleDelFlagJSON         = "del_flag"

	resourceRoleUUIDStruct            = "RoleUUID"
	resourceRoleResourceGroupIDStruct = "ResourceGroupID"

	resourceRoleIDStruct            = "ID"
	resourceRoleResourceGroupStruct = "ResourceGroup"
	resourceRoleUsersStruct         = "Users"
)

// @Tags mysql cluster
// @Summary	get all mysql clusters
// @Accept	application/json
// @Produce	application/json
// @Param	token	body string	true "token"
// @Success	200 {string} string "{"mysql_clusters":[{"middleware_cluster_id":1,"monitor_system_id":1,"env_id":1,"del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","last_update_time":"2021-02-23T20:57:24.603009+08:00","id":1,"cluster_name":"cluster_name_init"}]}"
// @Router	/api/v1/metadata/mysql-cluster [get]
func GetResourceRole(c *gin.Context) {
	// init service
	s := metadata.NewResourceRoleServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetResourceRoleAll, err)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	fmt.Println("ok")
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetResourceRoleAll, jsonBytes).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetResourceRoleAll)
}

// @Tags	mysql cluster
// @Summary	get mysql cluster by id
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	id		body int	true "mysql cluster id"
// @Produce	application/json
// @Success	200 {string} string "{"mysql_clusters":[{"del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","id":1,"monitor_system_id":1,"env_id":1,"last_update_time":"2021-02-23T20:57:24.603009+08:00","cluster_name":"cluster_name_init","middleware_cluster_id":1}]}"
// @Router	/api/v1/metadata/mysql-cluster/get [get]
func GetResourceRoleByID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, resourceRoleIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceRoleIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceRoleServiceWithDefault()
	// get entity
	err = s.GetByID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetResourceRoleByID, id, err)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetResourceRoleByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetResourceRoleByID, id)
}

// @Tags mysql cluster
// @Summary get mysql cluster by uuid
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	name	body string	true "mysql cluster name"
// @Produce  application/json
// @Success 200 {string} string "{"mysql_clusters":[{"del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","id":1,"monitor_system_id":1,"env_id":1,"last_update_time":"2021-02-23T20:57:24.603009+08:00","cluster_name":"cluster_name_init","middleware_cluster_id":1}]}"
// @Router /api/v1/metadata/mysql-cluster/cluster-name [get]
func GetResourceRoleByUUID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	roleUUID, err := jsonparser.GetString(data, resourceRoleRoleUUIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceRoleRoleUUIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceRoleServiceWithDefault()
	// get entity
	err = s.GetByRoleUUID(roleUUID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetResourceRoleByUUID, err, roleUUID)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetResourceRoleByUUID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetResourceRoleByUUID, roleUUID)
}

// @Tags	mysql cluster
// @Summary	get mysql servers by id
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	id		body int	true "mysql cluster id"
// @Produce	application/json
// @Success	200 {string} string "{"mysql_servers":[{"id":1,"server_name":"192-168-10-219","service_name":"192-168-10-219:3306","deployment_type":1,"last_update_time":"2021-12-21T09:16:20.184065+08:00","cluster_id":1,"host_ip":"192.168.10.219","port_num":3306,"version":"5.7","del_flag":0,"create_time":"2021-09-02T11:16:06.561525+08:00"}]}"
// @Router	/api/v1/metadata/mysql-cluster/mysql-server [get]
func GetResourceGroupByResourceRoleID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, resourceRoleIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceRoleIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceRoleServiceWithDefault()
	// get entity
	err = s.GetResourceGroupByID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLServers, id, err)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceRoleResourceGroupStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)

	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLServers, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLServers, id)
}

// @Tags	mysql cluster
// @Summary	get mysql cluster users
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	id		body int	true "mysql cluster id"
// @Produce	application/json
// @Success	200 {string} string "{"users":[{"id":1,"employee_id":"100001","account_name":"zs001","last_update_time":"2021-11-22T13:46:20.430926+08:00","mobile":"13012345678","role":3,"del_flag":0,"user_name":"zhangsan","department_name":"arch","email":"allinemailtest@163.com","telephone":"01012345678","create_time":"2021-10-25T09:21:50.364327+08:00"}]}"
// @Router	/api/v1/metadata/mysql-cluster/user [get]
func GetUsersByResourceRoleID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, resourceRoleIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceRoleIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceRoleServiceWithDefault()
	// get entity
	err = s.GetUsersByID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAppUsers, id, err)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceRoleUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAppUsers, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAppUsers, id)
}

// @Tags	mysql cluster
// @Summary	get app users
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	id		body int	true "mysql cluster id"
// @Produce	application/json
// @Success	200 {string} string "{"users":[{"id":1,"employee_id":"100001","account_name":"zs001","last_update_time":"2021-11-22T13:46:20.430926+08:00","mobile":"13012345678","role":3,"del_flag":0,"user_name":"zhangsan","department_name":"arch","email":"allinemailtest@163.com","telephone":"01012345678","create_time":"2021-10-25T09:21:50.364327+08:00"}]}"
// @Router	/api/v1/metadata/mysql-cluster/app-user [get]
func GetUsersByResourceRoleUUID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	roleUUID, err := jsonparser.GetString(data, resourceRoleRoleUUIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceRoleIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceRoleServiceWithDefault()
	// get entity
	err = s.GetUsersByRoleUUID(roleUUID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAppUsers, err, roleUUID)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceRoleUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAppUsers, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAppUsers, roleUUID)
}

// @Tags	mysql cluster
// @Summary	add user map
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	id		body int	true "mysql cluster id"
// @Param	user_id	body int	true "user id"
// @Produce	application/json
// @Success	200 {string} string "{"users":[{"id":1,"employee_id":"100001","email":"allinemailtest@163.com","role":3,"del_flag":0,"user_name":"zhangsan","department_name":"arch","account_name":"zs001","telephone":"01012345678","mobile":"13012345678","create_time":"2021-10-25T09:21:50.364327+08:00","last_update_time":"2021-11-22T13:46:20.430926+08:00"}}"
// @Router	/api/v1/metadata/mysql-cluster/add-user [post]
func ResourceRoleAddUser(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, resourceRoleIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceRoleIDJSON)
		return
	}
	userID, err := jsonparser.GetInt(data, resourceRoleUserIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceRoleUserIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceRoleServiceWithDefault()
	// update entities
	err = s.AddUser(int(id), int(userID))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataResourceRoleAddUser, err, id, userID)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceRoleUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, errors.Trace(err))
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataResourceRoleAddUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataResourceRoleAddUser, id, userID)
}

// @Tags	mysql cluster
// @Summary	delete user map
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	id		body int	true "mysql cluster id"
// @Param	user_id	body int	true "user id"
// @Produce	application/json
// @Success	200 {string} string "{"users":[{"id":1,"employee_id":"100001","email":"allinemailtest@163.com","role":3,"del_flag":0,"user_name":"zhangsan","department_name":"arch","account_name":"zs001","telephone":"01012345678","mobile":"13012345678","create_time":"2021-10-25T09:21:50.364327+08:00","last_update_time":"2021-11-22T13:46:20.430926+08:00"}]}"
// @Router	/api/v1/metadata/mysql-cluster/delete-user [post]
func ResourceRoleDeleteUser(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, resourceRoleIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceRoleIDJSON)
		return
	}
	userID, err := jsonparser.GetInt(data, resourceRoleUserIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceRoleUserIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceRoleServiceWithDefault()
	// update entities
	err = s.DeleteUser(int(id), int(userID))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataResourceRoleDeleteUser, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceRoleUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataResourceRoleDeleteUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataResourceRoleDeleteUser, id, userID)
}

// @Tags	mysql cluster
// @Summary	add a new mysql cluster
// @Accept	application/json
// @Param	token					body string	true "token"
// @Param	cluster_name			body string	true  "mysql cluster name"
// @Param	middleware_cluster_id	body int	false "middleware cluster id"
// @Param	monitor_system_id		body int	false "monitor system id"
// @Param	env_id					body string	true  "env id"
// @Produce	application/json
// @Success	200 {string} string "{"mysql_clusters":[{"id":3,"cluster_name":"api_test","monitor_system_id":0,"env_id":1,"create_time":"2022-03-01T08:30:43.428343+08:00","middleware_cluster_id":0,"del_flag":0,"last_update_time":"2022-03-01T08:30:43.428343+08:00"}]}"
// @Router	/api/v1/metadata/mysql-cluster [post]
func AddResourceRole(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.ResourceRoleInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, roleUUIDExists := fields[resourceRoleUUIDStruct]
	if !roleUUIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, resourceRoleRoleUUIDJSON)
		return
	}
	_, resourceGroupIDExists := fields[resourceRoleResourceGroupIDStruct]
	if !resourceGroupIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, resourceRoleResourceGroupIDJSON)
		return
	}

	// init service
	s := metadata.NewResourceRoleServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddResourceRole,
			fields[resourceRoleUUIDStruct],
			fields[resourceRoleResourceGroupIDStruct],
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddResourceRole, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddResourceRole,
		fields[resourceRoleUUIDStruct],
		fields[resourceRoleResourceGroupIDStruct],
	)
}

// @Tags	mysql cluster
// @Summary	update mysql cluster by id
// @Accept	application/json
// @Param	token					body string	true  "token"
// @Param	id						body int	true  "mysql cluster id"
// @Param	cluster_name			body string	false "mysql cluster name"
// @Param	middleware_cluster_id	body int	false "middleware cluster id"
// @Param	monitor_system_id		body int	false "monitor system id"
// @Param	env_id					body string	false "env id"
// @Param	del_flag				body int	false "delete flag"
// @Produce	application/json
// @Success	200 {string} string "{"mysql_clusters":[{"middleware_cluster_id":0,"monitor_system_id":0,"last_update_time":"2022-03-01T08:30:43.428343+08:00","id":3,"cluster_name":"test","env_id":1,"del_flag":0,"create_time":"2022-03-01T08:30:43.428343+08:00"}]}"
// @Router	/api/v1/metadata/mysql-cluster/update [post]
func UpdateResourceRoleByID(c *gin.Context) {
	var fields map[string]interface{}
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.ResourceRoleInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	idInterface, idExists := fields[resourceRoleIDStruct]
	if !idExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, resourceRoleIDJSON)
		return
	}
	id, ok := idInterface.(int)
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, resourceRoleIDJSON)
		return
	}
	_, roleUUIDExists := fields[resourceRoleUUIDStruct]
	_, roleResourceGroupIDExists := fields[resourceRoleResourceGroupIDStruct]
	_, delFlagExists := fields[envDelFlagStruct]
	if !roleUUIDExists &&
		!roleResourceGroupIDExists &&
		!delFlagExists {
		resp.ResponseNOK(
			c, message.ErrFieldNotExists,
			fmt.Sprintf("%s, %s and %s",
				resourceRoleRoleUUIDJSON,
				resourceRoleResourceGroupIDJSON,
				resourceRoleDelFlagJSON))
		return
	}
	// init service
	s := metadata.NewResourceRoleServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateResourceRole, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateResourceRole, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.DebugMetadataUpdateResourceRole, fields[resourceRoleUUIDStruct])
}

// @Tags	mysql cluster
// @Summary	update mysql cluster by id
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	id		body int	true "mysql cluster id"
// @Produce	application/json
// @Success	200 {string} string "{"mysql_clusters":[{"cluster_name":"test","env_id":1,"del_flag":0,"create_time":"2022-03-01T08:30:43.428343+08:00","last_update_time":"2022-03-01T08:32:25.715563+08:00","id":3,"middleware_cluster_id":0,"monitor_system_id":0}]}"
// @Router	/api/v1/metadata/mysql-cluster/delete [post]
func DeleteResourceRoleByID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, resourceRoleIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceRoleIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceRoleServiceWithDefault()
	// insert into middleware
	err = s.Delete(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteResourceRole,
			err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, errors.Trace(err))
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteResourceRole, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteResourceRole, id)
}
