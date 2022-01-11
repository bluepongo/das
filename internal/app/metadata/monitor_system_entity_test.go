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
	testMonitorSystemID                   = 1
	testMonitorSystemSystemName           = "pmm2"
	testMonitorSystemSystemType           = 2
	testMonitorSystemHostIP               = "192.168.10.219"
	testMonitorSystemPortNum              = 80
	testMonitorSystemPortNumSlow          = 9000
	testMonitorSystemBaseUrl              = "/prometheus"
	testMonitorSystemEnvID                = 1
	testMonitorSystemDelFlag              = 0
	testMonitorSystemCreateTimeString     = "2021-01-21 10:00:00.000000"
	testMonitorSystemLastUpdateTimeString = "2021-01-21 13:00:00.000000"
)

var testMonitorSystemInfo *MonitorSystemInfo

func init() {
	testInitDASMySQLPool()
	testMonitorSystemInfo = testInitNewMonitorSystemInfo()
}

func testInitNewMonitorSystemInfo() *MonitorSystemInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(testMonitorSystemCreateTimeString)
	lastUpdateTime, _ := now.Parse(testMonitorSystemLastUpdateTimeString)
	return NewMonitorSystemInfoWithGlobal(
		testMonitorSystemID,
		testMonitorSystemSystemName,
		testMonitorSystemSystemType,
		testMonitorSystemHostIP,
		testMonitorSystemPortNum,
		testMonitorSystemPortNumSlow,
		testMonitorSystemBaseUrl,
		testMonitorSystemEnvID,
		testMonitorSystemDelFlag,
		createTime,
		lastUpdateTime,
	)
}

func TestMonitorSystemEntityAll(t *testing.T) {
	TestMonitorSystemInfo_Identity(t)
	TestMonitorSystemInfo_GetSystemName(t)
	TestMonitorSystemInfo_GetSystemType(t)
	TestMonitorSystemInfo_GetHostIP(t)
	TestMonitorSystemInfo_GetPortNum(t)
	TestMonitorSystemInfo_GetPortNumSlow(t)
	TestMonitorSystemInfo_GetBaseURL(t)
	TestMonitorSystemInfo_GetEnvID(t)
	TestMonitorSystemInfo_GetDelFlag(t)
	TestMonitorSystemInfo_GetCreateTime(t)
	TestMonitorSystemInfo_GetLastUpdateTime(t)
	TestMonitorSystemInfo_Set(t)
	TestMonitorSystemInfo_Delete(t)
	TestMonitorSystemInfo_MarshalJSON(t)
	TestMonitorSystemInfo_MarshalJSONWithFields(t)
}

func TestMonitorSystemInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := testInitNewMonitorSystemInfo()
	asst.Equal(testMonitorSystemID, monitorSystemInfo.Identity(), "test Identity() failed")
}

func TestMonitorSystemInfo_GetSystemName(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := testInitNewMonitorSystemInfo()
	asst.Equal(testMonitorSystemSystemName, monitorSystemInfo.GetSystemName(), "test GetSystemName() failed")
}

func TestMonitorSystemInfo_GetSystemType(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := testInitNewMonitorSystemInfo()
	asst.Equal(testMonitorSystemSystemType, monitorSystemInfo.GetSystemType(), "test GetSystemType() failed")
}

func TestMonitorSystemInfo_GetHostIP(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := testInitNewMonitorSystemInfo()
	asst.Equal(testMonitorSystemHostIP, monitorSystemInfo.GetHostIP(), "test GetHostIP() failed")
}

func TestMonitorSystemInfo_GetPortNum(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := testInitNewMonitorSystemInfo()
	asst.Equal(testMonitorSystemPortNum, monitorSystemInfo.GetPortNum(), "test GetPortNum() failed")
}

func TestMonitorSystemInfo_GetPortNumSlow(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := testInitNewMonitorSystemInfo()
	asst.Equal(testMonitorSystemPortNumSlow, monitorSystemInfo.GetPortNumSlow(), "test GetPortNumSlow() failed")
}

func TestMonitorSystemInfo_GetBaseURL(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := testInitNewMonitorSystemInfo()
	asst.Equal(testMonitorSystemBaseUrl, monitorSystemInfo.GetBaseURL(), "test GetBaseURL() failed")
}

func TestMonitorSystemInfo_GetEnvID(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := testInitNewMonitorSystemInfo()
	asst.Equal(testMonitorSystemEnvID, monitorSystemInfo.GetEnvID(), "test GetEnvID() failed")
}

func TestMonitorSystemInfo_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := testInitNewMonitorSystemInfo()
	asst.Equal(testMonitorSystemDelFlag, monitorSystemInfo.GetDelFlag(), "test GetDelFlag() failed")
}

func TestMonitorSystemInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := testInitNewMonitorSystemInfo()
	asst.True(reflect.DeepEqual(monitorSystemInfo.CreateTime, monitorSystemInfo.GetCreateTime()), "test GetCreateTime() failed")
}

func TestMonitorSystemInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo := testInitNewMonitorSystemInfo()
	asst.True(reflect.DeepEqual(monitorSystemInfo.LastUpdateTime, monitorSystemInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestMonitorSystemInfo_Set(t *testing.T) {
	asst := assert.New(t)

	err := testMonitorSystemInfo.Set(map[string]interface{}{monitorSystemSystemNameStruct: testMonitorSystemUpdateSystemName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testMonitorSystemUpdateSystemName, testMonitorSystemInfo.GetSystemName(), "test Set() failed")
	err = testMonitorSystemInfo.Set(map[string]interface{}{monitorSystemSystemNameStruct: testMonitorSystemSystemName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testMonitorSystemSystemName, testMonitorSystemInfo.GetSystemName(), "test Set() failed")
}

func TestMonitorSystemInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	testMonitorSystemInfo.Delete()
	asst.Equal(1, testMonitorSystemInfo.GetDelFlag(), "test Delete() failed")
	err := testMonitorSystemInfo.Set(map[string]interface{}{monitorSystemDelFlagStruct: constant.ZeroInt})
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	asst.Equal(constant.ZeroInt, testMonitorSystemInfo.GetDelFlag(), "test Delete() failed")
}

func TestMonitorSystemInfo_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testMonitorSystemInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	t.Log(string(jsonBytes))
}

func TestMonitorSystemInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testMonitorSystemInfo.MarshalJSONWithFields(monitorSystemSystemNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	t.Log(string(jsonBytes))
}
