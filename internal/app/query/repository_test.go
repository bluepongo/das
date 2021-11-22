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
	testDASMySQLAddr = "192.168.10.219:3306"
	testDASMySQLName = "das"
	testDASMySQLUser = "root"
	testDASMySQLPass = "root"
)

var (
	testDASRepo        *DASRepo
	testMySQLRepo      *MySQLRepo
	testClickhouseRepo *ClickhouseRepo
)

func init() {
	testInitDASMySQLPool()
	testInitViper()

	testDASRepo = NewDASRepoWithGlobal()
	testMySQLRepo = testInitMySQLRepo()
	testClickhouseRepo = testInitClickhouseRepo()
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
	viper.Set(config.QueryMinRowsExaminedKey, 1)
}

func testInitMySQLRepo() *MySQLRepo {
	mysqlServerService := metadata.NewMySQLServerServiceWithDefault()
	err := mysqlServerService.GetByID(testPMM1MySQLServerID)
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
	err := mysqlServerService.GetByID(testPMM2MySQLServerID)
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

func TestQueryRepositoryAll(t *testing.T) {
	TestDASRepo_Save(t)
	TestMySQLRepo_GetByServiceNames(t)
	TestMySQLRepo_GetByDBName(t)
	TestMySQLRepo_GetBySQLID(t)
	TestClickhouseRepo_GetByServiceNames(t)
	TestClickhouseRepo_GetByDBName(t)
	TestClickhouseRepo_GetBySQLID(t)
}

func TestDASRepo_Save(t *testing.T) {
	asst := assert.New(t)

	err := testDASRepo.Save(
		testPMM2MySQLClusterID,
		testPMM2MySQLServerID,
		testPMM2DBID,
		testPMM2SQLID,
		time.Now().Add(-constant.Week),
		time.Now(),
		defaultLimit,
		defaultOffset,
	)
	asst.Nil(err, common.CombineMessageWithError("test Save() failed", err))
}

func TestMySQLRepo_GetByServiceNames(t *testing.T) {
	asst := assert.New(t)

	queries, err := testMySQLRepo.GetByServiceNames([]string{testPMM1ServiceName})
	asst.Nil(err, common.CombineMessageWithError("test GetByServiceNames() failed", err))
	asst.GreaterOrEqual(len(queries), constant.ZeroInt, "test TestMySQLRepo_GetByServiceNames() failed")
}

func TestMySQLRepo_GetByDBName(t *testing.T) {
	asst := assert.New(t)

	queries, err := testMySQLRepo.GetByDBName(testPMM1ServiceName, testPMM1DBName)
	asst.Nil(err, common.CombineMessageWithError("test GetByDBName() failed", err))
	asst.GreaterOrEqual(len(queries), constant.ZeroInt, "test GetByDBName() failed")
}

func TestMySQLRepo_GetBySQLID(t *testing.T) {
	asst := assert.New(t)

	query, err := testMySQLRepo.GetBySQLID(testPMM1ServiceName, testPMM1SQLID)
	asst.Nil(err, common.CombineMessageWithError("test GetBySQLID() failed", err))
	asst.NotNil(query, "test GetBySQLID() failed")
}

func TestClickhouseRepo_GetByServiceNames(t *testing.T) {
	asst := assert.New(t)

	queries, err := testClickhouseRepo.GetByServiceNames([]string{testPMM2ServiceName})
	asst.Nil(err, common.CombineMessageWithError("test GetByServiceNames() failed", err))
	asst.GreaterOrEqual(len(queries), constant.ZeroInt, "test GetByServiceNames() Failed")
}

func TestClickhouseRepo_GetByDBName(t *testing.T) {
	asst := assert.New(t)

	queries, err := testClickhouseRepo.GetByDBName(testPMM2ServiceName, testPMM2DBName)
	asst.Nil(err, common.CombineMessageWithError("test GetByDBName() failed", err))
	asst.GreaterOrEqual(len(queries), constant.ZeroInt, "test GetByDBName() Failed")
}

func TestClickhouseRepo_GetBySQLID(t *testing.T) {
	asst := assert.New(t)

	query, err := testClickhouseRepo.GetBySQLID(testPMM2ServiceName, testPMM2SQLID)
	asst.Nil(err, common.CombineMessageWithError("test GetBySQLID() failed", err))
	asst.NotNil(query, "test GetBySQLID() Failed")
}
