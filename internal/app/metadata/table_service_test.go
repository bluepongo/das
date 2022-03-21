package metadata

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/stretchr/testify/assert"
)

var testTableService *TableService

func init() {
	initTableRepo()
	testTableService = NewTableService(testTableRepo)
}

func TestTableServiceAll(t *testing.T) {
	TestTableService_GetTables(t)
	TestTableService_GetByDBName(t)
	TestTableService_GetStatisticsByDBNameAndTableName(t)
	TestTableService_AnalyzeTableByDBIDAndTableName(t)
	TestTableService_AnalyzeTableByHostInfoAndDBNameAndTableName(t)
	TestTableService_Marshal(t)
	TestTableService_MarshalWithFields(t)
}
func TestTableService_GetTables(t *testing.T) {
	asst := assert.New(t)

	tables := testTableService.GetTables()
	asst.Equal(0, len(tables), "test GetTables() failed")
}
func TestTableService_GetByDBName(t *testing.T) {
	asst := assert.New(t)

	err := testTableService.GetByDBName(testTableSchema)
	asst.Nil(err, common.CombineMessageWithError("test GetByDBName() failed", err))
	asst.Equal(31, len(testTableService.GetTables()), "test GetByDBName() failed")
}
func TestTableService_GetStatisticsByDBNameAndTableName(t *testing.T) {
	asst := assert.New(t)

	err := testTableService.GetStatisticsByDBNameAndTableName(testTableSchema, testTableName)
	asst.Nil(err, common.CombineMessageWithError("test GetStatisticsByDBNameAndTableName() failed", err))
	asst.Equal(31, len(testTableService.GetTables()), "test GetStatisticsByDBNameAndTableName() failed")
}
func TestTableService_AnalyzeTableByDBIDAndTableName(t *testing.T) {
	// TODO: compelete service test AnalyzeTableByDBIDAndTableName
}
func TestTableService_AnalyzeTableByHostInfoAndDBNameAndTableName(t *testing.T) {
	// TODO: compelete repo test AnalyzeTableByHostInfoAndDBNameAndTableName
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
