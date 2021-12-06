package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

var testMonitorSystemService *MonitorSystemService

func init() {
	testInitDASMySQLPool()
	testMonitorSystemService = NewMonitorSystemServiceWithDefault()
}

func TestMonitorSystemServiceAll(t *testing.T) {
	TestMonitorSystemService_GetMonitorSystems(t)
	TestMonitorSystemService_GetAll(t)
	TestMonitorSystemService_GetByID(t)
	TestMonitorSystemService_GetByEnv(t)
	TestMonitorSystemService_GetByHostInfo(t)
	TestMonitorSystemService_Create(t)
	TestMonitorSystemService_Update(t)
	TestMonitorSystemService_Delete(t)
	TestMonitorSystemService_Marshal(t)
	TestMonitorSystemService_MarshalWithFields(t)
}

func TestMonitorSystemService_GetMonitorSystems(t *testing.T) {
	asst := assert.New(t)

	err := testMonitorSystemService.GetAll()
	asst.Nil(err, "test GetMonitorSystems() failed")
	asst.Equal(2, len(testMonitorSystemService.GetMonitorSystems()), constant.ZeroInt, "test GetMonitorSystems() failed")
}

func TestMonitorSystemService_GetAll(t *testing.T) {
	asst := assert.New(t)

	err := testMonitorSystemService.GetAll()
	asst.Nil(err, "test GetAll() failed")
	asst.Equal(2, len(testMonitorSystemService.GetMonitorSystems()), constant.ZeroInt, "test GetAll() failed")
}

func TestMonitorSystemService_GetByID(t *testing.T) {
	asst := assert.New(t)

	err := testMonitorSystemService.GetByID(testMonitorSystemID)
	asst.Nil(err, "test GetByID() failed")
	asst.Equal(testMonitorSystemID, testMonitorSystemService.GetMonitorSystems()[constant.ZeroInt].Identity(), "test GetByID() failed")
}

func TestMonitorSystemService_GetByEnv(t *testing.T) {
	asst := assert.New(t)

	err := testMonitorSystemService.GetByEnv(testMonitorSystemID)
	asst.Nil(err, "test GetByEnv() failed")
	asst.Equal(testMonitorSystemID, testMonitorSystemService.GetMonitorSystems()[constant.ZeroInt].GetEnvID(), "test GetByEnv() failed")
}

func TestMonitorSystemService_GetByHostInfo(t *testing.T) {
	asst := assert.New(t)

	err := testMonitorSystemService.GetByHostInfo(testMonitorSystemHostIP, testMonitorSystemPortNum)
	asst.Nil(err, "test GetByHostInfo() failed")
	asst.Equal(testMonitorSystemID, testMonitorSystemService.GetMonitorSystems()[constant.ZeroInt].Identity(), "test GetByHostInfo() failed")
}

func TestMonitorSystemService_Create(t *testing.T) {
	asst := assert.New(t)

	err := testMonitorSystemService.Create(map[string]interface{}{
		monitorSystemSystemNameStruct:  testMonitorSystemNewSystemName,
		monitorSystemSystemTypeStruct:  testMonitorSystemSystemType,
		monitorSystemHostIPStruct:      testMonitorSystemHostIP,
		monitorSystemPortNumStruct:     testMonitorSystemNewPortNum,
		monitorSystemPortNumSlowStruct: testMonitorSystemPortNumSlow,
		monitorSystemBaseUrlStruct:     testMonitorSystemBaseUrl,
		monitorSystemEnvIDStruct:       testMonitorSystemEnvID,
	})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testMonitorSystemService.Delete(testMonitorSystemService.GetMonitorSystems()[constant.ZeroInt].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMonitorSystemService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testMonitorSystemService.Update(entity.Identity(), map[string]interface{}{monitorSystemSystemNameStruct: testMonitorSystemUpdateSystemName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testMonitorSystemService.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testMonitorSystemUpdateSystemName, testMonitorSystemService.GetMonitorSystems()[constant.ZeroInt].GetSystemName(), "test Update() failed")
	// delete
	err = testMonitorSystemService.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMonitorSystemService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testMonitorSystemService.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMonitorSystemService_Marshal(t *testing.T) {
	asst := assert.New(t)

	err := testMonitorSystemService.GetByID(testMonitorSystemID)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	jsonBytes, err := testMonitorSystemService.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}

func TestMonitorSystemService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	err := testMonitorSystemService.GetByID(testMonitorSystemID)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	jsonBytes, err := testMonitorSystemService.MarshalWithFields(monitorSystemMonitorSystemsStruct)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}
