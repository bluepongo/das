package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

var testMySQLClusterService *MySQLClusterService

func init() {
	initDASMySQLPool()
	testMySQLClusterService = NewMySQLClusterServiceWithDefault()
}

func TestMySQLClusterServiceAll(t *testing.T) {
	TestMySQLClusterService_GetMySQLClusters(t)
	TestMySQLClusterService_GetAll(t)
	TestMySQLClusterService_GetByID(t)
	TestMySQLClusterService_GetByName(t)
	TestMySQLClusterService_GetMySQLServersByID(t)
	TestMySQLClusterService_GetMasterServersByID(t)
	TestMySQLClusterService_GetDBsByID(t)
	TestMySQLClusterService_GetAppOwnersByID(t)
	TestMySQLClusterService_GetDBOwnersByID(t)
	TestMySQLClusterService_GetAllOwnersByID(t)
	TestMySQLClusterService_Create(t)
	TestMySQLClusterService_Update(t)
	TestMySQLClusterService_Delete(t)
	TestMySQLClusterService_Marshal(t)
	TestMySQLClusterService_MarshalWithFields(t)
}

func TestMySQLClusterService_GetMySQLClusters(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLServers() failed", err))
	asst.Equal(2, len(testMySQLClusterService.GetMySQLClusters()), "test GetMySQLServers() failed")
}

func TestMySQLClusterService_GetAll(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(2, len(testMySQLClusterService.GetMySQLClusters()), "test GetAll() failed")
}

func TestMySQLClusterService_GetByID(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.GetByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testMySQLClusterID, testMySQLClusterService.GetMySQLClusters()[constant.ZeroInt].Identity(), "test GetByID() failed")
}

func TestMySQLClusterService_GetByName(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.GetByName(testMySQLClusterClusterName)
	asst.Nil(err, common.CombineMessageWithError("test GetByName() failed", err))
	asst.Equal(testMySQLClusterClusterName, testMySQLClusterService.GetMySQLClusters()[constant.ZeroInt].GetClusterName(), "test GetByName() failed")
}

func TestMySQLClusterService_GetMySQLServersByID(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.GetMySQLServersByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLServersByID() failed", err))
	asst.Equal(1, len(testMySQLClusterService.GetMySQLServers()), "test GetMySQLServersByID() failed")
}

func TestMySQLClusterService_GetMasterServersByID(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.GetMasterServersByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetMasterServersByID() failed", err))
	asst.Equal(1, len(testMySQLClusterService.GetMySQLServers()), "test GetMasterServersByID() failed")
}

func TestMySQLClusterService_GetDBsByID(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.GetDBsByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBsByID() failed", err))
	asst.Equal(2, len(testMySQLClusterService.GetDBs()), "test GetDBsByID() failed")
}

func TestMySQLClusterService_GetAppOwnersByID(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.GetAppOwnersByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetAppOwnersByID() failed", err))
	asst.Equal(1, testMySQLClusterService.GetOwners()[constant.ZeroInt].Identity(), "test GetAppOwnersByID() failed")
}

func TestMySQLClusterService_GetDBOwnersByID(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.GetDBOwnersByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBOwnersByID() failed", err))
	asst.Equal(1, testMySQLClusterService.GetOwners()[constant.ZeroInt].Identity(), "test GetDBOwnersByID() failed")
}

func TestMySQLClusterService_GetAllOwnersByID(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.GetAllOwnersByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetAllOwnersByID() failed", err))
	asst.Equal(1, testMySQLClusterService.GetOwners()[constant.ZeroInt].Identity(), "test GetAllOwnersByID() failed")
}

func TestMySQLClusterService_Create(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.Create(map[string]interface{}{
		mysqlClusterClusterNameStruct:         testMySQLClusterNewClusterName,
		mysqlCLusterMiddlewareClusterIDStruct: testMySQLClusterMiddlewareClusterID,
		mysqlClusterMonitorSystemIDStruct:     testMySQLClusterMonitorSystemID,
		mysqlClusterOwnerIDStruct:             testMySQLClusterOwnerID,
		mysqlClusterEnvIDStruct:               testMySQLClusterEnvID,
	})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testMySQLClusterService.Delete(testMySQLClusterService.GetMySQLClusters()[constant.ZeroInt].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMySQLClusterService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))

	err = testMySQLClusterService.Update(entity.Identity(), map[string]interface{}{mysqlClusterClusterNameStruct: testMySQLClusterUpdateClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testMySQLClusterService.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testMySQLClusterUpdateClusterName, testMySQLClusterService.GetMySQLClusters()[constant.ZeroInt].GetClusterName(), "test Update() failed")
	// delete
	err = testMySQLClusterService.Delete(testMySQLClusterService.GetMySQLClusters()[constant.ZeroInt].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMySQLClusterService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testMySQLClusterService.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMySQLClusterService_Marshal(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testMySQLClusterService.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}

func TestMySQLClusterService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testMySQLClusterService.MarshalWithFields(mysqlClusterMySQLClustersStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	t.Log(string(jsonBytes))
}
