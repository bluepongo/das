package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

var testAppService *AppService

func init() {
	initDASMySQLPool()
	testAppRepo = NewAppRepoWithGlobal()
	testAppService = NewAppServiceWithDefault()
}

func TestAppServiceAll(t *testing.T) {
	TestAppService_GetEntities(t)
	TestAppService_GetAll(t)
	TestAppService_GetByID(t)
	TestAppService_GetDBsByID(t)
	TestAppService_Create(t)
	TestAppService_Update(t)
	TestAppService_Delete(t)
	TestAppService_Marshal(t)
	TestAppService_MarshalWithFields(t)
	TestAppService_DeleteDB(t)
	TestAppService_AddDB(t)
}

func TestAppService_GetEntities(t *testing.T) {
	asst := assert.New(t)

	err := testAppService.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := testAppService.GetApps()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestAppService_GetAll(t *testing.T) {
	asst := assert.New(t)

	err := testAppService.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	entities := testAppService.GetApps()
	asst.Greater(len(entities), constant.ZeroInt, common.CombineMessageWithError("test GetAll() failed", err))
}

func TestAppService_GetByID(t *testing.T) {
	asst := assert.New(t)

	err := testAppService.GetByID(testAppAppID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	id := testAppService.GetApps()[constant.ZeroInt].Identity()
	asst.Equal(testAppAppID, id, common.CombineMessageWithError("test GetByID() failed", err))
}

func TestAppService_GetDBsByID(t *testing.T) {
	asst := assert.New(t)

	err := testAppService.GetDBsByID(testAppAppID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBSByID() failed", err))
	asst.Equal(1, len(testAppService.DBs), common.CombineMessageWithError("test GetDBSByID() failed", err))
}

func TestAppService_Create(t *testing.T) {
	asst := assert.New(t)

	err := testAppService.Create(map[string]interface{}{appAppNameStruct: testAppNewAppName, appLevelStruct: testAppLevel})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	err = testAppService.GetAppByName(testAppNewAppName)
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testAppRepo.Delete(testAppService.GetApps()[constant.ZeroInt].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestAppService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateApp()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testAppService.Update(entity.Identity(), map[string]interface{}{appAppNameStruct: testAppUpdateAppName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testAppService.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	appName := testAppService.GetApps()[constant.ZeroInt].GetAppName()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testAppUpdateAppName, appName, common.CombineMessageWithError("test Update() failed", err))
	// delete
	err = testAppRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestAppService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateApp()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testAppService.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestAppService_Marshal(t *testing.T) {
	asst := assert.New(t)

	err := testAppService.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	data, err := testAppService.Marshal()
	t.Log(string(data))
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
}

func TestAppService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateApp()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	err = testAppService.GetByID(entity.Identity())
	appsBytes, err := testAppService.MarshalWithFields(appAppsStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	t.Log(string(appsBytes))
	// delete
	err = testAppRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
}

func TestAppService_AddDB(t *testing.T) {
	asst := assert.New(t)
	entity, err := testCreateApp()
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
	err = testAppService.AddDB(entity.Identity(), testAppDBID)
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
	err = testAppService.GetDBsByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
	asst.Equal(testAppDBID, testAppService.GetDBs()[constant.ZeroInt].Identity())
	err = testAppService.DeleteDB(entity.Identity(), testAppDBID)
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
	err = testAppRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
}

func TestAppService_DeleteDB(t *testing.T) {
	asst := assert.New(t)
	entity, err := testCreateApp()
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
	err = testAppService.AddDB(entity.Identity(), testAppDBID)
	asst.Nil(err, "test DeleteDB() failed")
	err = testAppService.DeleteDB(entity.Identity(), testAppDBID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
	err = testAppService.GetDBsByID(entity.Identity())
	asst.Zero(len(testAppService.GetDBs()))
	err = testAppRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
}
