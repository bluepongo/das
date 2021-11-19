package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

var testEnvService *EnvService

func init() {
	testInitDASMySQLPool()
	testEnvService = NewEnvServiceWithDefault()
}

func TestEnvServiceAll(t *testing.T) {
	TestEnvService_GetEnvs(t)
	TestEnvService_GetAll(t)
	TestEnvService_GetByID(t)
	TestEnvService_GetEnvByName(t)
	TestEnvService_Create(t)
	TestEnvService_Update(t)
	TestEnvService_Delete(t)
	TestEnvService_Marshal(t)
	TestEnvService_MarshalWithFields(t)
}

func TestEnvService_GetEnvs(t *testing.T) {
	asst := assert.New(t)

	err := testEnvService.GetAll()
	asst.Nil(err, "test GetEnvs() failed")
	asst.Equal(6, len(testEnvService.GetEnvs()), "test GetEnvs() failed")
}

func TestEnvService_GetAll(t *testing.T) {
	asst := assert.New(t)

	err := testEnvService.GetAll()
	asst.Nil(err, "test GetAll() failed")
	asst.Equal(6, len(testEnvService.GetEnvs()), "test GetAll() failed")
}

func TestEnvService_GetByID(t *testing.T) {
	asst := assert.New(t)

	err := testEnvService.GetByID(testEnvEnvID)
	asst.Nil(err, "test GetByID() failed")
	asst.Equal(testEnvEnvName, testEnvService.GetEnvs()[constant.ZeroInt].GetEnvName(), "test GetByID() failed")
}

func TestEnvService_GetEnvByName(t *testing.T) {
	asst := assert.New(t)

	err := testEnvService.GetEnvByName(testEnvEnvName)
	asst.Nil(err, "test GetEnvByName() failed")
	asst.Equal(testEnvEnvID, testEnvService.Envs[constant.ZeroInt].Identity(), "test GetEnvByName() failed")
}

func TestEnvService_Create(t *testing.T) {
	asst := assert.New(t)

	err := testEnvService.Create(map[string]interface{}{envEnvNameStruct: testEnvNewEnvName})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	asst.Equal(testEnvNewEnvName, testEnvService.GetEnvs()[constant.ZeroInt].GetEnvName(), "test Create() failed")
	// delete
	err = testEnvService.Delete(testEnvService.GetEnvs()[constant.ZeroInt].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestEnvService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateEnv()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testEnvService.Update(entity.Identity(), map[string]interface{}{envEnvNameStruct: testEnvUpdateEnvName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testEnvService.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testEnvUpdateEnvName, testEnvService.GetEnvs()[constant.ZeroInt].GetEnvName(), "test Update() failed")
	// delete
	err = testEnvService.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestEnvService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateEnv()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testEnvService.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestEnvService_Marshal(t *testing.T) {
	asst := assert.New(t)

	err := testEnvService.GetByID(testEnvEnvID)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	jsonBytes, err := testEnvService.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}

func TestEnvService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	err := testEnvService.GetByID(testEnvEnvID)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	jsonBytes, err := testEnvService.MarshalWithFields(envEnvNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	t.Log(string(jsonBytes))
}
