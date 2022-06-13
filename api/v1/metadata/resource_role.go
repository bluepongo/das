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

// @Tags resource role
// @Summary	get all resource roles
// @Accept	application/json
// @Produce	application/json
// @Param	token	body string	true "token"
// @Success	200 {string} string "{"resource_roles":[{"del_flag":0,"create_time":"2022-06-12T09:03:23.298572+08:00","last_update_time":"2022-06-12T09:21:36.667854+08:00","id":1,"role_uuid":"test_role_uuid","role_name":"test_role","resource_group_id":1}]}"
// @Router	/api/v1/metadata/resource-role [get]
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

// @Tags	resource role
// @Summary	get resource role by id
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	id		body int	true "resource role id"
// @Produce	application/json
// @Success	200 {string} string "{"resource_roles":[{"id":1,"role_uuid":"test_role_uuid","role_name":"test_role","resource_group_id":1,"del_flag":0,"create_time":"2022-06-12T09:03:23.298572+08:00","last_update_time":"2022-06-12T09:21:36.667854+08:00"}]}"
// @Router	/api/v1/metadata/resource-role/get [get]
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

// @Tags resource role
// @Summary get resource role by uuid
// @Accept	application/json
// @Param	token		body string	true "token"
// @Param	role_uuid	body string	true "resource role uuid"
// @Produce  application/json
// @Success 200 {string} string "{"resource_roles":[{"id":1,"role_uuid":"test_role_uuid","role_name":"test_role","resource_group_id":1,"del_flag":0,"create_time":"2022-06-12T09:03:23.298572+08:00","last_update_time":"2022-06-12T09:21:36.667854+08:00"}]}"
// @Router /api/v1/metadata/resource-role/role-uuid [get]
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

// @Tags	resource role
// @Summary	get resource group by resource role id
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	id		body int	true "resource role id"
// @Produce	application/json
// @Success	200 {string} string "{"resource_group":{"id":1,"group_uuid":"uuid","group_name":"test","del_flag":0,"create_time":"2022-06-12T09:02:39.376944+08:00","last_update_time":"2022-06-12T09:02:39.376944+08:00"}}"
// @Router	/api/v1/metadata/resource-role/resource-group [get]
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

// @Tags	resource role
// @Summary	get resource role users
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	id		body int	true "resource role id"
// @Produce	application/json
// @Success	200 {string} string "{"users":[{"id":14,"mobile":"","role":3,"del_flag":0,"user_name":"tester","department_name":"","employee_id":"","account_name":"test","email":"929059501@qq.com","telephone":"","create_time":"2021-12-06T18:08:03.736262+08:00","last_update_time":"2021-12-06T18:08:03.736262+08:00"}]}"
// @Router	/api/v1/metadata/resource-role/user [get]
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

// @Tags	resource role
// @Summary	get resource role users by resource role uuid
// @Accept	application/json
// @Param	token		body string	true "token"
// @Param	role_uuid	body int	true "resource role uuid"
// @Produce	application/json
// @Success	200 {string} string "{"users":[{"account_name":"test","email":"929059501@qq.com","telephone":"","mobile":"","role":3,"user_name":"tester","employee_id":"","del_flag":0,"create_time":"2021-12-06T18:08:03.736262+08:00","id":14,"department_name":"","last_update_time":"2021-12-06T18:08:03.736262+08:00"}]}"
// @Router	/api/v1/metadata/resource-role/user/role-uuid [get]
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

// @Tags	resource role
// @Summary	add user map
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	id		body int	true "resource role id"
// @Param	user_id	body int	true "user id"
// @Produce	application/json
// @Success	200 {string} string "{"users":[{"account_name":"test","email":"929059501@qq.com","telephone":"","mobile":"","role":3,"id":14,"department_name":"","employee_id":"","del_flag":0,"create_time":"2021-12-06T18:08:03.736262+08:00","last_update_time":"2021-12-06T18:08:03.736262+08:00","user_name":"tester"}]}"
// @Router	/api/v1/metadata/resource-role/add-user [post]
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

// @Tags	resource role
// @Summary	delete user map
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	id		body int	true "resource role id"
// @Param	user_id	body int	true "user id"
// @Produce	application/json
// @Success	200 {string} string "{"users":[]}"
// @Router	/api/v1/metadata/resource-role/delete-user [post]
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

// @Tags	resource role
// @Summary	add a new resource role
// @Accept	application/json
// @Param	token				body string	true  "token"
// @Param	role_name			body string	false "resource role name"
// @Param	role_uuid			body int	true  "resource role uuid"
// @Param	resource_group_id	body int	true  "resource group id"
// @Produce	application/json
// @Success	200 {string} string "{"resource_roles":[{"id":2,"role_uuid":"new_test_role_uuid","role_name":"new_test_role","resource_group_id":1,"del_flag":0,"create_time":"2022-06-13T01:08:21.739904+08:00","last_update_time":"2022-06-13T01:08:21.739904+08:00"}]}"
// @Router	/api/v1/metadata/resource-role [post]
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

// @Tags	resource role
// @Summary	update resource role by id
// @Accept	application/json
// @Param	token				body string	true  "token"
// @Param	id					body int	true  "resource role id"
// @Param	role_name			body string	false "resource role name"
// @Param	role_uuid			body int	false "resource role uuid"
// @Param	resource_group_id	body int	false "resource group id"
// @Param	del_flag			body int	false "delete flag"
// @Produce	application/json
// @Success	200 {string} string "{"resource_roles":[{"resource_group_id":1,"del_flag":0,"create_time":"2022-06-13T01:08:21.739904+08:00","last_update_time":"2022-06-13T01:08:21.739904+08:00","id":2,"role_uuid":"update_test_role_id","role_name":"update_test_role"}]}}"
// @Router	/api/v1/metadata/resource-role/update [post]
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

// @Tags	resource role
// @Summary	delete resource role by id
// @Accept	application/json
// @Param	token	body string	true "token"
// @Param	id		body int	true "resource role id"
// @Produce	application/json
// @Success	200 {string} string "{"resource_roles":[{"role_name":"update_test_role","resource_group_id":1,"del_flag":0,"create_time":"2022-06-13T01:08:21.739904+08:00","last_update_time":"2022-06-13T01:10:42.472215+08:00","id":2,"role_uuid":"update_test_role_id"}]}"
// @Router	/api/v1/metadata/resource-role/delete [post]
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
