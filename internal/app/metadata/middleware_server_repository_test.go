package metadata

import (
	"testing"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	testMiddlewareServerNewServerName    = "test_new_server_name"
	testMiddlewareServerUpdateServerName = "test_update_server_name"
	testMiddlewareServerNewPortNum       = 33062
)

var testMiddlewareServerRepo *MiddlewareServerRepo

func init() {
	testInitDASMySQLPool()
	testMiddlewareServerRepo = NewMiddlewareServerRepoWithGlobal()
}

func testCreateMiddlewareServer() (metadata.MiddlewareServer, error) {
	middlewareServerInfo := NewMiddlewareServerInfoWithDefault(
		testMiddlewareServerClusterID,
		testMiddlewareServerNewServerName,
		testMiddlewareServerMiddlewareRole,
		testMiddlewareServerHostIP,
		testMiddlewareServerNewPortNum,
	)
	entity, err := testMiddlewareServerRepo.Create(middlewareServerInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func TestMiddlewareServerRepoAll(t *testing.T) {
	TestMiddlewareServerRepo_Execute(t)
	TestMiddlewareClusterRepo_Transaction(t)
	TestMiddlewareServerRepo_GetAll(t)
	TestMiddlewareServerRepo_GetByClusterID(t)
	TestMiddlewareServerRepo_GetByID(t)
	TestMiddlewareServerRepo_GetByHostInfo(t)
	TestMiddlewareServerRepo_GetID(t)
	TestMiddlewareServerRepo_Create(t)
	TestMiddlewareServerRepo_Update(t)
	TestMiddlewareServerRepo_Delete(t)
}
func TestMiddlewareServerRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := testMiddlewareServerRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestMiddlewareServerRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_middleware_server_info(cluster_id, server_name, middleware_role, host_ip, port_num) values(?, ?, ?, ?, ?);`
	tx, err := testMiddlewareServerRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, testMiddlewareServerClusterID, testMiddlewareServerNewServerName, testMiddlewareServerMiddlewareRole, testMiddlewareServerHostIP, testMiddlewareServerPortNum)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select server_name from t_meta_middleware_server_info where server_name=?`
	result, err := tx.Execute(sql, testMiddlewareServerNewServerName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	middlewareServerName, err := result.GetString(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if middlewareServerName != testMiddlewareServerNewServerName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := testMiddlewareServerRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		if entity.GetServerName() == testMiddlewareServerNewServerName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestMiddlewareServerRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := testMiddlewareServerRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(1, len(entities), "test GetAll() failed")
}

func TestMiddlewareServerRepo_GetByClusterID(t *testing.T) {
	asst := assert.New(t)

	entities, err := testMiddlewareServerRepo.GetByClusterID(testMiddlewareServerClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetByClusterID() failed", err))
	asst.Equal(testMiddlewareServerServerName, entities[constant.ZeroInt].GetServerName(), "test GetByClusterID() failed")
}

func TestMiddlewareServerRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := testMiddlewareServerRepo.GetByID(testMiddlewareServerID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testMiddlewareServerServerName, entity.GetServerName(), "test GetByID() failed")
}

func TestMiddlewareServerRepo_GetByHostInfo(t *testing.T) {
	asst := assert.New(t)

	entity, err := testMiddlewareServerRepo.GetByHostInfo(testMiddlewareServerHostIP, testMiddlewareServerPortNum)
	asst.Nil(err, common.CombineMessageWithError("test GetByHostInfo() failed", err))
	asst.Equal(testMiddlewareServerID, entity.Identity(), "test GetByHostInfo() failed")
}

func TestMiddlewareServerRepo_GetID(t *testing.T) {
	asst := assert.New(t)

	id, err := testMiddlewareServerRepo.GetID(testMiddlewareServerHostIP, testMiddlewareServerPortNum)
	asst.Nil(err, common.CombineMessageWithError("test GetID() failed", err))
	asst.Equal(testMiddlewareServerID, id, "test GetID() failed")
}

func TestMiddlewareServerRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMiddlewareServer()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testMiddlewareServerRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMiddlewareServerRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMiddlewareServer()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{middlewareServerServerNameStruct: testMiddlewareServerUpdateServerName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testMiddlewareServerRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = testMiddlewareServerRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testMiddlewareServerUpdateServerName, entity.GetServerName(), "test Update() failed")
	// delete
	err = testMiddlewareServerRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMiddlewareServerRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMiddlewareServer()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = testMiddlewareServerRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
