package query

import (
	"fmt"
	"os"
	"testing"

	"github.com/romberli/das/config"
	"github.com/romberli/das/global"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	// modify the connection information
	testDBAddr   = "192.168.10.219:3306"
	testDBDBName = "das"
	testDBDBUser = "root"
	testDBDBPass = "root"

	defaultQueryInfoSQLID           = "sql_id"
	defaultQueryInfoFingerprint     = "fingerprint"
	defaultQueryInfoExample         = "example"
	defaultQueryInfoDBName          = "db"
	defaultQueryInfoExecCount       = 1
	defaultQueryInfoTotalExecTime   = 2.1
	defaultQueryInfoAvgExecTime     = 3.2
	defaultQueryInfoRowsExaminedMax = 4
)

var pmmVersion int

func init() {
	viper.Set(config.DBMonitorMySQLUserKey, config.DefaultDBMonitorMySQLUser)
	viper.Set(config.DBMonitorMySQLPassKey, config.DefaultDBMonitorMySQLPass)

	viper.Set(config.DBMonitorClickhouseUserKey, config.DefaultDBMonitorClickhouseUser)
	viper.Set(config.DBMonitorClickhousePassKey, config.DefaultDBMonitorClickhousePass)

	viper.Set(config.QueryMinRowsExaminedKey, 1)

	// pmmVersion = 1
	pmmVersion = 2

	err := initGlobalMySQLPool()
	if err != nil {
		log.Errorf("initGlobalMySQLPool() failed. error:\n%s", err.Error())
		os.Exit(constant.DefaultAbnormalExitCode)
	}
}

func initGlobalMySQLPool() error {
	cfg := mysql.NewConfig(testDBAddr, testDBDBName, testDBDBUser, testDBDBPass)
	poolConfig := mysql.NewPoolConfigWithConfig(cfg, mysql.DefaultMaxConnections, mysql.DefaultInitConnections,
		mysql.DefaultMaxIdleConnections, mysql.DefaultMaxIdleTime, mysql.DefaultKeepAliveInterval)
	log.Debugf("pool config: %v", poolConfig)
	var err error
	global.DASMySQLPool, err = mysql.NewPoolWithPoolConfig(poolConfig)

	return err
}

func TestQueryAll(t *testing.T) {
	TestQuery_GetSQLID(t)
	TestQuery_GetFingerprint(t)
	TestQuery_GetExample(t)
	TestQuery_GetDBName(t)
	TestQuery_GetExecCount(t)
	TestQuery_GetTotalExecTime(t)
	TestQuery_GetAvgExecTime(t)
	TestQuery_GetRowsExaminedMax(t)

	// Test PMM1.x
	pmmVersion = 1
	TestQuerier_GetByMySQLClusterID(t)
	TestQuerier_GetByMySQLServerID(t)
	TestQuerier_GetByDBID(t)
	TestQuerier_GetBySQLID(t)

	// Test PMM2.x
	pmmVersion = 2
	TestQuerier_GetByMySQLClusterID(t)
	TestQuerier_GetByMySQLServerID(t)
	TestQuerier_GetByDBID(t)
	TestQuerier_GetBySQLID(t)
}

func initNewQueryInfo() *Query {
	return &Query{
		defaultQueryInfoSQLID,
		defaultQueryInfoFingerprint,
		defaultQueryInfoExample,
		defaultQueryInfoDBName,
		defaultQueryInfoExecCount,
		defaultQueryInfoTotalExecTime,
		defaultQueryInfoAvgExecTime,
		defaultQueryInfoRowsExaminedMax,
	}
}

func TestQuery_GetSQLID(t *testing.T) {
	asst := assert.New(t)

	queryInfo := initNewQueryInfo()
	asst.Equal(defaultQueryInfoSQLID, queryInfo.GetSQLID(), "test GetUserName() failed")
}
func TestQuery_GetFingerprint(t *testing.T) {
	asst := assert.New(t)

	queryInfo := initNewQueryInfo()
	asst.Equal(defaultQueryInfoFingerprint, queryInfo.GetFingerprint(), "test GetUserName() failed")
}
func TestQuery_GetExample(t *testing.T) {
	asst := assert.New(t)

	queryInfo := initNewQueryInfo()
	asst.Equal(defaultQueryInfoExample, queryInfo.GetExample(), "test GetUserName() failed")
}
func TestQuery_GetDBName(t *testing.T) {
	asst := assert.New(t)

	queryInfo := initNewQueryInfo()
	asst.Equal(defaultQueryInfoDBName, queryInfo.GetDBName(), "test GetUserName() failed")
}
func TestQuery_GetExecCount(t *testing.T) {
	asst := assert.New(t)

	queryInfo := initNewQueryInfo()
	asst.Equal(defaultQueryInfoExecCount, queryInfo.GetExecCount(), "test GetUserName() failed")
}
func TestQuery_GetTotalExecTime(t *testing.T) {
	asst := assert.New(t)

	queryInfo := initNewQueryInfo()
	asst.Equal(defaultQueryInfoTotalExecTime, queryInfo.GetTotalExecTime(), "test GetUserName() failed")
}
func TestQuery_GetAvgExecTime(t *testing.T) {
	asst := assert.New(t)

	queryInfo := initNewQueryInfo()
	asst.Equal(defaultQueryInfoAvgExecTime, queryInfo.GetAvgExecTime(), "test GetUserName() failed")
}
func TestQuery_GetRowsExaminedMax(t *testing.T) {
	asst := assert.New(t)

	queryInfo := initNewQueryInfo()
	asst.Equal(defaultQueryInfoRowsExaminedMax, queryInfo.GetRowsExaminedMax(), "test GetUserName() failed")
}

func TestQuerier_GetByMySQLClusterID(t *testing.T) {
	asst := assert.New(t)

	querier := NewQuerierWithGlobal(NewConfigWithDefault())

	querierMySQLClusterID := 0
	switch pmmVersion {
	case 1:
		querierMySQLClusterID = testPMM1MySQLClusterID
	case 2:
		querierMySQLClusterID = testPMM2MySQLClusterID
	default:
		err := fmt.Errorf("PMM with version:%d is not supported for now", pmmVersion)
		asst.Nil(err, common.CombineMessageWithError("test GetByMySQLClusterID() failed", err))
	}

	queries, err := querier.GetByMySQLClusterID(querierMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetByMySQLClusterID() failed", err))

	asst.NotZero(len(queries), "test GetByMySQLClusterID() failed")
}

func TestQuerier_GetByMySQLServerID(t *testing.T) {
	asst := assert.New(t)

	querierMySQLServerID := 0
	switch pmmVersion {
	case 1:
		querierMySQLServerID = testPMM1MySQLServerID
	case 2:
		querierMySQLServerID = testPMM2MySQLServerID
	default:
		err := fmt.Errorf("PMM with version:%d is not supported for now", pmmVersion)
		asst.Nil(err, common.CombineMessageWithError("test GetByMySQLServerID() failed", err))
	}

	querier := NewQuerierWithGlobal(NewConfigWithDefault())
	queries, err := querier.GetByMySQLServerID(querierMySQLServerID)
	asst.Nil(err, common.CombineMessageWithError("test GetByMySQLServerID() failed", err))

	asst.NotZero(len(queries), "test GetByMySQLServerID() failed")
}

func TestQuerier_GetByDBID(t *testing.T) {
	asst := assert.New(t)

	querierDBID := 0
	querierMySQLServerID := 0
	switch pmmVersion {
	case 1:
		querierMySQLServerID = testPMM1MySQLServerID
		querierDBID = testPMM1DBID
	case 2:
		querierMySQLServerID = testPMM2MySQLServerID
		querierDBID = testPMM2DBID
	default:
		err := fmt.Errorf("PMM with version:%d is not supported for now", pmmVersion)
		asst.Nil(err, common.CombineMessageWithError("test GetByDBID() failed", err))
	}

	querier := NewQuerierWithGlobal(NewConfigWithDefault())
	queries, err := querier.GetByDBID(querierMySQLServerID, querierDBID)
	asst.Nil(err, common.CombineMessageWithError("test GetByDBID() failed", err))

	asst.NotZero(len(queries), "test GetByDBID() failed")
}
func TestQuerier_GetBySQLID(t *testing.T) {
	asst := assert.New(t)

	var querierSQLID string
	querierMySQLServerID := 0
	switch pmmVersion {
	case 1:
		querierMySQLServerID = testPMM1MySQLServerID
		querierSQLID = testPMM1SQLID
	case 2:
		querierMySQLServerID = testPMM2MySQLServerID
		querierSQLID = testPMM2SQLID
	default:
		err := fmt.Errorf("PMM with version:%d is not supported for now", pmmVersion)
		asst.Nil(err, common.CombineMessageWithError("test GetByDBID() failed", err))
	}

	querier := NewQuerierWithGlobal(NewConfigWithDefault())
	queries, err := querier.GetBySQLID(querierMySQLServerID, querierSQLID)
	asst.Nil(err, common.CombineMessageWithError("test GetBySQLID() failed", err))

	asst.NotZero(len(queries), "test GetBySQLID() failed")
}
