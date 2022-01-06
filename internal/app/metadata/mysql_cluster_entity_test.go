package metadata

import (
	"reflect"
	"testing"

	"github.com/jinzhu/now"
	"github.com/romberli/das/config"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	testMySQLClusterID                   = 1
	testMySQLClusterClusterName          = "mysql-cluster-pmm2"
	testMySQLClusterMiddlewareClusterID  = 0
	testMySQLClusterMonitorSystemID      = 1
	testMySQLClusterOwnerID              = 1
	testMySQLClusterEnvID                = 1
	testMySQLClusterDelFlag              = 0
	testMySQLClusterCreateTimeString     = "2021-01-21 10:00:00.000000"
	testMySQLClusterLastUpdateTimeString = "2021-01-21 13:00:00.000000"
)

var testMySQLClusterInfo *MySQLClusterInfo

func init() {
	testInitDASMySQLPool()
	initDBApplicationUserAndPass()
	testMySQLClusterInfo = testInitNewMySQLClusterInfo()
}

func initDBApplicationUserAndPass() {
	viper.Set(config.DBApplicationMySQLUserKey, config.DefaultDBApplicationMySQLUser)
	viper.Set(config.DBApplicationMySQLPassKey, config.DefaultDBApplicationMySQLPass)
}

func testInitNewMySQLClusterInfo() *MySQLClusterInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(testMySQLClusterCreateTimeString)
	lastUpdateTime, _ := now.Parse(testMySQLClusterLastUpdateTimeString)
	mysqlClusterInfo := NewMySQLClusterInfoWithGlobal(
		testMySQLClusterID,
		testMySQLClusterClusterName,
		testMySQLClusterMiddlewareClusterID,
		testMySQLClusterMonitorSystemID,
		testMySQLClusterOwnerID,
		testMySQLClusterEnvID,
		testMySQLClusterDelFlag,
		createTime,
		lastUpdateTime,
	)

	return mysqlClusterInfo
}

func TestMySQLClusterEntityAll(t *testing.T) {
	TestMySQLClusterInfo_Identity(t)
	TestMySQLClusterInfo_Get(t)
	TestMySQLClusterInfo_GetMySQLServers(t)
	TestMySQLClusterInfo_GetMasterServer(t)
	TestMySQLClusterInfo_GetDBs(t)
	TestMySQLClusterInfo_GetAppOwners(t)
	TestMySQLClusterInfo_GetDBOwners(t)
	TestMySQLClusterInfo_GetAllOwners(t)
	TestMySQLClusterInfo_Set(t)
	TestMySQLClusterInfo_Delete(t)
	TestMySQLClusterInfo_MarshalJSON(t)
	TestMySQLClusterInfo_MarshalJSONWithFields(t)
}

func TestMySQLClusterInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := testInitNewMySQLClusterInfo()
	asst.Equal(testMySQLClusterID, mysqlClusterInfo.Identity(), "test Identity() failed")
}

func TestMySQLClusterInfo_Get(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testMySQLClusterClusterName, testMySQLClusterInfo.GetClusterName(), "test GetClusterName() failed")
	asst.Equal(testMySQLClusterMiddlewareClusterID, testMySQLClusterInfo.GetMiddlewareClusterID(), "test GetMiddlewareClusterID() failed")
	asst.Equal(testMySQLClusterMonitorSystemID, testMySQLClusterInfo.GetMonitorSystemID(), "test GetMonitorSystemID() failed")
	asst.Equal(testMySQLClusterOwnerID, testMySQLClusterInfo.GetOwnerID(), "test GetOwnerID() failed")
	asst.Equal(testMySQLClusterEnvID, testMySQLClusterInfo.GetEnvID(), "test GetEnvID() failed")
	asst.Equal(testMySQLClusterDelFlag, testMySQLClusterInfo.GetDelFlag(), "test DelFlag() failed")
	asst.True(reflect.DeepEqual(testMySQLClusterInfo.CreateTime, testMySQLClusterInfo.GetCreateTime()), "test GetCreateTime() failed")
	asst.True(reflect.DeepEqual(testMySQLClusterInfo.LastUpdateTime, testMySQLClusterInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestMySQLClusterInfo_GetMySQLServers(t *testing.T) {
	asst := assert.New(t)

	mysqlServers, err := testMySQLClusterInfo.GetMySQLServers()
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLServers() failed", err))
	asst.Equal(1, mysqlServers[constant.ZeroInt].Identity(), "test GetMySQLServers() failed")
}

func TestMySQLClusterInfo_GetMasterServer(t *testing.T) {
	asst := assert.New(t)

	masterServers, err := testMySQLClusterInfo.GetMasterServers()
	asst.Nil(err, common.CombineMessageWithError("test GetMasterServers() failed", err))
	asst.Equal(1, masterServers[constant.ZeroInt].Identity(), "test GetMasterServers() failed", err)
}

func TestMySQLClusterInfo_GetDBs(t *testing.T) {
	asst := assert.New(t)

	dbs, err := testMySQLClusterInfo.GetDBs()
	asst.Nil(err, common.CombineMessageWithError("test GetDBs() failed", err))
	asst.Equal(2, len(dbs), "test GetDBs() failed", err)
}

func TestMySQLClusterInfo_GetAppOwners(t *testing.T) {
	asst := assert.New(t)

	owners, err := testMySQLClusterInfo.GetAppOwners()
	// fmt.Print(owners[0].Identity(), owners[1].Identity(), owners[2].Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetAppOwners() failed", err))
	asst.Equal(14, owners[constant.ZeroInt].Identity(), "test GetAppOwners() failed", err)
	asst.Equal(2, len(owners), "test GetAppOwners() failed", err)
}

func TestMySQLClusterInfo_GetDBOwners(t *testing.T) {
	asst := assert.New(t)

	owners, err := testMySQLClusterInfo.GetDBOwners()
	asst.Nil(err, common.CombineMessageWithError("test GetDBOwners() failed", err))
	asst.Equal(1, owners[constant.ZeroInt].Identity(), "test GetDBOwners() failed", err)
}

func TestMySQLClusterInfo_GetAllOwners(t *testing.T) {
	asst := assert.New(t)

	owners, err := testMySQLClusterInfo.GetAllOwners()
	// fmt.Print(owners[0].Identity(), owners[1].Identity(), owners[2].Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetAllOwners() failed", err))
	asst.Equal(3, len(owners), "test GetAllOwners() failed", err)
}

func TestMySQLClusterInfo_Set(t *testing.T) {
	asst := assert.New(t)

	err := testMySQLClusterInfo.Set(map[string]interface{}{mysqlClusterClusterNameStruct: testMySQLClusterUpdateClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testMySQLClusterUpdateClusterName, testMySQLClusterInfo.GetClusterName(), "test Set() failed")
	err = testMySQLClusterInfo.Set(map[string]interface{}{mysqlClusterClusterNameStruct: testMySQLClusterClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testMySQLClusterClusterName, testMySQLClusterInfo.GetClusterName(), "test Set() failed")
}

func TestMySQLClusterInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	testMySQLClusterInfo.Delete()
	asst.Equal(1, testMySQLClusterInfo.GetDelFlag(), "test Delete() failed")
	err := testMySQLClusterInfo.Set(map[string]interface{}{mysqlClusterDelFlagStruct: constant.ZeroInt})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(constant.ZeroInt, testMySQLClusterInfo.GetDelFlag(), "test Set() failed")
}

func TestMySQLClusterInfo_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testMySQLClusterInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	t.Log(string(jsonBytes))
}

func TestMySQLClusterInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testMySQLClusterInfo.MarshalJSONWithFields(mysqlClusterClusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	t.Log(string(jsonBytes))
}
