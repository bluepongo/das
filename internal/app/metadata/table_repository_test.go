package metadata

import (
	"os"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
)

const (
	testTableSchema = "mysql"
	testTableName   = "db"
)

var testTableRepo *TableRepo

func init() {
	initTableRepo()
}

func initTableRepo() {
	conn, err := mysql.NewConn(testDASMySQLAddr, constant.EmptyString, testDASMySQLUser, testDASMySQLPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitDASMySQLPool() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
	testTableRepo = NewTableRepo(conn)
}

func TestTableRepoAll(t *testing.T) {
	TestTableRepo_Execute(t)
	TestTableRepo_GetTableStatistics(t)
	TestTableRepo_GetIndexStatistics(t)
	TestTableRepo_GetCreateStatement(t)
	TestTableRepo_GetByDBName(t)
	TestTableRepo_GetStatisticsByDBNameAndTableName(t)
	TestTableRepo_AnalyzeTableByDBIDAndTableName(t)
	TestTableRepo_AnalyzeTableByHostInfoAndDBNameAndTableName(t)
}

func TestTableRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := testTableRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestTableRepo_GetTableStatistics(t *testing.T) {
	asst := assert.New(t)

	result, err := testTableRepo.GetTableStatistics(testTableSchema, testTableName)
	asst.Nil(err, common.CombineMessageWithError("test GetTableStatistics() failed", err))
	asst.Equal(1, len(result), "test GetTableStatistics() failed")
}

func TestTableRepo_GetIndexStatistics(t *testing.T) {
	asst := assert.New(t)

	result, err := testTableRepo.GetIndexStatistics(testTableSchema, testTableName)
	asst.Nil(err, common.CombineMessageWithError("test GetIndexStatistics() failed", err))
	asst.Equal(4, len(result), "test GetIndexStatistics() failed")
}

func TestTableRepo_GetCreateStatement(t *testing.T) {
	asst := assert.New(t)

	result, err := testTableRepo.GetCreateStatement(testTableSchema, testTableName)
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.NotEqual("", result, "test GetAll() failed")
}

func TestTableRepo_GetByDBName(t *testing.T) {
	asst := assert.New(t)

	result, err := testTableRepo.GetByDBName(testTableSchema)
	asst.Nil(err, common.CombineMessageWithError("test GetByDBName() failed", err))
	asst.NotEqual(0, len(result), "test GetByDBName() failed")
}

func TestTableRepo_GetStatisticsByDBNameAndTableName(t *testing.T) {
	asst := assert.New(t)

	resultTableStatistics, resultIndexStatistics, resultCreateStatement, err := testTableRepo.GetStatisticsByDBNameAndTableName(testTableSchema, testTableName)
	asst.Nil(err, common.CombineMessageWithError("test GetStatisticsByDBNameAndTableName() failed", err))
	asst.NotEqual(0, len(resultTableStatistics), "test GetStatisticsByDBNameAndTableName() failed")
	asst.NotEqual(0, len(resultIndexStatistics), "test GetStatisticsByDBNameAndTableName() failed")
	asst.NotEqual(0, len(resultCreateStatement), "test GetStatisticsByDBNameAndTableName() failed")

}

func TestTableRepo_AnalyzeTableByDBIDAndTableName(t *testing.T) {
	// TODO: compelete repo test AnalyzeTableByDBIDAndTableName
}
func TestTableRepo_AnalyzeTableByHostInfoAndDBNameAndTableName(t *testing.T) {
	// TODO: compelete repo test AnalyzeTableByHostInfoAndDBNameAndTableName
}
