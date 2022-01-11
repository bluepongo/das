package metadata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

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
	TestDBService_AddDBApp(t)
	TestDBService_DeleteDBApp(t)
	TestDBService_Marshal(t)
	TestDBService_MarshalWithFields(t)
}

func TestDBService_GetDBs(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test GetDBs() failed", err))
	s := NewDBService(dbRepo)
	err = s.GetAll()
	asst.Nil(err, "test GetDBs() failed")
	entities := s.GetDBs()
	asst.Greater(len(entities), constant.ZeroInt, "test GetDBs() failed")
	// delete
	err = deleteDBByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetDBs() failed", err))
}

func TestDBService_GetAll(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	s := NewDBService(dbRepo)
	err = s.GetAll()
	asst.Nil(err, "test GetAll() failed")
	entities := s.GetDBs()
	asst.Greater(len(entities), constant.ZeroInt, "test GetAll() failed")
	// delete
	err = deleteDBByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
}

func TestDBService_GetByEnv(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test GetByEnv() failed", err))
	s := NewDBService(dbRepo)
	err = s.GetByEnv(defaultDBInfoEnvID)
	asst.Nil(err, "test GetByEnv() failed")
	envId := s.DBs[constant.ZeroInt].GetEnvID()
	asst.Equal(defaultDBInfoEnvID, envId, "test GetByEnv() failed")
	// delete
	err = deleteDBByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetByEnv() failed", err))
}

func TestDBService_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	s := NewDBService(dbRepo)
	err = s.GetByID(entity.Identity())
	asst.Nil(err, "test GetByID() failed")
	dbName := s.DBs[constant.ZeroInt].GetDBName()
	asst.Equal(defaultDBInfoDBName, dbName, "test GetByID() failed")
	// delete
	err = deleteDBByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
}

func TestDBService_GetByNameAndClusterInfo(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	s := NewDBService(dbRepo)
	err = s.GetByNameAndClusterInfo(entity.GetDBName(), entity.GetClusterID(), entity.GetClusterType())
	asst.Nil(err, "test GetByID() failed")
	dbName := s.DBs[constant.ZeroInt].GetDBName()
	asst.Equal(defaultDBInfoDBName, dbName, "test GetByID() failed")
	// delete
	err = deleteDBByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
}

func TestDBService_GetAppsByID(t *testing.T) {
	asst := assert.New(t)

	s := NewDBService(dbRepo)
	err := s.GetAppsByID(1)
	asst.Nil(err, "test GetAppsByID() failed")
	apps := s.Apps
	asst.NotNil(apps, "test GetAppsByID() failed")
}

func TestDBService_GetMySQLClusterByID(t *testing.T) {
	asst := assert.New(t)

	s := NewDBService(dbRepo)
	err := s.GetMySQLClusterByID(1)
	asst.Nil(err, "test GetMySQLClusterByID() failed")
	mysqlCluster := s.MySQLCluster
	asst.NotNil(mysqlCluster, "test GetMySQLClusterByID() failed")
}

func TestDBService_GetAppOwnersByID(t *testing.T) {
	asst := assert.New(t)

	s := NewDBService(dbRepo)
	err := s.GetAppOwnersByID(1)
	asst.Nil(err, "test GetAppOwnersByID() failed")
	appOwners := s.AppOwners
	asst.NotNil(appOwners, "test GetAppOwnersByID() failed")
}

func TestDBService_GetDBOwnersByID(t *testing.T) {
	asst := assert.New(t)

	s := NewDBService(dbRepo)
	err := s.GetDBOwnersByID(1)
	asst.Nil(err, "test GetDBOwnersByID() failed")
	dbOwners := s.DBOwners
	asst.NotNil(dbOwners, "test GetDBOwnersByID() failed")
}

func TestDBService_GetAllOwnersByID(t *testing.T) {
	asst := assert.New(t)

	s := NewDBService(dbRepo)
	err := s.GetAllOwnersByID(1)
	asst.Nil(err, "test GetAllOwnersByID() failed")
	allOwners := s.AllOwners
	asst.NotNil(allOwners, "test GetAllOwnersByID() failed")
}

func TestDBService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewDBService(dbRepo)
	err := s.Create(map[string]interface{}{dbDBNameStruct: defaultDBInfoDBName,
		dbClusterIDStruct: defaultDBInfoClusterID, dbEnvIDStruct: defaultDBInfoEnvID})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteDBByID(s.DBs[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestDBService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewDBService(dbRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{dbDBNameStruct: newDBName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	dbName := s.GetDBs()[constant.ZeroInt].GetDBName()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newDBName, dbName)
	// delete
	err = deleteDBByID(s.DBs[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestDBService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewDBService(dbRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteDBByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestDBService_AddDBApp(t *testing.T) {
	asst := assert.New(t)

	s := NewDBService(dbRepo)

	err := s.AddApp(1, 3)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	apps, err := s.DBRepo.GetAppsByID(1)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	asst.NotNil(apps, common.CombineMessageWithError("test AddApp() failed", err))
}

func TestDBService_DeleteDBApp(t *testing.T) {
	asst := assert.New(t)

	s := NewDBService(dbRepo)
	err := s.DeleteApp(1, 3)
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	apps, err := s.DBRepo.GetAppsByID(1)
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	asst.Equal(0, len(apps), "test DeleteApp() failed")
}

func TestDBService_Marshal(t *testing.T) {
	var entitiesUnmarshal []*DBInfo

	asst := assert.New(t)

	s := NewDBService(dbRepo)
	err := s.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	data, err := s.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	err = json.Unmarshal(data, &entitiesUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	entities := s.GetDBs()
	for i := 0; i < len(entities); i++ {
		entity := entities[i]
		entityUnmarshal := entitiesUnmarshal[i]
		asst.True(dbEqual(entity.(*DBInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestDBService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createDB()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewDBService(dbRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(dbDBNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(dbDBNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteDBByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
