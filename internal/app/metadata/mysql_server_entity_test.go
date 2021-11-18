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
	testMySQLServerID                   = 1
	testMySQLServerClusterID            = 1
	testMySQLServerServerName           = "192-168-10-219"
	testMySQLServerServiceName          = "192-168-10-219:3306"
	testMySQLServerHostIP               = "192.168.10.219"
	testMySQLServerPortNum              = 3306
	testMySQLServerDeploymentType       = 1
	testMySQLServerVersion              = "5.7"
	testMySQLServerDelFlag              = 0
	testMySQLServerCreateTimeString     = "2021-01-21 10:00:00.000000"
	testMySQLServerLastUpdateTimeString = "2021-01-21 13:00:00.000000"
)

var testMySQLServerInfo *MySQLServerInfo

func init() {
	initDASMySQLPool()
	testMySQLServerInfo = testInitNewMySQLServerInfo()
}

func testInitNewMySQLServerInfo() *MySQLServerInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(testMySQLServerCreateTimeString)
	lastUpdateTime, _ := now.Parse(testMySQLServerLastUpdateTimeString)
	return NewMySQLServerInfoWithGlobal(
		testMySQLServerID,
		testMySQLServerClusterID,
		testMySQLServerServerName,
		testMySQLServerServiceName,
		testMySQLServerHostIP,
		testMySQLServerPortNum,
		testMySQLServerDeploymentType,
		testMySQLServerVersion,
		testMySQLServerDelFlag,
		createTime,
		lastUpdateTime)
}

func TestMySQLServerEntityAll(t *testing.T) {
	TestMySQLServerInfo_Identity(t)
	TestMySQLServerInfo_Get(t)
	TestMySQLServerInfo_IsMaster(t)
	TestMySQLServerInfo_GetMySQLCluster(t)
	TestMySQLServerInfo_GetMonitorSystem(t)
	TestMySQLServerInfo_Set(t)
	TestMySQLServerInfo_Delete(t)
	TestMySQLServerInfo_MarshalJSON(t)
	TestMySQLServerInfo_MarshalJSONWithFields(t)
}

func TestMySQLServerInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testMySQLServerID, testMySQLServerInfo.Identity(), "test Identity() failed")
}

func TestMySQLServerInfo_Get(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testMySQLServerClusterID, testMySQLServerInfo.GetClusterID(), "test GetClusterID() failed")
	asst.Equal(testMySQLServerHostIP, testMySQLServerInfo.GetHostIP(), "test GetHostIP() failed")
	asst.Equal(testMySQLServerPortNum, testMySQLServerInfo.GetPortNum(), "test GetPortNum() failed")
	asst.Equal(testMySQLServerDeploymentType, testMySQLServerInfo.GetDeploymentType(), "test GetDeploymentType() failed")
	asst.Equal(testMySQLServerVersion, testMySQLServerInfo.GetVersion(), "test GetVersion() failed")
	asst.Equal(testMySQLServerDelFlag, testMySQLServerInfo.GetDelFlag(), "test GetDelFlag() failed")
	asst.True(reflect.DeepEqual(testMySQLServerInfo.CreateTime, testMySQLServerInfo.GetCreateTime()), "test GetCreateTime() failed")
	asst.True(reflect.DeepEqual(testMySQLServerInfo.LastUpdateTime, testMySQLServerInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestMySQLServerInfo_IsMaster(t *testing.T) {
	asst := assert.New(t)

	isMaster, err := testMySQLServerInfo.IsMaster()
	asst.Nil(err, common.CombineMessageWithError("test IsMaster() failed", err))
	asst.True(isMaster, "test IsMaster() failed")
}

func TestMySQLServerInfo_GetMySQLCluster(t *testing.T) {
	asst := assert.New(t)

	mysqlCluster, err := testMySQLServerInfo.GetMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLCluster() failed", err))
	asst.Equal(testMySQLServerClusterID, mysqlCluster.Identity(), "test GetMySQLCluster() failed")
}

func TestMySQLServerInfo_GetMonitorSystem(t *testing.T) {
	asst := assert.New(t)

	monitorSystemInfo, err := testMySQLServerInfo.GetMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test GetMonitorSystem() failed", err))
	asst.Equal(1, monitorSystemInfo.Identity(), "test GetMonitorSystem() failed")
}

func TestMySQLServerInfo_Set(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLServerInfo.Set(map[string]interface{}{mysqlServerServerNameStruct: testMySQLServerUpdateServerName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testMySQLServerUpdateServerName, testMySQLServerInfo.GetServerName(), "test Set() failed")
	err = testMySQLServerInfo.Set(map[string]interface{}{mysqlServerServerNameStruct: testMySQLServerServerName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testMySQLServerServerName, testMySQLServerInfo.GetServerName(), "test Set() failed")
}

func TestMySQLServerInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	testMySQLServerInfo.Delete()
	asst.Equal(1, testMySQLServerInfo.GetDelFlag(), "test Delete() failed")
	err := testMySQLServerInfo.Set(map[string]interface{}{mysqlServerDelFlagStruct: constant.ZeroInt})
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	asst.Equal(constant.ZeroInt, testMySQLServerInfo.GetDelFlag(), "test Delete() failed")
}

func TestMySQLServerInfo_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testMySQLServerInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}

func TestMySQLServerInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testMySQLServerInfo.MarshalJSONWithFields(mysqlServerServerNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	t.Log(string(jsonBytes))
}
