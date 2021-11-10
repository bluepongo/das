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
)

func initNewAppInfo() *AppInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(testAppCreateTimeString)
	lastUpdateTime, _ := now.Parse(testAppLastUpdateTimeString)
	return NewAppInfo(
		testAppRepo,
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
	TestAppInfo_GetDBS(t)
}

func TestAppInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	asst.Equal(testAppAppID, appSystemInfo.Identity(), "test Identity() failed")
}

func TestAppInfo_GetAppName(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	asst.Equal(testAppAppName, appSystemInfo.GetAppName(), "test GetAppName() failed")
}

func TestAppInfo_GetLevel(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	asst.Equal(testAppLevel, appSystemInfo.GetLevel(), "test GetLevel() failed")
}

func TestAppInfo_GetOwnerID(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	asst.Equal(testAppOwnerID, appSystemInfo.GetOwnerID(), "test GetLevel() failed")
}

func TestAppInfo_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	asst.Equal(constant.ZeroInt, appSystemInfo.GetDelFlag(), "test GetDelFlag() failed")
}

func TestAppInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	asst.True(reflect.DeepEqual(appSystemInfo.CreateTime, appSystemInfo.GetCreateTime()), "test GetCreateTime failed")
}

func TestAppInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	asst.True(reflect.DeepEqual(appSystemInfo.LastUpdateTime, appSystemInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestAppInfo_GetDBS(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	dbs, err := appSystemInfo.GetDBs()
	fmt.Println(dbs)
	asst.Equal(nil, err, "test GetDBs() failed")

}

func TestAppInfo_Set(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	newAppName := "new_appSystem"
	err := appSystemInfo.Set(map[string]interface{}{"AppName": newAppName})
	asst.Nil(err, common.CombineMessageWithError("test Get() failed", err))
	asst.Equal(newAppName, appSystemInfo.AppName, "test Set() failed")
}

func TestAppInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	appSystemInfo.Delete()
	asst.Equal(1, appSystemInfo.GetDelFlag(), "test Delete() failed")
}

func TestAppInfo_MarshalJSON(t *testing.T) {
	var appSystemInfoUnmarshal *AppInfo

	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	data, err := appSystemInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &appSystemInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(appSystemStructEqual(appSystemInfo, appSystemInfoUnmarshal), common.CombineMessageWithError("test MarshalJSON() failed", err))
}

func TestAppInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	data, err := appSystemInfo.MarshalJSONWithFields(appAppNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{testAppAppNameJSON: testAppAppName})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
}

func TestAppInfo_AddAppDB(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	err := appSystemInfo.AddDB(2)
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
	dbs, err := appSystemInfo.GetDBs()
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
	asst.Equal(2, len(dbs), common.CombineMessageWithError("test AddDB() failed", err))
	// delete
	err = appSystemInfo.DeleteDB(2)
	asst.Nil(err, common.CombineMessageWithError("test AddDB() failed", err))
}

func TestAppInfo_DeleteAppDB(t *testing.T) {
	asst := assert.New(t)

	appSystemInfo := initNewAppInfo()
	err := appSystemInfo.DeleteDB(1)
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))

	dbs, err := appSystemInfo.GetDBs()
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
	asst.Zero(len(dbs), common.CombineMessageWithError("test DeleteDB() failed", err))
	// add
	err = appSystemInfo.AddDB(1)
	asst.Nil(err, common.CombineMessageWithError("test DeleteDB() failed", err))
}

func entityGetDBIDList(appSystemInfo *AppInfo) (dbIDList []int, err error) {
	dbs, err := appSystemInfo.GetDBs()
	if err != nil {
		return nil, err
	}
	for _, db := range dbs {
		dbIDList = append(dbIDList, db.Identity())
	}
	return dbIDList, nil
}
