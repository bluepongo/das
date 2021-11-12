package metadata

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/jinzhu/now"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	testEnvEnvID                = 1
	testEnvEnvName              = "online"
	testEnvDelFlag              = 0
	testEnvCreateTimeString     = "2021-01-21 10:00:00.000000"
	testEnvLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	testEnvEnvNameJSON          = "env_name"
)

var testEnvInfo *EnvInfo

func init() {
	testEnvInfo = initNewEnvInfo()
}

func initNewEnvInfo() *EnvInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)
	createTime, _ := now.Parse(testEnvCreateTimeString)
	lastUpdateTime, _ := now.Parse(testEnvLastUpdateTimeString)
	return NewEnvInfoWithGlobal(testEnvEnvID, testEnvEnvName, testEnvDelFlag, createTime, lastUpdateTime)
}

func equal(a, b *EnvInfo) bool {
	return a.ID == b.ID && a.EnvName == b.EnvName && a.DelFlag == b.DelFlag && a.CreateTime == b.CreateTime && a.LastUpdateTime == b.LastUpdateTime
}

func TestEnvEntityAll(t *testing.T) {
	TestEnvInfo_Identity(t)
	TestEnvInfo_GetEnvName(t)
	TestEnvInfo_GetDelFlag(t)
	TestEnvInfo_GetCreateTime(t)
	TestEnvInfo_GetLastUpdateTime(t)
	TestEnvInfo_Set(t)
	TestEnvInfo_Delete(t)
	TestEnvInfo_MarshalJSON(t)
	TestEnvInfo_MarshalJSONWithFields(t)
}

func TestEnvInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testEnvEnvID, testEnvInfo.Identity(), "test Identity() failed")
}

func TestEnvInfo_GetEnvName(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testEnvEnvName, testEnvInfo.GetEnvName(), "test GetEnvName() failed")
}

func TestEnvInfo_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testEnvDelFlag, testEnvInfo.GetDelFlag(), "test GetDelFlag() failed")
}

func TestEnvInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	asst.True(reflect.DeepEqual(testEnvInfo.CreateTime, testEnvInfo.GetCreateTime()), "test GetCreateTime() failed")
}

func TestEnvInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	asst.True(reflect.DeepEqual(testEnvInfo.LastUpdateTime, testEnvInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestEnvInfo_Set(t *testing.T) {
	asst := assert.New(t)

	err := testEnvInfo.Set(map[string]interface{}{envEnvNameStruct: testEnvUpdateEnvName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testEnvUpdateEnvName, testEnvInfo.GetEnvName(), "test Set() failed")
	err = testEnvInfo.Set(map[string]interface{}{envEnvNameStruct: testEnvEnvName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
}

func TestEnvInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	testEnvInfo.Delete()
	asst.Equal(1, testEnvInfo.GetDelFlag(), "test Delete() failed")
	err := testEnvInfo.Set(map[string]interface{}{envDelFlagStruct: constant.ZeroInt})
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestEnvInfo_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testEnvInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	t.Log(string(jsonBytes))
}

func TestEnvInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testEnvInfo.MarshalJSONWithFields(envEnvNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{testEnvEnvNameJSON: testEnvEnvName})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(jsonBytes), "test MarshalJSONWithFields() failed")
}
