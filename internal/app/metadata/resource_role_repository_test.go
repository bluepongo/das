package metadata

import (
	"testing"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	testResourceRoleNewRoleName    = "test_new_role_name"
	testResourceRoleUpdateRoleName = "test_update_role_name"
	testResourceRoleUserID         = 1
)

var testResourceRoleRepo *ResourceRoleRepo

func init() {
	testInitDASMySQLPool()
	testResourceRoleRepo = NewResourceRoleRepoWithGlobal()
}

func testCreateResourceRole() (metadata.ResourceRole, error) {
	resourceRoleInfo := NewResourceRoleInfoWithDefault(
		testResourceRoleNewRoleName,
		testResourceRoleResourceGroupID)
	entity, err := testResourceRoleRepo.Create(resourceRoleInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func TestResourceRoleRepoAll(t *testing.T) {
	TestResourceRoleRepo_Execute(t)
	TestResourceRoleRepo_Transaction(t)
	TestResourceRoleRepo_Create(t)
	TestResourceRoleRepo_GetAll(t)
	TestResourceRoleRepo_GetByID(t)
	TestResourceRoleRepo_GetID(t)
	TestResourceRoleRepo_GetResourceGroupByID(t)
	TestResourceRoleRepo_GetUsersByID(t)
	TestResourceRoleRepo_AddUser(t)
	TestResourceRoleRepo_DeleteUser(t)
	TestResourceRoleRepo_Update(t)
	TestResourceRoleRepo_Delete(t)
}

func TestResourceRoleRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := testResourceRoleRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestResourceRoleRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `
	insert into t_meta_mysql_cluster_info(
		cluster_name, middleware_cluster_id, monitor_system_id, env_id)
	values(?,?,?,?);`

	tx, err := testResourceRoleRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(
		sql,
		testResourceRoleNewRoleName,
		testResourceRoleMiddlewareClusterID,
		testResourceRoleMonitorSystemID,
		testResourceRoleOwnerID,
		testResourceRoleEnvID,
	)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select cluster_name from t_meta_mysql_cluster_info where cluster_name=?`
	result, err := tx.Execute(sql, testResourceRoleNewRoleName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	resourceRoleName, err := result.GetString(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if resourceRoleName != testResourceRoleNewRoleName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := testResourceRoleRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		if entity.GetRoleName() == testResourceRoleNewRoleName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestResourceRoleRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := testResourceRoleRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(2, len(entities), "test GetAll() failed")
}

func TestResourceRoleRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := testResourceRoleRepo.GetByID(testResourceRoleID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testResourceRoleRoleName, entity.GetRoleName(), "test GetByID() failed")
}

func TestResourceRoleRepo_GetID(t *testing.T) {
	asst := assert.New(t)

	id, err := testResourceRoleRepo.GetID(testResourceRoleRoleName)
	asst.Nil(err, common.CombineMessageWithError("test GetID() failed", err))
	asst.Equal(1, id, "test GetID() failed")
}

func TestResourceRoleRepo_GetByRoleUUID(t *testing.T) {
	asst := assert.New(t)

	entity, err := testResourceRoleRepo.GetByRoleUUID(testResourceRoleRoleUUID)
	asst.Nil(err, common.CombineMessageWithError("test GetByName() failed", err))
	asst.Equal(testResourceRoleRoleName, entity.GetRoleName(), "test GetByName() failed")
}

func TestResourceRoleRepo_GetResourceGroupByID(t *testing.T) {
	asst := assert.New(t)

	users, err := testResourceRoleRepo.GetUsersByID(testResourceRoleID)
	asst.Nil(err, common.CombineMessageWithError("test GetUsersByID() failed", err))
	asst.Equal(1, len(users), "test GetUsersByID() failed")
}

func TestResourceRoleRepo_GetUsersByID(t *testing.T) {
	asst := assert.New(t)

	users, err := testResourceRoleRepo.GetUsersByID(testResourceRoleID)
	asst.Nil(err, common.CombineMessageWithError("test GetUsersByID() failed", err))
	asst.Equal(1, len(users), "test GetUsersByID() failed")
}

func TestResourceRoleRepo_AddUser(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateResourceRole()
	asst.Nil(err, common.CombineMessageWithError("test AddResourceRoleUser() failed", err))
	err = testResourceRoleRepo.AddUser(entity.Identity(), testResourceRoleUserID)
	asst.Nil(err, common.CombineMessageWithError("test AddResourceRoleUser() failed", err))
	users, err := entity.GetUsers()
	asst.Nil(err, common.CombineMessageWithError("test AddResourceRoleUser() failed", err))
	asst.Equal(1, len(users), "test AddResourceRoleUser() failed")
	// delete
	err = testResourceRoleRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddResourceRoleUser() failed", err))
}

func TestResourceRoleRepo_DeleteUser(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateResourceRole()
	asst.Nil(err, common.CombineMessageWithError("test DeleteResourceRoleUser() failed", err))
	err = testResourceRoleRepo.DeleteUser(entity.Identity(), testResourceRoleUserID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteResourceRoleUser() failed", err))
	users, err := entity.GetUsers()
	asst.Nil(err, common.CombineMessageWithError("test DeleteResourceRoleUser() failed", err))
	asst.Zero(len(users), "test DeleteResourceRoleUser() failed")
	// delete
	err = testResourceRoleRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DeleteResourceRoleUser() failed", err))
}

func TestResourceRoleRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateResourceRole()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testResourceRoleRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestResourceRoleRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateResourceRole()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{resourceRoleRoleNameStruct: testResourceRoleUpdateRoleName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testResourceRoleRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = testResourceRoleRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testResourceRoleUpdateRoleName, entity.GetRoleName(), "test Update() failed")
	// delete
	err = testResourceRoleRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestResourceRoleRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateResourceRole()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testResourceRoleRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
