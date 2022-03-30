package privilege

import (
	"testing"

	"github.com/romberli/das/internal/dependency/privilege"
	"github.com/stretchr/testify/assert"
)

var testService privilege.Service

func init() {
	testInitDASMySQLPool()
	testInitViper()

	testService = NewServiceWithDefault(testLoginName)
}

func TestService_All(t *testing.T) {
	TestService_GetLoginName(t)
	TestService_CheckMySQLServerByID(t)
	TestService_CheckMySQLServerByHostInfo(t)
	TestService_CheckDBByID(t)
}

// GetUser returns the user
func TestService_GetLoginName(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testLoginName, testService.GetLoginName(), "test TestService_GetLoginName() failed")
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
