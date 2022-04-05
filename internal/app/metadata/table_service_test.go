package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const testTableDBID = 1

var testTableService *TableService

func init() {
	testInitDASMySQLPool()
	testTableInitViper()

	testTableService = newTableService(NewTableRepoWithDefault())
}

func TestTableServiceAll(t *testing.T) {
	TestTableService_GetTables(t)
	TestTableService_GetByDBID(t)
	TestTableService_GetStatisticsByDBIDAndTableName(t)
	TestTableService_GetStatisticsByHostInfoAndDBNameAndTableName(t)
	TestTableService_AnalyzeTableByDBIDAndTableName(t)
	TestTableService_AnalyzeTableByHostInfoAndDBNameAndTableName(t)
	TestTableService_Marshal(t)
	TestTableService_MarshalWithFields(t)
}

func TestTableService_GetTables(t *testing.T) {
	asst := assert.New(t)

	tables := testTableService.GetTables()
	asst.Equal(constant.ZeroInt, len(tables), "test GetTables() failed")
}

func TestTableService_GetByDBID(t *testing.T) {
	asst := assert.New(t)

	err := testTableService.GetByDBID(testTableDBID, testTableLoginName)
	asst.Nil(err, common.CombineMessageWithError("test GetByDBID() failed", err))
	asst.NotNil(testTableService.GetTables(), "test GetByDBID() failed")
}

func TestTableService_GetStatisticsByDBIDAndTableName(t *testing.T) {
	asst := assert.New(t)

	err := testTableService.GetStatisticsByDBIDAndTableName(testTableDBID, testTableTableName, testTableLoginName)
	asst.Nil(err, common.CombineMessageWithError("test GetStatisticsByDBIDAndTableName() failed", err))
	asst.NotNil(testTableService.GetTableStatistics(), "test GetStatisticsByDBIDAndTableName() failed")
	asst.NotNil(testTableService.GetIndexStatistics(), "test GetStatisticsByDBIDAndTableName() failed")
	asst.NotEqual(constant.EmptyString, testTableService.GetCreateStatement(), "test GetStatisticsByDBIDAndTableName() failed")
}

func TestTableService_GetStatisticsByHostInfoAndDBNameAndTableName(t *testing.T) {
	asst := assert.New(t)

	err := testTableService.GetStatisticsByHostInfoAndDBNameAndTableName(testTableHostIP, testTablePortNum, testTableDBName, testTableTableName, testTableLoginName)
	asst.Nil(err, common.CombineMessageWithError("test GetStatisticsByHostInfoAndDBNameAndTableName() failed", err))
	asst.NotNil(testTableService.GetTableStatistics(), "test GetStatisticsByHostInfoAndDBNameAndTableName() failed")
	asst.NotNil(testTableService.GetIndexStatistics(), "test GetStatisticsByHostInfoAndDBNameAndTableName() failed")
	asst.NotEqual(constant.EmptyString, testTableService.GetCreateStatement(), "test GetStatisticsByHostInfoAndDBNameAndTableName() failed")
}

func TestTableService_AnalyzeTableByDBIDAndTableName(t *testing.T) {
	asst := assert.New(t)

	err := testTableService.AnalyzeTableByDBIDAndTableName(testTableDBID, testTableTableName, testTableLoginName)
	asst.Nil(err, common.CombineMessageWithError("test AnalyzeTableByDBIDAndTableName() failed", err))
}

func TestTableService_AnalyzeTableByHostInfoAndDBNameAndTableName(t *testing.T) {
	asst := assert.New(t)

	err := testTableService.AnalyzeTableByHostInfoAndDBNameAndTableName(testTableHostIP, testTablePortNum, testTableDBName, testTableTableName, testTableLoginName)
	asst.Nil(err, common.CombineMessageWithError("test AnalyzeTableByHostInfoAndDBNameAndTableName() failed", err))
}

func TestTableService_Marshal(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testTableService.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}

func TestTableService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testTableService.MarshalWithFields(mysqlClusterMySQLClustersStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	t.Log(string(jsonBytes))
}
