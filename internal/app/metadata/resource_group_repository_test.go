package metadata

import (
	"testing"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	testResourceGroupID              = 1
	testResourceGroupUUID            = "5c6c6d73-eac2-11ec-8183-001c42d502a6"
	testResourceGroupName            = "resource_group_01"
	testResourceGroupGroupUUID       = "9b22fd82-e96c-11ec-8183-001c42d502a6"
	testResourceGroupNewGroupName    = "test_new_resource_group"
	testResourceGroupUpdateGroupName = "test_update_resource_group"
)

var testResourceGroupRepo *ResourceGroupRepo

func init() {
	testInitDASMySQLPool()
	testResourceGroupRepo = NewResourceGroupRepoWithGlobal()
}

func testCreateResourceGroup() (metadata.ResourceGroup, error) {
	resourceGroup := NewResourceGroupInfoWithDefault(
		testResourceGroupGroupUUID,
		testResourceGroupNewGroupName,
	)

	entity, err := testResourceGroupRepo.Create(resourceGroup)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func TestResourceGroupRepo_All(t *testing.T) {
	TestResourceGroupRepo_Execute(t)
	TestResourceGroupRepo_Transaction(t)
	TestResourceGroupRepo_GetAll(t)
	TestResourceGroupRepo_GetByID(t)
	TestResourceGroupRepo_GetByGroupUUID(t)
	TestResourceGroupRepo_GetResourceRolesByID(t)
	TestResourceGroupRepo_GetMySQLClustersByID(t)
	TestResourceGroupRepo_GetMySQLServersByID(t)
	TestResourceGroupRepo_GetMiddlewareClustersByID(t)
	TestResourceGroupRepo_GetMiddlewareServersByID(t)
	TestResourceGroupRepo_GetUsersByID(t)
	TestResourceGroupRepo_GetDASAdminUsersByID(t)
	TestResourceGroupRepo_Create(t)
	TestResourceGroupRepo_Update(t)
	TestResourceGroupRepo_Delete(t)
	TestResourceGroupRepo_AddMySQLCluster(t)
	TestResourceGroupRepo_DeleteMySQLCluster(t)
	TestResourceGroupRepo_AddMiddlewareCluster(t)
	TestResourceGroupRepo_DeleteMiddlewareCluster(t)
}

func TestResourceGroupRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := testResourceGroupRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestResourceGroupRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_resource_group_info(group_uuid, group_name) values(?,?);`

	tx, err := testResourceGroupRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(
		sql,
		testResourceGroupGroupUUID,
		testResourceGroupNewGroupName,
	)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select group_name from t_meta_resource_group_info where group_uuid=?`
	result, err := tx.Execute(sql, testResourceGroupGroupUUID)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	groupName, err := result.GetString(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if groupName != testResourceGroupNewGroupName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := testMySQLServerRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		if entity.GetServerName() == testMySQLServerNewServerName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestResourceGroupRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := testResourceGroupRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(1, len(entities), "test GetAll() failed")
}

func TestResourceGroupRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := testResourceGroupRepo.GetByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testResourceGroupName, entity.GetGroupName(), "test GetByGroupUUID() failed")
}

func TestResourceGroupRepo_GetByGroupUUID(t *testing.T) {
	asst := assert.New(t)

	entity, err := testResourceGroupRepo.GetByGroupUUID(testResourceGroupUUID)
	asst.Nil(err, common.CombineMessageWithError("test GetByGroupUUID() failed", err))
	asst.Equal(testResourceGroupName, entity.GetGroupName(), "test GetByGroupUUID() failed")
}

func TestResourceGroupRepo_GetID(t *testing.T) {
	asst := assert.New(t)

	id, err := testResourceGroupRepo.GetID(testResourceGroupUUID)
	asst.Nil(err, common.CombineMessageWithError("test GetID() failed", err))
	asst.Equal(testResourceGroupID, id, "test GetID() failed")
}

func TestResourceGroupRepo_GetResourceRolesByID(t *testing.T) {
	asst := assert.New(t)

	resourceRoles, err := testResourceGroupRepo.GetResourceRolesByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test GetResourceRoles() failed", err))
	asst.Equal(2, len(resourceRoles), "test GetResourceRoles() failed")
}

func TestResourceGroupRepo_GetMySQLClustersByID(t *testing.T) {
	asst := assert.New(t)

	mysqlClusters, err := testResourceGroupRepo.GetMySQLClustersByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLClustersByID() failed", err))
	asst.Equal(2, len(mysqlClusters), "test GetMySQLClustersByID() failed")
}

func TestResourceGroupRepo_GetMySQLServersByID(t *testing.T) {
	asst := assert.New(t)

	mysqlServers, err := testResourceGroupRepo.GetMySQLServersByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLServersByID() failed", err))
	asst.Equal(2, len(mysqlServers), "test GetMySQLServersByID() failed")
}

func TestResourceGroupRepo_GetMiddlewareClustersByID(t *testing.T) {
	asst := assert.New(t)

	middlewareClusters, err := testResourceGroupRepo.GetMiddlewareClustersByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test GetMiddlewareClustersByID() failed", err))
	asst.Equal(1, len(middlewareClusters), "test GetMiddlewareClustersByID() failed")
}

func TestResourceGroupRepo_GetMiddlewareServersByID(t *testing.T) {
	asst := assert.New(t)

	middlewareServers, err := testResourceGroupRepo.GetMiddlewareServersByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test GetMiddlewareServersByID() failed", err))
	asst.Equal(1, len(middlewareServers), "test GetMiddlewareServersBy() failed")
}

func TestResourceGroupRepo_GetUsersByID(t *testing.T) {
	asst := assert.New(t)

	users, err := testResourceGroupRepo.GetUsersByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test GetUsersByID() failed", err))
	asst.Equal(1, len(users), "test GetUsersByID() failed")
}

func TestResourceGroupRepo_GetDASAdminUsersByID(t *testing.T) {
	asst := assert.New(t)

	users, err := testResourceGroupRepo.GetDASAdminUsersByID(testResourceGroupID)
	asst.Nil(err, common.CombineMessageWithError("test GetDASAdminUsersByID() failed", err))
	asst.Equal(1, len(users), "test GetDASAdminUsersByID() failed")
}

func TestResourceGroupRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateResourceGroup()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testResourceGroupRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestResourceGroupRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateResourceGroup()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{resourceGroupGroupNameStruct: testResourceGroupUpdateGroupName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testResourceGroupRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = testResourceGroupRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testResourceGroupUpdateGroupName, entity.GetGroupName(), "test Update() failed")
	// delete
	err = testResourceGroupRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestResourceGroupRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateResourceGroup()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testResourceGroupRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestResourceGroupRepo_AddMySQLCluster(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupRepo.DeleteMySQLCluster(testResourceGroupID, testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLCluster() failed", err))
	err = testResourceGroupRepo.AddMySQLCluster(testResourceGroupID, testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLCluster() failed", err))
}

func TestResourceGroupRepo_DeleteMySQLCluster(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupRepo.DeleteMySQLCluster(testResourceGroupID, testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLCluster() failed", err))
	err = testResourceGroupRepo.AddMySQLCluster(testResourceGroupID, testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLCluster() failed", err))
}

func TestResourceGroupRepo_AddMiddlewareCluster(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupRepo.DeleteMiddlewareCluster(testResourceGroupID, testMiddlewareClusterClusterID)
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLCluster() failed", err))
	err = testResourceGroupRepo.AddMiddlewareCluster(testResourceGroupID, testMiddlewareClusterClusterID)
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLCluster() failed", err))
}

func TestResourceGroupRepo_DeleteMiddlewareCluster(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupRepo.DeleteMiddlewareCluster(testResourceGroupID, testMiddlewareClusterClusterID)
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLCluster() failed", err))
	err = testResourceGroupRepo.AddMiddlewareCluster(testResourceGroupID, testMiddlewareClusterClusterID)
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLCluster() failed", err))
}
