package privilege

import (
	"os"
	"testing"

	"github.com/romberli/das/config"
	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/privilege"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	testDASMySQLAddr = "192.168.137.11:3306"
	testDASMySQLName = "das"
	testDASMySQLUser = "root"
	testDASMySQLPass = "root"

	testLoginName      = "zhangs"
	testMySQLClusterID = 1
	testMySQLServerID  = 1
	testMySQLHostIP    = "192.168.137.11"
	testMYSQLPortNum   = 3306
	testDBID           = 1
)

var testRepository privilege.Repository

func init() {
	testInitDASMySQLPool()
	testInitViper()

	testRepository = NewRepositoryWithGlobal()
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

func testInitViper() {
	viper.Set(config.DBApplicationMySQLUserKey, testDASMySQLUser)
	viper.Set(config.DBApplicationMySQLPassKey, testDASMySQLPass)
	viper.Set(config.PrivilegeEnabledKey, true)
}

func TestRepository_All(t *testing.T) {
	TestRepository_Execute(t)
	TestRepository_GetMySQLServerClusterIDListByLoginName(t)
	TestRepository_GetMySQLClusterIDByMySQLServerID(t)
	TestRepository_GetMySQLClusterIDByHostInfo(t)
	TestRepository_GetMySQLClusterIDByDBID(t)
}

func TestRepository_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := "select 1;"
	result, err := testRepository.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestRepository_GetMySQLServerClusterIDListByLoginName(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterIDList, err := testRepository.GetMySQLClusterIDListByLoginName(testLoginName)
	asst.Nil(err, "test GetMySQLClusterIDListByLoginName() failed")
	asst.Equal(testMySQLClusterID, mysqlClusterIDList[constant.ZeroInt], "test GetMySQLClusterIDListByLoginName() failed")
}

func TestRepository_GetMySQLClusterIDByMySQLServerID(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterID, err := testRepository.GetMySQLClusterIDByMySQLServerID(testMySQLServerID)
	asst.Nil(err, "test GetMySQLClusterIDByMySQLServerID() failed")
	asst.Equal(testMySQLClusterID, mysqlClusterID, "test GetMySQLClusterIDByMySQLServerID() failed")
}

func TestRepository_GetMySQLClusterIDByHostInfo(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterID, err := testRepository.GetMySQLClusterIDByHostInfo(testMySQLHostIP, testMYSQLPortNum)
	asst.Nil(err, "test GetMySQLClusterIDByHostInfo() failed")
	asst.Equal(testMySQLClusterID, mysqlClusterID, "test GetMySQLClusterIDByHostInfo() failed")

}

func TestRepository_GetMySQLClusterIDByDBID(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterID, err := testRepository.GetMySQLClusterIDByDBID(testDBID)
	asst.Nil(err, "test GetMySQLClusterIDByDBID() failed")
	asst.Equal(testMySQLClusterID, mysqlClusterID, "test GetMySQLClusterIDByDBID() failed")
}
