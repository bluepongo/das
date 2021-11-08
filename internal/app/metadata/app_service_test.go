package metadata

import (
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

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

	s := NewAppService(appRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEnvs() failed")
	entities := s.GetApps()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEnvs() failed")
}

func TestAppService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewAppService(appRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEnvs() failed")
	entities := s.GetApps()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEnvs() failed")
}

func TestAppService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewAppService(appRepo)
	err := s.GetByID(2)
	asst.Nil(err, "test GetByID() failed")
	id := s.Apps[constant.ZeroInt].Identity()
	asst.Equal(2, id, "test GetByID() failed")
}

func TestAppService_GetDBSByID(t *testing.T) {
	asst := assert.New(t)

	s := NewAppService(appRepo)
	err := s.GetDBsByID(1)
	asst.Nil(err, "test GetDBsByID() failed")
	asst.Equal(2, len(s.DBs), "test GetDBsByID() failed")

}

func TestAppService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewAppService(appRepo)
	err := s.Create(map[string]interface{}{appAppNameStruct: defaultAppInfoAppName, appLevelStruct: defaultAppInfoLevel})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteAppByID(s.Apps[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestAppService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createApp()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewAppService(appRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{appAppNameStruct: newAppName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	appName := s.Apps[constant.ZeroInt].GetAppName()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newAppName, appName)
	// delete
	err = deleteAppByID(s.Apps[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestAppService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createApp()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewAppService(appRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteAppByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestAppService_Marshal(t *testing.T) {
	asst := assert.New(t)

	s := NewAppService(appRepo)
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
	s := NewAppService(appRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(appAppsStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf(`{"apps":[%s]}`, string(dataEntity)))
	// delete
	err = deleteAppByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestAppService_AddDB(t *testing.T) {
	asst := assert.New(t)
	entity, err := createApp()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewAppService(appRepo)
	s.GetDBsByID(1)
	asst.Equal(nil, err, "test DeleteDB() failed")
	dbID := s.DBs[0].Identity()
	asst.Nil(err, common.CombineMessageWithError("entity.GetDBIDList() failed", err))
	err = s.AddDB(entity.Identity(), dbID)
	deleteAppByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
}

func TestAppService_DeleteDB(t *testing.T) {
	asst := assert.New(t)
	entity, err := createApp()
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
	s := NewAppService(appRepo)
	s.GetDBsByID(1)
	asst.Equal(nil, err, "test DeleteDB() failed")
	dbID := s.DBs[0].Identity()
	asst.Nil(err, common.CombineMessageWithError("entity.GetDBIDList() failed", err))
	err = s.DeleteDB(entity.Identity(), dbID)
	deleteAppByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
}
