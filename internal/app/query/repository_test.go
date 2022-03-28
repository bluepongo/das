package query

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/romberli/das/config"
	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/clickhouse"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	testDASMySQLAddr = "192.168.137.11:3306"
	testDASMySQLName = "das"
	testDASMySQLUser = "root"
	testDASMySQLPass = "root"

	testPMM1ServiceName    = "192-168-137-11-mysql"
	testPMM1MySQLClusterID = 2
	testPMM1MySQLServerID  = 2
	testPMM1MySQLHostIP    = "192.168.137.11"
	testPMM1MySQLPortNum   = 3306
	testPMM1DBID           = 3
	testPMM1DBName         = "das"
	testPMM1SQLID          = "999ECD050D719733"

	testPMM2ServiceName    = "192-168-137-11-mysql"
	testPMM2MySQLClusterID = 1
	testPMM2MySQLServerID  = 1
	testPMM2MySQLHostIP    = "192.168.137.11"
	testPMM2MySQLPortNum   = 3306
	testPMM2DBID           = 1
	testPMM2DBName         = "das"
	testPMM2SQLID          = "F4F85858E527B409"

	testMinRowsExamined = 1
)

var (
	testPMMVersion     int
	testServiceName    string
	testMySQLClusterID int
	testMySQLServerID  int
	testMySQLHostIP    string
	testMySQLPortNum   int
	testDBID           int
	testDBName         string
	testSQLID          string

	testDASRepo        *DASRepo
	testMySQLRepo      *MySQLRepo
	testClickhouseRepo *ClickhouseRepo
)

func init() {
	testPMMVersion = 2
	testInitMySQLInfo()
	testInitDASMySQLPool()
	testInitViper()

	testDASRepo = NewDASRepoWithGlobal()

	switch testPMMVersion {
	case 1:
		testMySQLRepo = testInitMySQLRepo()
	case 2:
		testClickhouseRepo = testInitClickhouseRepo()
	default:
		log.Errorf(fmt.Sprintf("pmm version should be 1 or 2, %d is not valid", testPMMVersion))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
}

func testInitMySQLInfo() {
	switch testPMMVersion {
	case 1:
		testServiceName = testPMM1ServiceName
		testMySQLClusterID = testPMM1MySQLClusterID
		testMySQLServerID = testPMM1MySQLServerID
		testMySQLHostIP = testPMM1MySQLHostIP
		testMySQLPortNum = testPMM1MySQLPortNum
		testDBID = testPMM1DBID
		testDBName = testPMM1DBName
		testSQLID = testPMM1SQLID
	case 2:
		testServiceName = testPMM2ServiceName
		testMySQLClusterID = testPMM2MySQLClusterID
		testMySQLServerID = testPMM2MySQLServerID
		testMySQLHostIP = testPMM2MySQLHostIP
		testMySQLPortNum = testPMM2MySQLPortNum
		testDBID = testPMM2DBID
		testDBName = testPMM2DBName
		testSQLID = testPMM2SQLID
	default:
		log.Errorf(fmt.Sprintf("pmm version should be 1 or 2, %d is not valid", testPMMVersion))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
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

func testInitViper() {
	viper.Set(config.DBMonitorMySQLUserKey, config.DefaultDBMonitorMySQLUser)
	viper.Set(config.DBMonitorMySQLPassKey, config.DefaultDBMonitorMySQLPass)
	viper.Set(config.DBMonitorClickhouseUserKey, config.DefaultDBMonitorClickhouseUser)
	viper.Set(config.DBMonitorClickhousePassKey, config.DefaultDBMonitorClickhousePass)
	viper.Set(config.QueryMinRowsExaminedKey, testMinRowsExamined)
}

func testInitMySQLRepo() *MySQLRepo {
	mysqlServerService := metadata.NewMySQLServerServiceWithDefault()
	err := mysqlServerService.GetByID(testMySQLServerID)
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitMySQLRepo() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
	monitorSystem, err := mysqlServerService.GetMySQLServers()[constant.ZeroInt].GetMonitorSystem()
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitMySQLRepo() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
	addr := fmt.Sprintf("%s:%d", monitorSystem.GetHostIP(), monitorSystem.GetPortNumSlow())
	dbUser := viper.GetString(config.DBMonitorMySQLUserKey)
	dbPass := viper.GetString(config.DBMonitorMySQLPassKey)
	conn, err := mysql.NewConn(addr, pmmMySQLDBName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitMySQLRepo() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}

	return NewMySQLRepo(NewConfigWithDefault(), conn)
}

func testInitClickhouseRepo() *ClickhouseRepo {
	mysqlServerService := metadata.NewMySQLServerServiceWithDefault()
	err := mysqlServerService.GetByID(testMySQLServerID)
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitClickhouseRepo() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
	monitorSystem, err := mysqlServerService.GetMySQLServers()[constant.ZeroInt].GetMonitorSystem()
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitClickhouseRepo() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
	addr := fmt.Sprintf("%s:%d", monitorSystem.GetHostIP(), monitorSystem.GetPortNumSlow())
	dbUser := viper.GetString(config.DBMonitorClickhouseUserKey)
	dbPass := viper.GetString(config.DBMonitorClickhousePassKey)
	conn, err := clickhouse.NewConnWithDefault(addr, pmmClickhouseDBName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitClickhouseRepo() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}

	return NewClickHouseRepo(NewConfigWithDefault(), conn)
}

func TestQueryRepository_All(t *testing.T) {
	TestDASRepo_Save(t)
	// test PMM1.x
	TestQueryRepository_PMM1(t)
	// test PMM2.x
	TestQueryRepository_PMM2(t)
}

func TestQueryRepository_PMM1(t *testing.T) {
	testPMMVersion = 1
	testInitMySQLInfo()
	TestMySQLRepo_GetByServiceNames(t)
	TestMySQLRepo_GetByDBName(t)
	TestMySQLRepo_GetBySQLID(t)
}

func TestQueryRepository_PMM2(t *testing.T) {
	testPMMVersion = 2
	testInitMySQLInfo()
	TestClickhouseRepo_GetByServiceNames(t)
	TestClickhouseRepo_GetByDBName(t)
	TestClickhouseRepo_GetBySQLID(t)
}

func TestDASRepo_Save(t *testing.T) {
	asst := assert.New(t)

	err := testDASRepo.Save(
		testMySQLClusterID,
		testMySQLServerID,
		testDBID,
		testSQLID,
		time.Now().Add(-constant.Week),
		time.Now(),
		defaultLimit,
		defaultOffset,
	)
	asst.Nil(err, common.CombineMessageWithError("test Save() failed", err))
}

func TestMySQLRepo_GetByServiceNames(t *testing.T) {
	asst := assert.New(t)

	queries, err := testMySQLRepo.GetByServiceNames([]string{testServiceName})
	asst.Nil(err, common.CombineMessageWithError("test GetByServiceNames() failed", err))
	asst.GreaterOrEqual(len(queries), constant.ZeroInt, "test TestMySQLRepo_GetByServiceNames() failed")
}

func TestMySQLRepo_GetByDBName(t *testing.T) {
	asst := assert.New(t)

	queries, err := testMySQLRepo.GetByDBName(testServiceName, testDBName)
	asst.Nil(err, common.CombineMessageWithError("test GetByDBName() failed", err))
	asst.GreaterOrEqual(len(queries), constant.ZeroInt, "test GetByDBName() failed")
}

func TestMySQLRepo_GetBySQLID(t *testing.T) {
	asst := assert.New(t)

	query, err := testMySQLRepo.GetBySQLID(testServiceName, testSQLID)
	asst.Nil(err, common.CombineMessageWithError("test GetBySQLID() failed", err))
	asst.NotNil(query, "test GetBySQLID() failed")
}

func TestClickhouseRepo_GetByServiceNames(t *testing.T) {
	asst := assert.New(t)

	queries, err := testClickhouseRepo.GetByServiceNames([]string{testServiceName})
	asst.Nil(err, common.CombineMessageWithError("test GetByServiceNames() failed", err))
	asst.GreaterOrEqual(len(queries), constant.ZeroInt, "test GetByServiceNames() Failed")
}

func TestClickhouseRepo_GetByDBName(t *testing.T) {
	asst := assert.New(t)

	queries, err := testClickhouseRepo.GetByDBName(testServiceName, testDBName)
	asst.Nil(err, common.CombineMessageWithError("test GetByDBName() failed", err))
	asst.GreaterOrEqual(len(queries), constant.ZeroInt, "test GetByDBName() Failed")
}

func TestClickhouseRepo_GetBySQLID(t *testing.T) {
	asst := assert.New(t)

	query, err := testClickhouseRepo.GetBySQLID(testServiceName, testSQLID)
	asst.Nil(err, common.CombineMessageWithError("test GetBySQLID() failed", err))
	asst.NotNil(query, "test GetBySQLID() Failed")
}
