package metadata

import (
	"reflect"
	"testing"

	"github.com/jinzhu/now"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	testMiddlewareClusterClusterID            = 1
	testMiddlewareClusterClusterName          = "middleware-cluster-1"
	testMiddlewareClusterOwnerID              = 1
	testMiddlewareClusterEnvID                = 1
	testMiddlewareClusterDelFlag              = 0
	testMiddlewareClusterCreateTimeString     = "2021-01-21 10:00:00.000000"
	testMiddlewareClusterLastUpdateTimeString = "2021-01-21 13:00:00.000000"
)

var testMiddlewareClusterInfo *MiddlewareClusterInfo

func init() {
	testInitDASMySQLPool()
	testMiddlewareClusterInfo = testInitNewMiddlewareClusterInfo()
}

func testInitNewMiddlewareClusterInfo() *MiddlewareClusterInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(testMiddlewareClusterCreateTimeString)
	lastUpdateTime, _ := now.Parse(testMiddlewareClusterLastUpdateTimeString)
	return NewMiddlewareClusterInfoWithGlobal(
		testMiddlewareClusterClusterID,
		testMiddlewareClusterClusterName,
		testMiddlewareClusterOwnerID,
		testMiddlewareClusterEnvID,
		testMiddlewareClusterDelFlag,
		createTime,
		lastUpdateTime,
	)
}

func TestMiddlewareClusterEntityAll(t *testing.T) {
	TestMiddlewareClusterInfo_Identity(t)
	TestMiddlewareClusterInfo_GetClusterName(t)
	TestMiddlewareClusterInfo_GetOwnerID(t)
	TestMiddlewareClusterInfo_GetEnvID(t)
	TestMiddlewareClusterInfo_GetDelFlag(t)
	TestMiddlewareClusterInfo_GetCreateTime(t)
	TestMiddlewareClusterInfo_GetLastUpdateTime(t)
	TestMiddlewareClusterInfo_GetCreateTime(t)
	TestMiddlewareClusterInfo_GetMiddlewareServers(t)
	TestMiddlewareClusterInfo_Set(t)
	TestMiddlewareClusterInfo_Delete(t)
	TestMiddlewareClusterInfo_MarshalJSON(t)
	TestMiddlewareClusterInfo_MarshalJSONWithFields(t)
}

func TestMiddlewareClusterInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testMiddlewareClusterClusterID, testMiddlewareClusterInfo.Identity(), "test Identity() failed")
}

func TestMiddlewareClusterInfo_GetClusterName(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testMiddlewareClusterClusterName, testMiddlewareClusterInfo.GetClusterName(), "test GetClusterName() failed")
}

func TestMiddlewareClusterInfo_GetOwnerID(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testMiddlewareClusterOwnerID, testMiddlewareClusterInfo.GetOwnerID(), "test GetOwnerID() failed")
}

func TestMiddlewareClusterInfo_GetEnvID(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testMiddlewareClusterEnvID, testMiddlewareClusterInfo.GetEnvID(), "test GetEnvID() failed")
}
func TestMiddlewareClusterInfo_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(constant.ZeroInt, testMiddlewareClusterInfo.GetDelFlag(), "test GetDelFlag() failed")
}

func TestMiddlewareClusterInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	asst.True(reflect.DeepEqual(testMiddlewareClusterInfo.CreateTime, testMiddlewareClusterInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestMiddlewareClusterInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	asst.True(reflect.DeepEqual(testMiddlewareClusterInfo.LastUpdateTime, testMiddlewareClusterInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestMiddlewareClusterInfo_GetMiddlewareServers(t *testing.T) {
	asst := assert.New(t)

	middlewareServers, err := testMiddlewareClusterInfo.GetMiddlewareServers()
	asst.Nil(err, common.CombineMessageWithError("test GetMiddlewareServersByID() failed", err))
	asst.Equal(1, len(middlewareServers), "test GetMiddlewareServersByID() failed")
}

func TestMiddlewareClusterInfo_Set(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareClusterInfo.Set(map[string]interface{}{middlewareClusterClusterNameStruct: testMiddlewareClusterUpdateClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testMiddlewareClusterUpdateClusterName, testMiddlewareClusterInfo.GetClusterName(), "test Set() failed")
	err = testMiddlewareClusterInfo.Set(map[string]interface{}{middlewareClusterClusterNameStruct: testMiddlewareClusterNewClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testMiddlewareClusterNewClusterName, testMiddlewareClusterInfo.GetClusterName(), "test Set() failed")
}

func TestMiddlewareClusterInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	testMiddlewareClusterInfo.Delete()
	asst.Equal(1, testMiddlewareClusterInfo.GetDelFlag(), "test Delete() failed")
	err := testMiddlewareClusterInfo.Set(map[string]interface{}{middlewareClusterDelFlagStruct: constant.ZeroInt})
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	asst.Equal(constant.ZeroInt, testMiddlewareClusterInfo.GetDelFlag(), "test Delete() failed")
}

func TestMiddlewareClusterInfo_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testMiddlewareClusterInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	t.Log(string(jsonBytes))
}

func TestMiddlewareClusterInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testMiddlewareClusterInfo.MarshalJSONWithFields(middlewareClusterClusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	t.Log(string(jsonBytes))
}
