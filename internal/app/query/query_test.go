package query

import (
	"testing"

	"github.com/romberli/das/config"
	"github.com/romberli/das/global"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	// modify the connection information
	queryTestDBAddr   = "192.168.10.220:3306"
	queryTestDBDBName = "das"
	queryTestDBDBUser = "root"
	queryTestDBDBPass = "root"
)

const (
	defaultQueryInfoSQLID           = "sql_id"
	defaultQueryInfoFingerprint     = "fingerprint"
	defaultQueryInfoExample         = "example"
	defaultQueryInfoDBName          = "db"
	defaultQueryInfoExecCount       = 1
	defaultQueryInfoTotalExecTime   = 2.1
	defaultQueryInfoAvgExecTime     = 3.2
	defaultQueryInfoRowsExaminedMax = 4

	defaultQuerierMySQLClusterID = 1
	defaultQuerierMySQLServerID  = 2
	defaultQuerierDBID           = 1
	defaultQuerierSQLID          = "1"
)

func init() {
	viper.Set(config.DBMonitorMySQLUserKey, config.DefaultDBMonitorMySQLUser)
	viper.Set(config.DBMonitorMySQLPassKey, config.DefaultDBMonitorMySQLPass)

	// viper.Set(config.DBMonitorClickHouseUserKey, config.DefaultDBMonitorClickHouseUser)
	// viper.Set(config.DBMonitorClickHousePassKey, config.DefaultDBMonitorClickHousePass)

	if err := initGlobalMySQLPool(); err != nil {
		panic(err)
	}
}

func initGlobalMySQLPool() error {
	dbAddr := queryTestDBAddr
	dbName := queryTestDBDBName
	dbUser := queryTestDBDBUser
	dbPass := queryTestDBDBPass
	maxConnections := mysql.DefaultMaxConnections
	initConnections := mysql.DefaultInitConnections
	maxIdleConnections := mysql.DefaultMaxIdleConnections
	maxIdleTime := mysql.DefaultMaxIdleTime
	keepAliveInterval := mysql.DefaultKeepAliveInterval

	config := mysql.NewConfig(dbAddr, dbName, dbUser, dbPass)
	poolConfig := mysql.NewPoolConfigWithConfig(config, maxConnections, initConnections, maxIdleConnections, maxIdleTime, keepAliveInterval)
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
	queries, err := querier.GetByMySQLClusterID(defaultQuerierMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetByMySQLClusterID() failed", err))

	asst.NotZero(len(queries), "test GetByMySQLClusterID() failed")
}
func TestQuerier_GetByMySQLServerID(t *testing.T) {
	asst := assert.New(t)

	querier := NewQuerierWithGlobal(NewConfigWithDefault())
	queries, err := querier.GetByMySQLServerID(defaultQuerierMySQLServerID)
	asst.Nil(err, common.CombineMessageWithError("test GetByMySQLServerID() failed", err))

	asst.NotZero(len(queries), "test GetByMySQLServerID() failed")
}
func TestQuerier_GetByDBID(t *testing.T) {
	asst := assert.New(t)

	querier := NewQuerierWithGlobal(NewConfigWithDefault())
	queries, err := querier.GetByDBID(defaultQuerierMySQLServerID, defaultQuerierDBID)
	asst.Nil(err, common.CombineMessageWithError("test GetByDBID() failed", err))

	asst.NotZero(len(queries), "test GetByDBID() failed")
}
func TestQuerier_GetBySQLID(t *testing.T) {
	asst := assert.New(t)

	querier := NewQuerierWithGlobal(NewConfigWithDefault())
	queries, err := querier.GetBySQLID(defaultQuerierMySQLServerID, defaultQuerierSQLID)
	asst.Nil(err, common.CombineMessageWithError("test GetBySQLID() failed", err))

	asst.NotZero(len(queries), "test GetBySQLID() failed")
}
