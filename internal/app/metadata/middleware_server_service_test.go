package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

var testMiddlewareServerService *MiddlewareServerService

func init() {
	testInitDASMySQLPool()
	testMiddlewareServerService = NewMiddlewareServerServiceWithDefault()
}

func TestMiddlewareServerServiceAll(t *testing.T) {
	TestMiddlewareServerService_GetMiddlewareServers(t)
	TestMiddlewareServerService_GetAll(t)
	TestMiddlewareServerService_GetByClusterID(t)
	TestMiddlewareServerService_GetByID(t)
	TestMiddlewareServerService_GetByHostInfo(t)
	TestMiddlewareServerService_Create(t)
	TestMiddlewareServerService_Update(t)
	TestMiddlewareServerService_Delete(t)
	TestMiddlewareServerService_Marshal(t)
	TestMiddlewareServerService_MarshalWithFields(t)
}

func TestMiddlewareServerService_GetMiddlewareServers(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareServerService.GetAll()
	asst.Nil(err, "test GetMiddlewareServersByID() failed")
	asst.Equal(1, len(testMiddlewareServerService.GetMiddlewareServers()), "test GetMiddlewareServers() failed")
}

func TestMiddlewareServerService_GetAll(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareServerService.GetAll()
	asst.Nil(err, "test GetAll() failed")
	asst.Equal(1, len(testMiddlewareServerService.GetMiddlewareServers()), "test GetAll() failed")
}

func TestMiddlewareServerService_GetByClusterID(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareServerService.GetByClusterID(testMiddlewareServerClusterID)
	asst.Nil(err, "test GetByClusterID() failed")
	asst.Equal(1, len(testMiddlewareServerService.GetMiddlewareServers()), "test GetByClusterID() failed")
}

func TestMiddlewareServerService_GetByID(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareServerService.GetByID(testMiddlewareServerID)
	asst.Nil(err, "test GetByID() failed")
	asst.Equal(testMiddlewareServerServerName, testMiddlewareServerService.GetMiddlewareServers()[constant.ZeroInt].GetServerName(), "test GetByID() failed")
}

func TestMiddlewareServerService_GetByHostInfo(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareServerService.GetByHostInfo(testMiddlewareServerHostIP, testMiddlewareServerPortNum)
	asst.Nil(err, "test GetByHostInfo() failed")
	asst.Equal(testMiddlewareServerID, testMiddlewareServerService.GetMiddlewareServers()[constant.ZeroInt].Identity(), "test GetByHostInfo() failed")
}

func TestMiddlewareServerService_Create(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareServerService.Create(map[string]interface{}{
		middlewareServerClusterIDStruct:      testMiddlewareServerClusterID,
		middlewareServerServerNameStruct:     testMiddlewareServerNewServerName,
		middlewareServerMiddlewareRoleStruct: testMiddlewareServerMiddlewareRole,
		middlewareServerHostIPStruct:         testMiddlewareServerHostIP,
		middlewareServerPortNumStruct:        testMiddlewareServerNewPortNum,
	})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testMiddlewareServerService.Delete(testMiddlewareServerService.GetMiddlewareServers()[constant.ZeroInt].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMiddlewareServerService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMiddlewareServer()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testMiddlewareServerService.Update(entity.Identity(), map[string]interface{}{
		middlewareServerServerNameStruct: testMiddlewareServerUpdateServerName,
	})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testMiddlewareServerUpdateServerName, testMiddlewareServerService.GetMiddlewareServers()[constant.ZeroInt].GetServerName(), "test Update() failed")
	// delete
	err = testMiddlewareServerService.Delete(testMiddlewareServerService.GetMiddlewareServers()[constant.ZeroInt].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMiddlewareServerService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMiddlewareServer()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testMiddlewareServerService.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMiddlewareServerService_Marshal(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareServerService.GetByID(testMiddlewareServerID)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	jsonBytes, err := testMiddlewareServerService.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}

func TestMiddlewareServerService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareServerService.GetByID(testMiddlewareServerID)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	jsonBytes, err := testMiddlewareServerService.MarshalWithFields(middlewareServerMiddlewareServersStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	t.Log(string(jsonBytes))
}
