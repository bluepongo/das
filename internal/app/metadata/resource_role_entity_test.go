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
	testResourceRoleID                   = 1
	testResourceRoleRoleName             = "role"
	testResourceRoleRoleUUID             = "role"
	testResourceRoleResourceGroupID      = 1
	testResourceRoleDelFlag              = 0
	testResourceRoleCreateTimeString     = "2021-01-21 10:00:00.000000"
	testResourceRoleLastUpdateTimeString = "2021-01-21 13:00:00.000000"

	testResourceRoleNewUserID = 100
)

var testResourceRoleInfo *ResourceRoleInfo

func init() {
	testInitDASMySQLPool()
	initDBApplicationUserAndPass()
	testResourceRoleInfo = testInitNewResourceRoleInfo()
}

func testInitNewResourceRoleInfo() *ResourceRoleInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(testResourceRoleCreateTimeString)
	lastUpdateTime, _ := now.Parse(testResourceRoleLastUpdateTimeString)
	ResourceRoleInfo := NewResourceRoleInfoWithGlobal(
		testResourceRoleID,
		testResourceRoleRoleUUID,
		testResourceRoleRoleName,
		testResourceRoleResourceGroupID,
		testResourceRoleDelFlag,
		createTime,
		lastUpdateTime,
	)

	return ResourceRoleInfo
}

func TestResourceRoleEntityAll(t *testing.T) {
	TestResourceRoleInfo_Identity(t)
	TestResourceRoleInfo_Get(t)
	TestResourceRoleInfo_GetResourceGroup(t)
	TestResourceRoleInfo_GetUsers(t)
	TestResourceRoleInfo_Set(t)
	TestResourceRoleInfo_Delete(t)
	TestResourceRoleInfo_AddUser(t)
	TestResourceRoleInfo_DeleteUser(t)
	TestResourceRoleInfo_MarshalJSON(t)
	TestResourceRoleInfo_MarshalJSONWithFields(t)
}

func TestResourceRoleInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	resourceRoleInfo := testInitNewResourceRoleInfo()
	asst.Equal(testResourceRoleID, resourceRoleInfo.Identity(), "test Identity() failed")
}

func TestResourceRoleInfo_Get(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResourceRoleRoleName, testResourceRoleInfo.GetRoleName(), "test GetRoleName() failed")
	asst.Equal(testResourceRoleRoleUUID, testResourceRoleInfo.GetRoleUUID(), "test GetRoleUUID() failed")
	asst.Equal(testResourceRoleResourceGroupID, testResourceRoleInfo.GetResourceGroupID(), "test GetResourceGroupID() failed")
	asst.Equal(testResourceRoleDelFlag, testResourceRoleInfo.GetDelFlag(), "test DelFlag() failed")
	asst.True(reflect.DeepEqual(testResourceRoleInfo.CreateTime, testResourceRoleInfo.GetCreateTime()), "test GetCreateTime() failed")
	asst.True(reflect.DeepEqual(testResourceRoleInfo.LastUpdateTime, testResourceRoleInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestResourceRoleInfo_GetResourceGroup(t *testing.T) {
	asst := assert.New(t)

	resourceGroup, err := testResourceRoleInfo.GetResourceGroup()
	asst.Nil(err, common.CombineMessageWithError("test GetResourceGroup() failed", err))
	asst.Equal(testResourceRoleResourceGroupID, resourceGroup.Identity(), "test GetResourceGroup() failed", err)
}

func TestResourceRoleInfo_GetUsers(t *testing.T) {
	asst := assert.New(t)

	users, err := testResourceRoleInfo.GetUsers()
	asst.Nil(err, common.CombineMessageWithError("test GetUsers() failed", err))
	asst.Equal(1, users[constant.ZeroInt].Identity(), "test GetUsers() failed", err)
}

func TestResourceRoleInfo_Set(t *testing.T) {
	asst := assert.New(t)

	err := testResourceRoleInfo.Set(map[string]interface{}{resourceRoleRoleNameStruct: testResourceRoleUpdateRoleUUID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testResourceRoleUpdateRoleUUID, testResourceRoleInfo.GetRoleName(), "test Set() failed")
	err = testResourceRoleInfo.Set(map[string]interface{}{resourceRoleRoleNameStruct: testResourceRoleRoleName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testResourceRoleRoleName, testResourceRoleInfo.GetRoleName(), "test Set() failed")
}

func TestResourceRoleInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	testResourceRoleInfo.Delete()
	asst.Equal(1, testResourceRoleInfo.GetDelFlag(), "test Delete() failed")
	err := testResourceRoleInfo.Set(map[string]interface{}{resourceRoleDelFlagStruct: constant.ZeroInt})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(constant.ZeroInt, testResourceRoleInfo.GetDelFlag(), "test Set() failed")
}
func TestResourceRoleInfo_AddUser(t *testing.T) {
	asst := assert.New(t)

	err := testResourceRoleInfo.AddUser(testResourceRoleNewUserID)
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
	users, err := testResourceRoleInfo.GetUsers()
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
	asst.Equal(2, len(users), common.CombineMessageWithError("test AddUser() failed", err))
	err = testResourceRoleInfo.DeleteUser(testResourceRoleNewUserID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteUser() failed", err))
}

func TestResourceRoleInfo_DeleteUser(t *testing.T) {
	asst := assert.New(t)

	err := testResourceRoleInfo.AddUser(testResourceRoleNewUserID)
	asst.Nil(err, common.CombineMessageWithError("test AddUser() failed", err))
	err = testResourceRoleInfo.DeleteUser(testResourceRoleNewUserID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteUser() failed", err))
	users, err := testResourceRoleInfo.GetUsers()
	asst.Nil(err, common.CombineMessageWithError("test DeleteUser() failed", err))
	for _, user := range users {
		asst.NotEqual(testResourceRoleNewUserID, user.Identity(), common.CombineMessageWithError("test DeleteUser() failed", err))
	}
}
func TestResourceRoleInfo_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testResourceRoleInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	t.Log(string(jsonBytes))
}

func TestResourceRoleInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testResourceRoleInfo.MarshalJSONWithFields(resourceRoleRoleNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	t.Log(string(jsonBytes))
}
