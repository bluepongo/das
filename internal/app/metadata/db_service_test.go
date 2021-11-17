package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

var testDBService *DBService

func init() {
	initDASMySQLPool()
	testDBService = NewDBServiceWithDefault()
}

func TestDBServiceAll(t *testing.T) {
	TestDBService_GetDBs(t)
	TestDBService_GetAll(t)
	TestDBService_GetByEnv(t)
	TestDBService_GetByID(t)
	TestDBService_GetByNameAndClusterInfo(t)
	TestDBService_GetAppsByID(t)
	TestDBService_GetMySQLClusterByID(t)
	TestDBService_GetAppOwnersByID(t)
	TestDBService_GetDBOwnersByID(t)
	TestDBService_GetAllOwnersByID(t)
	TestDBService_Create(t)
	TestDBService_Update(t)
	TestDBService_Delete(t)
	TestDBService_AddApp(t)
	TestDBService_DeleteApp(t)
	TestDBService_Marshal(t)
	TestDBService_MarshalWithFields(t)
}

func TestDBService_GetDBs(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetAll()
	asst.Nil(err, "test GetDBs() failed")
	asst.Equal(3, len(testDBService.GetDBs()), constant.ZeroInt, "test GetDBs() failed")
}

func TestDBService_GetAll(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetAll()
	asst.Nil(err, "test GetAll() failed")
	asst.Equal(3, len(testDBService.GetDBs()), constant.ZeroInt, "test GetAll() failed")
}

func TestDBService_GetByEnv(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetByEnv(testDBEnvID)
	asst.Nil(err, "test GetByEnv() failed")
	asst.Equal(testDBEnvID, testDBService.GetDBs()[constant.ZeroInt].GetEnvID(), "test GetByEnv() failed")
}

func TestDBService_GetByID(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetByID(testDBDBID)
	asst.Nil(err, "test GetByID() failed")
	asst.Equal(testDBDBName, testDBService.GetDBs()[constant.ZeroInt].GetDBName(), "test GetByID() failed")
}

func TestDBService_GetByNameAndClusterInfo(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetByNameAndClusterInfo(testDBDBName, testDBClusterID, testDBClusterType)
	asst.Nil(err, "test GetByID() failed")
	asst.Equal(testDBDBID, testDBService.GetDBs()[constant.ZeroInt].Identity(), "test GetByID() failed")
}

func TestDBService_GetAppsByID(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetAppsByID(testDBDBID)
	asst.Nil(err, "test GetAppsByID() failed")
	asst.Equal(testDBAppID, testDBService.GetApps()[constant.ZeroInt].Identity(), "test GetAppsByID() failed")
}

func TestDBService_GetMySQLClusterByID(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetMySQLClusterByID(testDBDBID)
	asst.Nil(err, "test GetMySQLClusterByID() failed")
	asst.Equal(testDBMySQLClusterID, testDBService.GetMySQLCluster().Identity(), "test GetMySQLClusterByID() failed")
}

func TestDBService_GetAppOwnersByID(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetAppOwnersByID(testDBDBID)
	asst.Nil(err, "test GetAppOwnersByID() failed")
	asst.Equal(testDBOwnerID, testDBService.GetOwners()[constant.ZeroInt].Identity(), "test GetAppOwnersByID() failed")
}

func TestDBService_GetDBOwnersByID(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetDBOwnersByID(testDBDBID)
	asst.Nil(err, "test GetDBOwnersByID() failed")
	asst.Equal(testDBOwnerID, testDBService.GetOwners()[constant.ZeroInt].Identity(), "test GetDBOwnersByID() failed")
}

func TestDBService_GetAllOwnersByID(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetAllOwnersByID(testDBDBID)
	asst.Nil(err, "test GetAllOwnersByID() failed")
	asst.Equal(testDBOwnerID, testDBService.GetOwners()[constant.ZeroInt].Identity(), "test GetAllOwnersByID() failed")
}

func TestDBService_Create(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.Create(
		map[string]interface{}{
			dbDBNameStruct:      testDBNewDBName,
			dbClusterIDStruct:   testDBClusterID,
			dbClusterTypeStruct: testDBClusterType,
			dbEnvIDStruct:       testDBEnvID,
		},
	)
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testDBService.Delete(testDBService.GetDBs()[constant.ZeroInt].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestDBService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateDB()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testDBService.Update(entity.Identity(), map[string]interface{}{dbDBNameStruct: testDBUpdateDBName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testDBService.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testDBUpdateDBName, testDBService.GetDBs()[constant.ZeroInt].GetDBName())
	// delete
	err = testDBService.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestDBService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateDB()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testDBService.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestDBService_AddApp(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.AddApp(testDBDBID, testDBNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	err = testDBService.GetAppsByID(testDBDBID)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	asst.Equal(2, len(testDBService.GetApps()), common.CombineMessageWithError("test AddApp() failed", err))
	err = testDBService.DeleteApp(testDBDBID, testDBNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
}

func TestDBService_DeleteApp(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.AddApp(testDBDBID, testDBNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	err = testDBService.DeleteApp(testDBDBID, testDBNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	err = testDBService.GetAppsByID(testDBDBID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	asst.Equal(1, len(testDBService.GetApps()), common.CombineMessageWithError("test AddApp() failed", err))

}

func TestDBService_Marshal(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	jsonBytes, err := testDBService.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}

func TestDBService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetByID(testDBDBID)
	jsonBytes, err := testDBService.MarshalWithFields(dbDBsStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	t.Log(string(jsonBytes))
}
