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
	resourceGroupIDJSON                  = "id"
	resourceGroupGroupUUIDJSON           = "group_uuid"
	resourceGroupGroupNameJSON           = "group_name"
	resourceGroupDelFlagJSON             = "del_flag"
	resourceGroupMySQLClusterIDJSON      = "mysql_cluster_id"
	resourceGroupMiddlewareClusterIDJSON = "middleware_cluster_id"

	resourceGroupIDStruct                 = "ID"
	resourceGroupGroupUUIDStruct          = "GroupUUID"
	resourceGroupGroupNameStruct          = "GroupName"
	resourceGroupResourceRolesStruct      = "ResourceRoles"
	resourceGroupMySQLClustersStruct      = "MySQLClusters"
	resourceGroupMySQLServersStruct       = "MySQLServers"
	resourceGroupMiddlewareClustersStruct = "MiddlewareClusters"
	resourceGroupMiddlewareServersStruct  = "MiddlewareServers"
	resourceGroupUsersStruct              = "Users"
	resourceGroupDelFlagStruct            = "DelFlag"
)

// @Tags	resource group
// @Summary	get all resource groups
// @Accept	application/json
// @Param	token body string true "token"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group [get]
func GetResourceGroup(c *gin.Context) {
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetResourceGroupAll, err)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetResourceGroupAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetResourceGroupAll)
}

// @Tags	resource group
// @Summary	get resource group by id
// @Accept	application/json
// @Param	token 	body 	string 	true 	"token"
// @Param	id		body	int		true	"resource group id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group/get [get]
func GetResourceGroupByID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, resourceGroupIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// get entity
	err = s.GetByID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetResourceGroupByID, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetResourceGroupByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetResourceGroupByID, id)
}

// @Tags	resource group
// @Summary	get resource group by resource group uuid
// @Accept	application/json
// @Param	token 	    body 	string 	true 	"token"
// @Param	group_uuid	body	int		true	"resource group uuid"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group/group-uuid [get]
func GetResourceGroupByGroupUUID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	groupUUID, err := jsonparser.GetString(data, resourceGroupGroupUUIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupGroupUUIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// get entity
	err = s.GetByGroupUUID(groupUUID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetResourceGroupByGroupUUID, err, groupUUID)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetResourceGroupByGroupUUID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetResourceGroupByGroupUUID, groupUUID)
}

// @Tags	resource group
// @Summary	get resource roles by id
// @Accept	application/json
// @Param	token 	body 	string 	true 	"token"
// @Param	id		body	string	true	"resource group id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group/resource-role/id [get]
func GetResourceRolesByGroupID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, resourceGroupIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// get entity
	err = s.GetResourceRolesByID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetResourceRolesByGroupID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupResourceRolesStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetResourceRolesByGroupID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetResourceRolesByGroupID, id)
}

// @Tags	resource group
// @Summary	get mysql clusters by id
// @Accept	application/json
// @Param	token 	body 	string 	true 	"token"
// @Param	id		body	string	true	"resource group id"
// @Produce	application/json
// @Success	200 {string} string "{"mysql_clusters":[{"middleware_cluster_id":1,"monitor_system_id":1,"env_id":1,"del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","last_update_time":"2021-02-23T20:57:24.603009+08:00","id":1,"cluster_name":"cluster_name_init"}]}"
// @Router	/api/v1/metadata/resource-group/mysql-cluster/id [get]
func GetMySQLClustersByGroupID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, resourceGroupIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// get entity
	err = s.GetMySQLClustersByID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLClustersByGroupID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupMySQLClustersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLClustersByGroupID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLClustersByGroupID, id)
}

// @Tags	resource group
// @Summary	get mysql servers by id
// @Accept	application/json
// @Param	token 	body 	string 	true 	"token"
// @Param	id		body	string	true	"resource group id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group/mysql-server/id [get]
func GetMySQLServersByGroupID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, resourceGroupIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// get entity
	err = s.GetMySQLServersByID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLServersByGroupID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupMySQLServersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLServersByGroupID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLServersByGroupID, id)
}

// @Tags	resource group
// @Summary	get middleware clusters by id
// @Accept	application/json
// @Param	token 	body 	string 	true 	"token"
// @Param	id		body	string	true	"resource group id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group/middleware-cluster/id [get]
func GetMiddlewareClustersByGroupID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, resourceGroupIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// get entity
	err = s.GetMiddlewareClustersByID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareClustersByGroupID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupMiddlewareClustersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareClustersByGroupID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareClustersByGroupID, id)
}

// @Tags	resource group
// @Summary	get middleware servers by id
// @Accept	application/json
// @Param	token 	body 	string 	true 	"token"
// @Param	id		body	string	true	"resource group id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group/middleware-server/id [get]
func GetMiddlewareServersByGroupID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, resourceGroupIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// get entity
	err = s.GetMiddlewareServersByID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareServersByGroupID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupMiddlewareServersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareServersByGroupID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareServersByGroupID, id)
}

// @Tags	resource group
// @Summary	get users by id
// @Accept	application/json
// @Param	token 	body 	string 	true 	"token"
// @Param	id		body	int		true	"resource group id"
// @Produce	application/json
// @Success	200 {string} string {"users":[{"id":1,"email":"allinemailtest@163.com","role":3,"del_flag":0,"last_update_time":"2021-11-22T13:46:20.430926+08:00","create_time":"2021-10-25T09:21:50.364327+08:00","user_name":"zhangsan","department_name":"arch","employee_id":"100001","account_name":"zs001","telephone":"01012345678","mobile":"13012345678"}]}
// @Router	/api/vi/metadata/resource-group/user/id [get]
func GetUsersByGroupID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, resourceGroupIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()

	// get entity
	err = s.GetUsersByID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetUsersByGroupID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetUsersByGroupID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetUsersByGroupID, id)
}

// @Tags	resource group
// @Summary	get users by id
// @Accept	application/json
// @Param	token 	body 	string 	true 	"token"
// @Param	id		body	int		true	"resource group id"
// @Produce	application/json
// @Success	200 {string} string {"users":[{"id":1,"email":"allinemailtest@163.com","role":3,"del_flag":0,"last_update_time":"2021-11-22T13:46:20.430926+08:00","create_time":"2021-10-25T09:21:50.364327+08:00","user_name":"zhangsan","department_name":"arch","employee_id":"100001","account_name":"zs001","telephone":"01012345678","mobile":"13012345678"}]}
// @Router	/api/vi/metadata/resource-group/das-admin/id [get]
func GetDASAdminUsersByGroupID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, resourceGroupIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()

	// get entity
	err = s.GetDASAdminUsersByID(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDASAdminUsersByGroupID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDASAdminUsersByGroupID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDASAdminUsersByGroupID, id)
}

// @Tags	resource group
// @Summary	get resource roles by resource group uuid
// @Accept	application/json
// @Param	token 		body 	string 	true	"token"
// @Param	group_uuid	body	string	true	"resource group uuid"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group/resource-role/group-uuid [get]
func GetResourceRolesByGroupUUID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	groupUUID, err := jsonparser.GetString(data, resourceGroupGroupUUIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupGroupUUIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// get entity
	err = s.GetResourceRolesByGroupUUID(groupUUID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetResourceRolesByGroupUUID, err, groupUUID)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupResourceRolesStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetResourceRolesByGroupUUID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetResourceRolesByGroupUUID, groupUUID)
}

// @Tags	resource group
// @Summary	get mysql clusters by resource group uuid
// @Accept	application/json
// @Param	token 		body 	string 	true 	"token"
// @Param	group_uuid	body	string	true	"resource group uuid"
// @Produce	application/json
// @Success	200 {string} string "{"mysql_clusters":[{"middleware_cluster_id":1,"monitor_system_id":1,"env_id":1,"del_flag":0,"create_time":"2021-02-23T20:57:24.603009+08:00","last_update_time":"2021-02-23T20:57:24.603009+08:00","id":1,"cluster_name":"cluster_name_init"}]}"
// @Router	/api/v1/metadata/resource-group/mysql-cluster/group-uuid [get]
func GetMySQLClustersByGroupUUID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	groupUUID, err := jsonparser.GetString(data, resourceGroupGroupUUIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupGroupUUIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// get entity
	err = s.GetMySQLClustersByGroupUUID(groupUUID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLClustersByGroupUUID, err, groupUUID)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupMySQLClustersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLClustersByGroupUUID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLClustersByGroupUUID, groupUUID)
}

// @Tags	resource group
// @Summary	get mysql servers by resource group uuid
// @Accept	application/json
// @Param	token 		body 	string 	true 	"token"
// @Param	group_uuid	body	string	true	"resource group uuid"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group/mysql-server/group-uuid [get]
func GetMySQLServersByGroupUUID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	groupUUID, err := jsonparser.GetString(data, resourceGroupGroupUUIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupGroupUUIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// get entity
	err = s.GetMySQLServersByGroupUUID(groupUUID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLServersByGroupUUID, err, groupUUID)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupMySQLServersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLServersByGroupUUID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLServersByGroupUUID, groupUUID)
}

// @Tags	resource group
// @Summary	get middleware clusters by resource group uuid
// @Accept	application/json
// @Param	token 		body 	string 	true 	"token"
// @Param	group_uuid	body	string	true	"resource group uuid"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group/middleware-cluster/group-uuid [get]
func GetMiddlewareClustersByGroupUUID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	groupUUID, err := jsonparser.GetString(data, resourceGroupGroupUUIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupGroupUUIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// get entity
	err = s.GetMiddlewareClustersByGroupUUID(groupUUID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareClustersByGroupUUID, err, groupUUID)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupMiddlewareClustersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareClustersByGroupUUID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareClustersByGroupUUID, groupUUID)
}

// @Tags	resource group
// @Summary	get middleware servers by resource group uuid
// @Accept	application/json
// @Param	token 		body 	string 	true 	"token"
// @Param	group_uuid	body	string	true	"resource group uuid"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group/middleware-server/group-uuid [get]
func GetMiddlewareServersByGroupUUID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	groupUUID, err := jsonparser.GetString(data, resourceGroupGroupUUIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupGroupUUIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// get entity
	err = s.GetMiddlewareServersByGroupUUID(groupUUID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareServersByGroupUUID, err, groupUUID)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupMiddlewareServersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareServersByGroupUUID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareServersByGroupUUID, groupUUID)
}

// @Tags	resource group
// @Summary	get users by resource group uuid
// @Accept	application/json
// @Param	token 		body 	string 	true 	"token"
// @Param	group_uuid	body	string	true	"resource group uuid"
// @Produce	application/json
// @Success	200 {string} string {"users":[{"id":1,"email":"allinemailtest@163.com","role":3,"del_flag":0,"last_update_time":"2021-11-22T13:46:20.430926+08:00","create_time":"2021-10-25T09:21:50.364327+08:00","user_name":"zhangsan","department_name":"arch","employee_id":"100001","account_name":"zs001","telephone":"01012345678","mobile":"13012345678"}]}
// @Router	/api/vi/metadata/resource-group/user/group-uuid [get]
func GetUsersByGroupUUID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	groupUUID, err := jsonparser.GetString(data, resourceGroupGroupUUIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupGroupUUIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()

	// get entity
	err = s.GetUsersByGroupUUID(groupUUID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetUsersByGroupUUID, err, groupUUID)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetUsersByGroupUUID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetUsersByGroupUUID, groupUUID)
}

// @Tags	resource group
// @Summary	get users by resource group uuid
// @Accept	application/json
// @Param	token 		body 	string 	true 	"token"
// @Param	group_uuid	body	string	true	"resource group uuid"
// @Produce	application/json
// @Success	200 {string} string {"users":[{"id":1,"email":"allinemailtest@163.com","role":3,"del_flag":0,"last_update_time":"2021-11-22T13:46:20.430926+08:00","create_time":"2021-10-25T09:21:50.364327+08:00","user_name":"zhangsan","department_name":"arch","employee_id":"100001","account_name":"zs001","telephone":"01012345678","mobile":"13012345678"}]}
// @Router	/api/vi/metadata/resource-group/das-admin/group-uuid [get]
func GetDASAdminUsersByGroupUUID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	groupUUID, err := jsonparser.GetString(data, resourceGroupGroupUUIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupGroupUUIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()

	// get entity
	err = s.GetDASAdminUsersByUUID(groupUUID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDASAdminUsersByGroupUUID, err, groupUUID)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupUsersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDASAdminUsersByGroupUUID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDASAdminUsersByGroupUUID, groupUUID)
}

// @Tags	resource group
// @Summary	add a new resource group
// @Accept	application/json
// @Param	token 		body 	string 	true 	"token"
// @Param	group_uuid	body	string	true	"resource group uuid"
// @Param	group_name	body	string  true 	"resource group name"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group [post]
func AddResourceGroup(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err)
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.ResourceGroupInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, ok := fields[resourceGroupGroupUUIDStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, resourceGroupGroupUUIDJSON)
		return
	}
	_, ok = fields[resourceGroupGroupNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, resourceGroupGroupNameStruct)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddResourceGroup, err, fields[resourceGroupGroupUUIDStruct], fields[resourceGroupGroupNameStruct])
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddResourceGroup, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddResourceGroup, fields[resourceGroupGroupUUIDStruct], fields[resourceGroupGroupNameStruct])
}

// @Tags	resource group
// @Summary	update resource group by id
// @Accept	application/json
// @Param	token 		body 	string 	true 	"token"
// @Param	id			body	int		true	"resource group id"
// @Param	group_uuid	body	string	false	"resource group uuid"
// @Param	group_name	body	string	false	"resource group name"
// @Param	del_flag	body	int		false	"delete flag"
// @Produce	application/json
// @Success	200 {string} string {"middleware_clusters":[{"last_update_time":"2021-11-18T15:39:52.927116+08:00","id":1,"cluster_name":"update_middleware_cluster","env_id":1,"del_flag":0,"create_time":"2021-11-09T18:06:57.917596+08:00"}]}
// @Router	/api/v1/metadata/resource-group/update [post]
func UpdateResourceGroupByID(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.ResourceGroupInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	idInterface, idExists := fields[resourceGroupIDStruct]
	if !idExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, resourceGroupIDJSON)
		return
	}
	id, ok := idInterface.(int)
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, resourceGroupIDJSON)
		return
	}
	_, resourceGroupGroupUUIDExists := fields[resourceGroupGroupUUIDStruct]
	_, resourceGroupNameExists := fields[resourceGroupGroupUUIDStruct]
	_, resourceGroupDelFlagExists := fields[resourceGroupDelFlagStruct]
	if !resourceGroupGroupUUIDExists && !resourceGroupNameExists && !resourceGroupDelFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s, %s and %s", resourceGroupGroupUUIDJSON, resourceGroupGroupNameJSON, resourceGroupDelFlagJSON))
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateResourceGroup, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateResourceGroup, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataUpdateResourceGroup, id)
}

// @Tags	resource group
// @Summary	delete resource group by id
// @Accept	application/json
// @Param	token 	body	string 	true 	"token"
// @Param	id		body	int		true	"resource group id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group/delete [post]
func DeleteResourceGroupByID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	id, err := jsonparser.GetInt(data, resourceGroupIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// update entities
	err = s.Delete(int(id))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteResourceGroup, err, id)
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
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteResourceGroup, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteResourceGroup, id)
}

// @Tags	resource group
// @Summary	add mysql cluster map
// @Accept	application/json
// @Param	token 				body 	string 	true 	"token"
// @Param	id					body	int		true	"resource group id"
// @Param	mysql_cluster_id	body	int		true	"mysql cluster id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group/add-mysql-cluster [post]
func ResourceGroupAddMySQLCluster(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}

	resourceGroupID, err := jsonparser.GetInt(data, resourceGroupIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupIDJSON)
		return
	}
	mysqlClusterID, err := jsonparser.GetInt(data, resourceGroupMySQLClusterIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupMySQLClusterIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// update entities
	err = s.AddMySQLCluster(int(resourceGroupID), int(mysqlClusterID))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataResourceGroupAddMySQLCluster, err, resourceGroupID, mysqlClusterID)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupMySQLClustersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataResourceGroupAddMySQLCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataResourceGroupAddMySQLCluster, resourceGroupID, mysqlClusterID)
}

// @Tags	resource group
// @Summary	delete mysql cluster map
// @Accept	application/json
// @Param	token 				body 	string 	true 	"token"
// @Param	id					body	int		true	"resource group id"
// @Param	mysql_cluster_id	body	int		true	"mysql cluster id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group/delete-mysql-cluster [post]
func ResourceGroupDeleteMySQLCluster(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	resourceGroupID, err := jsonparser.GetInt(data, resourceGroupIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupIDJSON)
		return
	}
	mysqlClusterID, err := jsonparser.GetInt(data, resourceGroupMySQLClusterIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupMySQLClusterIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// update entities
	err = s.DeleteMySQLCluster(int(resourceGroupID), int(mysqlClusterID))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataResourceGroupDeleteMySQLCluster, err, resourceGroupID, mysqlClusterID)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupMySQLClustersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataResourceGroupDeleteMySQLCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataResourceGroupDeleteMySQLCluster, resourceGroupID, mysqlClusterID)
}

// @Tags	resource group
// @Summary	add middleware cluster map
// @Accept	application/json
// @Param	token 					body 	string 	true 	"token"
// @Param	id						body	int		true	"resource group id"
// @Param	middleware_cluster_id	body	int		true	"middleware cluster id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group/add-middleware-cluster [post]
func ResourceGroupAddMiddlewareCluster(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}

	resourceGroupID, err := jsonparser.GetInt(data, resourceGroupIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupIDJSON)
		return
	}
	middlewareClusterID, err := jsonparser.GetInt(data, resourceGroupMiddlewareClusterIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupMiddlewareClusterIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// update entities
	err = s.AddMiddlewareCluster(int(resourceGroupID), int(middlewareClusterID))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataResourceGroupAddMiddlewareCluster, err, resourceGroupID, middlewareClusterID)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupMiddlewareClustersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataResourceGroupAddMiddlewareCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataResourceGroupAddMiddlewareCluster, resourceGroupID, middlewareClusterID)
}

// @Tags	resource group
// @Summary	delete middleware cluster map
// @Accept	application/json
// @Param	token 					body 	string 	true 	"token"
// @Param	id						body	int		true	"resource group id"
// @Param	middleware_cluster_id	body	int		true	"middleware cluster id"
// @Produce	application/json
// @Success	200 {string} string ""
// @Router	/api/v1/metadata/resource-group/delete-middleware-cluster [post]
func ResourceGroupDeleteMiddlewareCluster(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	resourceGroupID, err := jsonparser.GetInt(data, resourceGroupIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupIDJSON)
		return
	}
	middlewareClusterID, err := jsonparser.GetInt(data, resourceGroupMiddlewareClusterIDJSON)
	if err != nil {
		resp.ResponseNOK(c, message.ErrFieldNotExistsOrWrongType, errors.Trace(err), resourceGroupMiddlewareClusterIDJSON)
		return
	}
	// init service
	s := metadata.NewResourceGroupServiceWithDefault()
	// update entities
	err = s.DeleteMiddlewareCluster(int(resourceGroupID), int(middlewareClusterID))
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataResourceGroupDeleteMiddlewareCluster, err, resourceGroupID, middlewareClusterID)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(resourceGroupMiddlewareClustersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataResourceGroupDeleteMiddlewareCluster, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataResourceGroupDeleteMiddlewareCluster, resourceGroupID, middlewareClusterID)
}
