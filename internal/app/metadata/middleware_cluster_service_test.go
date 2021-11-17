package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

var testMiddlewareClusterService *MiddlewareClusterService

func init() {
	initDASMySQLPool()
	testMiddlewareClusterService = NewMiddlewareClusterServiceWithDefault()
}

func TestMiddlewareClusterServiceAll(t *testing.T) {
	TestMiddlewareClusterService_GetMiddlewareClusters(t)
	TestMiddlewareClusterService_GetMiddlewareServers(t)
	TestMiddlewareClusterService_GetAll(t)
	TestMiddlewareClusterService_GetByEnv(t)
	TestMiddlewareClusterService_GetByID(t)
	TestMiddlewareClusterService_GetByName(t)
	TestMiddlewareClusterService_GetMiddlewareServers(t)
	TestMiddlewareClusterService_Create(t)
	TestMiddlewareClusterService_Update(t)
	TestMiddlewareClusterService_Delete(t)
	TestMiddlewareClusterService_Marshal(t)
	TestMiddlewareClusterService_MarshalWithFields(t)
}

func TestMiddlewareClusterService_GetMiddlewareClusters(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareClusterService.GetAll()
	asst.Nil(err, "test GetMiddlewareClusters() failed")
	asst.Equal(1, len(testMiddlewareClusterService.GetMiddlewareClusters()), "test GetMiddlewareClusters() failed")
}

func TestMiddlewareClusterService_GetMiddlewareServers(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareClusterService.GetMiddlewareServersByID(testMiddlewareClusterClusterID)
	asst.Nil(err, "test GetMiddlewareClusters() failed")
	asst.Equal(1, len(testMiddlewareClusterService.GetMiddlewareServers()), "test GetMiddlewareServers() failed")
}

func TestMiddlewareClusterService_GetAll(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareClusterService.GetAll()
	asst.Nil(err, "test GetAll() failed")
	asst.Equal(1, len(testMiddlewareClusterService.GetMiddlewareClusters()), "test GetAll() failed")

}

func TestMiddlewareClusterService_GetByEnv(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareClusterService.GetByEnv(testMiddlewareClusterEnvID)
	asst.Nil(err, "test GetByEnv() failed")
	asst.Equal(testMiddlewareClusterEnvID, testMiddlewareClusterService.GetMiddlewareClusters()[constant.ZeroInt].GetEnvID(), "test GetByEnvID() failed")
}

func TestMiddlewareClusterService_GetByID(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareClusterService.GetByID(testMiddlewareClusterClusterID)
	asst.Nil(err, "test GetByID() failed")
	asst.Equal(testMiddlewareClusterClusterName, testMiddlewareClusterService.GetMiddlewareClusters()[constant.ZeroInt].GetClusterName(), "test GetByID() failed")
}

func TestMiddlewareClusterService_GetByName(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareClusterService.GetByName(testMiddlewareClusterClusterName)
	asst.Nil(err, "test GetByName() failed")
	asst.Equal(testMiddlewareClusterClusterName, testMiddlewareClusterService.GetMiddlewareClusters()[constant.ZeroInt].GetClusterName(), "test GetByName() failed")
}

func TestMiddlewareClusterService_GetMiddlewareServersByID(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareClusterService.GetMiddlewareServersByID(testMiddlewareClusterClusterID)
	asst.Nil(err, "test GetMiddlewareServersByID() failed")
	asst.Equal(1, len(testMiddlewareClusterService.GetMiddlewareServers()), "test GetMiddlewareServersByID() failed")
}

func TestMiddlewareClusterService_Create(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareClusterService.Create(map[string]interface{}{
		middlewareClusterClusterNameStruct: testMiddlewareClusterNewClusterName,
		middlewareClusterOwnerIDStruct:     testMiddlewareClusterOwnerID,
		middlewareClusterEnvIDStruct:       testMiddlewareClusterEnvID,
	})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testMiddlewareClusterService.Delete(testMiddlewareClusterService.GetMiddlewareClusters()[constant.ZeroInt].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMiddlewareClusterService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMiddlewareCluster()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testMiddlewareClusterService.Update(entity.Identity(), map[string]interface{}{
		middlewareClusterClusterNameStruct: testMiddlewareClusterUpdateClusterName,
	})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testMiddlewareClusterService.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testMiddlewareClusterUpdateClusterName, testMiddlewareClusterService.GetMiddlewareClusters()[constant.ZeroInt].GetClusterName(), "test Update() failed", err)
	// delete
	err = testMiddlewareClusterService.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMiddlewareClusterService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMiddlewareCluster()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testMiddlewareClusterService.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMiddlewareClusterService_Marshal(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareClusterService.GetByID(testMiddlewareClusterClusterID)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	jsonBytes, err := testMiddlewareClusterService.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}

func TestMiddlewareClusterService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareClusterService.GetByID(testMiddlewareClusterClusterID)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	jsonBytes, err := testMiddlewareClusterService.MarshalWithFields(middlewareClusterMiddlewareClustersStruct)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}
