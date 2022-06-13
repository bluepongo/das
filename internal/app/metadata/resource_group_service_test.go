package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

var testResourceGroupService *ResourceGroupService

func init() {
	testInitDASMySQLPool()
	testResourceGroupService = NewResourceGroupServiceWithDefault()
}

func TestResourceGroupService_All(t *testing.T) {
	TestResourceGroupService_GetAll(t)
	TestResourceGroupService_GetByID(t)
	TestResourceGroupService_GetByGroupUUID(t)
	TestResourceGroupService_GetResourceRolesByID(t)
	TestResourceGroupService_GetMySQLClustersByID(t)
	TestResourceGroupService_GetMySQLServersByID(t)
	TestResourceGroupService_GetMiddlewareClustersByID(t)
	TestResourceGroupService_GetMiddlewareServersByID(t)
	TestResourceGroupService_GetUsersByID(t)
	TestResourceGroupService_GetDASAdminUsersByID(t)
	TestResourceGroupService_GetResourceRolesByGroupUUID(t)
	TestResourceGroupService_GetMySQLClustersByGroupUUID(t)
	TestResourceGroupService_GetMySQLServersByGroupUUID(t)
	TestResourceGroupService_GetMiddlewareClustersByGroupUUID(t)
	TestResourceGroupService_GetMiddlewareServersByGroupUUID(t)
	TestResourceGroupService_GetUsersByGroupUUID(t)
	TestResourceGroupService_GetDASAdminUsersByGroupUUID(t)
	TestResourceGroupService_Create(t)
	TestResourceGroupService_Update(t)
	TestResourceGroupService_Delete(t)
	TestResourceGroupService_AddMySQLCluster(t)
	TestResourceGroupService_DeleteMySQLCluster(t)
	TestResourceGroupService_AddMiddlewareCluster(t)
	TestResourceGroupService_DeleteMiddlewareCluster(t)
	TestResourceGroupService_Marshal(t)
	TestResourceGroupService_MarshalWithFields(t)
}

func TestResourceGroupService_GetAll(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(1, len(testResourceGroupService.GetResourceGroups()), "test GetAll() failed")
}

func TestResourceGroupService_GetByID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testResourceGroupID, testResourceGroupService.GetResourceGroups()[constant.ZeroInt].Identity(), "test GetByID() failed")
}

func TestResourceGroupService_GetByGroupUUID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetByGroupUUID(testResourceGroupGroupUUID)
	asst.Nil(err, common.CombineMessageWithError("test GetByGroupUUID() failed", err))
	asst.Equal(testResourceGroupGroupUUID, testResourceGroupService.GetResourceGroups()[constant.ZeroInt].GetGroupUUID(), "test GetByGroupUUID() failed")
}

func TestResourceGroupService_GetResourceRolesByID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetResourceRolesByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test GetResourceRolesByID() failed", err))
	asst.Equal(2, len(testResourceGroupService.GetResourceRoles()), "test GetResourceRolesByID() failed")
}

func TestResourceGroupService_GetMySQLClustersByID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetMySQLClustersByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLClustersByID() failed", err))
	asst.Equal(2, len(testResourceGroupService.GetMySQLClusters()), "test GetMySQLClustersByID() failed")
}

func TestResourceGroupService_GetMySQLServersByID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetMySQLServersByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLServersByID() failed", err))
	asst.Equal(2, len(testResourceGroupService.GetMySQLServers()), "test GetMySQLServersByID() failed")
}

func TestResourceGroupService_GetMiddlewareClustersByID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetMiddlewareClustersByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test GetMiddlewareClustersByID() failed", err))
	asst.Equal(1, len(testResourceGroupService.GetMiddlewareClusters()), "test GetMiddlewareClustersByID() failed")
}

func TestResourceGroupService_GetMiddlewareServersByID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetMiddlewareServersByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test GetMiddlewareServersByID() failed", err))
	asst.Equal(1, len(testResourceGroupService.GetMiddlewareServers()), "test GetMiddlewareServersByID() failed")
}

func TestResourceGroupService_GetUsersByID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetUsersByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test GetUsersByID() failed", err))
	asst.Equal(1, len(testResourceGroupService.GetUsers()), "test GetUsersByID() failed")
}

func TestResourceGroupService_GetDASAdminUsersByID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetDASAdminUsersByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test GetDASAdminUsersByID() failed", err))
	asst.Equal(1, len(testResourceGroupService.GetUsers()), "test GetDASAdminUsersByID() failed")
}

func TestResourceGroupService_GetResourceRolesByGroupUUID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetResourceRolesByGroupUUID(testResourceGroupGroupUUID)
	asst.Nil(err, common.CombineMessageWithError("test GetResourceRolesByGroupUUID() failed", err))
	asst.Equal(2, len(testResourceGroupService.GetResourceRoles()), "test GetResourceRolesByGroupUUID() failed")
}

func TestResourceGroupService_GetMySQLClustersByGroupUUID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetMySQLClustersByGroupUUID(testResourceGroupGroupUUID)
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLClustersByGroupUUID() failed", err))
	asst.Equal(2, len(testResourceGroupService.GetMySQLClusters()), "test GetMySQLClustersByGroupUUID() failed")
}

func TestResourceGroupService_GetMySQLServersByGroupUUID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetMySQLServersByGroupUUID(testResourceGroupGroupUUID)
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLServersByGroupUUID() failed", err))
	asst.Equal(2, len(testResourceGroupService.GetMySQLServers()), "test GetMySQLServersByGroupUUID() failed")
}

func TestResourceGroupService_GetMiddlewareClustersByGroupUUID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetMiddlewareClustersByGroupUUID(testResourceGroupGroupUUID)
	asst.Nil(err, common.CombineMessageWithError("test GetMiddlewareClustersByGroupUUID() failed", err))
	asst.Equal(1, len(testResourceGroupService.GetMiddlewareClusters()), "test GetMiddlewareClustersByGroupUUID() failed")
}

func TestResourceGroupService_GetMiddlewareServersByGroupUUID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetMiddlewareServersByGroupUUID(testResourceGroupGroupUUID)
	asst.Nil(err, common.CombineMessageWithError("test GetMiddlewareServersByGroupUUID() failed", err))
	asst.Equal(1, len(testResourceGroupService.GetMiddlewareServers()), "test GetMiddlewareServersByGroupUUID() failed")
}

func TestResourceGroupService_GetUsersByGroupUUID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetUsersByGroupUUID(testResourceGroupGroupUUID)
	asst.Nil(err, common.CombineMessageWithError("test GetUsersByGroupUUID() failed", err))
	asst.Equal(1, len(testResourceGroupService.GetUsers()), "test GetUsersByGroupUUID() failed")
}

func TestResourceGroupService_GetDASAdminUsersByGroupUUID(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetDASAdminUsersByGroupUUID(testResourceGroupGroupUUID)
	asst.Nil(err, common.CombineMessageWithError("test GetDASAdminUsersByGroupUUID() failed", err))
	asst.Equal(1, len(testResourceGroupService.GetUsers()), "test GetDASAdminUsersByGroupUUID() failed")
}

func TestResourceGroupService_Create(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.Create(map[string]interface{}{
		resourceGroupGroupUUIDStruct: testResourceGroupNewGroupUUID,
		resourceGroupGroupNameStruct: testResourceGroupNewGroupName,
	})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testResourceGroupService.Delete(testResourceGroupService.GetResourceGroups()[constant.ZeroInt].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestResourceGroupService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateResourceGroup()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))

	err = testResourceGroupService.Update(entity.Identity(), map[string]interface{}{resourceGroupGroupNameStruct: testResourceGroupUpdateGroupName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testResourceGroupService.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testResourceGroupUpdateGroupName, testResourceGroupService.GetResourceGroups()[constant.ZeroInt].GetGroupName(), "test Update() failed")
	// delete
	err = testResourceGroupService.Delete(testResourceGroupService.GetResourceGroups()[constant.ZeroInt].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestResourceGroupService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateResourceGroup()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testResourceGroupService.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestResourceGroupService_AddMySQLCluster(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.DeleteMySQLCluster(testResourceGroupID, testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLCluster() failed", err))
	err = testResourceGroupService.AddMySQLCluster(testResourceGroupID, testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLCluster() failed", err))
}

func TestResourceGroupService_DeleteMySQLCluster(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.DeleteMySQLCluster(testResourceGroupID, testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteMySQLCluster() failed", err))
	err = testResourceGroupService.AddMySQLCluster(testResourceGroupID, testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteMySQLCluster() failed", err))
}

func TestResourceGroupService_AddMiddlewareCluster(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.DeleteMiddlewareCluster(testResourceGroupID, testMiddlewareClusterClusterID)
	asst.Nil(err, common.CombineMessageWithError("test AddMiddlewareCluster() failed", err))
	err = testResourceGroupService.AddMiddlewareCluster(testResourceGroupID, testMiddlewareClusterClusterID)
	asst.Nil(err, common.CombineMessageWithError("test AddMiddlewareCluster() failed", err))
}

func TestResourceGroupService_DeleteMiddlewareCluster(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.DeleteMiddlewareCluster(testResourceGroupID, testMiddlewareClusterClusterID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteMiddlewareCluster() failed", err))
	err = testResourceGroupService.AddMiddlewareCluster(testResourceGroupID, testMiddlewareClusterClusterID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteMiddlewareCluster() failed", err))
}

func TestResourceGroupService_Marshal(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	jsonBytes, err := testResourceGroupService.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}

func TestResourceGroupService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	jsonBytes, err := testResourceGroupService.MarshalWithFields(resourceGroupGroupNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	t.Log(string(jsonBytes))
}

func TestMarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupService.GetMiddlewareServersByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	jsonBytes, err := testResourceGroupService.MarshalWithFields(resourceGroupMiddlewareServersStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	t.Log(string(jsonBytes))
}
