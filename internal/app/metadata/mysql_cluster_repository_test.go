package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"

	"github.com/romberli/das/internal/dependency/metadata"
)

const (
	testMySQLClusterNewClusterName    = "test_new_cluster_name"
	testMySQLClusterUpdateClusterName = "test_update_cluster_name"
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
	TestMySQLClusterRepo_GetAppOwners(t)
	TestMySQLClusterRepo_GetDBOwners(t)
	TestMySQLClusterRepo_GetAllOwners(t)
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
		cluster_name, middleware_cluster_id, monitor_system_id, owner_id, env_id)
	values(?,?,?,?,?);`

	tx, err := testMySQLClusterRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(
		sql,
		testMySQLClusterNewClusterName,
		testMySQLClusterMiddlewareClusterID,
		testMySQLClusterMonitorSystemID,
		testMySQLClusterOwnerID,
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
	asst.Equal(2, len(dbs), "test GetDBsByID() failed")
}

func TestMySQLClusterRepo_GetAppOwners(t *testing.T) {
	asst := assert.New(t)

	owners, err := testMySQLClusterRepo.GetAppOwnersByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetAppOwners() failed", err))
	asst.Equal(14, owners[constant.ZeroInt].Identity(), "test GetAppOwners() failed")
	asst.Equal(2, len(owners), "test GetAppOwners() failed")
}

func TestMySQLClusterRepo_GetDBOwners(t *testing.T) {
	asst := assert.New(t)

	owners, err := testMySQLClusterRepo.GetDBOwnersByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBOwners() failed", err))
	asst.Equal(1, owners[constant.ZeroInt].Identity(), "test GetDBOwners() failed")
}

func TestMySQLClusterRepo_GetAllOwners(t *testing.T) {
	asst := assert.New(t)

	owners, err := testMySQLClusterRepo.GetAllOwnersByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetAllOwners() failed", err))
	asst.Equal(14, owners[constant.ZeroInt].Identity(), "test GetAllOwners() failed")
	asst.Equal(3, len(owners), "test GetAllOwners() failed")
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
