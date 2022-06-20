package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

var testMySQLClusterService *MySQLClusterService

func init() {
	testInitDASMySQLPool()
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
	TestMySQLClusterService_GetResourceGroupByID(t)
	TestMySQLClusterService_AddUser(t)
	TestMySQLClusterService_DeleteUser(t)
	TestMySQLClusterService_GetAppUsersByID(t)
	TestMySQLClusterService_GetDBUsersByID(t)
	TestMySQLClusterService_GetAllUsersByID(t)
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

func TestMySQLClusterService_GetResourceGroupByID(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.GetResourceGroupByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetResourceGroupByID() failed", err))
	asst.Equal(1, testMySQLClusterService.GetResourceGroup().Identity(), "test GetResourceGroupByID() failed")
}

func TestMySQLClusterService_GetUsersByID(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.GetUsersByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBSByID() failed", err))
	asst.Equal(1, len(testMySQLClusterService.Users), common.CombineMessageWithError("test GetDBSByID() failed", err))
}

func TestMySQLClusterService_AddUser(t *testing.T) {
	asst := assert.New(t)
	entity, err := testCreateMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
	err = testMySQLClusterService.AddUser(entity.Identity(), testMySQLClusterUserID)
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
	err = testMySQLClusterService.GetUsersByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
	asst.Equal(testMySQLClusterUserID, testMySQLClusterService.GetUsers()[constant.ZeroInt].Identity())
	err = testMySQLClusterService.DeleteUser(entity.Identity(), testMySQLClusterUserID)
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
	err = testMySQLClusterRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
}

func TestMySQLClusterService_DeleteUser(t *testing.T) {
	asst := assert.New(t)
	entity, err := testCreateMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test DeleteUser() failed", err))
	err = testMySQLClusterService.AddUser(entity.Identity(), testMySQLClusterUserID)
	asst.Nil(err, "test DeleteUser() failed")
	err = testMySQLClusterService.DeleteUser(entity.Identity(), testMySQLClusterUserID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteUser() failed", err))
	err = testMySQLClusterService.GetUsersByID(entity.Identity())
	asst.Zero(len(testMySQLClusterService.GetUsers()))
	err = testMySQLClusterRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DeleteUser() failed", err))
}

func TestMySQLClusterService_GetAppUsersByID(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.GetAppUsersByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetAppUsersByID() failed", err))
	asst.Equal(15, testMySQLClusterService.Users[constant.ZeroInt].Identity(), "test GetAppUsersByID() failed")
	asst.Equal(1, len(testMySQLClusterService.Users), "test GetAppUsersByID() failed")
}

func TestMySQLClusterService_GetDBUsersByID(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.GetDBUsersByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBUsersByID() failed", err))
	asst.Equal(1, testMySQLClusterService.Users[constant.ZeroInt].Identity(), "test GetDBUsersByID() failed")
}

func TestMySQLClusterService_GetAllUsersByID(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.GetAllUsersByID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetAllUsersByID() failed", err))
	asst.Equal(15, testMySQLClusterService.Users[constant.ZeroInt].Identity(), "test GetAllUsersByID() failed")
	asst.Equal(2, len(testMySQLClusterService.Users), "test GetAllOwnersByID() failed")
}

func TestMySQLClusterService_Create(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterService.Create(map[string]interface{}{
		mysqlClusterClusterNameStruct:         testMySQLClusterNewClusterName,
		mysqlCLusterMiddlewareClusterIDStruct: testMySQLClusterMiddlewareClusterID,
		mysqlClusterMonitorSystemIDStruct:     testMySQLClusterMonitorSystemID,
		// mysqlClusterOwnerIDStruct:             testMySQLClusterOwnerID,
		mysqlClusterEnvIDStruct: testMySQLClusterEnvID,
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
