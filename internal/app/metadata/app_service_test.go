package metadata

import (
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

var testAppService *AppService

func init() {
	testAppService = NewAppServiceWithDefault()
}

func TestAppServiceAll(t *testing.T) {
	TestAppService_GetEntities(t)
	TestAppService_GetAll(t)
	TestAppService_GetByID(t)
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

func TestAppService_GetDBSByID(t *testing.T) {
	asst := assert.New(t)

	s := NewAppService(testAppRepo)
	err := s.GetDBsByID(testAppAppID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBSByID() failed", err))
	asst.Equal(1, len(s.DBs), common.CombineMessageWithError("test GetDBSByID() failed", err))
}

func TestAppService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewAppService(testAppRepo)
	err := s.Create(map[string]interface{}{appAppNameStruct: testAppAppName, appLevelStruct: testAppLevel})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteAppByID(s.Apps[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestAppService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createApp()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewAppService(testAppRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{appAppNameStruct: testAppAppName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	appName := s.Apps[constant.ZeroInt].GetAppName()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testAppAppName, appName, common.CombineMessageWithError("test Update() failed", err))
	// delete
	err = deleteAppByID(s.Apps[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestAppService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createApp()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewAppService(testAppRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteAppByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestAppService_Marshal(t *testing.T) {
	asst := assert.New(t)

	s := NewAppService(testAppRepo)
	err := s.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	data, err := s.Marshal()
	t.Log(string(data))
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
}

func TestAppService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createApp()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewAppService(testAppRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(appAppsStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf(`{"apps":[%s]}`, string(dataEntity)))
	// delete
	err = deleteAppByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
}

func TestAppService_AddDB(t *testing.T) {
	asst := assert.New(t)
	entity, err := createApp()
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
	s := NewAppService(testAppRepo)
	s.GetDBsByID(1)
	asst.Equal(nil, err, "test AddDB() failed")
	dbID := s.DBs[0].Identity()
	asst.Nil(err, common.CombineMessageWithError("entity.AddDB() failed", err))
	err = s.AddDB(entity.Identity(), dbID)
	deleteAppByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
}

func TestAppService_DeleteDB(t *testing.T) {
	asst := assert.New(t)
	entity, err := createApp()
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
	s := NewAppService(testAppRepo)
	s.GetDBsByID(1)
	asst.Equal(nil, err, "test DeleteDB() failed")
	dbID := s.DBs[0].Identity()
	asst.Nil(err, common.CombineMessageWithError("entity.DeleteDB() failed", err))
	err = s.DeleteDB(entity.Identity(), dbID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
	err = deleteAppByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
}
