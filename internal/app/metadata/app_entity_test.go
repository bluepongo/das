package metadata

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/jinzhu/now"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	testAppAppID                = 1
	testAppAppName              = "app1"
	testAppLevel                = 1
	testAppOwnerID              = 1
	testAppDelFlag              = 0
	testAppCreateTimeString     = "2021-01-21 10:00:00.000000"
	testAppLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	testAppAppNameJSON          = "app_name"

	testAppNewDBID   = 2
	testAppNewUserID = 16
)

var testAppInfo *AppInfo

func init() {
	testInitDASMySQLPool()
	testAppInfo = testInitNewAppInfo()
}

func testInitNewAppInfo() *AppInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(testAppCreateTimeString)
	lastUpdateTime, _ := now.Parse(testAppLastUpdateTimeString)
	return NewAppInfoWithGlobal(
		testAppAppID,
		testAppAppName,
		testAppLevel,
		testAppOwnerID,
		testAppDelFlag,
		createTime,
		lastUpdateTime,
	)
}

func appSystemStructEqual(a, b *AppInfo) bool {
	return a.ID == b.ID && a.AppName == b.AppName && a.DelFlag == b.DelFlag && a.CreateTime == b.CreateTime && a.LastUpdateTime == b.LastUpdateTime && a.Level == b.Level && a.OwnerID == b.OwnerID
}

func TestAppEntityAll(t *testing.T) {
	TestAppInfo_Identity(t)
	TestAppInfo_GetAppName(t)
	TestAppInfo_GetLevel(t)
	TestAppInfo_GetOwnerID(t)
	TestAppInfo_GetDelFlag(t)
	TestAppInfo_GetCreateTime(t)
	TestAppInfo_GetLastUpdateTime(t)
	TestAppInfo_Set(t)
	TestAppInfo_Delete(t)
	TestAppInfo_MarshalJSON(t)
	TestAppInfo_MarshalJSONWithFields(t)
	TestAppInfo_AddAppDB(t)
	TestAppInfo_DeleteAppDB(t)
	TestAppInfo_AddAppUser(t)
	TestAppInfo_DeleteAppUser(t)
	TestAppInfo_GetDBS(t)
	TestAppInfo_GetUsers(t)

}

func TestAppInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testAppAppID, testAppInfo.Identity(), "test Identity() failed")
}

func TestAppInfo_GetAppName(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testAppAppName, testAppInfo.GetAppName(), "test GetAppName() failed")
}

func TestAppInfo_GetLevel(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testAppLevel, testAppInfo.GetLevel(), "test GetLevel() failed")
}

func TestAppInfo_GetOwnerID(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testAppOwnerID, testAppInfo.GetOwnerID(), "test GetLevel() failed")
}

func TestAppInfo_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(constant.ZeroInt, testAppInfo.GetDelFlag(), "test GetDelFlag() failed")
}

func TestAppInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	asst.True(reflect.DeepEqual(testAppInfo.CreateTime, testAppInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestAppInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	asst.True(reflect.DeepEqual(testAppInfo.LastUpdateTime, testAppInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestAppInfo_GetDBS(t *testing.T) {
	asst := assert.New(t)

	dbs, err := testAppInfo.GetDBs()
	fmt.Println(dbs)
	asst.Equal(nil, err, "test GetDBs() failed")

}

func TestAppInfo_GetUsers(t *testing.T) {
	asst := assert.New(t)

	users, err := testAppInfo.GetUsers()
	fmt.Println(users)
	asst.Equal(nil, err, "test GetUsers() failed")
	asst.Equal(2, len(users), "test GetUsers() failed")
}

func TestAppInfo_Set(t *testing.T) {
	asst := assert.New(t)

	err := testAppInfo.Set(map[string]interface{}{"AppName": testAppNewAppName})
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(testAppNewAppName, testAppInfo.AppName, "test Set() failed")
	err = testAppInfo.Set(map[string]interface{}{"AppName": testAppAppName})
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
}

func TestAppInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	testAppInfo.Delete()
	asst.Equal(1, testAppInfo.GetDelFlag(), "test Delete() failed")
}

func TestAppInfo_MarshalJSON(t *testing.T) {
	var appSystemInfoUnmarshal *AppInfo

	asst := assert.New(t)

	data, err := testAppInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &appSystemInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(appSystemStructEqual(testAppInfo, appSystemInfoUnmarshal), common.CombineMessageWithError("test MarshalJSON() failed", err))
}

func TestAppInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	data, err := testAppInfo.MarshalJSONWithFields(appAppNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{testAppAppNameJSON: testAppAppName})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
}

func TestAppInfo_AddAppDB(t *testing.T) {
	asst := assert.New(t)

	err := testAppInfo.AddDB(testAppNewDBID)
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
	dbs, err := testAppInfo.GetDBs()
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
	asst.Equal(1, len(dbs), common.CombineMessageWithError("test AddDB() failed", err))
	// delete
	err = testAppInfo.DeleteDB(testAppNewDBID)
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
}

func TestAppInfo_DeleteAppDB(t *testing.T) {
	asst := assert.New(t)

	err := testAppInfo.AddDB(testAppNewDBID)
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
	err = testAppInfo.DeleteDB(testAppNewDBID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
	dbs, err := testAppInfo.GetDBs()
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
	asst.Equal(testAppDBID, dbs[constant.ZeroInt].Identity(), common.CombineMessageWithError("test DeleteDB() failed", err))
}

func TestAppInfo_AddAppUser(t *testing.T) {
	asst := assert.New(t)

	err := testAppInfo.AddUser(testAppNewUserID)
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
	users, err := testAppInfo.GetUsers()
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
	asst.Equal(3, len(users), common.CombineMessageWithError("test AddUser() failed", err))
	// delete
	err = testAppInfo.DeleteUser(testAppNewUserID)
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
}

func TestAppInfo_DeleteAppUser(t *testing.T) {
	asst := assert.New(t)

	err := testAppInfo.AddUser(testAppNewUserID)
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
	err = testAppInfo.DeleteUser(testAppNewUserID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteUser() failed", err))
	users, err := testAppInfo.GetUsers()
	asst.Nil(err, common.CombineMessageWithError("test DeleteUser() failed", err))
	asst.Equal(testAppUserID, users[constant.ZeroInt].Identity(), common.CombineMessageWithError("test DeleteUser() failed", err))
}
