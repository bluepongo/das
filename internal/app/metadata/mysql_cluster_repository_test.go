package metadata

import (
	"testing"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	testMySQLClusterNewClusterName    = "test_new_cluster_name"
	testMySQLClusterUpdateClusterName = "test_update_cluster_name"
	testMySQLClusterUserID            = 1
)

var testMySQLClusterRepo *MySQLClusterRepo

func init() {
	testInitDASMySQLPool()
	testMySQLClusterRepo = NewMySQLClusterRepoWithGlobal()
}

func testCreateMySQLCluster() (metadata.MySQLCluster, error) {
	mysqlClusterInfo := NewMySQLClusterInfoWithDefault(
		testMySQLClusterNewClusterName,
		testMySQLClusterEnvID)
	entity, err := testMySQLClusterRepo.Create(mysqlClusterInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func TestMySQLClusterRepoAll(t *testing.T) {
	TestMySQLClusterRepo_Execute(t)
	TestMySQLClusterRepo_Transaction(t)
	TestMySQLClusterRepo_Create(t)
	TestMySQLClusterRepo_GetAll(t)
	TestMySQLClusterRepo_GetByEnv(t)
	TestMySQLClusterRepo_GetByID(t)
	TestMySQLClusterRepo_GetByName(t)
	TestMySQLClusterRepo_GetID(t)
	TestMySQLClusterRepo_GetDBsByID(t)
	TestMySQLClusterRepo_GetResourceGroupByID(t)
	TestMySQLClusterRepo_GetUsersByID(t)
	TestMySQLClusterRepo_AddMySQLClusterUser(t)
	TestMySQLClusterRepo_DeleteMySQLClusterUser(t)
	TestMySQLClusterRepo_GetAppUsers(t)
	TestMySQLClusterRepo_GetDBUsers(t)
	TestMySQLClusterRepo_GetAllUsers(t)
	TestMySQLClusterRepo_Update(t)
	TestMySQLClusterRepo_Delete(t)
}

func TestMySQLClusterRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := testMySQLClusterRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestMySQLClusterRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `
	insert into t_meta_mysql_cluster_info(
		cluster_name, middleware_cluster_id, monitor_system_id, env_id)
	values(?,?,?,?);`

	tx, err := testMySQLClusterRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(
		sql,
		testMySQLClusterNewClusterName,
		testMySQLClusterMiddlewareClusterID,
		testMySQLClusterMonitorSystemID,
		testMySQLClusterEnvID,
	)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select cluster_name from t_meta_mysql_cluster_info where cluster_name=?`
	result, err := tx.Execute(sql, testMySQLClusterNewClusterName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	mysqlClusterName, err := result.GetString(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if mysqlClusterName != testMySQLClusterNewClusterName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := testMySQLClusterRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		if entity.GetClusterName() == testMySQLClusterNewClusterName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestMySQLClusterRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := testMySQLClusterRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(2, len(entities), "test GetAll() failed")
}

func TestMySQLClusterRepo_GetByEnv(t *testing.T) {
	asst := assert.New(t)

	entities, err := testMySQLClusterRepo.GetByEnv(testMySQLClusterEnvID)
	asst.Nil(err, common.CombineMessageWithError("test GetByEnv() failed", err))
	asst.Equal(2, len(entities), "test GetByEnv() failed")
}

func TestMySQLClusterRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := testMySQLClusterRepo.GetByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testMySQLClusterClusterName, entity.GetClusterName(), "test GetByID() failed")
}

func TestMySQLClusterRepo_GetByName(t *testing.T) {
	asst := assert.New(t)

	entity, err := testMySQLClusterRepo.GetByName(testMySQLClusterClusterName)
	asst.Nil(err, common.CombineMessageWithError("test GetByName() failed", err))
	asst.Equal(testMySQLClusterClusterName, entity.GetClusterName(), "test GetByName() failed")
}

func TestMySQLClusterRepo_GetID(t *testing.T) {
	asst := assert.New(t)

	id, err := testMySQLClusterRepo.GetID(testMySQLClusterClusterName)
	asst.Nil(err, common.CombineMessageWithError("test GetID() failed", err))
	asst.Equal(1, id, "test GetID() failed")
}

func TestMySQLClusterRepo_GetDBsByID(t *testing.T) {
	asst := assert.New(t)

	dbs, err := testMySQLClusterRepo.GetDBsByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBsByID() failed", err))
	asst.Equal(1, len(dbs), "test GetDBsByID() failed")
}

func TestMySQLClusterRepo_GetResourceGroupByID(t *testing.T) {
	asst := assert.New(t)

	resourceGroup, err := testMySQLClusterRepo.GetResourceGroupByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetResourceGroupByID() failed", err))
	asst.Equal(1, resourceGroup.Identity(), "test GetResourceGroupByID() failed")
}

func TestMySQLClusterRepo_GetUsersByID(t *testing.T) {
	asst := assert.New(t)

	users, err := testMySQLClusterRepo.GetUsersByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetUsersByID() failed", err))
	asst.Equal(1, len(users), "test GetUsersByID() failed")
}

func TestMySQLClusterRepo_AddMySQLClusterUser(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLClusterUser() failed", err))
	err = testMySQLClusterRepo.AddUser(entity.Identity(), testMySQLClusterUserID)
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLClusterUser() failed", err))
	users, err := entity.GetUsers()
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLClusterUser() failed", err))
	asst.Equal(1, len(users), "test AddMySQLClusterUser() failed")
	// delete
	err = testMySQLClusterRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLClusterUser() failed", err))
}

func TestMySQLClusterRepo_DeleteMySQLClusterUser(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test DeleteMySQLClusterUser() failed", err))
	err = testMySQLClusterRepo.DeleteUser(entity.Identity(), testMySQLClusterUserID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteMySQLClusterUser() failed", err))
	users, err := entity.GetUsers()
	asst.Nil(err, common.CombineMessageWithError("test DeleteMySQLClusterUser() failed", err))
	asst.Zero(len(users), "test DeleteMySQLClusterUser() failed")
	// delete
	err = testMySQLClusterRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DeleteMySQLClusterUser() failed", err))
}

func TestMySQLClusterRepo_GetAppUsers(t *testing.T) {
	asst := assert.New(t)

	users, err := testMySQLClusterRepo.GetAppUsersByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetAppUsers() failed", err))
	asst.Equal(1, users[constant.ZeroInt].Identity(), "test GetAppUsers() failed")
	asst.Equal(1, len(users), "test GetAppUsers() failed")
}

func TestMySQLClusterRepo_GetDBUsers(t *testing.T) {
	asst := assert.New(t)

	users, err := testMySQLClusterRepo.GetDBUsersByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBUsers() failed", err))
	asst.Equal(1, users[constant.ZeroInt].Identity(), "test GetDBUsers() failed")
}

func TestMySQLClusterRepo_GetAllUsers(t *testing.T) {
	asst := assert.New(t)

	users, err := testMySQLClusterRepo.GetAllUsersByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetAllUsers() failed", err))
	asst.Equal(1, users[constant.ZeroInt].Identity(), "test GetAllUsers() failed")
	asst.Equal(1, len(users), "test GetAllUsers() failed")
}

func TestMySQLClusterRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testMySQLClusterRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMySQLClusterRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{mysqlClusterClusterNameStruct: testMySQLClusterUpdateClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testMySQLClusterRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = testMySQLClusterRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testMySQLClusterUpdateClusterName, entity.GetClusterName(), "test Update() failed")
	// delete
	err = testMySQLClusterRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMySQLClusterRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testMySQLClusterRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
