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
	testMiddlewareServerID                   = 1
	testMiddlewareServerClusterID            = 1
	testMiddlewareServerServerName           = "middleware-server-1"
	testMiddlewareServerMiddlewareRole       = 1
	testMiddlewareServerHostIP               = "192.168.10.219"
	testMiddlewareServerPortNum              = 33061
	testMiddlewareServerDelFlag              = 0
	testMiddlewareServerCreateTimeString     = "2021-01-21 10:00:00.000000"
	testMiddlewareServerLastUpdateTimeString = "2021-01-21 13:00:00.000000"
)

var testMiddlewareServerInfo *MiddlewareServerInfo

func init() {
	initDASMySQLPool()
	testMiddlewareServerInfo = initNewMiddlewareServerInfo()
}

func initNewMiddlewareServerInfo() *MiddlewareServerInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(testMiddlewareServerCreateTimeString)
	lastUpdateTime, _ := now.Parse(testMiddlewareServerLastUpdateTimeString)
	return NewMiddlewareServerInfo(
		testMiddlewareServerRepo,
		testMiddlewareServerID,
		testMiddlewareServerClusterID,
		testMiddlewareServerServerName,
		testMiddlewareServerMiddlewareRole,
		testMiddlewareServerHostIP,
		testMiddlewareServerPortNum,
		testMiddlewareServerDelFlag,
		createTime,
		lastUpdateTime,
	)
}

func TestMiddlewareServerEntityAll(t *testing.T) {
	TestMiddlewareServerInfo_Identity(t)
	TestMiddlewareServerInfo_GetClusterID(t)
	TestMiddlewareServerInfo_GetServerName(t)
	TestMiddlewareServerInfo_GetMiddlewareRole(t)
	TestMiddlewareServerInfo_GetHostIP(t)
	TestMiddlewareServerInfo_GetPortNum(t)
	TestMiddlewareServerInfo_GetDelFlag(t)
	TestMiddlewareServerInfo_GetCreateTime(t)
	TestMiddlewareServerInfo_GetLastUpdateTime(t)
	TestMiddlewareServerInfo_Set(t)
	TestMiddlewareServerInfo_Delete(t)
	TestMiddlewareServerInfo_MarshalJSON(t)
	TestMiddlewareServerInfo_MarshalJSONWithFields(t)
}

func TestMiddlewareServerInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testMiddlewareServerID, testMiddlewareServerInfo.Identity(), "test Identity() failed")
}

func TestMiddlewareServerInfo_GetClusterID(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testMiddlewareServerClusterID, testMiddlewareServerInfo.GetClusterID(), "test GetClusterID() failed")
}

func TestMiddlewareServerInfo_GetServerName(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testMiddlewareServerServerName, testMiddlewareServerInfo.GetServerName(), "test GetServerName() failed")
}

func TestMiddlewareServerInfo_GetMiddlewareRole(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testMiddlewareServerMiddlewareRole, testMiddlewareServerInfo.GetMiddlewareRole(), "test GetServerName() failed")
}

func TestMiddlewareServerInfo_GetHostIP(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testMiddlewareServerHostIP, testMiddlewareServerInfo.GetHostIP(), "test GetServerName() failed")
}

func TestMiddlewareServerInfo_GetPortNum(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testMiddlewareServerPortNum, testMiddlewareServerInfo.GetPortNum(), "test GetServerName() failed")
}

func TestMiddlewareServerInfo_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testMiddlewareServerDelFlag, testMiddlewareServerInfo.GetDelFlag(), "test GetServerName() failed")
}

func TestMiddlewareServerInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	asst.True(reflect.DeepEqual(testMiddlewareServerInfo.CreateTime, testMiddlewareServerInfo.GetCreateTime()), "test GetCreateTime() failed")
}

func TestMiddlewareServerInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	asst.True(reflect.DeepEqual(testMiddlewareServerInfo.LastUpdateTime, testMiddlewareServerInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestMiddlewareServerInfo_Set(t *testing.T) {
	asst := assert.New(t)

	err := testMiddlewareServerInfo.Set(map[string]interface{}{middlewareServerServerNameStruct: testMiddlewareServerUpdateServerName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testMiddlewareServerUpdateServerName, testMiddlewareServerInfo.GetServerName(), "test Set() failed")
	err = testMiddlewareServerInfo.Set(map[string]interface{}{middlewareServerServerNameStruct: testMiddlewareServerServerName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testMiddlewareServerServerName, testMiddlewareServerInfo.GetServerName(), "test Set() failed")
}

func TestMiddlewareServerInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	testMiddlewareServerInfo.Delete()
	asst.Equal(1, testMiddlewareServerInfo.GetDelFlag(), "test Delete() failed")
	err := testMiddlewareServerInfo.Set(map[string]interface{}{middlewareServerDelFlagStruct: constant.ZeroInt})
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	asst.Equal(constant.ZeroInt, testMiddlewareServerInfo.GetDelFlag(), "test Delete() failed")
}

func TestMiddlewareServerInfo_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testMiddlewareServerInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	t.Log(string(jsonBytes))
}

func TestMiddlewareServerInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testMiddlewareServerInfo.MarshalJSONWithFields(middlewareServerServerNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	t.Log(string(jsonBytes))
}
