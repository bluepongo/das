package metadata

import (
	"testing"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	testUserNewUserName    = "test_new_user_name"
	testUserUpdateUserName = "test_update_user_name"
	testUserNewEmployeeID  = "200001"
	testUserNewAccountName = "test_new_account_name"
	testUserNewEmail       = "test_new_email@163.com"
	testUserNewTelephone   = "02112345678"
	testUserNewMobile      = "13112345678"
)

var testUserRepo *UserRepo

func init() {
	initDASMySQLPool()
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
	TestUserRepo_GetByName(t)
	TestUserRepo_GetByEmployeeID(t)
	TestUserRepo_GetByAccountName(t)
	TestUserRepo_GetByEmail(t)
	TestUserRepo_GetByTelephone(t)
	TestUserRepo_GetByMobile(t)
	TestUserRepo_GetID(t)
	TestUserRepo_Create(t)
	TestUserRepo_Update(t)
	TestUserRepo_Delete(t)
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

	entity, err := testUserRepo.GetByID(testUserID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testUserID, entity.Identity(), "test GetByID() failed")
}

func TestUserRepo_GetByName(t *testing.T) {
	asst := assert.New(t)

	entities, err := testUserRepo.GetByName(testUserUserName)
	asst.Nil(err, common.CombineMessageWithError("test GetByName() failed", err))
	asst.Equal(1, len(entities), "test GetByName() failed")
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
