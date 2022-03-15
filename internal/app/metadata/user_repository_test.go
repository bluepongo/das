package metadata

import (
	"testing"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	testUserNewUserName         = "test_new_user_name"
	testUserUpdateUserName      = "test_update_user_name"
	testUserNewEmployeeID       = "200001"
	testUserNewAccountName      = "test_new_account_name"
	testUserNewEmail            = "test_new_account_name@163.com"
	testUserNewTelephone        = "02112345678"
	testUserNewMobile           = "13112345678"
	testUserAppID               = 1
	testUserDBID                = 1
	testUserDB2ID               = 2
	testUserMiddlewareClusterID = 1
	testUserMySQLClusterID      = 1
	testUserMySQLServerID       = 1
	testUser2ID                 = 15
)

var testUserRepo *UserRepo

func init() {
	testInitDASMySQLPool()
	testUserRepo = NewUserRepoWithGlobal()
}

func testCreateUser() (metadata.User, error) {
	userInfo := NewUserInfoWithDefault(
		testUserNewUserName,
		testUserDepartmentName,
		testUserNewAccountName,
		testUserNewEmail,
		testUserRole,
	)
	entity, err := testUserRepo.Create(userInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func TestUserRepoAll(t *testing.T) {
	TestUserRepo_Execute(t)
	TestUserRepo_Transaction(t)
	TestUserRepo_GetAll(t)
	TestUserRepo_GetByID(t)
	TestUserRepo_GetByUserName(t)
	TestUserRepo_GetByEmployeeID(t)
	TestUserRepo_GetByAccountName(t)
	TestUserRepo_GetByEmail(t)
	TestUserRepo_GetByTelephone(t)
	TestUserRepo_GetByMobile(t)
	TestUserRepo_GetID(t)
	TestUserRepo_Create(t)
	TestUserRepo_Update(t)
	TestUserRepo_Delete(t)
	TestUserRepo_GetAppsByUserID(t)
	TestUserRepo_GetDBsByUserID(t)
	TestUserRepo_GetMiddlewareClustersByUserID(t)
	TestUserRepo_GetMySQLClustersByUserID(t)
	TestUserRepo_GetAllMySQLServersByUserID(t)
	TestUserRepo_GetByAccountNameOrEmployeeID(t)

}

func TestUserRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := testUserRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestUserRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `
		insert into t_meta_user_info(user_name,department_name,employee_id,account_name,email,telephone,mobile,role)
		values(?,?,?,?,?,?,?,?);`
	tx, err := testUserRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql,
		testUserNewUserName,
		testUserDepartmentName,
		testUserNewEmployeeID,
		testUserNewAccountName,
		testUserNewEmail,
		testUserNewTelephone,
		testUserNewMobile,
		testUserRole)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select account_name from t_meta_user_info where account_name=?`
	result, err := tx.Execute(sql, testUserNewAccountName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	accountName, err := result.GetString(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if accountName != testUserNewAccountName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := testUserRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		if entity.GetAccountName() == testUserNewAccountName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestUserRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := testUserRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(1, len(entities), "test GetAll() failed")
}

func TestUserRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := testUserRepo.GetByID(99)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testUserID, entity.Identity(), "test GetByID() failed")
}

func TestUserRepo_GetByUserName(t *testing.T) {
	asst := assert.New(t)

	entities, err := testUserRepo.GetByUserName(testUserUserName)
	asst.Nil(err, common.CombineMessageWithError("test GetByUserName() failed", err))
	asst.Equal(1, len(entities), "test GetByUserName() failed")
}

func TestUserRepo_GetByEmployeeID(t *testing.T) {
	asst := assert.New(t)

	entity, err := testUserRepo.GetByEmployeeID(testUserEmployeeID)
	asst.Nil(err, common.CombineMessageWithError("test GetByEmployeeID failed", err))
	asst.Equal(testUserEmployeeID, entity.GetEmployeeID(), "test GetByEmployeeID failed")
}

func TestUserRepo_GetByAccountName(t *testing.T) {
	asst := assert.New(t)

	entity, err := testUserRepo.GetByAccountName(testUserAccountName)
	asst.Nil(err, common.CombineMessageWithError("test GetByAccountName failed", err))
	asst.Equal(testUserAccountName, entity.GetAccountName(), "test GetByAccountName failed")
}

func TestUserRepo_GetByAccountNameOrEmployeeID(t *testing.T) {
	asst := assert.New(t)

	entity, err := testUserRepo.GetByAccountNameOrEmployeeID(testUserAccountName)
	asst.Nil(err, common.CombineMessageWithError("test GetByAccountNameOrEmployeeID failed", err))
	asst.Equal(testUserAccountName, entity.GetAccountName(), "test GetByAccountNameOrEmployeeID failed")

	entity, err = testUserRepo.GetByAccountNameOrEmployeeID(testUserEmployeeID)
	asst.Nil(err, common.CombineMessageWithError("test GetByAccountNameOrEmployeeID failed", err))
	asst.Equal(testUserAccountName, entity.GetAccountName(), "test GetByAccountNameOrEmployeeID failed")
}

func TestUserRepo_GetByEmail(t *testing.T) {
	asst := assert.New(t)

	entity, err := testUserRepo.GetByEmail(testUserEmail)
	asst.Nil(err, common.CombineMessageWithError("test GetByEmail failed", err))
	asst.Equal(testUserEmail, entity.GetEmail(), "test GetByEmail failed")

}

func TestUserRepo_GetByTelephone(t *testing.T) {
	asst := assert.New(t)

	entity, err := testUserRepo.GetByTelephone(testUserTelephone)
	asst.Nil(err, common.CombineMessageWithError("test GetByTelephone failed", err))
	asst.Equal(testUserTelephone, entity.GetTelephone(), "test GetByTelephone failed")
}

func TestUserRepo_GetByMobile(t *testing.T) {
	asst := assert.New(t)

	entity, err := testUserRepo.GetByMobile(testUserMobile)
	asst.Nil(err, common.CombineMessageWithError("test GetByMobile failed", err))
	asst.Equal(testUserMobile, entity.GetMobile(), "test GetByMobile failed")
}

func TestUserRepo_GetID(t *testing.T) {
	asst := assert.New(t)

	userID, err := testUserRepo.GetID(testUserAccountName)
	asst.Nil(err, common.CombineMessageWithError("test GetID failed", err))
	asst.Equal(testUserID, userID, "test GetID failed")
}

func TestUserRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateUser()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testUserRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestUserRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateUser()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{userUserNameStruct: testUserUpdateUserName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testUserRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = testUserRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testUserUpdateUserName, entity.GetUserName(), "test Update() failed")
	// delete
	err = testUserRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestUserRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateUser()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testUserRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestUserRepo_GetAppsByUserID(t *testing.T) {
	asst := assert.New(t)

	apps, err := testUserRepo.GetAppsByUserID(testUser2ID)
	asst.Nil(err, common.CombineMessageWithError("test GetAppsByUserID() failed", err))
	asst.Equal(testUserAppID, apps[constant.ZeroInt].Identity(), "test GetAppsByUserID() failed")
}

func TestUserRepo_GetDBsByUserID(t *testing.T) {
	asst := assert.New(t)

	dbs, err := testUserRepo.GetDBsByUserID(testUserID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBsByUserID() failed", err))
	asst.Equal(testUserDB2ID, dbs[constant.ZeroInt].Identity(), "test GetDBsByUserID() failed")
}

func TestUserRepo_GetMiddlewareClustersByUserID(t *testing.T) {
	asst := assert.New(t)

	mcs, err := testUserRepo.GetMiddlewareClustersByUserID(testUserID)
	asst.Nil(err, common.CombineMessageWithError("test GetMiddlewareClustersByUserID() failed", err))
	asst.Equal(testUserMiddlewareClusterID, mcs[constant.ZeroInt].Identity(), "test GetMiddlewareClustersByUserID() failed")
}

func TestUserRepo_GetMySQLClustersByUserID(t *testing.T) {
	asst := assert.New(t)

	mcs, err := testUserRepo.GetMySQLClustersByUserID(testUserID)
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLClustersByUserID() failed", err))
	asst.Equal(testUserMySQLClusterID, mcs[constant.ZeroInt].Identity(), "test GetMySQLClustersByUserID() failed")
}

func TestUserRepo_GetAllMySQLServersByUserID(t *testing.T) {
	asst := assert.New(t)

	mss, err := testUserRepo.GetAllMySQLServersByUserID(testUserID)
	asst.Nil(err, common.CombineMessageWithError("test GetAllMySQLServersByUserID() failed", err))
	asst.Equal(testUserMySQLServerID, mss[constant.ZeroInt].Identity(), "test GetAllMySQLServersByUserID() failed")
}
