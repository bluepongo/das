package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

var testMySQLServerService *MySQLServerService

func init() {
	testInitDASMySQLPool()
	testMySQLServerService = NewMySQLServerServiceWithDefault()
}

func TestMySQLServerServiceAll(t *testing.T) {
	TestMySQLServerService_GetMySQLServers(t)
	TestMySQLServerService_GetAll(t)
	TestMySQLServerService_GetByClusterID(t)
	TestMySQLServerService_GetByID(t)
	TestMySQLServerService_GetByHostInfo(t)
	TestMySQLServerService_IsMaster(t)
	TestMySQLServerService_GetMySQLClusterByID(t)
	TestMySQLServerService_Create(t)
	TestMySQLServerService_Update(t)
	TestMySQLServerService_Delete(t)
	TestMySQLServerService_Marshal(t)
	TestMySQLServerService_MarshalWithFields(t)
}

func TestMySQLServerService_GetMySQLServers(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLServerService.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLServers() failed", err))
	asst.Equal(2, len(testMySQLServerService.GetMySQLServers()), "test GetMySQLServers() failed")
}

func TestMySQLServerService_GetAll(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLServerService.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(2, len(testMySQLServerService.GetMySQLServers()), "test GetAll() failed")
}

func TestMySQLServerService_GetByClusterID(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLServerService.GetByClusterID(testMySQLServerClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetByClusterID() failed", err))
	asst.Equal(testMySQLServerID, testMySQLServerService.GetMySQLServers()[constant.ZeroInt].Identity(), "test GetByClusterID() failed")
}

func TestMySQLServerService_GetByID(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLServerService.GetByID(testMySQLServerID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testMySQLServerID, testMySQLServerService.GetMySQLServers()[constant.ZeroInt].Identity(), "test GetByID() failed")
}

func TestMySQLServerService_GetByHostInfo(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLServerService.GetByHostInfo(testMySQLServerHostIP, testMySQLServerPortNum)
	asst.Nil(err, common.CombineMessageWithError("test GetByHostInfo() failed", err))
	asst.Equal(testMySQLServerID, testMySQLServerService.GetMySQLServers()[constant.ZeroInt].Identity(), "test GetByHostInfo() failed")
}

func TestMySQLServerService_IsMaster(t *testing.T) {
	asst := assert.New(t)

	isMaster, err := testMySQLServerService.IsMaster(testMySQLServerHostIP, testMySQLServerPortNum)
	asst.Nil(err, common.CombineMessageWithError("test IsMaster() failed", err))
	asst.True(isMaster, "test IsMaster() failed")
}

func TestMySQLServerService_GetMySQLClusterByID(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLServerService.GetMySQLClusterByID(testMySQLServerID)
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLClusterByID() failed", err))
	asst.Equal(testMySQLServerClusterID, testMySQLServerService.GetMySQLCluster().Identity(), "test GetMySQLClusterByID() failed")
}

func TestMySQLServerService_Create(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLServerService.Create(map[string]interface{}{
		mysqlServerClusterIDStruct:      testMySQLServerClusterID,
		mysqlServerServerNameStruct:     testMySQLServerNewServerName,
		mysqlServerServiceNameStruct:    testMySQLServerServiceName,
		mysqlServerHostIPStruct:         testMySQLServerHostIP,
		mysqlServerPortNumStruct:        testMySQLServerNewPortNum,
		mysqlServerDeploymentTypeStruct: testMySQLServerDeploymentType,
		mysqlServerVersionStruct:        testMySQLServerVersion})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testMySQLServerService.Delete(testMySQLServerService.MySQLServers[constant.ZeroInt].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMySQLServerService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testMySQLServerService.Update(entity.Identity(), map[string]interface{}{
		mysqlServerServerNameStruct: testMySQLServerUpdateServerName,
	})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testMySQLServerUpdateServerName, testMySQLServerService.GetMySQLServers()[constant.ZeroInt].GetServerName(), "test Update() failed")
	// delete
	err = testMySQLServerService.Delete(testMySQLServerService.MySQLServers[constant.ZeroInt].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMySQLServerService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testMySQLServerService.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMySQLServerService_Marshal(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLServerService.GetByID(testMySQLServerID)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	jsonBytes, err := testMySQLServerService.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}

func TestMySQLServerService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLServerService.GetByID(testMySQLServerID)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	jsonBytes, err := testMySQLServerService.MarshalWithFields(mysqlServerServerNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	t.Log(string(jsonBytes))
}
