package query

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

var testQuerier *Querier

func init() {
	testQuerier = newQuerier(NewConfigWithDefault(), testDASRepo)
}

func TestQuery_All(t *testing.T) {
	// test PMM1.x
	TestQuery_PMM1(t)
	// test PMM2.x
	TestQuery_PMM2(t)
}

func TestQuery_PMM1(t *testing.T) {
	testPMMVersion = 1
	TestQuery_PMM(t)
}

func TestQuery_PMM2(t *testing.T) {
	testPMMVersion = 2
	TestQuery_PMM(t)
}

func TestQuery_PMM(t *testing.T) {
	testInitMySQLInfo()

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
