package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"

	"github.com/romberli/das/internal/dependency/metadata"
)

const (
	testEnvNewEnvName    = "test_new_env_name"
	testEnvUpdateEnvName = "test_update_env_name"
)

var testEnvRepo *EnvRepo

func init() {
	testInitDASMySQLPool()
	testEnvRepo = NewEnvRepoWithGlobal()
}

func testCreateEnv() (metadata.Env, error) {
	envInfo := NewEnvInfoWithDefault(testEnvNewEnvName)
	entity, err := testEnvRepo.Create(envInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func TestEnvRepoAll(t *testing.T) {
	TestEnvRepo_Execute(t)
	TestEnvRepo_Transaction(t)
	TestEnvRepo_GetAll(t)
	TestEnvRepo_GetByID(t)
	TestEnvRepo_GetID(t)
	TestEnvRepo_GetEnvByName(t)
	TestEnvRepo_Create(t)
	TestEnvRepo_Update(t)
	TestEnvRepo_Delete(t)
}

func TestEnvRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := testEnvRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestEnvRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_env_info(env_name) values(?);`
	tx, err := testEnvRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, testEnvNewEnvName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select env_name from t_meta_env_info where env_name=?`
	result, err := tx.Execute(sql, testEnvNewEnvName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	envName, err := result.GetString(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if envName != testEnvNewEnvName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	envs, err := testEnvRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, env := range envs {
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		if env.GetEnvName() == testEnvNewEnvName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestEnvRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := testEnvRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(6, len(entities), "test GetAll() failed")
}

func TestEnvRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := testEnvRepo.GetByID(testEnvEnvID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testEnvEnvName, entity.GetEnvName(), "test GetByID() failed")
}

func TestEnvRepo_GetID(t *testing.T) {
	asst := assert.New(t)

	entity, err := testEnvRepo.GetEnvByName(testEnvEnvName)
	asst.Nil(err, common.CombineMessageWithError("test GetID() failed", err))
	asst.Equal(testEnvEnvID, entity.Identity(), "test GetID() failed")
}

func TestEnvRepo_GetEnvByName(t *testing.T) {
	asst := assert.New(t)

	entity, err := testEnvRepo.GetEnvByName(testEnvEnvName)
	asst.Nil(err, common.CombineMessageWithError("test GetEnvByName() failed", err))
	asst.Equal(testEnvEnvID, entity.Identity(), "test GetEnvByName() failed")
}

func TestEnvRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateEnv()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	asst.Equal(testEnvNewEnvName, entity.GetEnvName(), "test Create() failed")
	// delete
	err = testEnvRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestEnvRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateEnv()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{envEnvNameStruct: testEnvUpdateEnvName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testEnvRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = testEnvRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testEnvUpdateEnvName, entity.GetEnvName(), "test Update() failed")
	// delete
	err = testEnvRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestEnvRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateEnv()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	err = testEnvRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
