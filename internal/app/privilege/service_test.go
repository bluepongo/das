package privilege

import (
	"os"
	"testing"

	"github.com/jinzhu/now"
	"github.com/romberli/das/config"
	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/app/metadata"
	depmeta "github.com/romberli/das/internal/dependency/metadata"
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

	testUserID                   = 1
	testUserUserName             = "zhangsan"
	testUserDepartmentName       = "arch"
	testUserEmployeeID           = "100001"
	testUserAccountName          = "zs001"
	testUserEmail                = "allinemailtest@163.com"
	testUserTelephone            = "01012345678"
	testUserMobile               = "13012345678"
	testUserRole                 = 3
	testUserDelFlag              = 0
	testUserCreateTimeString     = "2021-01-21 10:00:00.000000"
	testUserLastUpdateTimeString = "2021-01-21 13:00:00.000000"

	testMySQLServerID    = 1
	testMySQLHostIP      = "192.168.137.11"
	testMYSQLPortNum     = 3306
	testDBID             = 1
	testMySQLClusterID   = 1
	testMySQLClusterType = 1
	testDBName           = "das"
)

var (
	testUser    depmeta.User
	testService privilege.Service
)

func init() {
	testInitDASMySQLPool()
	testInitViper()
	testUser = testInitNewUserInfo()
	testService = NewServiceWithDefault(testUser)
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

func testInitNewUserInfo() depmeta.User {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(testUserCreateTimeString)
	lastUpdateTime, _ := now.Parse(testUserLastUpdateTimeString)
	return metadata.NewUserInfoWithGlobal(
		testUserID,
		testUserUserName,
		testUserDepartmentName,
		testUserEmployeeID,
		testUserAccountName,
		testUserEmail,
		testUserTelephone,
		testUserMobile,
		testUserRole,
		testUserDelFlag,
		createTime,
		lastUpdateTime,
	)
}

func TestUserServiceAll(t *testing.T) {
	TestService_GetUser(t)
	TestService_CheckMySQLServerByID(t)
	TestService_CheckMySQLServerByHostInfo(t)
	TestService_CheckDBByID(t)
	TestService_CheckDBByNameAndClusterInfo(t)
	TestService_CheckDBByNameAndHostInfo(t)
}

// GetUser returns the user
func TestService_GetUser(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testUser.Identity(), testService.GetUser().Identity(), "test GetUser() failed")
}

func TestService_CheckMySQLServerByID(t *testing.T) {
	asst := assert.New(t)

	err := testService.CheckMySQLServerByID(testMySQLServerID)
	asst.Nil(err, "test CheckMySQLServerByID() failed")
}

func TestService_CheckMySQLServerByHostInfo(t *testing.T) {
	asst := assert.New(t)

	err := testService.CheckMySQLServerByHostInfo(testMySQLHostIP, testMYSQLPortNum)
	asst.Nil(err, "test CheckMySQLServerByHostInfo() failed")
}

func TestService_CheckDBByID(t *testing.T) {
	asst := assert.New(t)

	err := testService.CheckDBByID(testDBID)
	asst.Nil(err, "test CheckDBByID() failed")
}

func TestService_CheckDBByNameAndClusterInfo(t *testing.T) {
	asst := assert.New(t)

	err := testService.CheckDBByNameAndClusterInfo(testDBName, testMySQLClusterID, testMySQLClusterType)
	asst.Nil(err, "test CheckDBByNameAndClusterInfo() failed")
}

func TestService_CheckDBByNameAndHostInfo(t *testing.T) {
	asst := assert.New(t)

	err := testService.CheckDBByNameAndHostInfo(testDBName, testMySQLHostIP, testMYSQLPortNum)
	asst.Nil(err, "test CheckMySQLServerByHostInfo() failed")
}
