package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
)

func TestQueryAll(t *testing.T) {
	TestQuery_GetSQLID(t)
	TestQuery_GetFingerprint(t)
	TestQuery_GetExample(t)
	TestQuery_GetDBName(t)
	TestQuery_GetExecCount(t)
	TestQuery_GetTotalExecTime(t)
	TestQuery_GetAvgExecTime(t)
	TestQuery_GetRowsExaminedMax(t)

	// TestQuerier_GetByMySQLClusterID(t)
	// TestQuerier_GetByMySQLServerID(t)
	// TestQuerier_GetByDBID(t)
	// TestQuerier_GetBySQLID(t)
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

// func TestQuerier_GetByMySQLClusterID(t *testing.T) {
// 	asst := assert.New(t)
// 	querier := NewQuerierWithGlobal(NewConfigWithDefault())
// 	queries, err := querier.GetByMySQLClusterID()
// 	asst.Nil(err, common.CombineMessageWithError("test GetByMySQLClusterID() failed", err))

// 	asst.Equal()
// }
// func TestQuerier_GetByMySQLServerID(t *testing.T) {
// 	asst := assert.New(t)
// 	querier := NewQuerierWithGlobal(NewConfigWithDefault())
// 	queries, err := querier.GetByMySQLServerID()
// 	asst.Nil(err, common.CombineMessageWithError("test GetByMySQLServerID() failed", err))

// 	asst.Equal()
// }
// func TestQuerier_GetByDBID(t *testing.T) {
// 	asst := assert.New(t)
// 	querier := NewQuerierWithGlobal(NewConfigWithDefault())
// 	queries, err := querier.GetByDBID()
// 	asst.Nil(err, common.CombineMessageWithError("test GetByDBID() failed", err))

// 	asst.Equal()
// }
// func TestQuerier_GetBySQLID(t *testing.T) {
// 	asst := assert.New(t)
// 	querier := NewQuerierWithGlobal(NewConfigWithDefault())
// 	queries, err := querier.GetBySQLID()
// 	asst.Nil(err, common.CombineMessageWithError("test GetBySQLID() failed", err))

// 	asst.Equal()
// }

// func createQuery() {
// 	querier := NewQuerierWithGlobal(NewConfigWithDefault())
// }
