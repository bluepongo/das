package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

var testResourceRoleService *ResourceRoleService

func init() {
	testInitDASMySQLPool()
	testResourceRoleService = NewResourceRoleServiceWithDefault()
}

func TestResourceRoleServiceAll(t *testing.T) {
	TestResourceRoleService_GetResourceRoles(t)
	TestResourceRoleService_GetResourceGroup(t)
	TestResourceRoleService_GetUsers(t)
	TestResourceRoleService_GetAll(t)
	TestResourceRoleService_GetByID(t)
	TestResourceRoleService_GetByRoleUUID(t)
	TestResourceRoleService_GetResourceGroupByID(t)
	TestResourceRoleService_GetUsersByID(t)
	TestResourceRoleService_AddUser(t)
	TestResourceRoleService_DeleteUser(t)
	TestResourceRoleService_Create(t)
	TestResourceRoleService_Update(t)
	TestResourceRoleService_Delete(t)
	TestResourceRoleService_Marshal(t)
	TestResourceRoleService_MarshalWithFields(t)
}

func TestResourceRoleService_GetResourceRoles(t *testing.T) {
	asst := assert.New(t)

	err := testResourceRoleService.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLServers() failed", err))
	asst.Equal(2, len(testResourceRoleService.GetResourceRoles()), "test GetMySQLServers() failed")
}

func TestResourceRoleService_GetResourceGroup(t *testing.T) {
	asst := assert.New(t)

	err := testResourceRoleService.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLServers() failed", err))
	asst.Equal(2, len(testResourceRoleService.GetResourceRoles()), "test GetMySQLServers() failed")
}
func TestResourceRoleService_GetUsers(t *testing.T) {
	asst := assert.New(t)

	err := testResourceRoleService.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLServers() failed", err))
	asst.Equal(2, len(testResourceRoleService.GetResourceRoles()), "test GetMySQLServers() failed")
}
func TestResourceRoleService_GetAll(t *testing.T) {
	asst := assert.New(t)

	err := testResourceRoleService.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(2, len(testResourceRoleService.GetResourceRoles()), "test GetAll() failed")
}

func TestResourceRoleService_GetByID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceRoleService.GetByID(testResourceRoleID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testResourceRoleID, testResourceRoleService.GetResourceRoles()[constant.ZeroInt].Identity(), "test GetByID() failed")
}

func TestResourceRoleService_GetByRoleUUID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceRoleService.GetByRoleUUID(testResourceRoleRoleUUID)
	asst.Nil(err, common.CombineMessageWithError("test GetByRoleUUID() failed", err))
	asst.Equal(testResourceRoleRoleUUID, testResourceRoleService.GetResourceRoles()[constant.ZeroInt].GetRoleUUID(), "test GetByRoleUUID() failed")
}

func TestResourceRoleService_GetResourceGroupByID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceRoleService.GetResourceGroupByID(testResourceRoleID)
	asst.Nil(err, common.CombineMessageWithError("test GetResourceGroupByID() failed", err))
	asst.Equal(1, testResourceRoleService.GetResourceGroup().Identity(), "test GetResourceGroupByID() failed")
}

func TestResourceRoleService_GetUsersByID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceRoleService.GetUsersByID(testResourceRoleID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBSByID() failed", err))
	asst.Equal(1, len(testResourceRoleService.Users), common.CombineMessageWithError("test GetDBSByID() failed", err))
}

func TestResourceRoleService_GetUsersByRoleUUID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceRoleService.GetUsersByID(testResourceRoleID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBSByID() failed", err))
	asst.Equal(1, len(testResourceRoleService.Users), common.CombineMessageWithError("test GetDBSByID() failed", err))
}
func TestResourceRoleService_AddUser(t *testing.T) {
	asst := assert.New(t)
	entity, err := testCreateResourceRole()
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
	err = testResourceRoleService.AddUser(entity.Identity(), testResourceRoleUserID)
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
	err = testResourceRoleService.GetUsersByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
	asst.Equal(testResourceRoleUserID, testResourceRoleService.GetUsers()[constant.ZeroInt].Identity())
	err = testResourceRoleService.DeleteUser(entity.Identity(), testResourceRoleUserID)
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
	err = testResourceRoleRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
}

func TestResourceRoleService_DeleteUser(t *testing.T) {
	asst := assert.New(t)
	entity, err := testCreateResourceRole()
	asst.Nil(err, common.CombineMessageWithError("test DeleteUser() failed", err))
	err = testResourceRoleService.AddUser(entity.Identity(), testResourceRoleUserID)
	asst.Nil(err, "test DeleteUser() failed")
	err = testResourceRoleService.DeleteUser(entity.Identity(), testResourceRoleUserID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteUser() failed", err))
	err = testResourceRoleService.GetUsersByID(entity.Identity())
	asst.Zero(len(testResourceRoleService.GetUsers()))
	err = testResourceRoleRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DeleteUser() failed", err))
}

func TestResourceRoleService_Create(t *testing.T) {
	asst := assert.New(t)

	err := testResourceRoleService.Create(map[string]interface{}{
		resourceRoleRoleUUIDStruct:        testResourceRoleNewRoleUUID,
		resourceRoleRoleNameStruct:        testResourceRoleRoleName,
		resourceRoleResourceGroupIDStruct: testResourceRoleResourceGroupID,
	})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testResourceRoleService.Delete(testResourceRoleService.GetResourceRoles()[constant.ZeroInt].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestResourceRoleService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateResourceRole()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))

	err = testResourceRoleService.Update(entity.Identity(), map[string]interface{}{resourceRoleRoleUUIDStruct: testResourceRoleUpdateRoleUUID})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testResourceRoleService.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testResourceRoleUpdateRoleUUID, testResourceRoleService.GetResourceRoles()[constant.ZeroInt].GetRoleUUID(), "test Update() failed")
	// delete
	err = testResourceRoleService.Delete(testResourceRoleService.GetResourceRoles()[constant.ZeroInt].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestResourceRoleService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateResourceRole()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testResourceRoleService.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestResourceRoleService_Marshal(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testResourceRoleService.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}

func TestResourceRoleService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testResourceRoleService.MarshalWithFields(resourceRoleRoleNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	t.Log(string(jsonBytes))
}
