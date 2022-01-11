package query

import (
	"fmt"
	"os"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
)

const (
	testPMM1ServiceName    = "192-168-10-220:3306"
	testPMM1MySQLClusterID = 2
	testPMM1MySQLServerID  = 2
	testPMM1DBID           = 3
	testPMM1DBName         = "sbtest"
	testPMM1SQLID          = "999ECD050D719733"

	testPMM2ServiceName    = "192-168-10-219:3306"
	testPMM2MySQLClusterID = 1
	testPMM2MySQLServerID  = 1
	testPMM2DBID           = 1
	testPMM2DBName         = "pmm_test"
	testPMM2SQLID          = "880A8A3142C8BC24"
)

var (
	testPMMVersion     int
	testMySQLClusterID int
	testMySQLServerID  int
	testDBID           int
	testSQLID          string

	testQuerier *Querier
)

func init() {
	testPMMVersion = 2
	testInitMySQLInfo()
	testInitDASMySQLPool()
	testInitViper()
	testQuerier = newQuerier(NewConfigWithDefault(), testDASRepo)
}

func testInitMySQLInfo() {
	switch testPMMVersion {
	case 1:
		testMySQLClusterID = testPMM1MySQLClusterID
		testMySQLServerID = testPMM1MySQLServerID
		testDBID = testPMM1DBID
		testSQLID = testPMM1SQLID
	case 2:
		testMySQLClusterID = testPMM2MySQLClusterID
		testMySQLServerID = testPMM2MySQLServerID
		testDBID = testPMM2DBID
		testSQLID = testPMM2SQLID
	default:
		log.Errorf(fmt.Sprintf("pmm version should be 1 or 2, %d is not valid", testPMMVersion))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
}

func TestQueryAll(t *testing.T) {
	// Test PMM1.x
	testPMMVersion = 1
	TestQuerier_GetByMySQLClusterID(t)
	TestQuerier_GetByMySQLServerID(t)
	TestQuerier_GetByDBID(t)
	TestQuerier_GetBySQLID(t)

	// Test PMM2.x
	testPMMVersion = 2
	TestQuerier_GetByMySQLClusterID(t)
	TestQuerier_GetByMySQLServerID(t)
	TestQuerier_GetByDBID(t)
	TestQuerier_GetBySQLID(t)
}

func TestQuerier_GetByMySQLClusterID(t *testing.T) {
	asst := assert.New(t)

	queries, err := testQuerier.GetByMySQLClusterID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetByMySQLClusterID() failed", err))
	asst.GreaterOrEqual(len(queries), constant.ZeroInt, "test GetByMySQLClusterID() failed")
}

func TestQuerier_GetByMySQLServerID(t *testing.T) {
	asst := assert.New(t)

	queries, err := testQuerier.GetByMySQLServerID(testMySQLServerID)
	asst.Nil(err, common.CombineMessageWithError("test GetByMySQLServerID() failed", err))
	asst.GreaterOrEqual(len(queries), constant.ZeroInt, "test GetByMySQLServerID() failed")
}

func TestQuerier_GetByDBID(t *testing.T) {
	asst := assert.New(t)

	queries, err := testQuerier.GetByDBID(testMySQLServerID, testDBID)
	asst.Nil(err, common.CombineMessageWithError("test GetByDBID() failed", err))
	asst.GreaterOrEqual(len(queries), constant.ZeroInt, "test GetByDBID() failed")
}

func TestQuerier_GetBySQLID(t *testing.T) {
	asst := assert.New(t)

	queries, err := testQuerier.GetBySQLID(testMySQLServerID, testSQLID)
	asst.Nil(err, common.CombineMessageWithError("test GetBySQLID() failed", err))
	asst.GreaterOrEqual(len(queries), constant.ZeroInt, "test GetBySQLID() failed")
}
