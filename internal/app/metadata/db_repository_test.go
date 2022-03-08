package metadata

import (
	"testing"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	testDBNewDBName    = "test_new_db_name"
	testDBUpdateDBName = "test_update_db_name"
	testDBAllDBNum     = 7
)

var testDBRepo *DBRepo

func init() {
	testInitDASMySQLPool()
	testDBRepo = NewDBRepoWithGlobal()
}

func testCreateDB() (metadata.DB, error) {
	dbInfo := NewDBInfoWithDefault(testDBDBName, testDBClusterID, testDBClusterType, testDBEnvID)
	entity, err := testDBRepo.Create(dbInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func TestDBRepoAll(t *testing.T) {
	TestDBRepo_Execute(t)
	TestDBRepo_GetAll(t)
	TestDBRepo_GetByEnv(t)
	TestDBRepo_GetByID(t)
	TestDBRepo_GetDBByNameAndClusterInfo(t)
	TestDBRepo_GetAppsByDBID(t)
	TestDBRepo_GetMySQLCLusterByID(t)
	TestDBRepo_GetAppUsersByDBID(t)
	TestDBRepo_GetUsersByDBID(t)
	TestDBRepo_GetAllUsersByDBID(t)
	TestDBRepo_Create(t)
	TestDBRepo_Update(t)
	TestDBRepo_Delete(t)
	TestDBRepo_AddDBApp(t)
	TestDBRepo_DeleteDBApp(t)
	TestDBRepo_DBAddUser(t)
	TestDBRepo_DBDeleteUser(t)
}

func TestDBRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := testDBRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestDBRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_db_info(db_name, cluster_id, cluster_type, env_id) values(?, ?, ?, ?);`
	tx, err := testDBRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, testDBNewDBName, testDBClusterID, testDBClusterType, testDBEnvID)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select db_name from t_meta_db_info where db_name = ?`
	result, err := tx.Execute(sql, testDBNewDBName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	dbName, err := result.GetString(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if dbName != testDBNewDBName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := testDBRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		asst.NotEqual(testDBNewDBName, entity.GetDBName(), "test Transaction() failed")
	}
}

func TestDBRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := testDBRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(testDBAllDBNum, len(entities), "test GetAll() failed")
}

func TestDBRepo_GetByEnv(t *testing.T) {
	asst := assert.New(t)

	entities, err := testDBRepo.GetByEnv(testDBEnvID)
	asst.Nil(err, common.CombineMessageWithError("test GetByEnv() failed", err))
	asst.Equal(testDBEnvID, entities[constant.ZeroInt].GetEnvID(), "test GetByEnv() failed")
}

func TestDBRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	db, err := testDBRepo.GetByID(testDBDBID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testDBDBName2, db.GetDBName(), "test GetByID() failed")
}

func TestDBRepo_GetDBByNameAndClusterInfo(t *testing.T) {
	asst := assert.New(t)

	db, err := testDBRepo.GetDBByNameAndClusterInfo(testDBDBName2, testDBClusterID, testDBClusterType)
	asst.Nil(err, common.CombineMessageWithError("test GetDBByNameAndClusterInfo() failed", err))
	asst.Equal(testDBDBID, db.Identity(), "test GetDBByNameAndClusterInfo() failed")
}

func TestDBRepo_GetAppsByDBID(t *testing.T) {
	asst := assert.New(t)

	apps, err := testDBRepo.GetAppsByDBID(testDBDBID)
	asst.Nil(err, common.CombineMessageWithError("test GetAppsByDBID() failed", err))
	asst.Equal(testDBAppID, apps[constant.ZeroInt].Identity(), "test GetAppsByDBID() failed")
}

func TestDBRepo_GetMySQLCLusterByID(t *testing.T) {
	asst := assert.New(t)

	mysqlCLuster, err := testDBRepo.GetMySQLCLusterByID(testDBDBID)
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLCLusterByID() failed", err))
	asst.Equal(testDBClusterID, mysqlCLuster.Identity(), "test GetMySQLCLusterByID() failed")
}

func TestDBRepo_GetAppUsersByDBID(t *testing.T) {
	asst := assert.New(t)

	appUsers, err := testDBRepo.GetAppUsersByDBID(testDBDBID)
	asst.Nil(err, common.CombineMessageWithError("test GetAppUsersByDBID() failed", err))
	asst.Equal(1, appUsers[constant.ZeroInt].Identity(), "test GetAppUsersByDBID() failed")
}

func TestDBRepo_GetUsersByDBID(t *testing.T) {
	asst := assert.New(t)

	dbUsers, err := testDBRepo.GetUsersByDBID(testDBDBID)
	asst.Nil(err, common.CombineMessageWithError("test GetUsersByDBID() failed", err))
	asst.Equal(testDBUserID, dbUsers[constant.ZeroInt].Identity(), "test GetAppUsersByDBID() failed")
}

func TestDBRepo_GetAllUsersByDBID(t *testing.T) {
	asst := assert.New(t)

	allUsers, err := testDBRepo.GetAllUsersByDBID(testDBDBID)
	asst.Nil(err, common.CombineMessageWithError("test GetAllUsersByDBID() failed", err))
	asst.Equal(1, allUsers[constant.ZeroInt].Identity(), "test GetAppUsersByDBID() failed")
}

func TestDBRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateDB()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testDBRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestDBRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateDB()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{dbDBNameStruct: testDBUpdateDBName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testDBRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	db, err := testDBRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testDBUpdateDBName, db.GetDBName(), "test Update() failed")
	// delete
	err = testDBRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestDBRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateDB()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = testDBRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestDBRepo_AddDBApp(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateDB()
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	err = testDBRepo.AddApp(entity.Identity(), testDBAppID)
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	apps, err := testDBRepo.GetAppsByDBID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test AddApp() failed", err))
	asst.Equal(testDBAppID, apps[constant.ZeroInt].Identity(), "test AddApp() failed")
	// delete
	err = testDBRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestDBRepo_DeleteDBApp(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateDB()
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	err = testDBRepo.AddApp(entity.Identity(), testDBAppID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	err = testDBRepo.DeleteApp(entity.Identity(), testDBAppID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	apps, err := testDBRepo.GetAppsByDBID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DeleteApp() failed", err))
	asst.Zero(len(apps), "test DeleteApp() failed")
	// delete
	err = testDBRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestDBRepo_DBAddUser(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateDB()
	asst.Nil(err, common.CombineMessageWithError("test DBAddUser() failed", err))
	err = testDBRepo.DBAddUser(entity.Identity(), testDBUserID)
	asst.Nil(err, common.CombineMessageWithError("test DBAddUser() failed", err))
	users, err := testDBRepo.GetUsersByDBID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DBAddUser() failed", err))
	asst.Equal(testDBUserID, users[constant.ZeroInt].Identity(), "test DBAddUser() failed")
	// delete
	err = testDBRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestDBRepo_DBDeleteUser(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateDB()
	asst.Nil(err, common.CombineMessageWithError("test DBDeleteUser() failed", err))
	err = testDBRepo.DBAddUser(entity.Identity(), testDBUserID)
	asst.Nil(err, common.CombineMessageWithError("test DBDeleteUser() failed", err))
	err = testDBRepo.DBDeleteUser(entity.Identity(), testDBUserID)
	asst.Nil(err, common.CombineMessageWithError("test DBDeleteUser() failed", err))
	users, err := testDBRepo.GetUsersByDBID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test DBDeleteUser() failed", err))
	asst.Zero(len(users), "test DBDeleteUser() failed")
	// delete
	err = testDBRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
