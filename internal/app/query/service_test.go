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
	defaultQuerySQLID           = "sql_id"
	defaultQueryFingerprint     = "fingerprint"
	defaultQueryExample         = "example"
	defaultQueryDBName          = "test"
	defaultQueryExecCount       = 1
	defaultQueryTotalExecTime   = 3.5
	defaultQueryAvgExecTime     = 1.5
	defaultQueryRowsExaminedMax = 10

	testPMM1ServiceName    = "192-168-10-220:3306"
	testPMM1MySQLClusterID = 2
	testPMM1MySQLServerID  = 2
	testPMM1DBID           = 3
	testPMM1DBName         = "sbtest"
	testPMM1SQLID          = "8245FBA6EC962AA4"

	testPMM2ServiceName    = "192-168-10-219:3306"
	testPMM2MySQLClusterID = 1
	testPMM2MySQLServerID  = 1
	testPMM2DBID           = 1
	testPMM2DBName         = "pmmtest"
	testPMM2SQLID          = "5418BB5EE2F4A162"
)

var (
	testServiceName    string
	testMySQLClusterID int
	testMySQLServerID  int
	testDBID           int
	testDBName         string
	testSQLID          string
	testService        *Service
)

func init() {
	// pmmVersion = 1
	pmmVersion = 2

	switch pmmVersion {
	case 1:
		testServiceName = testPMM1ServiceName
		testMySQLClusterID = testPMM1MySQLClusterID
		testMySQLServerID = testPMM1MySQLServerID
		testDBID = testPMM1DBID
		testDBName = testPMM1DBName
		testSQLID = testPMM1SQLID
		viper.Set(config.DBMonitorMySQLUserKey, config.DefaultDBMonitorMySQLUser)
		viper.Set(config.DBMonitorMySQLPassKey, config.DefaultDBMonitorMySQLPass)
	case 2:
		testServiceName = testPMM2ServiceName
		testMySQLClusterID = testPMM2MySQLClusterID
		testMySQLServerID = testPMM2MySQLServerID
		testDBID = testPMM2DBID
		testDBName = testPMM2DBName
		testSQLID = testPMM2SQLID
		viper.Set(config.DBMonitorClickhouseUserKey, config.DefaultDBMonitorClickhouseUser)
		viper.Set(config.DBMonitorClickhousePassKey, config.DefaultDBMonitorClickhousePass)
	}

	testService = createService()

}

func initQueryRepo() *DASRepo {
	var err error

	global.DASMySQLPool, err = mysql.NewPoolWithDefault(testDBAddr, testDBDBName, testDBDBUser, testDBDBPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initQueryRepo() failed", err))
		return nil
	}
	return newDASRepo(global.DASMySQLPool)
}

func initNewQuery() *Query {
	return &Query{
		defaultQuerySQLID,
		defaultQueryFingerprint,
		defaultQueryExample,
		defaultQueryDBName,
		defaultQueryExecCount,
		defaultQueryTotalExecTime,
		defaultQueryAvgExecTime,
		defaultQueryRowsExaminedMax,
	}
}

func createService() *Service {
	return newService(NewConfigWithDefault(), initQueryRepo())
}

func TestService_All(t *testing.T) {
	TestService_GetConfig(t)
	TestService_GetQueries(t)
	TestService_GetByMySQLClusterID(t)
	TestService_GetByMySQLServerID(t)
	TestService_GetByDBID(t)
	TestService_GetBySQLID(t)
	TestService_Marshal(t)
}

func TestService_GetConfig(t *testing.T) {
	asst := assert.New(t)

	limit := testService.GetConfig().GetLimit()
	asst.Equal(defaultLimit, limit, "test GetConfig() failed")
}

func TestService_GetQueries(t *testing.T) {
	asst := assert.New(t)

	testService.queries = append(testService.queries, initNewQuery())
	sqlID := testService.GetQueries()[0].GetSQLID()
	asst.Equal(defaultQuerySQLID, sqlID, "test GetQueries() failed")
}

func TestService_GetByMySQLClusterID(t *testing.T) {
	asst := assert.New(t)

	err := testService.GetByMySQLClusterID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetByMySQLClusterID() failed", err))
}

func TestService_GetByMySQLServerID(t *testing.T) {
	asst := assert.New(t)

	err := testService.GetByMySQLServerID(testMySQLServerID)
	asst.Nil(err, common.CombineMessageWithError("test GetByMySQLServerID() failed", err))
}

func TestService_GetByDBID(t *testing.T) {
	asst := assert.New(t)

	err := testService.GetByDBID(testMySQLServerID, testDBID)
	asst.Nil(err, common.CombineMessageWithError("test GetByDBID() failed", err))
}

func TestService_GetBySQLID(t *testing.T) {
	asst := assert.New(t)

	err := testService.GetBySQLID(testMySQLServerID, testSQLID)
	asst.Nil(err, common.CombineMessageWithError("test GetBySQLID() failed", err))
}

func TestService_Marshal(t *testing.T) {
	asst := assert.New(t)

	_, err := testService.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
}
