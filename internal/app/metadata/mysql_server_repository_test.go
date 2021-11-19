package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"

	"github.com/romberli/das/internal/dependency/metadata"
)

const (
	testMySQLServerNewServerName    = "test_new_mysql_server"
	testMySQLServerUpdateServerName = "test_update_mysql_server"
	testMySQLServerNewPortNum       = 33061
)

var testMySQLServerRepo *MySQLServerRepo

func init() {
	testInitDASMySQLPool()
	testMySQLServerRepo = NewMySQLServerRepoWithGlobal()
}

func testCreateMySQLServer() (metadata.MySQLServer, error) {
	mysqlServerInfo := NewMySQLServerInfoWithDefault(
		testMySQLServerClusterID,
		testMySQLServerNewServerName,
		testMySQLServerHostIP,
		testMySQLServerNewPortNum,
		testMySQLServerDeploymentType)
	entity, err := testMySQLServerRepo.Create(mysqlServerInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func TestMySQLServerRepoAll(t *testing.T) {
	TestMySQLServerRepo_Execute(t)
	TestMySQLServerRepo_Transaction(t)
	TestMySQLServerRepo_Create(t)
	TestMySQLServerRepo_GetAll(t)
	TestMySQLServerRepo_GetByClusterID(t)
	TestMySQLServerRepo_GetByID(t)
	TestMySQLServerRepo_GetByHostInfo(t)
	TestMySQLServerRepo_GetID(t)
	TestMySQLServerRepo_IsMaster(t)
	TestMySQLServerRepo_Update(t)
	TestMySQLServerRepo_Delete(t)
}

func TestMySQLServerRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := testMySQLServerRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestMySQLServerRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `
	insert into t_meta_mysql_server_info(
		cluster_id, server_name, service_name, host_ip, port_num, deployment_type, version)
	values(?,?,?,?,?,?,?);`

	tx, err := testMySQLServerRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(
		sql,
		testMySQLServerClusterID,
		testMySQLServerNewServerName,
		testMySQLServerServiceName,
		testMySQLServerHostIP,
		testMySQLServerNewPortNum,
		testMySQLServerDeploymentType,
		testMySQLServerVersion,
	)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select server_name from t_meta_mysql_server_info where host_ip=? and port_num=?`
	result, err := tx.Execute(sql, testMySQLServerHostIP, testMySQLServerNewPortNum)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	serverName, err := result.GetString(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if serverName != testMySQLServerNewServerName {
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

func TestMySQLServerRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := testMySQLServerRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(2, len(entities), "test GetAll() failed")
}

func TestMySQLServerRepo_GetByClusterID(t *testing.T) {
	asst := assert.New(t)

	entities, err := testMySQLServerRepo.GetByClusterID(testMySQLServerClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetByClusterID() failed", err))
	asst.Equal(1, len(entities), "test GetByClusterID() failed")
}

func TestMySQLServerRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := testMySQLServerRepo.GetByID(testMySQLServerClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testMySQLServerServerName, entity.GetServerName(), "test GetByID() failed")
}

func TestMySQLServerRepo_GetByHostInfo(t *testing.T) {
	asst := assert.New(t)

	entity, err := testMySQLServerRepo.GetByHostInfo(testMySQLServerHostIP, testMySQLServerPortNum)
	asst.Nil(err, common.CombineMessageWithError("test GetByHostInfo() failed", err))
	asst.Equal(testMySQLServerID, entity.Identity(), "test GetByHostInfo() failed")
}

func TestMySQLServerRepo_GetID(t *testing.T) {
	asst := assert.New(t)

	id, err := testMySQLServerRepo.GetID(testMySQLServerHostIP, testMySQLServerPortNum)
	asst.Nil(err, common.CombineMessageWithError("test GetID() failed", err))
	asst.Equal(testMySQLServerID, id, "test GetID() failed")
}

func TestMySQLServerRepo_IsMaster(t *testing.T) {
	asst := assert.New(t)

	isMaster, err := testMySQLServerRepo.IsMaster(testMySQLServerHostIP, testMySQLServerPortNum)
	asst.Nil(err, common.CombineMessageWithError("test IsMaster() failed", err))
	asst.True(isMaster, "test IsMaster() failed")
}

func TestMySQLServerRepo_GetMySQLClusterByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := testMySQLServerRepo.GetMySQLClusterByID(testMySQLServerID)
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLClusterByID() failed", err))

	asst.Equal(testMySQLServerClusterID, entity.Identity(), "test GetMySQLClusterByID() failed")
}

func TestMySQLServerRepo_GetMonitorSystem(t *testing.T) {
	asst := assert.New(t)

	monitorSystem, err := testMySQLServerRepo.GetMonitorSystem(testMySQLServerID)
	asst.Nil(err, common.CombineMessageWithError("test GetMonitorSystem() failed", err))
	asst.Equal(1, monitorSystem.Identity(), "test GetMonitorSystem() failed")
}

func TestMySQLServerRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testMySQLServerRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMySQLServerRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{mysqlServerServerNameStruct: testMySQLServerUpdateServerName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testMySQLServerRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = testMySQLServerRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testMySQLServerUpdateServerName, entity.GetServerName(), "test Update() failed")
	// delete
	err = testMySQLServerRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMySQLServerRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testMySQLServerRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
