package metadata

import (
	"testing"

	"github.com/jinzhu/now"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	testTableStatisticDBName           = "das"
	testTableStatisticTableName        = "t_meta_db_info"
	testTableStatisticTableRows        = 1
	testTableStatisticSize             = 1
	testTableStatisticSizeMB           = 1
	testTableStatisticAvgRowLength     = 1
	testTableStatisticAutoIncrement    = 1
	testTableStatisticEngine           = "InnoDB"
	testTableStatisticCharSet          = "utf8mb4"
	testTableStatisticCollation        = "utf8mb4_general_ci"
	testTableStatisticCreateTimeString = "2021-01-21 10:00:00.000000"

	testIndexStatisticDBName      = "das"
	testIndexStatisticTableName   = "t_meta_db_info"
	testIndexStatisticIndexName   = "idx01_db_name_cluster_id_cluster_type_env_id"
	testIndexStatisticSequence    = 1
	testIndexStatisticColumnName  = "db_name"
	testIndexStatisticCardinality = 1
	testIndexStatisticIsUnique    = true
	testIndexStatisticIsNullable  = false
)

var (
	testTableInfo      *TableInfo
	testTableStatistic *TableStatistic
	testIndexStatistic *IndexStatistic
)

func init() {
	testInitDASMySQLPool()
	testTableInitViper()

	testTableRepo = testInitTableRepo()
	testTableStatistic = testInitTableStatistic()
	testIndexStatistic = testInitIndexStatistic()
	testTableInfo = testInitTableInfo()
}

func testInitTableStatistic() *TableStatistic {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)
	createTime, _ := now.Parse(testTableStatisticCreateTimeString)

	return &TableStatistic{
		testTableStatisticDBName,
		testTableStatisticTableName,
		testTableStatisticTableRows,
		testTableStatisticSize,
		testTableStatisticSizeMB,
		testTableStatisticAvgRowLength,
		testTableStatisticAutoIncrement,
		testTableStatisticEngine,
		testTableStatisticCharSet,
		testTableStatisticCollation,
		createTime,
	}
}

func testInitIndexStatistic() *IndexStatistic {
	return &IndexStatistic{
		testIndexStatisticDBName,
		testIndexStatisticTableName,
		testIndexStatisticIndexName,
		testIndexStatisticSequence,
		testIndexStatisticColumnName,
		testIndexStatisticCardinality,
		testIndexStatisticIsUnique,
		testIndexStatisticIsNullable,
	}
}

func testInitTableInfo() *TableInfo {
	return &TableInfo{
		testTableRepo,
		testTableDBName,
		testTableTableName,
	}
}

func TestTableEntityAll(t *testing.T) {
	TestTableStatistic_Get(t)
	TestTableStatistic_MarshalJSON(t)

	TestIndexStatistic_Get(t)
	TestIndexStatistic_MarshalJSON(t)

	TestTableInfo_Get(t)
	TestTableInfo_GetTableStatistics(t)
	TestTableInfo_GetIndexStatistics(t)
	TestTableInfo_GetCreateStatement(t)
	TestTableInfo_MarshalJSON(t)
	TestTableInfo_MarshalJSONWithFields(t)
}

func TestTableStatistic_Get(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testTableStatistic.DBName, testTableStatistic.GetDBName(), "test GetDBName() failed")
	asst.Equal(testTableStatistic.TableName, testTableStatistic.GetTableName(), "test GetTableName() failed")
	asst.Equal(testTableStatistic.Rows, testTableStatistic.GetTableRows(), "test GetTableRows() failed")
	asst.Equal(testTableStatistic.Size, testTableStatistic.GetSize(), "test GetSize() failed")
	asst.Equal(testTableStatistic.SizeMB, testTableStatistic.GetSizeMB(), "test GetSizeMB() failed")
	asst.Equal(testTableStatistic.AvgRowLength, testTableStatistic.GetAvgRowLength(), "test GetAvgRowLength() failed")
	asst.Equal(testTableStatistic.AutoIncrement, testTableStatistic.GetAutoIncrement(), "test GetAutoIncrement() failed")
	asst.Equal(testTableStatistic.Engine, testTableStatistic.GetEngine(), "test GetEngine() failed")
	asst.Equal(testTableStatistic.CharSet, testTableStatistic.GetCharSet(), "test GetCharSet() failed")
	asst.Equal(testTableStatistic.Collation, testTableStatistic.GetCollation(), "test GetCollation() failed")
	asst.Equal(testTableStatistic.CreateTime, testTableStatistic.GetCreateTime(), "test GetCreateTime() failed")
}

func TestTableStatistic_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testTableStatistic.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	t.Log(string(jsonBytes))
}

func TestIndexStatistic_Get(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testIndexStatistic.DBName, testIndexStatistic.GetDBName(), "test GetDBName() failed")
	asst.Equal(testIndexStatistic.TableName, testIndexStatistic.GetTableName(), "test GetTableName() failed")
	asst.Equal(testIndexStatistic.IndexName, testIndexStatistic.GetIndexName(), "test GetIndexName() failed")
	asst.Equal(testIndexStatistic.Sequence, testIndexStatistic.GetSequence(), "test GetSequence() failed")
	asst.Equal(testIndexStatistic.ColumnName, testIndexStatistic.GetColumnName(), "test GetColumnName() failed")
	asst.Equal(testIndexStatistic.Cardinality, testIndexStatistic.GetCardinality(), "test GetCardinality() failed")
	asst.Equal(testIndexStatistic.Unique, testIndexStatistic.IsUnique(), "test IsUnique() failed")
	asst.Equal(testIndexStatistic.Nullable, testIndexStatistic.IsNullable(), "test IsNullable() failed")
}

func TestIndexStatistic_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testIndexStatistic.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	t.Log(string(jsonBytes))
}

func TestTableInfo_Get(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testTableInfo.DBName, testTableInfo.GetDBName(), "test GetDBName() failed")
	asst.Equal(testTableInfo.TableName, testTableInfo.GetTableName(), "test GetTableName() failed")
}

func TestTableInfo_GetTableStatistics(t *testing.T) {
	asst := assert.New(t)

	tableStatistics, err := testTableInfo.GetTableStatistics()
	asst.Nil(err, common.CombineMessageWithError("test GetTableStatistics() failed", err))
	asst.Equal(testTableTableName, tableStatistics[constant.ZeroInt].GetTableName(), "test GetTableStatistics() failed")
}

func TestTableInfo_GetIndexStatistics(t *testing.T) {
	asst := assert.New(t)

	indexStatistic, err := testTableInfo.GetIndexStatistics()
	asst.Nil(err, common.CombineMessageWithError("test GetIndexStatistics() failed", err))
	asst.Equal(testTableTableName, indexStatistic[constant.ZeroInt].GetTableName(), "test GetIndexStatistics() failed")
}

func TestTableInfo_GetCreateStatement(t *testing.T) {
	asst := assert.New(t)

	createStatement, err := testTableInfo.GetCreateStatement()
	asst.Nil(err, common.CombineMessageWithError("test GetCreateStatement() failed", err))
	asst.NotEqual(constant.EmptyString, createStatement, "test GetCreateStatement() failed")
}

func TestTableInfo_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testTableInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	t.Log(string(jsonBytes))
}

func TestTableInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testTableInfo.MarshalJSONWithFields(dbNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	t.Log(string(jsonBytes))
}
