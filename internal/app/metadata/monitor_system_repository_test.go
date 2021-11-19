package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"

	"github.com/romberli/das/internal/dependency/metadata"
)

const (
	testMonitorSystemNewSystemName    = "test_new_system_name"
	testMonitorSystemUpdateSystemName = "test_update_system_name"
	testMonitorSystemNewPortNum       = 8080
)

var testMonitorSystemRepo *MonitorSystemRepo

func init() {
	testInitDASMySQLPool()
	testMonitorSystemRepo = NewMonitorSystemRepoWithGlobal()
}

func testCreateMonitorSystem() (metadata.MonitorSystem, error) {
	monitorSystemInfo := NewMonitorSystemInfoWithDefault(
		testMonitorSystemNewSystemName,
		testMonitorSystemSystemType,
		testMonitorSystemHostIP,
		testMonitorSystemNewPortNum,
		testMonitorSystemPortNumSlow,
		testMonitorSystemBaseUrl,
		testMonitorSystemEnvID,
	)
	entity, err := testMonitorSystemRepo.Create(monitorSystemInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func TestMonitorSystemRepoAll(t *testing.T) {
	TestMonitorSystemRepo_Execute(t)
	TestMonitorSystemRepo_Transaction(t)
	TestMonitorSystemRepo_GetAll(t)
	TestMonitorSystemRepo_GetByEnv(t)
	TestMonitorSystemRepo_GetByID(t)
	TestMonitorSystemRepo_Create(t)
	TestMonitorSystemRepo_Update(t)
	TestMonitorSystemRepo_Delete(t)
}

func TestMonitorSystemRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := testMonitorSystemRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestMonitorSystemRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_monitor_system_info(system_name, system_type, host_ip, port_num, port_num_slow, base_url, env_id) values(?, ?, ?, ?, ?, ?, ?);`
	tx, err := testMonitorSystemRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(
		sql,
		testMonitorSystemNewSystemName,
		testMonitorSystemSystemType,
		testMonitorSystemHostIP,
		testMonitorSystemNewPortNum,
		testMonitorSystemPortNumSlow,
		testMonitorSystemBaseUrl,
		testMonitorSystemEnvID,
	)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select system_name from t_meta_monitor_system_info where system_name = ?`
	result, err := tx.Execute(sql, testMonitorSystemNewSystemName)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	monitorSystemName, err := result.GetString(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if monitorSystemName != testMonitorSystemNewSystemName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := testMonitorSystemRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		if entity.GetSystemName() == testMonitorSystemNewSystemName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestMonitorSystemRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := testMonitorSystemRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(2, len(entities), "test GetAll() failed")
}

func TestMonitorSystemRepo_GetByEnv(t *testing.T) {
	asst := assert.New(t)

	entities, err := testMonitorSystemRepo.GetByEnv(testMonitorSystemEnvID)
	asst.Nil(err, common.CombineMessageWithError("test GetByEnv() failed", err))
	asst.Equal(testMonitorSystemEnvID, entities[constant.ZeroInt].GetEnvID(), common.CombineMessageWithError("test GetByEnv() failed", err))
}

func TestMonitorSystemRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := testMonitorSystemRepo.GetByID(testMonitorSystemID)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testMonitorSystemSystemName, entity.GetSystemName(), "test GetByID() failed")
}

func TestMonitorSystemRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = testMonitorSystemRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMonitorSystemRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{monitorSystemSystemNameStruct: testMonitorSystemUpdateSystemName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = testMonitorSystemRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = testMonitorSystemRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testMonitorSystemUpdateSystemName, entity.GetSystemName(), "test Update() failed")
	// delete
	err = testMonitorSystemRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMonitorSystemRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := testCreateMonitorSystem()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = testMonitorSystemRepo.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
