package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

var testUserService *UserService

func init() {
	testInitDASMySQLPool()
	testUserService = NewUserServiceWithDefault()
}

func TestUserServiceAll(t *testing.T) {
	TestUserService_GetUsers(t)
	TestUserService_GetAll(t)
	TestUserService_GetByID(t)
	TestUserService_GetByEmployeeID(t)
	TestUserService_GetByAccountName(t)
	TestUserService_GetByEmail(t)
	TestUserService_GetByTelephone(t)
	TestUserService_GetByMobile(t)
	TestUserService_Create(t)
	TestUserService_Update(t)
	TestUserService_Delete(t)
	TestUserService_Marshal(t)
	TestUserService_MarshalWithFields(t)
}

func TestUserService_GetUsers(t *testing.T) {
	asst := assert.New(t)

	err := testUserService.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetUsers() failed", err))
	asst.Equal(1, len(testUserService.GetUsers()), "test GetUsers() failed")
}

func TestUserService_GetAll(t *testing.T) {
	asst := assert.New(t)

	err := testUserService.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(1, len(testUserService.GetUsers()), "test GetAll() failed")
}

func TestUserService_GetByID(t *testing.T) {
	asst := assert.New(t)

	err := testUserService.GetByID(testUserID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testUserID, testUserService.GetUsers()[constant.ZeroInt].Identity(), "test GetByID() failed")
}

func TestUserService_GetByEmployeeID(t *testing.T) {
	asst := assert.New(t)

	err := testUserService.GetByEmployeeID(testUserEmployeeID)
	asst.Nil(err, common.CombineMessageWithError("test GetByEmployeeID() failed", err))
	asst.Equal(testUserEmployeeID, testUserService.GetUsers()[constant.ZeroInt].GetEmployeeID(), "test GetByEmployeeID() failed")
}

func TestUserService_GetByAccountName(t *testing.T) {
	asst := assert.New(t)

	err := testUserService.GetByAccountName(testUserAccountName)
	asst.Nil(err, common.CombineMessageWithError("test GetByAccountName() failed", err))
	asst.Equal(testUserAccountName, testUserService.Users[constant.ZeroInt].GetAccountName(), "test GetByAccountName() failed")
}

func TestUserService_GetByEmail(t *testing.T) {
	asst := assert.New(t)

	err := testUserService.GetByEmail(testUserEmail)
	asst.Nil(err, common.CombineMessageWithError("test GetByEmail() failed", err))
	asst.Equal(testUserEmail, testUserService.Users[constant.ZeroInt].GetEmail(), "test GetByEmail() failed")
}

func TestUserService_GetByTelephone(t *testing.T) {
	asst := assert.New(t)

	err := testUserService.GetByTelephone(testUserTelephone)
	asst.Nil(err, common.CombineMessageWithError("test GetTelephone() failed", err))
	asst.Equal(testUserTelephone, testUserService.Users[constant.ZeroInt].GetTelephone(), "test GetTelephone() failed")
}

func TestUserService_GetByMobile(t *testing.T) {
	asst := assert.New(t)

	err := testUserService.GetByMobile(testUserMobile)
	asst.Nil(err, common.CombineMessageWithError("test GetByMobile() failed", err))
	asst.Equal(testUserMobile, testUserService.Users[constant.ZeroInt].GetMobile(), "test GetByMobile() failed")
}

func TestUserService_Create(t *testing.T) {
	asst := assert.New(t)

	err := testUserService.Create(map[string]interface{}{
		userUserNameStruct:       testUserNewUserName,
		userDepartmentNameStruct: testUserDepartmentName,
		userEmployeeIDStruct:     testUserNewEmployeeID,
		userAccountNameStruct:    testUserNewAccountName,
		userEmailStruct:          testUserNewEmail,
		userTelephoneStruct:      testUserNewTelephone,
		userMobileStruct:         testUserNewMobile,
		userRoleStruct:           testUserRole,
	})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testUserService.Delete(testUserService.GetUsers()[constant.ZeroInt].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestUserService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateUser()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testUserService.Update(entity.Identity(), map[string]interface{}{userUserNameStruct: testUserUpdateUserName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testUserService.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testUserUpdateUserName, testUserService.GetUsers()[constant.ZeroInt].GetUserName(), "test Update() failed")
	// delete
	err = testUserService.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestUserService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateUser()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testUserService.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestUserService_Marshal(t *testing.T) {
	asst := assert.New(t)

	err := testUserService.GetByID(testUserID)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	jsonBytes, err := testUserService.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}

func TestUserService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	err := testUserService.GetByID(testUserID)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	jsonBytes, err := testUserService.MarshalWithFields(userUsersStruct)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}

func TestUserService_GetAppsByID(t *testing.T) {
	asst := assert.New(t)

	err := testUserService.GetAppsByID(testUser2ID)
	asst.Nil(err, "test GetAppsByID() failed")
	asst.Equal(testUserAppID, testUserService.GetApps()[constant.ZeroInt].Identity(), "test GetAppsByID() failed")
}

func TestUserService_AddApp(t *testing.T) {
	asst := assert.New(t)

	err := testUserService.AddApp(testUser2ID, testUserNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	err = testUserService.GetAppsByID(testUser2ID)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	asst.Equal(2, len(testUserService.GetApps()), common.CombineMessageWithError("test AddApp() failed", err))
	err = testUserService.DeleteApp(testUser2ID, testUserNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
}

func TestUserService_DeleteApp(t *testing.T) {
	asst := assert.New(t)

	err := testUserService.AddApp(testUser2ID, testUserNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	err = testUserService.DeleteApp(testUser2ID, testUserNewAppID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	err = testUserService.GetAppsByID(testUser2ID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	asst.Equal(1, len(testUserService.GetApps()), common.CombineMessageWithError("test AddApp() failed", err))

}
