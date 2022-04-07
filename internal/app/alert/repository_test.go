package alert

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
	testDASMySQLAddr = "192.168.137.11:3306"
	testDASMySQLName = "das"
	testDASMySQLUser = "root"
	testDASMySQLPass = "root"

	testToAddrs = "allinemailtest@163.com"
	testCCAddrs = "allinemailtest@163.com"
	testSubject = "test subject"
	testContent = "test content"

	testMessage = "test-message"
	testConfig  = `{"pass": "****"}`
)

var testRepo *Repository

func init() {
	testInitDASMySQLPool()
	testInitViper()
	testRepo = newRepository(global.DASMySQLPool)
}

func testInitDASMySQLPool() {
	var err error

	if global.DASMySQLPool == nil {
		global.DASMySQLPool, err = mysql.NewPoolWithDefault(testDASMySQLAddr, testDASMySQLName, testDASMySQLUser, testDASMySQLPass)
		if err != nil {
			log.Error(common.CombineMessageWithError("testInitDASMySQLPool() failed", err))
			os.Exit(constant.DefaultAbnormalExitCode)
		}
	}
}

func TestRepo_All(t *testing.T) {
	TestRepository_Execute(t)
	TestRepository_Save(t)
}

func TestRepository_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1`
	_, err := testRepo.Execute(sql)

	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))

}

func TestRepository_Save(t *testing.T) {
	asst := assert.New(t)

	err := testRepo.Save(testSMTPURL, testToAddrs, testCCAddrs, testSubject, testContent, testConfig, testMessage)
	asst.Nil(err, common.CombineMessageWithError("test Save() failed", err))
}
