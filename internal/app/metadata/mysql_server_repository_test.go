package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"

	"github.com/romberli/das/internal/dependency"
)

const (
	testInitServerID = 1
	testInitHostIP         = "host_ip_init"
	testInitPortNum        = 3306
	testTransactionServerID = 2
	testTransactionHostIP  = "host_ip_need_rollback"
	testTransactionPortNum = 3308
	testInsertHostIP       = "host_ip_insert"
	testInsertPortNum      = 3307
	testUpdateHostIP       = "host_ip_update"
	testUpdatePortNum      = 3309
)

var mysqlServerRepo = initMySQLServerRepo()

func initMySQLServerRepo() *MySQLServerRepo {
	pool, err := mysql.NewMySQLPoolWithDefault(addr, dbName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initMySQLServerRepo() failed", err))
		return nil
	}

	return NewMySQLServerRepo(pool)
}

func createMySQLServer() (dependency.Entity, error) {
	mysqlServerInfo := NewMySQLServerInfoWithDefault(
		defaultMySQLServerInfoClusterID,
		defaultMySQLServerInfoServerName,
		defaultMySQLServerInfoHostIP,
		defaultMySQLServerInfoPortNum,
		defaultMySQLServerInfoDeploymentType)
	entity, err := mysqlServerRepo.Create(mysqlServerInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteMySQLServerByID(id string) error {
	sql := `delete from t_meta_mysql_server_info where id = ?`
	_, err := mysqlServerRepo.Execute(sql, id)
	return err
}

func TestMySQLServerRepoAll(t *testing.T) {
	TestMySQLServerRepo_Execute(t)
	TestMySQLServerRepo_Create(t)
	TestMySQLServerRepo_GetAll(t)
	TestMySQLServerRepo_GetByID(t)
	TestMySQLServerRepo_Update(t)
	TestMySQLServerRepo_Delete(t)
}

func TestMySQLServerRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := mysqlServerRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, int(r), "test Execute() failed")
}

func TestMySQLServerRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `
	insert into t_meta_mysql_server_info(
		id, cluster_id, server_name, host_ip, port_num, deployment_type, version) 
	values(?,?,?,?,?,?,?);`

	tx, err := mysqlServerRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql,
		testTransactionServerID,
		defaultMySQLServerInfoClusterID,
		defaultMySQLServerInfoServerName,
		testTransactionHostIP,
		testTransactionPortNum,
		defaultMySQLServerInfoDeploymentType,
		defaultMySQLServerInfoVersion)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select host_ip, port_num from t_meta_mysql_server_info where host_ip=? and port_num=?`
	result, err := tx.Execute(sql, testTransactionHostIP, testTransactionPortNum)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	hostIP, err := result.GetString(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if hostIP != testTransactionHostIP {
		asst.Fail("test Transaction() failed")
	}
	portNum, err := result.GetInt(0, 1)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if portNum != testTransactionPortNum {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := mysqlServerRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		hostIP, err := entity.Get(hostIPStruct)
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		portNum, err := entity.Get(portNumStruct)
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		if hostIP == testTransactionHostIP && portNum == testTransactionPortNum {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestMySQLServerRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	sql := `
	insert into t_meta_mysql_server_info(
		id, cluster_id, server_name, host_ip, port_num, deployment_type, version) 
	values(?,?,?,?,?,?,?);`

	// init data avoid empty set
	_, err := mysqlServerRepo.Execute(sql,
		testInitServerID,
		defaultMySQLServerInfoClusterID,
		defaultMySQLServerInfoServerName,
		testInitHostIP,
		testInitPortNum,
		defaultMySQLServerInfoDeploymentType,
		defaultMySQLServerInfoVersion)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))

	entities, err := mysqlServerRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	hostIP, err := entities[0].Get("HostIP")
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(testInitHostIP, hostIP.(string), "test GetAll() failed")
	portNum, err := entities[0].Get("PortNum")
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(testInitPortNum, portNum.(int), "test GetAll() failed")
}

func TestMySQLServerRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := mysqlServerRepo.GetByID("1")
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	hostIP, err := entity.Get(hostIPStruct)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testInitHostIP, hostIP.(string), "test GetByID() failed")
	portNum, err := entity.Get(portNumStruct)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(testInitPortNum, portNum.(int), "test GetByID() failed")
}

func TestMySQLServerRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMySQLServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMySQLServerRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{
		hostIPStruct:  testUpdateHostIP,
		portNumStruct: testUpdatePortNum})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = mysqlServerRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = mysqlServerRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	hostIP, err := entity.Get(hostIPStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testUpdateHostIP, hostIP, "test Update() failed")
	portNum, err := entity.Get(portNumStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(testUpdatePortNum, portNum, "test Update() failed")
	// delete
	err = deleteMySQLServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMySQLServerRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMySQLServer()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMySQLServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}