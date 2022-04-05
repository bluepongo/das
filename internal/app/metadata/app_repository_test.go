package metadata

import (
	"os"
	"testing"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
)

const (
	testDASMySQLAddr = "192.168.137.11:3306"
	testDASMySQLName = "das"
	testDASMySQLUser = "root"
	testDASMySQLPass = "root"

	testAppNewAppName    = "test_new_app_name"
	testAppUpdateAppName = "test_update_app_name"
	testAppDBID          = 1
	testAppUserID        = 1
	testAppUserID2       = 14
)

var testAppRepo *AppRepo

func init() {
	testInitDASMySQLPool()
	testAppRepo = NewAppRepoWithGlobal()
}

func testInitDASMySQLPool() {
	var err error

	if global.DASMySQLPool == nil {
		global.DASMySQLPool, err = mysql.NewPoolWithDefault(testDASMySQLAddr, testDASMySQLName, testDASMySQLUser, testDASMySQLPass)
		if err != nil {
			log.Error(common.CombineMessageWithError("testInitDASMySQLPool() failed", err))
			os.Exit(constant.DefaultAbnormalExitCode)
		}
	}
}

func testCreateApp() (metadata.App, error) {
	appSystemInfo := NewAppInfoWithDefault(
		testAppNewAppName,
		testAppLevel,
	)
	entity, err := testAppRepo.Create(appSystemInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func TestAppRepoAll(t *testing.T) {
	TestAppRepo_Execute(t)
	TestAppRepo_GetAll(t)
	TestAppRepo_GetByID(t)
	TestAppRepo_GetAppByName(t)
	TestAppRepo_Create(t)
	TestAppRepo_Update(t)
	TestAppRepo_Delete(t)
	TestAppRepo_AddAppDB(t)
	TestAppRepo_DeleteAppDB(t)
	TestAppRepo_AddAppUser(t)
	TestAppRepo_DeleteAppUser(t)
	TestAppRepo_GetDBsByAppID(t)
	TestAppRepo_GetUsersByAppID(t)

}

func TestAppRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := testAppRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestAppRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_app_info(app_name,level) values(?,?);`
	tx, err := testAppRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, testAppNewAppName, testAppLevel)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select app_name from t_meta_app_info where app_name=?`
	result, err := tx.Execute(sql, testAppNewAppName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	appName, err := result.GetString(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	asst.Equal(appName, testAppNewAppName, "test Transaction() failed")
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := testAppRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		asst.NotEqual(entity.GetAppName(), testAppNewAppName, "test Transaction() failed")
	}
}

func TestAppRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := testAppRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(3, len(entities), "test GetAll() failed")
}

func TestAppRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := testAppRepo.GetByID(testAppAppID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	appName := entity.GetAppName()
	asst.Equal(testAppAppName, appName, "test GetByID() failed")
}

func TestAppRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateApp()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testAppRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestAppRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateApp()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{appAppNameStruct: testAppUpdateAppName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testAppRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = testAppRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	appName := entity.GetAppName()
	asst.Equal(testAppUpdateAppName, appName, "test Update() failed")
	// delete
	err = testAppRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestAppRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateApp()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = testAppRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestAppRepo_GetAppByName(t *testing.T) {
	asst := assert.New(t)

	_, err := testCreateApp()
	asst.Nil(err, common.CombineMessageWithError("test GetAppByName() failed", err))
	entity, err := testAppRepo.GetAppByName(testAppNewAppName)
	asst.Nil(err, common.CombineMessageWithError("test GetAppByName() failed", err))
	asst.Equal(testAppNewAppName, entity.GetAppName(), common.CombineMessageWithError("test GetAppByName() failed", err))
	// delete
	err = testAppRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetAppByName() failed", err))
}

func TestAppRepo_GetDBsByAppID(t *testing.T) {
	asst := assert.New(t)

	dbs, err := testAppRepo.GetDBsByAppID(testAppAppID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBsByAppID() failed", err))
	asst.Equal(1, len(dbs), "test GetDBsByAppID() failed")

}

func TestAppRepo_GetUsersByAppID(t *testing.T) {
	asst := assert.New(t)

	users, err := testAppRepo.GetUsersByAppID(testAppAppID)
	asst.Nil(err, common.CombineMessageWithError("test GetUsersByAppID() failed", err))
	asst.Equal(2, len(users), "test GetUsersByAppID() failed")

}
func TestAppRepo_AddAppDB(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateApp()
	asst.Nil(err, common.CombineMessageWithError("test AddAppDB() failed", err))
	err = testAppRepo.AddDB(entity.Identity(), testAppDBID)
	asst.Nil(err, common.CombineMessageWithError("test AddAppDB() failed", err))
	dbs, err := entity.GetDBs()
	asst.Nil(err, common.CombineMessageWithError("test AddAppDB() failed", err))
	asst.Equal(1, len(dbs), "test AddAppDB() failed")
	// delete
	err = testAppRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddAppDB() failed", err))
}

func TestAppRepo_DeleteAppDB(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateApp()
	asst.Nil(err, common.CombineMessageWithError("test DeleteAppDB() failed", err))
	err = testAppRepo.DeleteDB(entity.Identity(), testAppDBID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteAppDB() failed", err))
	dbs, err := entity.GetDBs()
	asst.Nil(err, common.CombineMessageWithError("test DeleteAppDB() failed", err))
	asst.Zero(len(dbs), "test DeleteAppDB() failed")
	// delete
	err = testAppRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DeleteAppDB() failed", err))
}

func TestAppRepo_AddAppUser(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateApp()
	asst.Nil(err, common.CombineMessageWithError("test AddAppUser() failed", err))
	err = testAppRepo.AddUser(entity.Identity(), testAppUserID)
	asst.Nil(err, common.CombineMessageWithError("test AddAppUser() failed", err))
	users, err := entity.GetUsers()
	asst.Nil(err, common.CombineMessageWithError("test AddAppUser() failed", err))
	asst.Equal(1, len(users), "test AddAppUser() failed")
	// delete
	err = testAppRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddAppUser() failed", err))
}

func TestAppRepo_DeleteAppUser(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateApp()
	asst.Nil(err, common.CombineMessageWithError("test DeleteAppUser() failed", err))
	err = testAppRepo.DeleteUser(entity.Identity(), testAppUserID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteAppUser() failed", err))
	users, err := entity.GetUsers()
	asst.Nil(err, common.CombineMessageWithError("test DeleteAppUser() failed", err))
	asst.Zero(len(users), "test DeleteAppUser() failed")
	// delete
	err = testAppRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DeleteAppUser() failed", err))
}
