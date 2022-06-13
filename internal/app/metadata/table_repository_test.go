package metadata

import (
	"os"
	"testing"

	"github.com/romberli/das/config"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	testTableHostIP    = "192.168.137.11"
	testTablePortNum   = 3306
	testTableDBName    = "das"
	testTableTableName = "t_meta_db_info"
	testTableLoginName = "zhangs"
)

var testTableRepo *TableRepo

func init() {
	testInitDASMySQLPool()
	testTableInitViper()

	testTableRepo = testInitTableRepo()
}

func testTableInitViper() {
	viper.Set(config.DBApplicationMySQLUserKey, config.DefaultDBApplicationMySQLUser)
	viper.Set(config.DBApplicationMySQLPassKey, config.DefaultDBApplicationMySQLPass)
}
func testInitTableRepo() *TableRepo {
	conn, err := mysql.NewConn(testDASMySQLAddr, constant.EmptyString, "root", "root")
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitTableRepo() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}

	return newTableRepo(conn)
}

func TestTableRepoAll(t *testing.T) {
	TestTableRepo_Close(t)
	TestTableRepo_InitMySQLConn(t)
	TestTableRepo_Execute(t)
	TestTableRepo_GetByDBName(t)
	TestTableRepo_GetTableStatistics(t)
	TestTableRepo_GetIndexStatistics(t)
	TestTableRepo_GetCreateStatement(t)
	TestTableRepo_GetStatisticsByDBNameAndTableName(t)
	TestTableRepo_AnalyzeTableByDBNameAndTableName(t)
}

func TestTableRepo_Close(t *testing.T) {
	asst := assert.New(t)

	err := testTableRepo.Close()
	asst.Nil(err, common.CombineMessageWithError("test Close() failed", err))
	err = testTableRepo.InitMySQLConn(testTableHostIP, testTablePortNum, testTableDBName)
	asst.Nil(err, common.CombineMessageWithError("test InitMySQLConn() failed", err))
}

func TestTableRepo_InitMySQLConn(t *testing.T) {
	asst := assert.New(t)

	err := testTableRepo.Close()
	asst.Nil(err, common.CombineMessageWithError("test InitMySQLConn() failed", err))
	testTableRepo.Conn = nil
	err = testTableRepo.InitMySQLConn(testTableHostIP, testTablePortNum, testTableDBName)
	asst.Nil(err, common.CombineMessageWithError("test InitMySQLConn() failed", err))
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

func TestTableRepo_GetByDBName(t *testing.T) {
	asst := assert.New(t)

	tables, err := testTableRepo.GetByDBName(testTableDBName)
	asst.Nil(err, common.CombineMessageWithError("test GetByDBName() failed", err))
	asst.NotNil(tables, "test GetByDBName() failed")
}

func TestTableRepo_GetTableStatistics(t *testing.T) {
	asst := assert.New(t)

	tableStatistics, err := testTableRepo.GetTableStatistics(testTableDBName, testTableTableName)
	asst.Nil(err, common.CombineMessageWithError("test GetTableStatistics() failed", err))
	asst.NotNil(tableStatistics, "test GetTableStatistics() failed")
}

func TestTableRepo_GetIndexStatistics(t *testing.T) {
	asst := assert.New(t)

	indexStatistics, err := testTableRepo.GetIndexStatistics(testTableDBName, testTableTableName)
	asst.Nil(err, common.CombineMessageWithError("test GetIndexStatistics() failed", err))
	asst.NotNil(indexStatistics, "test GetIndexStatistics() failed")
}

func TestTableRepo_GetCreateStatement(t *testing.T) {
	asst := assert.New(t)

	createStatement, err := testTableRepo.GetCreateStatement(testTableDBName, testTableTableName)
	asst.Nil(err, common.CombineMessageWithError("test GetCreateStatement() failed", err))
	asst.NotEqual(constant.EmptyString, createStatement, "test GetCreateStatement() failed")
}

func TestTableRepo_GetStatisticsByDBNameAndTableName(t *testing.T) {
	asst := assert.New(t)

	tableStatistics, indexStatistics, createStatement, err := testTableRepo.GetStatisticsByDBNameAndTableName(testTableDBName, testTableTableName)
	asst.Nil(err, common.CombineMessageWithError("test GetStatisticsByDBNameAndTableName() failed", err))
	asst.NotNil(tableStatistics, "test GetStatisticsByDBNameAndTableName() failed")
	asst.NotNil(indexStatistics, "test GetStatisticsByDBNameAndTableName() failed")
	asst.NotEqual(constant.EmptyString, createStatement, "test GetStatisticsByDBNameAndTableName() failed")

}

func TestTableRepo_AnalyzeTableByDBNameAndTableName(t *testing.T) {
	asst := assert.New(t)

	err := testTableRepo.AnalyzeTableByDBNameAndTableName(testTableDBName, testTableTableName)
	asst.Nil(err, common.CombineMessageWithError("test AnalyzeTableByDBNameAndTableName() failed", err))
}
