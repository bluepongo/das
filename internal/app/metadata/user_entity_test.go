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
	testUserID                   = 1
	testUserUserName             = "zhangsan"
	testUserDepartmentName       = "arch"
	testUserEmployeeID           = "100001"
	testUserAccountName          = "zs001"
	testUserEmail                = "rombertmp@163.com"
	testUserTelephone            = "01012345678"
	testUserMobile               = "13012345678"
	testUserRole                 = 3
	testUserDelFlag              = 0
	testUserCreateTimeString     = "2021-01-21 10:00:00.000000"
	testUserLastUpdateTimeString = "2021-01-21 13:00:00.000000"
)

var testUserInfo *UserInfo

func init() {
	initDASMySQLPool()
	testUserInfo = testInitNewUserInfo()
}

func testInitNewUserInfo() *UserInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(testUserCreateTimeString)
	lastUpdateTime, _ := now.Parse(testUserLastUpdateTimeString)
	return NewUserInfoWithGlobal(
		testUserID,
		testUserUserName,
		testUserDepartmentName,
		testUserEmployeeID,
		testUserAccountName,
		testUserEmail,
		testUserTelephone,
		testUserMobile,
		testUserRole,
		testUserDelFlag,
		createTime,
		lastUpdateTime,
	)
}

func TestUserEntityAll(t *testing.T) {
	TestUserInfo_Identity(t)
	TestUserInfo_GetUserName(t)
	TestUserInfo_GetDepartmentName(t)
	TestUserInfo_GetEmployeeID(t)
	TestUserInfo_GetAccountName(t)
	TestUserInfo_GetEmail(t)
	TestUserInfo_GetTelephone(t)
	TestUserInfo_GetMobile(t)
	TestUserInfo_GetDelFlag(t)
	TestUserInfo_GetCreateTime(t)
	TestUserInfo_GetLastUpdateTime(t)
	TestUserInfo_Set(t)
	TestUserInfo_Delete(t)
	TestUserInfo_MarshalJSON(t)
	TestUserInfo_MarshalJSONWithFields(t)
}

func TestUserInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testUserID, testUserInfo.Identity(), "test Identity() failed")
}

func TestUserInfo_GetUserName(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testUserUserName, testUserInfo.GetUserName(), "test GetUserName() failed")
}

func TestUserInfo_GetDepartmentName(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testUserDepartmentName, testUserInfo.GetDepartmentName(), "test GetDepartmentName() failed")
}

func TestUserInfo_GetEmployeeID(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testUserEmployeeID, testUserInfo.GetEmployeeID(), "test GetEmployeeID() failed")
}

func TestUserInfo_GetAccountName(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testUserAccountName, testUserInfo.GetAccountName(), "test GetAccountName() failed")
}

func TestUserInfo_GetEmail(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testUserEmail, testUserInfo.GetEmail(), "test GetEmail() failed")
}

func TestUserInfo_GetTelephone(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testUserTelephone, testUserInfo.GetTelephone(), "test GetTelephone() failed")
}

func TestUserInfo_GetMobile(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testUserMobile, testUserInfo.GetMobile(), "test GetMobile() failed")
}

func TestUserInfo_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testUserDelFlag, testUserInfo.GetDelFlag(), "test GetDelFlag() failed")
}

func TestUserInfo_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	asst.True(reflect.DeepEqual(testUserInfo.CreateTime, testUserInfo.GetCreateTime()), "test GetCreateTime() failed")
}

func TestUserInfo_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	asst.True(reflect.DeepEqual(testUserInfo.LastUpdateTime, testUserInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestUserInfo_Set(t *testing.T) {
	asst := assert.New(t)

	err := testUserInfo.Set(map[string]interface{}{userUserNameStruct: testUserUpdateUserName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testUserUpdateUserName, testUserInfo.GetUserName(), "test Set() failed")
	err = testUserInfo.Set(map[string]interface{}{userUserNameStruct: testUserUserName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testUserUserName, testUserInfo.GetUserName(), "test Set() failed")
}

func TestUserInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	testUserInfo.Delete()
	asst.Equal(1, testUserInfo.GetDelFlag(), "test Delete() failed")
	err := testUserInfo.Set(map[string]interface{}{userDelFlagStruct: constant.ZeroInt})
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	asst.Equal(constant.ZeroInt, testUserInfo.GetDelFlag(), "test Delete() failed")
}

func TestUserInfo_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testUserInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}

func TestUserInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testUserInfo.MarshalJSONWithFields(userUserNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}
