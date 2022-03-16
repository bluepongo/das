package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

var testDBService *DBService

func init() {
	testInitDASMySQLPool()
	testDBService = NewDBServiceWithDefault()
}

func TestDBServiceAll(t *testing.T) {
	TestDBService_GetDBs(t)
	TestDBService_GetAll(t)
	TestDBService_GetByEnv(t)
	TestDBService_GetByID(t)
	TestDBService_GetDBByNameAndClusterInfo(t)
	TestDBService_GetDBByNameAndHostInfo(t)
	TestDBService_GetDBsByHostInfo(t)
	TestDBService_GetAppsByDBID(t)
	TestDBService_GetMySQLClusterByID(t)
	TestDBService_GetAppUsersByDBID(t)
	TestDBService_GetUsersByDBID(t)
	TestDBService_GetAllUsersByDBID(t)
	TestDBService_Create(t)
	TestDBService_Update(t)
	TestDBService_Delete(t)
	TestDBService_AddApp(t)
	TestDBService_DeleteApp(t)
	TestDBService_DBAddUser(t)
	TestDBService_DBDeleteUser(t)
	TestDBService_Marshal(t)
	TestDBService_MarshalWithFields(t)
}

func TestDBService_GetDBs(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetAll()
	asst.Nil(err, "test GetDBs() failed")
	asst.Equal(testDBAllDBNum, len(testDBService.GetDBs()), "test GetDBs() failed")
}

func TestDBService_GetAll(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetAll()
	asst.Nil(err, "test GetAll() failed")
	asst.Equal(testDBAllDBNum, len(testDBService.GetDBs()), "test GetAll() failed")
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
	asst.Equal(testDBDBName2, testDBService.GetDBs()[constant.ZeroInt].GetDBName(), "test GetByID() failed")
}

func TestDBService_GetDBByNameAndClusterInfo(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetDBByNameAndClusterInfo(testDBDBName2, testDBClusterID, testDBClusterType)
	asst.Nil(err, "test GetDBByNameAndClusterInfo() failed")
	asst.Equal(testDBDBID, testDBService.GetDBs()[constant.ZeroInt].Identity(), "test GetDBByNameAndClusterInfo() failed")
}

func TestDBService_GetDBByNameAndHostInfo(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetDBByNameAndHostInfo(testDBDBName2, testDBHostIP, testDBPortNum)
	asst.Nil(err, "test GetDBByNameAndHostInfo() failed")
	asst.Equal(testDBDBID, testDBService.GetDBs()[constant.ZeroInt].Identity(), "test GetDBByNameAndHostInfo() failed")
}

func TestDBService_GetDBsByHostInfo(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetDBsByHostInfo(testDBHostIP, testDBPortNum)
	asst.Nil(err, "test GetDBsByHostInfo() failed")
	asst.Equal(testDBDBID, len(testDBService.GetDBs()), "test GetDBsByHostInfo() failed")
}

func TestDBService_GetAppsByDBID(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetAppsByDBID(testDBDBID)
	asst.Nil(err, "test GetAppsByDBID() failed")
	asst.Equal(testDBAppID, testDBService.GetApps()[constant.ZeroInt].Identity(), "test GetAppsByDBID() failed")
}

func TestDBService_GetMySQLClusterByID(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetMySQLClusterByID(testDBDBID)
	asst.Nil(err, "test GetMySQLClusterByID() failed")
	asst.Equal(testDBMySQLClusterID, testDBService.GetMySQLCluster().Identity(), "test GetMySQLClusterByID() failed")
}

func TestDBService_GetAppUsersByDBID(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetAppUsersByDBID(testDBDBID)
	asst.Nil(err, "test GetAppUsersByDBID() failed")
	asst.Equal(1, testDBService.GetUsers()[constant.ZeroInt].Identity(), "test GetAppUsersByDBID() failed")
}

func TestDBService_GetUsersByDBID(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetUsersByDBID(testDBDBID)
	asst.Nil(err, "test GetUsersByDBID() failed")
	asst.Equal(2, testDBService.GetUsers()[constant.ZeroInt].Identity(), "test GetUsersByDBID() failed")
}

func TestDBService_GetAllUsersByDBID(t *testing.T) {
	asst := assert.New(t)

	err := testDBService.GetAllUsersByDBID(testDBDBID)
	asst.Nil(err, "test GetAllUsersByDBID() failed")
	asst.Equal(1, testDBService.GetUsers()[constant.ZeroInt].Identity(), "test GetAllUsersByDBID() failed")
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

	entity, err := testCreateDB()
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	err = testDBService.AddApp(entity.Identity(), testDBNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	err = testDBService.GetAppsByDBID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	asst.Equal(testDBNewAppID, testDBService.GetApps()[constant.ZeroInt].Identity())
	err = testDBService.DeleteApp(entity.Identity(), testDBNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	// delete
	err = testDBRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
}

func TestDBService_DeleteApp(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateDB()
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	err = testDBService.AddApp(entity.Identity(), testDBNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	err = testDBService.DeleteApp(entity.Identity(), testDBNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	err = testDBService.GetAppsByDBID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	asst.Equal(0, len(testDBService.GetApps()), "test DeleteApp() failed")
	// delete
	err = testDBRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
}

func TestDBService_DBAddUser(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateDB()
	asst.Nil(err, common.CombineMessageWithError("test DBAddUser() failed", err))
	err = testDBService.DBAddUser(entity.Identity(), testDBNewUserID)
	asst.Nil(err, common.CombineMessageWithError("test DBAddUser() failed", err))
	err = testDBService.GetUsersByDBID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DBAddUser() failed", err))
	asst.Equal(testDBNewUserID, testDBService.GetUsers()[constant.ZeroInt].Identity(), "test DBAddUser() failed")
	err = testDBService.DBDeleteUser(entity.Identity(), testDBNewUserID)
	asst.Nil(err, common.CombineMessageWithError("test DBAddUser() failed", err))
	// delete
	err = testDBRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DBAddUser() failed", err))
}

func TestDBService_DBDeleteUser(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateDB()
	asst.Nil(err, common.CombineMessageWithError("test DBDeleteUser() failed", err))
	err = testDBService.DBAddUser(entity.Identity(), testDBNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test DBDeleteUser() failed", err))
	err = testDBService.DBDeleteUser(entity.Identity(), testDBNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test DBDeleteUser() failed", err))
	err = testDBService.GetUsersByDBID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DBDeleteUser() failed", err))
	asst.Equal(0, len(testDBService.GetUsers()), "test DBDeleteUser() failed")
	// delete
	err = testDBRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DBDeleteUser() failed", err))
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
