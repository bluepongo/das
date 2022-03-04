package metadata

import (
	"testing"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	testMiddlewareClusterNewClusterName    = "test_new_cluster_name"
	testMiddlewareClusterUpdateClusterName = "test_update_cluster_name"
)

var testMiddlewareClusterRepo *MiddlewareClusterRepo

func init() {
	testInitDASMySQLPool()
	testMiddlewareClusterRepo = NewMiddlewareClusterRepoWithGlobal()
}

func testCreateMiddlewareCluster() (metadata.MiddlewareCluster, error) {
	middlewareClusterInfo := NewMiddlewareClusterInfoWithDefault(
		testMiddlewareClusterNewClusterName,
		testMiddlewareClusterEnvID,
	)
	entity, err := testMiddlewareClusterRepo.Create(middlewareClusterInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func TestMiddlewareClusterRepoAll(t *testing.T) {
	TestMiddlewareClusterRepo_Execute(t)
	TestMiddlewareClusterRepo_Transaction(t)
	TestMiddlewareClusterRepo_GetAll(t)
	TestMiddlewareClusterRepo_GetByEnv(t)
	TestMiddlewareClusterRepo_GetByID(t)
	TestMiddlewareClusterRepo_GetByName(t)
	TestMiddlewareClusterRepo_GetUsersByMiddlewareClusterID(t)
	TestMiddlewareClusterRepo_GetID(t)
	TestMiddlewareClusterRepo_Create(t)
	TestMiddlewareClusterRepo_Update(t)
	TestMiddlewareClusterRepo_Delete(t)
	TestMiddlewareClusterRepo_AddUser(t)
	TestMiddlewareClusterRepo_DeleteUser(t)

}
func TestMiddlewareClusterRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := testMiddlewareClusterRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestMiddlewareClusterRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_middleware_cluster_info(cluster_name, env_id) values(?, ?);`
	tx, err := testMiddlewareClusterRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, testMiddlewareClusterNewClusterName, testMiddlewareClusterEnvID)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select cluster_name from t_meta_middleware_cluster_info where cluster_name=?`
	result, err := tx.Execute(sql, testMiddlewareClusterNewClusterName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	middlewareClusterName, err := result.GetString(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if middlewareClusterName != testMiddlewareClusterNewClusterName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := testMiddlewareClusterRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		if entity.GetClusterName() == testMiddlewareClusterNewClusterName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestMiddlewareClusterRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := testMiddlewareClusterRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(1, len(entities), "test GetAll() failed")
}

func TestMiddlewareClusterRepo_GetByEnv(t *testing.T) {
	asst := assert.New(t)

	entities, err := testMiddlewareClusterRepo.GetByEnv(testMiddlewareClusterEnvID)
	asst.Nil(err, common.CombineMessageWithError("test GetByEnv() failed", err))
	asst.Equal(1, len(entities), "test GetByEnv() failed")
}

func TestMiddlewareClusterRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := testMiddlewareClusterRepo.GetByID(testMiddlewareClusterClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testMiddlewareClusterClusterName, entity.GetClusterName(), "test GetByID() failed")
}

func TestMiddlewareClusterRepo_GetByName(t *testing.T) {
	asst := assert.New(t)

	entity, err := testMiddlewareClusterRepo.GetByName(testMiddlewareClusterClusterName)
	asst.Nil(err, common.CombineMessageWithError("test GetByName() failed", err))
	asst.Equal(testMiddlewareClusterClusterName, entity.GetClusterName(), "test GetByID() failed")
}

func TestMiddlewareClusterRepo_GetID(t *testing.T) {
	asst := assert.New(t)

	id, err := testMiddlewareClusterRepo.GetID(testMiddlewareClusterClusterName, testMiddlewareClusterEnvID)
	asst.Nil(err, common.CombineMessageWithError("test GetID() failed", err))
	asst.Equal(testMiddlewareClusterClusterID, id, "test GetID() failed")
}

func TestMiddlewareClusterRepo_GetUsersByMiddlewareClusterID(t *testing.T) {
	asst := assert.New(t)

	users, err := testMiddlewareClusterRepo.GetUsersByMiddlewareClusterID(testMiddlewareClusterClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetUsersByMiddlewareClusterID() failed", err))
	asst.Equal(1, len(users), "test GetUsersByMiddlewareClusterID() failed", err)
}

func TestMiddlewareClusterRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMiddlewareCluster()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	asst.Equal(testMiddlewareClusterNewClusterName, entity.GetClusterName(), "test Create() failed")
	// delete
	err = testMiddlewareClusterRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMiddlewareClusterRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMiddlewareCluster()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{middlewareClusterClusterNameStruct: testMiddlewareClusterUpdateClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testMiddlewareClusterRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = testMiddlewareClusterRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testMiddlewareClusterUpdateClusterName, entity.GetClusterName(), "test Update() failed")
	// delete
	err = testMiddlewareClusterRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMiddlewareClusterRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMiddlewareCluster()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = testMiddlewareClusterRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMiddlewareClusterRepo_AddUser(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMiddlewareCluster()
	asst.Nil(err, common.CombineMessageWithError("test MiddlewareClusterAddUser() failed", err))
	err = testMiddlewareClusterRepo.AddUser(entity.Identity(), testMiddlewareClusterNewUserID)
	asst.Nil(err, common.CombineMessageWithError("test MiddlewareClusterAddUser() failed", err))
	users, err := entity.GetUsersByMiddlewareClusterID()
	asst.Nil(err, common.CombineMessageWithError("test MiddlewareClusterAddUser() failed", err))
	asst.Equal(1, len(users), "test MiddlewareClusterAddUser() failed", err)
	// delete
	err = testMiddlewareClusterRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test MiddlewareClusterAddUser() failed", err))
}

func TestMiddlewareClusterRepo_DeleteUser(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMiddlewareCluster()
	asst.Nil(err, common.CombineMessageWithError("test MiddlewareClusterDeleteUser() failed", err))
	err = testMiddlewareClusterRepo.DeleteUser(entity.Identity(), testUserID)
	asst.Nil(err, common.CombineMessageWithError("test MiddlewareClusterDeleteUser() failed", err))
	users, err := entity.GetUsersByMiddlewareClusterID()
	asst.Nil(err, common.CombineMessageWithError("test MiddlewareClusterDeleteUser() failed", err))
	asst.Zero(len(users), "test MiddlewareClusterDeleteUser() failed")
	// delete
	err = testMiddlewareClusterRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test MiddlewareClusterDeleteUser() failed", err))
}
