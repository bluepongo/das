package sqladvisor

import (
	"os"
	"testing"

	"github.com/romberli/das/global"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
)

const (
	testDASMySQLAddr = "192.168.10.219:3306"
	testDASMySQLName = "das"
	testDASMySQLUser = "root"
	testDASMySQLPass = "root"

	testDBID    = 1
	testSQLText = "select * from t_meta_db_info where create_time<'2021-01-01';"
	testAdvice  = "[\n {\n  \"ID\": \"B95017DB61875675\",\n  \"Fingerprint\": \"select * from t_meta_db_info where create_time\\u003c?\",\n  \"Score\": 95,\n  \"Sample\": \"select * from t_meta_db_info where create_time\\u003c'2021-01-01'\",\n  \"Explain\": null,\n  \"HeuristicRules\": [\n    {\n      \"Item\": \"COL.001\",\n      \"Severity\": \"L1\",\n      \"Summary\": \"不建议使用 SELECT * 类型查询\",\n      \"Content\": \"当表结构变更时，使用 * 通配符选择所有列将导致查询的含义和行为会发生更改，可能导致查询返回更多的数据。\",\n      \"Case\": \"select * from tbl where id=1\",\n      \"Position\": 0\n    }\n  ],\n  \"IndexRules\": null,\n  \"Tables\": [\n    \"`soar`.`t_meta_db_info`\"\n  ]\n}\n]"
	testMessage = ""
)

var testRepo *Repository

func init() {
	testInitDASMySQLPool()
	testRepo = NewRepositoryWithGlobal()
}

func testInitDASMySQLPool() {
	var err error

	global.DASMySQLPool, err = mysql.NewPoolWithDefault(testDASMySQLAddr, testDASMySQLName, testDASMySQLUser, testDASMySQLPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitDASMySQLPool() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
}

func TestRepositoryAll(t *testing.T) {
	TestRepository_Execute(t)
	TestRepository_Save(t)
}

func TestRepository_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := "select 1;"
	result, err := testRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestRepository_Save(t *testing.T) {
	asst := assert.New(t)

	err := testRepo.Save(testDBID, testSQLText, testAdvice, testMessage)
	asst.Nil(err, common.CombineMessageWithError("test Save() failed", err))
}
