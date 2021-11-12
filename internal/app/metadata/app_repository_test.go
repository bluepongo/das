package metadata

import (
	"testing"

	"github.com/romberli/das/global"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"

	"github.com/romberli/das/internal/dependency/metadata"
)

const (
	testDASMySQLAddr = "192.168.10.210:3306"
	testDASMySQLName = "das"
	testDASMySQLUser = "root"
	testDASMySQLPass = "root"

	testAppNewAppName    = "test_new_app_name"
	testAppUpdateAppName = "test_update_app_name"
	testAppDBID          = 1
)

var testAppRepo *AppRepo

func init() {
	initDASMySQLPool()
	testAppRepo = NewAppRepoWithGlobal()
}

func initDASMySQLPool() {
	var err error

	global.DASMySQLPool, err = mysql.NewPoolWithDefault(testDASMySQLAddr, testDASMySQLName, testDASMySQLUser, testDASMySQLPass)
	log.Infof("pool: %v, error: %v", global.DASMySQLPool, err)
	if err != nil {
		log.Error(common.CombineMessageWithError("initRepository() failed", err))
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
	TestAppRepo_GetDBsByID(t)

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

	sql := `insert into t_meta_app_info(app_name,level,owner_id) values(?,?,?);`
	tx, err := testAppRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, testAppNewAppName, testAppLevel, testAppOwnerID)
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

func TestAppRepo_GetDBsByID(t *testing.T) {
	asst := assert.New(t)

	dbs, err := testAppRepo.GetDBsByID(testAppAppID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBsByID() failed", err))
	asst.Equal(1, len(dbs), "test GetDBsByID() failed")

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
