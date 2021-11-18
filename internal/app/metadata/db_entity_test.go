package metadata

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/jinzhu/now"
	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	testDBDBID                 = 1
	testDBDBName               = "pmm_test"
	testDBClusterID            = 1
	testDBClusterType          = 1
	testDBOwnerID              = 1
	testDBEnvID                = 1
	testDBDelFlag              = 0
	testDBCreateTimeString     = "2021-01-21 10:00:00.000000"
	testDBLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	testDBDBNameJSON           = "db_name"
	testDBAppID                = 1
	testDBNewAppID             = 2
	testDBMySQLClusterID       = 1
)

var testDBInfo *DBInfo

func init() {
	testDBInfo = testInitNewDBInfo()
}

func testInitNewDBInfo() *DBInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(testDBCreateTimeString)
	lastUpdateTime, _ := now.Parse(testDBLastUpdateTimeString)

	return NewDBInfoWithGlobal(
		testDBDBID,
		testDBDBName,
		testDBClusterID,
		testDBClusterType,
		testDBOwnerID,
		testDBEnvID,
		testDBDelFlag,
		createTime,
		lastUpdateTime,
	)
}

func dbEqual(a, b *DBInfo) bool {
	return a.ID == b.ID && a.DBName == b.DBName && a.ClusterID == b.ClusterID && a.ClusterType == b.ClusterType &&
		a.OwnerID == b.OwnerID && a.EnvID == b.EnvID && a.DelFlag == b.DelFlag && a.CreateTime == b.CreateTime &&
		a.LastUpdateTime == b.LastUpdateTime
}

func TestDBEntityAll(t *testing.T) {
	TestDBInfo_Identity(t)
	TestDBInfo_GetDBName(t)
	TestDBInfo_GetClusterID(t)
	TestDBInfo_GetClusterType(t)
	TestDBInfo_GetOwnerID(t)
	TestDBInfo_GetEnvID(t)
	TestDBInfo_GetDelFlag(t)
	TestDBInfo_GetCreateTime(t)
	TestDBInfo_GetLastUpdateTime(t)
	TestDBInfo_GetApps(t)
	TestDBInfo_GetMySQLClusterByID(t)
	TestDBInfo_GetAppOwners(t)
	TestDBInfo_GetDBOwners(t)
	TestDBInfo_GetAllOwners(t)
	TestDBInfo_Set(t)
	TestDBInfo_Delete(t)
	TestDBInfo_AddApp(t)
	TestDBInfo_DeleteApp(t)
	TestDBInfo_MarshalJSON(t)
	TestDBInfo_MarshalJSONWithFields(t)
}

func TestDBInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testDBDBID, testDBInfo.Identity(), "test Identity() failed")
}

func TestDBInfo_GetDBName(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testDBDBName, testDBInfo.GetDBName(), "test GetDBName() failed")
}

func TestDBInfo_GetClusterID(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testDBClusterID, testDBInfo.GetClusterID(), "test GetClusterID() failed")
}

func TestDBInfo_GetClusterType(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testDBClusterType, testDBInfo.GetClusterType(), "test GetClusterType() failed")
}

func TestDBInfo_GetOwnerID(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testDBOwnerID, testDBInfo.GetOwnerID(), "test GetOwnerID() failed")
}

func TestDBInfo_GetEnvID(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testDBEnvID, testDBInfo.GetEnvID(), "test GetEnvID() failed")
}

func TestDBInfo_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testDBDelFlag, testDBInfo.GetDelFlag(), "test GetDelFlag() failed")
}

func TestDBInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	asst.True(reflect.DeepEqual(testDBInfo.CreateTime, testDBInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestDBInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	asst.True(reflect.DeepEqual(testDBInfo.LastUpdateTime, testDBInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestDBInfo_GetMySQLClusterByID(t *testing.T) {
	asst := assert.New(t)

	mysqlCluster, err := testDBInfo.GetMySQLCluster()
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLClusterByID() failed", err))
	asst.Equal(testDBClusterID, mysqlCluster.Identity(), "test GetMySQLClusterByID() failed")
}

func TestDBInfo_GetAppOwners(t *testing.T) {
	asst := assert.New(t)

	appOwners, err := testDBInfo.GetAppOwners()
	asst.Nil(err, common.CombineMessageWithError("test GetAppOwners() failed", err))
	asst.Equal(testAppOwnerID, appOwners[constant.ZeroInt].Identity(), "test GetAppOwners() failed")
}

func TestDBInfo_GetDBOwners(t *testing.T) {
	asst := assert.New(t)

	dbOwners, err := testDBInfo.GetDBOwners()
	asst.Nil(err, common.CombineMessageWithError("test GetDBOwners() failed", err))
	asst.Equal(testAppOwnerID, dbOwners[constant.ZeroInt].Identity(), "test GetDBOwners() failed")
}

func TestDBInfo_GetAllOwners(t *testing.T) {
	asst := assert.New(t)

	allOwners, err := testDBInfo.GetAllOwners()
	asst.Nil(err, common.CombineMessageWithError("test GetAllOwners() failed", err))
	asst.Equal(testAppOwnerID, allOwners[constant.ZeroInt].Identity(), "test GetAllOwners() failed")
}

func TestDBInfo_GetApps(t *testing.T) {
	asst := assert.New(t)

	apps, err := testDBInfo.GetApps()
	asst.Nil(err, common.CombineMessageWithError("test GetApps() failed", err))
	asst.Equal(testDBAppID, apps[constant.ZeroInt].Identity(), "test GetApps() failed")
}

func TestDBInfo_Set(t *testing.T) {
	asst := assert.New(t)

	err := testDBInfo.Set(map[string]interface{}{dbDBNameStruct: testDBNewDBName})
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(testDBNewDBName, testDBInfo.GetDBName(), "test Set() failed")
	err = testDBInfo.Set(map[string]interface{}{dbDBNameStruct: testDBDBName})
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
}

func TestDBInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	testDBInfo.Delete()
	asst.Equal(1, testDBInfo.DelFlag, "test Delete() failed")
	err := testDBInfo.Set(map[string]interface{}{dbDelFlagStruct: constant.ZeroInt})
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestDBInfo_AddApp(t *testing.T) {
	var apps []metadata.App

	asst := assert.New(t)

	err := testDBInfo.AddApp(testDBNewAppID)
	apps, err = testDBInfo.GetAppsByID(testDBInfo.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	asst.Equal(2, len(apps), "test AddApp() failed")
	// delete
	err = testDBInfo.DeleteApp(testDBNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
}

func TestDBInfo_DeleteApp(t *testing.T) {
	var apps []metadata.App

	asst := assert.New(t)

	err := testDBInfo.AddApp(testDBNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	err = testDBInfo.DeleteApp(testDBNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	apps, err = testDBInfo.GetAppsByID(testDBInfo.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	asst.Equal(1, len(apps), "test AddApp() failed")
}

func TestDBInfo_MarshalJSON(t *testing.T) {
	var dbInfoUnmarshal *DBInfo

	asst := assert.New(t)

	data, err := testDBInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &dbInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(dbEqual(testDBInfo, dbInfoUnmarshal), "test MarshalJSON() failed")
}

func TestDBInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	data, err := testDBInfo.MarshalJSONWithFields(dbDBNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{testDBDBNameJSON: testDBDBName})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}
