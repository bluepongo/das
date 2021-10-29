package metadata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

func initNewMySQLService() *MySQLClusterService {
	initMySQLClusterRepo()
	return NewMySQLClusterService(mysqlClusterRepo)
}

func TestMySQLClusterServiceAll(t *testing.T) {
	TestMySQLClusterService_GetMySQLServers(t)
	TestMySQLClusterService_GetAll(t)
	TestMySQLClusterService_GetByID(t)
	TestMySQLClusterService_Create(t)
	TestMySQLClusterService_Update(t)
	TestMySQLClusterService_Delete(t)
	TestMySQLClusterService_Marshal(t)
	TestMySQLClusterService_MarshalWithFields(t)
}

func TestMySQLClusterService_GetMySQLServers(t *testing.T) {
	asst := assert.New(t)

	s := initNewMySQLService()
	err := s.GetAll()
	asst.Nil(err, "test GetEnvs() failed")
	entities := s.GetMySQLClusters()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEnvs() failed")
}

func TestMySQLClusterService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := initNewMySQLService()
	err := s.GetAll()
	asst.Nil(err, "test GetEnvs() failed")
	entities := s.GetMySQLClusters()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEnvs() failed")
}

func TestMySQLClusterService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := initNewMySQLService()
	err := s.GetByID(1)
	asst.Nil(err, "test GetByID() failed")
	id := s.MySQLClusters[constant.ZeroInt].Identity()
	asst.Equal("1", id, "test GetByID() failed")
}

func TestMySQLClusterService_GetByName(t *testing.T) {
	asst := assert.New(t)

	s := initNewMySQLService()
	err := s.GetByName("name")
	asst.Nil(err, "test GetByName() failed")
	id := s.MySQLClusters[constant.ZeroInt].Identity()
	asst.Equal("name", id, "test GetByName() failed")
}

func TestMySQLClusterService_GetDBsByID(t *testing.T) {
	asst := assert.New(t)

	s := initNewMySQLService()
	err := s.GetDBsByID(1)
	asst.Nil(err, "test GetDBsByID() failed")
	id := s.Databases[constant.ZeroInt].Identity()
	asst.Equal(1, id, "test GetByID() failed")
}

func TestMySQLClusterService_GetAppOwnersByID(t *testing.T) {
	asst := assert.New(t)

	s := initNewMySQLService()
	err := s.GetAppOwnersByID(1)
	asst.Nil(err, "test GetAppOwnersByID() failed")
	id := s.Owners[constant.ZeroInt].Identity()
	asst.Equal(1, id, "test GetByID() failed")
}

func TestMySQLClusterService_GetDBOwnersByID(t *testing.T) {
	asst := assert.New(t)

	s := initNewMySQLService()
	err := s.GetDBOwnersByID(1)
	asst.Nil(err, "test GetDBOwnersByID() failed")
	id := s.Owners[constant.ZeroInt].Identity()
	asst.Equal(1, id, "test GetByID() failed")
}

func TestMySQLClusterService_GetAllOwnersByID(t *testing.T) {
	asst := assert.New(t)

	s := initNewMySQLService()
	err := s.GetAllOwnersByID(1)
	asst.Nil(err, "test GetAllOwnersByID() failed")
	id := s.Owners[constant.ZeroInt].Identity()
	asst.Equal(1, id, "test GetByID() failed")
}

func TestMySQLClusterService_Create(t *testing.T) {
	asst := assert.New(t)

	s := initNewMySQLService()

	err := s.Create(map[string]interface{}{
		clusterNameStruct:         testInsertClusterName,
		middlewareClusterIDStruct: defaultMySQLClusterInfoMiddlewareClusterID,
		monitorSystemIDStruct:     defaultMySQLClusterInfoMonitorSystemID,
		ownerIDStruct:             defaultMySQLClusterInfoOwnerID,
		envIDStruct:               defaultMySQLClusterInfoEnvID,
	})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMySQLClusterByID(s.MySQLClusters[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMySQLClusterService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := initNewMySQLService()
	err = s.Update(entity.Identity(), map[string]interface{}{clusterNameStruct: testUpdateClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	mysqlClusterName := s.GetMySQLClusters()[constant.ZeroInt].GetClusterName()
	asst.Equal(testUpdateClusterName, mysqlClusterName)
	// delete
	err = deleteMySQLClusterByID(s.MySQLClusters[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMySQLClusterService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := initNewMySQLService()
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMySQLClusterByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMySQLClusterService_Marshal(t *testing.T) {
	var entitiesUnmarshal []*MySQLClusterInfo

	asst := assert.New(t)

	s := initNewMySQLService()
	err := s.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	data, err := s.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	err = json.Unmarshal(data, &entitiesUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	entities := s.GetMySQLClusters()
	for i := 0; i < len(entities); i++ {
		entity := entities[i]
		entityUnmarshal := entitiesUnmarshal[i]
		asst.True(equalMySQLClusterInfo(entity.(*MySQLClusterInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestMySQLClusterService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := initNewMySQLService()
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteMySQLClusterByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
